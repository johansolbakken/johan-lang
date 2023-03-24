package parser

import "johan-lang/internal/lexer"

type Parser struct {
	Tokens []lexer.Token
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{Tokens: tokens}
}

func (p *Parser) Parse() Node {
	stack := NewNodeStack()
	position := 0

	for {
		if position < len(p.Tokens) {
			stack.Push(NewNode(p.Tokens[position]))
			position++
		}

		if stack.HasSequence(stringPatteern) {
			stack.Pop()               // CloseParen
			stringData := stack.Pop() // String
			stack.Pop()               // OpenParen
			identifier := stack.Pop()
			node := &Node{
				Type: NodeType_PrintStatement,
				Children: []*Node{
					identifier,
					stringData,
				},
			}
			stack.Push(node)
			continue
		}

		if position >= len(p.Tokens) {
			if stack.Size() == 1 {
				break
			}

			panic("Syntax error")
		}
	}

	root := Node{
		Type: NodeType_Root,
		Children: []*Node{
			stack.Pop(),
		},
	}

	return root
}
