package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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
		if value, ok := v.(int); ok { // 整数が入ってきた
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
		} else if value, ok := v.(float64); ok { // 小数が入ってきた
			if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
				// 加算する
				Memory[key[k]].AssignFloat(tmp + value)
			} else {
				// 加算しない
				Memory[key[k]].AssignFloat(value)
			}
		} else if value, ok := v.(string); ok { // 文字列が入ってきた
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
