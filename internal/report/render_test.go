package report

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
)

func TestJSONReportIsDeterministicAndParseable(t *testing.T) {
	project, err := javaproject.Scan(filepath.Join("..", "javaproject", "testdata", "pure-java-lib"))
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
}

func TestMarkdownReportIncludesRiskAndMigrationOrder(t *testing.T) {
	project, err := javaproject.Scan(filepath.Join("..", "javaproject", "testdata", "pure-java-lib"))
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	text := string(Markdown(project))
	for _, want := range []string{
		"# Java-To-Go Migration Report",
		"## Unsupported Feature Counts",
		"## Recommended Migration Order",
		"`src/main/java/com/example/Calculator.java`",
		"Add struct field lowering before transpiling Java fields.",
		"`instance int add(2 params)`",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in markdown report:\n%s", want, text)
		}
	}
}
