package nc2csv

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
	//"gonum.org/v1/gonum/spatial/r3"
)

func forNewLine(rs *[]rune, lines *[]string, in chan string) {

	// フラグをリセットする
	setting.IsOptionalSkip = false
	setting.IsProhibitAssignAxis = false
	// Gのキューを実行する
	flushGqueue()

	// この行を実行した後の状態を出力する
	outputOneLine := axis.genOnelineCsv()

	in <- outputOneLine

	// setting.CountLFは1からだけどlinesは0から
	//log.Printf("l.%v : %v\n", setting.CountLF-1, lines[setting.CountLF-1])

	setting.CountLF++
	if string((*rs)[len(*rs)-1:]) == "\n" && setting.CountLF == len(*lines) {
		// ファイル最後に改行があるファイルとないファイルに対応する
		return
	}
}

var axis Axis
var rowLines []string

// 出力用に座標を持っておくだけ
type Axis struct { // {{{
	X        float64
	Y        float64
	Z        float64
	A        float64
	B        float64
	C        float64
	dX       float64
	dY       float64
	dZ       float64
	dA       float64
	dB       float64
	dC       float64
	vX       float64
	vY       float64
	vZ       float64
	vA       float64
	vB       float64
	vC       float64
	distance float64 // 移動距離
	timeMin  float64 // 移動時間
} // }}}

func lengthVec(v *mat.VecDense) float64 { // ベクトルの大きさ {{{
	vl := v.Len()
	out := 0.0
	for i := 0; i < vl; i++ {
		t := v.AtVec(i)
		out += t * t
	}
	return out
} // }}}

// G00/01用
func shortcutDegree(delta *float64) { // {{{
	if *delta > 180 {
		*delta = 360 - *delta
	} else if *delta < -180 {
		*delta = 360 + *delta
	} else if *delta == 180 || *delta == -180 {
		log.Fatal(fmt.Sprintf("エラー : l.%d : 180度ちょうどの移動は禁止です。", setting.CountLF))
	}
} // }}}

