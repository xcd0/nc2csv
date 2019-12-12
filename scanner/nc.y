%{

package main

import (
	"fmt"
	"text/scanner"
	"os"
	"strings"
)

type Expression interface{}

type Token struct {
	token   int
	literal string
}

type NumExpr struct {
	literal string
}

type BinOpExpr struct {
	left     Expression
	operator rune
	right    Expression
}

%}

%union {
	token Token
}

%token <token> numbercharactor
%token <token> pm
%token <token> dot
%token <token> nceof
%token <token> skip
%token <token> eob
%token <token> commentStart
%token <token> commentEnd
%token <token> onum
%token <token> preparation
%token <token> feed
%token <token> spindle
%token <token> tool
%token <token> miscellaneous
%token <token> otherblockD
%token <token> otherblockH
%token <token> otherblockP
%token <token> otherblockR
%token <token> axis

%type <expr> num
%type <expr> unum
%type <expr> float
%type <expr> ufloat
%type <expr> int
%type <expr> uint
%type <expr> ncdata
%type <expr> nc
%type <expr> ncprogram
%type <expr> block
%type <expr> blockOnum
%type <expr> blockSkip
%type <expr> blockComment
%type <expr> blockPreparation
%type <expr> blockFeed
%type <expr> blockSpindle
%type <expr> blockTool
%type <expr> blockMiscellaneous
%type <expr> other

%%

numbercharactor:		'0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9'
pm:						'+' | '-'
dot:					'.'
nceof:					'%'
skip:					'/'
eob:					'\n'
commentStart:			'('
commentEnd:				')'
onum:					'O'
preparation:			'G'
feed:					'F'
spindle:				'S'
tool:					'T'
miscellaneous:			'M'
otherblockD:			'D'
otherblockH:			'H'
otherblockP:			'P'
otherblockR:			'R'
axis:					'X' | 'Y' | 'Z' | 'A' | 'B' | 'C' | 'I' | 'J' | 'K' | 'U' | 'V' | 'W'

num:					pm unum | unum
unum:					ufloat | uint
float:					int  dot | int  dot uint
ufloat:					uint  dot | uint  dot uint
int:					pm uint
uint:					numbercharactor uint | numbercharactor

ncdata:					nceof nc nceof
nc:						blockOnum ncprogram
ncprogram:				block ncprogram | block | eob

block:					blockSkip
						| blockComment block
						| blockPreparation block
						| blockFeed block
						| blockSpindle block
						| blockTool block
						| blockMiscellaneous block

blockOnum:				onum uint
blockSkip:				skip block eob
blockComment:			commentStart any commentEnd
blockPreparation:		preparation unum
blockFeed:				feed num
blockSpindle:			spindle num
blockTool:				tool unum
blockMiscellaneous:		miscellaneous unum
other:					otherblockD num
						| otherblockH num
						| otherblockP num
						| otherblockR num

%%

type Lexer struct {
	scanner.Scanner
	result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = NUMBER
	}
	lval.token = Token{token: token, literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

func main() {
	l := new(Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	yyParse(l)
	fmt.Printf("%#v\n", l.result)
}

