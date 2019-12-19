package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"./lexer"    // 字句解析器
	_ "./parser" // 構文解析器
)

func main() {
	flag.Parse()
	apath, _ := filepath.Abs(flag.Arg(0))

	// 前処理
	rowInput := ReadText(apath)          // 読み込んでstringに
	spacedInput := insertSpace(rowInput) // 空白を入れる 不要かも
	input := DeleteComment(spacedInput)  // コメント()を削除

	l := lexer.NewLexer(input)
	ts := make([]lexer.Token, 0, 1000)
	// トークン毎に出力する
	for {
		tok := l.NextToken()
		ts = append(ts, tok)
		if tok.Type == "EOF" {
			break
		}
	}
	/*
		for _, t := range ts {
			if t.Type == "EOB" || t.Type == "EOF" {
				fmt.Printf("{%v, \"%v\"}\n", t.Type, t.Literal)
			} else {
				fmt.Printf("{%v, \"%v\"},", t.Type, t.Literal)
			}
		}
	*/

	// 保存テスト
	fpath := "./save.txt"
	save(fpath, ts)
	var tsload []lexer.Token
	load(fpath, &tsload)
	fmt.Println(tsload)

	return

}
