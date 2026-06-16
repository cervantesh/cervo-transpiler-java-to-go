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

func TestLowerUsesTypedScopeAndFieldAssignment(t *testing.T) {
	source := `package demo;

public class Counter {
  private int value;

  public int bump(int delta) {
    int next = delta + 1;
    value = next;
    return value;
  }
}`
	tree, syntax := parser.ParseSource("Counter.java", source)
	if len(syntax) != 0 {
		t.Fatalf("syntax diagnostics: %#v", syntax)
	}
	file, diagnostics := Lower("Counter.java", tree)
	if len(diagnostics) != 0 {
		t.Fatalf("lower diagnostics: %#v", diagnostics)
	}
	if len(file.Classes) != 1 || len(file.Classes[0].Fields) != 1 {
		t.Fatalf("expected one class field, got %#v", file.Classes)
	}
	if file.Classes[0].Fields[0].Symbol != "demo.Counter.value" {
		t.Fatalf("expected field symbol, got %#v", file.Classes[0].Fields[0])
	}
	method := file.Classes[0].Methods[0]
	decl, ok := method.Body[0].(ir.VarDeclStmt)
	if !ok {
		t.Fatalf("expected var decl, got %#v", method.Body[0])
	}
	binary, ok := decl.Value.(ir.BinaryExpr)
	if !ok || binary.Type.Kind != ir.KindInt {
		t.Fatalf("expected typed int binary from local scope, got %#v", decl.Value)
	}
	assign, ok := method.Body[1].(ir.AssignStmt)
	if !ok || assign.Name != "c.value" {
		t.Fatalf("expected receiver field assignment, got %#v", method.Body[1])
	}
	ret, ok := method.Body[2].(ir.ReturnStmt)
	if !ok {
		t.Fatalf("expected return, got %#v", method.Body[2])
	}
	field, ok := ret.Value.(ir.FieldExpr)
	if !ok || field.Type.Kind != ir.KindInt {
		t.Fatalf("expected typed field return, got %#v", ret.Value)
	}
}