// G02/03用
func calcRotateCenter(a *Axis, 半径 float64) (*mat.VecDense, float64) { // 回転中心座標を求める {{{1
	// 試しに日本語変数を使ってみる

	// 始点s 終点e 中点m 回転中心c
	// 中点を通る法線ベクトルn (G02/G03で変わる)
	始点の位置ベクトル := mat.NewVecDense(3, []float64{a.X, a.Y, a.Z})
	終点の位置ベクトル := mat.NewVecDense(3, []float64{Reference("X").Float(), Reference("Y").Float(), Reference("Z").Float()})

	始点から終点への方向ベクトル := mat.NewVecDense(3, nil)
	始点から終点への方向ベクトル.SubVec(終点の位置ベクトル, 始点の位置ベクトル) // 差分 d = e - s

	中点の位置ベクトル := mat.NewVecDense(3, nil)
	中点の位置ベクトル.AddScaledVec(始点の位置ベクトル, 0.5, 始点から終点への方向ベクトル) // 中点 m = s + 0.5 * ( e - s )

	始点から中点への方向ベクトル := mat.NewVecDense(3, nil)
	始点から中点への方向ベクトル.SubVec(中点の位置ベクトル, 始点の位置ベクトル) // 始点から中点への方向ベクトル mdv = m - s
	始点と中点の距離 := lengthVec(始点から中点への方向ベクトル)

	始点から中点への単位方向ベクトル := mat.NewVecDense(3, nil)
	始点から中点への単位方向ベクトル.ScaleVec(1.0/始点と中点の距離, 始点から中点への方向ベクトル)

	始点から中点への単位方向ベクトルを行列にしたもの := mat.NewDense(3, 3, []float64{
		始点から中点への単位方向ベクトル.AtVec(0), 0, 0, // 行列とベクトルの積の関数をライブラリから見つけられなかったので
		始点から中点への単位方向ベクトル.AtVec(1), 0, 0, // 始点から中点への単位方向ベクトルを
		始点から中点への単位方向ベクトル.AtVec(2), 0, 0, // 行列に突っ込んで回転行列と積を取って回転させる
	})

	/* // 座標系に合わせた回転行列{{{2
	# 回転行列 3次元
	  XY平面 G17        ZX平面 G18          YZ平面 G19
	| cos -sin 0  |   |  cos  0  sin |    |  1   0   0   |
	| sin cos  0  |   |  0    1   0  |    |  0  cos -sin |
	|  0   0   1  |   | -sin  0  cos |    |  0  sin cos  |

	## G02 (時計回りなので-90°)

	  XY平面 G17        ZX平面 G18         YZ平面 G19
	|  0   1   0  |    |  0   0  -1  |    |  1   0   0  |
	| -1   0   0  |    |  0   1   0  |    |  0   0   1  |
	|  0   0   1  |    |  1   0   0  |    |  0  -1   0  |

	## G03 (反時計回りなので90°)

	  XY平面 G17        ZX平面 G18         YZ平面 G19
	|  0  -1   0  |    |  0   0   1  |    |  1   0   0  |
	|  1   0   0  |    |  0   1   0  |    |  0   0  -1  |
	|  0   0   1  |    | -1   0   0  |    |  0   1   0  |
	*/
	var 回転行列 *mat.Dense                 // 回転行列
	if setting.PlaneDesignation == 17 { // 17 XY平面指定
		if setting.CutMode == 2 {
			回転行列 = mat.NewDense(3, 3, []float64{
				0, 1, 0,
				-1, 0, 0,
				0, 0, 1})
		} else if setting.CutMode == 3 {
			回転行列 = mat.NewDense(3, 3, []float64{
				0, -1, 0,
				1, 0, 0,
				0, 0, 1})
		}
	} else if setting.PlaneDesignation == 18 { // 18 ZX平面指定
		if setting.CutMode == 2 {
			回転行列 = mat.NewDense(3, 3, []float64{
				0, 0, -1,
				0, 1, 0,
				1, 0, 0})
		} else if setting.CutMode == 3 {
			回転行列 = mat.NewDense(3, 3, []float64{
				0, 0, 1,
				0, 1, 0,
				-1, 0, 0})
		}
	} else if setting.PlaneDesignation == 19 { // 19 YZ平面指定
		if setting.CutMode == 2 {
			回転行列 = mat.NewDense(3, 3, []float64{
				1, 0, 0,
				0, 0, 1,
				0, -1, 0})
		} else if setting.CutMode == 3 {
			回転行列 = mat.NewDense(3, 3, []float64{
				1, 0, 0,
				0, 0, -1,
				0, 1, 0})
		}
	} // }}}2

	始点から中点への単位方向ベクトルを行列にしたもの.Mul(回転行列, 始点から中点への単位方向ベクトルを行列にしたもの)
	単位法線ベクトル := mat.NewVecDense(3, nil)
	// mat.Vecterからmat.VecDenseへのキャストがわからなくてとりあえず1掛けた
	単位法線ベクトル.ScaleVec(1, 始点から中点への単位方向ベクトルを行列にしたもの.RowView(0)) // 単位法線ベクトル 3x3行列から1列目だけ取り出す
	// 中点から回転中心までの距離 d = sqrt( r^2 - ( m - s )^2 )
	中点から回転中心までの距離 := math.Sqrt(math.Pow(半径, 2) - math.Pow(始点と中点の距離, 2))
	中点から回転中心までの方向ベクトル := mat.NewVecDense(3, nil)
	中点から回転中心までの方向ベクトル.ScaleVec(中点から回転中心までの距離, 単位法線ベクトル)

	回転中心の位置ベクトル := mat.NewVecDense(3, nil)
	回転中心の位置ベクトル.AddVec(中点の位置ベクトル, 中点から回転中心までの方向ベクトル)

	回転角rad := math.Asin(始点と中点の距離 / 半径)
	//回転角度 := math.Asin(始点と中点の距離/半径) * 360.0 / math.Pi   // asin(l/r) * 360 / 2pi * 2

	// 日本語ながい..

	return 回転中心の位置ベクトル, 回転角rad
} // }}}1

