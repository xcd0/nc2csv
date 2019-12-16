package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	apath, _ := filepath.Abs(arg)
	rowInput := ReadText(apath)
	spacedInput := insertSpace(rowInput)

	fmt.Printf("%v\n", rowInput)
	fmt.Printf("-----------\n")
	fmt.Printf("%v\n", spacedInput)
	fmt.Printf("-----------\n")

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

var regAZ = regexp.MustCompile(`.[A-Z]`)

func insertSpace(input string) string { // {{{
	//アルファベットの前に半角空白を入れる
	// 行頭は入れない
	output := ""

	flagCommentStart := false

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		// 行内の文字を1文字づつチェック
		for j, runeChar := range line {

			// コメントの中身はそのまま出力する
			// ()はコメント 複数行コメントにも対応
			if runeChar == '(' {
				flagCommentStart = true
				// ただし(の前には入れる
				if j != 0 {
					output += " "
				}
			}
			if flagCommentStart {
				if runeChar == ')' {
					flagCommentStart = false
				}
				// 空白を入れない
				output += string(runeChar)
				continue
			}

			if j != 0 && // 行頭でない
				'A' <= runeChar && // アルファベット大文字
				'Z' >= runeChar {
				output += " "
			}
			output += string(runeChar)
		}
		output += "\n" // 改行なくなってるので追加
	}
	return output
}

// }}}

func ReadText(path string) string { // {{{
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
} // }}}

func ConvertNewline(str, nlcode string) string { // {{{
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}

// }}}
