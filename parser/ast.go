package parser

import (
	"fmt"
	"strings"
)

type StatementType int

const (
	DeclarationStatementType StatementType = iota
	TypedInitStatementType
	FuncDefStatementType
	WhileStatementType
	IfElseStatementType
	IfStatementType
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

type WhileStatement struct {
	BaseStatement
	Condition BaseExpression
	Body      []BaseStatement
}

type IfElseStatement struct {
	BaseStatement
	Condition BaseExpression
	Truthy    []BaseStatement
	Falsy     []BaseStatement
}

type IfStatement struct {
	BaseStatement
	Condition BaseExpression
	Body      []BaseStatement
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

type NotEqualExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type EqualExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type GTEExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type LTEExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type GreaterThanExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type LessThanExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type AddExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type SubExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type MulExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type DivExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type ModExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type ExpExpression struct {
	BaseExpression
	Left  BaseExpression
	Right BaseExpression
}

type NonStdAllocExpression struct {
	Size BaseExpression
}

type NonStdDeallocExpression struct {
	Pointer BaseExpression
}

type NonStdSyscallExpression struct {
	Syscall   BaseExpression
	Arguments []BaseExpression
}

type NonStdAddrOfExpression struct {
	Target Token
}

type NonStdDerefExpression struct {
	Target BaseExpression
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

type ReturnStatement struct {
	BaseStatement
	Value BaseExpression
}

type TypedInitStatement struct {
	BaseStatement
	TypedDeclaration
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
	NotEqualExpressionType
	EqualExpressionType
	GTEExpressionType
	LTEExpressionType
	GreaterThanExpressionType
	LessThanExpressionType
	AddExpressionType
	SubExpressionType
	MulExpressionType
	DivExpressionType
	ModExpressionType
	ExpExpressionType
	IntExpressionType
	FuncCallExpressionType
	NonStdAllocExpressionType
	NonStdDeallocExpressionType
	NonStdSyscallExpressionType
	NonStdAddrOfExpressionType
	NonStdDerefExpressionType
)

func (n StatementType) String() string {
	switch n {
	case TypedInitStatementType:
		return "TypedInitStatement"
	case DeclarationStatementType:
		return "DeclarationStatement"
	case FuncDefStatementType:
		return "FuncDefStatement"
	case ReturnStatementType:
		return "ReturnStatement"
	case WhileStatementType:
		return "WhileStatement"
	case IfElseStatementType:
		return "IfElseStatement"
	case IfStatementType:
		return "IfStatement"
	case ExpressionStatementType:
		return "ExpressionStatement"
	default:
		panic("unexhaustive statementtype")
	}
}

func (n DeclarationStatement) StatementType() StatementType { return DeclarationStatementType }
func (n TypedInitStatement) StatementType() StatementType   { return TypedInitStatementType }
func (n FuncDefStatement) StatementType() StatementType     { return FuncDefStatementType }
func (n ReturnStatement) StatementType() StatementType      { return ReturnStatementType }
func (n WhileStatement) StatementType() StatementType       { return WhileStatementType }
func (n IfElseStatement) StatementType() StatementType      { return IfElseStatementType }
func (n IfStatement) StatementType() StatementType          { return IfStatementType }
func (n ExpressionStatement) StatementType() StatementType  { return ExpressionStatementType }

func (n DeclarationStatement) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.StatementType(), n.Identifier, n.DeclType)
}
func (n TypedInitStatement) String() string {
	return fmt.Sprintf("%s(%s, %s, %s)", n.StatementType(), n.Identifier, n.DeclType, n.Value)
}
func (n FuncDefStatement) String() string {
	return fmt.Sprintf("%s(%s, %s, %s, %s)", n.StatementType(), n.Identifier, n.ReturnType, n.Parameters, n.Body)
}
func (n WhileStatement) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.StatementType(), n.Condition, n.Body)
}
func (n IfElseStatement) String() string {
	return fmt.Sprintf("%s(%s, %s, %s)", n.StatementType(), n.Condition, n.Truthy, n.Falsy)
}
func (n IfStatement) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.StatementType(), n.Condition, n.Body)
}
func (n ReturnStatement) String() string {
	return fmt.Sprintf("%s(%s)", n.StatementType(), n.Value)
}
func (n ExpressionStatement) String() string {
	return fmt.Sprintf("%s(%s)", n.StatementType(), n.Expression)
}

func nstr(nesting int) string {
	return strings.Repeat("    ", nesting)
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

func (n TypedInitStatement) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(\n%s%s,\n%s%s,\n%s%s\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Identifier,
		nstr(nesting+1),
		n.DeclType,
		nstr(nesting+1),
		n.Value,
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
func (n WhileStatement) StringNested(nesting int) string {
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
		"%s%s(\n%s%s,\n%s[\n%s\n%s]\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Condition,
		nstr(nesting+1),
		body_str,
		nstr(nesting+1),
		nstr(nesting),
	)
}
func (n IfElseStatement) StringNested(nesting int) string {
	t_str := ""
	tfirst := true
	for i := range n.Truthy {
		if tfirst {
			t_str += n.Truthy[i].StringNested(nesting + 2)
			tfirst = false
		} else {
			t_str += ",\n" + n.Truthy[i].StringNested(nesting+2)
		}
	}
	f_str := ""
	ffirst := true
	for i := range n.Falsy {
		if ffirst {
			f_str += n.Falsy[i].StringNested(nesting + 2)
			ffirst = false
		} else {
			f_str += ",\n" + n.Falsy[i].StringNested(nesting+2)
		}
	}
	return fmt.Sprintf(
		"%s%s(\n%s%s,\n%s[\n%s\n%s],\n%s[\n%s\n%s]\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Condition,
		nstr(nesting+1),
		t_str,
		nstr(nesting+1),
		nstr(nesting+1),
		f_str,
		nstr(nesting+1),
		nstr(nesting),
	)
}
func (n IfStatement) StringNested(nesting int) string {
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
		"%s%s(\n%s%s,\n%s[\n%s\n%s]\n%s)",
		nstr(nesting),
		n.StatementType(),
		nstr(nesting+1),
		n.Condition,
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
		// fmt.Sprintf("%s<will, because bug, segfault if printed>", nstr(nesting+1)),
		nstr(nesting),
	)
}

