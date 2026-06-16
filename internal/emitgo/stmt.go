package emitgo

import (
	"go/ast"
	"go/token"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func stmts(in []ir.Stmt) []ast.Stmt {
	out := make([]ast.Stmt, 0, len(in))
	for _, stmt := range in {
		if stmt == nil {
			continue
		}
		out = append(out, stmtNode(stmt))
	}
	return out
}

func stmtNode(s ir.Stmt) ast.Stmt {
	switch v := s.(type) {
	case ir.VarDeclStmt:
		if v.Value != nil {
			return &ast.AssignStmt{Lhs: []ast.Expr{ast.NewIdent(v.Name)}, Tok: token.DEFINE, Rhs: []ast.Expr{expr(v.Value)}}
		}
		return &ast.DeclStmt{Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{&ast.ValueSpec{
				Names: []*ast.Ident{ast.NewIdent(v.Name)},
				Type:  typeExpr(v.Type),
			}},
		}}
	case ir.AssignStmt:
		return &ast.AssignStmt{Lhs: []ast.Expr{ast.NewIdent(v.Name)}, Tok: tokenForAssign(v.Op), Rhs: []ast.Expr{expr(v.Value)}}
	case ir.ExprStmt:
		return &ast.ExprStmt{X: expr(v.Expr)}
	case ir.ReturnStmt:
		if v.Value == nil {
			return &ast.ReturnStmt{}
		}
		return &ast.ReturnStmt{Results: []ast.Expr{expr(v.Value)}}
	case ir.IfStmt:
		return &ast.IfStmt{Cond: expr(v.Cond), Body: &ast.BlockStmt{List: stmts(v.Then)}, Else: elseNode(v.Else)}
	case ir.WhileStmt:
		return &ast.ForStmt{Cond: expr(v.Cond), Body: &ast.BlockStmt{List: stmts(v.Body)}}
	case ir.ForStmt:
		return &ast.ForStmt{Init: simpleStmt(v.Init), Cond: expr(v.Cond), Post: simpleStmt(v.Post), Body: &ast.BlockStmt{List: stmts(v.Body)}}
	default:
		return &ast.EmptyStmt{}
	}
}

func simpleStmt(stmt ir.Stmt) ast.Stmt {
	if stmt == nil {
		return nil
	}
	return stmtNode(stmt)
}

func elseNode(in []ir.Stmt) ast.Stmt {
	if len(in) == 0 {
		return nil
	}
	return &ast.BlockStmt{List: stmts(in)}
}

func tokenForAssign(op string) token.Token {
	switch op {
	case "=":
		return token.ASSIGN
	case "+=":
		return token.ADD_ASSIGN
	case "-=":
		return token.SUB_ASSIGN
	default:
		return token.ASSIGN
	}
}

func typeExpr(t ir.Type) ast.Expr {
	if t.Kind == ir.KindArray && t.Elem != nil {
		return &ast.ArrayType{Elt: typeExpr(*t.Elem)}
	}
	if t.Kind == ir.KindPointer && t.Elem != nil {
		return &ast.StarExpr{X: typeExpr(*t.Elem)}
	}
	return ast.NewIdent(t.GoName())
}
