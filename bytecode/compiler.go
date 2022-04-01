package bytecode

import (
	"eud-lang/parser"
	"fmt"
)

type Compiler struct {
	instructions []Instruction
}

func Compile(ast parser.BaseNode) (Program, error) {
	ctx := Compiler{
		instructions: []Instruction{},
	}
	err := compileBaseNode(&ctx, ast)
	if err != nil {
		return Program{}, err
	}
	return Program{
		Instructions: ctx.instructions,
	}, nil
}

func compileBaseNode(ctx *Compiler, node parser.BaseNode) error {
	switch node.Type() {
	default:
		return compileExpression(ctx, node.(parser.ExpressionNode))
	}
}

func compileExpression(ctx *Compiler, node parser.ExpressionNode) error {
	switch node.Type() {
	case parser.AddNodeType:
		return compileAddNode(ctx, node.(parser.AddNode))
	case parser.SubNodeType:
		return compileSubNode(ctx, node.(parser.SubNode))
	case parser.MulNodeType:
		return compileMulNode(ctx, node.(parser.MulNode))
	case parser.DivNodeType:
		return compileDivNode(ctx, node.(parser.DivNode))
	case parser.ExpNodeType:
		return compileExpNode(ctx, node.(parser.ExpNode))
	case parser.IntNodeType:
		return compileIntLiteral(ctx, node.(parser.IntLiteral))
	default:
		return fmt.Errorf("unknown or unexpected node type '%s'", node.Type())
	}
}

func compileAddNode(ctx *Compiler, node parser.AddNode) error {
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

func compileSubNode(ctx *Compiler, node parser.SubNode) error {
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

func compileMulNode(ctx *Compiler, node parser.MulNode) error {
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

func compileDivNode(ctx *Compiler, node parser.DivNode) error {
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

func compileExpNode(ctx *Compiler, node parser.ExpNode) error {
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
