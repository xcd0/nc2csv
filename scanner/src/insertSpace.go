package main

import (
	"strings"
)

func isAZ(runeChar rune) bool {
	return 'A' <= runeChar && 'Z' >= runeChar
}

func insertSpace(input string) string {
	//アルファベットの前に半角空白を入れる
	// 行頭は入れない
	// アルファベットが続くときは入れない
	output := ""

	flagCommentStart := false
	preChar := '\n'

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// 行内の文字を1文字づつチェック
		for j, runeChar := range line {

			// コメントの中身はそのまま出力する
			// ()はコメント 複数行コメントにも対応
			if runeChar == '(' {
				flagCommentStart = true
				// ただし(の前には入れる
				if j != 0 {
					output += " "
				}
			}
			if flagCommentStart {
				if runeChar == ')' {
					flagCommentStart = false
				}
				// 空白を入れない
				output += string(runeChar)
				continue
			}

			// 行頭でなく、一つ前の文字がアルファベットでない、アルファベット
			if j != 0 && !isAZ(preChar) && isAZ(runeChar) {
				output += " "
			}
			output += string(runeChar)
			preChar = runeChar
		}
		output += "\n" // 改行なくなってるので追加
	}
	return output
}
