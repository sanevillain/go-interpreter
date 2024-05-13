package parser

import (
	"sanevillain/go-interpreter/ast"
	"sanevillain/go-interpreter/lexer"
	"sanevillain/go-interpreter/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // == or !=
	LESSGREATER // > or <
	SUM         // + or -
	PRODUCT     // * or /
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQ:     EQUALS,
	token.NOT_EQ: EQUALS,

	token.LT: LESSGREATER,
	token.GT: LESSGREATER,

	token.PLUS:  SUM,
	token.MINUS: SUM,

	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,

	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

type (
	prefixParseFn func() ast.Expression
	// the argument is the "left side" of the operator
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	prefixParseFuncs map[token.TokenType]prefixParseFn
	infixParseFuncs  map[token.TokenType]infixParseFn

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	return nil
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
}

func (p *Parser) Errors() []string {
	return nil
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func (p *Parser) parseStatement() ast.Statement {
	return nil
}

func (p *Parser) parseLetStatement() ast.Statement {
	return nil
}

func (p *Parser) parseReturnStatement() ast.Statement {
	return nil
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return nil
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	return nil
}

func (p *Parser) parseBoolean() ast.Expression {
	return nil
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	return nil
}

func (p *Parser) parseHashLiteral() ast.Expression {
	return nil
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	return nil
}

func (p *Parser) parseIfExpression() ast.Expression {
	return nil
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	return nil
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	return nil
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	return nil
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	return nil
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	return nil
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	return nil
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	return nil
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	return nil
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	return nil
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
}

func (p *Parser) peekPrecedence() int {
	return 0
}

func (p *Parser) curPrecedence() int {
	return 0
}

func (p *Parser) parseIdentifier() ast.Expression {
	return nil
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return nil
}

func (p *Parser) nextToken() {
}

// expectPeek advances the parser on to the next token only
// if the supplied token argument matches with the current token type.
func (p *Parser) expectPeek(t token.TokenType) bool {
	return false
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return false
}

func (p *Parser) peekError(t token.TokenType) {
}
