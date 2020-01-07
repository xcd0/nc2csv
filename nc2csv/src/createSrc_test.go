package main

import "testing"

func TestCreateRunFunction(t *testing.T) {

	const input = // {{{
	`%
O0001(ROBO 4X)
(COORD=CENTER)
(A-90.)
G49
G69
M72
M69
M03S1000(ROT_INITIAL)
G90G00
G90G10L11P#4120R-1.0
X0.Y0.
G43Z50.
G90
G01X-6.Y0.Z50.A-90.F1000
G00Z31.A-90.
G01X-3.F5000
X0.
Z26.F1000
A-90.912F2204
` // }}}

	const out = // {{{
	`func (p *program) run(pc int) {
	if pc == 0 || pc > p.length {
		e := fmt.Sprintf("実行エラー : l.%d : その行は存在しません。エラーです。", pc)
		panic(e)
	}

	NOP := false
	switch pc {
	case 1:
		NOP = true
	case 2:
		Assign(O, 1)
	case 3:
		NOP = true
	case 4:
		NOP = true
	case 5:
		Assign(G, 49)
	case 6:
		Assign(G, 69)
	case 7:
		Assign(M, 72)
	case 8:
		Assign(M, 69)
	case 9:
		Assign(M, 3)
		Assign(S, 1000)
	case 10:
		// 一行に同じGが出たとき
		Assign(G, 90)
		Assign(G, 00)
	case 11:
		Assign(G, 90)
		Assign(G, 10)
		Assign(L, 11)
		Assign(P, Hash[412])
		Assign(R, -1.0)
	case 12:
		Assign(X, 0.0)
		Assign(Y, 0.0)
	case 13:
		Assign(G, 90)
		Assign(G, 90)
	case 14:
		Assign(G, 90)
	case 15:
		Assign(G, 1)
		Assign(X, -6.0)
		Assign(Y, 0.0)
		Assign(Z, 50.0)
		Assign(A, -90.0)
		Assign(F, 10000)
	case 16:
		Assign(G, 0)
		Assign(Z, 31.0)
		Assign(A, -90.0)
	case 17:
		Assign(G, 1)
		Assign(X, -3.0)
		Assign(F, 5000)
	case 18:
		Assign(X, 0.0)
	case 19:
		Assign(Z, 26.1)
		Assign(F, 1000)
	case 20:
		Assign(A, -90.912)
		Assign(F, F2204)
	}
	if !NOP {
		runBlock() //未実装
	}
}` // }}}

}
