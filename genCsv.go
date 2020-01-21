package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var in = make(chan string, 1000) // 別スレッドに投げるバッファ
var out = make(chan string, 2)   // 別スレッドからもらうバッファ
var done = make(chan bool)       // 別スレッドの終了通知をもらうバッファ

func genCsv(apath *string) *string { // {{{

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
				if isEOF := forLF(&i, &ln, &rs, &lines); isEOF {
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

func forLF(i *int, ln *int, rs *[]rune, lines *[]string) bool { // {{{
	// 戻り値はtrueの時continueする
	// フラグをリセットする
	setting.IsOptionalSkip, setting.IsProhibitAssignAxis = false, false
	flushGqueue()                         // Gのキューを実行する Gは先にすべて実行する
	outputOneLine := axis.genOnelineCsv() // この行を実行した後の状態を出力する
	in <- outputOneLine
	setting.CountLF++
	if string((*rs)[len(*rs)-1:]) == "\n" && setting.CountLF == len(*lines) {
		// ファイル最後に改行があるファイルとないファイルに対応する

		// ここで終了させる
		*i = len(*rs)     // for i のやつ
		*ln = len(*lines) // for lnのやつ
		return true
	}
	return false
} // }}}

func forOtherCharactor(r *rune, i *int, ln *int, rs *[]rune, lines *[]string) bool { // {{{
	// 戻り値はtrueの時continueする
	if setting.IsOptionalSkip || setting.IsProhibitAssignAxis {
		// この行は何もしない
		// 改行までどの文字が来ても無視する
	} else {
		// readLetters()はアルファベットまたは_がつづく間読み取って返す
		literal := readLetters(rs, *i)
		if IsReserved(literal) {
			// 予約語
			// GOTO IF WHILE THEN の予定？
			// TODO
			if IsImplementedWord(literal) {
				// 実装済み予約語
				switch literal {
				case "EOF":
					fmt.Println("")
					log.Printf(fmt.Sprintf("注意 : l.%v : EOF です。正常終了します。", setting.CountLF))
					// ここで終了させる
					*i = len(*rs)     // for i のやつ
					*ln = len(*lines) // for lnのやつ
					return true
				default:
					// TODO
				}
			} else if IsImplementedCharactor(literal) {
				// 実装済み予約語 %とかGとか
			} else {
				// 未実装予約語
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", setting.CountLF, literal))
			}
		}

		// アルファベット+数値を一個ずつデコードする
		if IsLetter(*r) {
			// X-10.とか G01とか Y#10とか
			forAssignToMemory(rs, r, i, len(*rs))
			// LookUpIdentはGOTOやIF,WHILE,ELSE,ENDのような予約語を予約語として
			// それ以外をIDENTIFIERとして返す
		} else if IsDot(*r) || IsDigit(*r) {
			// 改行のあとすぐに数値単体で来た時など
			// 数値はアルファベットのあとにしか来ないはず
			// ^10や^.50など
			fmt.Println("")
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
		} else {
			// 来ないはず
			// 異常値
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
			if *i+5 > 0 {
				tmp := string((*rs)[*i-6 : *i])
				log.Fatal(
					fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 直前の値 %s 文字カウンタ %d",
						setting.CountLF, *r, tmp, *i))
			}
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
		}
	}
	return false
} // }}}

func forAssignToMemory(rs *[]rune, r *rune, i *int, l int) { // {{{
	// メモリに代入する
	// X1.0とか
	// ここにはrがアルファベットの場合しか来ないはず

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
		isDigit := IsDigit(next)                                           //数値
		isPMDigit := (IsPM(next) && *i+2 <= l && IsDigit((*rs)[*i+2]))     // +-付きの数値
		isHashDigit := (IsHash(next) && *i+2 <= l && IsDigit((*rs)[*i+2])) // #付きの数値
		if isDigit || isPMDigit || isHashDigit {
			// 後ろの数値を読む #はスキップして読み込む
			numStr := readOptionNumber(&next, rs, i) // 中でiが進む
			// 文字 + 数値 or 文字 + 変数
			// G01 とか X#10とか
			if *r == 'G' {
				// Gは特別扱い
				// G専用のQueueに突っ込む
				// G90だったらEnqueueForG(90) みたいにする
				if IsHash(next) {
					// まあないはずだけど G#10 みたいに変数使ってきたらという想定
					// EnqueueForG(Hash[90]) みたいにする
					enqueueForG(Hash(numStr).String())
				} else {
					// G01みたいなの
					// GOTO とかは来ない
					// EnqueueForG(1) みたいにする
					enqueueForG(numStr)
				}
			} else {
				// G以外のもの M08 とか X1. とか Y#10 とか

				// G90G00X100.とかでは、X100.の時点でGのキューに要素がある
				// G以外の代入が走る前にGを処理する
				flushGqueue()
				if IsHash(next) {
					Assign(string(*r), Hash(numStr).String()) // X#100とか Assign(P, Hash[412]) みたいにする
				} else {
					// Assign(O, 90)とかAssign(X, "10")とかAssign(X, "10.0")とかみたいにする
					Assign(string(*r), numStr)
				}
			}
			*r = (*rs)[*i] // readOptionNumber中でiが数値分スキップしているのでpreの保存用に更新しておく
		}
	}
} // }}}

