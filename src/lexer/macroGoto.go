package lexer

func macroGoto(l *Lexer, tok *Token) bool {
	// GOTO // G01とかと間違わない
	if l.ch == 'G' {
		// 4文字覗く
		fourChars := l.PeekRunes(4)
		if string(fourChars) == "GOTO" {
			// GOTOと一致
			// とりあえず改行までliteralに蓄積して返す
			*tok = NewToken(GOTO, l.ch)
			literal := string(l.ch) // G
			// 行末まで読み込む
			for !IsEOB(l.PeekChar()) && //       後ろが;でない
				!IsNewLine(l.PeekChar()) { //   改行でない
				// 間読み込む
				l.ReadChar()
				literal += string(l.ch)
			}
			// 次の文字がセミコロンの状態で現在の文字まで読み込む
			if IsNewLine(l.PeekChar()) || IsEOB(l.PeekChar()) {
				l.ReadChar()
			}
			tok.Literal = literal
			return true
		}
	}
	return false
}
