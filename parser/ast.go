package parser

type Position interface {
	Row() int
	Col() int
}

type TokenType int

const (
	IDENTIFIER TokenType = iota
	KEYWORD
	INT
	FLOAT
	CHAR
	STRING
)

type Token interface {
	Position
	TokenType() TokenType
	Value() string
}

type NodeType int

const (
	Import NodeType = iota
	TypeDef
	TraitDef
	StructDef
	FuncDef
	Return
	Match
	Switch
	For
	While
	Break
	Continue
	If
	IfEls
	InferredInitialization
	TypedInitialization
	TypedObjectLiteral
	ObjectLiteral
	TypedArrayLiteral
	ArrayLiteral
	LambdaLiteral
	Assignment
	Ternary
	LogicalOr
	LogicalAnd
	BitwiseOr
	BitwiseXor
	BitwiseAnd
	Equality
	Inequality
	CompareLT
	CompareLTE
	CompareGT
	CompareGTE
	BitShiftLeft
	BitShiftRight
	Addition
	Subtraction
	Multiplication
	Division
	Remainder
	Exponentation
	LogicalNot
	BitwiseNot
	UnaryPlus
	UnaryNegation
	PreIncrement
	PreDecrement
	PostIncrement
	PostDecrement
	MemberAccess
	ComputedMember
	FuncCall
	IntLiteral
	FloatLiteral
	CharLiteral
	StringLiteral
	VarAccess
)

type Node interface {
	Position
	NodeType() NodeType
}

type Statement interface {
	Node
}

type ImportNode interface {
	Statement
}

type TypeDefNode interface {
	Statement
	Target() IdentifierNode
}

type TraitDefNode interface {
	Statement
	Target() IdentifierNode
}

type StructDefNode interface {
	Statement
	Target() IdentifierNode
}

type FuncDefNode interface {
	Statement
	Target() IdentifierNode
	DefType() Type
	Params() []TypedDeclarationNode
	Body() []Statement
}

type ReturnValueNode interface {
	Statement
	Value() Expression
}

type ReturnNode interface {
	Statement
}

type MatchNode interface {
	Statement
}

type SwitchNode interface {
	Statement
}

type ForNode interface {
	Statement
	Declaration() Expression
	Condition() Expression
	Iteration() Expression
	Body() Expression
}

type WhileNode interface {
	Statement
	Condition() Expression
	Body() Expression
}

type BreakNode interface {
	Statement
}

type ContinueNode interface {
	Statement
}

type IfNode interface {
	Statement
	Condition() Expression
	Body() Expression
}

type IfElsNode interface {
	Statement
	Condition() Expression
	Truthy() Expression
	Falsy() Expression
}

type InferredInitializationNode interface {
	Statement
	Target() Declareable
	Source() Expression
}

type TypedInitializationNode interface {
	Statement
	DefType() Type
	Target() Declareable
	Source() Expression
}

type Expression interface {
	Node
}

type TypedObjectLiteralNode interface {
	Expression
	ObjectLiteralNode
	DefType() Type
}

type ObjectLiteralNode interface {
	Expression
	Values() []KeyValuePair
}

type KeyValuePair interface {
	key() Token
	Value() Expression
}

type TypedArrayLiteralNode interface {
	Expression
	ArrayLiteralNode
	DefType() Type
}

type ArrayLiteralNode interface {
	Expression
	Values() []Expression
}

type LambdaLiteralNode interface {
	DefType() Type
	Expression
	Params() []Expression
	Body() Expression
}

type TypedDeclarationNode interface {
	DefType() Type
	Target() Declareable
}

type ArrayTypeNode interface {
	Type
	DefType() Type
	Values() []Declareable
}

type ObjectTypeNode interface {
	Type
	DefType() Type
	Fields() []TypedDeclarationNode
}

type LambdaTypeNode interface {
	Type
	DefType() Type
	Params() []TypedDeclarationNode
}

type KeywordTypeNode interface {
	Type
	Value() Token
}

type IdentifierTypeNode interface {
	IdentifierNode
	Type
}

type Type interface{}

type Declareable interface {
}

type UnpackedArrayNode interface {
	Declareable
	Targets() []Declareable
}

type UnpackedObjectNode interface {
	Declareable
	Targets() []Declareable
}

type RenamedIdentifierNode interface {
	Declareable
	Source() []IdentifierNode
	Target() []IdentifierNode
}

type AssignmentNode interface {
	Expression
	Target() Expression
	Source() Expression
}

type TernaryNode interface {
	Expression
	Condition() Expression
	Truthy() Expression
	Falsy() Expression
}

type BinaryOperationNode interface {
	Expression
	Left() Expression
	Right() Expression
}

type LogicalOrNode = BinaryOperationNode
type LogicalAndNode = BinaryOperationNode
type BitwiseOrNode = BinaryOperationNode
type BitwiseXorNode = BinaryOperationNode
type BitwiseAndNode = BinaryOperationNode
type EqualityNode = BinaryOperationNode
type InequalityNode = BinaryOperationNode
type CompareLTNode = BinaryOperationNode
type CompareLTENode = BinaryOperationNode
type CompareGTNode = BinaryOperationNode
type CompareGTENode = BinaryOperationNode
type BitShiftLeftNode = BinaryOperationNode
type BitShiftRightNode = BinaryOperationNode
type AdditionNode = BinaryOperationNode
type SubtractionNode = BinaryOperationNode
type MultiplicationNode = BinaryOperationNode
type DivisionNode = BinaryOperationNode
type RemainderNode = BinaryOperationNode
type ExponentationNode = BinaryOperationNode

type UnaryOperationNode interface {
	Expression
	Source() Expression
}

type LogicalNotNode = UnaryOperationNode
type BitwiseNotNode = UnaryOperationNode
type UnaryPlusNode = UnaryOperationNode
type UnaryNegationNode = UnaryOperationNode
type PreIncrementNode = UnaryOperationNode
type PreDecrementNode = UnaryOperationNode
type PostIncrementNode = UnaryOperationNode
type PostDecrementNode = UnaryOperationNode

type MemberAccessNode interface {
	Expression
	Source() Expression
	Specifier() Token
}

type ComputedMemberNode interface {
	Expression
	Source() Expression
	Specifier() Expression
}

type FuncCallNode interface {
	Expression
	Source() Expression
	Args() []Expression
}

type IntLiteralNode interface {
	Expression
	Value() Token
}

type FloatLiteralNode interface {
	Expression
	Value() Token
}

type CharLiteralNode interface {
	Expression
	Value() Token
}

type StringLiteralNode interface {
	Expression
	Value() Token
}

type VarAccessNode interface {
	Expression
	IdentifierNode
}

type IdentifierNode interface {
	Declareable
	Value() Token
}