func forOptionalSkipBlock(rs *[]rune, i int) { // {{{
	// オプショナルスキップブロック
	// オプショナルスキップはメモリの値を読んで
	// trueなら無視する。

	if t := (*rs)[i+1]; '1' <= t && t <= '9' {
		// 番号付きオプショナルスキップブロック
		tmpNum, _ := strconv.Atoi(string(t))
		if OptionalSkip[tmpNum] == false {
			// 無視しない
			setting.IsOptionalSkip = false
		} else {
			// 無視する
			setting.IsOptionalSkip = true
		}
	} else {
		if OptionalSkip[0] {
			// 無視しない
			setting.IsOptionalSkip = false
		} else {
			// 無視する
			setting.IsOptionalSkip = true
		}
	}
} // }}}

// 最小設定単位
type ISUnit struct { // {{{
	mm  float64
	in  float64
	deg float64
} // }}}

type commonSetting struct { // {{{
	IsMm                 bool    // mm か inchか
	IS                   *ISUnit // 機械によって設定されている最小設定単位
	FeedG00              float64 // G00送り速度 初期値
	FeedG01              float64 // G00送り速度 初期値
	IsG90                bool    // true : アブソリュート指令 false : インクリメンタル指令
	PlaneDesignation     int     // G17/18/19
	CutMode              int     // G00 01 02 03 を 0, 1, 2, 3 であらわす
	IsProhibitAssignAxis bool    // そのブロックでの座標値への代入を禁ずる
	CountLF              int
	IsOptionalSkip       bool
	CumulativeTime       float64 // 累積時間
} // }}}

var ( // {{{
	setting = commonSetting{}

	// 全体を格納するスライス
	Memory = make([]Value, 10000)

	// G専用Queue
	Gqueue = make([]Value, 0, 100)

	// G65で使われる座標をそのまま使う
	key = map[string]int{ // {{{ iotaでもいいけどなんだかんだ見るので
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7, // 引数として使用できない
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12, // 引数として使用できない
		"M": 13,
		"N": 14, // 引数として使用できない
		"O": 15, // 引数として使用できない
		"P": 16, // 引数として使用できない ことになっているが...G65
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	} // }}}
	// 引数指定2
	_I = make([]int, 10) // _I[0]は使用しない
	_J = make([]int, 10) // _I[0]は使用しない
	_K = make([]int, 10) // _I[0]は使用しない

	// オプショナルスキップブロック ボタンがONならそこで停止する
	// とりあえず/,/1,/2,/3,/4,/5,/6,/7,/8,/9までの10個分
	// 実際には外部テキストファイルとかにボタン設定書いてもらうのがいいと思う
	OptionalSkip = make([]bool, 10) // 初期値false

	// オプショナルストップブロック
	// M01でボタンがONならそこで停止する
	// エミュレーションではキー入力町するのがいいと思う
	OptionalStop = make([]bool, 10) // 初期値false

	ISA = ISUnit{0.01, 0.001, 0.01}
	ISB = ISUnit{0.001, 0.0001, 0.001}
	ISC = ISUnit{0.0001, 0.00001, 0.0001}
	ISD = ISUnit{0.00001, 0.000001, 0.00001}
	ISE = ISUnit{0.000001, 0.0000001, 0.000001}

	rowInput *string
) // }}}

