package aisidecar

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/report"
)

func TestExplainWithCannedProviderWritesAdvisoryForJSONReport(t *testing.T) {
	reportPath := filepath.Join(t.TempDir(), "migration-report.json")
	outPath := filepath.Join(t.TempDir(), "ai-review.md")
	snippetPath := filepath.Join(t.TempDir(), "generated.go")
	writeFile(t, snippetPath, "package sample\n\nfunc Answer() int { return 42 }\n")
	project := javaproject.Project{
		Root: "sample-project",
		Files: []javaproject.File{{
			RelativePath: "src/main/java/com/acme/Legacy.java",
			Risk:         "high",
			Unsupported: []javaproject.Feature{{
				Code:           "JTG1007",
				Name:           "try/catch exceptions",
				Recommendation: "Design explicit Go error returns before transpiling exceptions.",
			}},
		}},
		Summary: javaproject.Summary{JavaFiles: 1, Packages: 1, Classes: 1, Unsupported: 1},
	}
	data, err := report.JSON(project)
	if err != nil {
		t.Fatalf("render report: %v", err)
	}
	writeFile(t, reportPath, string(data))

	result, err := Explain(Options{ReportPath: reportPath, OutPath: outPath, Provider: "canned", SnippetPaths: []string{snippetPath}})
	if err != nil {
		t.Fatalf("Explain failed: %v", err)
	}
	if result.Provider != "canned" || !result.ParsedReport {
		t.Fatalf("unexpected result: %#v", result)
	}
	text := readFile(t, outPath)
	for _, want := range []string{
		"# AI Advisory Migration Review",
		"Advisory output only",
		"Provider: `canned`",
		"src/main/java/com/acme/Legacy.java",
		"JTG1007 try/catch exceptions",
		"Generated-Code Snippets Reviewed",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in advisory output:\n%s", want, text)
		}
	}
}

func TestExplainCannedMarkdownFallbackIsAdvisory(t *testing.T) {
	reportPath := filepath.Join(t.TempDir(), "migration-report.md")
	outPath := filepath.Join(t.TempDir(), "ai-review.md")
	writeFile(t, reportPath, "# Java-To-Go Migration Report\n\nMarkdown only.\n")

	result, err := Explain(Options{ReportPath: reportPath, OutPath: outPath, Provider: "offline"})
	if err != nil {
		t.Fatalf("Explain failed: %v", err)
	}
	if result.Provider != "canned" || result.ParsedReport {
		t.Fatalf("unexpected result: %#v", result)
	}
	text := readFile(t, outPath)
	for _, want := range []string{"Advisory output only", "not a structured JSON report", "j2go report <project> --format json"} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in advisory output:\n%s", want, text)
		}
	}
}

func TestExplainExternalProviderIsOptIn(t *testing.T) {
	reportPath := filepath.Join(t.TempDir(), "migration-report.md")
	outPath := filepath.Join(t.TempDir(), "ai-review.md")
	writeFile(t, reportPath, "# report\n")
	t.Setenv("J2GO_AI_COMMAND", "")

	if _, err := Explain(Options{ReportPath: reportPath, OutPath: outPath, Provider: "external"}); err == nil {
		t.Fatalf("expected external provider to require J2GO_AI_COMMAND")
	}

	if _, err := Explain(Options{ReportPath: reportPath, OutPath: outPath, Provider: "canned"}); err != nil {
		t.Fatalf("expected canned provider to work without external command: %v", err)
	}
}

func TestExplainExternalProviderWrapsCommandOutputAsAdvisory(t *testing.T) {
	reportPath := filepath.Join(t.TempDir(), "migration-report.md")
	outPath := filepath.Join(t.TempDir(), "ai-review.md")
	writeFile(t, reportPath, "# report\n")
	command := "printf 'external ok\\n'"
	if runtime.GOOS == "windows" {
		command = "Write-Output 'external ok'"
	}
	t.Setenv("J2GO_AI_COMMAND", command)

	result, err := Explain(Options{ReportPath: reportPath, OutPath: outPath, Provider: "external"})
	if err != nil {
		t.Fatalf("Explain failed: %v", err)
	}
	if result.Provider != "external" {
		t.Fatalf("unexpected provider: %#v", result)
	}
	text := readFile(t, outPath)
	for _, want := range []string{"# AI Advisory Migration Review", "Provider: `external`", "external ok"} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in advisory output:\n%s", want, text)
		}
	}
}

func writeFile(t *testing.T, path string, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}
