package unsupported

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/semantic"
)

func TestDetectUnsupportedFixtureCodes(t *testing.T) {
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

			diagnostics := Detect(path, string(source))
			if !hasCode(diagnostics, code) {
				t.Fatalf("expected code %s in %#v", code, diagnostics)
			}
			for _, diagnostic := range diagnostics {
				if diagnostic.Recommendation == "" {
					t.Fatalf("expected recommendation in %#v", diagnostic)
				}
			}
		})
	}
}

func TestDetectReportsMultipleUnsupportedFeatures(t *testing.T) {
	source := `package demo;

import java.util.List;

public class Demo {
  public static void main(String[] args) {
    System.out.println(1);
  }
}`

	diagnostics := Detect("Demo.java", source)
	for _, code := range []string{"JTG1001", "JTG1002"} {
		if !hasCode(diagnostics, code) {
			t.Fatalf("expected %s in %#v", code, diagnostics)
		}
	}
}

func TestDetectIgnoresCommentsLiteralsAndMainArrayParameter(t *testing.T) {
	source := `public class Demo {
  public static void main(String[] args) {
    // try import package interface
    String text = "new List<String>";
    System.out.println(text);
  }
}`

	diagnostics := Detect("Demo.java", source)
	if len(diagnostics) != 0 {
		t.Fatalf("expected no unsupported diagnostics, got %#v", diagnostics)
	}
}

func hasCode(diagnostics []semantic.Diagnostic, code string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			return true
		}
	}
	return false
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
