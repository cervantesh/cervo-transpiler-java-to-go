package semantic

import (
	"fmt"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

const (
	CodeUnknownSymbol    = "JTG2002"
	CodeIncompatibleType = "JTG2003"
)

func AnalyzeIR(file ir.File) []Diagnostic {
	analyzer := irAnalyzer{file: file}
	for _, fn := range file.Funcs {
		analyzer.analyzeFunc(fn)
	}
	for _, class := range file.Classes {
		for _, fn := range class.Methods {
			analyzer.analyzeFunc(fn)
		}
	}
	return analyzer.diagnostics
}

type irAnalyzer struct {
	file        ir.File
	diagnostics []Diagnostic
}

func (a *irAnalyzer) analyzeFunc(fn ir.Func) {
	scope := NewScope(nil)
	if fn.Receiver != nil {
		if err := scope.Define(Symbol{Name: fn.Receiver.Name, Type: fn.Receiver.Type}); err != nil {
			a.add(fn.Receiver.Span, CodeDuplicateSymbol, "duplicate symbol: receiver "+fn.Receiver.Name)
		}
	}
	for _, param := range fn.Params {
		if err := scope.Define(Symbol{Name: param.Name, Type: param.Type}); err != nil {
			a.add(param.Span, CodeDuplicateSymbol, "duplicate symbol: parameter "+param.Name)
		}
	}
	for _, stmt := range fn.Body {
		a.analyzeStmt(scope, fn.ReturnType, stmt)
	}
}

func (a *irAnalyzer) analyzeStmt(scope *Scope, returnType ir.Type, stmt ir.Stmt) {
	switch value := stmt.(type) {
	case ir.VarDeclStmt:
		if value.Value != nil {
			a.checkAssignable(value.Span, value.Type, a.exprType(scope, value.Value))
		}
		if err := scope.Define(Symbol{Name: value.Name, Type: value.Type}); err != nil {
			a.add(value.Span, CodeDuplicateSymbol, "duplicate symbol: local variable "+value.Name)
		}
	case ir.AssignStmt:
		expected, ok := a.assignTargetType(scope, value.Name)
		if !ok {
			a.add(value.Span, CodeUnknownSymbol, "symbol not found: "+value.Name)
			a.exprType(scope, value.Value)
			return
		}
		a.checkAssignable(value.Span, expected, a.exprType(scope, value.Value))
	case ir.ExprStmt:
		a.exprType(scope, value.Expr)
	case ir.ReturnStmt:
		if value.Value == nil {
			return
		}
		a.checkAssignable(value.Span, returnType, a.exprType(scope, value.Value))
	case ir.IfStmt:
		a.checkBoolean(value.Span, a.exprType(scope, value.Cond))
		a.analyzeNested(scope, returnType, value.Then)
		a.analyzeNested(scope, returnType, value.Else)
	case ir.WhileStmt:
		a.checkBoolean(value.Span, a.exprType(scope, value.Cond))
		a.analyzeNested(scope, returnType, value.Body)
	case ir.ForStmt:
		child := NewScope(scope)
		if value.Init != nil {
			a.analyzeStmt(child, returnType, value.Init)
		}
		if value.Cond != nil {
			a.checkBoolean(value.Span, a.exprType(child, value.Cond))
		}
		if value.Post != nil {
			a.analyzeStmt(child, returnType, value.Post)
		}
		for _, inner := range value.Body {
			a.analyzeStmt(child, returnType, inner)
		}
	}
}

func (a *irAnalyzer) assignTargetType(scope *Scope, name string) (ir.Type, bool) {
	if symbol, ok := scope.Lookup(name); ok {
		return symbol.Type, true
	}
	parts := strings.Split(name, ".")
	if len(parts) != 2 {
		return ir.Type{}, false
	}
	receiver, ok := scope.Lookup(parts[0])
	if !ok {
		return ir.Type{}, false
	}
	receiverType := receiver.Type
	if receiverType.Kind == ir.KindPointer && receiverType.Elem != nil {
		receiverType = *receiverType.Elem
	}
	if receiverType.Kind != ir.KindObject {
		return ir.Type{}, false
	}
	for _, class := range a.file.Classes {
		if class.Name != receiverType.Name && class.Symbol != receiverType.Name {
			continue
		}
		for _, field := range class.Fields {
			if field.Name == parts[1] {
				return field.Type, true
			}
		}
	}
	return ir.Type{}, false
}

func (a *irAnalyzer) analyzeNested(parent *Scope, returnType ir.Type, stmts []ir.Stmt) {
	child := NewScope(parent)
	for _, stmt := range stmts {
		a.analyzeStmt(child, returnType, stmt)
	}
}

func (a *irAnalyzer) exprType(scope *Scope, expression ir.Expr) ir.Type {
	switch value := expression.(type) {
	case nil:
		return ir.Type{Kind: ir.KindInvalid}
	case ir.LiteralExpr:
		return value.Type
	case ir.NameExpr:
		if symbol, ok := scope.Lookup(value.Name); ok {
			return symbol.Type
		}
		a.add(value.Span, CodeUnknownSymbol, "symbol not found: "+value.Name)
		return value.Type
	case ir.UnaryExpr:
		inner := a.exprType(scope, value.Expr)
		if value.Op == "!" {
			a.checkBoolean(ir.Span{}, inner)
			return ir.Type{Kind: ir.KindBoolean}
		}
		return inner
	case ir.BinaryExpr:
		left := a.exprType(scope, value.Left)
		right := a.exprType(scope, value.Right)
		return binaryResultType(value.Op, left, right)
	case ir.CallExpr:
		for _, arg := range value.Args {
			a.exprType(scope, arg)
		}
		return value.Type
	case ir.FieldExpr:
		return value.Type
	case ir.AddressExpr:
		a.exprType(scope, value.Expr)
		return value.Type
	case ir.CompositeLitExpr:
		for _, field := range value.Fields {
			a.exprType(scope, field.Value)
		}
		return value.Type
	default:
		return ir.Type{Kind: ir.KindInvalid}
	}
}

func (a *irAnalyzer) checkBoolean(span ir.Span, actual ir.Type) {
	if actual.Kind == ir.KindInvalid {
		return
	}
	if actual.Kind != ir.KindBoolean {
		a.add(span, CodeIncompatibleType, fmt.Sprintf("type incompatible: expected boolean, got %s", actual.GoName()))
	}
}

func (a *irAnalyzer) checkAssignable(span ir.Span, expected ir.Type, actual ir.Type) {
	if expected.Kind == ir.KindInvalid || actual.Kind == ir.KindInvalid {
		return
	}
	if !a.isAssignable(expected, actual) {
		a.add(span, CodeIncompatibleType, fmt.Sprintf("type incompatible: expected %s, got %s", expected.GoName(), actual.GoName()))
	}
}

func (a *irAnalyzer) isAssignable(expected ir.Type, actual ir.Type) bool {
	if typeEqual(expected, actual) {
		return true
	}
	if expected.Kind == ir.KindObject && actual.Kind == ir.KindPointer && actual.Elem != nil {
		if a.isInterface(expected.Name) {
			return true
		}
		return typeEqual(expected, *actual.Elem)
	}
	return false
}

func (a *irAnalyzer) isInterface(name string) bool {
	for _, iface := range a.file.Interfaces {
		if iface.Name == name || iface.Symbol == name {
			return true
		}
	}
	return false
}

func (a *irAnalyzer) add(span ir.Span, code string, message string) {
	if span.File == "" {
		span.File = a.file.Path
	}
	a.diagnostics = append(a.diagnostics, Diagnostic{
		File:    span.File,
		Line:    span.Line,
		Column:  span.Column,
		Code:    code,
		Message: message,
	})
}

func binaryResultType(op string, left ir.Type, right ir.Type) ir.Type {
	switch op {
	case "==", "!=", "<", "<=", ">", ">=", "&&", "||":
		return ir.Type{Kind: ir.KindBoolean}
	case "+":
		if left.Kind == ir.KindString || right.Kind == ir.KindString {
			return ir.Type{Kind: ir.KindString}
		}
	}
	if left.Kind == ir.KindDouble || right.Kind == ir.KindDouble {
		return ir.Type{Kind: ir.KindDouble}
	}
	if left.Kind == ir.KindInt && right.Kind == ir.KindInt {
		return ir.Type{Kind: ir.KindInt}
	}
	return ir.Type{Kind: ir.KindInvalid}
}
