package migrate

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunMigratesSimpleLibraryAndGeneratedGoCompiles(t *testing.T) {
	projectRoot := writeMigratableProject(t)
	outDir := filepath.Join(t.TempDir(), "go-lib")
	reportPath := filepath.Join(t.TempDir(), "migration-report.md")

	result, err := Run(projectRoot, Options{OutDir: outDir, ReportPath: reportPath, ModulePath: "example.com/migrated"})
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if result.Summary.JavaFiles != 2 || result.Summary.Generated != 2 || result.Summary.Skipped != 0 || result.Summary.Diagnostics != 0 {
		t.Fatalf("unexpected summary: %#v", result.Summary)
	}

	for _, path := range []string{
		filepath.Join(outDir, "go.mod"),
		filepath.Join(outDir, "com", "acme", "math", "Calculator.go"),
		filepath.Join(outDir, "com", "acme", "model", "Counter.go"),
		reportPath,
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected %s: %v", path, err)
		}
	}

	calculator, err := os.ReadFile(filepath.Join(outDir, "com", "acme", "math", "Calculator.go"))
	if err != nil {
		t.Fatalf("read generated Calculator.go: %v", err)
	}
	for _, want := range []string{"package math", "func Add(a int, b int) int", "return a + b"} {
		if !strings.Contains(string(calculator), want) {
			t.Fatalf("expected %q in generated Calculator.go:\n%s", want, string(calculator))
		}
	}

	counter, err := os.ReadFile(filepath.Join(outDir, "com", "acme", "model", "Counter.go"))
	if err != nil {
		t.Fatalf("read generated Counter.go: %v", err)
	}
	for _, want := range []string{"package model", "type Counter struct", "func NewCounter(value int) *Counter", "func (c *Counter) Value() int"} {
		if !strings.Contains(string(counter), want) {
			t.Fatalf("expected %q in generated Counter.go:\n%s", want, string(counter))
		}
	}

	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	for _, want := range []string{"# Java-To-Go Migration Report", "## Migration Execution Summary", "| Generated files | 2 |", "None."} {
		if !strings.Contains(string(report), want) {
			t.Fatalf("expected %q in migration report:\n%s", want, string(report))
		}
	}

	runGoTest(t, outDir)
}

func TestRunDryRunWritesReportWithoutGeneratedOutput(t *testing.T) {
	projectRoot := writeMigratableProject(t)
	tmp := t.TempDir()
	outDir := filepath.Join(tmp, "go-lib")
	reportPath := filepath.Join(tmp, "migration-report.md")
	logPath := filepath.Join(tmp, "j2go.log")

	result, err := Run(projectRoot, Options{OutDir: outDir, ReportPath: reportPath, DryRun: true, LogPath: logPath})
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if !result.Summary.DryRun || result.Summary.Generated != 2 || result.Summary.Skipped != 0 {
		t.Fatalf("unexpected dry-run summary: %#v", result.Summary)
	}
	if _, err := os.Stat(outDir); !os.IsNotExist(err) {
		t.Fatalf("expected dry-run to skip generated output dir, stat err=%v", err)
	}
	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("expected dry-run report: %v", err)
	}
	if !strings.Contains(string(report), "| Dry run | true |") {
		t.Fatalf("expected dry-run marker in report:\n%s", string(report))
	}
	log := readFile(t, logPath)
	for _, want := range []string{"command=migrate", "dryRun=true", "generated=2", "diagnostics=0"} {
		if !strings.Contains(log, want) {
			t.Fatalf("expected %q in log:\n%s", want, log)
		}
	}
}

func TestRunPartialMigrationSkipsUnsupportedFiles(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "ok", "Ok.java"), `package com.acme.ok;

public class Ok {
    public static int Answer() {
        return 42;
    }
}`)
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "broken", "Broken.java"), `package com.acme.broken;

public class Broken {
    public static void Fail() {
        {
            int x = 1;
        }
    }
}`)

	outDir := filepath.Join(t.TempDir(), "go-lib")
	reportPath := filepath.Join(t.TempDir(), "migration-report.md")
	result, err := Run(root, Options{OutDir: outDir, ReportPath: reportPath})
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if result.Summary.Generated != 1 || result.Summary.Skipped != 1 || result.Summary.Diagnostics == 0 {
		t.Fatalf("unexpected partial summary: %#v diagnostics=%#v", result.Summary, result.Diagnostics)
	}
	if _, err := os.Stat(filepath.Join(outDir, "com", "acme", "ok", "Ok.go")); err != nil {
		t.Fatalf("expected supported file output: %v", err)
	}
	if _, err := os.Stat(filepath.Join(outDir, "com", "acme", "broken", "Broken.go")); !os.IsNotExist(err) {
		t.Fatalf("expected unsupported file to be skipped, stat err=%v", err)
	}
	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	for _, want := range []string{"| Generated files | 1 |", "| Skipped files | 1 |", "JTG1020"} {
		if !strings.Contains(string(report), want) {
			t.Fatalf("expected %q in migration report:\n%s", want, string(report))
		}
	}
}

