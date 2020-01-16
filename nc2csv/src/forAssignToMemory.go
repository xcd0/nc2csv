package main

import (
	"fmt"
	"log"

	"./util"
)

func forAssignToMemory(rs *[]rune, r *rune, i *int, l int) {
	// ここにはrがアルファベットの場合しか来ない

	// M08 G00 X1. Y2
	// こういうのはすべて変数への代入と見なす
	// var M = 8 みたいなかんじ
	// 但しGだけ特別扱いする
	// NCの場合varに当たるprefixがないので
	// GOTO,IF,WHILEでないアルファベットが来たら変数代入だとみなす
	// 異常値だがAAと来たらAAも変数代入になる？

	// readIdentifier()はアルファベットまたは_がつづく間読み取って返す
	literal := util.ReadLetters(rs, *i)
	if util.IsReserved(literal) {
		// 予約語
		// GOTO IF WHILE THEN の予定？
		// TODO
		if util.IsImplemented(literal) {
			// 実装済み予約語
		} else {
			// 未実装予約語
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", setting.CountLF, literal))
		}
	} else if *i+1 >= l {
		// 謎
		// 文字数が超過している。
		log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, *r))
	} else {
		// 文字数チェックOK
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