func initialize(apath *string) { // {{{

	// 入力ファイルを開く
	rowInput = readText(apath) // NCを読み込んでstringに変換、改行コードを統一

	rowLines = strings.Split(*rowInput, "\n")

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

// G専用Queue int, float64, string対応
func enqueueForG(v interface{}) { // {{{
	// vの型判定
	if value, ok := v.(int); ok {
		// int
		Gqueue = append(Gqueue,
			Value{
				bInt: true,
				f:    float64(value),
			},
		)
	} else if value, ok := v.(float64); ok {
		// float
		Gqueue = append(Gqueue,
			Value{
				bInt: false,
				f:    value,
			},
		)
		Memory[key["G"]].assignFloat(value)
	} else if value, ok := v.(string); ok {
		// 小数点があるかどうか調べる
		if strings.Contains(value, ".") {
			// float
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("\n")
				log.Fatal(fmt.Sprintf("エラー : バグ : 文字列からfloat64への変換に失敗しました。%%v %v : err %v", v, err))
			}
			Gqueue = append(Gqueue,
				Value{
					bInt: false,
					f:    v,
				},
			)
		} else {
			// int
			v, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("\n")
				log.Fatal(fmt.Sprintf("エラー : バグ : 文字列からintへの変換に失敗しました。%%v %v : err %v", v, err))
			}
			Gqueue = append(Gqueue,
				Value{
					bInt: true,
					f:    float64(v),
				},
			)
		}
	}

} // }}}

func dequeueForG() (Value, error) { // {{{
	if len(Gqueue) == 0 {
		return Value{
			bInt: false,
			f:    0,
		}, nil
	}
	ret := Gqueue[0]
	Gqueue = Gqueue[1:]
	return ret, nil
} // }}}

