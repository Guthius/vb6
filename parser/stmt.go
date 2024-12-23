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
	p.expect(lexer.Const)
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

	/*
		Public Item(1 To MAX_ITEMS) As ItemRec
	*/

	identifier := p.expect(lexer.Identifier).Value

	var ranges []ast.RangeExpr
	if p.peek() == lexer.LParen {
		ranges = parseRangeExpr(p)
	}

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
		IsArray:    ranges != nil,
		Ranges:     ranges,
		Value:      value,
	}
}

func parseRangeExpr(p *parser) []ast.RangeExpr {
	ranges := make([]ast.RangeExpr, 0)

	p.expect(lexer.LParen)
	for p.peek() != lexer.RParen {
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
		Identifier: funcName,
		Args:       args,
	}
}

func parseDeclareStmt(p *parser) ast.Stmt {
	p.expect(lexer.Declare)
	p.expect(lexer.Function)
	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.Lib)
	lib := p.expect(lexer.String).Value
	p.expect(lexer.Alias)
	alias := p.expect(lexer.String).Value
	p.expect(lexer.LParen)
	args := parseArgList(p)
	p.expect(lexer.RParen)

	var returnType ast.TypeExpr
	if p.peek() == lexer.As {
		p.next()
		returnType = parseTypeExpr(p)
	}

	p.expectOrEof(lexer.LineBreak)

	return ast.DeclareStmt{
		Identifier: identifier,
		Lib:        lib,
		Alias:      alias,
		Args:       args,
		ReturnType: returnType,
	}
}

func parseArgList(p *parser) []ast.ArgExpr {
	args := make([]ast.ArgExpr, 0)
	for p.peek() != lexer.RParen {
		byRef := false
		switch p.peek() {
		case lexer.ByVal:
			p.next()
		case lexer.ByRef:
			byRef = true
			p.next()
		}

		name := p.expect(lexer.Identifier).Value
		p.expect(lexer.As)

		argType := parseTypeExpr(p)
		args = append(args, ast.ArgExpr{
			ByRef:      byRef,
			Identifier: name,
			Type:       argType,
		})

		if p.peek() != lexer.Comma {
			break
		}

		p.next()
	}

	return args
}

func parseFunctionStmt(p *parser) ast.Stmt {
	p.expect(lexer.Function)
	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.LParen)
	args := parseArgList(p)
	p.expect(lexer.RParen)
	p.expect(lexer.As)
	returnType := parseTypeExpr(p)
	p.expectOrEof(lexer.LineBreak)
	body := parseBlockStmt(p, lexer.EndFunction)
	p.expect(lexer.EndFunction)

	return ast.FunctionStmt{
		Public:     true,
		Identifier: identifier,
		Args:       args,
		ReturnType: returnType,
		Body:       body,
	}
}

func parseDimStmt(p *parser) ast.Stmt {
	p.expect(lexer.Dim)
	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.As)
	dataType := parseTypeExpr(p)
	p.expectOrEof(lexer.LineBreak)

	return ast.VarDeclStmt{
		Public:     false,
		Identifier: identifier,
		Type:       dataType,
	}
}

func parseBlockStmt(p *parser, end ...lexer.Kind) ast.BlockStmt {
	body := make(ast.BlockStmt, 0)

	for !p.isEof() {
		p.skipLineBreaks()
		if p.isEof() {
			break
		}

		for _, kind := range end {
			if p.peek() == kind {
				return body
			}
		}

		body = append(body, p.parseStmt())
	}

	return body
}

func parseIfStmt(p *parser) ast.Stmt {
	p.expect(lexer.If)
	condition := parseExpr(p, assignment)
	p.expect(lexer.Then)
	p.expectOrEof(lexer.LineBreak)
	body := parseBlockStmt(p, lexer.ElseIf, lexer.Else, lexer.EndIf)

	var elseIf []ast.ElseIfStmt
	if p.peek() == lexer.ElseIf {
		elseIf = make([]ast.ElseIfStmt, 0)
		for p.peek() == lexer.ElseIf {
			p.next()
			elseIfCondition := parseExpr(p, assignment)
			p.expect(lexer.Then)
			p.expectOrEof(lexer.LineBreak)
			elseBody := parseBlockStmt(p, lexer.ElseIf, lexer.Else, lexer.EndIf)
			elseIf = append(elseIf, ast.ElseIfStmt{
				Condition: elseIfCondition,
				Body:      elseBody,
			})
		}
	}

	var elseBody ast.BlockStmt
	if p.peek() == lexer.Else {
		p.next()
		elseBody = parseBlockStmt(p, lexer.ElseIf, lexer.Else, lexer.EndIf)
	}

	p.expect(lexer.EndIf)

	return ast.IfStmt{
		Condition: condition,
		Body:      body,
		ElseIf:    elseIf,
		Else:      elseBody,
	}
}

func parseExitFunctionStmt(p *parser) ast.Stmt {
	p.expect(lexer.ExitFunction)
	p.expectOrEof(lexer.LineBreak)
	return ast.ExitFunctionStmt{}
}

func parseForStmt(p *parser) ast.Stmt {
	p.expect(lexer.For)
	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.Equal)
	start := parseExpr(p, assignment)
	p.expect(lexer.To)
	end := parseExpr(p, assignment)
	var step ast.Expr
	if p.peek() == lexer.Step {
		p.next()
		step = parseExpr(p, assignment)
	}
	p.expectOrEof(lexer.LineBreak)
	body := parseBlockStmt(p, lexer.Next)
	p.expect(lexer.Next)
	next := p.expect(lexer.Identifier).Value
	if next != identifier {
		panic("Next identifier must match For identifier")
	}
	p.expectOrEof(lexer.LineBreak)

	return ast.ForStmt{
		Identifier: identifier,
		Start:      start,
		End:        end,
		Step:       step,
		Body:       body,
	}
}

func parseSubStmt(p *parser) ast.Stmt {
	p.expect(lexer.Sub)
	identifier := p.expect(lexer.Identifier).Value
	p.expect(lexer.LParen)
	args := parseArgList(p)
	p.expect(lexer.RParen)
	p.expectOrEof(lexer.LineBreak)
	body := parseBlockStmt(p, lexer.EndSub)
	p.expect(lexer.EndSub)

	return ast.SubStmt{
		Identifier: identifier,
		Args:       args,
		Body:       body,
	}
}
