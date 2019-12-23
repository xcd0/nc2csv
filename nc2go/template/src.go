package main

import (
	"fmt"
)

var ( // {{{
	// 全体を格納するスライス
	Memory = make([]Value, 10000)

	// G65で使われる座標をそのまま使う
	key = map[string]int{
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
	}

	// 引数指定2
	_I = make([]int, 10) // _I[0]は使用しない
	_J = make([]int, 10) // _I[0]は使用しない
	_K = make([]int, 10) // _I[0]は使用しない
) // }}}

func main() {
	// inputは最下行に定義
	preprocess()

	// 一行実行
	//program.length = {{ .length }}
	program.length = 20

	pc := 1 // プログラムカウンタ // 中でGOTOの時にPCを弄れるようにする
	for {
		program.run(&pc)
		pc++
	}
}

// メモリの内容に従って実行
// 未実装
func runBlock(pc *int) {
	fmt.Printf("G%v,X%v,Y%v,Z%v,A%v,B%v,C%v,F%v,S%v,T%v",
		Reference("G"),
		Reference("X"),
		Reference("Y"),
		Reference("Z"),
		Reference("A"),
		Reference("B"),
		Reference("C"),
		Reference("F"),
		Reference("S"),
		Reference("T"),
	)
}

func preprocess() { // {{{
	// 引数指定2の設定
	for i := 1; i <= 10; i++ {
		_I[i] = 1 + 3*i
		_J[i] = 2 + 3*i
		_K[i] = 3 + 3*i
	}
} // }}}

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

// }}}

// }}}

// テンプレートで代入される
//const input = `{{ .input }}`
// 例
const input = // {{{
`%
O0001(ROBO 4X)
(COORD=CENTER)
(A-90.)
G49
G69
M72
M69
M03S1000(ROT_INITIAL)
G90G00
G90G10L11P#4120R-1.0
X0.Y0.
G43Z50.
G90
G01X-6.Y0.Z50.A-90.F1000
G00Z31.A-90.
G01X-3.F5000
X0.
Z26.F1000
A-90.912F2204
` // }}}

// const program = {{ .program }}
var program Program

type Program struct {
	length int
}

// ここはテンプレートとして読み込む
// 例 こういう関数を出力させる予定
func (p *Program) run(pc *int) { // {{{
	if *pc == 0 || *pc > p.length {
		// 範囲外アクセス エラー
		// 1から始まるのでpc == p.lengthはOK
		e := fmt.Sprintf("実行エラー : l.%d : その行は存在しません。エラーです。", *pc)
		panic(e) // 予約語でない複数文字列は書式エラー
	}

	NOP := false
	switch *pc {
	case 1: //  1  %
		NOP = true
	case 2: //  2  O0001(ROBO 4X)
		Assign("O", 1)
	case 3: //  3  (COORD=CENTER)
		NOP = true
	case 4: //  4  (A-90.)
		NOP = true
	case 5: //  5  G49
		Assign("G", 49)
	case 6: //  6  G69
		Assign("G", 69)
	case 7: //  7  M72
		Assign("M", 72)
	case 8: //  8  M69
		Assign("M", 69)
	case 9: //  9  M03S1000(ROT_INITIAL)
		Assign("M", 3)
		Assign("S", 1000)
	case 10: // 10  G90G00
		// 一行に同じGが出たとき
		Assign("G", 90)
		Assign("G", 00)
	case 11: // 11  G90G10L11P#4120R-1.0
		Assign("G", 90)
		Assign("G", 10)
		Assign("L", 11)
		Assign("P", Memory[412])
		Assign("R", -1.0)
	case 12: // 12  X0.Y0.
		Assign("X", 0.0)
		Assign("Y", 0.0)
	case 13: // 13  G43Z50.
		Assign("G", 90)
		Assign("G", 90)
	case 14: // 14  G90
		Assign("G", 90)
	case 15: // 15  G01X-6.Y0.Z50.A-90.F1000
		Assign("G", 1)
		Assign("X", -6.0)
		Assign("Y", 0.0)
		Assign("Z", 50.0)
		Assign("A", -90.0)
		Assign("F", 10000)
	case 16: // 16  G00Z31.A-90.
		Assign("G", 0)
		Assign("Z", 31.0)
		Assign("A", -90.0)
	case 17: // 17  G01X-3.F5000
		Assign("G", 1)
		Assign("X", -3.0)
		Assign("F", 5000)
	case 18: // 18  X0.
		Assign("X", 0.0)
	case 19: // 19  Z26.F1000
		Assign("Z", 26.1)
		Assign("F", 1000)
	case 20: // 20  A-90.912F2204
		Assign("A", -90.912)
		Assign("F", 2204)
	}
	if !NOP {
		runBlock(pc) //未実装
	}
}

/*
NCの例
 1  %
 2  O0001(ROBO 4X)
 3  (COORD=CENTER)
 4  (A-90.)
 5  G49
 6  G69
 7  M72
 8  M69
 9  M03S1000(ROT_INITIAL)
10  G90G00
11  G90G10L11P#4120R-1.0
12  X0.Y0.
13  G43Z50.
14  G90
15  G01X-6.Y0.Z50.A-90.F1000
16  G00Z31.A-90.
17  G01X-3.F5000
18  X0.
19  Z26.F1000
20  A-90.912F2204
*/
// }}}

// 元のNCプログラムも埋め込む
//var nc = `{{ .nc }`
var nc = // {{{
`%
O0001(ROBO 4X)
(COORD=CENTER)
(A-90.)
G49
G69
M72
M69
M03S1000(ROT_INITIAL)
G90G00
G90G10L11P#4120R-1.0
X0.Y0.
G43Z50.
G90
G01X-6.Y0.Z50.A-90.F1000
G00Z31.A-90.
G01X-3.F5000
X0.
Z26.F1000
A-90.912F2204
`

// }}}
