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
	Token token.Token // let
	Name  *Identifier
	Value Expression
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
	Token       token.Token // return
	ReturnValue Expression
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
	Token      token.Token // the first token of the expression
	Expression Expression
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
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // !, -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var b strings.Builder

	b.WriteString("(")
	b.WriteString(pe.Operator)
	b.WriteString(pe.Right.String())
	b.WriteString(")")

	return b.String()
}

type InfixExpression struct {
	Token    token.Token // +, -, /, *, ==, !=, <, >
	Left     Expression
	Operator string
	Right    Expression
}

func (pe *InfixExpression) expressionNode()      {}
func (pe *InfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *InfixExpression) String() string {
	var b strings.Builder

	b.WriteString("(")
	b.WriteString(pe.Left.String())
	b.WriteString(" ")
	b.WriteString(pe.Operator)
	b.WriteString(" ")
	b.WriteString(pe.Right.String())
	b.WriteString(")")

	return b.String()
}

type Boolean struct {
	Token token.Token // true, false
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token // if
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var b strings.Builder

	b.WriteString("if")
	b.WriteString(ie.Condition.String())
	b.WriteString(" ")
	b.WriteString(ie.Consequence.String())

	// else if optional
	if ie.Alternative != nil {
		b.WriteString("else ")
		b.WriteString(ie.Alternative.String())
	}

	return b.String()
}

type BlockStatement struct {
	Token      token.Token // {
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var b strings.Builder

	for _, s := range bs.Statements {
		b.WriteString(s.String())
	}

	return b.String()
}

type FunctionLiteral struct {
	Token      token.Token // fn
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var b strings.Builder

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	b.WriteString(fl.TokenLiteral())
	b.WriteString("(")
	b.WriteString(strings.Join(params, ", "))
	b.WriteString(") ")
	b.WriteString(fl.Body.String())

	return b.String()
}

type CallExpression struct {
	Token     token.Token // (
	Function  Expression  // Identifier of Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var b strings.Builder

	params := []string{}
	for _, p := range ce.Arguments {
		params = append(params, p.String())
	}

	b.WriteString(ce.Function.String())
	b.WriteString("(")
	b.WriteString(strings.Join(params, ", "))
	b.WriteString(")")

	return b.String()
}

type ArrayLiteral struct {
	Token    token.Token // [
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var b strings.Builder

	elements := []string{}
	for _, p := range al.Elements {
		elements = append(elements, p.String())
	}

	b.WriteString("[")
	b.WriteString(strings.Join(elements, ", "))
	b.WriteString("]")

	return b.String()
}

type IndexExpression struct {
	Token token.Token // [
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var b strings.Builder

	b.WriteString("(")
	b.WriteString(ie.Left.String())
	b.WriteString("[")
	b.WriteString(ie.Index.String())
	b.WriteString("])")

	return b.String()
}

type HashLiteral struct {
	Token token.Token // {
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var b strings.Builder

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	b.WriteString("{")
	b.WriteString(strings.Join(pairs, ", "))
	b.WriteString("}")

	return b.String()
}