func TestRunMigratesInternalImportsAndBasicTests(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "math", "Calculator.java"), `package com.acme.math;

public class Calculator {
    public static int add(int a, int b) {
        return a + b;
    }
}`)
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "service", "CalculatorService.java"), `package com.acme.service;

import com.acme.math.Calculator;

public class CalculatorService {
    public static int total() {
        return Calculator.add(2, 3);
    }
}`)
	writeFile(t, filepath.Join(root, "src", "test", "java", "com", "acme", "math", "CalculatorTest.java"), `package com.acme.math;

public class CalculatorTest {
    public static void smoke() {
        Calculator.add(1, 2);
    }
}`)

	outDir := filepath.Join(t.TempDir(), "go-lib")
	result, err := Run(root, Options{OutDir: outDir, ModulePath: "example.com/migrated"})
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if result.Summary.Generated != 3 || result.Summary.Skipped != 0 || result.Summary.Diagnostics != 0 {
		t.Fatalf("unexpected summary: %#v diagnostics=%#v", result.Summary, result.Diagnostics)
	}

	calculator := readFile(t, filepath.Join(outDir, "com", "acme", "math", "Calculator.go"))
	for _, want := range []string{"package math", "func Add(a int, b int) int", "return a + b"} {
		if !strings.Contains(calculator, want) {
			t.Fatalf("expected %q in Calculator.go:\n%s", want, calculator)
		}
	}
	service := readFile(t, filepath.Join(outDir, "com", "acme", "service", "CalculatorService.go"))
	for _, want := range []string{
		"package service",
		"import \"example.com/migrated/com/acme/math\"",
		"func Total() int",
		"return math.Add(2, 3)",
	} {
		if !strings.Contains(service, want) {
			t.Fatalf("expected %q in CalculatorService.go:\n%s", want, service)
		}
	}
	testFile := readFile(t, filepath.Join(outDir, "com", "acme", "math", "CalculatorTest_test.go"))
	for _, want := range []string{"package math", "import \"testing\"", "func TestSmoke(t *testing.T)", "Add(1, 2)"} {
		if !strings.Contains(testFile, want) {
			t.Fatalf("expected %q in CalculatorTest_test.go:\n%s", want, testFile)
		}
	}
	runGoTest(t, outDir)
}

func TestRunReportsGoSymbolCollisions(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "one", "First.java"), `package com.acme.one;

public class First {
    public static int value() {
        return 1;
    }
}`)
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "one", "Second.java"), `package com.acme.one;

public class Second {
    public static int Value() {
        return 2;
    }
}`)

	outDir := filepath.Join(t.TempDir(), "go-lib")
	result, err := Run(root, Options{OutDir: outDir})
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if result.Summary.Generated != 0 || result.Summary.Skipped != 2 || result.Summary.Diagnostics != 2 {
		t.Fatalf("unexpected summary: %#v diagnostics=%#v", result.Summary, result.Diagnostics)
	}
	for _, diagnostic := range result.Diagnostics {
		if !strings.Contains(diagnostic.Error(), "JTG2001") || !strings.Contains(diagnostic.Error(), "duplicate Go symbol after migration: Value") {
			t.Fatalf("expected Go symbol collision diagnostic, got %v", diagnostic)
		}
	}
}

func writeMigratableProject(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "pom.xml"), `<project></project>`)
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "math", "Calculator.java"), `package com.acme.math;

public class Calculator {
    public static int Add(int a, int b) {
        return a + b;
    }
}`)
	writeFile(t, filepath.Join(root, "src", "main", "java", "com", "acme", "model", "Counter.java"), `package com.acme.model;

public class Counter {
    private int value;

    public Counter(int value) {
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

func runGoTest(t *testing.T, dir string) {
	t.Helper()
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOWORK=off")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated Go tests failed: %v\n%s", err, string(output))
	}
}
