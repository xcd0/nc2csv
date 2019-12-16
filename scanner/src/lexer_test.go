package main

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `%
O0001 (ROBO 4X)
(COORD=CENTER)
(A-90.)
G49
/G49
G69
M72
M69
M03 S1000 (ROT_INITIAL)
G90 G00
G90 G10 L11 P#4120 R-1.0
X0. Y0.
G43 Z50.
G90
G01 X-6. Y0. Z50. A-90. F1000
G00 Z31. A-90.
G01 X-3. F5000
X-3./2
X0.
Z26. F1000
A-90.912 F2204
`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{NCEOF, "%"},
		{ONUM, "O0001"},
		{COMMENTSTART, "(ROBO 4X)"},
		{COMMENTSTART, "(COORD=CENTER)"},
		{COMMENTSTART, "(A-90.)"},
		{PREPARATION, "G49"},
		{SKIP, "/G49"},
		{PREPARATION, "G69"},
		{MISCELLANEOUS, "M72"},
		{MISCELLANEOUS, "M69"},
		{MISCELLANEOUS, "M03"},
		{SPINDLE, "S1000"},
		{COMMENTSTART, "(ROT_INITIAL)"},
		{PREPARATION, "G90"},
		{PREPARATION, "G00"},
		{PREPARATION, "G90"},
		{PREPARATION, "G10"},
		{IDENT, "L11"},
		{IDENT, "P#4120"},
		{AXIS, "R-1.0"},
		{AXIS, "X0."},
		{AXIS, "Y0."},
		{PREPARATION, "G43"},
		{AXIS, "Z50."},
		{PREPARATION, "G90"},
		{PREPARATION, "G01"},
		{AXIS, "X-6."},
		{AXIS, "Y0."},
		{AXIS, "Z50."},
		{AXIS, "A-90."},
		{FEED, "F1000"},
		{PREPARATION, "G00"},
		{AXIS, "Z31."},
		{AXIS, "A-90."},
		{PREPARATION, "G01"},
		{AXIS, "X-3."},
		{FEED, "F5000"},
		{AXIS, "X-3./2"},
		{AXIS, "X0."},
		//{EOF, ""},
	}

	l := New(input)

	/*
		for _ = range tests {
			tok := l.NextToken()
			fmt.Printf("{%v, \"%v\"},\n", tok.Type, tok.Literal)
		}
	*/

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("%v %v %v\n", i, tt, tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
