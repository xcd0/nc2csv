package parser

import (
	"go/token"

	"../lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *pparser {
	p := &Parser{l: l}
	// 2つトークンを読み込む。curToeknとpeekTokenの両方がセットされる。
	p.nextToekn()
	p.nextToekn()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	p.curToekn = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *sat.Program {
	return nil
}
