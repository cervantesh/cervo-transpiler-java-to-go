package semantic

import "fmt"

type Diagnostic struct {
	File           string
	Line           int
	Column         int
	Code           string
	Message        string
	Recommendation string
}

func (d Diagnostic) Error() string {
	message := ""
	if d.Code == "" {
		message = fmt.Sprintf("%s:%d:%d: %s", d.File, d.Line, d.Column, d.Message)
	} else {
		message = fmt.Sprintf("%s:%d:%d: %s: %s", d.File, d.Line, d.Column, d.Code, d.Message)
	}
	if d.Recommendation != "" {
		message += "\n  recommendation: " + d.Recommendation
	}
	return message
}
