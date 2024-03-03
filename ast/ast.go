package ast

import (
	"sanevillain/go-interpreter/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var b strings.Builder

	for _, stmt := range p.Statements {
		b.WriteString(stmt.String())
	}

	return b.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) <= 0 {
		return ""
	}

	return p.Statements[0].TokenLiteral()
}

// LetStatement is of the form: let <identifier> = <expr>;
type LetStatement struct {
	Name  *Identifier
	Value Expression
	Token token.Token // the 'let' token
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var b strings.Builder

	b.WriteString(ls.TokenLiteral())
	b.WriteString(" ")
	b.WriteString(ls.Name.String())
	b.WriteString(" = ")

	if ls.Value != nil {
		b.WriteString(ls.Value.String())
	}

	b.WriteString(";")

	return b.String()
}

// ReturnStatement is of the form: return <expr>;
type ReturnStatement struct {
	ReturnValue Expression
	Token       token.Token // the 'return' token
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var b strings.Builder

	b.WriteString(rs.TokenLiteral())
	b.WriteString(" ")

	if rs.ReturnValue != nil {
		b.WriteString(rs.ReturnValue.String())
	}

	b.WriteString(";")

	return b.String()
}

type ExpressionStatement struct {
	Expression Expression
	Token      token.Token // the first token of the expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Value string
	Token token.Token // the token.IDENT token
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
