package migrate

import (
	"sort"
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/javaproject"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/semantic"
)

type migrationContext struct {
	file            javaproject.File
	model           semantic.ProjectModel
	modulePath      string
	imports         map[string]ir.Import
	importedClasses map[string]*semantic.ClassSymbol
	currentClasses  map[string]*semantic.ClassSymbol
	localTypes      map[string]string
	localFuncs      map[string]string
}

func rewriteForMigration(file *ir.File, projectFile javaproject.File, model semantic.ProjectModel, modulePath string) {
	ctx := newMigrationContext(projectFile, model, modulePath)
	ctx.collectLocalNames(file)
	ctx.rewriteFile(file)
}

func newMigrationContext(file javaproject.File, model semantic.ProjectModel, modulePath string) *migrationContext {
	ctx := &migrationContext{
		file:            file,
		model:           model,
		modulePath:      modulePath,
		imports:         map[string]ir.Import{},
		importedClasses: map[string]*semantic.ClassSymbol{},
		currentClasses:  map[string]*semantic.ClassSymbol{},
		localTypes:      map[string]string{},
		localFuncs:      map[string]string{},
	}
	if pkg, ok := model.Packages[file.PackageName]; ok {
		for name, class := range pkg.Classes {
			ctx.currentClasses[name] = class
		}
	}
	for _, importDecl := range file.Imports {
		if importDecl.Wildcard {
			if pkg, ok := model.Packages[importDecl.Name]; ok {
				for name, class := range pkg.Classes {
					ctx.importedClasses[name] = class
				}
			}
			continue
		}
		if class, ok := model.Classes[importDecl.Name]; ok {
			ctx.importedClasses[class.Name] = class
		}
	}
	return ctx
}

func (ctx *migrationContext) collectLocalNames(file *ir.File) {
	for _, class := range file.Classes {
		ctx.localTypes[class.Name] = exportIdentifier(class.Name)
	}
	for _, iface := range file.Interfaces {
		ctx.localTypes[iface.Name] = exportIdentifier(iface.Name)
	}
	for _, fn := range file.Funcs {
		ctx.localFuncs[fn.Name] = exportFunctionName(fn.Name)
	}
	for _, class := range file.Classes {
		for _, fn := range class.Methods {
			ctx.localFuncs[fn.Name] = exportFunctionName(fn.Name)
		}
	}
}

func (ctx *migrationContext) rewriteFile(file *ir.File) {
	file.PackageName = goPackageName(ctx.file.PackageName)
	for i := range file.Classes {
		file.Classes[i].Name = ctx.rewriteLocalTypeName(file.Classes[i].Name)
		for fieldIndex := range file.Classes[i].Fields {
			file.Classes[i].Fields[fieldIndex].Type = ctx.rewriteType(file.Classes[i].Fields[fieldIndex].Type)
		}
		for methodIndex := range file.Classes[i].Methods {
			ctx.rewriteFunc(&file.Classes[i].Methods[methodIndex])
		}
	}
	for i := range file.Interfaces {
		file.Interfaces[i].Name = ctx.rewriteLocalTypeName(file.Interfaces[i].Name)
		for methodIndex := range file.Interfaces[i].Methods {
			ctx.rewriteFunc(&file.Interfaces[i].Methods[methodIndex])
		}
	}
	for i := range file.Funcs {
		ctx.rewriteFunc(&file.Funcs[i])
	}
	ctx.rewriteTestFile(file)
	file.Imports = ctx.sortedImports()
}

func (ctx *migrationContext) rewriteFunc(fn *ir.Func) {
	fn.Name = ctx.rewriteFunctionName(fn.Name)
	fn.ReturnType = ctx.rewriteType(fn.ReturnType)
	if fn.Receiver != nil {
		fn.Receiver.Type = ctx.rewriteType(fn.Receiver.Type)
	}
	for i := range fn.Params {
		fn.Params[i].Type = ctx.rewriteType(fn.Params[i].Type)
	}
	for i := range fn.Body {
		fn.Body[i] = ctx.rewriteStmt(fn.Body[i])
	}
}

