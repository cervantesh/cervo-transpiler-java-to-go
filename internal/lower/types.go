package lower

import (
	"strings"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

func MapJavaType(name string) ir.Type {
	if strings.HasSuffix(name, "[]") {
		elem := MapJavaType(strings.TrimSuffix(name, "[]"))
		return ir.Type{Kind: ir.KindArray, Elem: &elem}
	}
	switch name {
	case "void":
		return ir.Type{Kind: ir.KindVoid}
	case "boolean":
		return ir.Type{Kind: ir.KindBoolean}
	case "int":
		return ir.Type{Kind: ir.KindInt}
	case "double":
		return ir.Type{Kind: ir.KindDouble}
	case "String", "java.lang.String":
		return ir.Type{Kind: ir.KindString}
	default:
		return ir.Type{Kind: ir.KindObject, Name: name}
	}
}
