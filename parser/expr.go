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
		panic(p.unexpected())
	}

	left := nfn(p)
	for tables.bp[p.peek()] > bp {
		kind = p.peek()

		lfn, ok := tables.led[kind]
		if !ok {
			panic(p.unexpected())
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
		return parseSymbolExpr(p)

	default:
		panic(p.unexpected())
	}
}

func parseSymbolExpr(p *parser) ast.Expr {
	identifier := p.expect(lexer.Identifier).Value
	if p.peek() == lexer.LParen {
		return parseCallExpr(p, identifier)
	}

	return ast.SymbolExpr{Name: identifier}
}

func parseCallExpr(p *parser, identifier string) ast.Expr {
	p.expect(lexer.LParen)

	args := []ast.Expr{}

	for p.peek() != lexer.RParen {
		args = append(args, parseExpr(p, assignment))
		if p.peek() == lexer.Comma {
			p.next()
		}
	}

	p.expect(lexer.RParen)

	return ast.CallExpr{
		Identifier: identifier,
		Args:       args,
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

var dataTypeMap = map[lexer.Kind]ast.DataType{
	lexer.BooleanType: ast.DtBoolean,
	lexer.ByteType:    ast.DtByte,
	lexer.IntegerType: ast.DtInteger,
	lexer.LongType:    ast.DtLong,
	lexer.SingleType:  ast.DtSingle,
	lexer.DoubleType:  ast.DtDouble,
	lexer.StringType:  ast.DtString,
}

func parseTypeExpr(p *parser) ast.TypeExpr {
	cur := p.next()
	if !cur.IsDataType() {
		if cur.Kind == lexer.Identifier {
			return ast.TypeExpr{
				Type:       ast.DtUserDefined,
				TypeName:   cur.Value,
				IsFixedLen: false,
				Len:        nil,
			}
		}
		panic("expected data type")
	}

	dataType, ok := dataTypeMap[cur.Kind]
	if !ok {
		panic(fmt.Errorf("unexpected token %s", lexer.TokenKindString(cur.Kind)))
	}

	if dataType == ast.DtString {
		if p.peek() == lexer.Multiply {
			p.next()
			len := parseExpr(p, assignment)
			return ast.TypeExpr{
				Type:       dataType,
				IsFixedLen: true,
				Len:        len,
			}
		}
	}

	return ast.TypeExpr{
		Type:       dataType,
		IsFixedLen: false,
		Len:        nil,
	}
}

func parseGroupExpr(p *parser) ast.Expr {
	p.expect(lexer.LParen)
	expr := parseExpr(p, defaultBindingPower)
	p.expect(lexer.RParen)
	return expr
}
