package parser

import (
	"fmt"
	"testing"

	"../ast"
	"../lexer"
)

func TestAssignStatements(t *testing.T) {

	input := `
#1000 = 10
#100 = 11
#101 = 12
`

	l := lexer.NewLexer(input)

	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"X1"},
		{"Y2."},
		{"Z003"},
		{"A004.000"},
		{"#1000=1"},
		{"#100=2"},
		{"#101=3"},
	}
	for i, tt := range tests {
		s := program.Statements[i]
		fmt.Printf("program.Statements[%v] : %v\n", i, s)
		fmt.Printf("program.Statements[%v].TokenLiteral() : %v\n", i, s.TokenLiteral())
		if !testAssignStatement(t, s, tt.expectedIdentifier) {
			return
		}
	}
}

func testAssignStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "=" {
		t.Errorf("s.TokenLiteral not '='. got=%T", s)
	}
	assignStatement, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("s not *ast.AssignStatement. got=%T", s)
		return false
	}
	if assignStatement.Name.Value != name {
		t.Errorf("assignStatement.Neame.Value not '%s'. got=%s",
			name,
			assignStatement.Name.Value,
		)
		return false
	}
	if assignStatement.Name.TokenLiteral() != name {
		t.Errorf("assignStatement.Neame.TokenLiteral() not '%s'. got=%s",
			name,
			assignStatement.Name.TokenLiteral(),
		)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error : %q", msg)
	}
	t.FailNow()
}

func TestGotoStatements(t *testing.T) {
	input := `
GOTO 10
GOTO020
GOTO030.0
GOTO40GOTO50
`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 5 {
		t.Errorf("len(program.Statements) does not contain 5 statements. got='%d'.",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		gs, ok := stmt.(*ast.GotoStatement)
		if !ok {
			t.Errorf(" stmt not *ast.GotoStatement. got=%T.", stmt)
		}
		if gs.TokenLiteral() != "GOTO" {
			gs, _ := stmt.(*ast.AssignStatement)
			t.Errorf("gs.TokenLiteral() not 'GOTO'. got=%s",
				gs.TokenLiteral(),
			)
		}
	}
}

// {{{
/*
var rowInput = `%
O0001(ROBO 4X)
(COORD=CENTER)
(A-90.)
G49
G69

M72
M69
M03S1000(ROT_INITIAL)
G90G00
G90G10L11P#4120R-1.0
X0.Y0.
G43Z50.
G90
G01X-6.Y0.Z50.A-90.F1000
G00Z31.A-90.
G01X-3.F5000
X0.
Z26.F1000
A-90.912F2204

( 変数代入のテスト )
#1000=3.141592653589793238

( if文のテスト )
#100=0
IF[#100EQ0]GOTO5
IF[#100NE0]THEN#100=0

( while文のテスト )

#100 = 2
#101 = 1
WHILE [#100 GE 0] DO 1
WHILE [#101 GE 0] DO 2
#101 = #101 - 1
END 2
#101 = 1
#100 = #100 - 1
END 1

( goto文のテスト )
GOTO3
(eof)
`
// }}} */