// Code generated from C:/dev/cervo-transpiler-java-to-go/.worktrees/issue-5-modern-mvp/grammar/JavaSubset.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // JavaSubset
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by JavaSubsetParser.
type JavaSubsetVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JavaSubsetParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#packageDecl.
	VisitPackageDecl(ctx *PackageDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#classDecl.
	VisitClassDecl(ctx *ClassDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#classMember.
	VisitClassMember(ctx *ClassMemberContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#methodDecl.
	VisitMethodDecl(ctx *MethodDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#constructorDecl.
	VisitConstructorDecl(ctx *ConstructorDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#typeTypeOrVoid.
	VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#dims.
	VisitDims(ctx *DimsContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#localVariableDecl.
	VisitLocalVariableDecl(ctx *LocalVariableDeclContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#statementExpression.
	VisitStatementExpression(ctx *StatementExpressionContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by JavaSubsetParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}
}
