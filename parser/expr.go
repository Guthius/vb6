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
	t := p.next()
	if !t.IsDataType() {
		panic("expected data type")
	}

	dataType, ok := dataTypeMap[t.Kind]
	if !ok {
		panic(fmt.Errorf("unexpected token %s", lexer.TokenKindString(t.Kind)))
	}

	if dataType == ast.DtString {
		if p.peek() == lexer.Multiply {
			p.next()
			token := p.expect(lexer.Number)
			len, err := strconv.Atoi(token.Value)
			if err != nil {
				panic(fmt.Errorf("invalid number %s", token.Value))
			}
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
		Len:        0,
	}
}
