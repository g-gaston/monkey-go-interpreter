package lexer

import "io"

type RunePeeker interface {
	io.RuneReader
	// PeekRune returns the next rune, if available, without consuming it.
	PeekRune() (rune, error)
}

type runePeeker struct {
	reader io.RuneReader
	peek   rune
	size   int
	err    error
}

func NewRunePeeker(reader io.RuneReader) RunePeeker {
	return &runePeeker{reader: reader}
}

func (p *runePeeker) PeekRune() (rune, error) {
	if p.peek == 0 {
		p.peek, p.size, p.err = p.reader.ReadRune()
		return p.peek, p.err
	}

	return p.peek, nil
}

func (p *runePeeker) ReadRune() (r rune, size int, err error) {
	if p.peek == 0 {
		return p.reader.ReadRune()
	}

	r, s, err := p.peek, p.size, p.err
	p.peek, p.size, p.err = 0, 0, nil

	return r, s, err
}
