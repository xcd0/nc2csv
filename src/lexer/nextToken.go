package lexer

import "../token"

/*
	* 変数操作(#)
	* 加減乗除(X10./2みたいな感じ)
		* / はオプショナルスキップブロックがあるのでチェックする
	はLiteralに吸収されているのでここでは解釈されない。
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.SkipWhitespace()

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case '/':
		// オプショナルスキップブロック
		// /,/1,/2,/3,/4,/5が来る可能性がある。
		// /#1 は#1=10にオプショナルスキップがついたものとみなす。
		tok = token.NewToken(token.SKIP, l.ch)
		literal := string(l.ch)
		// /の後に数値が来たら
		if IsDigit(l.PeekChar()) {
			// 一字だけ読み込む
			l.ReadChar()
			literal += string(l.ch)
		}
		tok.Literal = literal
	case '(':
		// 前処理で削除したのでない
		panic("Error : 前処理で削除したはずの'('が見つかりました。前処理プログラムのバグです。")
		/*
				tok = token.NewToken(token.COMMENTSTART, l.ch, 'c')
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
				tok = token.NewToken(token.COMMENTEND, l.ch, 'c')
		*/
	case '%':
		tok = token.NewToken(token.NCEOF, l.ch)
	case '#':
		// 変数
		// #数値 の#と数値含めて変数名とする
		tok = token.NewToken(token.VARIABLE, l.ch)
		literal := string(l.ch)
		for IsDigit(l.PeekPreChar()) {
			l.ReadChar()
			literal += string(l.ch)
		}
		tok.Literal = literal
	case '\n':
		tok = token.NewToken(token.EOB, l.ch)
	case ';':
		tok = token.NewToken(token.EOB, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
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
				tok.Type = token.LookupIdent(tok.Literal)

				if IsAxis(l.ch) {
					// XYZABCIJKUVWRなど
					tok = token.NewToken(token.AXIS, l.ch)
				}

				switch l.ch {
				case 'G':
					tok = token.NewToken(token.PREPARATION, l.ch)
				case 'M':
					tok = token.NewToken(token.MISCELLANEOUS, l.ch)
				case 'F':
					tok = token.NewToken(token.FEED, l.ch)
				case 'S':
					tok = token.NewToken(token.SPINDLE, l.ch)
				case 'T':
					tok = token.NewToken(token.TOOL, l.ch)
				case 'O':
					tok = token.NewToken(token.ONUM, l.ch)
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
				tok.Type = token.INT
				tok.Literal = l.ReadNumber()
				return tok
			} else {
				// 異常値
				tok = token.NewToken(token.ILLEGAL, l.ch)
				return tok
			}
		}
	}

	l.ReadChar()
	return tok
}
