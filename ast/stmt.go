package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) Stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (n ExprStmt) Stmt() {}

type ConstDeclStmt struct {
	Public     bool
	Identifier string
	Value      Expr
}

func (n ConstDeclStmt) Stmt() {}

type VarDeclStmt struct {
	Public     bool
	Identifier string
	Type       TypeExpr
	Value      Expr
}

func (n VarDeclStmt) Stmt() {}

type TypeStmt struct {
	Identifier string
	Fields     []FieldDeclExpr
}

func (n TypeStmt) Stmt() {}
