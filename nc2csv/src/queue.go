package main

import "log"

// G専用Queue
var Gqueue = make([]Value, 0, 100)

// G専用Queue
func EnqueueForG(v interface{}) {
	// vの型判定
	if value, ok := v.(int); ok {
		// int
		Gqueue = append(Gqueue,
			Value{
				bInt: true,
				f:    float64(value),
			},
		)
	} else if value, ok := v.(float64); ok {
		// float
		Gqueue = append(Gqueue,
			Value{
				bInt: false,
				f:    value,
			},
		)
		Memory[key["G"]].AssignFloat(value)
	}

}

func DequeueForG() (Value, error) {
	if len(Gqueue) == 0 {
		return Value{
			bInt: false,
			f:    0,
		}, nil
	}
	ret := Gqueue[0]
	Gqueue = Gqueue[1:]
	return ret, nil
}

func FlushGqueue() {
	if len(Gqueue) != 0 {
		// Gの指定がある
		// Gqueueから取り出して処理する
		for i := len(Gqueue); i > 0; i-- {

			if Setting.IsProhibitAssignAxis {
				// スキップフラグ立ってたら キューを空にして 終わる
				Gqueue = Gqueue[len(Gqueue):]
				break
			}
			v, _ := DequeueForG()
			n := int(v.f)
			switch {
			case n == 0:
				Setting.CutMode = 0
				Assign("F", Setting.FeedG00)
			case n == 1:
				Setting.CutMode = 1
				Assign("F", Setting.FeedG01)
			case n == 2:
				Setting.CutMode = 2
				Assign("F", Setting.FeedG01)
				log.Printf("注意 : l.%d : G02 です。", Setting.CountLF)
			case n == 3:
				Setting.CutMode = 3
				Assign("F", Setting.FeedG01)
				log.Printf("注意 : l.%d : G03 です。", Setting.CountLF)
			case n == 90:
				Setting.IsG90 = true
			case n == 91:
				Setting.IsG90 = false

			// 以下未実装
			case n == 4: // ドウェル
				fallthrough
			case n == 5: // 未指定
				fallthrough
			case n == 6: // 放物線補完
				fallthrough
			case n == 7: // 未指定
				fallthrough
			case n <= 8 && n <= 9: // 加速 減速
				fallthrough
			case n <= 10 && n <= 16: // 未指定
				fallthrough
			case n <= 17 && n <= 19: // 平面指定
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
				// 以降のパラメーターを座標として読み込んまない
				// 座標値への代入を禁止する // これが立っていると改行まで無視する
				Setting.IsProhibitAssignAxis = true
				// 未実装
				log.Printf("注意 : l.%d : G%2d は未実装です。無視します。", Setting.CountLF, n)
			}
		}
	}
}