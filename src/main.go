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
	input := DeleteComment(spacedInput)

	l := New(input)

	// トークン毎に出力する
	for {
		tok := l.NextToken()
		if tok.Type == "EOF" {
			break
		}
		if tok.Type != "EOB" {
			fmt.Printf("{%v, \"%v\"},", tok.Type, tok.Literal)
		} else {
			fmt.Printf("{%v, \"%v\"}\n", tok.Type, tok.Literal)
		}

	}

	return

}
