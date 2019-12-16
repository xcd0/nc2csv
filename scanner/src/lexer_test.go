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
		{IDENT, "O"},
		{INT, "0001"},
		{IDENT, "(ROBO 4X)"},
		{IDENT, "(COORD=CENTER)"},
		{IDENT, "(A-90.)"},
		{IDENT, "G49"},
		{IDENT, "G69"},
		{IDENT, "M72"},
		{IDENT, "M69"},
		{NCEOF, "%"},
		{IDENT, "O"},
		{IDENT, "(ROBO 4X)"},
		{IDENT, "ROBO"},
		{INT, "4"},
		{IDENT, "X"},
		{COMMENTEND, ")"},
		{COMMENTSTART, "("},
		{IDENT, "COORD"},

		{NCEOF, "%"},
		{IDENT, "O"},
		{INT, "0001"},
		{IDENT, "(ROBO 4X)"},
		{IDENT, "(COORD=CENTER)"},
		{IDENT, "(A-90.)"},
		{IDENT, "G49"},
		{IDENT, "G69"},
		{IDENT, "M72"},
		{IDENT, "M69"},
		{NCEOF, "%"},
		{IDENT, "O"},
		{IDENT, "(ROBO 4X)"},
		{IDENT, "ROBO"},
		{INT, "4"},
		{IDENT, "X"},
		{COMMENTEND, ")"},
		{COMMENTSTART, "("},
		{IDENT, "COORD"},

		/*
			M03 S1000 (ROT_INITIAL)
			G90 G00
			G90 G10 L11 P#4120 R-1.0
			X0. Y0.
			G43 Z50.
			G90
			G01 X-6. Y0. Z50. A-90. F1000
			G00 Z31. A-90.
			G01 X-3. F5000
			X0.
			Z26. F1000
			A-90.912 F2204
						{LET, "let"},
						{IDENT, "five"},
						{ASSIGN, "="},
						{INT, "5"},
						{SEMICOLON, ";"},
						{LET, "let"},
						{IDENT, "ten"},
						{ASSIGN, "="},
						{INT, "10"},
						{SEMICOLON, ";"},
						{LET, "let"},
						{IDENT, "add"},
						{ASSIGN, "="},
						{FUNCTION, "fn"},
						{LPAREN, "("},
						{IDENT, "x"},
						{COMMA, ","},
						{IDENT, "y"},
						{RPAREN, ")"},
						{LBRACE, "{"},
						{IDENT, "x"},
						{PLUS, "+"},
						{IDENT, "y"},
						{SEMICOLON, ";"},
						{RBRACE, "}"},
						{SEMICOLON, ";"},
						{LET, "let"},
						{IDENT, "result"},
						{ASSIGN, "="},
						{IDENT, "add"},
						{LPAREN, "("},
						{IDENT, "five"},
						{COMMA, ","},
						{IDENT, "ten"},
						{RPAREN, ")"},
						{SEMICOLON, ";"},
						{BANG, "!"},
						{MINUS, "-"},
						{SLASH, "/"},
						{ASTERISK, "*"},
						{INT, "5"},
						{SEMICOLON, ";"},
						{INT, "5"},
						{LT, "<"},
						{INT, "10"},
						{GT, ">"},
						{INT, "5"},
						{SEMICOLON, ";"},
						{IF, "if"},
						{LPAREN, "("},
						{INT, "5"},
						{LT, "<"},
						{INT, "10"},
						{RPAREN, ")"},
						{LBRACE, "{"},
						{RETURN, "return"},
						{TRUE, "true"},
						{SEMICOLON, ";"},
						{RBRACE, "}"},
						{ELSE, "else"},
						{LBRACE, "{"},
						{RETURN, "return"},
						{FALSE, "false"},
						{SEMICOLON, ";"},
						{RBRACE, "}"},
						{INT, "10"},
						{EQ, "=="},
						{INT, "10"},
						{SEMICOLON, ";"},
						{INT, "10"},
						{NOT_EQ, "!="},
						{INT, "9"},
						{SEMICOLON, ";"},
		*/
		{EOF, ""},
	}

	l := New(input)

	for _ = range tests {
		tok := l.NextToken()
		fmt.Printf("{%v, \"%v\"},\n", tok.Type, tok.Literal)
	}

	/*
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
	*/
}
