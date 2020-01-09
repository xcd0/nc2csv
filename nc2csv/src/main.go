package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"./util"
)

var outputFile *os.File

func main() {

	log.SetFlags(log.Llongfile)

	flag.Parse()
	// 引数
	apath, _ := filepath.Abs(flag.Arg(0))
	//apath, _ := filepath.Abs("./test/nc10")

	// 保存先ファイルを開く
	outputDir := filepath.Dir(apath)
	outputName := filepath.Base(apath) + ".csv"
	outputFilePath := filepath.Join(outputDir, outputName)
	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		// Openエラー処理
		log.Fatal(err)
	}
	defer outputFile.Close()

	// NCを読み込んでstringに変換、改行コードを統一
	rowInput := util.ReadText(apath)

	clearInput := util.DeleteComment(rowInput) // コメントを削除

	// 初期化処理
	Initialize(&rowInput)

	// 処理の本体
	csv := genCsv(clearInput)

	//fmt.Println(csv)

	outputFile.Write(([]byte)(csv))
}
