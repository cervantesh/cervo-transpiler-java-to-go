package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var supportedCases = []string{
	"hello_print",
	"variables",
	"if_else",
	"while_loop",
	"method_return",
	"arithmetic_precedence",
	"parenthesized_expression",
	"unary_minus",
	"boolean_and_or",
	"boolean_not",
	"equality_inequality",
	"comparison_bounds",
	"string_variable",
	"double_arithmetic",
	"comments_handling",
	"no_fmt_import",
	"return_without_value",
	"static_void_method_call",
	"nested_control",
	"multiple_parameters",
}

var unsupportedCases = []unsupportedCase{
	{Name: "unsupported_field", Code: "JTG1016", Feature: "class fields"},
	{Name: "unsupported_package_import", Code: "JTG1001", Feature: "package declarations"},
	{Name: "unsupported_exception", Code: "JTG1007", Feature: "try/catch exceptions"},
	{Name: "unsupported_instance_method", Code: "JTG1017", Feature: "instance methods"},
	{Name: "unsupported_overload", Code: "JTG1018", Feature: "method overloading"},
	{Name: "unsupported_interface", Code: "JTG1003", Feature: "interfaces"},
	{Name: "unsupported_inheritance", Code: "JTG1004", Feature: "inheritance"},
	{Name: "unsupported_generics", Code: "JTG1014", Feature: "generics"},
	{Name: "unsupported_array_indexing", Code: "JTG1015", Feature: "arrays and indexing"},
}

type unsupportedCase struct {
	Name    string
	Code    string
	Feature string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	exe, err := buildLegacy(root)
	if err != nil {
		return err
	}
	generatedDir := filepath.Join(root, "tests", "generated")
	if err := os.MkdirAll(generatedDir, 0755); err != nil {
		return err
	}
	for _, name := range supportedCases {
		if err := runSupportedCase(root, exe, generatedDir, name); err != nil {
			return err
		}
		fmt.Printf("PASS %s\n", name)
	}
	if err := runSyntaxErrorCase(root, exe, generatedDir); err != nil {
		return err
	}
	fmt.Println("PASS syntax_error")
	for _, tc := range unsupportedCases {
		if err := runUnsupportedCase(root, exe, generatedDir, tc); err != nil {
			return err
		}
		fmt.Printf("PASS %s\n", tc.Name)
	}
	total := len(supportedCases) + 1 + len(unsupportedCases)
	if total != 30 {
		return fmt.Errorf("expected exactly 30 evidence cases, got %d", total)
	}
	fmt.Printf("PASS evidence_case_count => %d\n", total)
	return nil
}

func repoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return os.Getwd()
	}
	return filepath.Clean(strings.TrimSpace(string(out))), nil
}

func buildLegacy(root string) (string, error) {
	buildDir := filepath.Join(root, "build")
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return "", err
	}
	bison, err := lookupTool("BISON", "bison", "win_bison")
	if err != nil {
		return "", err
	}
	flex, err := lookupTool("FLEX", "flex", "win_flex")
	if err != nil {
		return "", err
	}
	cxx, err := lookupTool("CXX", "g++")
	if err != nil {
		return "", err
	}
	parserCPP := filepath.Join(buildDir, "parser.cpp")
	lexerCPP := filepath.Join(buildDir, "lexer.cpp")
	if err := runChecked(root, bison, "-d", "-o", parserCPP, filepath.Join(root, "src", "parser.y")); err != nil {
		return "", err
	}
	if err := runChecked(root, flex, "-o", lexerCPP, filepath.Join(root, "src", "lexer.l")); err != nil {
		return "", err
	}
	exe := filepath.Join(buildDir, executableName("javago"))
	args := []string{
		"-std=c++17",
		"-I" + filepath.Join(root, "src"),
		"-I" + buildDir,
		parserCPP,
		lexerCPP,
		filepath.Join(root, "src", "ast.cpp"),
		filepath.Join(root, "src", "diagnostics.cpp"),
		filepath.Join(root, "src", "generator.cpp"),
		filepath.Join(root, "src", "main.cpp"),
		"-o", exe,
	}
	if err := runChecked(root, cxx, args...); err != nil {
		return "", err
	}
	fmt.Printf("Built %s\n", exe)
	return exe, nil
}

