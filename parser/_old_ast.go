package parser

type AST_Position interface {
	AST_Row() int
	AST_Col() int
}

type AST_TokenType int

const (
	AST_IDENTIFIER AST_TokenType = iota
	AST_KEYWORD
	AST_INT
	AST_FLOAT
	AST_CHAR
	AST_STRING
)

type AST_Token interface {
	AST_Position
	AST_TokenType() AST_TokenType
	AST_Value() string
}

type AST_NodeType int

const (
	AST_Import AST_NodeType = iota
	AST_TypeDef
	AST_TraitDef
	AST_StructDef
	AST_FuncDef
	AST_Return
	AST_Match
	AST_Switch
	AST_For
	AST_While
	AST_Break
	AST_Continue
	AST_If
	AST_IfEls
	AST_InferredInitialization
	AST_TypedInitialization
	AST_TypedObjectLiteral
	AST_ObjectLiteral
	AST_TypedArrayLiteral
	AST_ArrayLiteral
	AST_LambdaLiteral
	AST_Assignment
	AST_Ternary
	AST_LogicalOr
	AST_LogicalAnd
	AST_BitwiseOr
	AST_BitwiseXor
	AST_BitwiseAnd
	AST_Equality
	AST_Inequality
	AST_CompareLT
	AST_CompareLTE
	AST_CompareGT
	AST_CompareGTE
	AST_BitShiftLeft
	AST_BitShiftRight
	AST_Addition
	AST_Subtraction
	AST_Multiplication
	AST_Division
	AST_Remainder
	AST_Exponentation
	AST_LogicalNot
	AST_BitwiseNot
	AST_UnaryPlus
	AST_UnaryNegation
	AST_PreIncrement
	AST_PreDecrement
	AST_PostIncrement
	AST_PostDecrement
	AST_MemberAccess
	AST_ComputedMember
	AST_FuncCall
	AST_IntLiteral
	AST_FloatLiteral
	AST_CharLiteral
	AST_StringLiteral
	AST_VarAccess
)

type AST_Node interface {
	AST_Position
	AST_NodeType() AST_NodeType
}

type AST_Statement interface {
	AST_Node
}

type AST_ImportNode interface {
	AST_Statement
}

type AST_TypeDefNode interface {
	AST_Statement
	AST_Target() AST_IdentifierNode
}

type AST_TraitDefNode interface {
	AST_Statement
	AST_Target() AST_IdentifierNode
}

type AST_StructDefNode interface {
	AST_Statement
	AST_Target() AST_IdentifierNode
}

type AST_FuncDefNode interface {
	AST_Statement
	AST_Target() AST_IdentifierNode
	AST_DefType() AST_Type
	AST_Params() []AST_TypedDeclarationNode
	AST_Body() []AST_Statement
}

type AST_ReturnValueNode interface {
	AST_Statement
	AST_Value() AST_Expression
}

type AST_ReturnNode interface {
	AST_Statement
}

type AST_MatchNode interface {
	AST_Statement
}

type AST_SwitchNode interface {
	AST_Statement
}

type AST_ForNode interface {
	AST_Statement
	AST_Declaration() AST_Expression
	AST_Condition() AST_Expression
	AST_Iteration() AST_Expression
	AST_Body() AST_Expression
}

type AST_WhileNode interface {
	AST_Statement
	AST_Condition() AST_Expression
	AST_Body() AST_Expression
}

type AST_BreakNode interface {
	AST_Statement
}

type AST_ContinueNode interface {
	AST_Statement
}

type AST_IfNode interface {
	AST_Statement
	AST_Condition() AST_Expression
	AST_Body() AST_Expression
}

type AST_IfElsNode interface {
	AST_Statement
	AST_Condition() AST_Expression
	AST_Truthy() AST_Expression
	AST_Falsy() AST_Expression
}

type AST_InferredInitializationNode interface {
	AST_Statement
	AST_Target() AST_Declareable
	AST_Source() AST_Expression
}

type AST_TypedInitializationNode interface {
	AST_Statement
	AST_DefType() AST_Type
	AST_Target() AST_Declareable
	AST_Source() AST_Expression
}

type AST_Expression interface {
	AST_Node
}

type AST_TypedObjectLiteralNode interface {
	AST_Expression
	AST_ObjectLiteralNode
	AST_DefType() AST_Type
}

type AST_ObjectLiteralNode interface {
	AST_Expression
	AST_Values() []AST_KeyValuePair
}

type AST_KeyValuePair interface {
	key() AST_Token
	AST_Value() AST_Expression
}

type AST_TypedArrayLiteralNode interface {
	AST_Expression
	AST_ArrayLiteralNode
	AST_DefType() AST_Type
}

type AST_ArrayLiteralNode interface {
	AST_Expression
	AST_Values() []AST_Expression
}

