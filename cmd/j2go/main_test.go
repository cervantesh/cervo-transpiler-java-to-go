package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
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

func TestRunTranspileSubcommandTranspilesFile(t *testing.T) {
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
	code := run([]string{"transpile", "-out", outDir, input}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected success, got %d; stderr=%s", code, stderr.String())
	}
	if _, err := os.Stat(filepath.Join(outDir, "Hello.go")); err != nil {
		t.Fatalf("expected generated file: %v", err)
	}
}

func TestRunScanProject(t *testing.T) {
	projectRoot := writePureJavaProject(t)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"scan", projectRoot}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected scan success, got %d; stderr=%s", code, stderr.String())
	}
	text := stdout.String()
	for _, want := range []string{"Java files: 2", "Package list:", "com.example", "Internal dependencies:", "class com.example.Calculator"} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in scan output:\n%s", want, text)
		}
	}
}

func TestRunReportJSON(t *testing.T) {
	projectRoot := writePureJavaProject(t)
	outPath := filepath.Join(t.TempDir(), "migration-report.json")
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"report", projectRoot, "--format", "json", "--out", outPath}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected report success, got %d; stderr=%s", code, stderr.String())
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	var project javaproject.Project
	if err := json.Unmarshal(data, &project); err != nil {
		t.Fatalf("report JSON did not parse: %v", err)
	}
	if project.Summary.JavaFiles != 2 {
		t.Fatalf("expected two Java files in report, got %#v", project.Summary)
	}
}

func TestRunReportMarkdown(t *testing.T) {
	projectRoot := writePureJavaProject(t)
	outPath := filepath.Join(t.TempDir(), "migration-report.md")
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"report", projectRoot, "--format", "markdown", "--out", outPath}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected report success, got %d; stderr=%s", code, stderr.String())
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	if !strings.Contains(string(data), "# Java-To-Go Migration Report") {
		t.Fatalf("expected markdown report, got:\n%s", string(data))
	}
}

func writePureJavaProject(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "pom.xml"), `<project></project>`)
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "example", "Calculator.java"), `package com.example;

public class Calculator {
    private int last;

    public int add(int a, int b) {
        last = a + b;
        return last;
    }
}`)
	writeFile(t, filepath.Join(root, "src", "test", "java", "com", "example", "CalculatorTest.java"), `package com.example.test;

import com.example.Calculator;

public class CalculatorTest {
    public static int smoke() {
        return Calculator.add(1, 2);
    }
}`)
	return root
}

func writeFile(t *testing.T, path string, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
