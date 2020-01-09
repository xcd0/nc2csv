package main

import (
	"strings"
)

// 最小設定単位
type ISUnit struct {
	mm  float64
	in  float64
	deg float64
}

type CommonSetting struct {
	IsMm                 bool    // mm か inchか
	IS                   *ISUnit // 機械によって設定されている最小設定単位
	FeedG00              float64 // G00送り速度 初期値
	FeedG01              float64 // G01送り速度 初期値
	IsG90                bool    // true : アブソリュート指令 false : インクリメンタル指令
	CutMode              int     // G00 01 02 03 を 0, 1, 2, 3 であらわす
	IsProhibitAssignAxis bool    // そのブロックでの座標値への代入を禁ずる
	CountLF              int
	IsOptionalSkip       bool
	CumulativeTime       float64 // 累積時間
}

var (
	Setting = CommonSetting{}

	// 全体を格納するスライス
	Memory = make([]Value, 10000)

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
)

func Initialize(rowInput *string) {
	rowLines = strings.Split(*rowInput, "\n")

	// 初期設定
	Setting.IsMm = true    // mmか
	Setting.IS = &ISC      // 最小設定単位の指定 // とりあえずISCとしてみる
	Setting.FeedG00 = 5000 // 早送り速度
	Setting.FeedG01 = 1000 // 送り速度
	Setting.IsG90 = true   // アブソリュート指令か
	Setting.CutMode = 0    // 切削モード

}
