package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func writeCsv(apath string, csv *string) {
	// 出力ファイルを開く
	outputDir := filepath.Dir(apath)
	outputName := filepath.Base(apath) + ".csv"
	// 入力ファイルと同じディレクトリに入力ファイル+.csvのファイルを作成する。
	outputFilePath := filepath.Join(outputDir, outputName)
	var outputFile *os.File
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
