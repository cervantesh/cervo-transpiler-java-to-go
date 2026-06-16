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
	classFields map[string]ir.Type
	receiver    string
}

func Lower(fileName string, tree gen.ICompilationUnitContext) (ir.File, []semantic.Diagnostic) {
	l := &Lowerer{file: fileName}
	out := ir.File{Path: fileName}
	if tree.PackageDecl() != nil {
		out.PackageName = tree.PackageDecl().QualifiedName().GetText()
	}

	for _, typeDecl := range tree.AllTypeDecl() {
		classDecl := typeDecl.ClassDecl()
		if classDecl == nil {
			continue
		}
		class := ir.Class{
			Name:   classDecl.Identifier().GetText(),
			Symbol: qualifiedSymbol(out.PackageName, classDecl.Identifier().GetText()),
			Span:   l.span(classDecl.GetStart()),
		}
		for _, member := range classDecl.ClassBody().AllClassMember() {
			if member.FieldDecl() != nil {
				class.Fields = append(class.Fields, l.lowerFields(member.FieldDecl())...)
				class.NeedsStruct = true
			}
		}
		hasConstructor := false
		for _, member := range classDecl.ClassBody().AllClassMember() {
			if member.ConstructorDecl() != nil {
				hasConstructor = true
				class.NeedsStruct = true
				out.Funcs = append(out.Funcs, l.constructor(member.ConstructorDecl(), class))
				continue
			}
			method := member.MethodDecl()
			if method == nil {
				continue
			}
			fn := l.method(method, class)
			if fn.Static {
				out.Funcs = append(out.Funcs, fn)
			} else {
				class.NeedsStruct = true
				class.Methods = append(class.Methods, fn)
			}
		}
		if class.NeedsStruct && !hasConstructor {
			out.Funcs = append(out.Funcs, l.defaultConstructor(class))
		}
		out.Classes = append(out.Classes, class)
	}

	return out, l.diagnostics
}

func (l *Lowerer) method(ctx gen.IMethodDeclContext, class ir.Class) ir.Func {
	name := ctx.Identifier().GetText()
	fn := ir.Func{
		Name:       name,
		Symbol:     qualifiedSymbol(class.Symbol, name),
		Span:       l.span(ctx.GetStart()),
		ReturnType: MapJavaType(ctx.TypeTypeOrVoid().GetText()),
		Static:     hasModifier(ctx.AllModifier(), "static"),
	}
	if !fn.Static {
		receiverName := receiverName(class.Name)
		receiverType := pointerToObject(class.Name)
		fn.Receiver = &ir.Param{Name: receiverName, Type: receiverType, Span: l.span(ctx.GetStart())}
	}
	if params := ctx.FormalParameters().FormalParameterList(); params != nil && fn.Name != "main" {
		for _, param := range params.AllFormalParameter() {
			fn.Params = append(fn.Params, ir.Param{
				Name: param.Identifier().GetText(),
				Type: MapJavaType(param.TypeType().GetText()),
				Span: l.span(param.GetStart()),
			})
		}
	}
	previousFields := l.classFields
	previousReceiver := l.receiver
	l.classFields = fieldMap(class.Fields)
	if fn.Receiver != nil {
		l.receiver = fn.Receiver.Name
	} else {
		l.receiver = ""
	}
	fn.Body = l.block(ctx.Block())
	l.classFields = previousFields
	l.receiver = previousReceiver
	return fn
}

func (l *Lowerer) lowerFields(ctx gen.IFieldDeclContext) []ir.Field {
	typ := MapJavaType(ctx.TypeType().GetText())
	out := []ir.Field{}
	for _, decl := range ctx.VariableDeclarators().AllVariableDeclarator() {
		out = append(out, ir.Field{Name: decl.Identifier().GetText(), Type: typ, Span: l.span(decl.GetStart())})
	}
	return out
}

func (l *Lowerer) constructor(ctx gen.IConstructorDeclContext, class ir.Class) ir.Func {
	params := []ir.Param{}
	if list := ctx.FormalParameters().FormalParameterList(); list != nil {
		for _, param := range list.AllFormalParameter() {
			params = append(params, ir.Param{Name: param.Identifier().GetText(), Type: MapJavaType(param.TypeType().GetText()), Span: l.span(param.GetStart())})
		}
	}
	returnType := pointerToObject(class.Name)
	return ir.Func{
		Name:       "New" + class.Name,
		Symbol:     qualifiedSymbol(class.Symbol, "New"+class.Name),
		Span:       l.span(ctx.GetStart()),
		Params:     params,
		ReturnType: returnType,
		Static:     true,
		Body: []ir.Stmt{ir.ReturnStmt{
			Span:  l.span(ctx.GetStart()),
			Value: l.constructorReturn(class, params),
		}},
	}
}

