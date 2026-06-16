package semantic

import (
	"testing"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func TestAnalyzeIRResolvesReceiverFieldAssignment(t *testing.T) {
	file := counterFileWithAssignment(ir.LiteralExpr{Value: "1", Type: ir.Type{Kind: ir.KindInt}})
	if diagnostics := AnalyzeIR(file); len(diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %#v", diagnostics)
	}
}

func TestAnalyzeIRReportsIncompatibleFieldAssignment(t *testing.T) {
	file := counterFileWithAssignment(ir.LiteralExpr{Value: "\"bad\"", Type: ir.Type{Kind: ir.KindString}})
	diagnostics := AnalyzeIR(file)
	if !hasAnalyzeCode(diagnostics, CodeIncompatibleType) {
		t.Fatalf("expected incompatible field assignment diagnostic, got %#v", diagnostics)
	}
}

func counterFileWithAssignment(value ir.Expr) ir.File {
	counterType := ir.Type{Kind: ir.KindObject, Name: "Counter"}
	return ir.File{
		Path: "Counter.java",
		Classes: []ir.Class{{
			Name:   "Counter",
			Symbol: "demo.Counter",
			Fields: []ir.Field{{
				Name:   "value",
				Symbol: "demo.Counter.value",
				Type:   ir.Type{Kind: ir.KindInt},
			}},
			Methods: []ir.Func{{
				Name:       "Set",
				ReturnType: ir.Type{Kind: ir.KindVoid},
				Receiver:   &ir.Param{Name: "c", Type: ir.Type{Kind: ir.KindPointer, Elem: &counterType}},
				Body: []ir.Stmt{ir.AssignStmt{
					Name:  "c.value",
					Op:    "=",
					Value: value,
				}},
			}},
		}},
	}
}

func hasAnalyzeCode(diagnostics []Diagnostic, code string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == code {
			return true
		}
	}
	return false
}
