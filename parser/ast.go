package parser

import "fmt"

type NodeType int

const (
	InvalidNodeType NodeType = iota
	AddNodeType
	SubNodeType
	MulNodeType
	DivNodeType
	ExpNodeType
	IntNodeType
	LParenNodeType
	RParenNodeType
)

func (n NodeType) String() string {
	switch n {
	case AddNodeType:
		return "add"
	case SubNodeType:
		return "sub"
	case MulNodeType:
		return "mul"
	case DivNodeType:
		return "div"
	case ExpNodeType:
		return "exp"
	case LParenNodeType:
		return "l_paren"
	case RParenNodeType:
		return "r_paren"
	case IntNodeType:
		return "int"
	case InvalidNodeType:
		return "invalid"
	default:
		return "invalid"
	}
}

type BaseNode interface {
	Type() NodeType
	String() string
}
type ExpressionNode interface {
	BaseNode
}
type LeftRightNode struct {
	ExpressionNode
	Left  ExpressionNode
	Right ExpressionNode
}

type AddNode struct {
	LeftRightNode
}

func (n AddNode) Type() NodeType {
	return AddNodeType
}

func (n AddNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right)
}

type SubNode struct {
	LeftRightNode
}

func (n SubNode) Type() NodeType {
	return SubNodeType
}

func (n SubNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right)
}

type MulNode struct {
	LeftRightNode
}

func (n MulNode) Type() NodeType {
	return MulNodeType
}

func (n MulNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right)
}

type DivNode struct {
	LeftRightNode
}

func (n DivNode) Type() NodeType {
	return DivNodeType
}

func (n DivNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right)
}

type ExpNode struct {
	LeftRightNode
}

func (n ExpNode) Type() NodeType {
	return ExpNodeType
}

func (n ExpNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.Type(), n.Left, n.Right)
}

type IntLiteral struct {
	ExpressionNode,
	Tok *Token
}

func (n IntLiteral) Type() NodeType {
	return IntNodeType
}

func (n IntLiteral) String() string {
	return string(n.Tok.String())
}

type Parser struct {
	tok *Token
}

func (p *Parser) makeExpression() ExpressionNode {
	return p.makeAddition()
}

func (p *Parser) makeAddition() ExpressionNode {
	left := p.makeSubtraction()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == AddToken {
		p.next()
		right := p.makeAddition()
		return AddNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeSubtraction() ExpressionNode {
	left := p.makeMultiplication()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == SubToken {
		p.next()
		right := p.makeSubtraction()
		return SubNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeMultiplication() ExpressionNode {
	left := p.makeDivision()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == MulToken {
		p.next()
		right := p.makeMultiplication()
		return MulNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeDivision() ExpressionNode {
	left := p.makeExponentation()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == DivToken {
		p.next()
		right := p.makeDivision()
		return DivNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeExponentation() ExpressionNode {
	left := p.makeValue()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == ExpToken {
		p.next()
		right := p.makeExponentation()
		return ExpNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeValue() ExpressionNode {
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

func (p *Parser) Parse(t *Token) ExpressionNode {
	p.tok = t
	return p.makeExpression()
}
