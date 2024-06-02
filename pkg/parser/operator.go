package parser

import (
	"github.com/g-gaston/monkey-go-interpreter/pkg/ast"
	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
)

type precedence int

const (
	_ precedence = iota
	lowest
	equals
	lessGreater
	sum
	product
	prefix
	call
)

type operatorParserRegistry[P any] map[token.Type]P

func (r operatorParserRegistry[P]) register(t token.Type, parser P) {
	r[t] = parser
}

func (r operatorParserRegistry[P]) get(t token.Type) P {
	return r[t]
}

var tokenToInfixMapping = map[string]ast.InfixOperator{
	"+": ast.Addition,
	"-": ast.Subtraction,
}

func tokenToInfixOperator(token token.Token) ast.InfixOperator {
	return tokenToInfixMapping[token.Literal]
}
