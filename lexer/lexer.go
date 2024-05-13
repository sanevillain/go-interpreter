package lexer

import "sanevillain/go-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	return nil
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	return tok
}

func (l *Lexer) skipWhitespace() {
}

func (l *Lexer) readIdentifier() string {
	return ""
}

func (l *Lexer) readNumber() string {
	return ""
}

func (l *Lexer) readString() string {
	return ""
}

func (l *Lexer) readChar() {
}

func (l *Lexer) peekChar() byte {
	return ' '
}

func isLetter(ch byte) bool {
	return false
}

func isDigit(ch byte) bool {
	return false
}

func isWhitespace(ch byte) bool {
	return false
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
