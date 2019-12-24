package util

import "log"

func ReadNumbers(rs *[]rune, i *int) string {
	pre := *i
	post := *i
	for (IsDigit((*rs)[post]) || IsDot((*rs)[post])) && post < len(*rs)-1 {
		post++
		log.Printf("len(rs) : %d ; post : %d", len(*rs), post)
	}
	// iを進めておく
	*i = post
	log.Printf("len(rs) : %d ; last : %c; ret: %s", len(*rs), (*rs)[len(*rs)-1:][0], string((*rs)[pre:]))

	if post == len(*rs)-1 && string((*rs)[len(*rs)-1:]) != "\n" {
		// ファイル終端で改行が含まれていないなら、更に+1する
		*i = post + 1
		if IsDot((*rs)[post]) {
			// 小数点で終わっていたら0を付与する
			return string((*rs)[pre:]) + "0"
		} else {
			return string((*rs)[pre:])
		}
	}

	if IsDot((*rs)[post]) {
		// 小数点で終わっていたら0を付与する
		return string((*rs)[pre:post]) + "0"
	} else {
		return string((*rs)[pre:post])
	}
}