func (n ExpressionType) String() string {
	switch n {
	case NotEqualExpressionType:
		return "inequal"
	case EqualExpressionType:
		return "equal"
	case GTEExpressionType:
		return "gte"
	case LTEExpressionType:
		return "lte"
	case GreaterThanExpressionType:
		return "gt"
	case LessThanExpressionType:
		return "lt"
	case AddExpressionType:
		return "add"
	case SubExpressionType:
		return "sub"
	case MulExpressionType:
		return "mul"
	case DivExpressionType:
		return "div"
	case ModExpressionType:
		return "mod"
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
	case NonStdAllocExpressionType:
		return "non_std_alloc"
	case NonStdDeallocExpressionType:
		return "non_std_dealloc"
	case NonStdSyscallExpressionType:
		return "non_std_syscall"
	case NonStdAddrOfExpressionType:
		return "non_std_addr_of"
	case NonStdDerefExpressionType:
		return "non_std_deref"
	case FuncCallExpressionType:
		return "func_call"
	default:
		panic("unexhaustive expressiontype")
	}
}

func (n VarAssignExpression) ExpressionType() ExpressionType     { return VarAssignExpressionType }
func (n NotEqualExpression) ExpressionType() ExpressionType      { return NotEqualExpressionType }
func (n EqualExpression) ExpressionType() ExpressionType         { return EqualExpressionType }
func (n GTEExpression) ExpressionType() ExpressionType           { return GTEExpressionType }
func (n LTEExpression) ExpressionType() ExpressionType           { return LTEExpressionType }
func (n GreaterThanExpression) ExpressionType() ExpressionType   { return GreaterThanExpressionType }
func (n LessThanExpression) ExpressionType() ExpressionType      { return LessThanExpressionType }
func (n AddExpression) ExpressionType() ExpressionType           { return AddExpressionType }
func (n SubExpression) ExpressionType() ExpressionType           { return SubExpressionType }
func (n MulExpression) ExpressionType() ExpressionType           { return MulExpressionType }
func (n DivExpression) ExpressionType() ExpressionType           { return DivExpressionType }
func (n ModExpression) ExpressionType() ExpressionType           { return ModExpressionType }
func (n ExpExpression) ExpressionType() ExpressionType           { return ExpExpressionType }
func (n FuncCallExpression) ExpressionType() ExpressionType      { return FuncCallExpressionType }
func (n NonStdAllocExpression) ExpressionType() ExpressionType   { return NonStdAllocExpressionType }
func (n NonStdDeallocExpression) ExpressionType() ExpressionType { return NonStdDeallocExpressionType }
func (n NonStdSyscallExpression) ExpressionType() ExpressionType { return NonStdSyscallExpressionType }
func (n NonStdAddrOfExpression) ExpressionType() ExpressionType  { return NonStdAddrOfExpressionType }
func (n NonStdDerefExpression) ExpressionType() ExpressionType   { return NonStdDerefExpressionType }
func (n VarAccessExpression) ExpressionType() ExpressionType     { return VarAccessExpressionType }
func (n IntLiteral) ExpressionType() ExpressionType              { return IntExpressionType }

func (n VarAssignExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Identifier, n.Value)
}
func (n NotEqualExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n EqualExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n GTEExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n LTEExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n GreaterThanExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n LessThanExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n AddExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n SubExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n MulExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n DivExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n ModExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n ExpExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Left, n.Right)
}
func (n NonStdAllocExpression) String() string {
	return fmt.Sprintf("%s(%s)", n.ExpressionType(), n.Size)
}
func (n NonStdDeallocExpression) String() string {
	return fmt.Sprintf("%s(%s)", n.ExpressionType(), n.Pointer)
}
func (n NonStdSyscallExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionType(), n.Syscall, n.Arguments)
}
func (n NonStdAddrOfExpression) String() string {
	return fmt.Sprintf("%s(%s)", n.ExpressionType(), n.Target)
}
func (n NonStdDerefExpression) String() string {
	return fmt.Sprintf("%s(%s)", n.ExpressionType(), n.Target)
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

func (n NotEqualExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n EqualExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n GTEExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n LTEExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n GreaterThanExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n LessThanExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n AddExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n SubExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n MulExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n DivExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n ModExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n ExpExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, %s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Left,
		n.Right,
	)
}
func (n NonStdAllocExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Size,
	)
}
func (n NonStdDeallocExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Pointer,
	)
}
func (n NonStdSyscallExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s, [%s])",
		nstr(nesting),
		n.ExpressionType(),
		n.Syscall,
		n.Arguments,
	)
}
func (n NonStdAddrOfExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Target,
	)
}
func (n NonStdDerefExpression) StringNested(nesting int) string {
	return fmt.Sprintf(
		"%s%s(%s)",
		nstr(nesting),
		n.ExpressionType(),
		n.Target,
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
