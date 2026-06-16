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
		"unsupported_package_import": "JTG1001",
		"unsupported_exception":      "JTG1007",
		"unsupported_overload":       "JTG1018",
		"unsupported_interface":      "JTG1003",
		"unsupported_inheritance":    "JTG1004",
		"unsupported_generics":       "JTG1014",
		"unsupported_array_indexing": "JTG1015",
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

func TestSyntaxErrorFixtureReturnsParserDiagnostic(t *testing.T) {
	root := workspaceRoot(t)
	path := filepath.Join(root, "tests", "fixtures", "syntax_error.java")
	source, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	result := TranspileSource(path, string(source))
	if len(result.Diagnostics) == 0 {
		t.Fatal("expected syntax diagnostics")
	}
	if len(result.Code) != 0 {
		t.Fatalf("expected no generated Go for syntax error, got:\n%s", string(result.Code))
	}
	if !diagnosticsContain(result.Diagnostics, "syntax_error.java:") {
		t.Fatalf("expected parser diagnostic with source position, got %#v", result.Diagnostics)
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

func TestAssignmentExpressionReturnsDiagnostic(t *testing.T) {
	source := `public class Demo {
  public static void main(String[] args) {
    int count = 1;
    System.out.println(count = 2);
  }
}`

	result := TranspileSource("Demo.java", source)
	if len(result.Code) != 0 {
		t.Fatalf("expected no generated Go, got:\n%s", string(result.Code))
	}
	if !diagnosticsContain(result.Diagnostics, "JTG1019") {
		t.Fatalf("expected assignment expression diagnostic in %#v", result.Diagnostics)
	}
	if diagnosticsContain(result.Diagnostics, "__assign_") {
		t.Fatalf("expected no assignment sentinel in diagnostics: %#v", result.Diagnostics)
	}
}

func TestStandaloneBlockReturnsDiagnostic(t *testing.T) {
	source := `public class Demo {
  public static void main(String[] args) {
    {
      System.out.println(1);
    }
  }
}`

	result := TranspileSource("Demo.java", source)
	if len(result.Code) != 0 {
		t.Fatalf("expected no generated Go, got:\n%s", string(result.Code))
	}
	if !diagnosticsContain(result.Diagnostics, "JTG1020") {
		t.Fatalf("expected standalone block diagnostic in %#v", result.Diagnostics)
	}
	if diagnosticsContain(result.Diagnostics, "__unsupported_block") {
		t.Fatalf("expected no unsupported block sentinel in diagnostics: %#v", result.Diagnostics)
	}
}

func TestDuplicateLocalVariableReturnsSemanticDiagnostic(t *testing.T) {
	source := `public class Demo {
  public static void main(String[] args) {
    int count = 1;
    int count = 2;
    System.out.println(count);
  }
}`

	result := TranspileSource("Demo.java", source)
	if !diagnosticsContain(result.Diagnostics, "JTG2001") {
		t.Fatalf("expected duplicate local diagnostic in %#v", result.Diagnostics)
	}
	if len(result.Code) != 0 {
		t.Fatalf("expected no generated Go, got:\n%s", string(result.Code))
	}
}

func TestUnknownSymbolReturnsSemanticDiagnostic(t *testing.T) {
	source := `public class Demo {
  public static void main(String[] args) {
    missing = 1;
  }
}`

	result := TranspileSource("Demo.java", source)
	if !diagnosticsContain(result.Diagnostics, "JTG2002") {
		t.Fatalf("expected unknown symbol diagnostic in %#v", result.Diagnostics)
	}
	if !diagnosticsContain(result.Diagnostics, "Demo.java:") {
		t.Fatalf("expected source position in %#v", result.Diagnostics)
	}
	if len(result.Code) != 0 {
		t.Fatalf("expected no generated Go, got:\n%s", string(result.Code))
	}
}

func TestIncompatibleAssignmentReturnsSemanticDiagnostic(t *testing.T) {
	source := `public class Demo {
  public static void main(String[] args) {
    int count = "wrong";
    System.out.println(count);
  }
}`

	result := TranspileSource("Demo.java", source)
	if !diagnosticsContain(result.Diagnostics, "JTG2003") {
		t.Fatalf("expected incompatible type diagnostic in %#v", result.Diagnostics)
	}
	if len(result.Code) != 0 {
		t.Fatalf("expected no generated Go, got:\n%s", string(result.Code))
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
