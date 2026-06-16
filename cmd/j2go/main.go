package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/pipeline"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/report"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/version"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) > 0 {
		switch args[0] {
		case "--version", "-version", "version":
			fmt.Fprintf(stdout, "j2go %s\n", version.Version)
			return 0
		case "scan":
			return runScan(args[1:], stdout, stderr)
		case "report":
			return runReport(args[1:], stderr)
		case "transpile":
			return runTranspile(args[1:], stderr)
		}
	}
	return runTranspile(args, stderr)
}

func runTranspile(args []string, stderr io.Writer) int {
	fs := flag.NewFlagSet("j2go transpile", flag.ContinueOnError)
	fs.SetOutput(stderr)
	outDir := fs.String("out", "", "directory for generated Go files")
	if err := fs.Parse(reorderValueFlags(args, "out")); err != nil {
		return 2
	}
	if *outDir == "" {
		fmt.Fprintln(stderr, "-out is required")
		return 2
	}
	inputs := fs.Args()
	if len(inputs) == 0 {
		fmt.Fprintln(stderr, "at least one Java file or directory is required")
		return 2
	}
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	for _, input := range inputs {
		if err := transpilePath(input, *outDir); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}
	return 0
}

func runScan(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := flag.NewFlagSet("j2go scan", flag.ContinueOnError)
	fs.SetOutput(stderr)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	inputs := fs.Args()
	if len(inputs) != 1 {
		fmt.Fprintln(stderr, "scan requires exactly one Java project path")
		return 2
	}
	project, err := javaproject.Scan(inputs[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	_, _ = stdout.Write(report.ScanText(project))
	return 0
}

func runReport(args []string, stderr io.Writer) int {
	fs := flag.NewFlagSet("j2go report", flag.ContinueOnError)
	fs.SetOutput(stderr)
	format := fs.String("format", "markdown", "report format: json or markdown")
	outPath := fs.String("out", "", "report output file")
	if err := fs.Parse(reorderValueFlags(args, "format", "out")); err != nil {
		return 2
	}
	inputs := fs.Args()
	if len(inputs) != 1 {
		fmt.Fprintln(stderr, "report requires exactly one Java project path")
		return 2
	}
	if *outPath == "" {
		fmt.Fprintln(stderr, "-out is required")
		return 2
	}

	project, err := javaproject.Scan(inputs[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	var rendered []byte
	switch *format {
	case "json":
		rendered, err = report.JSON(project)
	case "markdown", "md":
		rendered = report.Markdown(project)
	default:
		fmt.Fprintf(stderr, "unsupported report format %q\n", *format)
		return 2
	}
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if err := os.MkdirAll(filepath.Dir(*outPath), 0755); err != nil && filepath.Dir(*outPath) != "." {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if err := os.WriteFile(*outPath, rendered, 0644); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}

func reorderValueFlags(args []string, names ...string) []string {
	valueFlags := map[string]bool{}
	for _, name := range names {
		valueFlags[name] = true
	}

	flags := []string{}
	positionals := []string{}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") || arg == "-" {
			positionals = append(positionals, arg)
			continue
		}

		flags = append(flags, arg)
		name := strings.TrimLeft(strings.SplitN(arg, "=", 2)[0], "-")
		if strings.Contains(arg, "=") || !valueFlags[name] || i+1 >= len(args) {
			continue
		}
		i++
		flags = append(flags, args[i])
	}
	return append(flags, positionals...)
}

func transpilePath(input string, outDir string) error {
	info, err := os.Stat(input)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return filepath.WalkDir(input, func(path string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || !strings.HasSuffix(path, ".java") {
				return nil
			}
			return transpileFile(path, outDir)
		})
	}
	return transpileFile(input, outDir)
}

func transpileFile(input string, outDir string) error {
	source, err := os.ReadFile(input)
	if err != nil {
		return err
	}
	result := pipeline.TranspileSource(input, string(source))
	if len(result.Diagnostics) > 0 {
		return result.Diagnostics[0]
	}
	base := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input)) + ".go"
	return os.WriteFile(filepath.Join(outDir, base), result.Code, 0644)
}
