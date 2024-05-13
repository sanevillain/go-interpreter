package lexer

import (
	"sanevillain/go-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	type Expected struct {
		Type    token.TokenType
		Literal string
	}

	testCases := []struct {
		input    string
		expected []Expected
	}{
		{
			"let five = 5;",
			[]Expected{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			"let ten = 10;",
			[]Expected{
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			`let add = fn(x, y) {
	           x + y;
	        };`,
			[]Expected{
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			"let result = add(five, ten);",
			[]Expected{
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			"!-/*5;",
			[]Expected{
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			"5 < 10 > 5;",
			[]Expected{
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			`if (5 < 10) {
                return true;
            } else {
                return false;
            }`,
			[]Expected{
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			"10 == 10;",
			[]Expected{
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			"10 != 9;",
			[]Expected{
				{token.INT, "10"},
				{token.NOT_EQ, "!="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			`"foobar"`,
			[]Expected{
				{token.STRING, "foobar"},
				{token.EOF, ""},
			},
		},
		{
			`"foo bar"`,
			[]Expected{
				{token.STRING, "foo bar"},
				{token.EOF, ""},
			},
		},
		{
			"[1, 2];",
			[]Expected{
				{token.LBRACKET, "["},
				{token.INT, "1"},
				{token.COMMA, ","},
				{token.INT, "2"},
				{token.RBRACKET, "]"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			`{"foo": bar}`,
			[]Expected{
				{token.LBRACE, "{"},
				{token.STRING, "foo"},
				{token.COLON, ":"},
				{token.IDENT, "bar"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			l := New(tC.input)

			for _, expected := range tC.expected {
				tok := l.NextToken()

				if tok.Type != expected.Type {
					t.Fatalf("tokentype wrong. expected=%q, got=%q", expected.Type, tok.Type)
				}

				if tok.Literal != expected.Literal {
					t.Fatalf("literal wrong. expected=%q, got=%q", expected.Literal, tok.Literal)
				}
			}
		})
	}
}
