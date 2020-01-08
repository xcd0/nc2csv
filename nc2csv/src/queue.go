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
			v, _ := DequeueForG()
			n := int(v.f)
			switch n {
			case 0:
				Setting.CutMode = 0
				Assign("F", Setting.FeedG00)
			case 1:
				Setting.CutMode = 1
				Assign("F", Setting.FeedG01)
			case 2:
				Setting.CutMode = 2
				Assign("F", Setting.FeedG01)
				log.Printf("注意 : l.%d : G02 です。", Setting.CountLF)
			case 3:
				Setting.CutMode = 3
				Assign("F", Setting.FeedG01)
				log.Printf("注意 : l.%d : G03 です。", Setting.CountLF)
			case 90:
				Setting.IsG90 = true
			case 91:
				Setting.IsG90 = false

			// 以下未実装
			case 4: // ドウェル
				fallthrough
			case 5: // 未指定
				fallthrough
			case 6: // 放物線補完
				fallthrough
			case 7: // 未指定
				fallthrough
			case 8, 9: // 加速 減速
				fallthrough
			case 10, 11, 12, 13, 14, 15, 16: // 未指定
				fallthrough
			case 17, 18, 19: // 平面指定
				fallthrough
			case 41, 42: // 工具径補正 左 右
				fallthrough
			case 43, 44: // 工具オフセット オフセットのキャンセル
				fallthrough
			case 45, 46, 47, 48, 49, 50, 51, 52: // 工具オフセット いろいろ
				fallthrough
			case 53, 54, 55, 56, 57, 58, 59: // シフト
				fallthrough
			case 60, 61: // 正確な位置決め 精密 普通
				fallthrough
			case 62: // 迅速位置決め
				fallthrough
			case 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79: // 未指定
				fallthrough
			case 80: // 固定サイクルのキャンセル
				fallthrough
			case 81, 82, 83, 84, 85, 86, 87, 88, 89: // 固定サイクル
				fallthrough
			case 92: // 座標系指定
				fallthrough
			case 93: // 時間の逆数で送りを指定
				fallthrough
			case 94: // 毎分当たりmm送り
				fallthrough
			case 95: // 主軸一回店当たり送り
				fallthrough
			case 96, 97: // 定切削速度 定切削速度のキャンセル
				fallthrough
			case 98, 99: // 未指定
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
