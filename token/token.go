package token

type TokenType string

const (
	IDENT = "IDENT"
	LET   = "LET"

	NUMBER = "NUMBER"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	STRING = "STRING"
	ARRAY  = "ARRAY"
	NULL   = "NULL"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	BANG = "!"

	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	LE     = "<="
	GT     = ">"
	GE     = ">="

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"

	FUNCTION = "func"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

type Token struct {
	Line    int
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"func":   FUNCTION,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"null":   NULL,
}

func LookupIdent(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
