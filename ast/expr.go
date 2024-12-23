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

type RangeExpr struct {
	LBound Expr
	UBound Expr
}

func (n RangeExpr) Expr() {}

type FieldDeclExpr struct {
	Identifier string
	Type       TypeExpr
	IsArray    bool
	Ranges     []RangeExpr
}

func (n FieldDeclExpr) Expr() {}

type CallExpr struct {
	Identifier string
	Args       []Expr
}

func (n CallExpr) Expr() {}

type ArgExpr struct {
	ByRef      bool
	Identifier string
	Type       TypeExpr
}

func (n ArgExpr) Expr() {}
