package util

import (
	"errors"
	"fmt"
	"log"
)

func IsAxIs(r rune) bool {
	return r == 'X' || r == 'Y' || r == 'Z' ||
		r == 'A' || r == 'B' || r == 'C' ||
		r == 'I' || r == 'J' || r == 'K' ||
		r == 'U' || r == 'V' || r == 'W' ||
		r == 'R'
}

func IsEob(r rune) bool {
	return ';' == r
}

func IsLf(r rune) bool {
	return '\n' == r
}

func IsLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func IsPM(r rune) bool {
	return r == '+' || r == '-'
}

func IsHash(r rune) bool {
	return r <= '#'
}

func IsDot(r rune) bool {
	return r == '.'
}

func IsNewline(r rune) bool {
	return r == '\n' || r == '\r'
}

func IsWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
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
	return IsImplementedCharactor(id) || IsImplementedWord(id)
}

func IsImplementedWord(id string) bool { // {{{
	// 実装したら増やす
	switch id {
	case "EOF":
		// この行で正常終了させる。
		return true

	// 未実装
	case "GOTO":
		log.Printf(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "IF":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "THEN":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "WHILE":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "DO":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "EQ":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "NE":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "LT":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "GT":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "LE":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "GE":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "AND":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "OR":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	case "XOR":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return false
	default:
		return false
	}
	return false
} // }}}

func IsImplementedCharactor(id string) bool { // {{{
	// 実装したら増やす
	switch id {
	case "%": // % 無視する
		return true

	case "[":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return true
	case "]":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return true
	case "+":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return true
	case "-":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return true
	case "*":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return true
	case "/":
		// オプショナルスキップはここには来ないはず来たらダメ
		//panic(fmt.Sprintf("プログラムエラー : オプショナルスキップの処理にバグがあります。"))

		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
		return true
	case "=":
		log.Fatal(fmt.Sprintf("警告 : 実装予定予約語 %v です。未実装です。", id))
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
} // }}}
