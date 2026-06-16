package pipeline

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var legacySupportedFixtures = []string{
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

func TestLegacySupportedFixturesCompileAndRunThroughModernPipeline(t *testing.T) {
	root := workspaceRoot(t)
	for _, fixture := range legacySupportedFixtures {
		t.Run(fixture, func(t *testing.T) {
			sourcePath := filepath.Join(root, "tests", "fixtures", fixture+".java")
			source, err := os.ReadFile(sourcePath)
			if err != nil {
				t.Fatalf("read source: %v", err)
			}

			result := TranspileSource(sourcePath, string(source))
			if len(result.Diagnostics) != 0 {
				t.Fatalf("diagnostics: %#v", result.Diagnostics)
			}

			outPath := filepath.Join(t.TempDir(), fixture+".go")
			if err := os.WriteFile(outPath, result.Code, 0644); err != nil {
				t.Fatalf("write generated Go: %v", err)
			}

			cmd := exec.Command("go", "run", outPath)
			cmd.Dir = root
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("generated Go did not run: %v\n%s", err, string(output))
			}
		})
	}
}
