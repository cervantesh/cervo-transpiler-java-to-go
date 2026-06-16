package emitgo

import (
	"go/ast"
	"go/token"
	"strconv"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func expr(e ir.Expr) ast.Expr {
	switch v := e.(type) {
	case nil:
		return ast.NewIdent("nil")
	case ir.LiteralExpr:
		return literal(v)
	case ir.NameExpr:
		return ast.NewIdent(v.Name)
	case ir.UnaryExpr:
		return &ast.UnaryExpr{Op: tokenForUnary(v.Op), X: expr(v.Expr)}
	case ir.BinaryExpr:
		return &ast.BinaryExpr{X: expr(v.Left), Op: tokenForBinary(v.Op), Y: expr(v.Right)}
	case ir.CallExpr:
		return call(v)
	case ir.FieldExpr:
		return &ast.SelectorExpr{X: ast.NewIdent(v.Target), Sel: ast.NewIdent(v.Field)}
	case ir.AddressExpr:
		return &ast.UnaryExpr{Op: token.AND, X: expr(v.Expr)}
	case ir.CompositeLitExpr:
		return compositeLit(v)
	default:
		return ast.NewIdent("nil")
	}
}

func literal(v ir.LiteralExpr) ast.Expr {
	switch v.Type.Kind {
	case ir.KindString:
		if len(v.Value) >= 2 && v.Value[0] == '"' {
			return &ast.BasicLit{Kind: token.STRING, Value: v.Value}
		}
		return &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(v.Value)}
	case ir.KindBoolean:
		return ast.NewIdent(v.Value)
	case ir.KindDouble:
		return &ast.BasicLit{Kind: token.FLOAT, Value: v.Value}
	default:
		return &ast.BasicLit{Kind: token.INT, Value: v.Value}
	}
}

func call(v ir.CallExpr) ast.Expr {
	target := v.Target
	name := v.Name
	if target == "System.out" && name == "println" {
		target = "fmt"
		name = "Println"
	}

	args := make([]ast.Expr, 0, len(v.Args))
	for _, arg := range v.Args {
		args = append(args, expr(arg))
	}

	if target == "" {
		return &ast.CallExpr{Fun: ast.NewIdent(name), Args: args}
	}
	return &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: ast.NewIdent(target), Sel: ast.NewIdent(name)},
		Args: args,
	}
}

func compositeLit(v ir.CompositeLitExpr) ast.Expr {
	elts := make([]ast.Expr, 0, len(v.Fields))
	for _, field := range v.Fields {
		elts = append(elts, &ast.KeyValueExpr{Key: ast.NewIdent(field.Key), Value: expr(field.Value)})
	}
	return &ast.CompositeLit{Type: ast.NewIdent(v.TypeName), Elts: elts}
}

func tokenForUnary(op string) token.Token {
	switch op {
	case "!":
		return token.NOT
	case "-":
		return token.SUB
	default:
		return token.ILLEGAL
	}
}

func tokenForBinary(op string) token.Token {
	switch op {
	case "+":
		return token.ADD
	case "-":
		return token.SUB
	case "*":
		return token.MUL
	case "/":
		return token.QUO
	case "%":
		return token.REM
	case "==":
		return token.EQL
	case "!=":
		return token.NEQ
	case "<":
		return token.LSS
	case "<=":
		return token.LEQ
	case ">":
		return token.GTR
	case ">=":
		return token.GEQ
	case "&&":
		return token.LAND
	case "||":
		return token.LOR
	default:
		return token.ILLEGAL
	}
}
