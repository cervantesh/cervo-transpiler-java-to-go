package ir

import "testing"

func TestTypeGoName(t *testing.T) {
	tests := map[Kind]string{
		KindVoid:    "",
		KindBoolean: "bool",
		KindInt:     "int",
		KindDouble:  "float64",
		KindString:  "string",
	}
	for kind, want := range tests {
		if got := (Type{Kind: kind}).GoName(); got != want {
			t.Fatalf("kind %v: expected %q, got %q", kind, want, got)
		}
	}
}

func TestArrayTypeGoName(t *testing.T) {
	typ := Type{Kind: KindArray, Elem: &Type{Kind: KindInt}}
	if got := typ.GoName(); got != "[]int" {
		t.Fatalf("expected []int, got %q", got)
	}
}
