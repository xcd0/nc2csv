package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"./util"
)

func genCsv(input string) string {
	in := make(chan string, 1000) // 別スレッドに投げるバッファ
	out := make(chan string, 2)   // 別スレッドからもらうバッファ
	done := make(chan bool)       // 別スレッドの終了通知をもらうバッファ

	// 別スレッドで受ける
	go srcOutput(in, out, done)

	rs := []rune(input)

	lines := strings.Split(input, "\n")

	l := len(rs)
	pre := '\n'

	// 適当な初期値
	Assign("F", 1000) // 送り速度 初期値

	Setting.CountLF = 1
	Setting.IsOptionalSkip = false // オプショナルスキップ
	// 座標を持っているメモリへの代入禁止フラグ
	// G01X1とかだとXに1を代入する
	// G43X1とかだとXの1は座標値の意味ではない
	// この時G43が来たらそのブロックでは座標値を持っているメモリへの代入を禁止する
	Setting.IsProhibitAssignAxis = false

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	in <- "line,NC,X,Y,Z,A,B,C,R,F,vX,vY,vZ,vA,vB,vC,time,cumulative time"

	// 行ごとにcase文を出力する
	for i := 0; i < l; i++ {
		r := rs[i]
		/*
			if r == '\n' {
				log.Printf("%s\n", "\\n")
			} else {
				log.Printf("%c\n", r)
			}
		*/
		switch r {
		case '/': // {{{
			if pre == '\n' && i+1 < l {
				// オプショナルスキップブロック
				// オプショナルスキップはメモリの値を読んで
				// trueなら無視する。

				if t := rs[i+1]; '1' <= t && t <= '9' {
					// 番号付きオプショナルスキップブロック
					tmpNum, _ := strconv.Atoi(string(r))
					if OptionalSkip[tmpNum] == false {
						// 無視しない
						Setting.IsOptionalSkip = false
					} else {
						// 無視する
						Setting.IsOptionalSkip = true
					}
				} else {
					if OptionalSkip[0] {
						// 無視しない
						Setting.IsOptionalSkip = false
					} else {
						// 無視する
						Setting.IsOptionalSkip = true
					}
				}
			} else {
				// 除算の/
				// 未実装予約語エラー
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 演算子 / %v は未実装です。", Setting.CountLF, r))
			}
		// }}}
		case '\n': // {{{

			// 進捗を表示
			fmt.Printf("\rprogress : % 3.2f %%  ", 100*float64(Setting.CountLF)/float64(len(lines)))

			// フラグをリセットする
			Setting.IsOptionalSkip = false
			Setting.IsProhibitAssignAxis = false
			// Gのキューを実行する
			FlushGqueue()

			// この行を実行した後の状態を出力する
			outputOneLine, time := axis.genOnelineCsv()

			Setting.CumulativeTime += time
			in <- outputOneLine

			// Setting.CountLFは1からだけどlinesは0から
			//log.Printf("l.%v : %v\n", Setting.CountLF-1, lines[Setting.CountLF-1])

			Setting.CountLF++
			if string(rs[len(rs)-1:]) == "\n" && Setting.CountLF == len(lines) {
				// ファイル最後に改行があるファイルとないファイルに対応する
				continue
			}

			// これ要るのだろうか...
			if pre == '\n' {
				// 空行
				continue
			}
			// }}}
		case '%':
			// 何もしない
			// 本来はプログラムの区切り
			// メモリのリセットをしてもいいかもしれない
		default:
			if Setting.IsOptionalSkip || Setting.IsProhibitAssignAxis {
				// この行は何もしない
				// 改行までどの文字が来ても無視する
			} else {
				// アルファベット+数値を一個ずつデコードする
				if util.IsLetter(r) {
					// M08 G00 X1. Y2
					// こういうのはすべて変数への代入と見なす
					// var M = 8 みたいなかんじ
					// 但しGだけ特別扱いする
					// NCの場合varに当たるprefixがないので
					// GOTO,IF,WHILEでないアルファベットが来たら変数代入だとみなす
					// 異常値だがAAと来たらAAも変数代入になる？

					// readIdentifier()はアルファベットまたは_がつづく間読み取って返す
					literal := util.ReadLetters(&rs, i)
					if util.IsReserved(literal) {
						// 予約語
						// GOTO IF WHILE THEN の予定？
						// TODO
						if util.IsImplemented(literal) {
							// 実装済み予約語
						} else {
							// 未実装予約語
							log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", Setting.CountLF, literal))
						}
					} else if i+1 >= l {
						// 謎
						// 文字数が超過している。
						log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", Setting.CountLF, r))
					} else {
						// 文字数チェックOK
						if next := rs[i+1]; '0' <= next && next <= '9' ||
							next == '-' || next == '+' || next == '#' {
							// 文字 + 数値 or 文字 + 変数
							// G01 とか X#10とか
							if r == 'G' {
								// Gは特別扱い
								// G専用のQueueに突っ込む
								// G90だったらEnqueueForG(90) みたいにする
								if next == '#' {
									// まあないはずだけど G#10 みたいに変数使ってきたらという想定
									numStr := readOptionNumber(&next, &rs, &i)
									// EnqueueForG(Hash[90]) みたいにする
									tmpNum, _ := strconv.Atoi(numStr)
									EnqueueForG(Hash(tmpNum))
									r = rs[i] // 中でiが数値分スキップしているのでpreの保存用に更新しておく
								} else {
									// G01みたいなの
									// GOTO とかは来ない
									numStr := readOptionNumber(&next, &rs, &i)
									// EnqueueForG(1) みたいにする
									tmpNum, _ := strconv.Atoi(numStr)
									EnqueueForG(tmpNum)
								}
							} else {
								// G以外のもの M08 とか X1. とか Y#10 とか

								// G90G00X100.とかでは、X100.の時点でGのキューに要素がある
								// G以外の代入が走る前にGを処理する
								FlushGqueue()
								if next == '#' {
									// X#100とか
									// この関数は#100とかの場合でも#を無視して"100"を返してくれる
									numStr := readOptionNumber(&next, &rs, &i)
									// Assign(P, Hash[412]) みたいにする
									Assign(string(r), Hash(numStr)) // 中でiは進んでいる
									r = rs[i]                       // 中でiが数値分スキップしているのでpreの保存用に更新しておく
								} else if util.IsLetter(r) {
									// X1.とか X-10.とか
									numStr := readOptionNumber(&next, &rs, &i)
									// Assign(O, 90) みたいにする
									// Assign(X, "10") Assign(X, "10.0") とかも可
									Assign(string(r), numStr) // 中でiは進んでいる
									r = rs[i]                 // 中でiが数値分スキップしているのでpreの保存用に更新しておく
								} else {
									// 謎
									// 来ないはず
									// 異常値
									// X! とか X(とか
									// X-とかは来ない
									tmp := string(rs[i-3 : i+3])
									log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 %s", Setting.CountLF, r, tmp))
								}
							}
						}
						//
					}
					// LookUpIdentはGOTOやIF,WHILE,ELSE,ENDのような予約語を予約語として
					// それ以外をIDENTIFIERとして返す
				} else if util.IsDot(r) {
					// 来ないはず
					// 小数点 // 確定で浮動小数点
					log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", Setting.CountLF, r))
				} else if util.IsDigit(r) {
					// 来ないはず
					// 数値 // 小数点があるかどうかわからないと確定しない
					log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", Setting.CountLF, r))
				} else {
					// 来ないはず
					// 異常値
					if i+5 > 0 {
						tmp := string(rs[i-6 : i])
						log.Fatal(
							fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 直前の値 %s 文字カウンタ %d",
								Setting.CountLF, r, tmp, i))
					}
					log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", Setting.CountLF, r))
				}
			}
		}
		pre = r
	}

	fmt.Printf("\rprogress : 100.00 %%\n")
	close(in)
	output := <-out
	// 別スレッドの終了待ち
	<-done

	// 変数代入の変換
	//convertedAssign := convertAssign(input)
	return output
}
