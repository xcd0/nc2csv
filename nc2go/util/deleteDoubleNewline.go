package util

func DeleteDoubleNewline(input string) string {
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
