package ast

import (
	"bytes"
	"strconv"

	"../token"
)

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

// 式 {{{
type Statement interface {
	Node
	StatementNode()
}

// }}}

// 値 {{{
type Expression interface {
	Node
	expressionNode()
}

// }}}

// 代入式 {{{
type AssignStatement struct {
	Token token.Token // token.ASSIGN
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) StatementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// }}}

// 配列 {{{

type ArrayExpression struct {
	Token token.Token // #
	Left  Expression  // # オブジェクト
	Index Expression  // 要素番号
}

func (a *ArrayExpression) expressionNode()      {}
func (a *ArrayExpression) TokenLiteral() string { return a.Token.Literal }
func (a *ArrayExpression) String() string {
	c := a.Index.TokenLiteral()
	n, _ := strconv.Atoi(c)
	return "( #" + a.Index.TokenLiteral() + " = " + token.Hash[n].String() + " )"
}

// }}}

// GOTO {{{
type GotoStatement struct {
	Token     token.Token
	GotoValue Expression
}

func (g *GotoStatement) StatementNode()       {}
func (g *GotoStatement) TokenLiteral() string { return g.Token.Literal }

// }}}

type IfExpression struct { // {{{
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// }}}

type ExpressionStatement struct { // {{{
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

// }}}

type BlockStatement struct { // WHILE block {{{
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// }}}

type InfixExpression struct { // 中置演算子 {{{
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// }}}