func (l *Lowerer) defaultConstructor(class ir.Class) ir.Func {
	return ir.Func{
		Name:       "New" + class.Name,
		Symbol:     qualifiedSymbol(class.Symbol, "New"+class.Name),
		Span:       class.Span,
		ReturnType: pointerToObject(class.Name),
		Static:     true,
		Body: []ir.Stmt{ir.ReturnStmt{
			Span: class.Span,
			Value: ir.AddressExpr{
				Type: pointerToObject(class.Name),
				Expr: ir.CompositeLitExpr{
					TypeName: class.Name,
					Type:     ir.Type{Kind: ir.KindObject, Name: class.Name},
				},
			},
		}},
	}
}

func (l *Lowerer) constructorReturn(class ir.Class, params []ir.Param) ir.Expr {
	paramByName := map[string]ir.Param{}
	for _, param := range params {
		paramByName[param.Name] = param
	}
	fields := []ir.KeyValueExpr{}
	for _, field := range class.Fields {
		if param, ok := paramByName[field.Name]; ok {
			fields = append(fields, ir.KeyValueExpr{Key: field.Name, Value: ir.NameExpr{Name: param.Name, Type: param.Type, Span: param.Span}})
		}
	}
	objectType := ir.Type{Kind: ir.KindObject, Name: class.Name}
	return ir.AddressExpr{
		Type: pointerToObject(class.Name),
		Expr: ir.CompositeLitExpr{
			TypeName: class.Name,
			Fields:   fields,
			Type:     objectType,
		},
	}
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
			Span:  l.span(decl.GetStart()),
			Value: value,
		})
	}
	return out
}

func (l *Lowerer) statement(ctx gen.IStatementContext) ir.Stmt {
	start := ""
	if ctx.GetStart() != nil {
		start = ctx.GetStart().GetText()
	}
	if ctx.Block() != nil {
		l.unsupported(ctx.GetStart(), "JTG1020", "standalone blocks", "Inline the block or add block-scope lowering before transpiling nested standalone blocks.")
		return nil
	}
	if ctx.ParExpression() != nil && start == "if" {
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
			Span: l.span(ctx.GetStart()),
			Cond: l.expression(ctx.ParExpression().Expression()),
			Then: thenBody,
			Else: elseBody,
		}
	}
	if ctx.ParExpression() != nil && start == "while" {
		children := ctx.AllStatement()
		body := []ir.Stmt{}
		if len(children) > 0 {
			body = l.statementBody(children[0])
		}
		return ir.WhileStmt{
			Span: l.span(ctx.GetStart()),
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
	if ctx.Expression() != nil && start == "return" {
		return ir.ReturnStmt{Span: l.span(ctx.GetStart()), Value: l.expression(ctx.Expression())}
	}
	if start == "return" {
		return ir.ReturnStmt{Span: l.span(ctx.GetStart())}
	}
	if ctx.StatementExpression() != nil {
		expression := ctx.StatementExpression().Expression()
		if assign := l.assignment(expression); assign != nil {
			return assign
		}
		return ir.ExprStmt{Span: l.span(ctx.GetStart()), Expr: l.expression(expression)}
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
	return ir.ForStmt{Span: l.span(ctx.GetStart()), Init: init, Cond: cond, Post: post, Body: body}
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
		Span:  l.span(ctx.GetStart()),
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
			l.unsupported(ctx.GetStart(), "JTG1019", "assignment expressions", "Move the assignment to a standalone statement before transpiling.")
			return ir.NameExpr{Name: assign.Name, Type: ir.Type{Kind: ir.KindInvalid}, Span: l.span(ctx.GetStart())}
		}
	}
	if ctx.GetPrefix() != nil {
		child := l.expression(ctx.Expression(0))
		op := ctx.GetPrefix().GetText()
		return ir.UnaryExpr{Op: op, Expr: child, Type: unaryType(op, child)}
	}
	if ctx.Identifier() != nil && len(ctx.AllExpression()) > 0 {
		children := ctx.AllExpression()
		if ctx.Arguments() != nil {
			return ir.CallExpr{Target: children[0].GetText(), Name: ctx.Identifier().GetText(), Args: l.arguments(ctx.Arguments()), Type: ir.Type{Kind: ir.KindInvalid}}
		}
		return ir.NameExpr{Name: ctx.GetText(), Type: ir.Type{Kind: ir.KindInvalid}, Span: l.span(ctx.GetStart())}
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
			return ir.CallExpr{Target: target, Name: ctx.Identifier().GetText(), Args: args, Type: ir.Type{Kind: ir.KindInvalid}}
		}
		if len(children) == 2 {
			left := l.expression(children[0])
			right := l.expression(children[1])
			return ir.BinaryExpr{Left: left, Op: op, Right: right, Type: binaryType(op, left, right)}
		}
	}
	return ir.NameExpr{Name: ctx.GetText(), Type: ir.Type{Kind: ir.KindInvalid}, Span: l.span(ctx.GetStart())}
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
		if l.receiver != "" {
			if typ, ok := l.classFields[identifier.GetText()]; ok {
				return ir.FieldExpr{Target: l.receiver, Field: identifier.GetText(), Type: typ}
			}
		}
		return ir.NameExpr{Name: identifier.GetText(), Type: ir.Type{Kind: ir.KindInvalid}, Span: l.span(ctx.GetStart())}
	}
	if ctx.QualifiedName() != nil && ctx.Arguments() != nil {
		name := ctx.QualifiedName().GetText()
		lastDot := strings.LastIndex(name, ".")
		target := ""
		callName := name
		if ctx.GetStart() != nil && ctx.GetStart().GetText() == "new" {
			callName = "New" + name
			if lastDot >= 0 {
				callName = "New" + name[lastDot+1:]
			}
			return ir.CallExpr{Target: "", Name: callName, Args: l.arguments(ctx.Arguments()), Type: pointerToObject(strings.TrimPrefix(callName, "New"))}
		}
		if lastDot >= 0 {
			target = name[:lastDot]
			callName = name[lastDot+1:]
		}
		return ir.CallExpr{Target: target, Name: callName, Args: l.arguments(ctx.Arguments()), Type: ir.Type{Kind: ir.KindInvalid}}
	}
	if ctx.Expression() != nil {
		return l.expression(ctx.Expression())
	}
	return ir.NameExpr{Name: ctx.GetText(), Type: ir.Type{Kind: ir.KindInvalid}, Span: l.span(ctx.GetStart())}
}

