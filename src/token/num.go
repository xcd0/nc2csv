package token

type Num struct {
	Int        bool
	ValueInt   int
	ValueFolat float64
}

func (n *Num) AssignInt(i int) {
	n.Int = true
	n.ValueInt = i
	n.ValueFloat = 0
}

func (n *Num) AssignFloat(f float64) {
	n.Int = false
	n.ValueFloat = f
	n.ValueInt = 0
}

func (n *Num) String() string {
	if n.Int {
		return string(ValueInt)
	} else {
		return string(ValueFloat)
	}
}
