package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つトークンを読み込む。curToeknとpeekTokenの両方がセットされる。
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
