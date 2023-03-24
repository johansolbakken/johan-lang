package symboltable

import (
	"johan-lang/internal/lexer"
	"johan-lang/internal/parser"
)

type StringList struct {
	Items []string
}

func NewStringList() *StringList {
	return &StringList{}
}

func (s *StringList) Add(item string) int {
	s.Items = append(s.Items, item)
	return len(s.Items) - 1
}

func ExtractStringFromAst(stringList *StringList, node *parser.Node) {
	if node.Token.Type == lexer.TokenType_String {
		node.StringIndex = stringList.Add(node.Token.StrValue)
	}

	for _, child := range node.Children {
		ExtractStringFromAst(stringList, child)
	}
}
