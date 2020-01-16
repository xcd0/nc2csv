package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	//"github.com/nsf/termbox-go"
)

func progressbar(progressVal int) {
	flag.Parse()

	if progressVal < 0 {
		fmt.Println("Error!")
		os.Exit(1)
	} else if progressVal > 1000 {
		progressVal = 1000
	}

	/*
		// ウィンドウサイズ取得
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		w, _ := termbox.Size()
		termbox.Close()
	*/
	w := 80

	// 進捗バーは画面の半分の長さとする
	width := int(w / 2)

	f := float64(progressVal) / 1000.0
	printnum := int(float64(width) * f)
	fp := f * 100.0

	// 出力
	bar := "  ( "
	if progressVal < 100.0 {
		bar += "  "
	} else if progressVal < 1000.0 {
		bar += " "
	}
	bar += strconv.FormatFloat(fp, 'f', 1, 64)
	bar += " % ) ["

	for i := 0; i < width; i++ {
		if i < printnum {
			bar += "#"
		} else if i == printnum {
			bar += ">"
		} else {
			bar += " "
		}
	}
	bar += "]"

	fmt.Printf("\r%v", bar)
}
