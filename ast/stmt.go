package ast

type BlockStmt []Stmt

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

type CallStmt struct {
	Identifier string
	Args       []Expr
}

func (n CallStmt) Stmt() {}

type DeclareStmt struct {
	Identifier string
	Lib        string
	Alias      string
	Args       []ArgExpr
	ReturnType TypeExpr
}

func (n DeclareStmt) Stmt() {}

type FunctionStmt struct {
	Public     bool
	Identifier string
	Args       []ArgExpr
	ReturnType TypeExpr
	Body       BlockStmt
}

func (n FunctionStmt) Stmt() {}

type ElseIfStmt struct {
	Condition Expr
	Body      BlockStmt
}

func (n ElseIfStmt) Stmt() {}

type IfStmt struct {
	Condition Expr
	Body      BlockStmt
	ElseIf    []ElseIfStmt
	Else      BlockStmt
}

func (n IfStmt) Stmt() {}
