package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

var _ Expression = &Prefix{}

type PrefixOperator string

const (
	Not      PrefixOperator = "!"
	Negative PrefixOperator = "-"
)

type Prefix struct {
	Token    token.Token
	Operator PrefixOperator
	Right    Expression
}

func (i *Prefix) TokenLiteral() string {
	return i.Token.Literal
}
