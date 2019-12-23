package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"./lexer"    // 字句解析器
	_ "./parser" // 構文解析器
	"./token"
	"./util"
)

func main() {
	flag.Parse()
	apath, _ := filepath.Abs(flag.Arg(0))

	// 前処理
	rowInput := util.ReadText(apath)          // 読み込んでstringに
	spacedInput := util.InsertSpace(rowInput) // 空白を入れる 不要かも
	input := util.DeleteComment(spacedInput)  // コメント()を削除

	l := lexer.NewLexer(input)
	ts := make([]token.Token, 0, 1000)
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
	util.Save(fpath, ts)
	var tsload []token.Token
	util.Load(fpath, &tsload)
	fmt.Println(tsload)

	return

}
