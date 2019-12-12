package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	// 第一引数にファイルのパスを受け取る
	switch flag.NArg() {
	case 1:
		// OK
	default:
		log.Fatal("引数を1つ指定してください。\n")
		return
	}

	// 引数をいい感じに整形する
	ii := Argparse(flag.Arg(0))

	// ファイルを読み込む
	inputString := ReadText(ii.Apath)

	// 分析する
	output := AnalizeString(inputString)

	// まだ実装してないので適当に出力する
	fmt.Printf("%v\n", output)
}
