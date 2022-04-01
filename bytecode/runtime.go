package bytecode

import (
	"fmt"
	"math"
)

type Runtime struct {
	Stack []uint64
	Pc    uint
	Sp    uint
}

func Run(p Program) (Runtime, error) {
	ctx := Runtime{
		Stack: make([]uint64, 8192),
		Pc:    0,
		Sp:    0,
	}
	for ctx.Pc < uint(len(p.Instructions)) {
		err := runInstruction(&ctx, p.Instructions[ctx.Pc])
		if err != nil {
			return ctx, err
		}
		ctx.Pc++
	}
	return ctx, nil
}

func runInstruction(ctx *Runtime, i Instruction) error {
	switch i.InstructionType() {
	case AddInstruction:
		ctx.Sp--
		b := ctx.Stack[ctx.Sp]
		ctx.Sp--
		ctx.Stack[ctx.Sp] += b
		ctx.Sp++
		return nil
	case SubtractInstruction:
		ctx.Sp--
		b := ctx.Stack[ctx.Sp]
		ctx.Sp--
		ctx.Stack[ctx.Sp] -= b
		ctx.Sp++
		return nil
	case MultiplyInstruction:
		ctx.Sp--
		b := ctx.Stack[ctx.Sp]
		ctx.Sp--
		ctx.Stack[ctx.Sp] *= b
		ctx.Sp++
		return nil
	case DivideInstruction:
		ctx.Sp--
		b := ctx.Stack[ctx.Sp]
		if b == 0 {
			return fmt.Errorf("division by zero")
		}
		ctx.Sp--
		ctx.Stack[ctx.Sp] /= b
		ctx.Sp++
		return nil
	case ExponentInstruction:
		ctx.Sp--
		b := ctx.Stack[ctx.Sp]
		ctx.Sp--
		a := ctx.Stack[ctx.Sp]
		ctx.Stack[ctx.Sp] = uint64(math.Pow(float64(a), float64(b)))
		ctx.Sp++
		return nil
	case PushInstruction:
		return runPush(ctx, i.(Push))
	default:
		return fmt.Errorf("instruction '%s' not implemented", i.InstructionType())
	}
}

func runPush(ctx *Runtime, i Push) error {
	switch i.Type {
	case I32:
		ctx.Stack[ctx.Sp] = uint64(i.Value)
		ctx.Sp++
		return nil
	default:
		return fmt.Errorf("not implemented")
	}
}
