package ast

import "sanevillain/go-interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST the parser produces.
type Program struct {
	Statements []Statement // Every program is a series of statements.
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) <= 0 {
		return ""
	}

	return p.Statements[0].TokenLiteral()
}

type Identifier struct {
	Value string
	Token token.Token // the token.IDENT token
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
