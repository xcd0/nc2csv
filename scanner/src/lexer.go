package main

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
	* 変数操作(#)
	* 加減乗除(X10./2みたいな感じ)
		* / はオプショナルスキップブロックがあるのでチェックする
	はLiteralに吸収されているのでここでは解釈されない。
*/
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	/*
		case '+':
			if isDigit(l.peekPreChar()) {
				// 前が数値だったら
				//ch := l.ch
				//l.readChar()
				//literal := string(ch) + string(l.ch)
				tok = newToken(ADD, l.ch)
			} else {
				tok = newToken(PLUS, l.ch)
			}
		case '-':
			if isDigit(l.peekPreChar()) {
				// 前が数値だったら
				//ch := l.ch
				//l.readChar()
				//literal := string(ch) + string(l.ch)
				tok = newToken(SUB, l.ch)
			} else {
				tok = newToken(MINUS, l.ch)
			}
		case '*':
			if isDigit(l.peekPreChar()) {
				// 前が数値だったら
				//ch := l.ch
				//l.readChar()
				//literal := string(ch) + string(l.ch)
				tok = newToken(MUL, l.ch)
			} else {
				tok = newToken(ASTERISK, l.ch)
			}
		case '/':
			if isDigit(l.peekPreChar()) {
				// 前が数値だったら
				//ch := l.ch
				//l.readChar()
				//literal := string(ch) + string(l.ch)
				tok = newToken(DIV, l.ch)
			} else if isNewLine(l.peekPreChar()) {
				tok = newToken(SKIP, l.ch)
			} else {
				tok = newToken(SLASH, l.ch)
			}
	*/
	case '/':
		if isNewLine(l.peekPreChar()) {
			tok = newToken(SKIP, l.ch)

			// コメントの文字列を読み込む
			literal := string(l.ch)
			for !isNewLine(l.peekChar()) {
				l.readChar()
				literal += string(l.ch)
			}
			// ) も含める
			//l.readChar()
			//literal += string(l.ch)
			tok.Literal = literal
		}
	case '(':
		tok = newToken(COMMENTSTART, l.ch)

		// コメントの文字列を読み込む
		literal := string(l.ch)
		for l.peekChar() != ')' {
			l.readChar()
			literal += string(l.ch)
		}
		// ) も含める
		l.readChar()
		literal += string(l.ch)
		tok.Literal = literal
	case ')':
		tok = newToken(COMMENTEND, l.ch)
	case '%':
		tok = newToken(NCEOF, l.ch)
	/*
		case '#':
			tok = newToken(VARIABLE, l.ch)
			literal := string(l.ch)
			for isDigit(l.peekPreChar()) {
				l.readChar()
				literal += string(l.ch)
			}
			tok.Literal = literal
	*/
	case '\n':
		tok = newToken(EOB, l.ch)
	case ';':
		tok = newToken(EOB, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF

	default:
		if isLetter(l.ch) {

			//tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)

			if isAxis(l.ch) {
				// XYZABCIJKUVWRなど
				tok = newToken(AXIS, l.ch)
			}

			if l.ch == 'G' {
				tok = newToken(PREPARATION, l.ch)
			}
			if l.ch == 'M' {
				tok = newToken(MISCELLANEOUS, l.ch)
			}
			if l.ch == 'F' {
				tok = newToken(FEED, l.ch)
			}
			if l.ch == 'S' {
				tok = newToken(SPINDLE, l.ch)
			}
			if l.ch == 'T' {
				tok = newToken(TOOL, l.ch)
			}
			if l.ch == 'O' {
				tok = newToken(ONUM, l.ch)
			}

			literal := string(l.ch)
			for !isEOB(l.peekChar()) && //       後ろが;でない
				!isLetter(l.peekChar()) && //    アルファベットでない
				!isNewLine(l.peekChar()) && //   改行でない
				!isWhitespace(l.peekChar()) { // 半角空白でない
				// 間読み込む
				l.readChar()
				literal += string(l.ch)
			}
			// 次の文字が半角空白やセミコロンの状態で現在の文字まで読み込む
			if isWhitespace(l.peekChar()) || isEOB(l.peekChar()) {
				l.readChar()
			}
			tok.Literal = literal
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// util {{{

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = getRuneAt(l.input, l.readPosition)
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekPreChar() rune {
	if l.position <= 0 {
		return 0
	} else {
		return getRuneAt(l.input, l.position-1)
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return getRuneAt(l.input, l.readPosition)
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isAxis(ch rune) bool {
	return ch == 'X' || ch == 'Y' || ch == 'Z' ||
		ch == 'A' || ch == 'B' || ch == 'C' ||
		ch == 'I' || ch == 'J' || ch == 'K' ||
		ch == 'U' || ch == 'V' || ch == 'W' ||
		ch == 'R'
}

func isEOB(ch rune) bool {
	return ';' == ch
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func isNewLine(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func newToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

// }}}
