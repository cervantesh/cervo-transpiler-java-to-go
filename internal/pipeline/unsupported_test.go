package pipeline

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnsupportedFixturesReturnStructuredDiagnostics(t *testing.T) {
	root := workspaceRoot(t)
	cases := map[string]string{
		"unsupported_field":           "JTG1016",
		"unsupported_package_import":  "JTG1001",
		"unsupported_exception":       "JTG1007",
		"unsupported_instance_method": "JTG1017",
		"unsupported_overload":        "JTG1018",
		"unsupported_interface":       "JTG1003",
		"unsupported_inheritance":     "JTG1004",
		"unsupported_generics":        "JTG1014",
		"unsupported_array_indexing":  "JTG1015",
	}

	for name, code := range cases {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(root, "tests", "fixtures", name+".java")
			source, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}

			result := TranspileSource(path, string(source))
			if len(result.Diagnostics) == 0 {
				t.Fatalf("expected diagnostics")
			}
			if len(result.Code) != 0 {
				t.Fatalf("expected no generated Go for unsupported source, got:\n%s", string(result.Code))
			}
			if !diagnosticsContain(result.Diagnostics, code) {
				t.Fatalf("expected code %s in %#v", code, result.Diagnostics)
			}
			if !diagnosticsContain(result.Diagnostics, "unsupported feature:") {
				t.Fatalf("expected unsupported feature message in %#v", result.Diagnostics)
			}
			if !diagnosticsContain(result.Diagnostics, "recommendation:") {
				t.Fatalf("expected recommendation in %#v", result.Diagnostics)
			}
		})
	}
}

func TestUnsupportedPackageImportReturnsAllDiagnostics(t *testing.T) {
	source := `package demo;

import java.util.List;

public class Demo {
  public static void main(String[] args) {
    System.out.println(1);
  }
}`

	result := TranspileSource("Demo.java", source)
	for _, code := range []string{"JTG1001", "JTG1002"} {
		if !diagnosticsContain(result.Diagnostics, code) {
			t.Fatalf("expected code %s in %#v", code, result.Diagnostics)
		}
	}
}

func diagnosticsContain(diagnostics []error, want string) bool {
	for _, diagnostic := range diagnostics {
		if strings.Contains(diagnostic.Error(), want) {
			return true
		}
	}
	return false
}
