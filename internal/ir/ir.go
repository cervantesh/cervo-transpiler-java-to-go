package ir

type File struct {
	PackageName string
	Classes     []Class
	Funcs       []Func
}

type Class struct {
	Name    string
	Methods []Func
}

type Func struct {
	Name       string
	Params     []Param
	ReturnType Type
	Body       []Stmt
	Static     bool
}

type Param struct {
	Name string
	Type Type
}

type Stmt interface {
	stmtNode()
}

type Expr interface {
	exprNode()
}

type VarDeclStmt struct {
	Name  string
	Type  Type
	Value Expr
}

func (VarDeclStmt) stmtNode() {}

type AssignStmt struct {
	Name  string
	Op    string
	Value Expr
}

func (AssignStmt) stmtNode() {}

type ExprStmt struct {
	Expr Expr
}

func (ExprStmt) stmtNode() {}

type ReturnStmt struct {
	Value Expr
}

func (ReturnStmt) stmtNode() {}

type IfStmt struct {
	Cond Expr
	Then []Stmt
	Else []Stmt
}

func (IfStmt) stmtNode() {}

type WhileStmt struct {
	Cond Expr
	Body []Stmt
}

func (WhileStmt) stmtNode() {}

type ForStmt struct {
	Init Stmt
	Cond Expr
	Post Stmt
	Body []Stmt
}

func (ForStmt) stmtNode() {}

type LiteralExpr struct {
	Value string
	Type  Type
}

func (LiteralExpr) exprNode() {}

type NameExpr struct {
	Name string
}

func (NameExpr) exprNode() {}

type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
}

func (BinaryExpr) exprNode() {}

type UnaryExpr struct {
	Op   string
	Expr Expr
}

func (UnaryExpr) exprNode() {}

type CallExpr struct {
	Target string
	Name   string
	Args   []Expr
}

func (CallExpr) exprNode() {}
