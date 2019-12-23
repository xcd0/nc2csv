package util

func ReadNumbers(rs *[]rune, i int) string {
	pre := i
	post := i
	for IsDigit(rs[post]) || IsDot(rs[post]) {
		post++
	}
	return string(rs[pre : post+1])
}
