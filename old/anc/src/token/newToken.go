package token

func NewToken(tokenType TokenType, ch rune) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
