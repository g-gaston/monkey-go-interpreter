package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

var _ Statement = &Return{}

type Return struct {
	Token token.Token
	Value Expression
}

func (l *Return) TokenLiteral() string {
	return l.Token.Literal
}
