package parser

import (
	"github.com/guthius/vb6/ast"
	"github.com/guthius/vb6/lexer"
)

type bindingPower int

const (
	defaultBindingPower bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmtHandler func(*parser) ast.Stmt
type nudHandler func(*parser) ast.Expr
type ledHandler func(*parser, ast.Expr, bindingPower) ast.Expr

var tables struct {
	stmt map[lexer.Kind]stmtHandler
	nud  map[lexer.Kind]nudHandler
	led  map[lexer.Kind]ledHandler
	bp   map[lexer.Kind]bindingPower
}

// init initializes the parser lookup tables.
func init() {
	tables.stmt = make(map[lexer.Kind]stmtHandler)
	tables.nud = make(map[lexer.Kind]nudHandler)
	tables.led = make(map[lexer.Kind]ledHandler)
	tables.bp = make(map[lexer.Kind]bindingPower)

	// Logical operators
	led(lexer.And, logical, parseBinaryExpr)
	led(lexer.Or, logical, parseBinaryExpr)

	// Relational operators
	led(lexer.Equal, relational, parseBinaryExpr)
	led(lexer.NotEqual, relational, parseBinaryExpr)
	led(lexer.GreaterThan, relational, parseBinaryExpr)
	led(lexer.GreaterThanOrEqual, relational, parseBinaryExpr)
	led(lexer.LessThan, relational, parseBinaryExpr)
	led(lexer.LessThanOrEqual, relational, parseBinaryExpr)

	// Additive operators
	led(lexer.Add, additive, parseBinaryExpr)
	led(lexer.Subtract, additive, parseBinaryExpr)

	// Multiplicative operators
	led(lexer.Multiply, multiplicative, parseBinaryExpr)
	led(lexer.Divide, multiplicative, parseBinaryExpr)
	led(lexer.DivideInt, multiplicative, parseBinaryExpr)
	led(lexer.Modulus, multiplicative, parseBinaryExpr)

	// Literals and identifiers
	nud(lexer.Number, primary, parsePrimaryExpr)
	nud(lexer.String, primary, parsePrimaryExpr)
	nud(lexer.Identifier, primary, parsePrimaryExpr)

	// Statements
	stmt(lexer.Public, parseDeclStmt)
	stmt(lexer.Private, parseDeclStmt)
	stmt(lexer.Const, parsePrivateConstDeclStmt)
}

func stmt(king lexer.Kind, handler stmtHandler) {
	tables.stmt[king] = handler
}

func nud(king lexer.Kind, bp bindingPower, handler nudHandler) {
	tables.nud[king] = handler
	tables.bp[king] = bp
}

func led(king lexer.Kind, bp bindingPower, handler ledHandler) {
	tables.led[king] = handler
	tables.bp[king] = bp
}
