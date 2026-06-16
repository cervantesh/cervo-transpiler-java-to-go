package emitgo

import (
	"strings"
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func TestEmitMainWithPrintln(t *testing.T) {
	file := ir.File{Funcs: []ir.Func{{
		Name:       "main",
		ReturnType: ir.Type{Kind: ir.KindVoid},
		Body: []ir.Stmt{
			ir.VarDeclStmt{Name: "count", Type: ir.Type{Kind: ir.KindInt}, Value: ir.LiteralExpr{Value: "2", Type: ir.Type{Kind: ir.KindInt}}},
			ir.ExprStmt{Expr: ir.CallExpr{Target: "System.out", Name: "println", Args: []ir.Expr{ir.NameExpr{Name: "count"}}}},
		},
	}}}
	code, err := Emit(file)
	if err != nil {
		t.Fatalf("Emit failed: %v", err)
	}
	text := string(code)
	for _, want := range []string{"package main", "import \"fmt\"", "count := 2", "fmt.Println(count)"} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in output:\n%s", want, text)
		}
	}
}

func TestEmitNoFmtWhenUnused(t *testing.T) {
	code, err := Emit(ir.File{Funcs: []ir.Func{{Name: "main", ReturnType: ir.Type{Kind: ir.KindVoid}}}})
	if err != nil {
		t.Fatalf("Emit failed: %v", err)
	}
	if strings.Contains(string(code), "fmt") {
		t.Fatalf("did not expect fmt import:\n%s", string(code))
	}
}

func TestEmitUsesIRPackageName(t *testing.T) {
	code, err := Emit(ir.File{PackageName: "sample", Funcs: []ir.Func{{Name: "Answer", ReturnType: ir.Type{Kind: ir.KindInt}, Body: []ir.Stmt{
		ir.ReturnStmt{Value: ir.LiteralExpr{Value: "42", Type: ir.Type{Kind: ir.KindInt}}},
	}}}})
	if err != nil {
		t.Fatalf("Emit failed: %v", err)
	}
	text := string(code)
	for _, want := range []string{"package sample", "func Answer() int"} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in output:\n%s", want, text)
		}
	}
}

func TestEmitExplicitImportsAndQualifiedTypes(t *testing.T) {
	modelType := ir.Type{Kind: ir.KindObject, Name: "model.User"}
	code, err := Emit(ir.File{
		PackageName: "service",
		Imports:     []ir.Import{{Path: "example.com/migrated/com/acme/model"}},
		Funcs: []ir.Func{{
			Name:       "Use",
			Params:     []ir.Param{{Name: "user", Type: modelType}},
			ReturnType: modelType,
			Body:       []ir.Stmt{ir.ReturnStmt{Value: ir.NameExpr{Name: "user", Type: modelType}}},
		}},
	})
	if err != nil {
		t.Fatalf("Emit failed: %v", err)
	}
	text := string(code)
	for _, want := range []string{
		"package service",
		"import \"example.com/migrated/com/acme/model\"",
		"func Use(user model.User) model.User",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("expected %q in output:\n%s", want, text)
		}
	}
}
