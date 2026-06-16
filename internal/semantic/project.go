package semantic

import (
	"sort"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
)

const (
	CodeDuplicateSymbol     = "JTG2001"
	CodeOverloadUnsupported = "JTG2004"
	CodeImportUnresolved    = "JTG2005"
)

type ProjectModel struct {
	Packages    map[string]*PackageSymbol
	Classes     map[string]*ClassSymbol
	Diagnostics []Diagnostic
}

type PackageSymbol struct {
	Name    string
	Classes map[string]*ClassSymbol
}

type ClassSymbol struct {
	Name         string
	Qualified    string
	PackageName  string
	File         string
	Span         SourceSpan
	Fields       map[string]FieldSymbol
	Methods      map[string][]MethodSymbol
	Constructors []MethodSymbol
}

type FieldSymbol struct {
	Name string
	Type ir.Type
	File string
	Span SourceSpan
}

type MethodSymbol struct {
	Name       string
	ReturnType ir.Type
	Static     bool
	Params     []Symbol
	File       string
	Span       SourceSpan
}

type SourceSpan struct {
	File   string
	Line   int
	Column int
}

func BuildProjectModel(project javaproject.Project) ProjectModel {
	model := ProjectModel{
		Packages: map[string]*PackageSymbol{},
		Classes:  map[string]*ClassSymbol{},
	}

	files := append([]javaproject.File(nil), project.Files...)
	sort.Slice(files, func(i, j int) bool {
		return files[i].RelativePath < files[j].RelativePath
	})

	for _, file := range files {
		if file.PackageName != "" {
			model.ensurePackage(file.PackageName)
		}
		for _, class := range file.Classes {
			model.addClass(file, class)
		}
	}

	for _, file := range files {
		model.resolveImports(file)
	}

	return model
}

func (m *ProjectModel) ensurePackage(name string) *PackageSymbol {
	if existing, ok := m.Packages[name]; ok {
		return existing
	}
	symbol := &PackageSymbol{Name: name, Classes: map[string]*ClassSymbol{}}
	m.Packages[name] = symbol
	return symbol
}

func (m *ProjectModel) addClass(file javaproject.File, class javaproject.Class) {
	qualified := class.Qualified
	if qualified == "" {
		qualified = class.Name
	}
	if _, exists := m.Classes[qualified]; exists {
		m.Diagnostics = append(m.Diagnostics, semanticDiagnostic(file.Path, class.Span, CodeDuplicateSymbol, "duplicate symbol: class "+qualified))
		return
	}

	symbol := &ClassSymbol{
		Name:         class.Name,
		Qualified:    qualified,
		PackageName:  file.PackageName,
		File:         file.Path,
		Span:         sourceSpan(file.Path, class.Span),
		Fields:       map[string]FieldSymbol{},
		Methods:      map[string][]MethodSymbol{},
		Constructors: make([]MethodSymbol, 0, len(class.Constructors)),
	}
	m.Classes[qualified] = symbol
	if file.PackageName != "" {
		m.ensurePackage(file.PackageName).Classes[class.Name] = symbol
	}

	for _, field := range class.Fields {
		if _, exists := symbol.Fields[field.Name]; exists {
			m.Diagnostics = append(m.Diagnostics, semanticDiagnostic(file.Path, field.Span, CodeDuplicateSymbol, "duplicate symbol: field "+qualified+"."+field.Name))
			continue
		}
		symbol.Fields[field.Name] = FieldSymbol{Name: field.Name, Type: JavaType(field.Type), File: file.Path, Span: sourceSpan(file.Path, field.Span)}
	}

	for _, constructor := range class.Constructors {
		symbol.Constructors = append(symbol.Constructors, MethodSymbol{
			Name:       constructor.Name,
			ReturnType: ir.Type{Kind: ir.KindObject, Name: qualified},
			Params:     params(constructor.Params),
			File:       file.Path,
			Span:       sourceSpan(file.Path, constructor.Span),
		})
	}

	for _, method := range class.Methods {
		methodSymbol := MethodSymbol{
			Name:       method.Name,
			ReturnType: JavaType(method.ReturnType),
			Static:     method.Static,
			Params:     params(method.Params),
			File:       file.Path,
			Span:       sourceSpan(file.Path, method.Span),
		}
		for _, existing := range symbol.Methods[method.Name] {
			if sameSignature(existing.Params, methodSymbol.Params) {
				m.Diagnostics = append(m.Diagnostics, semanticDiagnostic(file.Path, method.Span, CodeDuplicateSymbol, "duplicate symbol: method "+qualified+"."+method.Name))
			} else {
				m.Diagnostics = append(m.Diagnostics, semanticDiagnostic(file.Path, method.Span, CodeOverloadUnsupported, "overload not supported: method "+qualified+"."+method.Name))
			}
		}
		symbol.Methods[method.Name] = append(symbol.Methods[method.Name], methodSymbol)
	}
}

func (m *ProjectModel) resolveImports(file javaproject.File) {
	for _, importDecl := range file.Imports {
		if importDecl.Wildcard {
			if _, ok := m.Packages[importDecl.Name]; !ok && m.looksInternal(importDecl.Name) {
				m.Diagnostics = append(m.Diagnostics, Diagnostic{File: file.Path, Code: CodeImportUnresolved, Message: "import unresolved: " + importDecl.Name})
			}
			continue
		}
		if _, ok := m.Classes[importDecl.Name]; ok {
			continue
		}
		if _, ok := m.Packages[importDecl.Name]; ok {
			continue
		}
		if m.looksInternal(importDecl.Name) {
			m.Diagnostics = append(m.Diagnostics, Diagnostic{File: file.Path, Code: CodeImportUnresolved, Message: "import unresolved: " + importDecl.Name})
		}
	}
}

func (m *ProjectModel) looksInternal(name string) bool {
	for packageName := range m.Packages {
		if name == packageName || strings.HasPrefix(name, packageName+".") {
			return true
		}
	}
	return false
}

func params(in []javaproject.Param) []Symbol {
	out := make([]Symbol, 0, len(in))
	seen := map[string]bool{}
	for _, param := range in {
		if seen[param.Name] {
			continue
		}
		seen[param.Name] = true
		out = append(out, Symbol{Name: param.Name, Type: JavaType(param.Type)})
	}
	return out
}

func sameSignature(left []Symbol, right []Symbol) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if !typeEqual(left[i].Type, right[i].Type) {
			return false
		}
	}
	return true
}

func typeEqual(left ir.Type, right ir.Type) bool {
	if left.Kind != right.Kind || left.Name != right.Name {
		return false
	}
	if left.Elem == nil || right.Elem == nil {
		return left.Elem == nil && right.Elem == nil
	}
	return typeEqual(*left.Elem, *right.Elem)
}

func sourceSpan(file string, span javaproject.Span) SourceSpan {
	return SourceSpan{File: file, Line: span.Line, Column: span.Column}
}

func semanticDiagnostic(file string, span javaproject.Span, code string, message string) Diagnostic {
	return Diagnostic{File: file, Line: span.Line, Column: span.Column, Code: code, Message: message}
}
