package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	/*
		// 第一引数にファイルのパスを受け取る
		switch flag.NArg() {
		case 1:
			// OK
		default:
			log.Fatal("引数を1つ指定してください。\n")
			return
		}
		arg := flag.Arg(0)
	*/
	//arg := "..\\A1.ncd"
	arg := "..\\A1_utf8_lf.ncd"

	// 引数をいい感じに整形する
	ii := Argparse(arg)

	// ファイルを読み込む
	inputString := ReadText(ii.Apath)

	// 分析する
	output := AnalizeString(inputString)

	// まだ実装してないので適当に出力する
	fmt.Printf("%v\n", output)
}
