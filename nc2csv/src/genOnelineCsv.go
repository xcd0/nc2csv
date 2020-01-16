package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
	//"gonum.org/v1/gonum/spatial/r3"
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

func (m *mat.Dense) MulMatVec(im, iv *mat.Dense) {
}

func (a *Axis) genOnelineCsv() (string, float64) {
	// 移動距離
	var dDistance float64
	// 移動時間
	var dTimeMin float64

	if setting.CutMode == 2 || setting.CutMode == 3 {
		// 回転中心座標を求める {{{
		pre := mat.NewDense(3, 1, []float64{a.X, a.Y, a.Z})
		now := mat.NewDense(3, 1, []float64{Reference("X").Float(), Reference("Y").Float(), Reference("Z").Float()})
		d := mat.NewDense(3, 1, nil)
		d.Sub(now, pre) // 差分
		mid := mat.NewDense(3, 1, nil)
		mid.Scale(0.5, now)
		mid.Add(pre)                    // 中点
		norm := mat.NewDense(3, 1, nil) // 法線ベクトル norm := mat.Dense

		rot := mat.Dense
		if setting.PlaneDesignation == 17 {
			if setting.CutMode == 2 {
				rot = mat.NewDense(3, 3, [][]float64{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}})
			} else if setting.CutMode == 3 {
				rot = mat.NewDense(3, 3, [][]float64{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}})
			}
		} else if setting.PlaneDesignation == 18 {
			if setting.CutMode == 2 {
				rot = mat.NewDense(3, 3, [][]float64{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}})
			} else if setting.CutMode == 3 {
				rot = mat.NewDense(3, 3, [][]float64{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}})
			}
		} else if setting.PlaneDesignation == 19 {
			if setting.CutMode == 2 {
				rot = mat.NewDense(3, 3, [][]float64{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}})
			} else if setting.CutMode == 3 {
				rot = mat.NewDense(3, 3, [][]float64{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}})
			}
		}
		norm.Mul(rot, d)
		// }}}
		if setting.CutMode == 2 {
			// G02
			r := Reference("R").Float()
			r := Reference("R").Float()

			//dDistance =
			a.dX, a.dY, a.dZ, a.dA, a.dB, a.dC = 0, 0, 0, 0, 0, 0
			a.vX, a.vY, a.vZ, a.vA, a.vB, a.vC = 0, 0, 0, 0, 0, 0
		} else if setting.CutMode == 3 {
			// G03
			r := Reference("R").Float()
			a.dX, a.dY, a.dZ, a.dA, a.dB, a.dC = 0, 0, 0, 0, 0, 0
			a.vX, a.vY, a.vZ, a.vA, a.vB, a.vC = 0, 0, 0, 0, 0, 0
		}
	} else {
		// G00/01
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

		shortcutDegree(&a.dA)
		shortcutDegree(&a.dB)
		shortcutDegree(&a.dC)
		dDistance = math.Sqrt(math.Pow(a.dX, 2) + math.Pow(a.dY, 2) + math.Pow(a.dZ, 2) + math.Pow(a.dA, 2) + math.Pow(a.dB, 2) + math.Pow(a.dC, 2))
		f := Reference("F").Float()
		if f == 0 {
			// 送り速度が 0
			log.Fatal(fmt.Sprintf("実行時エラー : l.%d : 送り速度が0です。", setting.CountLF))
		}
		if dDistance != 0 && f != 0 {
			// Fは分単位
			dTimeMin = dDistance / f
		}
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
	out += "," + fmt.Sprintf("%.6f", dDistance)
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
