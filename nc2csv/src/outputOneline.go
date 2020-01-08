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

func (a *Axis) outputOneline() string {

	a.dX = (Reference("X").Float() - a.X)
	a.dY = (Reference("Y").Float() - a.Y)
	a.dZ = (Reference("Z").Float() - a.Z)
	a.dA = (Reference("A").Float() - a.A)
	a.dB = (Reference("B").Float() - a.B)
	a.dC = (Reference("C").Float() - a.C)
	// 移動距離
	var dEuclideanDistance float64
	dEuclideanDistance = math.Sqrt(math.Pow(a.dX, 2) + math.Pow(a.dY, 2) + math.Pow(a.dZ, 2) + math.Pow(a.dA, 2) + math.Pow(a.dB, 2) + math.Pow(a.dC, 2))
	// 移動時間
	var dTimeMin float64
	f := Reference("F").Float()
	if f == 0 {
		// 送り速度が 0
		log.Fatal(fmt.Sprintf("実行時エラー : l.%d : 送り速度が0です。", Setting.CountLF))
	}
	if dEuclideanDistance != 0 && f != 0 {
		// Fは分単位
		dTimeMin = dEuclideanDistance / f
	}

	// 元ncプログラムの行、NC、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
	// 元ncプログラムの行
	out := fmt.Sprintf("%d", Setting.CountLF)
	// その行のncプログラム
	out += "," + rowLines[Setting.CountLF-1]
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
}
