package main

import "strings"

func AnalizeString(input string) string {
	// 改行で分ける
	lines := strings.Split(input, "\n")

	output := ""

	type state struct {
		x    float64
		y    float64
		z    float64
		a    float64
		b    float64
		c    float64
		mode string
	}

	for _, line := range lines { // 一行ずつ
		// スキップする行ならスキップ
		if RegComment.MatchString(line) ||
			RegN.MatchString(line) ||
			RegO.MatchString(line) ||
			RegM1.MatchString(line) ||
			false {
			continue
		}
		if RegO.MatchString(line) {
			// O番号 スキップ
			continue
		}
	}

	return output
}
