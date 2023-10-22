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
				// Value: &ast.Identifier{
				// 	Token: intToken(5),
				// 	Value: "5",
				// },
			},
			&ast.Let{
				Token: letToken(),
				Name:  identifier("y"),
				// Value: &ast.Identifier{
				// 	Token: intToken(10),
				// 	Value: "10",
				// },
			},
			&ast.Let{
				Token: letToken(),
				Name:  identifier("foobar"),
				// Value: &ast.Identifier{
				// 	Token: intToken(838383),
				// 	Value: "838383",
				// },
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
			},
			&ast.Return{
				Token: returnToken(),
			},
			&ast.Return{
				Token: returnToken(),
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
	g := NewWithT(t)

	input := `foobar;`

	wantProgram := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Token:      identifierToken("foobar"),
				Expression: identifier("foobar"),
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

func intToken(n int) token.Token {
	return token.Token{
		Type:    token.Let,
		Literal: fmt.Sprintf("%d", n),
	}
}

func identifier(name string) *ast.Identifier {
	return &ast.Identifier{
		Token: identifierToken(name),
		Value: name,
	}
}
