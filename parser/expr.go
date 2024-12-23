package parser

import (
	"fmt"
	"strconv"

	"github.com/guthius/vb6/ast"
	"github.com/guthius/vb6/lexer"
)

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	kind := p.peek()

	nfn, ok := tables.nud[kind]
	if !ok {
		panic(fmt.Errorf("unexpected token %s", lexer.TokenKindString(kind)))
	}

	left := nfn(p)
	for tables.bp[p.peek()] > bp {
		kind = p.peek()

		lfn, ok := tables.led[kind]
		if !ok {
			panic(fmt.Errorf("unexpected token %s", lexer.TokenKindString(kind)))
		}

		left = lfn(p, left, tables.bp[p.peek()])
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	kind := p.peek()
	switch kind {
	case lexer.Number:
		t := p.next()
		v, err := strconv.ParseFloat(t.Value, 64)
		if err != nil {
			panic(fmt.Errorf("invalid number %s", t.Value))
		}
		return ast.NumberExpr{Value: v}

	case lexer.String:
		t := p.next()
		return ast.StringExpr{Value: t.Value}

	case lexer.Identifier:
		t := p.next()
		return ast.SymbolExpr{Name: t.Value}

	default:
		panic(fmt.Errorf("unexpected token %s", lexer.TokenKindString(kind)))
	}
}

func parseBinaryExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operator := p.next()
	right := parseExpr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
