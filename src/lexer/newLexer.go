package lexer

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.length = len([]rune(input))
	l.ReadChar()
	return l
}
