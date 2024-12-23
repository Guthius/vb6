package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) Stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (n ExprStmt) Stmt() {}
