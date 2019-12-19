package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"./lexer"
)

func main() {
	flag.Parse()
	apath, _ := filepath.Abs(flag.Arg(0))

	// 前処理
	rowInput := ReadText(apath)          // 読み込んでstringに
	spacedInput := insertSpace(rowInput) // 空白を入れる 不要かも
	input := DeleteComment(spacedInput)  // コメント()を削除

	l := lexer.NewLexer(input)
	// トークン毎に出力する
	for {
		tok := l.NextToken()
		if tok.Type == "EOF" {
			break
		}
		if tok.Type != "EOB" {
			fmt.Printf("{%v, \"%v\"},", tok.Type, tok.Literal)
		} else {
			fmt.Printf("{%v, \"%v\"}\n", tok.Type, tok.Literal)
		}

	}

	return

}
