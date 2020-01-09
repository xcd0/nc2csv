package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"./util"
)

func main() {

	log.SetFlags(log.Llongfile)

	flag.Parse()
	// 引数
	apath, _ := filepath.Abs(flag.Arg(0))

	// NCを読み込んでstringに変換、改行コードを統一
	rowInput := util.ReadText(apath)

	clearInput := util.DeleteComment(rowInput) // コメントを削除

	// 初期化処理
	Initialize(&rowInput)

	// 処理の本体
	csv := genCsv(clearInput)

	fmt.Println(csv)

	/*
		// 保存テスト
		fpath := "./save.txt"
		util.Save(fpath, ts)
		var tsload []token.Token
		util.Load(fpath, &tsload)
		fmt.Println(tsload)
	*/

}
