package util

func ReadLetters(rs *[]rune, i int) string {
	pre := i
	post := i
	for IsLetter((*rs)[post]) {
		post++
	}
	return string((*rs)[pre:post])
}
