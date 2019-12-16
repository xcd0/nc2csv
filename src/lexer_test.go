package main

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	rowInput := `%
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
`

	spacedInput := insertSpace(rowInput)
	l := New(spacedInput)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{NCEOF, "%"},
		{EOB, ";"},
		{ONUM, "O0001"},
		{COMMENTSTART, "(ROBO 4X)"},
		{EOB, ";"},
		{COMMENTSTART, "(COORD=CENTER)"},
		{EOB, ";"},
		{COMMENTSTART, "(A-90.)"},
		{EOB, ";"},
		{PREPARATION, "G49"},
		{EOB, ";"},
		{PREPARATION, "G69"},
		{EOB, ";"},
		{MISCELLANEOUS, "M72"},
		{EOB, ";"},
		{MISCELLANEOUS, "M69"},
		{EOB, ";"},
		{MISCELLANEOUS, "M03"},
		{SPINDLE, "S1000"},
		{COMMENTSTART, "(ROT_INITIAL)"},
		{EOB, ";"},
		{PREPARATION, "G90"},
		{PREPARATION, "G00"},
		{EOB, ";"},
		{PREPARATION, "G90"},
		{PREPARATION, "G10"},
		{IDENT, "L11"},
		{IDENT, "P#4120"},
		{AXIS, "R-1.0"},
		{EOB, ";"},
		{AXIS, "X0."},
		{AXIS, "Y0."},
		{EOB, ";"},
		{PREPARATION, "G43"},
		{AXIS, "Z50."},
		{EOB, ";"},
		{PREPARATION, "G90"},
		{EOB, ";"},
		{PREPARATION, "G01"},
		{AXIS, "X-6."},
		{AXIS, "Y0."},
		{AXIS, "Z50."},
		{AXIS, "A-90."},
		{FEED, "F1000"},
		{EOB, ";"},
		{PREPARATION, "G00"},
		{AXIS, "Z31."},
		{AXIS, "A-90."},
		{EOB, ";"},
		{PREPARATION, "G01"},
		{AXIS, "X-3."},
		{FEED, "F5000"},
		{EOB, ";"},
		{AXIS, "X0."},
		{EOB, ";"},
		{AXIS, "Z26."},
		{FEED, "F1000"},
		{EOB, ";"},
		{AXIS, "A-90.912"},
		{FEED, "F2204"},
	}

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
