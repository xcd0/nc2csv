package lexer

func (l *Lexer) SkipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func (l *Lexer) IsEOF() bool {
	return l.position >= l.length
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= l.length {
		l.ch = 0
	} else {
		l.ch = GetRuneAt(l.input, l.readPosition)
	}
	l.position = l.readPosition
	l.readPosition += 1
}

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

func (l *Lexer) ReadIdentifier() string {
	position := l.position
	for IsLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) ReadNumber() string {
	position := l.position
	for IsDigit(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func IsAxis(ch rune) bool {
	return ch == 'X' || ch == 'Y' || ch == 'Z' ||
		ch == 'A' || ch == 'B' || ch == 'C' ||
		ch == 'I' || ch == 'J' || ch == 'K' ||
		ch == 'U' || ch == 'V' || ch == 'W' ||
		ch == 'R'
}

func IsEOB(ch rune) bool {
	return ';' == ch
}

func IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func IsNewLine(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func NewToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func GetRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}
