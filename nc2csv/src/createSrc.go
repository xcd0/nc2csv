package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"./util"
)

var axis Axis
var lines []string
var rowLines []string

// 最小設定単位
type ISUnit struct {
	mm  float64
	in  float64
	deg float64
}

type CommonSetting struct {
	IsMm    bool    // mm か inchか
	IS      *ISUnit // 機械によって設定されている最小設定単位
	FeedG00 float64 // G00送り速度 初期値
	FeedG01 float64 // G01送り速度 初期値
	IsG90   bool    // true : アブソリュート指令 false : インクリメンタル指令
	CutMode int     // G00 01 02 03 を 0, 1, 2, 3 であらわす
}

var ( // {{{1
	Setting = CommonSetting{}

	// 全体を格納するスライス
	Memory = make([]Value, 10000)

	// G65で使われる座標をそのまま使う
	key = map[string]int{ // {{{2
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
	} // }}} 2

	// G専用Queue
	Gqueue = make([]Value, 0, 100)

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
) // }}} 1

// G専用Queue {{{
func EnqueueForG(v interface{}) {
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
		Memory[key["G"]].AssignFloat(value)
	}

}

func DequeueForG() (Value, error) {
	if len(Gqueue) == 0 {
		return Value{
			bInt: false,
			f:    0,
		}, nil
	}
	ret := Gqueue[0]
	Gqueue = Gqueue[1:]
	return ret, nil
}

// }}}

// メモリ {{{

type Value struct {
	bInt bool
	f    float64
}

func Hash(v interface{}) *Value {
	out := 0
	if value, ok := v.(int); ok {
		out = value
	} else if value, ok := v.(float64); ok {
		out = int(value)
	} else if value, ok := v.(string); ok {
		// 小数点があるかどうか調べる
		if strings.Contains(".", value) {
			// 小数点がある
			n, _ := strconv.ParseFloat(value, 64)
			out = int(n)
		} else {
			out, _ = strconv.Atoi(value)
		}
	} else {
		// エラー
		log.Fatal(fmt.Sprintf("書式エラー : %v はエラーです。", v))
	}
	return &Memory[out]
}

func Reference(k string) *Value {
	return &Memory[key[k]]
}

func Assign(k string, v interface{}) {

	// 代入時にインクリメンタル指令もすべて計算して座標値を入れる
	if Setting.IsG90 { // アブソリュート指令 {{{
		if value, ok := v.(int); ok {
			// 座標値の数値代入において整数が代入されたときは、
			// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
			if k == "X" || k == "Y" || k == "Z" || k == "R" {
				if Setting.IsMm {
					Memory[key[k]].AssignFloat(float64(value) * Setting.IS.mm)
				} else {
					Memory[key[k]].AssignFloat(float64(value) * Setting.IS.in)
				}
			} else if k == "A" || k == "B" || k == "C" {
				Memory[key[k]].AssignFloat(float64(value) * Setting.IS.deg)
			} else {
				Memory[key[k]].AssignInt(value)
			}
		} else if value, ok := v.(float64); ok {
			Memory[key[k]].AssignFloat(value)
		} else if value, ok := v.(string); ok {
			// 小数点があるかどうか調べる
			if strings.Contains(value, ".") {
				// 小数点がある
				n, _ := strconv.ParseFloat(value, 64)
				Memory[key[k]].AssignFloat(n)
			} else {
				n, _ := strconv.Atoi(value)
				// 座標値の数値代入において整数が代入されたときは、
				// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
				if k == "X" || k == "Y" || k == "Z" || k == "R" {
					if Setting.IsMm {
						Memory[key[k]].AssignFloat(float64(n) * Setting.IS.mm)
					} else {
						Memory[key[k]].AssignFloat(float64(n) * Setting.IS.in)
					}
				} else if k == "A" || k == "B" || k == "C" {
					Memory[key[k]].AssignFloat(float64(n) * Setting.IS.deg)
				} else {
					Memory[key[k]].AssignInt(n)
				}
			}
		} // }}}
	} else { // インクリメンタル指令
		// 座標値のみ 一旦取り出して加算して代入する
		tmp := 0.0
		if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
			tmp = Memory[key[k]].Float()
		}
		if value, ok := v.(int); ok {
			// 座標値の数値代入において整数が代入されたときは、
			// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
			if k == "X" || k == "Y" || k == "Z" || k == "R" {
				// 加算する
				if Setting.IsMm {
					Memory[key[k]].AssignFloat(tmp + float64(value)*Setting.IS.mm)
				} else {
					Memory[key[k]].AssignFloat(tmp + float64(value)*Setting.IS.in)
				}
			} else if k == "A" || k == "B" || k == "C" {
				// 加算する
				Memory[key[k]].AssignFloat(tmp + float64(value)*Setting.IS.deg)
			} else {
				// 座標値ではないので加算しない
				Memory[key[k]].AssignInt(value)
			}
		} else if value, ok := v.(float64); ok {
			if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
				// 加算する
				Memory[key[k]].AssignFloat(tmp + value)
			} else {
				// 加算しない
				Memory[key[k]].AssignFloat(value)
			}
		} else if value, ok := v.(string); ok {
			// 小数点があるかどうか調べる
			if strings.Contains(value, ".") {
				// 小数点がある
				n, _ := strconv.ParseFloat(value, 64)
				if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
					// 加算する
					Memory[key[k]].AssignFloat(tmp + n)
				} else {
					// 加算しない
					Memory[key[k]].AssignFloat(n)
				}
			} else {
				n, _ := strconv.Atoi(value)
				// 座標値の数値代入において整数が代入されたときは、
				// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
				if k == "X" || k == "Y" || k == "Z" || k == "R" {
					// 加算する
					if Setting.IsMm {
						Memory[key[k]].AssignFloat(tmp + float64(n)*Setting.IS.mm)
					} else {
						Memory[key[k]].AssignFloat(tmp + float64(n)*Setting.IS.in)
					}
				} else if k == "A" || k == "B" || k == "C" {
					// 加算する
					Memory[key[k]].AssignFloat(tmp + float64(n)*Setting.IS.deg)
				} else {
					// 加算しない
					Memory[key[k]].AssignInt(n)
				}
			}
		}
	}

}

