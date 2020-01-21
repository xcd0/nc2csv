package nc2csv

import "strconv"

// オプショナルスキップブロックについての処理
func forOptionalSkipBlock(rs *[]rune, i int) { // {{{
	// オプショナルスキップブロック
	// オプショナルスキップはメモリの値を読んで
	// trueなら無視する。

	if t := (*rs)[i+1]; '1' <= t && t <= '9' {
		// 番号付きオプショナルスキップブロック
		tmpNum, _ := strconv.Atoi(string(t))
		if optionalSkip[tmpNum] == false {
			// 無視しない
			setting.IsOptionalSkip = false
		} else {
			// 無視する
			setting.IsOptionalSkip = true
		}
	} else {
		if optionalSkip[0] {
			// 無視しない
			setting.IsOptionalSkip = false
		} else {
			// 無視する
			setting.IsOptionalSkip = true
		}
	}
} // }}}
