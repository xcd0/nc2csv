package nc2csv

// 予約語リスト
var keywords = []string{ // {{{
	// 実装済み予約語
	"%",
	"(",
	")",
	// 実装予定予約語
	"GOTO",
	"IF",
	"THEN",
	"WHILE",
	"DO",
	"END",
	"[",
	"]",
	"=",
	"+",
	"-",
	"*",
	"/",
	"#",
	"EQ",
	"NE",
	"LT",
	"GT",
	"LE",
	"GE",
	"AND",
	"OR",
	"XOR",
	"EOF",
	// デバッグ用 EOF

	// 未実装予約語
	/*
		"ADD",
		"SUB",
		"MUL",
		"DIV",
		"COS",
		"SIN",
		"TAN",
		"ACOS",
		"ASIN",
		"ATAN",
	*/
} // }}}
