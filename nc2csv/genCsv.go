package nc2csv

import (
	"fmt"
	"log"
	"strings"
)

// NCプログラムを一字づつ読み込んで処理します。
// 改行でCSVを1行出力して別スレッドに投げます。
func GenCsv(apath *string) *string { // {{{

	initialize(apath) // 初期化処理

	// 別スレッドで受ける
	go receiver(in, out, done)

	lines := strings.Split(*rowInput, "\n") // CSVに入力NCを埋め込むために使う
	pre := '\n'
	flagCommentStart := false

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	in <- "line,NC,X,Y,Z,A,B,C,R,F,moveing distance,vX,vY,vZ,vA,vB,vC,time,cumulative time"

	// 内部でlnを書き換えたいのでこれだとダメ for ln, line := range lines {
	for ln := 0; ln < len(lines); ln++ {
		// 行ごとの処理

		progressbar(int(1000 * float64(setting.CountLF) / float64(len(lines)))) // 進捗を表示
		line := lines[ln] + "\n"                                                // 改行が消えているので付け足す
		rs := []rune(line)                                                      // 一行を読み込んでrsに入れる
		for i := 0; i < len(rs); i++ {                                          // 内部でiを書き換えたいのでこれだとダメ for i, r := range rs {
			// 	1文字づつ処理する
			r := rs[i]
			//if r == '\n' { log.Printf("%s\n", "\\n") } else { log.Printf("%c\n", r) }
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
				if pre == '\n' && i+1 < len(rs) {
					forOptionalSkipBlock(&rs, i)
				} else {
					// 未実装予約語エラー // 除算の/
					log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 演算子 / %v は未実装です。", setting.CountLF, r))
				}
			case '\n':
				if isEOF := forNewLine(&i, &ln, &rs, &lines); isEOF {
					continue
				}
			default:
				if isEOF := forOtherCharactor(&r, &i, &ln, &rs, &lines); isEOF {
					continue
				}
			}
			pre = r
		}
	}

	close(in)       // 別スレッドに終了を通知
	output := <-out // 出力を得る
	<-done          // 別スレッドの終了待ち
	fmt.Println("")
	return &output
} // }}}

// 別スレッドで呼び出される // 受け取ってファイルへ出力ための文字列を作る
func receiver(in chan string, out chan string, done chan bool) { // {{{
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
} // }}}

func initialize(apath *string) { // {{{

	// 適当な初期値
	// これは外部ファイル的なものから読み込むべき
	Assign("F", 1) // 送り速度 初期値

	// 初期設定
	setting.IsMm = true            // mmか
	setting.IS = &ISC              // 最小設定単位の指定 // とりあえずISCとしてみる
	setting.FeedG00 = 99999        // 早送り速度初期値
	setting.FeedG01 = 1            // 送り速度初期値
	setting.IsG90 = true           // アブソリュート指令か
	setting.PlaneDesignation = 17  // 平面指定 とりあえずG17
	setting.CutMode = 0            // 切削モード
	setting.CountLF = 1            // 処理中の行番号
	setting.IsOptionalSkip = false // オプショナルスキップ
	// 座標を持っているメモリへの代入禁止フラグ
	// G01X1とかだとXに1を代入するが、G43X1とかだとXの1は座標値の意味ではない
	// この時G43が来たらそのブロックでは座標値を持っているメモリへの代入を禁止する
	setting.IsProhibitAssignAxis = false
} // }}}
