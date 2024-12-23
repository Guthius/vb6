package ast

type DataType int

const (
	DtBoolean DataType = iota
	DtByte
	DtInteger
	DtLong
	DtSingle
	DtDouble
	DtString
)

type TypeExpr struct {
	Type       DataType
	IsFixedLen bool
	Len        int
}

func (n TypeExpr) Expr() {}
