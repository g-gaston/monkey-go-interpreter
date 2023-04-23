package lexer

import (
	"io"
	"unicode"

	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
)

type Lexer struct {
	peeker RunePeeker
}

func New(peeker RunePeeker) *Lexer {
	return &Lexer{
		peeker: peeker,
	}
}

func (r *Lexer) NextToken() (token.Token, error) {
	r.skipAllWhiteSpace()

	rune, _, err := r.peeker.ReadRune()
	if err == io.EOF {
		return token.Token{Type: token.EOF}, nil
	}
	if err != nil {
		return token.Token{}, err
	}


	// TODO: this can probably be simplified by using a lookup table
	switch rune {
	case '=':
		return r.parseEqualsStart()
	case '!':
		return r.parseEqualsBang()
	case '-':
		return token.Token{Type: token.Minus, Literal: string(rune)}, nil
	case '+':
		return token.Token{Type: token.Plus, Literal: string(rune)}, nil
	case ',':
		return token.Token{Type: token.Comma, Literal: string(rune)}, nil
	case ';':
		return token.Token{Type: token.Semicolon, Literal: string(rune)}, nil
	case '(':
		return token.Token{Type: token.LParen, Literal: string(rune)}, nil
	case ')':
		return token.Token{Type: token.RParen, Literal: string(rune)}, nil
	case '{':
		return token.Token{Type: token.LBrace, Literal: string(rune)}, nil
	case '}':
		return token.Token{Type: token.RBrace, Literal: string(rune)}, nil
	case '*':
		return token.Token{Type: token.Asterisk, Literal: string(rune)}, nil
	case '/':
		return token.Token{Type: token.Slash, Literal: string(rune)}, nil
	case '<':
		return token.Token{Type: token.LowerThan, Literal: string(rune)}, nil
	case '>':
		return token.Token{Type: token.GreaterThan, Literal: string(rune)}, nil
	default:
		return r.parseMultiCharSymbol(rune)
	}
}

func (r *Lexer) skipAllWhiteSpace() {
	for ru, err := r.peeker.PeekRune(); err == nil && unicode.IsSpace(ru); ru, err = r.peeker.PeekRune() {
		r.peeker.ReadRune()
	}
}

func (r *Lexer) parseMultiCharSymbol(currentRune rune) (token.Token, error) {
	if unicode.IsLetter(currentRune) {
		return r.parseWord(currentRune)
	} else if unicode.IsDigit(currentRune) {
		return r.parseNumber(currentRune)
	}

	return token.Token{Type: token.Illegal, Literal: string(currentRune)}, nil
}

func (r *Lexer) parseWord(currentRune rune) (token.Token, error) {
	letterRunes := []rune{currentRune}
	ru, err := r.peeker.PeekRune()
	for ; err == nil && unicode.IsLetter(ru); ru, err = r.peeker.PeekRune() {
		letterRunes = append(letterRunes, ru)
		r.peeker.ReadRune()
	}

	word := string(letterRunes)

	if t, ok := token.IsKeyword(word); ok {
		return token.Token{Type: t, Literal: word}, nil
	}

	return token.Token{Type: token.Ident, Literal: word}, nil
}

func (r *Lexer) parseNumber(currentRune rune) (token.Token, error) {
	digitRunes := []rune{currentRune}
	ru, err := r.peeker.PeekRune()
	for ; err == nil && unicode.IsDigit(ru); ru, err = r.peeker.PeekRune() {
		digitRunes = append(digitRunes, ru)
		r.peeker.ReadRune()
	}

	return token.Token{Type: token.Int, Literal: string(digitRunes)}, nil
}

func (r *Lexer) parseEqualsStart() (token.Token, error) {
	if ru, err := r.peeker.PeekRune(); err == nil && ru == '=' {
		r.peeker.ReadRune()
		return token.Token{Type: token.Equal, Literal: "=="}, nil
	}

	return token.Token{Type: token.Assign, Literal: "="}, nil
}

func (r *Lexer) parseEqualsBang() (token.Token, error) {
	if ru, err := r.peeker.PeekRune(); err == nil && ru == '=' {
		r.peeker.ReadRune()
		return token.Token{Type: token.NotEqual, Literal: "!="}, nil
	}

	return token.Token{Type: token.Bang, Literal: "!"}, nil
}
