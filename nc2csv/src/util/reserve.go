package util

const (
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

	COS  = "COS"
	SIN  = "SIN"
	TAN  = "TAN"
	ACOS = "ACOS"
	ASIN = "ASIN"
	ATAN = "ATAN"

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
)

var keywords = map[string]string{

	// 実装済み予約語

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
