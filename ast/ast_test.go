package ast

import (
	"sanevillain/go-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	expect := "let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Value: "myVar",
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
				},
				Value: &Identifier{
					Value: "anotherVar",
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
				},
			},
		},
	}

	if str := program.String(); str != expect {
		t.Errorf("program.String() wrong, got=%q", str)
	}
}
