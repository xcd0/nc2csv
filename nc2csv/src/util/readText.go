package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xcd0/go-nkf"
)

func ReadText(path string) string {

	// 与えられたパスの文字列について
	// そのパスにあるファイルをテキストファイルとして読み込む

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル%vが読み込めません\n", path)
		log.Fatal(err)
		return ""
	}
	// ファイルの文字コード変換
	charset, err := nkf.CharDet(b)
	if err != nil {
		/*
			fmt.Fprintf(os.Stderr, "文字コード変換に失敗しました\nutf8を使用してください\n")
			log.Println(err)
			panic(err)
			return ""
		*/
		return ConvertNewline(string(b), "\n")
	}

	str, err := nkf.ToUtf8(string(b), charset)

	str = ConvertNewline(str, "\n")

	return str
}
