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
		{
			name:  "with simple infix operator +",
			input: `5 + 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    plusToken(),
							Left:     literal(5),
							Operator: ast.Addition,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple infix operator *",
			input: `5 * 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    asteriskToken(),
							Left:     literal(5),
							Operator: ast.Multiplication,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple infix operator /",
			input: `5 / 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    slashToken(),
							Left:     literal(5),
							Operator: ast.Division,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple infix operator >",
			input: `5 > 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    greaterThanToken(),
							Left:     literal(5),
							Operator: ast.GreaterThan,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple infix operator <",
			input: `5 < 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    lessThanToken(),
							Left:     literal(5),
							Operator: ast.LessThan,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple infix operator ==",
			input: `5 == 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    equalToken(),
							Left:     literal(5),
							Operator: ast.Equal,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "with simple infix operator !=",
			input: `5 != 5;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: &ast.Infix{
							Token:    notEqualToken(),
							Left:     literal(5),
							Operator: ast.NotEqual,
							Right:    literal(5),
						},
					},
				},
			},
		},
		{
			name:  "-a * b",
			input: `-a * b;`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: minusToken(),
						Expression: &ast.Infix{
							Token: asteriskToken(),
							Left: &ast.Prefix{
								Token:    minusToken(),
								Right:    identifier("a"),
								Operator: ast.Negative,
							},
							Operator: ast.Multiplication,
							Right:    identifier("b"),
						},
					},
				},
			},
		},
		{
			name:  "!-a",
			input: `!-a`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: bangToken(),
						Expression: &ast.Prefix{
							Token: bangToken(),
							Right: &ast.Prefix{
								Token:    minusToken(),
								Right:    identifier("a"),
								Operator: ast.Negative,
							},
							Operator: ast.Not,
						},
					},
				},
			},
		},
		{
			name:  "a + b + c",
			input: `a + b + c`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: identifierToken("a"),
						Expression: &ast.Infix{
							Token: plusToken(),
							Left: &ast.Infix{
								Token:    plusToken(),
								Left:     identifier("a"),
								Right:    identifier("b"),
								Operator: ast.Addition,
							},
							Operator: ast.Addition,
							Right:    identifier("c"),
						},
					},
				},
			},
		},
		{
			name:  "a + b - c",
			input: `a + b - c`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: identifierToken("a"),
						Expression: &ast.Infix{
							Token: minusToken(),
							Left: &ast.Infix{
								Token:    plusToken(),
								Left:     identifier("a"),
								Right:    identifier("b"),
								Operator: ast.Addition,
							},
							Operator: ast.Subtraction,
							Right:    identifier("c"),
						},
					},
				},
			},
		},
		{
			name:  "a * b * c",
			input: `a * b * c`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: identifierToken("a"),
						Expression: &ast.Infix{
							Token: asteriskToken(),
							Left: &ast.Infix{
								Token:    asteriskToken(),
								Left:     identifier("a"),
								Right:    identifier("b"),
								Operator: ast.Multiplication,
							},
							Operator: ast.Multiplication,
							Right:    identifier("c"),
						},
					},
				},
			},
		},
		{
			name:  "a * b / c",
			input: `a * b / c`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: identifierToken("a"),
						Expression: &ast.Infix{
							Token: slashToken(),
							Left: &ast.Infix{
								Token:    asteriskToken(),
								Left:     identifier("a"),
								Right:    identifier("b"),
								Operator: ast.Multiplication,
							},
							Operator: ast.Division,
							Right:    identifier("c"),
						},
					},
				},
			},
		},
		{
			name:  "a + b / c",
			input: `a + b / c`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: identifierToken("a"),
						Expression: &ast.Infix{
							Token: plusToken(),
							Left:  identifier("a"),
							Right: &ast.Infix{
								Token:    slashToken(),
								Left:     identifier("b"),
								Right:    identifier("c"),
								Operator: ast.Division,
							},
							Operator: ast.Addition,
						},
					},
				},
			},
		},
		{
			name:  "a + b * c + d / e - f",
			input: `a + b * c + d / e - f`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: identifierToken("a"),
						Expression: sub(
							add(
								add("a", multiply("b", "c")),
								divide("d", "e"),
							),
							"f",
						),
					},
				},
			},
		},
		{
			name:  "5 > 4 == 3 < 4",
			input: `5 > 4 == 3 < 4`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: equal(
							greaterThan(5, 4),
							lessThan(3, 4),
						),
					},
				},
			},
		},
		{
			name:  "5 < 4 != 3 > 4",
			input: `5 < 4 != 3 > 4`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(5),
						Expression: notEqual(
							lessThan(5, 4),
							greaterThan(3, 4),
						),
					},
				},
			},
		},
		{
			name:  "3 + 4 * 5 == 3 * 1 + 4 * 5",
			input: `3 + 4 * 5 == 3 * 1 + 4 * 5`,
			wantProgram: &ast.Root{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{
						Token: intToken(3),
						Expression: equal(
							add(3, multiply(4, 5)),
							add(multiply(3, 1), multiply(4, 5)),
						),
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

func plusToken() token.Token {
	return token.Token{
		Type:    token.Plus,
		Literal: "+",
	}
}

func asteriskToken() token.Token {
	return token.Token{
		Type:    token.Asterisk,
		Literal: "*",
	}
}

func slashToken() token.Token {
	return token.Token{
		Type:    token.Slash,
		Literal: "/",
	}
}

func greaterThanToken() token.Token {
	return token.Token{
		Type:    token.GreaterThan,
		Literal: ">",
	}
}

func lessThanToken() token.Token {
	return token.Token{
		Type:    token.LowerThan,
		Literal: "<",
	}
}

func equalToken() token.Token {
	return token.Token{
		Type:    token.Equal,
		Literal: "==",
	}
}

func notEqualToken() token.Token {
	return token.Token{
		Type:    token.NotEqual,
		Literal: "!=",
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

// castExpression is a helper function to cast a string or an ast.Expression to an ast.Expression.
// Useful to compose expected ASTs for tests.
func castExpression(e any) ast.Expression {
	switch e := e.(type) {
	case string:
		return identifier(e)
	case int:
		return literal(int64(e))
	case ast.Expression:
		return e
	}
	return nil
}

func add(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    plusToken(),
		Left:     castExpression(a),
		Right:    castExpression(b),
		Operator: ast.Addition,
	}
}

func sub(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    minusToken(),
		Left:     castExpression(a),
		Operator: ast.Subtraction,
		Right:    castExpression(b),
	}
}

func multiply(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    asteriskToken(),
		Left:     castExpression(a),
		Operator: ast.Multiplication,
		Right:    castExpression(b),
	}
}

func divide(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    slashToken(),
		Left:     castExpression(a),
		Operator: ast.Division,
		Right:    castExpression(b),
	}
}

func greaterThan(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    greaterThanToken(),
		Left:     castExpression(a),
		Operator: ast.GreaterThan,
		Right:    castExpression(b),
	}
}

func lessThan(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    lessThanToken(),
		Left:     castExpression(a),
		Operator: ast.LessThan,
		Right:    castExpression(b),
	}
}

func equal(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    equalToken(),
		Left:     castExpression(a),
		Operator: ast.Equal,
		Right:    castExpression(b),
	}
}

func notEqual(a, b any) *ast.Infix {
	return &ast.Infix{
		Token:    notEqualToken(),
		Left:     castExpression(a),
		Operator: ast.NotEqual,
		Right:    castExpression(b),
	}
}
