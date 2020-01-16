package main

import (
	"fmt"
	"log"
	"math"
)

var axis Axis
var rowLines []string

// 表示用に座標を持っておくだけ
type Axis struct {
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
}

func (a *Axis) genOnelineCsv() (string, float64) {

	a.dX = (Reference("X").Float() - a.X)
	a.dY = (Reference("Y").Float() - a.Y)
	a.dZ = (Reference("Z").Float() - a.Z)

	// 回転する軸は0~360に丸めているので 境界で処理が必要
	// 例えばA359.0 A0はdA==1になる
	// 180度より大きい移動は近回りする
	// 180度ピッタリの移動はエラーにする
	a.dA = (Reference("A").Float() - a.A)
	a.dB = (Reference("B").Float() - a.B)
	a.dC = (Reference("C").Float() - a.C)

	if setting.CountLF == 527 {
		log.Println("aa")
	}

	shortcutDegree(&a.dA)
	shortcutDegree(&a.dB)
	shortcutDegree(&a.dC)
	// 移動距離
	var dEuclideanDistance float64
	dEuclideanDistance = math.Sqrt(math.Pow(a.dX, 2) + math.Pow(a.dY, 2) + math.Pow(a.dZ, 2) + math.Pow(a.dA, 2) + math.Pow(a.dB, 2) + math.Pow(a.dC, 2))
	// 移動時間
	var dTimeMin float64
	f := Reference("F").Float()
	if f == 0 {
		// 送り速度が 0
		log.Fatal(fmt.Sprintf("実行時エラー : l.%d : 送り速度が0です。", setting.CountLF))
	}
	if dEuclideanDistance != 0 && f != 0 {
		// Fは分単位
		dTimeMin = dEuclideanDistance / f
	}

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	// 元ncプログラムの行
	out := fmt.Sprintf("%d", setting.CountLF)
	// その行のncプログラム
	out += ",\"" + rowLines[setting.CountLF-1] + "\""
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
	out += "," + fmt.Sprintf("%.6f", vX)
	out += "," + fmt.Sprintf("%.6f", vY)
	out += "," + fmt.Sprintf("%.6f", vZ)
	out += "," + fmt.Sprintf("%.6f", vA)
	out += "," + fmt.Sprintf("%.6f", vB)
	out += "," + fmt.Sprintf("%.6f", vC)
	// 移動に要する時間
	out += "," + fmt.Sprintf("%.10f", dTimeMin)
	out += "," + fmt.Sprintf("%.10f", setting.CumulativeTime)

	// 保存
	a.X = Reference("X").Float()
	a.Y = Reference("Y").Float()
	a.Z = Reference("Z").Float()
	a.A = Reference("A").Float()
	a.B = Reference("B").Float()
	a.C = Reference("C").Float()

	//log.Printf("%v\n", out)
	return out, dTimeMin
}

func shortcutDegree(delta *float64) {
	if *delta > 180 {
		*delta = 360 - *delta
	} else if *delta < -180 {
		*delta = 360 + *delta
	} else if *delta == 180 || *delta == -180 {
		log.Fatal(fmt.Sprintf("エラー : l.%d : 180度ちょうどの移動は禁止です。", setting.CountLF))
	}
}
