package main

func DeleteComment(input string) string {
	//アルファベットの前に半角空白を入れる
	// 行頭は入れない
	// アルファベットが続くときは入れない
	output := ""

	flagCommentStart := false

	rs := []rune(input)

	for _, r := range rs {
		// コメントの中身はそのまま出力する
		// ()はコメント 複数行コメントにも対応
		if r == '(' {
			flagCommentStart = true
			// ただし(の前には入れる
			continue
		}
		if flagCommentStart {
			if r == ')' {
				flagCommentStart = false
			}
			// 空白を入れない
			continue
		}
		// 行頭でなく、一つ前の文字がアルファベットでない、アルファベット
		output += string(r)
	}
	return output
}
