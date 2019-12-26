//%{
package main

import (
	"fmt"
)

const (
	memoryNum           = 10000  // メモリは何番地まで持つか
	minimalStraightMove = 0.0001 // 直線移動軸最⼩移動量
	minimalRotateMove   = 0.001  // 回転軸最⼩移動量
)

// データ型

// メモリセル
type Value float64 // {{{

func (v *Value) Put(f float64) {
	v.f = f
}
func (v *Value) Get() (float64, error) {
	return v.f, nil
}
func (v *Value) Add(l, r *Value) {
	v.Put(l.f + r.f)
}
func (v *Value) Sub(l, r *Value) {
	v.Put(l.f - r.f)
}
func (v *Value) Mul(l, r *Value) {
	v.Put(l.f * r.f)
}
func (v *Value) Div(l, r *Value) {
	v.Put(l.f / r.f)
}
func (v *Value) String() string {
	return fmt.Sprintf("%f", v.f)
} // }}}

// メモリ全体
type Memory struct { // {{{1
	memory       *[]Value
	_i           *[]int
	_j           *[]int
	_k           *[]int
	optionalSkip *[]int
	optionalStop *[]int
	variableKey  *map[string]int
	gQueue       *[]Value
}

func NewMemory() *Memory { // {{{2
	var cells = make(Value, memoryNum)

	// 引数指定2 IJKIJKIJK...って記法 [0]は使用しない
	var i, j, k = make(int, 10), make(int, 10), make(int, 10)

	// オプショナルスキップブロック ボタンがONならそこで停止する /,/1,/2,/3,/4,/5,/6,/7,/8,/9までの10個分
	// 実際には外部テキストファイルとかにボタン設定書いてもらうのがいいと思う
	var oSkip = make(bool, 10) // 初期値false

	// オプショナルストップブロック M01でボタンがONならそこで停止する エミュレーションではキー入力待ちするのがいいと思う
	var oStop = make(bool, 10) // 初期値false

	// G65で使われる座標をそのまま使う
	var v = map[string]int{ // {{{3
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
	} // }}}3

	var gq = make([]Value, 0, 100)

	var memory = Memory{
		memory:       &cells,
		_i:           &i,
		_j:           &j,
		_k:           &k,
		optionalSkip: &oSkip,
		optionalStop: &oStop,
		variableKey:  &v,
		gQueue:       &gq,
	}

	return &memory
}

// }}}2

// メモリから値を読み出す。文字でのアクセスや変数番号でのアクセスに対応。
// ex1) X  → x  := m.Get("X")
// ex2) #1 → h1 := m.Get(1)
func (m *Memory) Get(v interface{}) (*Value, error) {
	if value, ok := v.(string); ok {
		// 文字でのアクセス
		return m.memory[variableKey[k]], nil
	} else if value, ok := v.(int); ok {
		// #数値でのアクセス
		return m.memory[variableKey[k]], nil
	}
	e := fmt.Sprintf("メモリアクセスエラー : Get( %v ) はエラーです。", v)
	return nil, error.New(e)
}

// メモリに値を保存する。文字でのアクセスや変数番号でのアクセスに対応。
// ex1) G01    → m.Put("G", 1)
// ex2) #1 = 1 → m.Put(1, 1)
func (m *Memory) Put(k interface{}, v interface{}) {
	if key, ok := k.(string); ok {
		if value, ok := v.(int); ok {
			Memory[variableKey[key]].AssignInt(value)
		} else if value, ok := v.(float64); ok {
			Memory[variableKey[key]].AssignFloat(value)
		}
	} else if num, ok := k.(int); ok {
		if value, ok := v.(int); ok {
			Memory[num].AssignInt(value)
		} else if value, ok := v.(float64); ok {
			Memory[num].AssignFloat(value)
		}
	}
}
func (m *Memory) EnqueueForG(v interface{}) { // G専用Queueのenqueue {{{2
	// vの型判定
	if value, ok := v.(int); ok {
		// int
		m.gQueue = append(m.gQueue,
			Value{
				bInt: true,
				i:    value,
				f:    0,
			},
		)
	} else if value, ok := v.(float64); ok {
		// float
		m.gQueue = append(m.gQueue,
			Value{
				bInt: false,
				i:    0,
				f:    value,
			},
		)
	}
}

// }}}2
func (m *Memory) DequeueForG() []Value { // G専用Queueのdequeue {{{2
	if len(m.gQueue) == 0 {
		return nil
	}
	ret := m.gQueue[0]
	m.gQueue = m.gQueue[1:]
	return ret
}

// }}}2

// }}}1

// 数値はすべてValueで持つ。 というかMemory上に持つ。 // フラグでintかfloatか制御する。
type Expression interface {
	Get()
}

type Number Value

func (n *Number) Get() float64 {
	// 本来判定が必要...
	// 生の値を取り出したいときだけ
	return n.GetFloat()
}

type Operator struct {
	Left     *Value
	Operator string
	Right    *Value
}

func (s Operator) Get() *Value {
	var ans = make(Value)
	switch s.Operator {
	case "+":
		return &ans.Add(r1, r2) // l = r1 + r2
	case "-":
		return &ans.Div(r1, r2)
	case "*":
		return &ans.Mul(r1, r2)
	case "/":
		return &ans.Div(r1, r2)
	}
}
func runBlock(m *Memory, pc *int) {
	// とりあえずメモリの中身の主要な部分を出力
	fmt.Printf("%5d : G%v,X%v,Y%v,Z%v,A%v,B%v,C%v,F%v,S%v,T%v",
		*pc,
		m.Get("G"),
		m.Get("X"), m.Get("Y"), m.Get("Z"),
		m.Get("A"), m.Get("B"), m.Get("C"),
		m.Get("F"), m.Get("S"), m.Get("T"),
	)
}

//%}
