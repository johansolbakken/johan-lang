package parser

import "johan-lang/internal/lexer"

var stringPatteern = []int{
	lexer.TokenType_Identifier,
	lexer.TokenType_OpenParen,
	lexer.TokenType_String,
	lexer.TokenType_CloseParen,
}
