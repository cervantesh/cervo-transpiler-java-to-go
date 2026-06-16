package aisidecar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
)

const advisoryHeader = "# AI Advisory Migration Review\n\n> Advisory output only. Deterministic reports and generated Go remain the source of truth.\n\n"

type Options struct {
	ReportPath   string
	OutPath      string
	Provider     string
	SnippetPaths []string
}

type Result struct {
	Provider     string
	Output       []byte
	ParsedReport bool
}

type Input struct {
	ReportPath   string               `json:"reportPath"`
	ReportFormat string               `json:"reportFormat"`
	Project      *javaproject.Project `json:"project,omitempty"`
	ReportText   string               `json:"reportText,omitempty"`
	Snippets     []Snippet            `json:"snippets,omitempty"`
}

type Snippet struct {
	Path     string `json:"path"`
	Language string `json:"language"`
	Body     string `json:"body"`
}

type Provider interface {
	Name() string
	Explain(Input) ([]byte, error)
}

func Explain(opts Options) (Result, error) {
	if opts.ReportPath == "" {
		return Result{}, fmt.Errorf("report path is required")
	}
	if opts.OutPath == "" {
		return Result{}, fmt.Errorf("out path is required")
	}

	input, parsed, err := loadInput(opts.ReportPath, opts.SnippetPaths)
	if err != nil {
		return Result{}, err
	}
	provider, err := selectProvider(opts.Provider)
	if err != nil {
		return Result{}, err
	}
	output, err := provider.Explain(input)
	if err != nil {
		return Result{}, err
	}
	if !bytes.HasPrefix(output, []byte(advisoryHeader)) {
		output = append([]byte(advisoryHeader), output...)
	}
	if err := os.MkdirAll(filepath.Dir(opts.OutPath), 0755); err != nil && filepath.Dir(opts.OutPath) != "." {
		return Result{}, err
	}
	if err := os.WriteFile(opts.OutPath, output, 0644); err != nil {
		return Result{}, err
	}
	return Result{Provider: provider.Name(), Output: output, ParsedReport: parsed}, nil
}

func loadInput(reportPath string, snippetPaths []string) (Input, bool, error) {
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return Input{}, false, err
	}
	input := Input{ReportPath: reportPath, ReportFormat: "markdown", ReportText: string(data)}
	var project javaproject.Project
	if err := json.Unmarshal(data, &project); err == nil && project.Root != "" {
		input.Project = &project
		input.ReportFormat = "json"
		input.ReportText = ""
	}
	for _, snippetPath := range snippetPaths {
		snippet, err := readSnippet(snippetPath)
		if err != nil {
			return Input{}, false, err
		}
		input.Snippets = append(input.Snippets, snippet)
	}
	return input, input.Project != nil, nil
}

func readSnippet(path string) (Snippet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Snippet{}, err
	}
	language := "text"
	if strings.EqualFold(filepath.Ext(path), ".go") {
		language = "go"
	}
	return Snippet{Path: path, Language: language, Body: string(data)}, nil
}

func selectProvider(name string) (Provider, error) {
	if name == "" {
		name = os.Getenv("J2GO_AI_PROVIDER")
	}
	if name == "" {
		name = "canned"
	}
	switch strings.ToLower(name) {
	case "canned", "offline":
		return cannedProvider{}, nil
	case "external":
		command := os.Getenv("J2GO_AI_COMMAND")
		if command == "" {
			return nil, fmt.Errorf("external AI provider requires J2GO_AI_COMMAND")
		}
		return externalProvider{command: command}, nil
	default:
		return nil, fmt.Errorf("unsupported AI provider %q", name)
	}
}

type cannedProvider struct{}

func (cannedProvider) Name() string {
	return "canned"
}

func (cannedProvider) Explain(input Input) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(advisoryHeader)
	fmt.Fprintf(&buffer, "Provider: `canned`\n\n")
	fmt.Fprintf(&buffer, "Source report: `%s`\n\n", input.ReportPath)
	if input.Project == nil {
		fmt.Fprintf(&buffer, "The input report was not a structured JSON report, so this advisory review is limited to generic migration guidance.\n\n")
		fmt.Fprintf(&buffer, "## Manual Migration Notes\n\n")
		fmt.Fprintf(&buffer, "- Regenerate the report with `j2go report <project> --format json` for file-level risk and diagnostic explanations.\n")
		fmt.Fprintf(&buffer, "- Keep generated code deterministic; treat this advisory output as review material only.\n")
		writeSnippetSummary(&buffer, input.Snippets)
		return buffer.Bytes(), nil
	}

	project := input.Project
	fmt.Fprintf(&buffer, "## Project Summary\n\n")
	fmt.Fprintf(&buffer, "- Java files: %d\n", project.Summary.JavaFiles)
	fmt.Fprintf(&buffer, "- Packages: %d\n", project.Summary.Packages)
	fmt.Fprintf(&buffer, "- Classes: %d\n", project.Summary.Classes)
	fmt.Fprintf(&buffer, "- Unsupported features: %d\n", project.Summary.Unsupported)
	fmt.Fprintf(&buffer, "- Diagnostics: %d\n\n", project.Summary.Diagnostics)

	writeHighestRiskFiles(&buffer, project.Files)
	writeManualNotes(&buffer, project.Files)
	writeReviewNotes(&buffer, input.Snippets)
	writeGoldenTestIdeas(&buffer, project.Files)
	writeSnippetSummary(&buffer, input.Snippets)
	return buffer.Bytes(), nil
}