func (v *Value) IsInt() bool {
	return v.bInt
}

func (v *Value) AssignInt(i int) {
	v.bInt = true
	v.f = float64(i)
}

func (v *Value) AssignFloat(f float64) {
	v.bInt = false
	v.f = f
}

func (v *Value) String() string {
	if v.bInt {
		return fmt.Sprintf("%d", int(v.f))
	} else {
		return fmt.Sprintf("%f", v.f)
	}
}

func (v *Value) Float() float64 {
	return v.f
}

func (v *Value) Int() int {
	return int(v.f)
}

// }}}

func Initialize(rowInput *string) { // {{{
	rowLines = strings.Split(*rowInput, "\n")

	// 初期設定
	Setting.IsMm = true    // mmか
	Setting.IS = &ISC      // 最小設定単位の指定 // とりあえずISCとしてみる
	Setting.FeedG00 = 5000 // 早送り速度
	Setting.FeedG01 = 1000 // 送り速度
	Setting.IsG90 = true   // アブソリュート指令か
	Setting.CutMode = 0    // 切削モード

} // }}}

func CreateSrc(rowInput string) string {

	Initialize(&rowInput)

	// コメントを削除
	clearInput := util.DeleteComment(rowInput)

	log.Printf("clearInput : %v", clearInput)

	runfunc := makeRunFunction(clearInput)

	log.Printf("runfunc : %v", runfunc)

	return runfunc
}

// 表示用に座標を持っておくだけ
type Axis struct { // {{{
	X  float64
	Y  float64
	Z  float64
	A  float64
	B  float64
	C  float64
	dX float64
	dY float64
	dZ float64
	dA float64
	dB float64
	dC float64
} // }}}

// mode : 0(アブソリュート指令) 1(インクリメンタル指令)
func (a *Axis) outputOneline(countLF int) string { // {{{

	a.dX = (Reference("X").Float() - a.X)
	a.dY = (Reference("Y").Float() - a.Y)
	a.dZ = (Reference("Z").Float() - a.Z)
	a.dA = (Reference("A").Float() - a.A)
	a.dB = (Reference("B").Float() - a.B)
	a.dC = (Reference("C").Float() - a.C)
	// 移動距離
	dEuclideanDistance := math.Sqrt(math.Pow(a.dX, 2) + math.Pow(a.dY, 2) + math.Pow(a.dZ, 2) + math.Pow(a.dA, 2) + math.Pow(a.dB, 2) + math.Pow(a.dC, 2))
	// 移動時間
	dTimeMin := 0.0
	f := Reference("F").Float()
	if f == 0 {
		// 送り速度が 0
	}
	if dEuclideanDistance != 0 && f != 0 {
		dTimeMin = dEuclideanDistance / f
	}

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	// 元ncプログラムの行
	out := fmt.Sprintf("%d", countLF)
	// その行のncプログラム
	out += "," + rowLines[countLF-1]
	// XYZABC の各位置
	out += "," + Reference("X").String()
	out += "," + Reference("Y").String()
	out += "," + Reference("Z").String()
	out += "," + Reference("A").String()
	out += "," + Reference("B").String()
	out += "," + Reference("C").String()
	out += "," + Reference("R").String()
	// プログラムの F
	out += "," + Reference("F").String()
	// XYZABC の各軸速度
	vX, vY, vZ, vA, vB, vC := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	if a.dX != 0 {
		vX = a.dX / dTimeMin
	}
	if a.dY != 0 {
		vY = a.dY / dTimeMin
	}
	if a.dZ != 0 {
		vZ = a.dZ / dTimeMin
	}
	if a.dA != 0 {
		vA = a.dA / dTimeMin
	}
	if a.dB != 0 {
		vB = a.dB / dTimeMin
	}
	if a.dC != 0 {
		vC = a.dC / dTimeMin
	}
	out += "," + fmt.Sprintf("%f", vX)
	out += "," + fmt.Sprintf("%f", vY)
	out += "," + fmt.Sprintf("%f", vZ)
	out += "," + fmt.Sprintf("%f", vA)
	out += "," + fmt.Sprintf("%f", vB)
	out += "," + fmt.Sprintf("%f", vC)
	// 移動に要する時間
	out += "," + fmt.Sprintf("%f", dTimeMin)

	// 保存
	a.X = Reference("X").Float()
	a.Y = Reference("Y").Float()
	a.Z = Reference("Z").Float()
	a.A = Reference("A").Float()
	a.B = Reference("B").Float()
	a.C = Reference("C").Float()

	log.Printf("%v\n", out)
	return out
} // }}}

