package lexer

import (
	"strconv"

	"../token"
)

func (l *Lexer) ReadChar() {
	if l.readPosition >= l.length {
		l.ch = 0
	} else {
		l.ch = GetRuneAt(l.input, l.readPosition)
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) ReadIdentifier() string {
	position := l.position
	for IsLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) ReadNumber() (string, token.Value) {
	position := l.position
	floatFlag := false
	for IsDigit(l.ch) || IsDot(l.ch) {
		if IsDot(l.ch) {
			floatFlag = true
		}
		l.ReadChar()
	}
	var v token.Value
	if floatFlag {
		f, _ := strconv.ParseFloat(l.input[position:l.position]+"0", 64)
		v.AssignFloat(f)
		return l.input[position:l.position], v
	} else {
		i, _ := strconv.Atoi(l.input[position:l.position])
		v.AssignInt(i)
		return l.input[position:l.position], v
	}
}

/*
func (l *Lexer) ReadNumber() string {
	position := l.position
	for IsDigit(l.ch) {
		l.ReadChar()
	}
	if l.ch == '.' {
		fmt.Println(">>>" + l.input[position:l.position] + "0")
		return l.input[position:l.position] + "0"
	} else {
		return l.input[position:l.position]
	}
}
*/
