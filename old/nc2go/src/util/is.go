package util

import (
	"errors"
	"fmt"
)

func IsAxIs(ch rune) bool {
	return ch == 'X' || ch == 'Y' || ch == 'Z' ||
		ch == 'A' || ch == 'B' || ch == 'C' ||
		ch == 'I' || ch == 'J' || ch == 'K' ||
		ch == 'U' || ch == 'V' || ch == 'W' ||
		ch == 'R'
}

func IsEob(ch rune) bool {
	return ';' == ch
}

func IsLf(ch rune) bool {
	return '\n' == ch
}

func IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func IsDot(ch rune) bool {
	return ch == '.'
}

func IsNewline(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func GetRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func GetRunes(rs *[]rune, i int, num int) (string, error) {
	if i+num < len((*rs)) {
		return string((*rs)[i : i+num]), nil
	} else {
		return "", errors.New("範囲外アクセス")
	}
}

// keywordsにあるか調べる
func IsReserved(identifier string) bool {
	// #等もkeywordsに含まれるが、
	// readlettersでアルファベットと_だけを切り出しているので該当しない。
	if _, ok := keywords[identifier]; ok {
		return true
	} else {
		return false
	}
}

func IsImplemented(id string) bool {
	// 実装したら増やす
	switch id {
	case "%": // % 無視する
		return true
	case "/": // オプショナルスキップはここには来ないはず来たらダメ
		panic(fmt.Sprintf("プログラムエラー : オプショナルスキップ"))
		return true

	// 変数
	// ただ代入すればいいだけなら全部実装している
	// まあまあ想定している奴
	case "O":
		return true
	case "G":
		return true
	case "M":
		return true

	case "X":
		return true
	case "Y":
		return true
	case "Z":
		return true

	case "F": // 送り速度
		return true
	case "S": // 主軸回転速度
		return true
	case "T": // 工具機能
		return true

	case "A": // X軸の回りの角度のディメンション
		return true
	case "B": // Y軸の回りの角度のディメンション
		return true
	case "C": // Z軸の回りの角度のディメンション
		return true
	case "D": // 特殊軸の回りの角度のディメンション又は第三の送り機能 <-謎
		return true
	case "E": // 特殊軸の回りの角度のディメンション又は第二の送り機能 <-謎
		return true

	// あんまり来てほしくないやつ

	case "I": // めんどいやつ
		return true
	case "J": // めんどいやつ
		return true
	case "K": // めんどいやつ
		return true

	case "U": // X軸に平行な第二の運動のディメンション
		return true
	case "V": // Y軸に平行な第二の運動のディメンション
		return true
	case "W": // Z軸に平行な第二の運動のディメンション
		return true

	case "P": // X軸に平行な第三の運動のディメンション <-謎 PはM65のサブルーチンの呼び出し時に使われるので特別扱いでもいいと思う
		return true
	case "Q": // Y軸に平行な第三の運動のディメンション <-謎
		return true
	case "R": // Z軸に平行な第三の運動のディメンション又は第三の運動のディメンション 普通に半径としてつかう
		return true

	// あんまり使われなさそうで想定してないやつ
	case "H": // 特に決まってない
		return true
	case "L": // 特に決まってない
		return true
	case "N": // シーケンス番号
		return true
	default:
		return false
	}
	return false
}
