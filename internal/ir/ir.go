package ir

type File struct {
	Path        string
	PackageName string
	Classes     []Class
	Funcs       []Func
}

type Class struct {
	Name    string
	Symbol  string
	Span    Span
	Methods []Func
}

type Func struct {
	Name       string
	Symbol     string
	Span       Span
	Params     []Param
	ReturnType Type
	Body       []Stmt
	Static     bool
}

type Param struct {
	Name string
	Type Type
	Span Span
}

type Span struct {
	File   string
	Line   int
	Column int
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
	Span  Span
	Value Expr
}

func (VarDeclStmt) stmtNode() {}

type AssignStmt struct {
	Name  string
	Op    string
	Span  Span
	Value Expr
}

func (AssignStmt) stmtNode() {}

type ExprStmt struct {
	Span Span
	Expr Expr
}

func (ExprStmt) stmtNode() {}

type ReturnStmt struct {
	Span  Span
	Value Expr
}

func (ReturnStmt) stmtNode() {}

type IfStmt struct {
	Span Span
	Cond Expr
	Then []Stmt
	Else []Stmt
}

func (IfStmt) stmtNode() {}

type WhileStmt struct {
	Span Span
	Cond Expr
	Body []Stmt
}

func (WhileStmt) stmtNode() {}

type ForStmt struct {
	Span Span
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
	Type Type
	Span Span
}

func (NameExpr) exprNode() {}

type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
	Type  Type
}

func (BinaryExpr) exprNode() {}

type UnaryExpr struct {
	Op   string
	Expr Expr
	Type Type
}

func (UnaryExpr) exprNode() {}

type CallExpr struct {
	Target string
	Name   string
	Args   []Expr
	Type   Type
}

func (CallExpr) exprNode() {}
