package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/report"
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
	t.Setenv("J2GO_AI_PROVIDER", "external")
	t.Setenv("J2GO_AI_COMMAND", "")

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

func TestRunTranspilePrintsAllUnsupportedDiagnostics(t *testing.T) {
	tmp := t.TempDir()
	input := filepath.Join(tmp, "Demo.java")
	outDir := filepath.Join(tmp, "out")
	source := `package demo;

import java.util.List;

public class Demo {
  public static void main(String[] args) {
    System.out.println(1);
  }
}`
	if err := os.WriteFile(input, []byte(source), 0644); err != nil {
		t.Fatalf("write source: %v", err)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	code := run([]string{"transpile", "-out", outDir, input}, &stdout, &stderr)
	if code != 1 {
		t.Fatalf("expected failure exit code 1, got %d; stderr=%s", code, stderr.String())
	}
	for _, want := range []string{"JTG1001", "JTG1002", "unsupported feature:", "recommendation:"} {
		if !strings.Contains(stderr.String(), want) {
			t.Fatalf("expected %q in stderr:\n%s", want, stderr.String())
		}
	}
	if _, err := os.Stat(filepath.Join(outDir, "Demo.go")); !os.IsNotExist(err) {
		t.Fatalf("expected no generated file, stat err=%v", err)
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

func TestRunMigrateDryRunWritesReport(t *testing.T) {
	projectRoot := writeMigratableCLIProject(t)
	tmp := t.TempDir()
	outDir := filepath.Join(tmp, "go-lib")
	reportPath := filepath.Join(tmp, "migration-report.md")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	t.Setenv("J2GO_AI_PROVIDER", "external")
	t.Setenv("J2GO_AI_COMMAND", "")

	code := run([]string{"migrate", projectRoot, "--out", outDir, "--report", reportPath, "--dry-run"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected migrate success, got %d; stderr=%s", code, stderr.String())
	}
	for _, want := range []string{"Migration summary:", "javaFiles=1", "generated=1", "dryRun=true", "Report: " + reportPath} {
		if !strings.Contains(stdout.String(), want) {
			t.Fatalf("expected %q in stdout:\n%s", want, stdout.String())
		}
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected empty stderr, got %q", stderr.String())
	}
	if _, err := os.Stat(outDir); !os.IsNotExist(err) {
		t.Fatalf("expected dry-run to skip generated output dir, stat err=%v", err)
	}
	data, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	if !strings.Contains(string(data), "## Migration Execution Summary") {
		t.Fatalf("expected execution summary in report:\n%s", string(data))
	}
}

func TestRunMigrateUsesConfigAndWritesLog(t *testing.T) {
	projectRoot := writeMigratableCLIProject(t)
	tmp := t.TempDir()
	outDir := filepath.Join(tmp, "go-lib")
	reportPath := filepath.Join(tmp, "migration-report.md")
	logPath := filepath.Join(tmp, "j2go.log")
	configPath := filepath.Join(tmp, "cervo-migration.yaml")
	writeFile(t, configPath, "version: 1\nproject:\n  source: "+slashPath(projectRoot)+"\n  module: example.com/generated\nmigration:\n  out: "+slashPath(outDir)+"\n  report: "+slashPath(reportPath)+"\n  dryRun: true\nlogs:\n  file: "+slashPath(logPath)+"\n")
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"migrate", "--config", configPath}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected migrate success, got %d; stderr=%s", code, stderr.String())
	}
	for _, want := range []string{"Migration summary:", "dryRun=true", "Report: " + slashPath(reportPath), "Log: " + slashPath(logPath)} {
		if !strings.Contains(stdout.String(), want) {
			t.Fatalf("expected %q in stdout:\n%s", want, stdout.String())
		}
	}
	if _, err := os.Stat(outDir); !os.IsNotExist(err) {
		t.Fatalf("expected config dry-run to skip generated output dir, stat err=%v", err)
	}
	log, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	for _, want := range []string{"command=migrate", "dryRun=true", "generated=1"} {
		if !strings.Contains(string(log), want) {
			t.Fatalf("expected %q in log:\n%s", want, string(log))
		}
	}
}

func TestRunConfigValidate(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "cervo-migration.yaml")
	writeFile(t, configPath, "version: 1\nproject:\n  source: ./java-lib\nmigration:\n  out: ./go-lib\n")
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"config", "validate", "--config", configPath}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected config validation success, got %d; stderr=%s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Config valid: "+configPath) {
		t.Fatalf("unexpected stdout:\n%s", stdout.String())
	}
}

func TestRunAIExplainWritesAdvisory(t *testing.T) {
	tmp := t.TempDir()
	reportPath := filepath.Join(tmp, "migration-report.json")
	outPath := filepath.Join(tmp, "ai-review.md")
	snippetPath := filepath.Join(tmp, "Generated.go")
	configPath := filepath.Join(tmp, "cervo-migration.yaml")
	project := javaproject.Project{
		Root: "sample-project",
		Files: []javaproject.File{{
			RelativePath: "src/main/java/com/example/Demo.java",
			Risk:         "low",
		}},
		Summary: javaproject.Summary{JavaFiles: 1, Packages: 1, Classes: 1},
	}
	data, err := report.JSON(project)
	if err != nil {
		t.Fatalf("render report: %v", err)
	}
	writeFile(t, reportPath, string(data))
	writeFile(t, snippetPath, "package example\n\nfunc Demo() {}\n")
	writeFile(t, configPath, "version: 1\nai:\n  provider: canned\n")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	code := run([]string{"ai", "explain", "--config", configPath, "--report", reportPath, "--out", outPath, "--snippet", snippetPath}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("expected ai explain success, got %d; stderr=%s", code, stderr.String())
	}
	for _, want := range []string{"AI advisory written:", "Provider: canned", "Structured report parsed: true"} {
		if !strings.Contains(stdout.String(), want) {
			t.Fatalf("expected %q in stdout:\n%s", want, stdout.String())
		}
	}
	output, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read AI advisory: %v", err)
	}
	for _, want := range []string{"# AI Advisory Migration Review", "Advisory output only", "Generated-Code Snippets Reviewed"} {
		if !strings.Contains(string(output), want) {
			t.Fatalf("expected %q in AI advisory:\n%s", want, string(output))
		}
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

func writeMigratableCLIProject(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "example", "simple", "NumberBox.java"), `package com.example.simple;

public class NumberBox {
    private int value;

    public NumberBox(int value) {
    }

    public int Value() {
        return value;
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

func slashPath(path string) string {
	return filepath.ToSlash(path)
}
