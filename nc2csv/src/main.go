package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"./util"
)

var outputFile *os.File
var rowInput string
var clearInput string

func main() {

	// ログ設定
	log.SetFlags(log.Llongfile)

	flag.Parse()
	// 引数
	apath, _ := filepath.Abs(flag.Arg(0))
	//apath, _ := filepath.Abs("./test/nc10")

	// 入力ファイルを開く
	rowInput = util.ReadText(apath) // NCを読み込んでstringに変換、改行コードを統一
	//clearInput = util.DeleteComment(rowInput) // コメントを削除

	// 初期化処理
	Initialize(&rowInput)

	// 処理の本体
	csv := genCsv()

	// ファイルに書き出す
	writeCsv(apath, csv)

}

func writeCsv(apath string, csv *string) {
	// 出力ファイルを開く
	outputDir := filepath.Dir(apath)
	outputName := filepath.Base(apath) + ".csv"
	// 入力ファイルと同じディレクトリに入力ファイル+.csvのファイルを作成する。
	outputFilePath := filepath.Join(outputDir, outputName)

	var err error
	outputFile, err = os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		// Openエラー処理
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputFile.Write(([]byte)(*csv))
}
