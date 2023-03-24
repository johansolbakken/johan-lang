package lexer

import "fmt"

type Lexer struct {
	Input    string
	Position int
	Tokens   []Token
}

func NewLexer(input string) *Lexer {
	return &Lexer{Input: input}
}

func (l *Lexer) Lex() []Token {
	for l.Position < len(l.Input) {
		l.lexNext()
	}

	return l.Tokens
}

func (l *Lexer) lexNext() {
	if l.isWhitespace() {
		l.lexWhitespace()
	} else if l.isLetter() {
		l.lexIdentifier()
	} else if l.isNumber() {
		l.lexNumber()
	} else if l.isString() {
		l.lexString()
	} else if l.isParen() {
		l.lexParen()
	} else {
		fmt.Println("Unexpected character: ", l.Input[l.Position])
	}
}

func (l *Lexer) lexWhitespace() {
	for l.isWhitespace() {
		l.Position++
	}
}

func (l *Lexer) lexIdentifier() {
	start := l.Position

	for l.isLetter() {
		l.Position++
	}

	l.Tokens = append(l.Tokens, Token{
		Type:   TokenType_Identifier,
		Lexeme: l.Input[start:l.Position],
	})
}

func (l *Lexer) lexNumber() {
	start := l.Position

	for l.isNumber() {
		l.Position++
	}

	l.Tokens = append(l.Tokens, Token{
		Type:     TokenType_Number,
		IntValue: l.parseInt(start, l.Position),
		Lexeme:   l.Input[start:l.Position],
	})
}

func (l *Lexer) lexString() {
	start := l.Position

	l.Position++

	for !l.isString() {
		l.Position++
	}

	l.Position++

	l.Tokens = append(l.Tokens, Token{
		Type:     TokenType_String,
		StrValue: l.Input[start+1 : l.Position-1],
		Lexeme:   l.Input[start:l.Position],
	})
}

func (l *Lexer) lexParen() {
	l.Tokens = append(l.Tokens, Token{
		Type:   l.parenType(),
		Lexeme: l.Input[l.Position : l.Position+1],
	})

	l.Position++
}

func (l *Lexer) isWhitespace() bool {
	return l.Input[l.Position] == ' ' || l.Input[l.Position] == '\t' || l.Input[l.Position] == '\n'
}

func (l *Lexer) isLetter() bool {
	return (l.Input[l.Position] >= 'a' && l.Input[l.Position] <= 'z') || (l.Input[l.Position] >= 'A' && l.Input[l.Position] <= 'Z')
}

func (l *Lexer) isNumber() bool {
	return l.Input[l.Position] >= '0' && l.Input[l.Position] <= '9'
}

func (l *Lexer) isString() bool {
	return l.Input[l.Position] == '"'
}

func (l *Lexer) isParen() bool {
	return l.Input[l.Position] == '(' || l.Input[l.Position] == ')'
}

func (l *Lexer) parenType() int {
	if l.Input[l.Position] == '(' {
		return TokenType_OpenParen
	}

	return TokenType_CloseParen
}

func (l *Lexer) parseInt(start, end int) int {
	var result int

	for i := start; i < end; i++ {
		result = result*10 + int(l.Input[i]-'0')
	}

	return result
}
