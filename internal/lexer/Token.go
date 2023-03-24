package lexer

const (
	TokenType_Identifier = iota
	TokenType_String     = iota
	TokenType_OpenParen  = iota
	TokenType_CloseParen = iota
	TokenType_Number
)

type Token struct {
	Type     int
	StrValue string
	IntValue int
	Lexeme   string
}

func TokenTypeToString(tokenType int) string {
	switch tokenType {
	case TokenType_Identifier:
		return "Identifier"
	case TokenType_String:
		return "String"
	case TokenType_OpenParen:
		return "OpenParen"
	case TokenType_CloseParen:
		return "CloseParen"
	case TokenType_Number:
		return "Number"
	}

	return "Unknown"
}

func (t Token) String() string {
	return TokenTypeToString(t.Type) + " " + t.Lexeme
}
