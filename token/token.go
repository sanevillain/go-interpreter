package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENTIFIERS + LITERALS
	IDENT = "IDENT"
	INT   = "INT"

	// OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	// DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}

	return IDENT
}
