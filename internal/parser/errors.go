package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Diagnostic struct {
	File    string
	Line    int
	Column  int
	Message string
}

func (d Diagnostic) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", d.File, d.Line, d.Column, d.Message)
}

type errorListener struct {
	file        string
	diagnostics []Diagnostic
}

func newErrorListener(file string) *errorListener {
	return &errorListener{file: file}
}

func (l *errorListener) Diagnostics() []Diagnostic {
	return append([]Diagnostic(nil), l.diagnostics...)
}

func (l *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line int,
	column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.diagnostics = append(l.diagnostics, Diagnostic{
		File:    l.file,
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

func (l *errorListener) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex int,
	stopIndex int,
	exact bool,
	ambigAlts *antlr.BitSet,
	configs *antlr.ATNConfigSet,
) {
}

func (l *errorListener) ReportAttemptingFullContext(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex int,
	stopIndex int,
	conflictingAlts *antlr.BitSet,
	configs *antlr.ATNConfigSet,
) {
}

func (l *errorListener) ReportContextSensitivity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex int,
	stopIndex int,
	prediction int,
	configs *antlr.ATNConfigSet,
) {
}
