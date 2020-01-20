package main

import "./util"

func readAndCutNumber(rs *[]rune, i *int) string {
	numStr := util.ReadNumbers(rs, i)
	numRunes := []rune(numStr)
	for i, n := range numStr {
		if n == '0' {
			if len(numStr) == 1 {
				// 00000みたいなのは0を消していったら
				// 全部消えてしまうので1つ残す
				return numStr
			}
			continue
		}
		if len(numStr) > 1 && numRunes[1] == '.' {
			// 0. みたいなの
			continue
		}
		// 先頭の0だけ捨てる
		numStr = numStr[i:]
		numRunes = []rune(numStr)
		break
	}
	return numStr
}
