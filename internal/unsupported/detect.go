package unsupported

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/semantic"
)

type line struct {
	number int
	text   string
}

type keywordRule struct {
	keyword        string
	code           string
	feature        string
	recommendation string
}

var keywordRules = []keywordRule{
	{"package", "JTG1001", "package declarations", "Remove the package declaration for the current subset or add package-to-module mapping."},
	{"import", "JTG1002", "import declarations", "Inline supported code or add import mapping before transpiling."},
	{"interface", "JTG1003", "interfaces", "Add an interface-to-Go-interface lowering strategy before using Java interfaces."},
	{"extends", "JTG1004", "inheritance", "Rewrite inheritance as composition or add an object model lowering pass."},
	{"implements", "JTG1005", "interface implementation", "Add Java interface mapping before transpiling implements clauses."},
	{"try", "JTG1007", "try/catch exceptions", "Design explicit Go error returns before transpiling exceptions."},
	{"catch", "JTG1008", "catch blocks", "Design explicit Go error returns before transpiling exceptions."},
	{"finally", "JTG1009", "finally blocks", "Map cleanup to defer before transpiling finally blocks."},
	{"throw", "JTG1010", "throw statements", "Map thrown exceptions to Go error values before transpiling throw."},
	{"throws", "JTG1011", "throws declarations", "Map checked exceptions to explicit Go error returns."},
}

var (
	classDeclPattern   = regexp.MustCompile(`\bclass\s+([A-Za-z_$][A-Za-z0-9_$]*)\b`)
	genericUsageRegexp = regexp.MustCompile(`([A-Za-z_][A-Za-z0-9_\.]*)\s*<\s*([A-Za-z_?][A-Za-z0-9_?\.]*)`)
)

func Detect(fileName string, source string) []semantic.Diagnostic {
	masked := maskCommentsAndLiterals(source)
	lines := splitLines(masked)
	diagnostics := []semantic.Diagnostic{}
	methodLines := map[string]int{}
	braceDepth := 0
	className := ""

	for _, current := range lines {
		trimmed := strings.TrimLeftFunc(current.text, unicode.IsSpace)
		if className == "" {
			if match := classDeclPattern.FindStringSubmatch(current.text); len(match) == 2 {
				className = match[1]
			}
		}

		for _, rule := range keywordRules {
			if position := findKeyword(current.text, rule.keyword); position >= 0 {
				diagnostics = append(diagnostics, diagnostic(fileName, current.number, position, rule.code, rule.feature, rule.recommendation))
			}
		}

		if position := strings.Index(current.text, "->"); position >= 0 {
			diagnostics = append(diagnostics, diagnostic(fileName, current.number, position, "JTG1012", "lambdas", "Rewrite the lambda as a named method or add functional-interface lowering."))
		}

		if trimmed != "" {
			if position := strings.Index(current.text, "@"); position >= 0 {
				diagnostics = append(diagnostics, diagnostic(fileName, current.number, position, "JTG1013", "annotations", "Remove annotations or add an annotation policy for migration metadata."))
			}
		}

		if genericUsageRegexp.MatchString(current.text) && strings.Contains(current.text, ">") {
			diagnostics = append(diagnostics, diagnostic(fileName, current.number, strings.Index(current.text, "<"), "JTG1014", "generics", "Specialize the generic type manually or add a Java-to-Go type-parameter mapping pass."))
		}

		if position := strings.Index(current.text, "["); position >= 0 && !isAllowedMainArrayParameter(current.text) {
			diagnostics = append(diagnostics, diagnostic(fileName, current.number, position, "JTG1015", "arrays and indexing", "Add array declarations, indexing, and length lowering before using arrays."))
		}

		if braceDepth == 1 && looksLikeMethodSignature(current.text) {
			methodName := methodNameFromSignature(current.text)
			if methodName != "" && methodName != className {
				if _, exists := methodLines[methodName]; exists {
					diagnostics = append(diagnostics, diagnostic(fileName, current.number, strings.Index(current.text, methodName), "JTG1018", "method overloading", "Rename overloaded methods or add overload name mangling before generating Go."))
				} else {
					methodLines[methodName] = current.number
				}
			}
		}

		braceDepth += braceDelta(current.text)
		if braceDepth < 0 {
			braceDepth = 0
		}
	}

	return diagnostics
}

