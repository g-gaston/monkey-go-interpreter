package parser

import (
	"errors"
	"strconv"

	// TODO: cleanup package name collision
	perrors "github.com/pkg/errors"

	"github.com/g-gaston/monkey-go-interpreter/pkg/ast"
	"github.com/g-gaston/monkey-go-interpreter/pkg/lexer"
	"github.com/g-gaston/monkey-go-interpreter/pkg/token"
)

type Parser struct {
	lexer         *lexer.Lexer
	current, peek token.Token
	errors        []Error
	prefixParsers operatorParserRegistry[prefixParser]
	infixParsers  operatorParserRegistry[infixParser]
}

type (
	prefixParser func() (ast.Expression, error)
	infixParser  func(ast.Expression) (ast.Expression, error)
)

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:         lexer,
		prefixParsers: make(operatorParserRegistry[prefixParser]),
		infixParsers:  make(operatorParserRegistry[infixParser]),
	}

	p.prefixParsers.register(token.Ident, p.parseIdentifier)
	p.prefixParsers.register(token.Int, p.parseLiteral)
	p.prefixParsers.register(token.Bang, p.parsePrefix)
	p.prefixParsers.register(token.Minus, p.parsePrefix)

	return p
}

func (p *Parser) Parse() (*ast.Root, error) {
	r := ast.Root{}

	// Read the initial tokens
	// We do it twice to populate both current and peek
	p.advanceToken()
	p.advanceToken()

	for p.current.Type != token.EOF {
		if statement, err := p.parseStatement(); err == nil {
			r.Statements = append(r.Statements, statement)
		} else {
			p.errors = append(p.errors, NewError(err, p.current))
		}
		p.advanceToken()
	}

	return &r, p.error()
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.current.Type {
	case token.Let:
		return p.parseLet()
	case token.Return:
		return p.parseReturn()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) Errors() []error {
	if len(p.errors) == 0 {
		return nil
	}

	e := make([]error, 0, len(p.errors))
	for _, err := range p.errors {
		e = append(e, err)
	}
	return e
}

func (p *Parser) error() error {
	return errors.Join(p.Errors()...)
}

func (p *Parser) advanceToken() {
	p.current = p.peek
	t, err := p.lexer.NextToken()
	for ; err != nil; t, err = p.lexer.NextToken() {
		p.errors = append(p.errors, NewError(err, p.peek))
	}
	p.peek = t
}

func (p *Parser) assertPeek(wantTokenType token.Type) error {
	if p.peek.Type != wantTokenType {
		return NewError(perrors.Errorf("expected token type %s but got %s", wantTokenType, p.peek.Type), p.peek)
	}
	return nil
}

func (p *Parser) parseLet() (*ast.Let, error) {
	l := &ast.Let{
		Token: p.current,
	}

	if err := p.assertPeek(token.Ident); err != nil {
		return nil, err
	}

	l.Name = &ast.Identifier{Token: p.peek, Value: p.peek.Literal}

	p.advanceToken()

	if err := p.assertPeek(token.Assign); err != nil {
		return nil, err
	}

	p.advanceToken()
	// now current is assign

	p.advanceToken()
	// now current is at the beginning of the expression so we start its parsing

	exp, err := p.parseExpression(0) // TODO: what precedence do we use here?
	if err != nil {
		return nil, err
	}

	l.Value = exp

	return l, nil
}

func (p *Parser) parseExpression(precedence precedence) (ast.Expression, error) {
	prefixParser := p.prefixParsers.get(p.current.Type)
	if prefixParser == nil {
		return nil, NewError(perrors.New("can't find a prefix operator for token"), p.peek)
	}

	left, err := prefixParser()
	if err != nil {
		return nil, err
	}
	p.advanceToken() // Advancing here so we pass tests.
	return left, nil
}

func (p *Parser) parseReturn() (*ast.Return, error) {
	r := &ast.Return{
		Token: p.current,
	}

	// now current is at the beginning of the expression
	p.advanceToken()

	exp, err := p.parseExpression(0) // TODO: what precedence do we use here?
	if err != nil {
		return nil, err
	}

	r.Value = exp

	return r, nil
}

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	s := &ast.ExpressionStatement{
		Token: p.current,
	}
	var err error
	s.Expression, err = p.parseExpression(lowest)
	if err != nil {
		return nil, err
	}

	if p.peek.Type == token.Semicolon {
		p.advanceToken()
	}

	return s, nil
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	return &ast.Identifier{Token: p.current, Value: p.current.Literal}, nil
}

func (p *Parser) parseLiteral() (ast.Expression, error) {
	value, err := strconv.ParseInt(p.current.Literal, 0, 64)
	if err != nil {
		return nil, err
	}
	return &ast.Literal{Token: p.current, Value: value}, nil
}

func (p *Parser) parsePrefix() (ast.Expression, error) {
	prefixExp := &ast.Prefix{
		Token: p.current,
	}
	switch p.current.Literal {
	case "!":
		prefixExp.Operator = ast.Not
	case "-":
		prefixExp.Operator = ast.Negative
	}

	p.advanceToken()

	right, err := p.parseExpression(prefix)
	if err != nil {
		return nil, err
	}

	prefixExp.Right = right
	return prefixExp, nil
}
