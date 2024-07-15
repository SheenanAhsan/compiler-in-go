package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// Special types
	ILLEGAL = "ILLEGAL" // token/character we don't know
	EOF = "EOF" // end of file

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// if tokentype found, return tokentype
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// else, return default IDENT, ex) variable x
	return IDENT
}