func diagnostic(fileName string, lineNumber int, zeroBasedColumn int, code string, feature string, recommendation string) semantic.Diagnostic {
	return semantic.Diagnostic{
		File:           fileName,
		Line:           lineNumber,
		Column:         zeroBasedColumn + 1,
		Code:           code,
		Message:        "unsupported feature: " + feature,
		Recommendation: recommendation,
	}
}

func maskCommentsAndLiterals(source string) string {
	masked := []rune(source)
	runes := []rune(source)
	inLineComment := false
	inBlockComment := false
	inString := false
	inChar := false
	escaped := false

	for i := 0; i < len(runes); i++ {
		current := runes[i]
		next := rune(0)
		if i+1 < len(runes) {
			next = runes[i+1]
		}

		switch {
		case inLineComment:
			if current == '\n' {
				inLineComment = false
			} else {
				masked[i] = ' '
			}
		case inBlockComment:
			if current == '*' && next == '/' {
				masked[i] = ' '
				masked[i+1] = ' '
				i++
				inBlockComment = false
			} else if current != '\n' {
				masked[i] = ' '
			}
		case inString:
			if current != '\n' {
				masked[i] = ' '
			}
			if escaped {
				escaped = false
			} else if current == '\\' {
				escaped = true
			} else if current == '"' {
				inString = false
			}
		case inChar:
			if current != '\n' {
				masked[i] = ' '
			}
			if escaped {
				escaped = false
			} else if current == '\\' {
				escaped = true
			} else if current == '\'' {
				inChar = false
			}
		case current == '/' && next == '/':
			masked[i] = ' '
			masked[i+1] = ' '
			i++
			inLineComment = true
		case current == '/' && next == '*':
			masked[i] = ' '
			masked[i+1] = ' '
			i++
			inBlockComment = true
		case current == '"':
			masked[i] = ' '
			inString = true
		case current == '\'':
			masked[i] = ' '
			inChar = true
		}
	}

	return string(masked)
}

func splitLines(source string) []line {
	if source == "" {
		return []line{{number: 1}}
	}
	parts := strings.Split(source, "\n")
	out := make([]line, 0, len(parts))
	for i, part := range parts {
		if i == len(parts)-1 && part == "" {
			continue
		}
		out = append(out, line{number: i + 1, text: strings.TrimSuffix(part, "\r")})
	}
	return out
}

func findKeyword(text string, keyword string) int {
	position := strings.Index(text, keyword)
	for position >= 0 {
		leftOK := position == 0 || !isIdentifierChar(rune(text[position-1]))
		after := position + len(keyword)
		rightOK := after >= len(text) || !isIdentifierChar(rune(text[after]))
		if leftOK && rightOK {
			return position
		}
		next := strings.Index(text[position+1:], keyword)
		if next < 0 {
			return -1
		}
		position += next + 1
	}
	return -1
}

func containsKeyword(text string, keyword string) bool {
	return findKeyword(text, keyword) >= 0
}

func isIdentifierChar(value rune) bool {
	return unicode.IsLetter(value) || unicode.IsDigit(value) || value == '_' || value == '$'
}

func isAllowedMainArrayParameter(text string) bool {
	return strings.Contains(text, "main") && strings.Contains(text, "String[] args")
}

func looksLikeMethodSignature(text string) bool {
	return strings.Contains(text, "(") && strings.Contains(text, ")") && !containsKeyword(text, "class")
}

func methodNameFromSignature(text string) string {
	paren := strings.Index(text, "(")
	if paren < 0 {
		return ""
	}
	end := paren
	for end > 0 && unicode.IsSpace(rune(text[end-1])) {
		end--
	}
	start := end
	for start > 0 && isIdentifierChar(rune(text[start-1])) {
		start--
	}
	return text[start:end]
}

func braceDelta(text string) int {
	delta := 0
	for _, value := range text {
		switch value {
		case '{':
			delta++
		case '}':
			delta--
		}
	}
	return delta
}
