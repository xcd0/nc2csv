package lexer

func (l *Lexer) IsEOF() bool {
	return l.position >= l.length
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
