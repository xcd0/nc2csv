package lexer

func GetRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}
