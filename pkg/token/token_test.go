package token_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
)

func TestTypeString(t *testing.T) {
	tests := []struct {
		name string
		t    token.Type
		want string
	}{
		{
			name: "Illegal",
			t:    token.Illegal,
			want: "ILLEGAL",
		},
		{
			name: "Unknown",
			t:    token.Type(math.MaxInt),
			want: "ILLEGAL",
		},
		{
			name: "EOF",
			t:    token.EOF,
			want: "EOF",
		},
		{
			name: "Ident",
			t:    token.Ident,
			want: "IDENT",
		},
		{
			name: "Int",
			t:    token.Int,
			want: "INT",
		},
		{
			name: "Assign",
			t:    token.Assign,
			want: "ASSIGN",
		},
		{
			name: "Plus",
			t:    token.Plus,
			want: "+",
		},
		{
			name: "Comma",
			t:    token.Comma,
			want: ",",
		},
		{
			name: "Semicolon",
			t:    token.Semicolon,
			want: ";",
		},
		{
			name: "LParen",
			t:    token.LParen,
			want: "(",
		},
		{
			name: "RParen",
			t:    token.RParen,
			want: ")",
		},
		{
			name: "LBrace",
			t:    token.LBrace,
			want: "{",
		},
		{
			name: "RBrance",
			t:    token.RBrace,
			want: "}",
		},
		{
			name: "Function",
			t:    token.Function,
			want: "FUNCTION",
		},
		{
			name: "Let",
			t:    token.Let,
			want: "LET",
		},
		{
			name: "true",
			t:    token.True,
			want: "TRUE",
		},
		{
			name: "false",
			t:    token.False,
			want: "FALSE",
		},
		{
			name: "if",
			t:    token.If,
			want: "IF",
		},
		{
			name: "else",
			t:    token.Else,
			want: "ELSE",
		},
		{
			name: "return",
			t:    token.Return,
			want: "RETURN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, token.TypeString(tt.t), tt.want)
		})
	}
}
