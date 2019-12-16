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

	for {
		tok := l.NextToken()
		if tok.Type == "EOF" {
			break
		}

		fmt.Printf("{%v, \"%v\"},\n", tok.Type, tok.Literal)
	}

	return

	/*
		var sc scanner.Scanner
		src := []byte(str)
		errorHandler := func(pos token.Position, msg string) { fmt.Printf("ERROR %v %v\n", pos, msg) }
		sc.Init(token.NewFileSet().AddFile("", -1, len(src)), src, errorHandler, 0)
		fmt.Printf("%6v %6v %6v\n", "pos", "tok", "lit")
		for {
			pos, tok, lit := sc.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%6v %6v %6v\n", pos, tok, lit)
		}
	*/
}
