package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	apath, _ := filepath.Abs(arg)
	str := ReadText(apath)

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
}

func ReadText(path string) string {

	// 与えられたパスの文字列について
	// そのパスにあるファイルをテキストファイルとして読み込む

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル%vが読み込めません\n", path)
		log.Println(err)
		panic(err)
		return ""
	}

	str := ConvertNewline(string(b), "\n")

	return str
}

func ConvertNewline(str, nlcode string) string {
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}
