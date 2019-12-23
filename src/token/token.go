package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // 関数名や変数名 add, foobar, x, y, ...
	INT        = "INT"        // 1343456
	FLOAT      = "FLOAT"      // 003.1415926

type Num struct {
	Int        bool
	ValueInt   int
	ValueFolat float64
}

func (n *Num) Assign(i int) {
	n.Int = true
	n.ValueInt = i
	n.ValueFloat = 0
}

func (n *Num) Assign(f float64) {
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

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	ASSIGN   = "=" // 変数#への代入で使われる

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
		SKIP          = "SKIP"
	*/

)

var keywords = map[string]TokenType{
	"=":   ASSIGN,
	"+":   PLUS,
	"-":   MINUS,
	"*":   ASTERISK,
	"/":   SLASH,
	"#":   ARRAY,
	"(":   COMMENTSTART,
	")":   COMMENTEND,
	"%":   NCEOF,
	"#":   VARIABLE,
	"EQ":  EQ,
	"NE":  NE,
	"LT":  LT,
	"LE":  LE,
	"GT":  GT,
	"GE":  GE,
	"OR":  OR,
	"ADD": ADD,
	"SUB": SUB,
	"MUL": MUL,
	"DIV": DIV,
	"XOR": XOR,
	"AND": AND,

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