type AST_LambdaLiteralNode interface {
	AST_DefType() AST_Type
	AST_Expression
	AST_Params() []AST_Expression
	AST_Body() AST_Expression
}

type AST_TypedDeclarationNode interface {
	AST_DefType() AST_Type
	AST_Target() AST_Declareable
}

type AST_ArrayTypeNode interface {
	AST_Type
	AST_DefType() AST_Type
	AST_Values() []AST_Declareable
}

type AST_ObjectTypeNode interface {
	AST_Type
	AST_DefType() AST_Type
	AST_Fields() []AST_TypedDeclarationNode
}

type AST_LambdaTypeNode interface {
	AST_Type
	AST_DefType() AST_Type
	AST_Params() []AST_TypedDeclarationNode
}

type AST_KeywordTypeNode interface {
	AST_Type
	AST_Value() AST_Token
}

type AST_IdentifierTypeNode interface {
	AST_IdentifierNode
	AST_Type
}

type AST_Type interface{}

type AST_Declareable interface {
}

type AST_UnpackedArrayNode interface {
	AST_Declareable
	AST_Targets() []AST_Declareable
}

type AST_UnpackedObjectNode interface {
	AST_Declareable
	AST_Targets() []AST_Declareable
}

type AST_RenamedIdentifierNode interface {
	AST_Declareable
	AST_Source() []AST_IdentifierNode
	AST_Target() []AST_IdentifierNode
}

type AST_AssignmentNode interface {
	AST_Expression
	AST_Target() AST_Expression
	AST_Source() AST_Expression
}

type AST_TernaryNode interface {
	AST_Expression
	AST_Condition() AST_Expression
	AST_Truthy() AST_Expression
	AST_Falsy() AST_Expression
}

type AST_BinaryOperationNode interface {
	AST_Expression
	AST_Left() AST_Expression
	AST_Right() AST_Expression
}

type AST_LogicalOrNode = AST_BinaryOperationNode
type AST_LogicalAndNode = AST_BinaryOperationNode
type AST_BitwiseOrNode = AST_BinaryOperationNode
type AST_BitwiseXorNode = AST_BinaryOperationNode
type AST_BitwiseAndNode = AST_BinaryOperationNode
type AST_EqualityNode = AST_BinaryOperationNode
type AST_InequalityNode = AST_BinaryOperationNode
type AST_CompareLTNode = AST_BinaryOperationNode
type AST_CompareLTENode = AST_BinaryOperationNode
type AST_CompareGTNode = AST_BinaryOperationNode
type AST_CompareGTENode = AST_BinaryOperationNode
type AST_BitShiftLeftNode = AST_BinaryOperationNode
type AST_BitShiftRightNode = AST_BinaryOperationNode
type AST_AdditionNode = AST_BinaryOperationNode
type AST_SubtractionNode = AST_BinaryOperationNode
type AST_MultiplicationNode = AST_BinaryOperationNode
type AST_DivisionNode = AST_BinaryOperationNode
type AST_RemainderNode = AST_BinaryOperationNode
type AST_ExponentationNode = AST_BinaryOperationNode

type AST_UnaryOperationNode interface {
	AST_Expression
	AST_Source() AST_Expression
}

type AST_LogicalNotNode = AST_UnaryOperationNode
type AST_BitwiseNotNode = AST_UnaryOperationNode
type AST_UnaryPlusNode = AST_UnaryOperationNode
type AST_UnaryNegationNode = AST_UnaryOperationNode
type AST_PreIncrementNode = AST_UnaryOperationNode
type AST_PreDecrementNode = AST_UnaryOperationNode
type AST_PostIncrementNode = AST_UnaryOperationNode
type AST_PostDecrementNode = AST_UnaryOperationNode

type AST_MemberAccessNode interface {
	AST_Expression
	AST_Source() AST_Expression
	AST_Specifier() AST_Token
}

type AST_ComputedMemberNode interface {
	AST_Expression
	AST_Source() AST_Expression
	AST_Specifier() AST_Expression
}

type AST_FuncCallNode interface {
	AST_Expression
	AST_Source() AST_Expression
	AST_Args() []AST_Expression
}

type AST_IntLiteralNode interface {
	AST_Expression
	AST_Value() AST_Token
}

type AST_FloatLiteralNode interface {
	AST_Expression
	AST_Value() AST_Token
}

type AST_CharLiteralNode interface {
	AST_Expression
	AST_Value() AST_Token
}

type AST_StringLiteralNode interface {
	AST_Expression
	AST_Value() AST_Token
}

type AST_VarAccessNode interface {
	AST_Expression
	AST_IdentifierNode
}

type AST_IdentifierNode interface {
	AST_Declareable
	AST_Value() AST_Token
}
