package lower

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser/gen"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/semantic"
)

type Lowerer struct {
	file        string
	diagnostics []semantic.Diagnostic
}

func Lower(fileName string, tree gen.ICompilationUnitContext) (ir.File, []semantic.Diagnostic) {
	l := &Lowerer{file: fileName}
	out := ir.File{}
	if tree.PackageDecl() != nil {
		out.PackageName = tree.PackageDecl().QualifiedName().GetText()
	}

	for _, typeDecl := range tree.AllTypeDecl() {
		classDecl := typeDecl.ClassDecl()
		if classDecl == nil {
			continue
		}
		class := ir.Class{Name: classDecl.Identifier().GetText()}
		for _, member := range classDecl.ClassBody().AllClassMember() {
			if member.FieldDecl() != nil {
				l.unsupported(member.FieldDecl().GetStart(), "JTG1016", "class fields", "Add struct field lowering before transpiling Java fields.")
				continue
			}
			if member.ConstructorDecl() != nil {
				l.unsupported(member.ConstructorDecl().GetStart(), "JTG1006", "constructors", "Add constructor lowering before transpiling Java constructors.")
				continue
			}
			method := member.MethodDecl()
			if method == nil {
				continue
			}
			fn := l.method(method)
			if fn.Static {
				out.Funcs = append(out.Funcs, fn)
			} else {
				class.Methods = append(class.Methods, fn)
			}
		}
		if len(class.Methods) > 0 {
			out.Classes = append(out.Classes, class)
		}
	}

	return out, l.diagnostics
}

func (l *Lowerer) method(ctx gen.IMethodDeclContext) ir.Func {
	fn := ir.Func{
		Name:       ctx.Identifier().GetText(),
		ReturnType: MapJavaType(ctx.TypeTypeOrVoid().GetText()),
		Static:     hasModifier(ctx.AllModifier(), "static"),
	}
	if params := ctx.FormalParameters().FormalParameterList(); params != nil && fn.Name != "main" {
		for _, param := range params.AllFormalParameter() {
			fn.Params = append(fn.Params, ir.Param{
				Name: param.Identifier().GetText(),
				Type: MapJavaType(param.TypeType().GetText()),
			})
		}
	}
	fn.Body = l.block(ctx.Block())
	return fn
}

func hasModifier(modifiers []gen.IModifierContext, name string) bool {
	for _, modifier := range modifiers {
		if modifier.GetText() == name {
			return true
		}
	}
	return false
}

func (l *Lowerer) block(ctx gen.IBlockContext) []ir.Stmt {
	out := []ir.Stmt{}
	for _, item := range ctx.AllBlockStatement() {
		if local := item.LocalVariableDecl(); local != nil {
			out = append(out, l.localVar(local)...)
			continue
		}
		if stmtCtx := item.Statement(); stmtCtx != nil {
			if stmt := l.statement(stmtCtx); stmt != nil {
				out = append(out, stmt)
			}
		}
	}
	return out
}

func (l *Lowerer) localVar(ctx gen.ILocalVariableDeclContext) []ir.Stmt {
	typ := MapJavaType(ctx.TypeType().GetText())
	out := []ir.Stmt{}
	for _, decl := range ctx.VariableDeclarators().AllVariableDeclarator() {
		var value ir.Expr
		if decl.Expression() != nil {
			value = l.expression(decl.Expression())
		}
		out = append(out, ir.VarDeclStmt{
			Name:  decl.Identifier().GetText(),
			Type:  typ,
			Value: value,
		})
	}
	return out
}

func (l *Lowerer) statement(ctx gen.IStatementContext) ir.Stmt {
	text := ctx.GetText()
	if ctx.Block() != nil && !strings.HasPrefix(text, "if") && !strings.HasPrefix(text, "while") && !strings.HasPrefix(text, "for") {
		return ir.ExprStmt{Expr: ir.CallExpr{Name: "__unsupported_block"}}
	}
	if ctx.ParExpression() != nil && strings.HasPrefix(text, "if") {
		children := ctx.AllStatement()
		thenBody := []ir.Stmt{}
		elseBody := []ir.Stmt{}
		if len(children) > 0 {
			thenBody = l.statementBody(children[0])
		}
		if len(children) > 1 {
			elseBody = l.statementBody(children[1])
		}
		return ir.IfStmt{
			Cond: l.expression(ctx.ParExpression().Expression()),
			Then: thenBody,
			Else: elseBody,
		}
	}
	if ctx.ParExpression() != nil && strings.HasPrefix(text, "while") {
		children := ctx.AllStatement()
		body := []ir.Stmt{}
		if len(children) > 0 {
			body = l.statementBody(children[0])
		}
		return ir.WhileStmt{
			Cond: l.expression(ctx.ParExpression().Expression()),
			Body: body,
		}
	}
	if ctx.ForControl() != nil {
		children := ctx.AllStatement()
		body := []ir.Stmt{}
		if len(children) > 0 {
			body = l.statementBody(children[0])
		}
		return l.forStmt(ctx.ForControl(), body)
	}
	if ctx.Expression() != nil && strings.HasPrefix(text, "return") {
		return ir.ReturnStmt{Value: l.expression(ctx.Expression())}
	}
	if strings.HasPrefix(text, "return") {
		return ir.ReturnStmt{}
	}
	if ctx.StatementExpression() != nil {
		expression := ctx.StatementExpression().Expression()
		if assign := l.assignment(expression); assign != nil {
			return assign
		}
		return ir.ExprStmt{Expr: l.expression(expression)}
	}
	return nil
}

