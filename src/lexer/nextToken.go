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
		// 変数代入扱いだが配列へのアクセスと見なす方が現代の解釈では適当。
		// なので配列へのアクセスとする。
		// #をポインタとし、#のあとの数値を配列のインデックスと見なす
		tok = token.NewToken(token.ARRAY, l.ch)
		/*
			literal := string(l.ch)
			for IsDigit(l.PeekPreChar()) {
				l.ReadChar()
				literal += string(l.ch)
			}
			tok.Literal = literal
		*/
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
				/* //古いコード {{{
				// 種別の決定
				tok.Literal = l.ReadIdentifier()
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
				*/ /// }}}

				// M08 G00 X1.
				// こういうのはすべて変数への代入と見なす
				// var M = 8 みたいなかんじ
				// golang の場合変数宣言と代入は var <identifier> = <expression> になる
				// NCの場合varに当たるprefixがないので
				// GOTO,IF,WHILEでないアルファベットが来たら変数代入だとみなす
				// 例えば異常値だがAAと来たらAAも変数代入になる

				// readIdentifier()はアルファベットまたは_がつづく間読み取って返す
				tok.Literal = l.readIdentifier()
				// LookUpIdentはGOTOやIF,WHILE,ELSE,ENDのような予約語を予約語として
				// それ以外をIDENTとして返す
				tok.Type = token.LookupIdent(tok.Literal)
			} else if IsDot(l.ch) {
				// 小数点
				tok.Type = token.FLOAT
				tok.Literal, _ = l.ReadNumber()
				return tok
			} else if IsDigit(l.ch) {
				// 数値
				// 小数点があるかどうかわからないと確定しない
				tok.Literal, n = l.ReadNumber()
				if n.Int {
					tok.Type = token.INT
				} else {
					tok.Type = token.FLOAT
				}
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
