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
	ExpressionNode,
	Left ExpressionNode
	Right ExpressionNode
}

func (n LeftRightNode) String() string {
	return fmt.Sprintf("%s(%s, %s)", n.ExpressionNode.Type(), n.Left, n.Right)
}

type AddNode struct {
	LeftRightNode
}

func (n AddNode) Type() NodeType {
	return AddNodeType
}

type SubNode struct {
	LeftRightNode
}

func (n SubNode) Type() NodeType {
	return SubNodeType
}

type MulNode struct {
	LeftRightNode
}

func (n MulNode) Type() NodeType {
	return MulNodeType
}

type DivNode struct {
	LeftRightNode
}

func (n DivNode) Type() NodeType {
	return DivNodeType
}

type ExpNode struct {
	LeftRightNode
}

func (n ExpNode) Type() NodeType {
	return ExpNodeType
}

type IntLiteral struct {
	ExpressionNode,
	Tok *Token
}

func (n IntLiteral) Type() NodeType {
	return IntNodeType
}

func (n IntLiteral) String() string {
	return string(n.Tok.Value)
}

type Parser struct {
	tok *Token
}

func (p *Parser) makeExpression() ExpressionNode {
	return p.makeAddition()
}

func (p *Parser) makeAddition() ExpressionNode {
	left := p.makeSubtraction()
	if p.tok.Type == AddToken {
		p.next()
		right := p.makeAddition()
		return AddNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return p.makeSubtraction()
	}
}

func (p *Parser) makeSubtraction() ExpressionNode {
	left := p.makeMultiplication()
	if p.tok.Type == SubToken {
		p.next()
		right := p.makeSubtraction()
		return SubNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return p.makeMultiplication()
	}
}

func (p *Parser) makeMultiplication() ExpressionNode {
	left := p.makeDivision()
	if p.tok.Type == MulToken {
		p.next()
		right := p.makeMultiplication()
		return MulNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return p.makeDivision()
	}
}

func (p *Parser) makeDivision() ExpressionNode {
	left := p.makeExponentation()
	if p.tok.Type == DivToken {
		p.next()
		right := p.makeDivision()
		return DivNode{LeftRightNode{Left: left, Right: right}}
	} else {
		return p.makeExponentation()
	}
}

func (p *Parser) makeExponentation() ExpressionNode {
	left := p.makeValue()
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
	if p.tok.Type == IntToken {
		return IntLiteral{
			Tok: p.tok,
		}
	} else if token.Type == LParenToken {
		expr := p.makeExpression()
		if p.tok.Type != RParenToken {
			panic("unexpected: tokenType != RParen")
		}
		return expr
	} else {
		panic("unexpected: token not valid for makeValue")
	}
}

func (p *Parser) next() {
	if p.tok == nil {
		fmt.Println("finished parsing")
	} else {
		p.tok = p.tok.Next
	}
}

func (p *Parser) Parse(t *Token) ExpressionNode {
	p.tok = t
	return p.makeExpression()
}
