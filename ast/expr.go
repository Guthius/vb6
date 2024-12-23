package ast

import (
	"github.com/guthius/vb6/lexer"
)

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) Expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) Expr() {}

type SymbolExpr struct {
	Name string
}

func (n SymbolExpr) Expr() {}

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) Expr() {}
