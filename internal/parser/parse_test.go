package parser

import "testing"

func TestParseValidMainClass(t *testing.T) {
	source := `public class Hello {
  public static void main(String[] args) {
    System.out.println("hello");
  }
}`
	tree, diagnostics := ParseSource("Hello.java", source)
	if len(diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %#v", diagnostics)
	}
	if tree == nil {
		t.Fatal("expected parse tree")
	}
}

func TestParseInvalidClassReportsLineAndColumn(t *testing.T) {
	source := `public class Broken {
  public static void main(String[] args) {
    int x =
  }
}`
	_, diagnostics := ParseSource("Broken.java", source)
	if len(diagnostics) == 0 {
		t.Fatal("expected syntax diagnostics")
	}
	if diagnostics[0].Line == 0 || diagnostics[0].Column < 0 {
		t.Fatalf("expected line and column, got %#v", diagnostics[0])
	}
}