func (ctx *migrationContext) rewriteTestFile(file *ir.File) {
	if ctx.file.SourceRoot != "src/test/java" {
		return
	}
	needsTesting := false
	for i := range file.Funcs {
		if file.Funcs[i].ReturnType.Kind != ir.KindVoid {
			continue
		}
		file.Funcs[i].Name = "Test" + exportIdentifier(strings.TrimPrefix(file.Funcs[i].Name, "Test"))
		file.Funcs[i].Params = []ir.Param{{
			Name: "t",
			Type: ir.Type{Kind: ir.KindPointer, Elem: &ir.Type{Kind: ir.KindObject, Name: "testing.T"}},
			Span: file.Funcs[i].Span,
		}}
		needsTesting = true
	}
	if needsTesting {
		ctx.imports["testing"] = ir.Import{Path: "testing"}
	}
}

func (ctx *migrationContext) rewriteStmt(stmt ir.Stmt) ir.Stmt {
	switch value := stmt.(type) {
	case ir.VarDeclStmt:
		value.Type = ctx.rewriteType(value.Type)
		value.Value = ctx.rewriteExpr(value.Value)
		return value
	case ir.AssignStmt:
		value.Value = ctx.rewriteExpr(value.Value)
		return value
	case ir.ExprStmt:
		value.Expr = ctx.rewriteExpr(value.Expr)
		return value
	case ir.ReturnStmt:
		value.Value = ctx.rewriteExpr(value.Value)
		return value
	case ir.IfStmt:
		value.Cond = ctx.rewriteExpr(value.Cond)
		for i := range value.Then {
			value.Then[i] = ctx.rewriteStmt(value.Then[i])
		}
		for i := range value.Else {
			value.Else[i] = ctx.rewriteStmt(value.Else[i])
		}
		return value
	case ir.WhileStmt:
		value.Cond = ctx.rewriteExpr(value.Cond)
		for i := range value.Body {
			value.Body[i] = ctx.rewriteStmt(value.Body[i])
		}
		return value
	case ir.ForStmt:
		value.Init = ctx.rewriteStmt(value.Init)
		value.Cond = ctx.rewriteExpr(value.Cond)
		value.Post = ctx.rewriteStmt(value.Post)
		for i := range value.Body {
			value.Body[i] = ctx.rewriteStmt(value.Body[i])
		}
		return value
	default:
		return stmt
	}
}

func (ctx *migrationContext) rewriteExpr(expr ir.Expr) ir.Expr {
	switch value := expr.(type) {
	case nil:
		return nil
	case ir.NameExpr:
		value.Type = ctx.rewriteType(value.Type)
		return value
	case ir.BinaryExpr:
		value.Left = ctx.rewriteExpr(value.Left)
		value.Right = ctx.rewriteExpr(value.Right)
		value.Type = ctx.rewriteType(value.Type)
		return value
	case ir.UnaryExpr:
		value.Expr = ctx.rewriteExpr(value.Expr)
		value.Type = ctx.rewriteType(value.Type)
		return value
	case ir.CallExpr:
		for i := range value.Args {
			value.Args[i] = ctx.rewriteExpr(value.Args[i])
		}
		value = ctx.rewriteCall(value)
		value.Type = ctx.rewriteType(value.Type)
		return value
	case ir.FieldExpr:
		value.Type = ctx.rewriteType(value.Type)
		return value
	case ir.AddressExpr:
		value.Expr = ctx.rewriteExpr(value.Expr)
		value.Type = ctx.rewriteType(value.Type)
		return value
	case ir.CompositeLitExpr:
		value.TypeName = ctx.rewriteLocalTypeName(value.TypeName)
		for i := range value.Fields {
			value.Fields[i].Value = ctx.rewriteExpr(value.Fields[i].Value)
		}
		value.Type = ctx.rewriteType(value.Type)
		return value
	default:
		return expr
	}
}

