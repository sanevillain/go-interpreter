package lexer

import (
	"sanevillain/go-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	// OPERATORS
	case '!': // BANG or NOTEQ
		tok = token.New(token.BANG, "!")

		if l.peekChar() == '=' {
			l.readChar()
			tok = token.New(token.NOT_EQ, "!=")
		}
	case '=': // EQ or ASSIGN
		tok = token.New(token.ASSIGN, "=")

		if l.peekChar() == '=' {
			l.readChar()
			tok = token.New(token.EQ, "==")
		}
	case '+':
		tok = token.New(token.PLUS, "+")
	case '-':
		tok = token.New(token.MINUS, "-")
	case '*':
		tok = token.New(token.ASTERISK, "*")
	case '/':
		tok = token.New(token.SLASH, "/")
	case '<':
		tok = token.New(token.LT, "<")
	case '>':
		tok = token.New(token.GT, ">")

		// DELIMITERS
	case ',':
		tok = token.New(token.COMMA, ",")
	case ';':
		tok = token.New(token.SEMICOLON, ";")
	case ':':
		tok = token.New(token.COLON, ":")
	case '(':
		tok = token.New(token.LPAREN, "(")
	case ')':
		tok = token.New(token.RPAREN, ")")
	case '{':
		tok = token.New(token.LBRACE, "{")
	case '}':
		tok = token.New(token.RBRACE, "}")
	case '[':
		tok = token.New(token.LBRACKET, "[")
	case ']':
		tok = token.New(token.RBRACKET, "]")

		// STRING LITERAL
	case '"':
		l.readChar()

		literal := l.readString()
		tok = token.New(token.STRING, literal)

		if l.ch != '"' {
			tok = token.New(token.ILLEGAL, "")
		}

		// EOF
	case 0:
		tok = token.New(token.EOF, "")

		// KEYWORDS or IDENTIFIERS + LITERALS
	default:
		if isDigit(l.ch) {
			return token.New(token.INT, l.readNumber())
		}

		if isLetter(l.ch) {
			switch literal := l.readIdentifier(); literal {
			case "fn":
				return token.New(token.FUNCTION, literal)
			case "let":
				return token.New(token.LET, literal)
			case "true":
				return token.New(token.TRUE, literal)
			case "false":
				return token.New(token.FALSE, literal)
			case "if":
				return token.New(token.IF, literal)
			case "else":
				return token.New(token.ELSE, literal)
			case "return":
				return token.New(token.RETURN, literal)
			default:
				return token.New(token.IDENT, literal)
			}
		}

		tok = token.New(token.ILLEGAL, "")
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position

	for l.ch != 0 && l.ch != '"' {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
