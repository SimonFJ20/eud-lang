package parser

import (
	"fmt"
	"strings"
)

type StatementType int

const (
	DeclarationStatementType StatementType = iota
	FuncDefStatementType
	ReturnStatementType
	ExpressionStatementType
)

type BaseStatement interface {
	StatementType() StatementType
	String() string
	StringNested(nesting int) string
}

type FuncDefStatement struct {
	BaseStatement
	Identifier Token
	ReturnType Type
	Parameters []TypedDeclaration
	Body       []BaseStatement
}

type BaseExpression interface {
	ExpressionType() ExpressionType
	String() string
	StringNested(nesting int) string
}

type VarAssignExpression struct {
	Identifier Token
	Value      BaseExpression
}

type LeftRightExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type AddExpression struct {
	LeftRightExpression
}

type SubExpression struct {
	LeftRightExpression
}

type MulExpression struct {
	LeftRightExpression
}

type DivExpression struct {
	LeftRightExpression
}

type ExpExpression struct {
	LeftRightExpression
}

type FuncCallExpression struct {
	Identifier BaseExpression
	Arguments  []BaseExpression
}

type VarAccessExpression struct {
	Identifier Token
}

type IntLiteral struct {
	BaseExpression,
	Tok *Token
}

type Parser struct {
	tok *Token
}

type ReturnStatement struct {
	BaseStatement
	Value BaseExpression
}

type DeclarationStatement struct {
	BaseStatement
	TypedDeclaration
}

type TypedDeclaration struct {
	DeclType   Type
	Identifier Token
}

type Type = Token

type ExpressionStatement struct {
	BaseStatement
	Expression BaseExpression
}

type ExpressionType int

const (
	InvalidExpressionType ExpressionType = iota
	VarAccessExpressionType
	VarAssignExpressionType
	AddExpressionType
	SubExpressionType
	MulExpressionType
	DivExpressionType
	ExpExpressionType
	IntExpressionType
	FuncCallExpressionType
)

func (n StatementType) String() string {
	switch n {
	case DeclarationStatementType:
		return "DeclarationStatement"
	case FuncDefStatementType:
		return "FuncDefStatement"
	case ExpressionStatementType:
		return "ExpressionStatement"
	default:
		panic("unexhaustive statementtype")
	}
}

func (n DeclarationStatement) StatementType() StatementType { return DeclarationStatementType }
func (n FuncDefStatement) StatementType() StatementType     { return FuncDefStatementType }
func (n ReturnStatement) StatementType() StatementType      { return ReturnStatementType }
func (n ExpressionStatement) StatementType() StatementType  { return ExpressionStatementType }

func (n DeclarationStatement) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.StatementType(), n.Identifier, n.DeclType)
}
func (n FuncDefStatement) String() string {
	return fmt.Sprintf("%s(%s, %s, %s, %s)", n.StatementType(), n.Identifier, n.ReturnType, n.Parameters, n.Body)
}
func (n ReturnStatement) String() string {
	return fmt.Sprintf("%s(%s)", n.StatementType(), n.Value)
}
func (n ExpressionStatement) String() string {
	return fmt.Sprintf("%s(%s)", n.StatementType(), n.Expression)
}

func nstr(nesting int) string {
	return strings.Repeat("|   ", nesting)
}

func (n DeclarationStatement) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(\n%s%s,\n%s%s\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Identifier,
		nstr(nesting+1),
		n.DeclType,
		nstr(nesting),
	)
}

func (n FuncDefStatement) StringNested(nesting int) string {
	body_str := ""
	first := true
	for i := range n.Body {
		if first {
			body_str += n.Body[i].StringNested(nesting + 2)
			first = false
		} else {
			body_str += ",\n" + n.Body[i].StringNested(nesting+2)
		}
	}
	return fmt.Sprintf(
		"%s%s(\n%s%s,\n%s%s,\n%s%s,\n%s[\n%s\n%s]\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Identifier,
		nstr(nesting+1),
		n.ReturnType,
		nstr(nesting+1),
		n.Parameters,
		nstr(nesting+1),
		body_str,
		nstr(nesting+1),
		nstr(nesting),
	)
}
func (n ReturnStatement) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(\n%s%s\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Value.StringNested(nesting+1),
		nstr(nesting),
	)
}
func (n ExpressionStatement) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(\n%s\n%s)",
		nstr(nesting),
		n.StatementType(),
		n.Expression.StringNested(nesting+1),
		nstr(nesting),
	)
}

func (n ExpressionType) String() string {
	switch n {
	case AddExpressionType:
		return "add"
	case SubExpressionType:
		return "sub"
	case MulExpressionType:
		return "mul"
	case DivExpressionType:
		return "div"
	case ExpExpressionType:
		return "exp"
	case IntExpressionType:
		return "int"
	case InvalidExpressionType:
		return "invalid"
	case VarAccessExpressionType:
		return "var_access"
	case VarAssignExpressionType:
		return "var_assign"
	case FuncCallExpressionType:
		return "func_call"
	default:
		panic("unexhaustive expressiontype")
	}
}

func (n VarAssignExpression) ExpressionType() ExpressionType { return VarAssignExpressionType }
func (n AddExpression) ExpressionType() ExpressionType       { return AddExpressionType }
func (n SubExpression) ExpressionType() ExpressionType       { return SubExpressionType }
func (n MulExpression) ExpressionType() ExpressionType       { return MulExpressionType }
func (n DivExpression) ExpressionType() ExpressionType       { return DivExpressionType }
func (n ExpExpression) ExpressionType() ExpressionType       { return ExpExpressionType }
func (n FuncCallExpression) ExpressionType() ExpressionType  { return FuncCallExpressionType }
func (n VarAccessExpression) ExpressionType() ExpressionType { return VarAccessExpressionType }
func (n IntLiteral) ExpressionType() ExpressionType          { return IntExpressionType }

func (n VarAssignExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Identifier, n.Value)
}

func (n LeftRightExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n FuncCallExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Identifier, n.Arguments)
}
func (n VarAccessExpression) String() string {
	return fmt.Sprintf("%s(%s)", n.ExpressionType(), n.Identifier)
}
func (n IntLiteral) String() string { return string(n.Tok.String()) }

func (n VarAssignExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Identifier,
		n.Value,
	)
}

func (n LeftRightExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left, n.Right,
	)
}
func (n FuncCallExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, [%s])",
		nstr(nesting),
		n.ExpressionType(),
		n.Identifier,
		n.Arguments,
	)
}
func (n VarAccessExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Identifier,
	)
}
func (n IntLiteral) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s",
		nstr(nesting),
		n.Tok.String(),
	)
}