func (ctx *migrationContext) rewriteCall(call ir.CallExpr) ir.CallExpr {
	if call.Target == "System.out" && call.Name == "println" {
		return call
	}
	if class, ok := ctx.resolveClass(call.Target); ok {
		call.Name = exportFunctionName(call.Name)
		if class.PackageName == ctx.file.PackageName {
			call.Target = ""
			return call
		}
		call.Target = ctx.importClassPackage(class)
		return call
	}
	if call.Target == "" {
		if className, ok := strings.CutPrefix(call.Name, "New"); ok {
			if class, ok := ctx.resolveClass(className); ok {
				call.Name = "New" + exportIdentifier(class.Name)
				if class.PackageName != ctx.file.PackageName {
					call.Target = ctx.importClassPackage(class)
				}
				return call
			}
		}
		call.Name = ctx.rewriteFunctionName(call.Name)
		return call
	}
	call.Name = exportFunctionName(call.Name)
	return call
}

func (ctx *migrationContext) rewriteType(t ir.Type) ir.Type {
	switch t.Kind {
	case ir.KindArray:
		if t.Elem != nil {
			elem := ctx.rewriteType(*t.Elem)
			t.Elem = &elem
		}
	case ir.KindPointer:
		if t.Elem != nil {
			elem := ctx.rewriteType(*t.Elem)
			t.Elem = &elem
		}
	case ir.KindObject:
		t.Name = ctx.rewriteObjectName(t.Name)
	}
	return t
}

func (ctx *migrationContext) rewriteObjectName(name string) string {
	if name == "" {
		return name
	}
	if strings.Contains(name, ".") {
		if class, ok := ctx.model.Classes[name]; ok {
			exported := exportIdentifier(class.Name)
			if class.PackageName == ctx.file.PackageName {
				return exported
			}
			return ctx.importClassPackage(class) + "." + exported
		}
		return name
	}
	if class, ok := ctx.resolveClass(name); ok {
		exported := exportIdentifier(class.Name)
		if class.PackageName == ctx.file.PackageName {
			return exported
		}
		return ctx.importClassPackage(class) + "." + exported
	}
	return ctx.rewriteLocalTypeName(name)
}

func (ctx *migrationContext) rewriteLocalTypeName(name string) string {
	if rewritten, ok := ctx.localTypes[name]; ok {
		return rewritten
	}
	return exportIdentifier(name)
}

func (ctx *migrationContext) rewriteFunctionName(name string) string {
	if rewritten, ok := ctx.localFuncs[name]; ok {
		return rewritten
	}
	return exportFunctionName(name)
}

func (ctx *migrationContext) resolveClass(simpleName string) (*semantic.ClassSymbol, bool) {
	if simpleName == "" {
		return nil, false
	}
	if strings.Contains(simpleName, ".") {
		if class, ok := ctx.model.Classes[simpleName]; ok {
			return class, true
		}
	}
	if class, ok := ctx.importedClasses[simpleName]; ok {
		return class, true
	}
	if class, ok := ctx.currentClasses[simpleName]; ok {
		return class, true
	}
	if class, ok := ctx.model.Classes[simpleName]; ok {
		return class, true
	}
	return nil, false
}

func (ctx *migrationContext) importClassPackage(class *semantic.ClassSymbol) string {
	alias := goPackageName(class.PackageName)
	importPath := strings.TrimRight(ctx.modulePath, "/") + "/" + strings.ReplaceAll(class.PackageName, ".", "/")
	ctx.imports[importPath] = ir.Import{Name: alias, Path: importPath}
	return alias
}

func (ctx *migrationContext) sortedImports() []ir.Import {
	paths := make([]string, 0, len(ctx.imports))
	for importPath := range ctx.imports {
		paths = append(paths, importPath)
	}
	sort.Strings(paths)
	out := make([]ir.Import, 0, len(paths))
	for _, importPath := range paths {
		out = append(out, ctx.imports[importPath])
	}
	return out
}

func exportFunctionName(name string) string {
	if name == "main" {
		return name
	}
	return exportIdentifier(name)
}

func exportIdentifier(name string) string {
	name = sanitizeIdentifier(name)
	if name == "" {
		return name
	}
	first := name[0]
	if first >= 'a' && first <= 'z' {
		return strings.ToUpper(name[:1]) + name[1:]
	}
	if first == '_' {
		return "Exported" + strings.TrimLeft(name, "_")
	}
	return name
}
