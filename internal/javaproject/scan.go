package javaproject

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/parser/gen"
)

func Scan(path string) (Project, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return Project{}, err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return Project{}, err
	}

	root := abs
	if !info.IsDir() {
		root = filepath.Dir(abs)
	}
	project := Project{
		Root:              slash(root),
		BuildFiles:        discoverBuildFiles(root),
		SourceRoots:       discoverSourceRoots(root),
		FeatureCounts:     map[string]int{},
		UnsupportedCounts: map[string]int{},
	}

	paths, err := discoverJavaFiles(abs, info)
	if err != nil {
		return Project{}, err
	}
	for _, filePath := range paths {
		file, err := scanFile(root, filePath)
		if err != nil {
			return Project{}, err
		}
		project.Files = append(project.Files, file)
	}

	finalize(&project)
	return project, nil
}

func discoverBuildFiles(root string) []BuildFile {
	candidates := []BuildFile{
		{Path: "pom.xml", Kind: "maven"},
		{Path: "build.gradle", Kind: "gradle"},
		{Path: "build.gradle.kts", Kind: "gradle-kotlin"},
	}
	out := []BuildFile{}
	for _, candidate := range candidates {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(candidate.Path))); err == nil {
			out = append(out, candidate)
		}
	}
	return out
}

func discoverSourceRoots(root string) []SourceRoot {
	candidates := []SourceRoot{
		{Path: "src/main/java", Kind: "main"},
		{Path: "src/test/java", Kind: "test"},
	}
	out := make([]SourceRoot, 0, len(candidates))
	for _, candidate := range candidates {
		if info, err := os.Stat(filepath.Join(root, filepath.FromSlash(candidate.Path))); err == nil && info.IsDir() {
			candidate.Exists = true
		}
		out = append(out, candidate)
	}
	return out
}

func discoverJavaFiles(path string, info os.FileInfo) ([]string, error) {
	if !info.IsDir() {
		if strings.EqualFold(filepath.Ext(path), ".java") {
			return []string{path}, nil
		}
		return nil, nil
	}

	files := []string{}
	err := filepath.WalkDir(path, func(current string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".git" || entry.Name() == "build" || entry.Name() == "target" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.EqualFold(filepath.Ext(current), ".java") {
			files = append(files, current)
		}
		return nil
	})
	sort.Strings(files)
	return files, err
}

func scanFile(root string, path string) (File, error) {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return File{}, err
	}
	source, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}

	tree, syntax := parser.ParseSource(path, string(source))
	file := File{
		Path:         slash(path),
		RelativePath: slash(rel),
		SourceRoot:   sourceRootFor(rel),
		Risk:         "low",
	}
	for _, diagnostic := range syntax {
		file.Diagnostics = append(file.Diagnostics, Diagnostic{
			File:    slash(diagnostic.File),
			Line:    diagnostic.Line,
			Column:  diagnostic.Column,
			Code:    "JTG0001",
			Message: diagnostic.Message,
		})
	}
	if tree != nil {
		extractCompilationUnit(&file, tree)
	}
	file.Risk = riskFor(file)
	return file, nil
}

func extractCompilationUnit(file *File, tree gen.ICompilationUnitContext) {
	if tree.PackageDecl() != nil {
		file.PackageName = tree.PackageDecl().QualifiedName().GetText()
		file.Features = append(file.Features, feature("package declarations", true, tree.PackageDecl().GetStart()))
	}
	for _, importDecl := range tree.AllImportDecl() {
		name := importDecl.QualifiedName().GetText()
		wildcard := strings.HasSuffix(importDecl.GetText(), ".*;")
		file.Imports = append(file.Imports, Import{Name: name, Wildcard: wildcard})
		file.Features = append(file.Features, feature("import declarations", true, importDecl.GetStart()))
	}
	for _, typeDecl := range tree.AllTypeDecl() {
		classDecl := typeDecl.ClassDecl()
		if classDecl == nil {
			continue
		}
		class := extractClass(file.PackageName, classDecl)
		file.Classes = append(file.Classes, class)
		file.Features = append(file.Features, feature("class declarations", true, classDecl.GetStart()))
		for _, field := range class.Fields {
			file.Features = append(file.Features, Feature{Name: "class fields", Supported: true, Span: field.Span})
		}
		for _, constructor := range class.Constructors {
			file.Features = append(file.Features, Feature{Name: "constructors", Supported: true, Span: constructor.Span})
		}
		for _, method := range class.Methods {
			if method.Static {
				file.Features = append(file.Features, Feature{Name: "static methods", Supported: true, Span: method.Span})
			} else {
				file.Features = append(file.Features, Feature{Name: "instance methods", Supported: true, Span: method.Span})
			}
		}
	}
}

func extractClass(packageName string, ctx gen.IClassDeclContext) Class {
	name := ctx.Identifier().GetText()
	class := Class{
		Name:      name,
		Qualified: qualified(packageName, name),
		Span:      span(ctx.GetStart()),
	}
	for _, member := range ctx.ClassBody().AllClassMember() {
		if field := member.FieldDecl(); field != nil {
			class.Fields = append(class.Fields, extractFields(field)...)
			continue
		}
		if method := member.MethodDecl(); method != nil {
			class.Methods = append(class.Methods, extractMethod(method))
			continue
		}
		if constructor := member.ConstructorDecl(); constructor != nil {
			class.Constructors = append(class.Constructors, extractConstructor(constructor))
		}
	}
	return class
}

