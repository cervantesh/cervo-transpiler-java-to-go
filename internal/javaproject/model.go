package javaproject

type Project struct {
	Root                 string         `json:"root"`
	BuildFiles           []BuildFile    `json:"buildFiles"`
	SourceRoots          []SourceRoot   `json:"sourceRoots"`
	Files                []File         `json:"files"`
	Packages             []string       `json:"packages"`
	InternalDependencies []Dependency   `json:"internalDependencies"`
	FeatureCounts        map[string]int `json:"featureCounts"`
	UnsupportedCounts    map[string]int `json:"unsupportedCounts"`
	Summary              Summary        `json:"summary"`
}

type BuildFile struct {
	Path string `json:"path"`
	Kind string `json:"kind"`
}

type SourceRoot struct {
	Path   string `json:"path"`
	Kind   string `json:"kind"`
	Exists bool   `json:"exists"`
}

type File struct {
	Path         string       `json:"path"`
	RelativePath string       `json:"relativePath"`
	SourceRoot   string       `json:"sourceRoot"`
	PackageName  string       `json:"packageName,omitempty"`
	Imports      []Import     `json:"imports,omitempty"`
	Classes      []Class      `json:"classes,omitempty"`
	Features     []Feature    `json:"features,omitempty"`
	Unsupported  []Feature    `json:"unsupported,omitempty"`
	Diagnostics  []Diagnostic `json:"diagnostics,omitempty"`
	Risk         string       `json:"risk"`
}

type Import struct {
	Name     string `json:"name"`
	Wildcard bool   `json:"wildcard,omitempty"`
}

type Class struct {
	Name         string        `json:"name"`
	Qualified    string        `json:"qualified"`
	Fields       []Field       `json:"fields,omitempty"`
	Methods      []Method      `json:"methods,omitempty"`
	Constructors []Constructor `json:"constructors,omitempty"`
	Span         Span          `json:"span"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Span Span   `json:"span"`
}

type Method struct {
	Name       string  `json:"name"`
	ReturnType string  `json:"returnType"`
	Static     bool    `json:"static"`
	Params     []Param `json:"params,omitempty"`
	Span       Span    `json:"span"`
}

type Constructor struct {
	Name   string  `json:"name"`
	Params []Param `json:"params,omitempty"`
	Span   Span    `json:"span"`
}

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Feature struct {
	Code           string `json:"code,omitempty"`
	Name           string `json:"name"`
	Supported      bool   `json:"supported"`
	Recommendation string `json:"recommendation,omitempty"`
	Span           Span   `json:"span,omitempty"`
}

type Diagnostic struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Dependency struct {
	FromFile string `json:"fromFile"`
	Import   string `json:"import"`
	Package  string `json:"package"`
}

type Span struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Summary struct {
	JavaFiles       int `json:"javaFiles"`
	Packages        int `json:"packages"`
	Classes         int `json:"classes"`
	Methods         int `json:"methods"`
	Constructors    int `json:"constructors"`
	Fields          int `json:"fields"`
	Diagnostics     int `json:"diagnostics"`
	Unsupported     int `json:"unsupported"`
	InternalImports int `json:"internalImports"`
	LowRiskFiles    int `json:"lowRiskFiles"`
	MediumRiskFiles int `json:"mediumRiskFiles"`
	HighRiskFiles   int `json:"highRiskFiles"`
}
