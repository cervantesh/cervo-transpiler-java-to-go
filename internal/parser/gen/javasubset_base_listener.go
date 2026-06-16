// Code generated from C:/dev/cervo-transpiler-java-to-go/.worktrees/issue-5-modern-mvp/grammar/JavaSubset.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // JavaSubset
import "github.com/antlr4-go/antlr/v4"

// BaseJavaSubsetListener is a complete listener for a parse tree produced by JavaSubsetParser.
type BaseJavaSubsetListener struct{}

var _ JavaSubsetListener = &BaseJavaSubsetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavaSubsetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavaSubsetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavaSubsetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavaSubsetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseJavaSubsetListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseJavaSubsetListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPackageDecl is called when production packageDecl is entered.
func (s *BaseJavaSubsetListener) EnterPackageDecl(ctx *PackageDeclContext) {}

// ExitPackageDecl is called when production packageDecl is exited.
func (s *BaseJavaSubsetListener) ExitPackageDecl(ctx *PackageDeclContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseJavaSubsetListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseJavaSubsetListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseJavaSubsetListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseJavaSubsetListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterClassDecl is called when production classDecl is entered.
func (s *BaseJavaSubsetListener) EnterClassDecl(ctx *ClassDeclContext) {}

// ExitClassDecl is called when production classDecl is exited.
func (s *BaseJavaSubsetListener) ExitClassDecl(ctx *ClassDeclContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseJavaSubsetListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseJavaSubsetListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassMember is called when production classMember is entered.
func (s *BaseJavaSubsetListener) EnterClassMember(ctx *ClassMemberContext) {}

// ExitClassMember is called when production classMember is exited.
func (s *BaseJavaSubsetListener) ExitClassMember(ctx *ClassMemberContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseJavaSubsetListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseJavaSubsetListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseJavaSubsetListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseJavaSubsetListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterConstructorDecl is called when production constructorDecl is entered.
func (s *BaseJavaSubsetListener) EnterConstructorDecl(ctx *ConstructorDeclContext) {}

// ExitConstructorDecl is called when production constructorDecl is exited.
func (s *BaseJavaSubsetListener) ExitConstructorDecl(ctx *ConstructorDeclContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseJavaSubsetListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseJavaSubsetListener) ExitModifier(ctx *ModifierContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseJavaSubsetListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseJavaSubsetListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseJavaSubsetListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseJavaSubsetListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseJavaSubsetListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseJavaSubsetListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterTypeTypeOrVoid is called when production typeTypeOrVoid is entered.
func (s *BaseJavaSubsetListener) EnterTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// ExitTypeTypeOrVoid is called when production typeTypeOrVoid is exited.
func (s *BaseJavaSubsetListener) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BaseJavaSubsetListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BaseJavaSubsetListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseJavaSubsetListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseJavaSubsetListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterDims is called when production dims is entered.
func (s *BaseJavaSubsetListener) EnterDims(ctx *DimsContext) {}

// ExitDims is called when production dims is exited.
func (s *BaseJavaSubsetListener) ExitDims(ctx *DimsContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseJavaSubsetListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseJavaSubsetListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseJavaSubsetListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseJavaSubsetListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJavaSubsetListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJavaSubsetListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseJavaSubsetListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseJavaSubsetListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDecl is called when production localVariableDecl is entered.
func (s *BaseJavaSubsetListener) EnterLocalVariableDecl(ctx *LocalVariableDeclContext) {}

// ExitLocalVariableDecl is called when production localVariableDecl is exited.
func (s *BaseJavaSubsetListener) ExitLocalVariableDecl(ctx *LocalVariableDeclContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJavaSubsetListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJavaSubsetListener) ExitStatement(ctx *StatementContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseJavaSubsetListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseJavaSubsetListener) ExitForControl(ctx *ForControlContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseJavaSubsetListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseJavaSubsetListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseJavaSubsetListener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseJavaSubsetListener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseJavaSubsetListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseJavaSubsetListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJavaSubsetListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJavaSubsetListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseJavaSubsetListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseJavaSubsetListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseJavaSubsetListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseJavaSubsetListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseJavaSubsetListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseJavaSubsetListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJavaSubsetListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJavaSubsetListener) ExitLiteral(ctx *LiteralContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseJavaSubsetListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseJavaSubsetListener) ExitQualifiedName(ctx *QualifiedNameContext) {}
