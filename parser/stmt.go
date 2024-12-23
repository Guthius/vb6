package parser

import (
	"github.com/guthius/vb6/ast"
	"github.com/guthius/vb6/lexer"
)

func (p *parser) skipLineBreaks() {
	for !p.isEof() {
		if p.peek() != lexer.LineBreak {
			break
		}
		p.next()
	}
}

func (p *parser) parseStmt() ast.Stmt {
	sfn, ok := tables.stmt[p.peek()]
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
		Public:     public,
		Identifier: identifier,
		Type:       dataType,
		Value:      value,
	}
}

func parseRangeExpr(p *parser) []ast.RangeExpr {
	ranges := make([]ast.RangeExpr, 0)

	p.expect(lexer.LParen)
	for {
		l := parseExpr(p, assignment)
		p.expect(lexer.To)
		u := parseExpr(p, assignment)
		ranges = append(ranges, ast.RangeExpr{LBound: l, UBound: u})
		if p.peek() != lexer.Comma {
			break
		}
		p.next()
	}

	p.expect(lexer.RParen)

	return ranges
}

func parseFieldDeclExpr(p *parser) ast.FieldDeclExpr {
	var ranges []ast.RangeExpr

	identifier := p.expect(lexer.Identifier).Value
	if p.peek() == lexer.LParen {
		ranges = parseRangeExpr(p)
	}

	p.expect(lexer.As)
	dataType := parseTypeExpr(p)
	p.expect(lexer.LineBreak)

	return ast.FieldDeclExpr{
		Identifier: identifier,
		Type:       dataType,
		IsArray:    ranges != nil,
		Ranges:     ranges,
	}
}

func parseTypeStmt(p *parser) ast.Stmt {
	p.expect(lexer.Type)
	identifier := p.expect(lexer.Identifier).Value
	p.next()

	fields := make([]ast.FieldDeclExpr, 0)
	for {
		p.skipLineBreaks()

		if p.peek() == lexer.EndType {
			break
		}

		fields = append(fields, parseFieldDeclExpr(p))
	}

	p.expect(lexer.EndType)

	return ast.TypeStmt{
		Identifier: identifier,
		Fields:     fields,
	}
}

func parseCallStmt(p *parser) ast.Stmt {
	p.expect(lexer.Call)
	funcName := p.expect(lexer.Identifier).Value
	p.expect(lexer.LParen)

	args := make([]ast.Expr, 0)
	for {
		args = append(args, parseExpr(p, comma))
		if p.peek() != lexer.Comma {
			break
		}
		p.next()
	}

	p.expect(lexer.RParen)

	return ast.CallStmt{
		Func: funcName,
		Args: args,
	}
}
