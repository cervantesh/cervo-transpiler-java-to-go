// Code generated from grammar/JavaSubset.g4 by ANTLR 4.13.1. DO NOT EDIT.

package gen // JavaSubset
import "github.com/antlr4-go/antlr/v4"

// JavaSubsetListener is a complete listener for a parse tree produced by JavaSubsetParser.
type JavaSubsetListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPackageDecl is called when entering the packageDecl production.
	EnterPackageDecl(c *PackageDeclContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterClassDecl is called when entering the classDecl production.
	EnterClassDecl(c *ClassDeclContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassMember is called when entering the classMember production.
	EnterClassMember(c *ClassMemberContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterMethodDecl is called when entering the methodDecl production.
	EnterMethodDecl(c *MethodDeclContext)

	// EnterConstructorDecl is called when entering the constructorDecl production.
	EnterConstructorDecl(c *ConstructorDeclContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterTypeTypeOrVoid is called when entering the typeTypeOrVoid production.
	EnterTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterDims is called when entering the dims production.
	EnterDims(c *DimsContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDecl is called when entering the localVariableDecl production.
	EnterLocalVariableDecl(c *LocalVariableDeclContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterStatementExpression is called when entering the statementExpression production.
	EnterStatementExpression(c *StatementExpressionContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPackageDecl is called when exiting the packageDecl production.
	ExitPackageDecl(c *PackageDeclContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitClassDecl is called when exiting the classDecl production.
	ExitClassDecl(c *ClassDeclContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassMember is called when exiting the classMember production.
	ExitClassMember(c *ClassMemberContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitMethodDecl is called when exiting the methodDecl production.
	ExitMethodDecl(c *MethodDeclContext)

	// ExitConstructorDecl is called when exiting the constructorDecl production.
	ExitConstructorDecl(c *ConstructorDeclContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitTypeTypeOrVoid is called when exiting the typeTypeOrVoid production.
	ExitTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitDims is called when exiting the dims production.
	ExitDims(c *DimsContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDecl is called when exiting the localVariableDecl production.
	ExitLocalVariableDecl(c *LocalVariableDeclContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitStatementExpression is called when exiting the statementExpression production.
	ExitStatementExpression(c *StatementExpressionContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)
}
