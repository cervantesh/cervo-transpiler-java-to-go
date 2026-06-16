package lower

import (
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/semantic"
)

func MapJavaType(name string) ir.Type {
	return semantic.JavaType(name)
}
