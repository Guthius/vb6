package parser

import (
	"github.com/guthius/vb6/ast"
	"github.com/guthius/vb6/lexer"
)

func (p *parser) parseStmt() ast.Stmt {
	kind := p.peek()

	sfn, ok := tables.stmt[kind]
	if !ok {
		expr := parseExpr(p, defaultBindingPower)
		p.expect(lexer.LineBreak, lexer.EOF)
		return ast.ExprStmt{Expr: expr}
	}

	return sfn(p)
}
