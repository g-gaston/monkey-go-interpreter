package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

var _ Expression = &Literal{}

type Literal struct {
	Token token.Token
	Value int64
}

func (i *Literal) TokenLiteral() string {
	return i.Token.Literal
}
