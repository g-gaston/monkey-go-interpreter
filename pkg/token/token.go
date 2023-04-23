package token

import "fmt"

type Type int

func (t Type) String() string {
	return TypeString(t)
}

type Token struct {
	Type    Type
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("token.Token{Type:%s, Literal:\"%s\"}", t.Type.String(), t.Literal)
}

const (
	Illegal Type = iota
	EOF

	Ident
	Int

	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash

	Equal
	NotEqual
	LowerThan
	GreaterThan

	Comma
	Semicolon

	LParen
	RParen
	LBrace
	RBrace

	Function
	Let
	True
	False
	If
	Else
	Return

	upperLimit
)

const numberOfTypes = upperLimit

var typeStrings = []string{
	"ILLEGAL",
	"EOF",
	"IDENT",
	"INT",
	"ASSIGN",
	"+",
	"-",
	"!",
	"*",
	"/",
	"==",
	"!=",
	"<",
	">",
	",",
	";",
	"(",
	")",
	"{",
	"}",
	"FUNCTION",
	"LET",
	"TRUE",
	"FALSE",
	"IF",
	"ELSE",
	"RETURN",
}

func TypeString(t Type) string {
	if t >= numberOfTypes {
		return "ILLEGAL"
	}

	return typeStrings[t]
}