func lookupTool(env string, names ...string) (string, error) {
	if value := os.Getenv(env); value != "" {
		return value, nil
	}
	for _, name := range names {
		if path, err := exec.LookPath(name); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("missing tool: %s", strings.Join(names, " or "))
}

func runSupportedCase(root string, exe string, generatedDir string, name string) error {
	inputPath := filepath.Join(root, "tests", "fixtures", name+".java")
	actual := filepath.Join(generatedDir, name+".go")
	expected := filepath.Join(root, "tests", "expected", name+".go")
	if err := runChecked(root, exe, inputPath, actual); err != nil {
		return fmt.Errorf("transpile %s: %w", name, err)
	}
	if err := runChecked(root, "gofmt", "-w", actual); err != nil {
		return fmt.Errorf("gofmt %s: %w", name, err)
	}
	actualText, err := readNormalized(actual)
	if err != nil {
		return err
	}
	expectedText, err := readNormalized(expected)
	if err != nil {
		return err
	}
	if actualText != expectedText {
		return fmt.Errorf("golden test failed: %s", name)
	}
	if err := runChecked(root, "go", "run", actual); err != nil {
		return fmt.Errorf("go run %s: %w", name, err)
	}
	return nil
}

func runSyntaxErrorCase(root string, exe string, generatedDir string) error {
	inputPath := filepath.Join(root, "tests", "fixtures", "syntax_error.java")
	actual := filepath.Join(generatedDir, "syntax_error.go")
	exitCode, stderr, err := runTranspilerFailure(root, exe, inputPath, actual)
	if err != nil {
		return err
	}
	if exitCode == 0 {
		return fmt.Errorf("syntax_error.java unexpectedly succeeded")
	}
	if !strings.Contains(stderr, "line") || !strings.Contains(stderr, "column") {
		return fmt.Errorf("syntax_error.java did not report line and column")
	}
	return nil
}

func runUnsupportedCase(root string, exe string, generatedDir string, tc unsupportedCase) error {
	inputPath := filepath.Join(root, "tests", "fixtures", tc.Name+".java")
	actual := filepath.Join(generatedDir, tc.Name+".go")
	_ = os.Remove(actual)
	exitCode, stderr, err := runTranspilerFailure(root, exe, inputPath, actual)
	if err != nil {
		return err
	}
	if exitCode == 0 {
		return fmt.Errorf("%s.java unexpectedly succeeded", tc.Name)
	}
	if _, err := os.Stat(actual); err == nil {
		return fmt.Errorf("%s.java wrote Go output despite unsupported diagnostics", tc.Name)
	} else if !os.IsNotExist(err) {
		return err
	}
	if !strings.Contains(stderr, tc.Code) {
		return fmt.Errorf("%s.java did not report diagnostic code %s", tc.Name, tc.Code)
	}
	feature := "unsupported feature: " + tc.Feature
	if !strings.Contains(stderr, feature) {
		return fmt.Errorf("%s.java did not report feature %s", tc.Name, tc.Feature)
	}
	if !strings.Contains(stderr, "recommendation:") {
		return fmt.Errorf("%s.java did not include a recommendation", tc.Name)
	}
	return nil
}

func runTranspilerFailure(root string, exe string, input string, output string) (int, string, error) {
	var stderr bytes.Buffer
	cmd := exec.Command(exe, input, output)
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err == nil {
		return 0, stderr.String(), nil
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode(), stderr.String(), nil
	}
	return -1, stderr.String(), err
}

func runChecked(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s failed: %w\n%s", name, strings.Join(args, " "), err, output.String())
	}
	return nil
}

func readNormalized(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	return text, nil
}

func executableName(name string) string {
	if runtime.GOOS == "windows" {
		return name + ".exe"
	}
	return name
}
