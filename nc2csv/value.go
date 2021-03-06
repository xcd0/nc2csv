package nc2csv

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// valueはメモリの1つのアドレスに格納される値の実態です。
// 入力が整数であったか小数であったかを保持する必要があり数値のほかにboolの値を保持しています。
type value struct {
	bInt bool
	f    float64
}

// メモリの値を参照します。
// 例としてReference("X") とするとXの値value型としてが返ってきます。
func Reference(k string) *value {
	return &memory[key[k]]
}

// メモリに値を代入します。
// 例としてAssign("X", 10.)とするとXに10.0を代入します。
// 第2引数はint,float64,stringが取れます。
func Assign(k string, v interface{}) {
	if setting.IsProhibitAssignAxis {
		// フラグ立ってたら何もしない
		return
	}
	// 代入時にインクリメンタル指令もすべて計算して座標値を入れる
	if setting.IsG90 { // アブソリュート指令 {{{
		if value, ok := v.(int); ok {
			// 座標値の数値代入において整数が代入されたときは、
			// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
			if k == "X" || k == "Y" || k == "Z" || k == "R" {
				if setting.IsMm {
					memory[key[k]].assignFloat(float64(value) * setting.IS.mm)
				} else {
					memory[key[k]].assignFloat(float64(value) * setting.IS.in)
				}
			} else if k == "A" || k == "B" || k == "C" {
				// 角度は剰余を取って0~360に丸める
				d := float64(value) * setting.IS.deg
				for d > 360 {
					d -= 360.0
				}
				if d < 0 {
					d += 360
				}
				memory[key[k]].assignFloat(d)

			} else {
				memory[key[k]].assignInt(value)
			}
		} else if value, ok := v.(float64); ok {
			if k == "A" || k == "B" || k == "C" {
				d := float64(value) * setting.IS.deg
				if d < 0 {
					d += 360
				}
				for d > 360 {
					d -= 360.0
				}
				memory[key[k]].assignFloat(d)
			} else {
				memory[key[k]].assignFloat(value)
			}
		} else if value, ok := v.(string); ok {
			// 小数点があるかどうか調べる
			if strings.Contains(value, ".") {
				// 小数点がある
				d, _ := strconv.ParseFloat(value, 64)
				if k == "A" || k == "B" || k == "C" {
					for d > 360 {
						d -= 360.0
					}
					if d < 0 {
						d += 360
					}
					memory[key[k]].assignFloat(d)
				} else {
					memory[key[k]].assignFloat(d)
				}
			} else {
				n, _ := strconv.Atoi(value)
				// 座標値の数値代入において整数が代入されたときは、
				// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
				if k == "X" || k == "Y" || k == "Z" || k == "R" {
					if setting.IsMm {
						memory[key[k]].assignFloat(float64(n) * setting.IS.mm)
					} else {
						memory[key[k]].assignFloat(float64(n) * setting.IS.in)
					}
				} else if k == "A" || k == "B" || k == "C" {
					d := float64(n) * setting.IS.deg
					for d > 360 {
						d -= 360.0
					}
					if d < 0 {
						d += 360
					}
					memory[key[k]].assignFloat(d)
				} else {
					// XYZABCR以外は整数値そのまま入れる
					// F1000とか
					memory[key[k]].assignInt(n)
				}
			}
		} // }}}
	} else { // インクリメンタル指令 {{{
		// 座標値のみ 一旦取り出して加算して代入する
		tmp := 0.0
		if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
			tmp = memory[key[k]].Float()
		}
		if value, ok := v.(int); ok { // 整数が入ってきた
			// 座標値の数値代入において整数が代入されたときは、
			// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
			if k == "X" || k == "Y" || k == "Z" || k == "R" {
				// 加算する
				if setting.IsMm {
					memory[key[k]].assignFloat(tmp + float64(value)*setting.IS.mm)
				} else {
					memory[key[k]].assignFloat(tmp + float64(value)*setting.IS.in)
				}
			} else if k == "A" || k == "B" || k == "C" {
				// 加算する
				memory[key[k]].assignFloat(tmp + float64(value)*setting.IS.deg)
			} else {
				// 座標値ではないので加算しない
				memory[key[k]].assignInt(value)
			}
		} else if value, ok := v.(float64); ok { // 小数が入ってきた
			if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
				// 加算する
				memory[key[k]].assignFloat(tmp + value)
			} else {
				// 加算しない
				memory[key[k]].assignFloat(value)
			}
		} else if value, ok := v.(string); ok { // 文字列が入ってきた
			// 小数点があるかどうか調べる
			if strings.Contains(value, ".") {
				// 小数点がある
				n, _ := strconv.ParseFloat(value, 64)
				if k == "X" || k == "Y" || k == "Z" || k == "R" || k == "A" || k == "B" || k == "C" {
					// 加算する
					memory[key[k]].assignFloat(tmp + n)
				} else {
					// 加算しない
					memory[key[k]].assignFloat(n)
				}
			} else {
				n, _ := strconv.Atoi(value)
				// 座標値の数値代入において整数が代入されたときは、
				// 最小値の定数倍としてv.fに代入し、v.bIntをfalseにする
				if k == "X" || k == "Y" || k == "Z" || k == "R" {
					// 加算する
					if setting.IsMm {
						memory[key[k]].assignFloat(tmp + float64(n)*setting.IS.mm)
					} else {
						memory[key[k]].assignFloat(tmp + float64(n)*setting.IS.in)
					}
				} else if k == "A" || k == "B" || k == "C" {
					// 加算する
					memory[key[k]].assignFloat(tmp + float64(n)*setting.IS.deg)
				} else {
					// 加算しない
					memory[key[k]].assignInt(n)
				}
			}
		}
	} // }}}
}

// #の記法でアドレス値を参照します。
// #10はHash(10),Hash(10.0),Hash("10")のように入力され、&memory[10]が返されます。
func Hash(v interface{}) *value { // {{{
	out := 0
	// ただ 入力引数の型をint, float64, stringの3つどれでもよくしているだけ
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
	return &memory[out]
} // }}}

// メモリの値が整数であるかどうかを返します。
func (v *value) IsInt() bool {
	return v.bInt
}

// メモリの値に整数を代入します。
func (v *value) assignInt(i int) {
	v.bInt = true
	v.f = float64(i)
}

// メモリの値に小数を代入します。
func (v *value) assignFloat(f float64) {
	v.bInt = false
	v.f = f
}

// メモリの値を文字列として返します。
func (v *value) String() string {
	if v.bInt {
		return fmt.Sprintf("%d", int(v.f))
	} else {
		return fmt.Sprintf("%.10f", v.f)
	}
}

// メモリの値を小数として返します。
func (v *value) Float() float64 {
	return v.f
}

// メモリの値を整数として返します。
func (v *value) Int() int {
	return int(v.f)
}
