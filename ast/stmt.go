package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) Stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (n ExprStmt) Stmt() {}

// Public Const START_X = MAX_MAPX / 2
type ConstDeclStmt struct {
	Public     bool
	Identifier string
	Value      Expr
}

func (n ConstDeclStmt) Stmt() {}

// Dim IPMask As String
// Public SpawnSeconds As Long
type VarDeclStmt struct {
	Identifier string
	Type       TypeExpr
	Value      Expr
	Public     bool
}

func (n VarDeclStmt) Stmt() {}
