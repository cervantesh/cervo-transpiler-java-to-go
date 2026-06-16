package report

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
)

func TestJSONReportIsDeterministicAndParseable(t *testing.T) {
	projectRoot := filepath.Join("..", "javaproject", "testdata", "pure-java-lib")
	project, err := javaproject.Scan(projectRoot)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	first, err := JSON(project)
	if err != nil {
		t.Fatalf("JSON failed: %v", err)
	}
	second, err := JSON(project)
	if err != nil {
		t.Fatalf("JSON failed: %v", err)
	}
	if string(first) != string(second) {
		t.Fatal("expected deterministic JSON output")
	}

	var decoded javaproject.Project
	if err := json.Unmarshal(first, &decoded); err != nil {
		t.Fatalf("report JSON was not parseable: %v", err)
	}
	if decoded.Summary.JavaFiles != 3 {
		t.Fatalf("expected Java file count in report, got %#v", decoded.Summary)
	}

	gotSummary, err := json.MarshalIndent(decoded.Summary, "", "  ")
	if err != nil {
		t.Fatalf("marshal summary: %v", err)
	}
	gotSummary = append(gotSummary, '\n')
	wantSummary, err := os.ReadFile(filepath.Join("testdata", "expected", "pure-java-lib-summary.json"))
	if err != nil {
		t.Fatalf("read summary golden: %v", err)
	}
	if !bytes.Equal(gotSummary, normalizeNewlines(wantSummary)) {
		t.Fatalf("summary golden mismatch\n--- got ---\n%s\n--- want ---\n%s", gotSummary, wantSummary)
	}
}

func TestMarkdownReportIncludesRiskAndMigrationOrder(t *testing.T) {
	projectRoot := filepath.Join("..", "javaproject", "testdata", "pure-java-lib")
	project, err := javaproject.Scan(projectRoot)
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	root, err := filepath.Abs(projectRoot)
	if err != nil {
		t.Fatalf("abs project root: %v", err)
	}
	text := normalizeRoot(string(Markdown(project)), root)
	want, err := os.ReadFile(filepath.Join("testdata", "expected", "pure-java-lib.md"))
	if err != nil {
		t.Fatalf("read markdown golden: %v", err)
	}
	if strings.TrimSpace(text) != strings.TrimSpace(string(normalizeNewlines(want))) {
		t.Fatalf("markdown golden mismatch\n--- got ---\n%s\n--- want ---\n%s", text, string(want))
	}

	for _, want := range []string{
		"# Java-To-Go Migration Report",
		"## Unsupported Feature Counts",
		"## Recommended Migration Order",
		"`src/main/java/com/example/Calculator.java`",
		"| class fields | 1 |",
		"None.",
		"`instance int add(2 params)`",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in markdown report:\n%s", want, text)
		}
	}
}

func normalizeRoot(text string, root string) string {
	text = strings.ReplaceAll(text, filepath.ToSlash(root), "<PROJECT_ROOT>")
	return strings.ReplaceAll(text, "\r\n", "\n")
}

func normalizeNewlines(text []byte) []byte {
	return []byte(strings.ReplaceAll(string(text), "\r\n", "\n"))
}
