package parser

import (
	"fmt"

	"github.com/guthius/vb6/ast"
	"github.com/guthius/vb6/lexer"
)

type parser struct {
	Tokens []lexer.Token
	Pos    int
}

func newParser(tokens []lexer.Token) *parser {
	return &parser{
		Tokens: tokens,
		Pos:    0,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	return parseBlockStmt(newParser(tokens), lexer.EOF)
}

func (p *parser) peek() lexer.Kind {
	return p.Tokens[p.Pos].Kind
}

func (p *parser) next() lexer.Token {
	p.Pos++
	return p.Tokens[p.Pos-1]
}

func (p *parser) isEof() bool {
	return p.Pos >= len(p.Tokens) || p.peek() == lexer.EOF
}

func (p *parser) expectError(kind lexer.Kind, err string) lexer.Token {
	k := p.peek()
	if k == kind {
		return p.next()
	}
	panic(fmt.Sprintf("%s (got %s, expected %s)",
		err,
		lexer.TokenKindString(k),
		lexer.TokenKindString(kind)))
}

func (p *parser) expect(kind lexer.Kind) lexer.Token {
	return p.expectError(kind, "unexpected token")
}

func (p *parser) expectOrEof(kind lexer.Kind) lexer.Token {
	k := p.peek()
	if k == kind || k == lexer.EOF {
		return p.next()
	}
	panic(fmt.Sprintf("unexpected token (got %v, expected %v)",
		lexer.TokenKindString(k),
		lexer.TokenKindString(kind)))
}
