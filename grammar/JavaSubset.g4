grammar JavaSubset;

compilationUnit
    : packageDecl? importDecl* typeDecl* EOF
    ;

packageDecl
    : 'package' qualifiedName ';'
    ;

importDecl
    : 'import' qualifiedName ('.' '*')? ';'
    ;

typeDecl
    : classDecl
    | interfaceDecl
    | ';'
    ;

classDecl
    : modifier* 'class' Identifier classBody
    ;

classBody
    : '{' classMember* '}'
    ;

interfaceDecl
    : modifier* 'interface' Identifier interfaceBody
    ;

interfaceBody
    : '{' interfaceMember* '}'
    ;

interfaceMember
    : interfaceMethodDecl
    | ';'
    ;

interfaceMethodDecl
    : modifier* typeTypeOrVoid Identifier formalParameters ';'
    ;

classMember
    : fieldDecl
    | methodDecl
    | constructorDecl
    | ';'
    ;

fieldDecl
    : modifier* typeType variableDeclarators ';'
    ;

methodDecl
    : modifier* typeTypeOrVoid Identifier formalParameters block
    ;

constructorDecl
    : modifier* Identifier formalParameters block
    ;

modifier
    : 'public'
    | 'private'
    | 'protected'
    | 'static'
    | 'final'
    ;

formalParameters
    : '(' formalParameterList? ')'
    ;

formalParameterList
    : formalParameter (',' formalParameter)*
    ;

formalParameter
    : 'final'* typeType Identifier
    ;

typeTypeOrVoid
    : typeType
    | 'void'
    ;

typeType
    : primitiveType dims?
    | qualifiedName dims?
    ;

primitiveType
    : 'boolean'
    | 'int'
    | 'double'
    | 'String'
    ;

dims
    : ('[' ']')+
    ;

variableDeclarators
    : variableDeclarator (',' variableDeclarator)*
    ;

variableDeclarator
    : Identifier ('=' expression)?
    ;

block
    : '{' blockStatement* '}'
    ;

blockStatement
    : localVariableDecl ';'
    | statement
    ;

localVariableDecl
    : 'final'* typeType variableDeclarators
    ;

statement
    : block
    | 'if' parExpression statement ('else' statement)?
    | 'while' parExpression statement
    | 'for' '(' forControl ')' statement
    | 'return' expression? ';'
    | statementExpression ';'
    | ';'
    ;

forControl
    : localVariableDecl? ';' expression? ';' forUpdate?
    ;

forUpdate
    : statementExpression (',' statementExpression)*
    ;

statementExpression
    : expression
    ;

parExpression
    : '(' expression ')'
    ;

expression
    : primary
    | expression '.' Identifier arguments?
    | prefix=('!' | '-') expression
    | expression bop=('*' | '/' | '%') expression
    | expression bop=('+' | '-') expression
    | expression bop=('<=' | '>=' | '>' | '<') expression
    | expression bop=('==' | '!=') expression
    | expression bop='&&' expression
    | expression bop='||' expression
    | <assoc=right> Identifier op=('=' | '+=' | '-=') expression
    ;

primary
    : literal
    | 'new' qualifiedName arguments
    | Identifier
    | qualifiedName arguments
    | '(' expression ')'
    ;

arguments
    : '(' expressionList? ')'
    ;

expressionList
    : expression (',' expression)*
    ;

literal
    : IntegerLiteral
    | FloatingPointLiteral
    | BooleanLiteral
    | StringLiteral
    | NullLiteral
    ;

qualifiedName
    : Identifier ('.' Identifier)*
    ;

BooleanLiteral
    : 'true'
    | 'false'
    ;

NullLiteral
    : 'null'
    ;

IntegerLiteral
    : DecimalIntegerLiteral
    ;

FloatingPointLiteral
    : [0-9]+ '.' [0-9]+
    ;

StringLiteral
    : '"' (~["\\\r\n] | EscapeSequence)* '"'
    ;

fragment DecimalIntegerLiteral
    : '0'
    | [1-9] [0-9]*
    ;

fragment EscapeSequence
    : '\\' ([btnfr"'\\] | 'u' HexDigit HexDigit HexDigit HexDigit)
    ;

fragment HexDigit
    : [0-9a-fA-F]
    ;

Identifier
    : JavaLetter JavaLetterOrDigit*
    ;

fragment JavaLetter
    : [a-zA-Z$_]
    ;

fragment JavaLetterOrDigit
    : [a-zA-Z0-9$_]
    ;

WS
    : [ \t\r\n\u000C]+ -> skip
    ;

LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

COMMENT
    : '/*' .*? '*/' -> skip
    ;
