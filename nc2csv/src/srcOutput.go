package main

import (
	"fmt"
)

// 別スレッドで呼び出される
// ファイルへ出力
func srcOutput(in chan string, out chan string, done chan bool) {
	output := ""
	for {
		text, flag := <-in
		if flag {
			output += fmt.Sprintf("%s\n", text)
		} else {
			output += fmt.Sprintf("\n")
			break
		}
	}
	// ファイルに追記する
	fmt.Fprintf(outputFile, output)

	out <- output
	done <- true
	return
}
