package ast

type Stmt interface {
	Stmt()
}

type Expr interface {
	Expr()
}
