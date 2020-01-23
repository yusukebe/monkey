package token

type TokenType string

toke Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456

	// 演算子
	ASSIGN = "="
	PLUS = "+"

	// デリミタ
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET = "LET"
)
