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
	DtUserDefined
)

type TypeExpr struct {
	Type       DataType
	TypeName   string // only for DtUserDefined
	IsFixedLen bool
	Len        Expr
}

func (n TypeExpr) Expr() {}
