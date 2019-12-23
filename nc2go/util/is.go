package lexer

func isaxis(ch rune) bool {
	return ch == 'x' || ch == 'y' || ch == 'z' ||
		ch == 'a' || ch == 'b' || ch == 'c' ||
		ch == 'i' || ch == 'j' || ch == 'k' ||
		ch == 'u' || ch == 'v' || ch == 'w' ||
		ch == 'r'
}

func iseob(ch rune) bool {
	return ';' == ch
}

func islf(ch rune) bool {
	return '\n' == ch
}

func isletter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'a' <= ch && ch <= 'z' || ch == '_'
}

func isdigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isdot(ch rune) bool {
	return ch == '.'
}

func isnewline(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func iswhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func getruneat(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func getrunes(rs *[]rune, i int, num int) (string, error) {
	r = rs[i]
	if i+num < len(rs) {
		return string(*rs[i : i+num]), nil
	} else {
		return "", errors.new("範囲外アクセス")
	}
}

// keywordsにあるか調べる
func isreserved(identifier string) bool {
	// #等もkeywordsに含まれるが、
	// readlettersでアルファベットと_だけを切り出しているので該当しない。
	if token, ok := keywords[identifier]; ok {
		return true
	} else {
		return false
	}
}

func isimplemented(id string) bool {
	// 実装したら増やす
	if id == "%" {
		return true
	} else {
		return false
	}
}
