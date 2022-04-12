package bytecode

type IElement interface{ GetType() string }
type Element struct {
	Type string `json:"type"`
	IElement
}

type Position struct {
	Type string `json:"type"`
	IElement
	Row      int    `json:"row"`
	Col      int    `json:"col"`
	Filename string `json:"filename"`
}

type Token struct {
	Type string `json:"type"`
	IElement
	TokenType string `json:"tokenType"`
	Value     string `json:"value"`
}

type IStatementNode interface{ GetType() string }
type StatementNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
}

type IExpressionNode interface{ GetType() string }
type ExpressionNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
}

type TypeNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Token Token `json:"token"`
}

type TypedDeclNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token    `json:"target"`
	ValueType TypeNode `json:"valueType"`
}

type FuncDefNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token            `json:"target"`
	ValueType TypeNode         `json:"valueType"`
	Params    []TypedDeclNode  `json:"params"`
	Body      []IStatementNode `json:"body"`
}

type ReturnNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Value IExpressionNode `json:"value"`
}

type WhileNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Condition IExpressionNode  `json:"value"`
	Body      []IStatementNode `json:"body"`
}

type IfElseNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Condition IExpressionNode  `json:"value"`
	Truthy    []IStatementNode `json:"truthy"`
	Falsy     []IStatementNode `json:"falsy"`
}

type IfNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Condition IExpressionNode  `json:"value"`
	Body      []IStatementNode `json:"body"`
}

type VarInitNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token           `json:"target"`
	ValueType TypeNode        `json:"valueType"`
	Value     IExpressionNode `json:"value"`
}

type VarDeclNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token    `json:"target"`
	ValueType TypeNode `json:"valueType"`
}

type AssignNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target Token           `json:"target"`
	Value  IExpressionNode `json:"value"`
}

type NotEqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type EqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type GreaterThanOrEqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type LessThanOrEqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type GreaterThanNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type LessThanNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type AddNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type SubNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type MulNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type DivNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type ModNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type ExpNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type FuncCallNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target IExpressionNode   `json:"target"`
	Args   []IExpressionNode `json:"args"`
}

type IntNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Token Token `json:"token"`
}

type VarNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Token Token `json:"token"`
}

func (e Element) GetType() string                { return e.Type }
func (e Position) GetType() string               { return e.Type }
func (e Token) GetType() string                  { return e.Type }
func (e StatementNode) GetType() string          { return e.Type }
func (e ExpressionNode) GetType() string         { return e.Type }
func (e TypeNode) GetType() string               { return e.Type }
func (e TypedDeclNode) GetType() string          { return e.Type }
func (e FuncDefNode) GetType() string            { return e.Type }
func (e ReturnNode) GetType() string             { return e.Type }
func (e WhileNode) GetType() string              { return e.Type }
func (e IfElseNode) GetType() string             { return e.Type }
func (e IfNode) GetType() string                 { return e.Type }
func (e VarInitNode) GetType() string            { return e.Type }
func (e VarDeclNode) GetType() string            { return e.Type }
func (e AssignNode) GetType() string             { return e.Type }
func (e NotEqualNode) GetType() string           { return e.Type }
func (e EqualNode) GetType() string              { return e.Type }
func (e GreaterThanOrEqualNode) GetType() string { return e.Type }
func (e LessThanOrEqualNode) GetType() string    { return e.Type }
func (e GreaterThanNode) GetType() string        { return e.Type }
func (e LessThanNode) GetType() string           { return e.Type }
func (e AddNode) GetType() string                { return e.Type }
func (e SubNode) GetType() string                { return e.Type }
func (e MulNode) GetType() string                { return e.Type }
func (e DivNode) GetType() string                { return e.Type }
func (e ModNode) GetType() string                { return e.Type }
func (e ExpNode) GetType() string                { return e.Type }
func (e FuncCallNode) GetType() string           { return e.Type }
func (e IntNode) GetType() string                { return e.Type }
func (e VarNode) GetType() string                { return e.Type }
