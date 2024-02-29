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
	Name  *Identifier // lfs
	Value Expression  // rhs
	Token token.Token // the 'let' token
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ReturnStatement is of the form: return <expr>;
type ReturnStatement struct {
	ReturnValue Expression
	Token       token.Token // the 'return' token
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct {
	Expression Expression
	Token      token.Token // the first token of the expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

type Identifier struct {
	Value string
	Token token.Token // the token.IDENT token
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
