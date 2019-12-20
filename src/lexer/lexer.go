package lexer

type Lexer struct {
	input        string
	length       int
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}