func makeRunFunction(input string) string {
	in := make(chan string, 100) // 別スレッドに投げるバッファ
	out := make(chan string, 2)  // 別スレッドからもらうバッファ
	done := make(chan bool)      // 別スレッドの終了通知をもらうバッファ

	// 別スレッドで受ける
	go srcOutput(in, out, done)

	rs := []rune(input)

	lines := strings.Split(input, "\n")

	countLF := 1
	l := len(rs)
	pre := '\n'

	// 適当な初期値
	Assign("F", 1000) // 送り速度 初期値

	bOptionalSkip := false

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	in <- "line,NC,X,Y,Z,A,B,C,R,F,vX,vY,vZ,vA,vB,vC,time"

	// 行ごとにcase文を出力する
	for i := 0; i < l; i++ {
		r := rs[i]
		if r == '\n' {
			log.Printf("%s\n", "\\n")
		} else {
			log.Printf("%c\n", r)
		}
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
						bOptionalSkip = false
					} else {
						// 無視する
						bOptionalSkip = true
					}
				} else {
					if OptionalSkip[0] {
						// 無視しない
						bOptionalSkip = false
					} else {
						// 無視する
						bOptionalSkip = true
					}
				}
			} else {
				// 除算の/
				// 未実装予約語エラー
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 演算子 / %v は未実装です。", countLF, r))
			}
		// }}}
		case '\n': // {{{

			if bOptionalSkip {
				// スキップする つまり何もしない。
				// 状態を変えずに出力する
				// といっても何もしないという処理はほかの文字に対して行うので
				// ここでは単にフラグをリセットする
				bOptionalSkip = false
			}

			// この行を実行した後の状態を出力する
			in <- axis.outputOneline(countLF)

			// countLFは1からだけどlinesは0から
			log.Printf("l.%v : %v\n", countLF-1, lines[countLF-1])

			countLF++
			if string(rs[len(rs)-1:]) == "\n" && countLF == len(lines) {
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
			if bOptionalSkip {
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
							log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", countLF, literal))
						}
					} else if i+1 >= l {
						// 謎
						// 文字数が超過している。
						log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", countLF, r))
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
									numStr := getOptionNumber(&next, &rs, &i)
									// EnqueueForG(Hash[90]) みたいにする
									tmpNum, _ := strconv.Atoi(numStr)
									EnqueueForG(Hash(tmpNum))
									r = rs[i] // 中でiが数値分スキップしているのでpreの保存用に更新しておく
								} else {
									// G01みたいなの
									// GOTO とかは来ない
									numStr := getOptionNumber(&next, &rs, &i)
									// EnqueueForG(1) みたいにする
									tmpNum, _ := strconv.Atoi(numStr)
									EnqueueForG(tmpNum)
								}
							} else {
								// G以外のもの M08 とか X1. とか Y#10 とか
								if next == '#' {
									// X#100とか
									// この関数は#100とかの場合でも#を無視して"100"を返してくれる
									numStr := getOptionNumber(&next, &rs, &i)
									// Assign(P, Hash[412]) みたいにする
									Assign(string(r), Hash(numStr)) // 中でiは進んでいる
									r = rs[i]                       // 中でiが数値分スキップしているのでpreの保存用に更新しておく
								} else if util.IsLetter(r) {
									// X1.とか X-10.とか
									numStr := getOptionNumber(&next, &rs, &i)
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
									log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 %s", countLF, r, tmp))
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

func readAndCutNumber(rs *[]rune, i *int) string { // {{{
	numStr := util.ReadNumbers(rs, i)
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
} // }}}

func getOptionNumber(next *rune, rs *[]rune, i *int) string { // {{{
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
} // }}}
