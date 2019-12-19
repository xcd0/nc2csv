package lexer

type Lexer struct {
	input        string
	length       int
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.length = len([]rune(input))
	l.ReadChar()
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

	l.SkipWhitespace()

	switch l.ch {
	case '=':
		tok = NewToken(ASSIGN, l.ch)
	case '/':
		if IsNewLine(l.PeekPreChar()) {
			tok = NewToken(SKIP, l.ch)

			// コメントの文字列を読み込む
			literal := string(l.ch)
			for !IsNewLine(l.PeekChar()) {
				l.ReadChar()
				literal += string(l.ch)
			}
			tok.Literal = literal
		}
	case '(':
		// 前処理で削除したのでない
		/*
				tok = NewToken(COMMENTSTART, l.ch, 'c')
				// コメントの文字列を読み込む
				literal := string(l.ch)
				for l.PeekChar() != ')' {
					l.ReadChar()
					literal += string(l.ch)
				}
				// ) も含める
				l.ReadChar()
				literal += string(l.ch)
				tok.Literal = literal
			case ')':
				tok = NewToken(COMMENTEND, l.ch, 'c')
		*/
	case '%':
		tok = NewToken(NCEOF, l.ch)
	case '#':
		tok = NewToken(VARIABLE, l.ch)
		literal := string(l.ch)
		for IsDigit(l.PeekPreChar()) {
			l.ReadChar()
			literal += string(l.ch)
		}
		tok.Literal = literal
	case '\n':
		tok = NewToken(EOB, l.ch)
	case ';':
		tok = NewToken(EOB, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		// 先に GOTO, IF, WHILE かどうか判別して、そうでなかったらG00とかに切り分ける。
		// 謎の文字AGとか 謎の記号〇とかが入ってきたらILLIGAL
		if macroGoto(l, &tok) {
			// trueならすでにtokが編集されているので特に何もしなくていい
			// falseならtokは何もされてないのでそのまま次の判定に進む
		} else if macroIf(l, &tok) {
			// GOTO におなじ
		} else if macroWhile(l, &tok) {
			// GOTO におなじ
		} else {
			// GOTOでもIFでもWHILEでもなかったとき
			if IsLetter(l.ch) {

				// 種別の決定

				//tok.Literal = l.ReadIdentifier()
				tok.Type = LookupIdent(tok.Literal)

				if IsAxis(l.ch) {
					// XYZABCIJKUVWRなど
					tok = NewToken(AXIS, l.ch)
				}

				switch l.ch {
				case 'G':
					tok = NewToken(PREPARATION, l.ch)
				case 'M':
					tok = NewToken(MISCELLANEOUS, l.ch)
				case 'F':
					tok = NewToken(FEED, l.ch)
				case 'S':
					tok = NewToken(SPINDLE, l.ch)
				case 'T':
					tok = NewToken(TOOL, l.ch)
				case 'O':
					tok = NewToken(ONUM, l.ch)
				}

				literal := string(l.ch)
				for !IsEOB(l.PeekChar()) && //       後ろが;でない
					!IsLetter(l.PeekChar()) && //    アルファベットでない
					!IsNewLine(l.PeekChar()) && //   改行でない
					!IsWhitespace(l.PeekChar()) { // 半角空白でない
					// 間読み込む
					l.ReadChar()
					literal += string(l.ch)
				}
				// 次の文字が半角空白やセミコロンの状態で現在の文字まで読み込む
				if IsWhitespace(l.PeekChar()) || IsEOB(l.PeekChar()) {
					l.ReadChar()
				}
				tok.Literal = literal
				return tok
			} else if IsDigit(l.ch) {
				// G01X1とかでは来ない
				// 行頭いきなり数値とかでないと来ない？
				// <- #100=20とかの20で来る
				tok.Type = INT
				tok.Literal = l.ReadNumber()
				return tok
			} else {
				// 異常値
				tok = NewToken(ILLEGAL, l.ch)
				return tok
			}
		}
	}

	l.ReadChar()
	return tok
}
