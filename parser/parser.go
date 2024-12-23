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

func (p *parser) expect(kind lexer.Kind) lexer.Token {
	if p.peek() == kind {
		return p.next()
	}
	panic(expectedToken(p.Tokens[p.Pos], kind))
}

func (p *parser) expectOrEof(kind lexer.Kind) lexer.Token {
	k := p.peek()
	if k == kind || k == lexer.EOF {
		return p.next()
	}
	panic(expectedToken(p.Tokens[p.Pos], kind))
}

func expectedToken(tok lexer.Token, expected lexer.Kind) string {
	return fmt.Sprintf("unexpected token %s ('%s') at line %d, column %d (expected %s)",
		lexer.TokenKindString(tok.Kind),
		tok.Value, tok.Line, tok.Column,
		lexer.TokenKindString(expected))
}

func unexpectedToken(tok lexer.Token) string {
	return fmt.Sprintf("unexpected token %s ('%s') at line %d, column %d",
		lexer.TokenKindString(tok.Kind),
		tok.Value, tok.Line, tok.Column)
}

func (p *parser) unexpected() string {
	return unexpectedToken(p.Tokens[p.Pos])
}
