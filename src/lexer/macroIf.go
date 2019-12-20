package lexer

import (
	"../token"
)

func macroIf(l *Lexer, tok *token.Token) bool {
	// if
	if l.ch == 'I' && l.PeekChar() == 'F' {
		*tok = token.NewToken(token.IF, l.ch)
		literal := string(l.ch)
		l.ReadChar()
		literal += string(l.ch)
		tok.Literal = literal

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
	return false
}
