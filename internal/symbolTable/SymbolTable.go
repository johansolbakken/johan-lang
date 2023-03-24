package symboltable

import "johan-lang/internal/parser"

type Symbol struct {
	Name     string
	StrValue string
	IntValue int
	Node     *parser.Node
}

type SymbolTable struct {
	Parent   *SymbolTable
	Children []*SymbolTable
	Symbols  []*Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{}
}

func (s *SymbolTable) AddSymbol(symbol *Symbol) {
	s.Symbols = append(s.Symbols, symbol)
}

func (s *SymbolTable) AddChild(child *SymbolTable) {
	s.Children = append(s.Children, child)
	child.Parent = s
}

func (s *SymbolTable) FindSymbol(name string) *Symbol {
	for _, symbol := range s.Symbols {
		if symbol.Name == name {
			return symbol
		}
	}

	if s.Parent != nil {
		return s.Parent.FindSymbol(name)
	}

	return nil
}
