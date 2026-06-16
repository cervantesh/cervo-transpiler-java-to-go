package semantic

import (
	"fmt"

	"github.com/cervantesh/cervo-transpiler-java-to-go/internal/ir"
)

type Symbol struct {
	Name string
	Type ir.Type
}

type Scope struct {
	parent  *Scope
	symbols map[string]Symbol
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent:  parent,
		symbols: map[string]Symbol{},
	}
}

func (s *Scope) Define(symbol Symbol) error {
	if _, exists := s.symbols[symbol.Name]; exists {
		return fmt.Errorf("duplicate symbol %q", symbol.Name)
	}
	s.symbols[symbol.Name] = symbol
	return nil
}

func (s *Scope) Lookup(name string) (Symbol, bool) {
	if symbol, exists := s.symbols[name]; exists {
		return symbol, true
	}
	if s.parent != nil {
		return s.parent.Lookup(name)
	}
	return Symbol{}, false
}
