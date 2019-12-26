package main

import (
	"fmt"
	"log"
)

func srcOutput(in chan string, out chan string, done chan bool) {
	output := `func (p *program) run(pc int) {
	if pc == 0 || pc > p.length {
		e := fmt.Sprintf("実行エラー : l.%d : その行は存在しません。エラーです。", pc)
		log.Fatal(e)
	}

	NOP := false // コメント行 runBlock()を呼ばない
	switch pc {
	case 1:
`
	count := 1
	for {
		text, flag := <-in
		if flag {
			if text == "break" {
				count++
				output += fmt.Sprintf("\tcase %v:\n", count)
				log.Println(output)
			} else {
				output += fmt.Sprintf("\t\t%s\n", text)
			}
		} else {
			output += fmt.Sprintf("\t}\n\tif !NOP {\n\t\trunBlock()\n\t}\n}\n")
			break
		}
	}
	log.Println(output)
	out <- output
	done <- true
	return
}
