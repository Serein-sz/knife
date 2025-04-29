package parser

import (
	"fmt"

	"github.com/Serein-sz/knife/ast"
	"github.com/Serein-sz/knife/lexer"
	"github.com/Serein-sz/knife/token"
)

const (
	_ = iota
	LOWEST
	EQUALS // ==
	// false == (2 < 3)
	LESS_GREATER // > <
	// 1 + (2 * 3)
	SUM
	PRODUCT
	// (-X) * Y
	PREFIX // !X -X
	CALL   // -x + (foo(1,2))
	INDEX
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESS_GREATER,
	token.GT:       LESS_GREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.ASTERISK: PRODUCT,
	token.SLASH:    PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

type (
	prefixHandlerFunc func() ast.Expression
	infixHandlerFunc  func(lhs ast.Expression) ast.Expression
)

type Parser struct {
	l                    lexer.Lexer
	curToken             token.Token
	peekToken            token.Token
	prefixHandlerFuncMap map[token.TokenType]prefixHandlerFunc
	infixHandlerFuncMap  map[token.TokenType]infixHandlerFunc
	errors               []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: *l, curToken: token.Token{}, peekToken: token.Token{}}

	p.prefixHandlerFuncMap = make(map[token.TokenType]prefixHandlerFunc)
	p.infixHandlerFuncMap = make(map[token.TokenType]infixHandlerFunc)
	p.prefixHandlerFuncMap[token.IDENT] = p.parseIdentifier
	p.prefixHandlerFuncMap[token.BANG] = p.parsePrefixExpression
	p.prefixHandlerFuncMap[token.NUMBER] = p.parseNumberLiteral
	p.prefixHandlerFuncMap[token.STRING] = p.parseStringLiteral
	p.infixHandlerFuncMap[token.PLUS] = p.parseInfixExpression
	p.infixHandlerFuncMap[token.MINUS] = p.parseInfixExpression
	p.infixHandlerFuncMap[token.ASTERISK] = p.parseInfixExpression
	p.infixHandlerFuncMap[token.SLASH] = p.parseInfixExpression
	p.infixHandlerFuncMap[token.LPAREN] = p.parseInfixExpression

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Error() error {
	if len(p.errors) == 0 {
		return nil
	}

	var s string
	for _, msg := range p.errors {
		s += "\t" + msg + "\n"
	}
	return fmt.Errorf("parser error: %v", s)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.peekToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.FUNCTION:
		return p.parseFunctionDefineStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseFunctionDefineStatement() ast.Statement {
	functionDefineStatement := &ast.FunctionDefineStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	functionDefineStatement.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	functionDefineStatement.Parameters = p.parseFunctionDefineParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	functionDefineStatement.Body = p.parseBlockStatement()

	if !p.curTokenTypeIs(token.RBRACE) {
		panic("the { is not closed")
	}
	return functionDefineStatement
}

func (p *Parser) parseFunctionDefineParameters() []*ast.Identifier {
	if p.peekTokenTypeIs(token.RPAREN) {
		// no params
		p.nextToken()
		return nil
	}
	p.nextToken()
	identifiers := []*ast.Identifier{p.parseIdentifier().(*ast.Identifier)}
	for p.peekTokenTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		identifiers = append(identifiers, p.parseIdentifier().(*ast.Identifier))
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return identifiers
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	blockStatement := &ast.BlockStatement{Token: p.curToken}
	p.nextToken()
	for !p.curTokenTypeIs(token.RBRACE) && !p.curTokenTypeIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			blockStatement.Statements = append(blockStatement.Statements, statement)
		}
		p.nextToken()
	}
	return blockStatement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	returnStatement := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	returnStatement.Value = p.parseExpression(LOWEST)
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return returnStatement
}

func (p *Parser) parseFunctionCallExpression(lhs ast.Expression) ast.Expression {
	functionCallExpression := &ast.FunctionCallExpression{
		Token:    p.curToken,
		Function: lhs,
	}
	functionCallExpression.Arguments = p.parseExpressionList(token.RPAREN)
	return functionCallExpression
}

func (p *Parser) parseExpressionList(close token.TokenType) []ast.Expression {
	if p.peekTokenTypeIs(close) {
		p.nextToken()
		return nil
	}
	p.nextToken()
	expression := p.parseExpression(LOWEST)
	expressions := []ast.Expression{expression}
	for p.peekTokenTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		expressions = append(expressions, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(close) {
		return nil
	}
	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}
	return expressions
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expressionStatement := &ast.ExpressionStatement{
		Token:      p.curToken,
		Expression: p.parseExpression(LOWEST),
	}
	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}
	return expressionStatement
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	letStatement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()
	letStatement.Value = p.parseExpression(LOWEST)
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return letStatement
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixHandler, ok := p.prefixHandlerFuncMap[p.curToken.Type]
	if !ok {
		msg := fmt.Sprintf("undefined prefix operator: %q", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	lhs := prefixHandler()
	if !p.peekTokenTypeIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infixHandler, ok := p.infixHandlerFuncMap[p.peekToken.Type]
		if ok {
			p.nextToken()
			lhs = infixHandler(lhs)
		}
	}
	return lhs
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	prefixExpression := &ast.PrefixExpression{
		Token: p.curToken,
		Op:    p.curToken.Literal,
		Rhs:   nil,
	}
	p.nextToken()
	prefixExpression.Rhs = p.parseExpression(PREFIX)
	return prefixExpression
}

func (p *Parser) parseInfixExpression(lhs ast.Expression) ast.Expression {
	if p.curTokenTypeIs(token.LPAREN) { // function calling
		return p.parseFunctionCallExpression(lhs)
	}
	infixExpression := &ast.InfixExpression{
		Token: p.curToken,
		Lhs:   lhs,
		Op:    p.curToken.Literal,
		Rhs:   nil,
	}
	curPrecedence := p.curPrecedence()
	p.nextToken()
	infixExpression.Rhs = p.parseExpression(curPrecedence)
	return infixExpression
}

func (p *Parser) parseNumberLiteral() ast.Expression {
	return &ast.NumberLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}


func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curTokenTypeIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenTypeIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be %s, but got %s", t, p.peekToken.Type))
}