func (l *Lowerer) span(token antlr.Token) ir.Span {
	if token == nil {
		return ir.Span{File: l.file}
	}
	return ir.Span{File: l.file, Line: token.GetLine(), Column: token.GetColumn()}
}

func qualifiedSymbol(packageName string, name string) string {
	if packageName == "" {
		return name
	}
	return packageName + "." + name
}

func unaryType(op string, child ir.Expr) ir.Type {
	if op == "!" {
		return ir.Type{Kind: ir.KindBoolean}
	}
	return exprType(child)
}

func binaryType(op string, left ir.Expr, right ir.Expr) ir.Type {
	switch op {
	case "==", "!=", "<", "<=", ">", ">=", "&&", "||":
		return ir.Type{Kind: ir.KindBoolean}
	case "+":
		if exprType(left).Kind == ir.KindString || exprType(right).Kind == ir.KindString {
			return ir.Type{Kind: ir.KindString}
		}
	}
	if exprType(left).Kind == ir.KindDouble || exprType(right).Kind == ir.KindDouble {
		return ir.Type{Kind: ir.KindDouble}
	}
	if exprType(left).Kind == ir.KindInt && exprType(right).Kind == ir.KindInt {
		return ir.Type{Kind: ir.KindInt}
	}
	return ir.Type{Kind: ir.KindInvalid}
}

func exprType(expression ir.Expr) ir.Type {
	switch value := expression.(type) {
	case ir.LiteralExpr:
		return value.Type
	case ir.NameExpr:
		return value.Type
	case ir.BinaryExpr:
		return value.Type
	case ir.UnaryExpr:
		return value.Type
	case ir.CallExpr:
		return value.Type
	case ir.FieldExpr:
		return value.Type
	case ir.AddressExpr:
		return value.Type
	case ir.CompositeLitExpr:
		return value.Type
	default:
		return ir.Type{Kind: ir.KindInvalid}
	}
}

func fieldMap(fields []ir.Field) map[string]ir.Type {
	out := map[string]ir.Type{}
	for _, field := range fields {
		out[field.Name] = field.Type
	}
	return out
}

func pointerToObject(name string) ir.Type {
	objectType := ir.Type{Kind: ir.KindObject, Name: name}
	return ir.Type{Kind: ir.KindPointer, Elem: &objectType}
}

func receiverName(className string) string {
	if className == "" {
		return "receiver"
	}
	name := strings.ToLower(className[:1])
	if name == "_" || name == "$" {
		return "receiver"
	}
	return name
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
