package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/xcd0/go-nkf"
)

func CreateSrc(rowInput string) string {

	output := makeRunFunction(input)

	return output
}

func makeRunFunction(input string) string {

	lines := strings.Split(input, "\n")

	output := `func (p *program) run(pc int) {
	if pc == 0 || pc > p.length {
		e := fmt.Sprintf("実行エラー : l.%d : その行は存在しません。エラーです。", pc)
		panic(e)
	}

	NOP := false
	switch pc { `

	flagCommentStart := false
	// 一行読み取る
	for _, line := range lines {
		// 行ごとにcase文
		output += fmt.Sprintf("\tcase %3v:\n", i)
		// コメントの処理
		rs := []rune(line)
		if rs[0] == '(' {
			flagCommentStart = false
			for _, r := range rs {
				// ()はコメント 複数行コメントにも対応
				if r == '(' {
					flagCommentStart = true
					continue
				}
				if flagCommentStart {
					if r == ')' {
						flagCommentStart = false
					}
					continue
				}
			}
			output += string(r)
		}

		output += fmt.Sprintf("\t\t%s\n", "")
	}

	rs := []rune(input)
	l := len(rs)
	ro := make([]rune, 0, l*2)
	var r rune
	var t rune
	countLF := 1
	for i := 0; i < l; i++ {
	}

	// コメント()を削除
	input := util.DeleteComment(spacedInput)
	// 空行の削除 // GOTOを実装するときに空白を削除すると行がずれてまずい //formedInput = util.DeleteDoubleNewline(input)
	// 変数代入の変換
	convertedAssign := convertAssign(input)
}

func convertAssign(input string) string { // {{{
	rs := []rune(input)
	l := len(rs)
	ro := make([]rune, 0, l*2)
	var r rune
	var t rune
	countLF := 1
	for i := 0; i < l; i++ {
		r = rs[i]
		if IsLetter(r) {
			// 予約語の判定

			// rs[i]がアルファベット
			// 後ろを見てアルファベットが続く間切り出す
			identifier := ReadLetters(rs, i)
			if IsReserved(identifier) && IsImplemented(identifier) {
				// 実装済み予約語
				// ここに実装

				if identifier == "%" { // % 無視する
					continue
				}

			} else if IsReserved(identifier) && !IsImplemented(identifier) {
				// 未実装予約語
				e := fmt.Sprintf("書式エラー : l.%d :予約語 %v は未実装です。", countLF, id)
				panic(e)
			}
			// 予約語でないものは変数
			// 変数は複数文字でないので 複数文字はエラー
			if len(identifier) != 1 {
				e := fmt.Sprintf("書式エラー : l.%d :予約語でない文字列\n%v はエラーです。", countLF, id)
				panic(e) // 予約語でない複数文字列は書式エラー
			}
			// 変数代入
			// 後ろの数値を読み取って文字列に追加
			tmp := fmt.Sprintf("Hash[%v] = %v;", identifier, ReadNumbers(rs, i))
			ro = append(ro, []rune(tmp)...)
			continue
		} else if IsLF(r) {
			countLF++
			t = ';'
			ro = append(ro, t)
			t = '\n'
		}
		ro = append(ro, t)
	}

	return string(ro)
}

// }}}

// read {{{
func ReadLetters(rs *[]rune, i int) string {
	pre := i
	post := i
	for IsLetter(rs[post]) {
		post++
	}
	return string(rs[pre : post+1])
}
func ReadNumbers(rs *[]rune, i int) string {
	pre := i
	post := i
	for IsDigit(rs[post]) || IsDot(rs[post]) {
		post++
	}
	return string(rs[pre : post+1])
}
func ReadText(path string) string {

	// 与えられたパスの文字列について
	// そのパスにあるファイルをテキストファイルとして読み込む

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル%vが読み込めません\n", path)
		log.Println(err)
		panic(err)
		return ""
	}
	// ファイルの文字コード変換
	charset, err := nkf.CharDet(b)
	if err != nil {
		/*
			fmt.Fprintf(os.Stderr, "文字コード変換に失敗しました\nutf8を使用してください\n")
			log.Println(err)
			panic(err)
			return ""
		*/
		return ConvertNewline(string(b), "\n")
	}

	str, err := nkf.ToUtf8(string(b), charset)

	str = ConvertNewline(str, "\n")

	return str
}

// }}}

