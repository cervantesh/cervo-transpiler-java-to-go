package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunVersion(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"--version"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d; stderr=%s", code, stderr.String())
	}
	if got := strings.TrimSpace(stdout.String()); got != "j2go dev" {
		t.Fatalf("expected version output, got %q", got)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}
}

func TestRunRequiresOutDir(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"Hello.java"}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("expected usage exit code 2, got %d", code)
	}
	if !strings.Contains(stderr.String(), "-out is required") {
		t.Fatalf("expected missing -out error, got %q", stderr.String())
	}
}

func TestRunRequiresInput(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"-out", t.TempDir()}, &stdout, &stderr)
	if code != 2 {
		t.Fatalf("expected usage exit code 2, got %d", code)
	}
	if !strings.Contains(stderr.String(), "at least one Java file or directory is required") {
		t.Fatalf("expected missing input error, got %q", stderr.String())
	}
}

func TestRunTranspilesFile(t *testing.T) {
	tmp := t.TempDir()
	input := filepath.Join(tmp, "Hello.java")
	outDir := filepath.Join(tmp, "out")
	source := `public class Hello {
  public static void main(String[] args) {
    System.out.println("hello");
  }
}`
	if err := os.WriteFile(input, []byte(source), 0644); err != nil {
		t.Fatalf("write source: %v", err)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	code := run([]string{"-out", outDir, input}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected success, got %d; stderr=%s", code, stderr.String())
	}

	generated, err := os.ReadFile(filepath.Join(outDir, "Hello.go"))
	if err != nil {
		t.Fatalf("read generated file: %v", err)
	}
	if !strings.Contains(string(generated), "fmt.Println(\"hello\")") {
		t.Fatalf("expected generated println, got:\n%s", string(generated))
	}
}
