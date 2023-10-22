package ast

import "strings"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

var _ Node = &Root{}

// Root is the beginning of the AST. It represents the whole program.
type Root struct {
	Statements []Statement
}

func (r *Root) TokenLiteral() string {
	// todo: optimize?
	var literals []string
	for _, s := range r.Statements {
		literals = append(literals, s.TokenLiteral())
	}

	return strings.Join(literals, " ")
}
