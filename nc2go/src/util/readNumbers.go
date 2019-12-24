package util

func ReadNumbers(rs *[]rune, i *int) string {
	pre := *i
	post := *i
	for IsDigit((*rs)[post]) || IsDot((*rs)[post]) {
		post++
	}
	// iを進めておく
	*i = post
	if IsDot((*rs)[post]) {
		// 小数点で終わっていたら0を付与する
		return string((*rs)[pre:post]) + "0"
	} else {
		return string((*rs)[pre:post])
	}
}
