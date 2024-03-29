package parser_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/g-gaston/monkey-go-interpreter/pkg/ast"
	"github.com/g-gaston/monkey-go-interpreter/pkg/lexer"
	"github.com/g-gaston/monkey-go-interpreter/pkg/parser"
	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
)

func TestParserParseLetStatements(t *testing.T) {
	g := NewWithT(t)

	input := `let x = 5;
let y = 10;
let foobar = 838383;`

	wantProgram := &ast.Root{
		Statements: []ast.Statement{
			&ast.Let{
				Token: letToken(),
				Name:  identifier("x"),
				Value: literal(5),
			},
			&ast.Let{
				Token: letToken(),
				Name:  identifier("y"),
				Value: literal(10),
			},
			&ast.Let{
				Token: letToken(),
				Name:  identifier("foobar"),
				Value: literal(838383),
			},
		},
	}

	l := lexer.New(
		lexer.NewRunePeeker(bufio.NewReader(strings.NewReader(input))),
	)

	p := parser.New(l)

	program, err := p.Parse()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(program).To(BeComparableTo(wantProgram))
}

func TestParserParseReturnStatements(t *testing.T) {
	g := NewWithT(t)

	input := `return 5;
return 10;
return 838383;`

	wantProgram := &ast.Root{
		Statements: []ast.Statement{
			&ast.Return{
				Token: returnToken(),
				Value: &ast.Literal{
					Token: intToken(5),
					Value: 5,
				},
			},
			&ast.Return{
				Token: returnToken(),
				Value: &ast.Literal{
					Token: intToken(10),
					Value: 10,
				},
			},
			&ast.Return{
				Token: returnToken(),
				Value: &ast.Literal{
					Token: intToken(838383),
					Value: 838383,
				},
			},
		},
	}

	l := lexer.New(
		lexer.NewRunePeeker(bufio.NewReader(strings.NewReader(input))),
	)

	p := parser.New(l)

	program, err := p.Parse()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(program).To(BeComparableTo(wantProgram))
}

func TestParserParseExpressionStatement(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		wantProgram *ast.Root
	}{
		{
			name:  "with simple identifier",
			input: `foobar;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      identifierToken("foobar"),
						Expression: identifier("foobar"),
					},
				},
			},
		},
		{
			name:  "with simple literal",
			input: `5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token:      intToken(5),
						Expression: literal(5),
					},
				},
			},
		},
		{
			name:  "with simple prefix expression bang",
			input: `!5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: bangToken(),
						Expression: &ast.Prefix{
							Token:    bangToken(),
							Operator: ast.Not,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple prefix expression -",
			input: `-15;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: minusToken(),
						Expression: &ast.Prefix{
							Token:    minusToken(),
							Operator: ast.Negative,
							Right:    literal(15),
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)
			l := lexer.New(
				lexer.NewRunePeeker(bufio.NewReader(strings.NewReader(tc.input))),
			)

			p := parser.New(l)

			program, err := p.Parse()
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(program).To(BeComparableTo(tc.wantProgram))
		})
	}
}

func letToken() token.Token {
	return token.Token{
		Type:    token.Let,
		Literal: "let",
	}
}

func returnToken() token.Token {
	return token.Token{
		Type:    token.Return,
		Literal: "return",
	}
}

func identifierToken(literal string) token.Token {
	return token.Token{
		Type:    token.Ident,
		Literal: literal,
	}
}

func intToken(n int64) token.Token {
	return token.Token{
		Type:    token.Int,
		Literal: fmt.Sprintf("%d", n),
	}
}

func bangToken() token.Token {
	return token.Token{
		Type:    token.Bang,
		Literal: "!",
	}
}

func minusToken() token.Token {
	return token.Token{
		Type:    token.Minus,
		Literal: "-",
	}
}

func identifier(name string) *ast.Identifier {
	return &ast.Identifier{
		Token: identifierToken(name),
		Value: name,
	}
}

func literal(value int64) *ast.Literal {
	return &ast.Literal{
		Token: intToken(value),
		Value: value,
	}
}
