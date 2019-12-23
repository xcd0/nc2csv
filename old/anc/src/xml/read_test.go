package main

import (
	"fmt"
	"testing"
)

func TestReadCsv(t *testing.T) {
	rowInput := `x,y,z,a,b,c
0,0,0,0,0,1
0,0,100,0,0,1
10,10,100,0,0,e`

	out := ReadCsv(rowInput)

	expect := "[[x y z a b c] [0 0 0 0 0 1] [0 0 100 0 0 1] [10 10 100 0 0 e]]"

	output := fmt.Sprintf("%v", out)

	if output != expect {
		t.Fatalf("tests[ %v ] - tokentype wrong. expected=%q, got=%q", "ReadCsv", expect, output)
	}
}

func TestReadXml(t *testing.T) {
	rowInput := `<?xml version="1.0"?>
<top>
<a oooo="oooooooo">
<b>b1</b>
<b>b2</b>
<c o1="o1" o2="o2" num="42">c1</c>
</a>
<d>
<e>eeeee</e>
</d>
</top>
`

	expect := `aa`

	out := ReadXml(rowInput)

	output := fmt.Sprintf("%v", out)

	fmt.Println(out)
	fmt.Println("--")
	fmt.Println(output)
	fmt.Println("--")

	if output != expect {
		t.Fatalf("tests[ %v ] - wrong. expected=%q, got=%q", "ReadXml", expect, output)
	}
}
