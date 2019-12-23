package lexer

func (l *Lexer) PeekPreChar() rune {
	if l.position <= 0 {
		return 0
	} else {
		return GetRuneAt(l.input, l.position-1)
	}
}

func (l *Lexer) PeekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return GetRuneAt(l.input, l.readPosition)
	}
}

func (l *Lexer) PeekRunes(num int) []rune {
	out := make([]rune, num)
	for i := 0; i < num; i++ {
		out[i] = GetRuneAt(l.input, l.position+i)
	}
	return out
}
