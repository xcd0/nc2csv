package nc2csv

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// G専用Queueに保存する int, float64, string対応
func enqueueForG(v interface{}) { // {{{
	// vの型判定
	if vl, ok := v.(int); ok {
		// int
		gqueue = append(gqueue,
			value{
				bInt: true,
				f:    float64(vl),
			},
		)
	} else if vl, ok := v.(float64); ok {
		// float
		gqueue = append(gqueue,
			value{
				bInt: false,
				f:    vl,
			},
		)
		memory[key["G"]].assignFloat(vl)
	} else if vl, ok := v.(string); ok {
		// 小数点があるかどうか調べる
		if strings.Contains(vl, ".") {
			// float
			v, err := strconv.ParseFloat(vl, 64)
			if err != nil {
				fmt.Printf("\n")
				log.Fatal(fmt.Sprintf("エラー : バグ : 文字列からfloat64への変換に失敗しました。%%v %v : err %v", v, err))
			}
			gqueue = append(gqueue,
				value{
					bInt: false,
					f:    v,
				},
			)
		} else {
			// int
			v, err := strconv.Atoi(vl)
			if err != nil {
				fmt.Printf("\n")
				log.Fatal(fmt.Sprintf("エラー : バグ : 文字列からintへの変換に失敗しました。%%v %v : err %v", v, err))
			}
			gqueue = append(gqueue,
				value{
					bInt: true,
					f:    float64(v),
				},
			)
		}
	}

} // }}}

// G専用Queueから値を取り出す
func dequeueForG() (value, error) { // {{{
	if len(gqueue) == 0 {
		return value{
			bInt: false,
			f:    0,
		}, nil
	}
	ret := gqueue[0]
	gqueue = gqueue[1:]
	return ret, nil
} // }}}

// G専用Queueの中身をすべて処理する
func flushGQueue() { // {{{1
	if len(gqueue) == 0 {
		return
	}
	// Gの指定がある
	// gqueueから取り出して処理する
	for i := len(gqueue); i > 0; i-- {

		if setting.IsProhibitAssignAxis {
			// スキップフラグ立ってたら キューを空にして 終わる
			gqueue = gqueue[len(gqueue):]
			break
		}
		v, _ := dequeueForG()
		if v.bInt == false { // {{{2
			// 小数
			switch {
			// 見つけたらエラー扱い(未対応で処理停止)を出してほしいワード
			// G43.4, G43.5 (XYZABCの扱いが大きく変わるため、来たら未対応を表明したい)
			case v.f == 43.4:
				fmt.Println("")
				log.Printf("エラー : l.%d : G43.4 は未対応です。行を無視します。", setting.CountLF)
				setting.IsProhibitAssignAxis = true
				continue
			case v.f == 43.5:
				fmt.Println("")
				log.Printf("エラー : l.%d : G43.5 は未対応です。行を無視します。", setting.CountLF)
				setting.IsProhibitAssignAxis = true
				continue
			default:
				// 未実装
				fmt.Println("")
				log.Printf("注意 : l.%d : G%2d.% .2f は未対応です。行を無視します。", setting.CountLF, int(v.f), v.f-float64(int(v.f)))
				// 以降のパラメーターを座標として読み込んまない
				// 座標値への代入を禁止する // これが立っていると改行まで無視する
				setting.IsProhibitAssignAxis = true
				continue
			}
			return
		} // }}}2

		// 以下整数
		n := int(v.f)
		switch { // {{{2
		case n == 0:
			setting.CutMode = 0
			// 一旦現在の送り速度を保存
			setting.FeedG01 = Reference("F").Float()
			Assign("F", setting.FeedG00)
		case n == 1:
			Assign("F", setting.FeedG01) // 直前がG00だったときの為に復元する
			setting.CutMode = 1
		case n == 2:
			Assign("F", setting.FeedG01) // 直前がG00だったときの為に復元する
			setting.CutMode = 2
			fmt.Println("")
			log.Printf("注意 : l.%d : G02 です。", setting.CountLF)
		case n == 3:
			Assign("F", setting.FeedG01) // 直前がG00だったときの為に復元する
			setting.CutMode = 3
			fmt.Println("")
			log.Printf("注意 : l.%d : G03 です。", setting.CountLF)
		case n == 17: // XY平面指定 G02/03を実装したら使われる
			setting.PlaneDesignation = 17
		case n == 18: // ZX平面指定 G02/03を実装したら使われる
			setting.PlaneDesignation = 18
		case n == 19: // YZ平面指定 G02/03を実装したら使われる
			setting.PlaneDesignation = 19
		case n == 90: // アブソリュート指令
			setting.IsG90 = true
		case n == 91: // インクリメンタル指令
			setting.IsG90 = false

		case n == 4: // ドウェル
			// スキップ
			fmt.Println("")
			log.Printf("注意 : l.%d : G04 ドウェルです。行を無視します。", setting.CountLF)
			setting.IsProhibitAssignAxis = true
			continue
		case n == 5: // AI輪郭制御モード
			// スキップ
			fmt.Println("")
			log.Printf("注意 : l.%d : G05 です。行を無視します。", setting.CountLF)
			setting.IsProhibitAssignAxis = true
			continue

		// 以下未実装
		case n == 6: // 放物線補完
			fallthrough
		case n == 7: // 未指定
			fallthrough
		case n <= 8 && n <= 9: // 加速 減速
			fallthrough
		case n <= 10 && n <= 16: // 未指定
			fallthrough
		case n <= 41 && n <= 42: // 工具径補正 左 右
			fallthrough
		case n <= 43 && n <= 44: // 工具オフセット オフセットのキャンセル
			fallthrough
		case n <= 45 && n <= 52: // 工具オフセット いろいろ
			fallthrough
		case n <= 53 && n <= 59: // シフト
			fallthrough
		case n <= 60 && n <= 61: // 正確な位置決め 精密 普通
			fallthrough
		case n == 62: // 迅速位置決め
			fallthrough
		case n <= 63 && n <= 79: // 未指定
			fallthrough
		case n == 80: // 固定サイクルのキャンセル
			fallthrough
		case n <= 81 && n <= 89: // 固定サイクル
			fallthrough
		case n == 92: // 座標系指定
			fallthrough
		case n == 93: // 時間の逆数で送りを指定
			fallthrough
		case n == 94: // 毎分当たりmm送り
			fallthrough
		case n == 95: // 主軸一回店当たり送り
			fallthrough
		case n <= 96 && n <= 97: // 定切削速度 定切削速度のキャンセル
			fallthrough
		case n <= 98 && n <= 99: // 未指定
			fallthrough
		default:
			// 未実装
			fmt.Println("")
			log.Printf("注意 : l.%d : G%2d は未対応です。行を無視します。", setting.CountLF, n)
			// 以降のパラメーターを座標として読み込んまない
			// 座標値への代入を禁止する // これが立っていると改行まで無視する
			setting.IsProhibitAssignAxis = true
			continue
		} // }}}2
	}
} // }}}1
