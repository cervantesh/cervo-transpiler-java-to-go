package pipeline

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoldenHello(t *testing.T) {
	root := workspaceRoot(t)
	sourcePath := filepath.Join(root, "modern-tests", "fixtures", "hello", "Hello.java")
	wantPath := filepath.Join(root, "modern-tests", "expected", "hello", "Hello.go")

	source, err := os.ReadFile(sourcePath)
	if err != nil {
		t.Fatalf("read source: %v", err)
	}
	want, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatalf("read expected: %v", err)
	}

	result := TranspileSource(sourcePath, string(source))
	if len(result.Diagnostics) != 0 {
		t.Fatalf("diagnostics: %#v", result.Diagnostics)
	}

	got := strings.TrimSpace(normalizeNewlines(string(result.Code)))
	expected := strings.TrimSpace(normalizeNewlines(string(want)))
	if got != expected {
		t.Fatalf("golden mismatch\n--- got ---\n%s\n--- want ---\n%s", got, expected)
	}
}

func normalizeNewlines(s string) string {
	return strings.ReplaceAll(s, "\r\n", "\n")
}

func workspaceRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find workspace root")
		}
		dir = parent
	}
}
