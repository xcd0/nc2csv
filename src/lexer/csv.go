package lexer

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	iconv "github.com/djimenez/iconv-go"
)

func rune2string(r rune) string {
	return fmt.Sprintf("%c", r)
}

func ReadCsv(rowInput string) [][]string {
	lines := strings.Split(rowInput, "\n")
	colNum := strings.Count(lines[0], ",") + 1
	out := make([][]string, len(lines))

	for i, line := range lines {
		if colNum != strings.Count(line, ",")+1 {
			break
		}
		tmpCol := make([]string, colNum)
		// 行内の文字を1文字づつチェック
		tmpCount := 0

		tmpRunes := make([]rune, 0)

		for _, runeChar := range line {
			flagSQ := false
			flagDQ := false
			if runeChar == ',' {
				if flagSQ || flagDQ {
					tmpRunes = append(tmpRunes, runeChar)
				} else {
					tmpCol[tmpCount] = string(tmpRunes)
					tmpCount++
					tmpRunes = make([]rune, 0)
				}
			} else if runeChar == '\'' {
				if flagSQ {
					flagSQ = false
				} else {
					flagSQ = true
				}
				// シングルクオートを追加するかどうかは考慮の余地あり
				tmpRunes = append(tmpRunes, runeChar)
			} else if runeChar == '"' {
				if flagDQ {
					flagDQ = false
				} else {
					flagDQ = true
				}
				// ダブルクオートを追加するかどうかは考慮の余地あり
				tmpRunes = append(tmpRunes, runeChar)
			} else {
				//tmpString += rune2string(runeChar)
				tmpRunes = append(tmpRunes, runeChar)
			}
		}
		tmpCol[tmpCount] = string(tmpRunes)
		out[i] = tmpCol
	}
	return out
}

func WriteCsv(data [][]string, outputPath string, flagSjis bool) {
	// O_WRONLY:書き込みモード開く, O_CREATE:無かったらファイルを作成
	file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0644)
	FailOnError(err, "csv.go")
	defer file.Close()

	err = file.Truncate(0) // ファイルを空っぽにする(実行2回目以降用)
	FailOnError(err, "csv.go")

	var writer *csv.Writer
	if flagSjis {
		converter, err := iconv.NewWriter(file, "utf-8", "sjis")
		FailOnError(err, "csv.go")
		writer = csv.NewWriter(converter)
	} else {
		writer = csv.NewWriter(file)
	}

	for _, line := range data {
		writer.Write(line)
	}
	writer.Flush()
}
