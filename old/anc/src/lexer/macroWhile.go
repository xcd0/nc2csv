package lexer

import (
	"strings"

	"../token"
)

func macroWhile(l *Lexer, tok *token.Token) bool {
	// WHILE // W1.とかと間違わない
	if l.ch == 'W' {
		// 5文字覗く
		fiveChars := l.PeekRunes(5)
		if string(fiveChars) == "WHILE" {
			// WHILEと一致
			// とりあえず END 数値 までliteralに蓄積して返す
			// WHILEは3重まで入れ子にできるので
			// 数値は最初のDO 数値 を読み取る
			runWhile(l, tok)
			return true
		}
	}
	return false
}

func runWhile(l *Lexer, tok *token.Token) {
	*tok = token.NewToken(token.WHILE, l.ch)
	literal := string(l.ch) // W

	// WHILE [#100 GE 0] DO 1
	// のようになっているので最初のWHILEの行の行末の数値を読み取る

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

	// これがWHILEの最上位の階層を表す数値
	whileNum := literal[len(literal)-1:]

	/*
		fmt.Println("-----")
		fmt.Printf("%v\n", literal)
		fmt.Println("-----")
		fmt.Printf("whileNum : %v\n", whileNum)
	*/

	// ↑の最後に読み込んだ数値(1~3)の
	// ENDまで読み込む
	tmp := ""
	for !l.IsEOF() {
		//fmt.Printf("(%v,%v)\n", l.ReadPosition, len(l.input))

		//fmt.Printf(":%v\n", string(getRuneAt(l.input, l.ReadPosition)))

		for !IsEOB(l.PeekChar()) && !IsNewLine(l.PeekChar()) { // 行末まで読み込む
			l.ReadChar()
			literal += string(l.ch)
			tmp += string(l.ch)
		}
		if IsNewLine(l.PeekChar()) {
			l.ReadChar()
			//literal += string(l.ch)
			literal += string("\\n")
			tmp += string(l.ch)
		}
		if IsEOB(l.PeekChar()) {
			l.ReadChar()
			literal += string(l.ch)
			tmp += string(l.ch)
		}
		// これで一行読み込んだ
		// literalの最後4文字にEND1みたいなのが入っているか調べる

		// 途中に空白があったり行末にセミコロンがあったりなかったりする
		tmp2 := strings.Replace(tmp, " ", "", -1)
		tmp = strings.Replace(tmp2, ";", "", -1)

		rs := []rune(tmp)
		if len(rs) < 4 {
			continue
		}
		last4 := rs[len(rs)-4:]
		//fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
		//fmt.Println(string(last4))
		//fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
		if string(last4[:3]) == "END" {
			if string(last4[3]) == whileNum {
				// 一致した
				break
			}
		}

	}

	tok.Literal = literal
}
