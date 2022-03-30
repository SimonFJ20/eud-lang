package parser

type Position interface {
	row() int
	col() int
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
	tokenType() TokenType
	value() string
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
	nodeType() NodeType
}

type Statement interface {
	Node
}

type ImportNode interface {
	Statement
}

type TypeDefNode interface {
	Statement
	target() IdentifierNode
}

type TraitDefNode interface {
	Statement
	target() IdentifierNode
}

type StructDefNode interface {
	Statement
	target() IdentifierNode
}

type FuncDefNode interface {
	Statement
	target() IdentifierNode
	defType() Type
	params() []TypedDeclarationNode
	body() []Statement
}

type ReturnValueNode interface {
	Statement
	value() Expression
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
	declaration() Expression
	condition() Expression
	iteration() Expression
	body() Expression
}

type WhileNode interface {
	Statement
	condition() Expression
	body() Expression
}

type BreakNode interface {
	Statement
}

type ContinueNode interface {
	Statement
}

type IfNode interface {
	Statement
	condition() Expression
	body() Expression
}

type IfElsNode interface {
	Statement
	condition() Expression
	truthy() Expression
	falsy() Expression
}

type InferredInitializationNode interface {
	Statement
	target() Declareable
	source() Expression
}

type TypedInitializationNode interface {
	Statement
	defType() Type
	target() Declareable
	source() Expression
}

type Expression interface {
	Node
}

type TypedObjectLiteralNode interface {
	Expression
	ObjectLiteralNode
	defType() Type
}

type ObjectLiteralNode interface {
	Expression
	values() []KeyValuePair
}

type KeyValuePair interface {
	key() Token
	value() Expression
}

type TypedArrayLiteralNode interface {
	Expression
	ArrayLiteralNode
	defType() Type
}

type ArrayLiteralNode interface {
	Expression
	values() []Expression
}

type LambdaLiteralNode interface {
	defType() Type
	Expression
	params() []Expression
	body() Expression
}

type TypedDeclarationNode interface {
	defType() Type
	target() Declareable
}

type ArrayTypeNode interface {
	Type
	defType() Type
	values() []Declareable
}

type ObjectTypeNode interface {
	Type
	defType() Type
	fields() []TypedDeclarationNode
}

type LambdaTypeNode interface {
	Type
	defType() Type
	params() []TypedDeclarationNode
}

type KeywordTypeNode interface {
	Type
	value() Token
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
	targets() []Declareable
}

type UnpackedObjectNode interface {
	Declareable
	targets() []Declareable
}

type RenamedIdentifierNode interface {
	Declareable
	source() []IdentifierNode
	target() []IdentifierNode
}

type AssignmentNode interface {
	Expression
	target() Expression
	source() Expression
}

type TernaryNode interface {
	Expression
	condition() Expression
	truthy() Expression
	falsy() Expression
}

type BinaryOperationNode interface {
	Expression
	left() Expression
	right() Expression
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
	source() Expression
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
	source() Expression
	specifier() Token
}

type ComputedMemberNode interface {
	Expression
	source() Expression
	specifier() Expression
}

type FuncCallNode interface {
	Expression
	source() Expression
	args() []Expression
}

type IntLiteralNode interface {
	Expression
	value() Token
}

type FloatLiteralNode interface {
	Expression
	value() Token
}

type CharLiteralNode interface {
	Expression
	value() Token
}

type StringLiteralNode interface {
	Expression
	value() Token
}

type VarAccessNode interface {
	Expression
	IdentifierNode
}

type IdentifierNode interface {
	Declareable
	value() Token
}
