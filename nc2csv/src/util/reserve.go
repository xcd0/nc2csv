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
	"(": COMMENTSTART,
	")": COMMENTEND,

	// 実装予定予約語

	"GOTO":  GOTO,
	"IF":    IF,
	"WHILE": WHILE,
	"END":   END,

	"=": ASSIGN,
	"+": PLUS,
	"-": MINUS,
	"*": ASTERISK,
	"/": SLASH,
	"#": ARRAY,

	"EQ": EQ,
	"NE": NE,
	"LT": LT,
	"GT": GT,
	"LE": LE,
	"GE": GE,

	"AND": AND,
	"OR":  OR,
	"XOR": XOR,

	// デバッグ用 EOF
	"EOF": EOF,

	// 未実装予約語
	/*
		"ADD":  ADD,
		"SUB":  SUB,
		"MUL":  MUL,
		"DIV":  DIV,
		"COS":  COS,
		"SIN":  SIN,
		"TAN":  TAN,
		"ACOS": ACOS,
		"ASIN": ASIN,
		"ATAN": ATAN,
	*/

	//"SKIP":  SKIP,
	//"EOB": EOB,

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
