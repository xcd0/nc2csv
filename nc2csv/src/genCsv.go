package main

import (
	"fmt"
	"log"
	"strings"

	"./util"
)

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
							// 実装済み予約語 GOTOとか
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
