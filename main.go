package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

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
	if flag.NArg() == 0 {
		log.Fatal("エラー : 引数が与えられていません。")
	}
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
	//outputFile, err = os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	outputFile, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("")
		fmt.Println("---------------")
		fmt.Println("")
		fmt.Println(*csv)
		fmt.Println("")
		fmt.Println("---------------")
		fmt.Println("")
		// Openエラー処理
		log.Println("エラー : 出力先ファイルが開けません。")
		log.Println("       : 他のプログラムでファイルを開いていませんか？")

		t := time.Now().Local()
		outputName = filepath.Base(apath) + "_" + fmt.Sprintf(t.Format("2006-01-02-15-04-05")) + ".csv"
		outputFilePath = filepath.Join(outputDir, outputName)
		outputFile, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(fmt.Sprintf("エラーメッセージ : %v\nエラー終了します。\n", err))
		}
		log.Println("       : 別ファイル名で保存します。")
	}
	defer outputFile.Close()

	outputFile.Write(([]byte)(*csv))

	fmt.Printf("\n\n")
	fmt.Println("NCデータからの生成を完了しました。")
	fmt.Printf("input  : %v\n", apath)
	fmt.Printf("output : %v\n", outputFilePath)
}
