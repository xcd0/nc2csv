package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	apath, _ := filepath.Abs(arg)

	rowInput := ReadText(apath)
	spacedInput := insertSpace(rowInput)

	//fmt.Println(spacedInput)

	l := New(spacedInput)

	// トークン毎に出力する
	for {
		tok := l.NextToken()
		if tok.Type == "EOF" {
			break
		}

		fmt.Printf("{%v, \"%v\"},\n", tok.Type, tok.Literal)
	}

	return

}
