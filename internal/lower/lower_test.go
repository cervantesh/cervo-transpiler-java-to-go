package lower

import (
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser"
)

func TestLowerMainPrintln(t *testing.T) {
	source := `public class Hello {
  public static void main(String[] args) {
    int count = 2;
    System.out.println(count + 3);
  }
}`
	tree, syntax := parser.ParseSource("Hello.java", source)
	if len(syntax) != 0 {
		t.Fatalf("syntax diagnostics: %#v", syntax)
	}
	file, diagnostics := Lower("Hello.java", tree)
	if len(diagnostics) != 0 {
		t.Fatalf("lower diagnostics: %#v", diagnostics)
	}
	if len(file.Funcs) != 1 {
		t.Fatalf("expected one top-level function, got %#v", file)
	}
	if file.Funcs[0].Name != "main" {
		t.Fatalf("expected main, got %q", file.Funcs[0].Name)
	}
	if _, ok := file.Funcs[0].Body[1].(ir.ExprStmt); !ok {
		t.Fatalf("expected println expression statement, got %#v", file.Funcs[0].Body[1])
	}
}

func TestMapJavaArrayType(t *testing.T) {
	typ := MapJavaType("String[]")
	if typ.Kind != ir.KindArray || typ.Elem == nil || typ.Elem.Kind != ir.KindString {
		t.Fatalf("expected string array, got %#v", typ)
	}
}
