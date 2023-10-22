package ast

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

var _ Statement = &ExpressionStatement{}

// TODO: should this receiver be a non pointer?
func (c *ExpressionStatement) TokenLiteral() string {
	return c.Token.Literal
}
