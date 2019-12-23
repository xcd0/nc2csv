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
	ext := filepath.Ext(apath)
	rowInput := ReadText(apath)

	if ext == ".csv" {
		out := ReadCsv(rowInput)
		fmt.Println(out)
	} else {
		out := ReadXml(rowInput)
		fmt.Println(out)
	}
}
