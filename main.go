package main

import (
	"flag"
	"log"
	"path/filepath"
)

func main() {

	// ログ設定
	log.SetFlags(log.Llongfile)

	flag.Parse()
	// 引数
	if flag.NArg() == 0 {
		log.Fatal("エラー : 引数が与えられていません。")
	}
	apath, _ := filepath.Abs(flag.Arg(0))

	// 処理の本体
	csv := genCsv(&apath)

	// ファイルに書き出す
	writeCsv(apath, csv)

}
