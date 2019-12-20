package parser

import "../lexer/token"

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// 全てのノードはNodeインターフェースを実装しなければならない
type Node interface {
	TokenLiteral() string // ノードが関連付けられているトークンのリテラル値を返す
}

// 式
type Statement interface {
	Node
	StatementNode()
}

// 値
type Expression interface {
	Node
	expressionNode()
}

// 代入式
type AssignStatement struct {
	Toke  token.Token // token.ASSIGN
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
