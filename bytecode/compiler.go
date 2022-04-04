package bytecode

import (
	"eud/parser"
	"fmt"
)

type Compiler struct {
	instructions []Instruction
}

func Compile(ast parser.BaseExpression) (Program, error) {
	ctx := Compiler{
		instructions: []Instruction{},
	}
	err := compileBaseExpression(&ctx, ast)
	if err != nil {
		return Program{}, err
	}
	return Program{
		Instructions: ctx.instructions,
	}, nil
}

func compileBaseExpression(ctx *Compiler, node parser.BaseExpression) error {
	switch node.Type() {
	default:
		return compileExpression(ctx, node.(parser.BaseExpression))
	}
}

func compileExpression(ctx *Compiler, node parser.BaseExpression) error {
	switch node.Type() {
	case parser.AddExpressionType:
		return compileAddExpression(ctx, node.(parser.AddExpression))
	case parser.SubExpressionType:
		return compileSubExpression(ctx, node.(parser.SubExpression))
	case parser.MulExpressionType:
		return compileMulExpression(ctx, node.(parser.MulExpression))
	case parser.DivExpressionType:
		return compileDivExpression(ctx, node.(parser.DivExpression))
	case parser.ExpExpressionType:
		return compileExpExpression(ctx, node.(parser.ExpExpression))
	case parser.IntExpressionType:
		return compileIntLiteral(ctx, node.(parser.IntLiteral))
	default:
		return fmt.Errorf("unknown or unexpected node type '%s'", node.Type())
	}
}

func compileAddExpression(ctx *Compiler, node parser.AddExpression) error {
	var err error = nil
	err = compileExpression(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileExpression(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Add{Type: I32})
	return nil
}

func compileSubExpression(ctx *Compiler, node parser.SubExpression) error {
	var err error = nil
	err = compileExpression(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileExpression(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Subtract{Type: I32})
	return nil
}

func compileMulExpression(ctx *Compiler, node parser.MulExpression) error {
	var err error = nil
	err = compileExpression(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileExpression(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Multiply{Type: I32})
	return nil
}

func compileDivExpression(ctx *Compiler, node parser.DivExpression) error {
	var err error = nil
	err = compileExpression(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileExpression(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Divide{Type: I32})
	return nil
}

func compileExpExpression(ctx *Compiler, node parser.ExpExpression) error {
	var err error = nil
	err = compileExpression(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileExpression(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Exponent{Type: I32})
	return nil
}

func compileIntLiteral(ctx *Compiler, node parser.IntLiteral) error {
	ctx.instructions = append(ctx.instructions, Push{Type: I32, Value: node.Tok.Value})
	return nil
}
