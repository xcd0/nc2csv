package main

import (
	"flag"
	"fmt"
	"go/token"
	"path/filepath"
)

func main() {

	flag.Parse()
	// 引数
	apath, _ := filepath.Abs(arg)

	// NCを読み込んでstringに変換、改行コードを統一
	rowInput := ReadText(apath)

	src := createSrc(rowInput)

	// 保存テスト
	fpath := "./save.txt"
	util.Save(fpath, ts)
	var tsload []token.Token
	util.Load(fpath, &tsload)
	fmt.Println(tsload)

	return

}
