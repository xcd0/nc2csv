package lexer

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	ASSIGN   = "=" // 変数#への代入で使われる

	ADD = "ADD"
	SUB = "SUB"
	MUL = "MUL"
	DIV = "DIV"

	EQ  = "EQ"  // ==
	NE  = "NE"  // !=
	LT  = "LT"  // <
	LE  = "LE"  // <=
	GT  = "GT"  // >
	GE  = "GE"  // >=
	OR  = "OR"  // ||
	XOR = "XOR" // &^
	AND = "AND" // &&

	COMMENTSTART = "COMMENTSTART" // '('
	COMMENTEND   = "COMMENTEND"   // ')'
	NCEOF        = "NCEOF"        // '%'
	VARIABLE     = "VARIABLE"     // '#'
	EOB          = "EOB"          // '\n'

	AXIS          = "AXIS"
	PREPARATION   = "PREPARATION"
	MISCELLANEOUS = "MISCELLANEOUS"
	FEED          = "FEED"
	SPINDLE       = "SPINDLE"
	TOOL          = "TOOL"
	ONUM          = "ONUM"
	SKIP          = "SKIP"

	GOTO  = "GOTO"
	IF    = "IF"
	WHILE = "WHILE"
	END   = "END"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"=":             ASSIGN,
	"+":             PLUS,
	"-":             MINUS,
	"*":             ASTERISK,
	"/":             SLASH,
	"ADD":           ADD,
	"SUB":           SUB,
	"MUL":           MUL,
	"DIV":           DIV,
	"EQ":            EQ,
	"NE":            NE,
	"LT":            LT,
	"LE":            LE,
	"GT":            GT,
	"GE":            GE,
	"OR":            OR,
	"XOR":           XOR,
	"AND":           AND,
	"(":             COMMENTSTART,
	")":             COMMENTEND,
	"%":             NCEOF,
	"#":             VARIABLE,
	"AXIS":          AXIS,
	"PREPARATION":   PREPARATION,
	"MISCELLANEOUS": MISCELLANEOUS,
	"FEED":          FEED,
	"SPINDLE":       SPINDLE,
	"TOOL":          TOOL,
	"ONUM":          ONUM,
	"SKIP":          SKIP,
	"EOB":           EOB,
	"GOTO":          GOTO,
	"IF":            IF,
	"WHILE":         WHILE,
	"END":           END,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
