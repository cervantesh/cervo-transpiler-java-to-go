package migrate

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/emitgo"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/lower"
	javaparser "github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/report"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/semantic"
)

type Options struct {
	OutDir     string
	ReportPath string
	DryRun     bool
	ModulePath string
}

type Summary struct {
	JavaFiles   int
	Generated   int
	Skipped     int
	Diagnostics int
	DryRun      bool
	ReportPath  string
	OutDir      string
}

type Result struct {
	Summary     Summary
	Diagnostics []error
}

func Run(projectPath string, opts Options) (Result, error) {
	if opts.OutDir == "" {
		return Result{}, fmt.Errorf("out dir is required")
	}
	if opts.ModulePath == "" {
		opts.ModulePath = "migrated"
	}
	project, err := javaproject.Scan(projectPath)
	if err != nil {
		return Result{}, err
	}
	result := Result{Summary: Summary{
		JavaFiles:  len(project.Files),
		DryRun:     opts.DryRun,
		ReportPath: opts.ReportPath,
		OutDir:     opts.OutDir,
	}}

	if !opts.DryRun {
		if err := os.MkdirAll(opts.OutDir, 0755); err != nil {
			return Result{}, err
		}
		if err := os.WriteFile(filepath.Join(opts.OutDir, "go.mod"), []byte("module "+opts.ModulePath+"\n\ngo 1.25\n"), 0644); err != nil {
			return Result{}, err
		}
	}

	files := append([]javaproject.File(nil), project.Files...)
	sort.Slice(files, func(i, j int) bool {
		return files[i].RelativePath < files[j].RelativePath
	})
	seenOutputs := map[string]bool{}
	for _, file := range files {
		code, diagnostics := transpileProjectFile(file)
		if len(diagnostics) > 0 {
			result.Summary.Skipped++
			result.Summary.Diagnostics += len(diagnostics)
			result.Diagnostics = append(result.Diagnostics, diagnostics...)
			continue
		}

		outPath := outputPath(opts.OutDir, file)
		if seenOutputs[outPath] {
			result.Summary.Skipped++
			result.Summary.Diagnostics++
			result.Diagnostics = append(result.Diagnostics, fmt.Errorf("%s: output collision: %s", file.Path, outPath))
			continue
		}
		seenOutputs[outPath] = true
		result.Summary.Generated++
		if opts.DryRun {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return Result{}, err
		}
		if err := os.WriteFile(outPath, code, 0644); err != nil {
			return Result{}, err
		}
	}
	if opts.ReportPath != "" {
		if err := writeReport(opts.ReportPath, renderMigrationReport(project, result)); err != nil {
			return Result{}, err
		}
	}
	return result, nil
}

func transpileProjectFile(file javaproject.File) ([]byte, []error) {
	source, err := os.ReadFile(file.Path)
	if err != nil {
		return nil, []error{err}
	}
	tree, syntax := javaparser.ParseSource(file.Path, stripProjectMetadata(string(source)))
	if len(syntax) != 0 {
		out := make([]error, 0, len(syntax))
		for _, diagnostic := range syntax {
			out = append(out, diagnostic)
		}
		return nil, out
	}
	irFile, lowerDiagnostics := lower.Lower(file.Path, tree)
	if len(lowerDiagnostics) != 0 {
		out := make([]error, 0, len(lowerDiagnostics))
		for _, diagnostic := range lowerDiagnostics {
			out = append(out, diagnostic)
		}
		return nil, out
	}
	irFile.PackageName = goPackageName(file.PackageName)
	semanticDiagnostics := semantic.AnalyzeIR(irFile)
	if len(semanticDiagnostics) != 0 {
		out := make([]error, 0, len(semanticDiagnostics))
		for _, diagnostic := range semanticDiagnostics {
			out = append(out, diagnostic)
		}
		return nil, out
	}
	code, err := emitgo.Emit(irFile)
	if err != nil {
		return nil, []error{err}
	}
	if _, err := parser.ParseFile(token.NewFileSet(), outputPath("", file), code, parser.AllErrors); err != nil {
		return nil, []error{fmt.Errorf("%s: generated Go did not parse: %w", file.Path, err)}
	}
	return code, nil
}

func stripProjectMetadata(source string) string {
	lines := strings.Split(source, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "package ") || strings.HasPrefix(trimmed, "import ") {
			lines[i] = ""
		}
	}
	return strings.Join(lines, "\n")
}

func outputPath(outDir string, file javaproject.File) string {
	packagePath := strings.ReplaceAll(file.PackageName, ".", "/")
	if packagePath == "" {
		packagePath = "."
	}
	base := strings.TrimSuffix(filepath.Base(file.Path), filepath.Ext(file.Path)) + ".go"
	return filepath.Join(outDir, filepath.FromSlash(packagePath), base)
}

func goPackageName(javaPackage string) string {
	if javaPackage == "" {
		return "main"
	}
	parts := strings.Split(javaPackage, ".")
	name := parts[len(parts)-1]
	if name == "" {
		return "main"
	}
	return sanitizeIdentifier(name)
}

func sanitizeIdentifier(name string) string {
	var buffer bytes.Buffer
	for i, r := range name {
		if r == '_' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || i > 0 && r >= '0' && r <= '9' {
			buffer.WriteRune(r)
			continue
		}
		buffer.WriteRune('_')
	}
	out := buffer.String()
	if out == "" || out[0] >= '0' && out[0] <= '9' {
		return "pkg_" + out
	}
	return out
}

func renderMigrationReport(project javaproject.Project, result Result) []byte {
	var buffer bytes.Buffer
	buffer.Write(report.Markdown(project))
	if !bytes.HasSuffix(buffer.Bytes(), []byte("\n")) {
		buffer.WriteByte('\n')
	}
	fmt.Fprintf(&buffer, "\n## Migration Execution Summary\n\n")
	fmt.Fprintf(&buffer, "| Metric | Count |\n")
	fmt.Fprintf(&buffer, "| --- | ---: |\n")
	fmt.Fprintf(&buffer, "| Java files | %d |\n", result.Summary.JavaFiles)
	fmt.Fprintf(&buffer, "| Generated files | %d |\n", result.Summary.Generated)
	fmt.Fprintf(&buffer, "| Skipped files | %d |\n", result.Summary.Skipped)
	fmt.Fprintf(&buffer, "| Diagnostics | %d |\n", result.Summary.Diagnostics)
	fmt.Fprintf(&buffer, "| Dry run | %t |\n\n", result.Summary.DryRun)
	fmt.Fprintf(&buffer, "### Migration Diagnostics\n\n")
	if len(result.Diagnostics) == 0 {
		fmt.Fprintf(&buffer, "None.\n")
		return buffer.Bytes()
	}
	for _, diagnostic := range result.Diagnostics {
		fmt.Fprintf(&buffer, "- `%s`\n", strings.ReplaceAll(diagnostic.Error(), "\n", " "))
	}
	return buffer.Bytes()
}

func writeReport(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil && filepath.Dir(path) != "." {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
