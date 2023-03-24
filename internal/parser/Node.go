package parser

import (
	"johan-lang/internal/lexer"
)

const (
	NodeType_Unkown         = iota
	NodeType_Root           = iota
	NodeType_Identifier     = iota
	NodeType_String         = iota
	NodeType_Number         = iota
	NodeType_Literal        = iota
	NodeType_PrintStatement = iota
)

type Node struct {
	Children    []*Node
	Token       lexer.Token
	Type        int
	StringIndex int
}

func NewNode(token lexer.Token) *Node {
	node := Node{}
	node.Token = token

	switch token.Type {
	case lexer.TokenType_Identifier:
		node.Type = NodeType_Identifier
	case lexer.TokenType_String:
		node.Type = NodeType_String
	case lexer.TokenType_Number:
		node.Type = NodeType_Number
	case lexer.TokenType_OpenParen:
		node.Type = NodeType_Literal
	case lexer.TokenType_CloseParen:
		node.Type = NodeType_Literal
	default:
		node.Type = NodeType_Unkown
	}

	return &node
}

func (n Node) String() string {
	if n.Token.Type != 0 {
		return n.Token.String()
	}

	switch n.Type {
	case NodeType_Root:
		return "Root"
	case NodeType_Identifier:
		return "Identifier"
	case NodeType_String:
		return "String"
	case NodeType_PrintStatement:
		return "PrintStatement"
	case NodeType_Number:
		return "Number"
	case NodeType_Literal:
		return "Literal"
	}

	return "Unknown"
}

func (node *Node) WriteAst(indent string) {
	println(indent + node.String())
	for _, child := range node.Children {
		child.WriteAst(indent + "  ")
	}
}