func (a *Axis) calcDistance() { // {{{
	// a.distanceを更新する
	if setting.CutMode == 2 || setting.CutMode == 3 {
		// G02/03

		// 変わるので
		a.vX, a.vY, a.vZ, a.vA, a.vB, a.vC = 0, 0, 0, 0, 0, 0
		// a.dが0の時calcTime()でa.vが更新されない → a.vが0のままになる
		a.dX, a.dY, a.dZ, a.dA, a.dB, a.dC = 0, 0, 0, 0, 0, 0

		r := Reference("R").Float()
		_, d := calcRotateCenter(a, math.Abs(r))
		if r < 0 {
			d = 360.0 - d // 回転角 rが負の時は大きく回る
		}
		a.distance = r * d // 円弧長 = 半径 * なす角
	} else {
		// G00/01

		// 近回りさせる
		shortcutDegree(&a.dA)
		shortcutDegree(&a.dB)
		shortcutDegree(&a.dC)
		a.distance = math.Sqrt(
			math.Pow(a.dX, 2) + math.Pow(a.dY, 2) + math.Pow(a.dZ, 2) +
				math.Pow(a.dA, 2) + math.Pow(a.dB, 2) + math.Pow(a.dC, 2))
	}
} // }}}

func (a *Axis) calcTime() { // {{{
	// 前チェック
	f := Reference("F").Float()
	if f == 0 {
		// 送り速度が 0
		log.Fatal(fmt.Sprintf("実行時エラー : l.%d : 送り速度が0です。", setting.CountLF))
	}
	// a.timeMinを更新する
	// 序にa.vも更新する
	a.timeMin = 0 // 一旦リセット
	if a.distance != 0 && f != 0 {
		// Fは分単位
		a.timeMin = a.distance / f
		// XYZABC の各軸速度
		a.vX, a.vY, a.vZ, a.vA, a.vB, a.vC = 0, 0, 0, 0, 0, 0
		if a.dX != 0 {
			a.vX = a.dX / a.timeMin
		}
		if a.dY != 0 {
			a.vY = a.dY / a.timeMin
		}
		if a.dZ != 0 {
			a.vZ = a.dZ / a.timeMin
		}
		if a.dA != 0 {
			a.vA = a.dA / a.timeMin
		}
		if a.dB != 0 {
			a.vB = a.dB / a.timeMin
		}
		if a.dC != 0 {
			a.vC = a.dC / a.timeMin
		}
	}
	// 合計時間も更新する
	setting.CumulativeTime += a.timeMin
} // }}}

func (a *Axis) genOnelineCsv() string { // {{{

	a.calcDistance() // 移動距離を計算
	a.calcTime()     // 移動時間を計算
	a.dX, a.dY, a.dZ = (Reference("X").Float() - a.X), (Reference("Y").Float() - a.Y), (Reference("Z").Float() - a.Z)
	// 180度より大きい移動は近回り。180度ピッタリの移動はエラーにする。
	// 回転する軸は0~360に丸めているので。 境界で処理が必要。 例えばA359.0 A0はdA==1になる。
	a.dA, a.dB, a.dC = (Reference("A").Float() - a.A), (Reference("B").Float() - a.B), (Reference("C").Float() - a.C)
	// 出力
	out := fmt.Sprintf("%d,\"%v\",%v,%v,%v,%v,%v,%v,%v,%v,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.10f,%.10f",
		setting.CountLF, rowLines[setting.CountLF-1],
		Reference("X").String(), Reference("Y").String(), Reference("Z").String(),
		Reference("A").String(), Reference("B").String(), Reference("C").String(),
		Reference("R").String(), Reference("F").String(),
		a.distance,
		a.vX, a.vY, a.vZ, a.vA, a.vB, a.vC,
		a.timeMin, setting.CumulativeTime,
	)

	// 保存
	a.X, a.Y, a.Z, a.A, a.B, a.C = Reference("X").Float(), Reference("Y").Float(), Reference("Z").Float(), Reference("A").Float(), Reference("B").Float(), Reference("C").Float()
	return out
} // }}}
