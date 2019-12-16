package main

import "fmt"

type NcState struct {
	x   float64
	y   float64
	z   float64
	a   float64
	b   float64
	c   float64
	f   float64
	dim int // 90 or 91
}

func (ns *NcState) show() {
	fmt.Printf("(%v,%v,%v,%v,%v,%v)", x, y, z, a, b, c, ", f=%v,dim=%v", f, dim)
}

func ncLineToBlocks() {
}
