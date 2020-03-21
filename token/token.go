// token/token.go

package token

// TokenType を様々な値として使用し、トークンのタイプを区別する.
type TokenType string

// Token .
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL is token, string is unknown.
	ILLEGAL = "ILLEGAL"
	// EOF is end of file.
	EOF = "EOF"

	// IDENT (識別子) + リテラル .
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT (整数).
	INT = "INT" // 13535

	// ASSIGN (演算子) .
	ASSIGN = "="
	// PLUS ASSIGN.
	PLUS = "+"

	// COMMA delimiter.
	COMMA = ","
	// SEMICOLON delimiter.
	SEMICOLON = ";"

	// LPAREN (左括弧)
	LPAREN = "("
	// RPAREN (右括弧)
	RPAREN = ")"
	// LBRACE (左波括弧)
	LBRACE = "{"
	// RBRACE (右波括弧)
	RBRACE = "}"

	// FUNCTION .
	FUNCTION = "FUNCTION"
	// LET .
	LET = "LET"
)
