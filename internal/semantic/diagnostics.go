package semantic

import "fmt"

type Diagnostic struct {
	File    string
	Line    int
	Column  int
	Code    string
	Message string
}

func (d Diagnostic) Error() string {
	if d.Code == "" {
		return fmt.Sprintf("%s:%d:%d: %s", d.File, d.Line, d.Column, d.Message)
	}
	return fmt.Sprintf("%s:%d:%d: %s: %s", d.File, d.Line, d.Column, d.Code, d.Message)
}
