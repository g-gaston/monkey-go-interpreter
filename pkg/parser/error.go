package parser

import (
	"errors"
	"fmt"

	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
)

type Error struct {
	err   error
	token token.Token
}

func NewError(err error, t token.Token) Error {
	e := Error{}
	if errors.As(err, &e) {
		return e
	}

	e.err = err
	e.token = t
	return e
}

func (e Error) Error() string {
	return fmt.Sprintf("invalid program at %s: %s", e.token, e.err)
}
