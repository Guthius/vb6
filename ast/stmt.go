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
	IsArray    bool
	Ranges     []RangeExpr
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

type ExitFunctionStmt struct{}

func (n ExitFunctionStmt) Stmt() {}

type ForStmt struct {
	Identifier string
	Start      Expr
	End        Expr
	Step       Expr
	Body       BlockStmt
}

func (n ForStmt) Stmt() {}

type SubStmt struct {
	Identifier string
	Args       []ArgExpr
	Body       BlockStmt
}

func (n SubStmt) Stmt() {}

type OptionExplicitStmt struct{}

func (n OptionExplicitStmt) Stmt() {}