func flushGqueue() { // {{{1
	if len(Gqueue) == 0 {
		return
	}
	// Gの指定がある
	// Gqueueから取り出して処理する
	for i := len(Gqueue); i > 0; i-- {

		if setting.IsProhibitAssignAxis {
			// スキップフラグ立ってたら キューを空にして 終わる
			Gqueue = Gqueue[len(Gqueue):]
			break
		}
		v, _ := dequeueForG()
		if v.bInt == false { // {{{2
			// 小数
			switch {
			// 見つけたらエラー扱い(未対応で処理停止)を出してほしいワード
			// G43.4, G43.5 (XYZABCの扱いが大きく変わるため、来たら未対応を表明したい)
			case v.f == 43.4:
				fmt.Println("")
				log.Printf("エラー : l.%d : G43.4 は未対応です。行を無視します。", setting.CountLF)
				setting.IsProhibitAssignAxis = true
				continue
			case v.f == 43.5:
				fmt.Println("")
				log.Printf("エラー : l.%d : G43.5 は未対応です。行を無視します。", setting.CountLF)
				setting.IsProhibitAssignAxis = true
				continue
			default:
				// 未実装
				fmt.Println("")
				log.Printf("注意 : l.%d : G%2d.% .2f は未対応です。行を無視します。", setting.CountLF, int(v.f), v.f-float64(int(v.f)))
				// 以降のパラメーターを座標として読み込んまない
				// 座標値への代入を禁止する // これが立っていると改行まで無視する
				setting.IsProhibitAssignAxis = true
				continue
			}
			return
		} // }}}2

		// 以下整数
		n := int(v.f)
		switch { // {{{2
		case n == 0:
			setting.CutMode = 0
			// 一旦現在の送り速度を保存
			setting.FeedG01 = Reference("F").Float()
			Assign("F", setting.FeedG00)
		case n == 1:
			Assign("F", setting.FeedG01) // 直前がG00だったときの為に復元する
			setting.CutMode = 1
		case n == 2:
			Assign("F", setting.FeedG01) // 直前がG00だったときの為に復元する
			setting.CutMode = 2
			fmt.Println("")
			log.Printf("注意 : l.%d : G02 です。", setting.CountLF)
		case n == 3:
			Assign("F", setting.FeedG01) // 直前がG00だったときの為に復元する
			setting.CutMode = 3
			fmt.Println("")
			log.Printf("注意 : l.%d : G03 です。", setting.CountLF)
		case n == 17: // XY平面指定 G02/03を実装したら使われる
			setting.PlaneDesignation = 17
		case n == 18: // ZX平面指定 G02/03を実装したら使われる
			setting.PlaneDesignation = 18
		case n == 19: // YZ平面指定 G02/03を実装したら使われる
			setting.PlaneDesignation = 19
		case n == 90: // アブソリュート指令
			setting.IsG90 = true
		case n == 91: // インクリメンタル指令
			setting.IsG90 = false

		case n == 4: // ドウェル
			// スキップ
			fmt.Println("")
			log.Printf("注意 : l.%d : G04 ドウェルです。行を無視します。", setting.CountLF)
			setting.IsProhibitAssignAxis = true
			continue
		case n == 5: // AI輪郭制御モード
			// スキップ
			fmt.Println("")
			log.Printf("注意 : l.%d : G05 です。行を無視します。", setting.CountLF)
			setting.IsProhibitAssignAxis = true
			continue

		// 以下未実装
		case n == 6: // 放物線補完
			fallthrough
		case n == 7: // 未指定
			fallthrough
		case n <= 8 && n <= 9: // 加速 減速
			fallthrough
		case n <= 10 && n <= 16: // 未指定
			fallthrough
		case n <= 41 && n <= 42: // 工具径補正 左 右
			fallthrough
		case n <= 43 && n <= 44: // 工具オフセット オフセットのキャンセル
			fallthrough
		case n <= 45 && n <= 52: // 工具オフセット いろいろ
			fallthrough
		case n <= 53 && n <= 59: // シフト
			fallthrough
		case n <= 60 && n <= 61: // 正確な位置決め 精密 普通
			fallthrough
		case n == 62: // 迅速位置決め
			fallthrough
		case n <= 63 && n <= 79: // 未指定
			fallthrough
		case n == 80: // 固定サイクルのキャンセル
			fallthrough
		case n <= 81 && n <= 89: // 固定サイクル
			fallthrough
		case n == 92: // 座標系指定
			fallthrough
		case n == 93: // 時間の逆数で送りを指定
			fallthrough
		case n == 94: // 毎分当たりmm送り
			fallthrough
		case n == 95: // 主軸一回店当たり送り
			fallthrough
		case n <= 96 && n <= 97: // 定切削速度 定切削速度のキャンセル
			fallthrough
		case n <= 98 && n <= 99: // 未指定
			fallthrough
		default:
			// 未実装
			fmt.Println("")
			log.Printf("注意 : l.%d : G%2d は未対応です。行を無視します。", setting.CountLF, n)
			// 以降のパラメーターを座標として読み込んまない
			// 座標値への代入を禁止する // これが立っていると改行まで無視する
			setting.IsProhibitAssignAxis = true
			continue
		} // }}}2
	}
} // }}}1

func progressbar(progressVal int) { // {{{
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
} // }}}
