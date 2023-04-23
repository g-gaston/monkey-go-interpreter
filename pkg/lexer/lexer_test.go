package lexer_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/g-gaston/monkey-go-interpreter/pkg/lexer"
	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestLexerNextToken(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantSequence []token.Token
		wantErr      string
	}{
		{
			name:  "one token",
			input: ";",
			wantSequence: []token.Token{
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.EOF},
			},
		},
		{
			name:  "multiple symbols",
			input: "=+(){},;",
			wantSequence: []token.Token{
				{Type: token.Assign, Literal: "="},
				{Type: token.Plus, Literal: "+"},
				{Type: token.LParen, Literal: "("},
				{Type: token.RParen, Literal: ")"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.RBrace, Literal: "}"},
				{Type: token.Comma, Literal: ","},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.EOF},
			},
		},
		{
			name: "actual code",
			input: `let five = 5;
let ten = 10;

let add = fn(x,y) {
	x + y;
};

let result = add(five, ten);

!-/*5;
5 < 10 > 5;
if (5 < 5) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`,
			wantSequence: []token.Token{
				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "five"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Int, Literal: "5"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "ten"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Int, Literal: "10"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "add"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Function, Literal: "fn"},
				{Type: token.LParen, Literal: "("},
				{Type: token.Ident, Literal: "x"},
				{Type: token.Comma, Literal: ","},
				{Type: token.Ident, Literal: "y"},
				{Type: token.RParen, Literal: ")"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.Ident, Literal: "x"},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Ident, Literal: "y"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RBrace, Literal: "}"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "result"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Ident, Literal: "add"},
				{Type: token.LParen, Literal: "("},
				{Type: token.Ident, Literal: "five"},
				{Type: token.Comma, Literal: ","},
				{Type: token.Ident, Literal: "ten"},
				{Type: token.RParen, Literal: ")"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Bang, Literal: "!"},
				{Type: token.Minus, Literal: "-"},
				{Type: token.Slash, Literal: "/"},
				{Type: token.Asterisk, Literal: "*"},
				{Type: token.Int, Literal: "5"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Int, Literal: "5"},
				{Type: token.LowerThan, Literal: "<"},
				{Type: token.Int, Literal: "10"},
				{Type: token.GreaterThan, Literal: ">"},
				{Type: token.Int, Literal: "5"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.If, Literal: "if"},
				{Type: token.LParen, Literal: "("},
				{Type: token.Int, Literal: "5"},
				{Type: token.LowerThan, Literal: "<"},
				{Type: token.Int, Literal: "5"},
				{Type: token.RParen, Literal: ")"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.Return, Literal: "return"},
				{Type: token.True, Literal: "true"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RBrace, Literal: "}"},
				{Type: token.Else, Literal: "else"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.Return, Literal: "return"},
				{Type: token.False, Literal: "false"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RBrace, Literal: "}"},

				{Type: token.Int, Literal: "10"},
				{Type: token.Equal, Literal: "=="},
				{Type: token.Int, Literal: "10"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.Int, Literal: "10"},
				{Type: token.NotEqual, Literal: "!="},
				{Type: token.Int, Literal: "9"},
				{Type: token.Semicolon, Literal: ";"},

				{Type: token.EOF},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(
				lexer.NewRunePeeker(bufio.NewReader(strings.NewReader(tt.input))),
			)
			var tok token.Token
			var err error
			var got []token.Token
			for tok, err = l.NextToken(); err == nil; tok, err = l.NextToken() {
				got = append(got, tok)
				if tok.Type == token.EOF {
					break
				}
			}

			assert.Equal(t, got, tt.wantSequence)
			if tt.wantErr == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
