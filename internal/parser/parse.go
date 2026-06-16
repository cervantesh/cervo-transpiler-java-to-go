package parser

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser/gen"
)

func ParseSource(fileName string, source string) (gen.ICompilationUnitContext, []Diagnostic) {
	input := antlr.NewInputStream(source)
	lexer := gen.NewJavaSubsetLexer(input)
	listener := newErrorListener(fileName)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := gen.NewJavaSubsetParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(listener)
	p.BuildParseTrees = true

	tree := p.CompilationUnit()
	return tree, listener.Diagnostics()
}
