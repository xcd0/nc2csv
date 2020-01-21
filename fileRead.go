package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"github.com/saintfish/chardet"
	//"golang.org/x/net/html/charset"
)

func readText(path *string) *string {

	// 与えられたパスの文字列について
	// そのパスにあるファイルをテキストファイルとして読み込む

	b, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル%vが読み込めません\n", path)
		log.Fatal(err)
		return nil
	}
	str := string(b)
	// ファイルの文字コード変換
	charset, err := charDet(b)
	if err != nil {
		/*
			fmt.Fprintf(os.Stderr, "文字コード変換に失敗しました\nutf8を使用してください\n")
			log.Println(err)
			panic(err)
			return ""
		*/
		str = convertNewline(&str, "\n")
		return &str
	}

	str, _ = toUtf8(string(b), charset)
	str = convertNewline(&str, "\n")
	return &str
}

func readOptionNumber(next *rune, rs *[]rune, i *int) string {
	if *next == '#' {
		*i++ // 文字をスキップして数値を読み込む
		*i++ // #をスキップして数値を読み込む
		numStr := readAndCutNumber(rs, i)
		*i-- // 最後の数値に合わせる
		return numStr
	}
	if *next == '+' {
		// +は捨てる
		*i++ // 文字をスキップして数値を読み込む
		*i++ // +をスキップして数値を読み込む
		numStr := readAndCutNumber(rs, i)
		// Assign(O, 90) みたいにする
		*i-- // 最後の数値に合わせる
		return numStr
	}
	if *next == '-' {
		// -は残したいので後でつける
		*i++ // 文字をスキップして数値を読み込む
		*i++ // -をスキップして数値を読み込む
		numStr := readAndCutNumber(rs, i)
		// Assign(O, 90) みたいにする
		*i-- // 最後の数値に合わせる
		return "-" + numStr
	}
	if IsLetter((*rs)[*i]) {
		*i++ // 文字をスキップして数値を読み込む
		numStr := readAndCutNumber(rs, i)
		// Assign(O, 90) みたいにする
		*i-- // 最後の数値に合わせる
		return numStr
	}
	// 未実装
	return "log.Fatal(\"" + string((*rs)[*i]) + " はエラーです。\")"
}

/*
func toUtf8(b []byte) []byte {
	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(b)

	// 文字コード変換
	bReader := bytes.NewReader(b)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	u8, _ := ioutil.ReadAll(reader)

	return u8
}
*/

func readAndCutNumber(rs *[]rune, i *int) string {
	numStr := readNumbers(rs, i)
	numRunes := []rune(numStr)
	for i, n := range numStr {
		if n == '0' {
			if len(numStr) == 1 {
				// 00000みたいなのは0を消していったら
				// 全部消えてしまうので1つ残す
				return numStr
			}
			continue
		}
		if len(numStr) > 1 && numRunes[1] == '.' {
			// 0. みたいなの
			continue
		}
		// 先頭の0だけ捨てる
		numStr = numStr[i:]
		numRunes = []rune(numStr)
		break
	}
	return numStr
}

func readNumbers(rs *[]rune, i *int) string {
	pre := *i
	post := *i
	for (IsDigit((*rs)[post]) || IsDot((*rs)[post])) && post < len(*rs)-1 {
		post++
		//log.Printf("len(rs) : %d ; post : %d", len(*rs), post)
	}
	// iを進めておく
	*i = post
	//log.Printf("len(rs) : %d ; last : %c; ret: %s", len(*rs), (*rs)[len(*rs)-1:][0], string((*rs)[pre:]))

	if post == len(*rs)-1 && string((*rs)[len(*rs)-1:]) != "\n" {
		// ファイル終端で改行が含まれていないなら、更に+1する
		*i = post + 1
		if IsDot((*rs)[post]) {
			// 小数点で終わっていたら0を付与する
			return string((*rs)[pre:]) + "0"
		} else {
			return string((*rs)[pre:])
		}
	}

	out := ""
	if IsDot((*rs)[post-1]) {
		// 小数点で終わっていたら0を付与する
		out = string((*rs)[pre:post]) + "0"
	} else {
		out = string((*rs)[pre:post])
	}
	if IsDot([]rune(out)[0]) {
		// .で始まっていたら0をつける
		out = "0" + out
	}
	return out
}

func readLetters(rs *[]rune, i int) string {
	pre := i
	post := i
	for IsLetter((*rs)[post]) {
		post++
	}
	return string((*rs)[pre:post])
}
