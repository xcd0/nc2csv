package util

func DeleteComment(input string) string {
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
