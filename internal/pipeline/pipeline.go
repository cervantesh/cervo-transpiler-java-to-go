package pipeline

import (
	"fmt"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/emitgo"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/lower"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser"
)

type Result struct {
	Code        []byte
	Diagnostics []error
}

func TranspileSource(fileName string, source string) Result {
	tree, syntax := parser.ParseSource(fileName, source)
	if len(syntax) != 0 {
		out := make([]error, 0, len(syntax))
		for _, diagnostic := range syntax {
			out = append(out, diagnostic)
		}
		return Result{Diagnostics: out}
	}

	irFile, semantic := lower.Lower(fileName, tree)
	if len(semantic) != 0 {
		out := make([]error, 0, len(semantic))
		for _, diagnostic := range semantic {
			out = append(out, diagnostic)
		}
		return Result{Diagnostics: out}
	}

	code, err := emitgo.Emit(irFile)
	if err != nil {
		return Result{Diagnostics: []error{fmt.Errorf("emit Go: %w", err)}}
	}
	return Result{Code: code}
}