// is {{{
func IsAxIs(ch rune) bool {
	return ch == 'x' || ch == 'y' || ch == 'z' ||
		ch == 'a' || ch == 'b' || ch == 'c' ||
		ch == 'i' || ch == 'j' || ch == 'k' ||
		ch == 'u' || ch == 'v' || ch == 'w' ||
		ch == 'r'
}

func IsEob(ch rune) bool {
	return ';' == ch
}

func IsLf(ch rune) bool {
	return '\n' == ch
}

func IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'a' <= ch && ch <= 'z' || ch == '_'
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
	r = rs[i]
	if i+num < len(rs) {
		return string(*rs[i : i+num]), nil
	} else {
		return "", errors.New("範囲外アクセス")
	}
}

// keywordsにあるか調べる
func IsReserved(identifier string) bool {
	// #等もkeywordsに含まれるが、
	// readlettersでアルファベットと_だけを切り出しているので該当しない。
	if token, ok := keywords[identifier]; ok {
		return true
	} else {
		return false
	}
}

func IsImplemented(id string) bool {
	// 実装したら増やす
	if id == "%" {
		return true
	} else {
		return false
	}
}

// }}}

const ( // {{{
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // 関数名や変数名 add, foobar, x, y, ...
	INT        = "INT"        // 1343456
	FLOAT      = "FLOAT"      // 003.1415926

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	ASSIGNEQ = "="      // 変数#への代入で変数と値の区切りに使われる
	ASSIGN   = "ASSIGN" // 代入で汎用的に使われる

	ARRAY = "#"

	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	DIV = "DIV"

	EQ = "EQ" // ==
	NE = "NE" // !=
	LT = "LT" // <
	LE = "LE" // <=
	GT = "GT" // >
	GE = "GE" // >=

	AND = "AND" // &&
	OR  = "OR"  // ||
	XOR = "XOR" // &^

	GOTO  = "GOTO"
	IF    = "IF"
	THEN  = "THEN"
	WHILE = "WHILE"
	END   = "END"

	COMMENTSTART = "COMMENTSTART" // '('
	COMMENTEND   = "COMMENTEND"   // ')'
	NCEOF        = "NCEOF"        // '%'
	EOB          = "EOB"          // '\n'
	VARIABLE     = "VARIABLE"     // '#'

	/*
		AXIS          = "AXIS"
		PREPARATION   = "PREPARATION"
		MISCELLANEOUS = "MISCELLANEOUS"
		FEED          = "FEED"
		SPINDLE       = "SPINDLE"
		TOOL          = "TOOL"
		ONUM          = "ONUM"
	*/
	SKIP = "SKIP"

	A = "A"
	B = "B"
	C = "C"
	D = "D"
	E = "E"
	F = "F"
	G = "G"
	H = "H"
	I = "I"
	J = "J"
	K = "K"
	L = "L"
	M = "M"
	N = "N"
	O = "O"
	P = "P"
	Q = "Q"
	R = "R"
	S = "S"
	T = "T"
	U = "U"
	V = "V"
	W = "W"
	X = "X"
	Y = "Y"
	Z = "Z"
)

// }}}

// 実装済み予約語 {{{
var keywords = map[string]string{

	"%": NCEOF,

	// 未実装予約語

	"=":    ASSIGN,
	"+":    PLUS,
	"-":    MINUS,
	"*":    ASTERISK,
	"/":    SLASH,
	"#":    ARRAY,
	"(":    COMMENTSTART,
	")":    COMMENTEND,
	"EQ":   EQ,
	"NE":   NE,
	"LT":   LT,
	"LE":   LE,
	"GT":   GT,
	"GE":   GE,
	"OR":   OR,
	"ADD":  ADD,
	"SUB":  SUB,
	"MUL":  MUL,
	"DIV":  DIV,
	"XOR":  XOR,
	"AND":  AND,
	"COS":  COS,
	"SIN":  SIN,
	"TAN":  TAN,
	"ACOS": ACOS,
	"ASIN": ASIN,
	"ATAN": ATAN,

	"SKIP":  SKIP,
	"GOTO":  GOTO,
	"IF":    IF,
	"WHILE": WHILE,
	"END":   END,

	"EOB": EOB,

	/*
		"AXIS":          AXIS,
		"PREPARATION":   PREPARATION,
		"MISCELLANEOUS": MISCELLANEOUS,
		"FEED":          FEED,
		"SPINDLE":       SPINDLE,
		"TOOL":          TOOL,
		"ONUM":          ONUM,
	*/
}

// }}}
