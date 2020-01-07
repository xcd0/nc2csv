package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"./util"
)

var lines []string

// 最小設定単位
type IS struct {
	mm  float64
	in  float64
	deg float64
}

var ISA IS
var ISB IS
var ISC IS
var ISD IS
var ISE IS

var ( // {{{1
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
	OptionalSkip = make(bool, 10) // 初期値false

	// オプショナルストップブロック
	// M01でボタンがONならそこで停止する
	// エミュレーションではキー入力町するのがいいと思う
	OptionalStop = make(bool, 10) // 初期値false
) // }}} 1

// G専用Queue {{{
func EnqueueForG(v interface{}) {
	// vの型判定
	if value, ok := v.(int); ok {
		// int
		Gqueue = append(Gqueue,
			Value{
				bInt: true,
				i:    value,
				f:    0,
			},
		)
	} else if value, ok := v.(float64); ok {
		// float
		Gqueue = append(Gqueue,
			Value{
				bInt: false,
				i:    0,
				f:    value,
			},
		)
		Memory[key[k]].AssignFloat(value)
	}

}

func DequeueForG() []Value {
	if len(Gqueue) == 0 {
		return nil
	}
	ret := Gqueue[0]
	Gqueue = Gqueue[1:]
	return ret
}

// }}}

// メモリ {{{

type Value struct {
	bInt bool
	i    int
	f    float64
}

func Reference(k string) Value {
	return Memory[key[k]]
}

func Assign(k string, v interface{}) {
	if value, ok := v.(int); ok {
		Memory[key[k]].AssignInt(value)
	} else if value, ok := v.(float64); ok {
		Memory[key[k]].AssignFloat(value)
	}
}

func (v *Value) IsInt() bool {
	return v.bInt
}

func (v *Value) AssignInt(i int) {
	v.bInt = true
	v.i = i
	v.f = 0
}

func (v *Value) AssignFloat(f float64) {
	v.bInt = false
	v.i = 0
	v.f = f
}

func (v *Value) String() string {
	if v.bInt {
		return string(v.i)
	} else {
		return fmt.Sprintf("%f", v.f)
	}
}

func (v *Value) Float() float64 {
	if v.bInt {
		return float64(v.i)
	} else {
		return v.f
	}
}

// }}}

func Initialize() { // {{{
	ISA.deg = 0.01
	ISA.in = 0.001
	ISA.mm = 0.01
	ISB.deg = 0.001
	ISB.in = 0.0001
	ISB.mm = 0.001
	ISC.deg = 0.0001
	ISC.in = 0.00001
	ISC.mm = 0.0001
	ISD.deg = 0.00001
	ISD.in = 0.000001
	ISD.mm = 0.00001
	ISE.deg = 0.000001
	ISE.in = 0.0000001
	ISE.mm = 0.000001
} // }}}

func CreateSrc(rowInput string) string {

	Initialize()

	// コメントを削除
	clearInput := util.DeleteComment(rowInput)

	log.Printf("clearInput : %v", clearInput)

	lines = strings.Split(rowInput, "\n")

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
func (a *Axis) outputOneline(mode, countLF int) string { // {{{

	if mode == 1 {
		panic("インクリメンタル指令は未実装です")
	}

	a.dX = (Reference("X").Float() - a.X)
	a.dY = (Reference("Y").Float() - a.Y)
	a.dZ = (Reference("Z").Float() - a.Z)
	a.dA = (Reference("A").Float() - a.A)
	a.dB = (Reference("B").Float() - a.B)
	a.dC = (Reference("C").Float() - a.C)
	// 移動距離
	dEuclideanDistance := math.Sqrt(math.Pow(a.dX, 2) + math.Pow(a.dY, 2) + math.Pow(a.dZ, 2) + math.Pow(a.dA, 2) + math.Pow(a.dB, 2) + math.Pow(a.dC, 2))
	// 移動時間
	dTimeMin := dEuclideanDistance / Reference("F").Float()

	// 元ncプログラムの行、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	// 元ncプログラムの行
	out = string(countLf)
	// XYZABC の各位置
	out += "," + Reference("X").String()
	out += "," + Reference("Y").String()
	out += "," + Reference("Z").String()
	out += "," + Reference("A").String()
	out += "," + Reference("B").String()
	out += "," + Reference("C").String()
	// プログラムの F
	out += "," + Reference("F").String()
	// XYZABC の各軸速度
	out += "," + string(a.dX/dTimeMin)
	out += "," + string(a.dY/dTimeMin)
	out += "," + string(a.dZ/dTimeMin)
	out += "," + string(a.dA/dTimeMin)
	out += "," + string(a.dB/dTimeMin)
	out += "," + string(a.dC/dTimeMin)
	// 移動に要する時間
	out += "," + string(dTimeMin)

	return out
} // }}}

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
			if pre == '\n' && i+1 < l {
				// オプショナルスキップブロック
				// オプショナルスキップはメモリの値を読んで
				// trueなら無視する。

				bOptionalSkip = true
				if t := rs[i+1]; '1' <= t && t <= '9' {
					// 番号付きオプショナルスキップブロック
					if OptionalSkip[strconv.Atoi(string(r))] == false {
						// 無視しない
						axis.outputOneline(G90_91, countLF)
					} else {
						// 無視する
						axis.outputOneline(G90_91, countLF)
					}
				} else {
					if OptionalSkip[0] {
						// 無視しない
					} else {
						// 無視する
					}
				}
			} else {
				// 除算の/
				// 未実装予約語エラー
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 演算子 / %v は未実装です。", countLF, r))
			}
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
