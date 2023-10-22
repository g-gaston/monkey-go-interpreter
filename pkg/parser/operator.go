package parser

import "github.com/g-gaston/monkey-go-interpreter/pkg/token"

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

type precedence int

type operatorParserRegistry[P any] map[token.Type]P

func (r operatorParserRegistry[P]) register(t token.Type, parser P) {
	r[t] = parser
}

func (r operatorParserRegistry[P]) get(t token.Type) P {
	return r[t]
}
