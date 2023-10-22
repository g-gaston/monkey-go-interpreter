package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

var _ Statement = &Let{}

type Let struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *Let) TokenLiteral() string {
	return l.Token.Literal
}
