package parser

import (
	"github.com/guthius/vb6/ast"
	"github.com/guthius/vb6/lexer"
)

func (p *parser) parseStmt() ast.Stmt {
	var kind lexer.Kind

	for !p.isEof() {
		kind = p.peek()
		if kind != lexer.LineBreak {
			break
		}
		p.next()
	}

	sfn, ok := tables.stmt[kind]
	if !ok {
		expr := parseExpr(p, defaultBindingPower)
		p.expectOrEof(lexer.LineBreak)
		return ast.ExprStmt{Expr: expr}
	}

	return sfn(p)
}

func parseConstDeclStmt(p *parser, public bool) ast.Stmt {
	if p.next().Kind != lexer.Const {
		panic("expected const")
	}

	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.Equal)
	value := parseExpr(p, assignment)
	p.expectOrEof(lexer.LineBreak)

	return ast.ConstDeclStmt{
		Public:     public,
		Identifier: identifier,
		Value:      value,
	}
}

func parsePrivateConstDeclStmt(p *parser) ast.Stmt {
	return parseConstDeclStmt(p, false)
}

func parseDeclStmt(p *parser) ast.Stmt {
	public := p.next().Kind == lexer.Public

	if p.peek() == lexer.Const {
		return parseConstDeclStmt(p, public)
	}

	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.As)
	dataType := parseTypeExpr(p)

	var value ast.Expr
	if p.peek() == lexer.Equal {
		p.next()
		value = parseExpr(p, assignment)
	}

	p.expectOrEof(lexer.LineBreak)

	return ast.VarDeclStmt{
		Identifier: identifier,
		Type:       dataType,
		Value:      value,
		Public:     public,
	}
}
