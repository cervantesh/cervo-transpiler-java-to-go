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

func TestLowerAddsTypedIRMetadata(t *testing.T) {
	source := `package demo;

public class Calculator {
  public static int add() {
    return 1 + 2;
  }
}`
	tree, syntax := parser.ParseSource("Calculator.java", source)
	if len(syntax) != 0 {
		t.Fatalf("syntax diagnostics: %#v", syntax)
	}
	file, diagnostics := Lower("Calculator.java", tree)
	if len(diagnostics) != 0 {
		t.Fatalf("lower diagnostics: %#v", diagnostics)
	}
	if file.Path != "Calculator.java" {
		t.Fatalf("expected file path metadata, got %#v", file)
	}
	if len(file.Classes) != 1 || file.Classes[0].Symbol != "demo.Calculator" {
		t.Fatalf("expected class symbol, got %#v", file.Classes)
	}
	if file.Classes[0].Span.Line == 0 {
		t.Fatalf("expected class source span, got %#v", file.Classes[0].Span)
	}
	if len(file.Funcs) != 1 || file.Funcs[0].Symbol != "demo.Calculator.add" {
		t.Fatalf("expected method symbol, got %#v", file.Funcs)
	}
	ret, ok := file.Funcs[0].Body[0].(ir.ReturnStmt)
	if !ok {
		t.Fatalf("expected return statement, got %#v", file.Funcs[0].Body)
	}
	binary, ok := ret.Value.(ir.BinaryExpr)
	if !ok {
		t.Fatalf("expected binary expression, got %#v", ret.Value)
	}
	if binary.Type.Kind != ir.KindInt {
		t.Fatalf("expected typed int binary expression, got %#v", binary.Type)
	}
}
