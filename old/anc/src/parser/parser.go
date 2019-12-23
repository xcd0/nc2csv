package parser

import (
	"fmt"

	"../ast"
	"../lexer"
	"../token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// 2つトークンを読み込む。curToeknとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead.",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
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
	case token.ASSIGN:
		return p.parseAssignStatement()
	case token.GOTO:
		return p.parseGotoStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	statement := &ast.AssignStatement{Token: p.curToken}
	// 2パターンある
	// 1. X1.0  Y#10
	// 2. #10=1.0
	// 2のパターンがややこしく、#が来ただけでは参照なのか代入なのかわからない。
	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) parseGotoStatement() *ast.GotoStatement {
	statement := &ast.GotoStatement{Token: p.curToken}
	for !p.curTokenIs(token.EOB) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
