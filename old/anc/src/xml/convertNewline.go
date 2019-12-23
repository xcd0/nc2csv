package main

import "strings"

func ConvertNewline(str, nlcode string) string {
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}
