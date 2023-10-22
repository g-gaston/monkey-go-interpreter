package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

var _ Expression = &Identifier{}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
