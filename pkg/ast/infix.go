package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

type InfixOperator string

const (
	Addition       InfixOperator = "+"
	Subtraction    InfixOperator = "-"
	Multiplication InfixOperator = "*"
	Division       InfixOperator = "/"
	GreaterThan    InfixOperator = ">"
	LessThan       InfixOperator = "<"
	Equal          InfixOperator = "=="
	NotEqual       InfixOperator = "!="
)

type Infix struct {
	Token    token.Token
	Operator InfixOperator
	Right    Expression
	Left     Expression
}

var _ Expression = &Infix{}

func (i *Infix) TokenLiteral() string {
	return i.Token.Literal
}