func extractFields(ctx gen.IFieldDeclContext) []Field {
	fields := []Field{}
	typ := ctx.TypeType().GetText()
	for _, variable := range ctx.VariableDeclarators().AllVariableDeclarator() {
		fields = append(fields, Field{Name: variable.Identifier().GetText(), Type: typ, Span: span(variable.GetStart())})
	}
	return fields
}

func extractMethod(ctx gen.IMethodDeclContext) Method {
	return Method{
		Name:       ctx.Identifier().GetText(),
		ReturnType: ctx.TypeTypeOrVoid().GetText(),
		Static:     hasModifier(ctx.AllModifier(), "static"),
		Params:     extractParams(ctx.FormalParameters()),
		Span:       span(ctx.GetStart()),
	}
}

func extractConstructor(ctx gen.IConstructorDeclContext) Constructor {
	return Constructor{
		Name:   ctx.Identifier().GetText(),
		Params: extractParams(ctx.FormalParameters()),
		Span:   span(ctx.GetStart()),
	}
}

func extractParams(ctx gen.IFormalParametersContext) []Param {
	if ctx == nil || ctx.FormalParameterList() == nil {
		return nil
	}
	params := []Param{}
	for _, param := range ctx.FormalParameterList().AllFormalParameter() {
		params = append(params, Param{Name: param.Identifier().GetText(), Type: param.TypeType().GetText()})
	}
	return params
}

func finalize(project *Project) {
	packages := map[string]bool{}
	for i := range project.Files {
		file := &project.Files[i]
		if file.PackageName != "" {
			packages[file.PackageName] = true
		}
		for _, feature := range file.Features {
			project.FeatureCounts[feature.Name]++
		}
		for _, unsupported := range file.Unsupported {
			project.UnsupportedCounts[unsupported.Name]++
		}
		project.Summary.JavaFiles++
		project.Summary.Classes += len(file.Classes)
		project.Summary.Diagnostics += len(file.Diagnostics)
		project.Summary.Unsupported += len(file.Unsupported)
		switch file.Risk {
		case "high":
			project.Summary.HighRiskFiles++
		case "medium":
			project.Summary.MediumRiskFiles++
		default:
			project.Summary.LowRiskFiles++
		}
		for _, class := range file.Classes {
			project.Summary.Fields += len(class.Fields)
			project.Summary.Constructors += len(class.Constructors)
			project.Summary.Methods += len(class.Methods)
		}
	}
	project.Packages = sortedKeys(packages)
	project.Summary.Packages = len(project.Packages)
	project.InternalDependencies = internalDependencies(project.Files, packages)
	project.Summary.InternalImports = len(project.InternalDependencies)
}

func internalDependencies(files []File, packages map[string]bool) []Dependency {
	out := []Dependency{}
	for _, file := range files {
		for _, importDecl := range file.Imports {
			for packageName := range packages {
				if importDecl.Name == packageName || strings.HasPrefix(importDecl.Name, packageName+".") {
					out = append(out, Dependency{FromFile: file.RelativePath, Import: importDecl.Name, Package: packageName})
				}
			}
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].FromFile != out[j].FromFile {
			return out[i].FromFile < out[j].FromFile
		}
		return out[i].Import < out[j].Import
	})
	return out
}

func riskFor(file File) string {
	if len(file.Diagnostics) > 0 || len(file.Unsupported) > 0 {
		return "high"
	}
	if len(file.Imports) > 0 || file.PackageName != "" {
		return "medium"
	}
	return "low"
}

func feature(name string, supported bool, token antlr.Token) Feature {
	return Feature{Name: name, Supported: supported, Span: span(token)}
}

func unsupported(code string, name string, recommendation string, source Span) Feature {
	return Feature{Code: code, Name: name, Supported: false, Recommendation: recommendation, Span: source}
}

func span(token antlr.Token) Span {
	if token == nil {
		return Span{}
	}
	return Span{Line: token.GetLine(), Column: token.GetColumn()}
}

func hasModifier(modifiers []gen.IModifierContext, name string) bool {
	for _, modifier := range modifiers {
		if modifier.GetText() == name {
			return true
		}
	}
	return false
}

func sourceRootFor(rel string) string {
	rel = slash(rel)
	switch {
	case strings.HasPrefix(rel, "src/main/java/"):
		return "src/main/java"
	case strings.HasPrefix(rel, "src/test/java/"):
		return "src/test/java"
	default:
		return "."
	}
}

func qualified(packageName string, name string) string {
	if packageName == "" {
		return name
	}
	return packageName + "." + name
}

func sortedKeys(in map[string]bool) []string {
	out := make([]string, 0, len(in))
	for key := range in {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

func slash(path string) string {
	return filepath.ToSlash(path)
}
