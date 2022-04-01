package parser

import (
	"./tokenizer"
)

type BaseNode interface {
}
type ExpressionNode struct {
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

type MulNode struct {
	LeftRightNode
}

type DivNode struct {
	LeftRightNode
}

type ExpNode struct {
	LeftRightNode
}

type IntLiteral struct {
	ExpressionNode,
	Token tokenizer.Token
}

type Parser struct {
	tokens []tokenizer.Token
	pos    int
	tok    tokenizer.Token
	done   bool
}

func ParserNew(tokens []tokenizer.Token) Parser {
	Parser{
		tokens: tokens,
		pos:    0,
	}

}
