package util

import "fmt"

var Hash = make([]Value, 10000)

type Value struct {
	bInt bool
	i    int
	f    float64
}

func (v *Value) IsInt() bool {
	return v.bInt
}

func (v *Value) AssignInt(i int) {
	v.bInt = true
	v.i = i
	v.f = 0
}

func (v *Value) AssignFloat(f float64) {
	v.bInt = false
	v.i = 0
	v.f = f
}

func (v *Value) String() string {
	if v.bInt {
		return string(v.i)
	} else {
		return fmt.Sprintf("%f", v.f)
	}
}
