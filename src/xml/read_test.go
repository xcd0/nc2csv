package main

import (
	"fmt"
	"testing"
)

func TestReadCSV(t *testing.T) {
	rowInput := `x,y,z,a,b,c
0,0,0,0,0,1
0,0,100,0,0,1
10,10,100,0,0,e`

	out := ReadCsv(rowInput)

	expect := "[[x y z a b c] [0 0 0 0 0 1] [0 0 100 0 0 1] [10 10 100 0 0 e]]"

	output := fmt.Sprintf("%v", out)

	if output != expect {
		t.Fatalf("tests[ %v ] - tokentype wrong. expected=%q, got=%q", "ReadCSV", expect, output)
	}
}
