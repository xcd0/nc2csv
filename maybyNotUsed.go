package main

func deleteDoubleNewline(input string) string {
	//アルファベットの前に半角空白を入れる
	// 行頭は入れない
	// アルファベットが続くときは入れない
	output := ""
	rs := []rune(input)
	var post rune
	for _, r := range rs {
		if post == '\n' && r == '\n' {
			post = r
			continue
		}
		output += string(r)
		post = r
	}
	return output
}

func deleteComment(input string) string {
	output := ""
	rs := []rune(input)
	flagCommentStart := false
	for _, r := range rs {
		// ()はコメント 複数行コメントにも対応
		if r == '(' {
			flagCommentStart = true
			continue
		}
		if flagCommentStart {
			if r == ')' {
				flagCommentStart = false
			}
			continue
		}
		output += string(r)
	}
	return output
}
