package main

import (
	"fmt"
	"log"
	"strings"

	"./util"
)

var lines []string

func CreateSrc(rowInput string) string {

	// コメントを削除
	clearInput := util.DeleteComment(rowInput)

	log.Printf("clearInput : %v", clearInput)

	lines = strings.Split(rowInput, "\n")

	runfunc := makeRunFunction(clearInput)

	log.Printf("runfunc : %v", runfunc)

	return runfunc
}

func makeRunFunction(input string) string {
	in := make(chan string, 100) // 別スレッドに投げるバッファ
	out := make(chan string, 2)  // 別スレッドからもらうバッファ
	done := make(chan bool)      // 別スレッドの終了通知をもらうバッファ

	// 別スレッドで受ける
	go srcOutput(in, out, done)

	rs := []rune(input)
	countLF := 1
	l := len(rs)
	pre := '\n'

	bOptionalSkip := false
	// 行ごとにcase文を出力する
	for i := 0; i < l; i++ {
		r := rs[i]
		log.Printf("%c\n", r)
		switch r {
		case '/': // {{{
			if pre == '\n' {
				if i+1 < l {
					// オプショナルスキップブロック
					// ただオプショナルスキップブロック用のif文を出力し、
					// 改行で閉じるだけ。
					bOptionalSkip = true
					if t := rs[i+1]; '1' <= t && t <= '9' {
						// 番号付きオプショナルスキップブロック
						in <- "if !OptionalSkip[" + string(r) + "] {"
					} else {
						in <- "if !OptionalSkip[0] {"
					}
				} else {
					// 文法エラーだけどとりあえずNOPにする
					in <- "NOP = true"
				}
			}
			// 除算の/
			// 未実装予約語エラー
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 演算子 / %v は未実装です。", countLF, r))
		// }}}
		case '\n': // {{{

			// countLFは1からだけどlinesは0から
			log.Printf("l.%v : %v\n", countLF-1, lines[countLF-1])

			countLF++
			if string(rs[len(rs)-1:]) == "\n" && countLF == len(lines) {
				// ファイル最後に改行があるファイルとないファイルに対応する
				continue
			}

			if pre == '\n' {
				// 空行
				in <- "NOP = true"
				in <- "break"
				continue
			}
			if bOptionalSkip {
				bOptionalSkip = false
				in <- "}"
			}
			in <- "break"
			// }}}
		case '%':
			in <- "NOP = true"
		default:
			// アルファベット+数値を一個ずつデコードする
			if util.IsLetter(r) {
				// M08 G00 X1.
				// こういうのはすべて変数への代入と見なす
				// var M = 8 みたいなかんじ
				// golang の場合変数宣言と代入は var <identifier> = <expression> になる
				// NCの場合varに当たるprefixがないので
				// GOTO,IF,WHILEでないアルファベットが来たら変数代入だとみなす
				// 異常値だがAAと来たらAAも変数代入になる？

				// readIdentifier()はアルファベットまたは_がつづく間読み取って返す
				literal := util.ReadLetters(&rs, i)
				if util.IsReserved(literal) {
					// 予約語
					//out <- "Assign(G, 49)"
					if util.IsImplemented(literal) {
						// 実装済み予約語
						// 上でcatchされていないとおかしい
					} else {
						// 未実装予約語
						log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", countLF, literal))
					}
				} else if i+1 < l {
					if next := rs[i+1]; '0' <= next && next <= '9' ||
						next == '-' ||
						next == '+' ||
						next == '#' {
						if r == 'G' {
							// Gは特別扱い
							// G専用のQueueに突っ込む
							// G90だったらEnqueueForG(90) みたいにする
							if next == '#' {
								numStr := getOptionNumber(&next, &rs, &i)
								// Assign(P, Hash[412]) みたいにする
								t := "EnqueueForG(Hash[" + numStr + "])"
								in <- t
								r = rs[i] // 中でiが数値分スキップしているのでpreの保存用に更新しておく
							}
							if util.IsLetter(r) {
								numStr := getOptionNumber(&next, &rs, &i)
								// Assign(O, 90) みたいにする
								t := "EnqueueForG(" + numStr + ")"
								in <- t
								r = rs[i] // 中でiが数値分スキップしているのでpreの保存用に更新しておく
							} else {
								// #[1] みたいなのは未実装
								log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", countLF, r))
							}
						} else {
							if next == '#' {
								numStr := getOptionNumber(&next, &rs, &i)
								// Assign(P, Hash[412]) みたいにする
								t := "Assign(" + string(r) + ", Hash[" + numStr + "])" // 中でiは進んでいる
								in <- t
								r = rs[i] // 中でiが数値分スキップしているのでpreの保存用に更新しておく
							} else if util.IsLetter(r) {
								numStr := getOptionNumber(&next, &rs, &i)
								// Assign(O, 90) みたいにする
								t := "Assign(" + string(r) + ", " + numStr + ")" // 中でiは進んでいる
								in <- t
								r = rs[i] // 中でiが数値分スキップしているのでpreの保存用に更新しておく
							} else {
								// 謎
								tmp := string(rs[i-3 : i+3])
								log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 %s", countLF, r, tmp))
							}
						}

					}
				} else {
					// 謎
					log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", countLF, r))
				}

				// LookUpIdentはGOTOやIF,WHILE,ELSE,ENDのような予約語を予約語として
				// それ以外をIDENTIFIERとして返す
			} else if util.IsDot(r) {
				// 来ないはず
				// 小数点 // 確定で浮動小数点
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", countLF, r))
			} else if util.IsDigit(r) {
				// 来ないはず
				// 数値 // 小数点があるかどうかわからないと確定しない
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", countLF, r))
			} else {
				// 来ないはず
				// 異常値
				if i+5 > 0 {
					tmp := string(rs[i-6 : i])
					log.Fatal(
						fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 直前の値 %s 文字カウンタ %d",
							countLF, r, tmp, i))
				}
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", countLF, r))
			}
		}

		pre = r
	}
	close(in)
	output := <-out
	// 別スレッドの終了待ち
	<-done

	// 変数代入の変換
	//convertedAssign := convertAssign(input)
	return output
}

func readAndCutNumber(rs *[]rune, i *int) string {
	numStr := util.ReadNumbers(rs, i)
	for i, n := range numStr {
		if n == '0' {
			if len(numStr) == 1 {
				// 00000みたいなのは0を消していったら
				// 全部消えてしまうので1つ残す
				return numStr
			}
			continue
		}
		// 先頭の0だけ捨てる
		numStr = numStr[i:]
		break
	}
	return numStr
}

func getOptionNumber(next *rune, rs *[]rune, i *int) string {
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
	if util.IsLetter((*rs)[*i]) {
		*i++ // 文字をスキップして数値を読み込む
		numStr := readAndCutNumber(rs, i)
		// Assign(O, 90) みたいにする
		*i-- // 最後の数値に合わせる
		return numStr
	}
	// 未実装
	return "log.Fatal(\"" + string((*rs)[*i]) + " はエラーです。\")"
}