func writeHighestRiskFiles(buffer *bytes.Buffer, files []javaproject.File) {
	fmt.Fprintf(buffer, "## Highest-Risk Files\n\n")
	sorted := append([]javaproject.File(nil), files...)
	rank := map[string]int{"high": 0, "medium": 1, "low": 2}
	sort.Slice(sorted, func(i, j int) bool {
		if rank[sorted[i].Risk] != rank[sorted[j].Risk] {
			return rank[sorted[i].Risk] < rank[sorted[j].Risk]
		}
		return sorted[i].RelativePath < sorted[j].RelativePath
	})
	for _, file := range sorted {
		if file.Risk == "low" && len(sorted) > 5 {
			continue
		}
		fmt.Fprintf(buffer, "- `%s`: %s risk", file.RelativePath, file.Risk)
		if len(file.Unsupported) > 0 {
			fmt.Fprintf(buffer, ", %d unsupported", len(file.Unsupported))
		}
		if len(file.Diagnostics) > 0 {
			fmt.Fprintf(buffer, ", %d diagnostics", len(file.Diagnostics))
		}
		fmt.Fprintf(buffer, "\n")
	}
	if len(sorted) == 0 {
		fmt.Fprintf(buffer, "None.\n")
	}
	fmt.Fprintf(buffer, "\n")
}

func writeManualNotes(buffer *bytes.Buffer, files []javaproject.File) {
	fmt.Fprintf(buffer, "## Manual Migration Notes\n\n")
	notes := map[string]string{}
	for _, file := range files {
		for _, unsupported := range file.Unsupported {
			key := unsupported.Name
			if unsupported.Code != "" {
				key = unsupported.Code + " " + key
			}
			if unsupported.Recommendation != "" {
				notes[key] = unsupported.Recommendation
			} else {
				notes[key] = "Review manually before relying on generated Go."
			}
		}
		for _, diagnostic := range file.Diagnostics {
			key := diagnostic.Code
			if key == "" {
				key = diagnostic.Message
			}
			notes[key] = diagnostic.Message
		}
	}
	if len(notes) == 0 {
		fmt.Fprintf(buffer, "- No manual blockers were reported by the deterministic JSON report.\n")
		fmt.Fprintf(buffer, "- Still review package names, exported API shape, and generated tests before using the migrated library.\n\n")
		return
	}
	keys := make([]string, 0, len(notes))
	for key := range notes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintf(buffer, "- `%s`: %s\n", key, notes[key])
	}
	fmt.Fprintf(buffer, "\n")
}

func writeReviewNotes(buffer *bytes.Buffer, snippets []Snippet) {
	fmt.Fprintf(buffer, "## Review Notes\n\n")
	fmt.Fprintf(buffer, "- Do not copy advisory text into generated output without human review.\n")
	fmt.Fprintf(buffer, "- Keep deterministic diagnostics as the authority for supported and unsupported constructs.\n")
	if len(snippets) > 0 {
		fmt.Fprintf(buffer, "- Review %d generated-code snippets for package naming, receiver naming, and exported API shape.\n", len(snippets))
	}
	fmt.Fprintf(buffer, "\n")
}

func writeGoldenTestIdeas(buffer *bytes.Buffer, files []javaproject.File) {
	fmt.Fprintf(buffer, "## Proposed Golden-Test Ideas\n\n")
	if len(files) == 0 {
		fmt.Fprintf(buffer, "- Add a fixture that proves empty projects produce a stable report.\n\n")
		return
	}
	for _, file := range files {
		if file.Risk == "high" || len(file.Unsupported) > 0 || len(file.Diagnostics) > 0 {
			fmt.Fprintf(buffer, "- Add a negative fixture for `%s` that asserts the current diagnostics stay stable.\n", file.RelativePath)
		}
	}
	fmt.Fprintf(buffer, "- Add a positive fixture for the lowest-risk package and run `go test ./...` on generated output.\n\n")
}

func writeSnippetSummary(buffer *bytes.Buffer, snippets []Snippet) {
	if len(snippets) == 0 {
		return
	}
	fmt.Fprintf(buffer, "## Generated-Code Snippets Reviewed\n\n")
	for _, snippet := range snippets {
		lines := 0
		if snippet.Body != "" {
			lines = strings.Count(snippet.Body, "\n") + 1
		}
		fmt.Fprintf(buffer, "- `%s` (%s, %d lines)\n", snippet.Path, snippet.Language, lines)
	}
	fmt.Fprintf(buffer, "\n")
}

type externalProvider struct {
	command string
}

func (p externalProvider) Name() string {
	return "external"
}

func (p externalProvider) Explain(input Input) ([]byte, error) {
	payload, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return nil, err
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-NoProfile", "-Command", p.command)
	} else {
		cmd = exec.Command("sh", "-c", p.command)
	}
	cmd.Stdin = bytes.NewReader(payload)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("external AI provider failed: %w\n%s", err, string(output))
	}
	var buffer bytes.Buffer
	buffer.WriteString(advisoryHeader)
	fmt.Fprintf(&buffer, "Provider: `external`\n\n")
	buffer.Write(output)
	if !bytes.HasSuffix(output, []byte("\n")) {
		buffer.WriteByte('\n')
	}
	return buffer.Bytes(), nil
}
