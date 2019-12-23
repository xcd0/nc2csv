package lexer

import (
	"fmt"
	"strconv"

	"../token"
)

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
func (l *Lexer) ReadNumber() (string, token.Num) {
	position := l.position
	floatFlag := false
	for IsDigit(l.ch) || IsDot(l.ch) {
		if IsDot(l.ch) {
			floatFlag = true
		}
		l.ReadChar()
	}
	var n token.Num
	if floatFlag {
		return l.input[position:l.position], n.AssignFloat(strcov.ParseFloat(l.input[position:l.position], 64))
	} else {
		return l.input[position:l.position], n.AssignInt(strconv.Atoi(l.input[position:l.position]))
	}
}
