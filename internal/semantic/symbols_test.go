package semantic

import (
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func TestScopeLookupFindsNearestSymbol(t *testing.T) {
	root := NewScope(nil)
	if err := root.Define(Symbol{Name: "x", Type: ir.Type{Kind: ir.KindInt}}); err != nil {
		t.Fatalf("define root symbol: %v", err)
	}
	child := NewScope(root)
	if err := child.Define(Symbol{Name: "x", Type: ir.Type{Kind: ir.KindString}}); err != nil {
		t.Fatalf("define child symbol: %v", err)
	}

	symbol, ok := child.Lookup("x")
	if !ok {
		t.Fatal("expected symbol")
	}
	if symbol.Type.Kind != ir.KindString {
		t.Fatalf("expected nearest string symbol, got %#v", symbol)
	}
}

func TestScopeRejectsDuplicateInSameScope(t *testing.T) {
	scope := NewScope(nil)
	if err := scope.Define(Symbol{Name: "x", Type: ir.Type{Kind: ir.KindInt}}); err != nil {
		t.Fatalf("first define failed: %v", err)
	}
	if err := scope.Define(Symbol{Name: "x", Type: ir.Type{Kind: ir.KindInt}}); err == nil {
		t.Fatal("expected duplicate symbol error")
	}
}
