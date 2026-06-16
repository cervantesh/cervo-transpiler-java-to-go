package pipeline

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestObjectModelFieldsConstructorAndReceiverCompile(t *testing.T) {
	source := `public class Counter {
  private int value;

  public Counter(int value) {
  }

  public int Value() {
    return value;
  }

  public static void main(String[] args) {
    Counter counter = new Counter(3);
    System.out.println(counter.Value());
  }
}`

	result := TranspileSource("Counter.java", source)
	if len(result.Diagnostics) != 0 {
		t.Fatalf("diagnostics: %#v", result.Diagnostics)
	}
	code := string(result.Code)
	for _, want := range []string{
		"type Counter struct",
		"value int",
		"func NewCounter(value int) *Counter",
		"return &Counter{value: value}",
		"func (c *Counter) Value() int",
		"return c.value",
		"counter := NewCounter(3)",
		"fmt.Println(counter.Value())",
	} {
		if !strings.Contains(code, want) {
			t.Fatalf("expected %q in generated Go:\n%s", want, code)
		}
	}

	output := runGeneratedGo(t, result.Code)
	if strings.TrimSpace(output) != "3" {
		t.Fatalf("expected generated program to print 3, got %q", output)
	}
}

func TestInstanceMethodWithoutFieldsCompiles(t *testing.T) {
	source := `public class NumberSource {
  public int Value() {
    return 7;
  }

  public static void main(String[] args) {
    NumberSource source = new NumberSource();
    System.out.println(source.Value());
  }
}`

	result := TranspileSource("NumberSource.java", source)
	if len(result.Diagnostics) != 0 {
		t.Fatalf("diagnostics: %#v", result.Diagnostics)
	}
	code := string(result.Code)
	for _, want := range []string{
		"type NumberSource struct",
		"func NewNumberSource() *NumberSource",
		"func (n *NumberSource) Value() int",
	} {
		if !strings.Contains(code, want) {
			t.Fatalf("expected %q in generated Go:\n%s", want, code)
		}
	}

	output := runGeneratedGo(t, result.Code)
	if strings.TrimSpace(output) != "7" {
		t.Fatalf("expected generated program to print 7, got %q", output)
	}
}

func TestSimpleInterfaceDeclarationCompiles(t *testing.T) {
	source := `public interface Named {
  String Name();
}

public class User {
  public String Name() {
    return "Ada";
  }

  public static void main(String[] args) {
    Named named = new User();
    System.out.println(named.Name());
  }
}`

	result := TranspileSource("User.java", source)
	if len(result.Diagnostics) != 0 {
		t.Fatalf("diagnostics: %#v", result.Diagnostics)
	}
	code := string(result.Code)
	for _, want := range []string{
		"type Named interface",
		"Name() string",
		"type User struct",
		"func NewUser() *User",
		"func (u *User) Name() string",
		"named := NewUser()",
		"fmt.Println(named.Name())",
	} {
		if !strings.Contains(code, want) {
			t.Fatalf("expected %q in generated Go:\n%s", want, code)
		}
	}

	output := runGeneratedGo(t, result.Code)
	if strings.TrimSpace(output) != "Ada" {
		t.Fatalf("expected generated program to print Ada, got %q", output)
	}
}

func runGeneratedGo(t *testing.T, code []byte) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "main.go")
	if err := os.WriteFile(path, code, 0644); err != nil {
		t.Fatalf("write generated Go: %v", err)
	}
	cmd := exec.Command("go", "run", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("generated Go did not run: %v\n%s", err, string(output))
	}
	return string(output)
}
