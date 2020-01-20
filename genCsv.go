package main

import (
	"fmt"
	"log"
	"strings"

	"./util"
)

// 別スレッドで呼び出される
// ファイルへ出力
func srcOutput(in chan string, out chan string, done chan bool) {
	output := ""
	for {
		text, flag := <-in
		if flag {
			output += fmt.Sprintf("%s\n", text)
		} else {
			output += fmt.Sprintf("\n")
			break
		}
	}

	out <- output
	done <- true
	return
}

func genCsv() *string {
	in := make(chan string, 1000) // 別スレッドに投げるバッファ
	out := make(chan string, 2)   // 別スレッドからもらうバッファ
	done := make(chan bool)       // 別スレッドの終了通知をもらうバッファ

	// 別スレッドで受ける
	go srcOutput(in, out, done)

	lines := strings.Split(rowInput, "\n") // CSVに入力NCを埋め込むために使う
	pre := '\n'
	flagCommentStart := false

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	in <- "line,NC,X,Y,Z,A,B,C,R,F,moveing distance,vX,vY,vZ,vA,vB,vC,time,cumulative time"

	// 内部でlnを書き換えたいのでこれだとダメ for ln, line := range lines {
	for ln := 0; ln < len(lines); ln++ {

		// 進捗を表示
		progressbar(int(1000 * float64(setting.CountLF) / float64(len(lines))))

		line := lines[ln] + "\n" // 改行が消えているので付け足す
		rs := []rune(line)

		// 	1文字づつ処理する
		l := len(rs)
		//log.Println(ln)
		// 内部でiを書き換えたいのでこれだとダメ for i, r := range rs {
		for i := 0; i < len(rs); i++ {
			r := rs[i]
			/*
				if r == '\n' {
					log.Printf("%s\n", "\\n")
				} else {
					log.Printf("%c\n", r)
				}
			*/
			if flagCommentStart && !(r == '\n' || r == ')') {
				// コメント中の改行とコメント終了以外の文字はスキップ
				continue
			}
			switch r {
			// ()はコメント 複数行コメントにも対応
			case '%': // 何もしない 本来はプログラムの区切り メモリのリセットをしてもいいかもしれない
			case '(':
				flagCommentStart = true
			case ')':
				flagCommentStart = false
			case '/':
				if pre == '\n' && i+1 < l {
					forOptionalSkipBlock(&rs, i)
				} else {
					// 除算の/
					// 未実装予約語エラー
					log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 演算子 / %v は未実装です。", setting.CountLF, r))
				}
			case '\n':
				forNewLine(&rs, &lines, in)
			default:
				if setting.IsOptionalSkip || setting.IsProhibitAssignAxis {
					// この行は何もしない
					// 改行までどの文字が来ても無視する
				} else {
					// ReadLetters()はアルファベットまたは_がつづく間読み取って返す
					literal := util.ReadLetters(&rs, i)
					if util.IsReserved(literal) {
						// 予約語
						// GOTO IF WHILE THEN の予定？
						// TODO
						if util.IsImplementedWord(literal) {
							// 実装済み予約語
							switch literal {
							case "EOF":
								fmt.Println("")
								log.Printf(fmt.Sprintf("注意 : l.%v : EOF です。正常終了します。", setting.CountLF))
								// ここで終了させる
								i = len(rs)     // for i のやつ
								ln = len(lines) // for lnのやつ
								continue
							default:
								// TODO
							}
						} else if util.IsImplementedCharactor(literal) {
							// 実装済み予約語 %とかGとか
						} else {
							// 未実装予約語
							log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", setting.CountLF, literal))
						}
					}

					// アルファベット+数値を一個ずつデコードする
					if util.IsLetter(r) {
						// X-10.とか G01とか Y#10とか
						forAssignToMemory(&rs, &r, &i, l)
						// LookUpIdentはGOTOやIF,WHILE,ELSE,ENDのような予約語を予約語として
						// それ以外をIDENTIFIERとして返す
					} else if util.IsDot(r) || util.IsDigit(r) {
						// 改行のあとすぐに数値単体で来た時など
						// 数値はアルファベットのあとにしか来ないはず
						// ^10や^.50など
						fmt.Println("")
						log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
					} else if util.IsDot(r) || util.IsDigit(r) {
					} else {
						// 来ないはず
						// 異常値
						log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
						if i+5 > 0 {
							tmp := string(rs[i-6 : i])
							log.Fatal(
								fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 直前の値 %s 文字カウンタ %d",
									setting.CountLF, r, tmp, i))
						}
						log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
					}
				}
			}
			pre = r
		}
	}

	close(in)
	output := <-out
	// 別スレッドの終了待ち
	<-done

	fmt.Println("")
	return &output
}

func forAssignToMemory(rs *[]rune, r *rune, i *int, l int) {
	// ここにはrがアルファベットの場合しか来ない

	// M08 G00 X1. Y2
	// こういうのはすべて変数への代入と見なす
	// var M = 8 みたいなかんじ
	// 但しGだけ特別扱いする
	// NCの場合varに当たるprefixがないので
	// GOTO,IF,WHILEでないアルファベットが来たら変数代入だとみなす
	// 異常値だがAAと来たらAAも変数代入になる？

	if *i+1 >= l {
		// 謎
		// 文字数が超過している。
		log.Fatal(fmt.Sprintf("書式エラー : l.%d : 文字数が超過します。", setting.CountLF, *r))
		log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, *r))
	} else {
		// 文字数チェックOK 次の文字をチェックする
		next := (*rs)[*i+1]
		isDigit := util.IsDigit(next)                                                //数値
		isPMDigit := (util.IsPM(next) && *i+2 <= l && util.IsDigit((*rs)[*i+2]))     // +-付きの数値
		isHashDigit := (util.IsHash(next) && *i+2 <= l && util.IsDigit((*rs)[*i+2])) // #付きの数値
		if isDigit || isPMDigit || isHashDigit {
			// 後ろの数値を読む #はスキップして読み込む
			numStr := readOptionNumber(&next, rs, i) // 中でiが進む
			// 文字 + 数値 or 文字 + 変数
			// G01 とか X#10とか
			if *r == 'G' {
				// Gは特別扱い
				// G専用のQueueに突っ込む
				// G90だったらEnqueueForG(90) みたいにする
				if util.IsHash(next) {
					// まあないはずだけど G#10 みたいに変数使ってきたらという想定
					// EnqueueForG(Hash[90]) みたいにする
					EnqueueForG(Hash(numStr).String())
				} else {
					// G01みたいなの
					// GOTO とかは来ない
					// EnqueueForG(1) みたいにする
					EnqueueForG(numStr)
				}
			} else {
				// G以外のもの M08 とか X1. とか Y#10 とか

				// G90G00X100.とかでは、X100.の時点でGのキューに要素がある
				// G以外の代入が走る前にGを処理する
				FlushGqueue()
				if util.IsHash(next) {
					Assign(string(*r), Hash(numStr).String()) // X#100とか Assign(P, Hash[412]) みたいにする
				} else {
					// Assign(O, 90)とかAssign(X, "10")とかAssign(X, "10.0")とかみたいにする
					Assign(string(*r), numStr)
				}
			}
			*r = (*rs)[*i] // readOptionNumber中でiが数値分スキップしているのでpreの保存用に更新しておく
		}
	}
}
