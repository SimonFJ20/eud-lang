package parser

import "fmt"

type StatementType int

const (
	DeclarationStatementType StatementType = iota
	FuncDefStatementType
	ExpressionStatementType
)

type BaseStatement interface {
	Type() StatementType
	String() string
}

type FuncDefStatement struct {
	BaseStatement
	Identifier Token
	DeclType   Type
	Parameters []TypedDeclaration
	Body       []BaseStatement
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

type BaseExpression interface {
	Type() ExpressionType
	String() string
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

func (p *Parser) makeExpression() BaseExpression {
	return p.makeAddition()
}

func (p *Parser) makeAddition() BaseExpression {
	left := p.makeSubtraction()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == AddToken {
		p.next()
		right := p.makeAddition()
		return AddExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeSubtraction() BaseExpression {
	left := p.makeMultiplication()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == SubToken {
		p.next()
		right := p.makeSubtraction()
		return SubExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeMultiplication() BaseExpression {
	left := p.makeDivision()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == MulToken {
		p.next()
		right := p.makeMultiplication()
		return MulExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeDivision() BaseExpression {
	left := p.makeExponentation()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == DivToken {
		p.next()
		right := p.makeDivision()
		return DivExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeExponentation() BaseExpression {
	left := p.makeValue()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == ExpToken {
		p.next()
		right := p.makeExponentation()
		return ExpExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeValue() BaseExpression {
	token := p.tok
	p.next()
	if token.Type == LParenToken {
		expr := p.makeExpression()
		if p.tok.Type != RParenToken {
			panic("unexpected: tokenType != RParen")
		}
		p.next()
		return expr
	} else if p.tok.Type == IntToken {
		return IntLiteral{
			Tok: p.tok,
		}
	} else if token.Type == IntToken {
		return IntLiteral{
			Tok: token,
		}
	} else {
		return p.makeExpression()
	}
}

func (p *Parser) next() {
	if p.tok.Next == nil {
		// finished parsing
	} else {
		p.tok = p.tok.Next
	}
}

func (p *Parser) Parse(t *Token) BaseStatement {
	p.tok = t
	return ExpressionStatement{Expression: p.makeExpression()}
}

func (n StatementType) String() string {
	switch n {
	case DeclarationStatementType:
		return "DeclarationStatement"
	case FuncDefStatementType:
		return "FuncDefStatement"
	case ExpressionStatementType:
		return "ExpressionStatement"
	default:
		panic("unexhaustive")
	}
}

func (n DeclarationStatement) Type() StatementType { return DeclarationStatementType }
func (n FuncDefStatement) Type() StatementType     { return FuncDefStatementType }
func (n ExpressionStatement) Type() StatementType  { return ExpressionStatementType }

func (n DeclarationStatement) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Identifier, n.DeclType)
}
func (n FuncDefStatement) String() string {
	return fmt.Sprintf("%s(%s, %s, [%s], [%s])", n.Type(), n.Identifier, n.DeclType, n.Parameters, n.Body)
}
func (n ExpressionStatement) String() string { return fmt.Sprintf("%s(%s)", n.Type(), n.Expression) }

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
	default:
		panic("unexhaustive")
	}
}

func (n VarAssignExpression) Type() ExpressionType { return VarAssignExpressionType }
func (n AddExpression) Type() ExpressionType       { return AddExpressionType }
func (n SubExpression) Type() ExpressionType       { return SubExpressionType }
func (n MulExpression) Type() ExpressionType       { return MulExpressionType }
func (n DivExpression) Type() ExpressionType       { return DivExpressionType }
func (n ExpExpression) Type() ExpressionType       { return ExpExpressionType }
func (n FuncCallExpression) Type() ExpressionType  { return FuncCallExpressionType }
func (n VarAccessExpression) Type() ExpressionType { return VarAccessExpressionType }
func (n IntLiteral) Type() ExpressionType          { return IntExpressionType }

func (n VarAssignExpression) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Identifier, n.Value)
}
func (n AddExpression) String() string { return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right) }
func (n SubExpression) String() string { return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right) }
func (n MulExpression) String() string { return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right) }
func (n DivExpression) String() string { return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right) }
func (n ExpExpression) String() string { return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right) }
func (n FuncCallExpression) String() string {
	return fmt.Sprintf("%s(%s, [%s])", n.Type(), n.Identifier, n.Arguments)
}
func (n VarAccessExpression) String() string { return fmt.Sprintf("%s(%s)", n.Type(), n.Identifier) }
func (n IntLiteral) String() string          { return string(n.Tok.String()) }
