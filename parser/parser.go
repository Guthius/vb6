package parser

import (
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
	body := make([]ast.Stmt, 0)

	p := newParser(tokens)
	for !p.isEof() {
		body = append(body, p.parseStmt())
	}

	return ast.BlockStmt{
		Body: body,
	}
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

func (p *parser) expect(kinds ...lexer.Kind) lexer.Token {
	kind := p.peek()
	for _, k := range kinds {
		if kind == k {
			return p.next()
		}
	}
	panic("unexpected token")
}
