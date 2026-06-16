package ir

type Kind int

const (
	KindInvalid Kind = iota
	KindVoid
	KindBoolean
	KindInt
	KindDouble
	KindString
	KindArray
	KindObject
)

type Type struct {
	Kind Kind
	Name string
	Elem *Type
}

func (t Type) GoName() string {
	switch t.Kind {
	case KindVoid:
		return ""
	case KindBoolean:
		return "bool"
	case KindInt:
		return "int"
	case KindDouble:
		return "float64"
	case KindString:
		return "string"
	case KindArray:
		if t.Elem == nil {
			return "[]any"
		}
		return "[]" + t.Elem.GoName()
	case KindObject:
		if t.Name != "" {
			return t.Name
		}
		return "any"
	default:
		return "any"
	}
}
