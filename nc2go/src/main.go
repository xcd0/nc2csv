package main

import (
	"fmt"
	"log"
	"path/filepath"

	"./util"
)

func main() {

	log.SetFlags(log.Llongfile)

	/*
		flag.Parse()
		// 引数
		apath, _ := filepath.Abs(flag.Arg(0))
	*/
	apath, _ := filepath.Abs("./test/nc10")

	// NCを読み込んでstringに変換、改行コードを統一
	rowInput := util.ReadText(apath)

	fmt.Println("read input file : " + rowInput)

	src := CreateSrc(rowInput)

	fmt.Println(src)

	/*
		// 保存テスト
		fpath := "./save.txt"
		util.Save(fpath, ts)
		var tsload []token.Token
		util.Load(fpath, &tsload)
		fmt.Println(tsload)
	*/

}
