/*
NCプログラムを読み込み、分析して、各軸座標位置、各行の送り速度、移動時間 などをある程度見やすく出力する。

in : nc プログラム (テキストファイル、拡張子不定)

out : 1行ごとの各軸の動作、送り速度の解析結果 (csv)
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xcd0/nc2csv/nc2csv"
)

func main() {

	// ログ設定
	log.SetFlags(log.Llongfile)

	flag.Parse()
	// 引数
	if flag.NArg() == 0 {
		fmt.Println("エラー : 引数が与えられていません。")
		fmt.Println("　　　 : 引数にはNCプログラムの書かれたテキストファイルを与えます。")
		fmt.Println("　　　 : また-hでこのプログラムのソースコードを出力するためのオプションが表示されます。")
		fmt.Println("　　　 : このバイナリにはソースコードが埋め込まれています。")
		fmt.Println("　　　 : 終了します。")
		os.Exit(1)
	}
	apath, _ := filepath.Abs(flag.Arg(0))

	// NCプログラムを読み込む 戻り値は読み込んだ文字列へのポインタ
	_ = nc2csv.ReadNcProgram(&apath)
	// 処理の本体
	csv := nc2csv.GenCsv(&apath)

	// ファイルに書き出す
	nc2csv.WriteCsv(apath, csv)

}
