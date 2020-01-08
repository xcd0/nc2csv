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
	//apath, _ := filepath.Abs(flag.Arg(0))
	apath, _ := filepath.Abs("./test/nc10")

	// NCを読み込んでstringに変換、改行コードを統一
	rowInput := util.ReadText(apath)

	clearInput := util.DeleteComment(rowInput) // コメントを削除

	Initialize(&rowInput) // 初期化処理
	csv := MakeCsv(clearInput)

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