func (l *Lowerer) statementBody(ctx gen.IStatementContext) []ir.Stmt {
	if ctx.Block() != nil {
		return l.block(ctx.Block())
	}
	if stmt := l.statement(ctx); stmt != nil {
		return []ir.Stmt{stmt}
	}
	return []ir.Stmt{}
}

func (l *Lowerer) forStmt(ctx gen.IForControlContext, body []ir.Stmt) ir.Stmt {
	var init ir.Stmt
	if local := ctx.LocalVariableDecl(); local != nil {
		vars := l.localVar(local)
		if len(vars) > 0 {
			init = vars[0]
		}
	}
	var cond ir.Expr
	if ctx.Expression() != nil {
		cond = l.expression(ctx.Expression())
	}
	var post ir.Stmt
	if update := ctx.ForUpdate(); update != nil {
		items := update.AllStatementExpression()
		if len(items) > 0 {
			expression := items[0].Expression()
			if assign := l.assignment(expression); assign != nil {
				post = assign
			} else {
				post = ir.ExprStmt{Expr: l.expression(expression)}
			}
		}
	}
	return ir.ForStmt{Init: init, Cond: cond, Post: post, Body: body}
}

func (l *Lowerer) assignment(ctx gen.IExpressionContext) ir.Stmt {
	if ctx.GetOp() == nil || ctx.Identifier() == nil {
		return nil
	}
	children := ctx.AllExpression()
	if len(children) != 1 {
		return nil
	}
	return ir.AssignStmt{
		Name:  ctx.Identifier().GetText(),
		Op:    ctx.GetOp().GetText(),
		Value: l.expression(children[0]),
	}
}

func (l *Lowerer) expression(ctx gen.IExpressionContext) ir.Expr {
	if ctx == nil {
		return nil
	}
	if primary := ctx.Primary(); primary != nil {
		return l.primary(primary)
	}
	if ctx.GetOp() != nil {
		if stmt := l.assignment(ctx); stmt != nil {
			assign := stmt.(ir.AssignStmt)
			return ir.CallExpr{Name: "__assign_" + assign.Op, Args: []ir.Expr{ir.NameExpr{Name: assign.Name}, assign.Value}}
		}
	}
	if ctx.GetPrefix() != nil {
		return ir.UnaryExpr{Op: ctx.GetPrefix().GetText(), Expr: l.expression(ctx.Expression(0))}
	}
	if ctx.Identifier() != nil && len(ctx.AllExpression()) > 0 {
		children := ctx.AllExpression()
		if ctx.Arguments() != nil {
			return ir.CallExpr{Target: children[0].GetText(), Name: ctx.Identifier().GetText(), Args: l.arguments(ctx.Arguments())}
		}
		return ir.NameExpr{Name: ctx.GetText()}
	}
	if ctx.GetBop() != nil {
		op := ctx.GetBop().GetText()
		children := ctx.AllExpression()
		if op == "." {
			target := ""
			if len(children) > 0 {
				target = children[0].GetText()
			}
			args := []ir.Expr{}
			if ctx.Arguments() != nil {
				args = l.arguments(ctx.Arguments())
			}
			return ir.CallExpr{Target: target, Name: ctx.Identifier().GetText(), Args: args}
		}
		if len(children) == 2 {
			return ir.BinaryExpr{Left: l.expression(children[0]), Op: op, Right: l.expression(children[1])}
		}
	}
	return ir.NameExpr{Name: ctx.GetText()}
}

func (l *Lowerer) primary(ctx gen.IPrimaryContext) ir.Expr {
	if literal := ctx.Literal(); literal != nil {
		text := literal.GetText()
		switch {
		case strings.HasPrefix(text, "\""):
			return ir.LiteralExpr{Value: text, Type: ir.Type{Kind: ir.KindString}}
		case text == "true" || text == "false":
			return ir.LiteralExpr{Value: text, Type: ir.Type{Kind: ir.KindBoolean}}
		case strings.Contains(text, "."):
			return ir.LiteralExpr{Value: text, Type: ir.Type{Kind: ir.KindDouble}}
		default:
			return ir.LiteralExpr{Value: text, Type: ir.Type{Kind: ir.KindInt}}
		}
	}
	if identifier := ctx.Identifier(); identifier != nil {
		return ir.NameExpr{Name: identifier.GetText()}
	}
	if ctx.QualifiedName() != nil && ctx.Arguments() != nil {
		name := ctx.QualifiedName().GetText()
		lastDot := strings.LastIndex(name, ".")
		target := ""
		callName := name
		if lastDot >= 0 {
			target = name[:lastDot]
			callName = name[lastDot+1:]
		}
		return ir.CallExpr{Target: target, Name: callName, Args: l.arguments(ctx.Arguments())}
	}
	if ctx.Expression() != nil {
		return l.expression(ctx.Expression())
	}
	return ir.NameExpr{Name: ctx.GetText()}
}

func (l *Lowerer) arguments(ctx gen.IArgumentsContext) []ir.Expr {
	if ctx == nil || ctx.ExpressionList() == nil {
		return nil
	}
	args := []ir.Expr{}
	for _, expression := range ctx.ExpressionList().AllExpression() {
		args = append(args, l.expression(expression))
	}
	return args
}

func (l *Lowerer) unsupported(token antlr.Token, code string, feature string, recommendation string) {
	message := "unsupported feature: " + feature
	if token == nil {
		l.diagnostics = append(l.diagnostics, semantic.Diagnostic{File: l.file, Code: code, Message: message, Recommendation: recommendation})
		return
	}
	l.diagnostics = append(l.diagnostics, semantic.Diagnostic{
		File:           l.file,
		Line:           token.GetLine(),
		Column:         token.GetColumn(),
		Code:           code,
		Message:        message,
		Recommendation: recommendation,
	})
}
