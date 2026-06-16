package emitgo

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func Emit(file ir.File) ([]byte, error) {
	fset := token.NewFileSet()
	out := &ast.File{
		Name:  ast.NewIdent("main"),
		Decls: declarations(file),
	}
	var buffer bytes.Buffer
	if err := format.Node(&buffer, fset, out); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func declarations(file ir.File) []ast.Decl {
	decls := []ast.Decl{}
	if needsFmt(file) {
		decls = append(decls, &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{&ast.ImportSpec{
				Path: &ast.BasicLit{Kind: token.STRING, Value: "\"fmt\""},
			}},
		})
	}
	for _, class := range file.Classes {
		if class.NeedsStruct {
			decls = append(decls, structDecl(class))
		}
	}
	for _, iface := range file.Interfaces {
		decls = append(decls, interfaceDecl(iface))
	}
	for _, fn := range file.Funcs {
		decls = append(decls, funcDecl(fn))
	}
	for _, class := range file.Classes {
		for _, fn := range class.Methods {
			decls = append(decls, funcDecl(fn))
		}
	}
	return decls
}

func structDecl(class ir.Class) ast.Decl {
	fields := make([]*ast.Field, 0, len(class.Fields))
	for _, field := range class.Fields {
		fields = append(fields, &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(field.Name)},
			Type:  typeExpr(field.Type),
		})
	}
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent(class.Name),
			Type: &ast.StructType{Fields: &ast.FieldList{List: fields}},
		}},
	}
}

func interfaceDecl(iface ir.Interface) ast.Decl {
	methods := make([]*ast.Field, 0, len(iface.Methods))
	for _, method := range iface.Methods {
		methods = append(methods, &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(method.Name)},
			Type: &ast.FuncType{
				Params:  &ast.FieldList{List: params(method.Params)},
				Results: results(method.ReturnType),
			},
		})
	}
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent(iface.Name),
			Type: &ast.InterfaceType{Methods: &ast.FieldList{List: methods}},
		}},
	}
}

func funcDecl(fn ir.Func) *ast.FuncDecl {
	decl := &ast.FuncDecl{
		Name: ast.NewIdent(fn.Name),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params(fn.Params)},
			Results: results(fn.ReturnType),
		},
		Body: &ast.BlockStmt{List: stmts(fn.Body)},
	}
	if fn.Receiver != nil {
		decl.Recv = &ast.FieldList{List: []*ast.Field{{
			Names: []*ast.Ident{ast.NewIdent(fn.Receiver.Name)},
			Type:  typeExpr(fn.Receiver.Type),
		}}}
	}
	return decl
}

func params(in []ir.Param) []*ast.Field {
	out := make([]*ast.Field, 0, len(in))
	for _, param := range in {
		out = append(out, &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(param.Name)},
			Type:  typeExpr(param.Type),
		})
	}
	return out
}

func results(t ir.Type) *ast.FieldList {
	if t.Kind == ir.KindVoid {
		return nil
	}
	return &ast.FieldList{List: []*ast.Field{{Type: typeExpr(t)}}}
}

func needsFmt(file ir.File) bool {
	for _, fn := range file.Funcs {
		if bodyNeedsFmt(fn.Body) {
			return true
		}
	}
	for _, class := range file.Classes {
		for _, fn := range class.Methods {
			if bodyNeedsFmt(fn.Body) {
				return true
			}
		}
	}
	return false
}

func bodyNeedsFmt(stmtsIn []ir.Stmt) bool {
	for _, stmt := range stmtsIn {
		switch v := stmt.(type) {
		case ir.ExprStmt:
			if exprNeedsFmt(v.Expr) {
				return true
			}
		case ir.VarDeclStmt:
			if exprNeedsFmt(v.Value) {
				return true
			}
		case ir.AssignStmt:
			if exprNeedsFmt(v.Value) {
				return true
			}
		case ir.ReturnStmt:
			if exprNeedsFmt(v.Value) {
				return true
			}
		case ir.IfStmt:
			if exprNeedsFmt(v.Cond) || bodyNeedsFmt(v.Then) || bodyNeedsFmt(v.Else) {
				return true
			}
		case ir.WhileStmt:
			if exprNeedsFmt(v.Cond) || bodyNeedsFmt(v.Body) {
				return true
			}
		case ir.ForStmt:
			if exprNeedsFmt(v.Cond) || bodyNeedsFmt(v.Body) {
				return true
			}
		}
	}
	return false
}

func exprNeedsFmt(expression ir.Expr) bool {
	switch v := expression.(type) {
	case ir.CallExpr:
		if v.Target == "System.out" && v.Name == "println" {
			return true
		}
		for _, arg := range v.Args {
			if exprNeedsFmt(arg) {
				return true
			}
		}
	case ir.BinaryExpr:
		return exprNeedsFmt(v.Left) || exprNeedsFmt(v.Right)
	case ir.UnaryExpr:
		return exprNeedsFmt(v.Expr)
	case ir.AddressExpr:
		return exprNeedsFmt(v.Expr)
	case ir.CompositeLitExpr:
		for _, field := range v.Fields {
			if exprNeedsFmt(field.Value) {
				return true
			}
		}
	}
	return false
}
