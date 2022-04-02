package bytecode

import (
	"fmt"
	"math"
)

type RuntimeValue interface{ Type() Type }
type U8Value struct{ Value uint8 }
type U16Value struct{ Value uint16 }
type U32Value struct{ Value uint32 }
type U64Value struct{ Value uint64 }
type I16Value struct{ Value int16 }
type I32Value struct{ Value int32 }
type I64Value struct{ Value int64 }
type F32Value struct{ Value float32 }
type F64Value struct{ Value float64 }
type CharValue struct{ Value int8 }
type UsizeValue struct{ Value uint64 }
type UptrValue struct{ Value uintptr }

func (v U8Value) Type() Type    { return U8 }
func (v U16Value) Type() Type   { return U16 }
func (v U32Value) Type() Type   { return U32 }
func (v U64Value) Type() Type   { return U64 }
func (v I16Value) Type() Type   { return I16 }
func (v I32Value) Type() Type   { return I32 }
func (v I64Value) Type() Type   { return I64 }
func (v F32Value) Type() Type   { return F32 }
func (v F64Value) Type() Type   { return F64 }
func (v CharValue) Type() Type  { return CHAR }
func (v UsizeValue) Type() Type { return USIZE }
func (v UptrValue) Type() Type  { return UPTR }

type Runtime struct {
	Stack []RuntimeValue
	Pc    uint
	Sp    uint
}

func (ctx *Runtime) Push(v RuntimeValue) {
	if ctx.Sp >= uint(len(ctx.Stack)) {
		panic("stack overflow")
	}
	ctx.Stack[ctx.Sp] = v
	ctx.Sp++
}

func (ctx *Runtime) Pop() RuntimeValue {
	if ctx.Sp <= 0 {
		panic("stack underflow")
	}
	ctx.Sp--
	return ctx.Stack[ctx.Sp]
}

func Run(p Program) Runtime {
	ctx := Runtime{
		Stack: make([]RuntimeValue, 8192),
		Pc:    0,
		Sp:    0,
	}
	for ctx.Pc < uint(len(p.Instructions)) {
		runInstruction(&ctx, p.Instructions[ctx.Pc])
		ctx.Pc++
	}
	return ctx
}

func binaryOperation[T any, R RuntimeValue](ctx *Runtime, i Instruction, t Type, op func(T, T, Type) T) {
	b := ctx.Pop()
	a := ctx.Pop()
	if a.Type() != t || b.Type() != t {
		panic("type mismatch")
	}
	r := a.(I32Value).Value + b.(I32Value).Value
	ctx.Push(I32Value{Value: r})
}

func runInstruction(ctx *Runtime, i Instruction) {
	switch i.InstructionType() {
	case AddInstruction:
		binaryOperation[int32, I32Value](ctx, i, I32, func(a int32, b int32, t Type) int32 { return a + b })
	case SubtractInstruction:
		binaryOperation[int32, I32Value](ctx, i, I32, func(a int32, b int32, t Type) int32 { return a - b })
	case MultiplyInstruction:
		binaryOperation[int32, I32Value](ctx, i, I32, func(a int32, b int32, t Type) int32 { return a * b })
	case DivideInstruction:
		binaryOperation[int32, I32Value](ctx, i, I32, func(a int32, b int32, t Type) int32 { return a / b })
	case ExponentInstruction:
		binaryOperation[int32, I32Value](ctx, i, I32, func(a int32, b int32, t Type) int32 {
			return int32(math.Pow(float64(a), float64(b)))
		})
	case PushInstruction:
		runPush(ctx, i.(Push))
	default:
		panic(fmt.Sprintf("instruction '%s' not implemented", i.InstructionType()))
	}
}

func runPush(ctx *Runtime, i Push) {
	switch i.Type {
	// case U8:
	// case U16:
	// case U32:
	// case U64:
	// case I16:
	case I32:
		ctx.Push(I32Value{Value: int32(i.Value)})
		return
		// case I64:
		// case F32:
		// case F64:
		// case CHAR:
		// case USIZE:
		// case UPTR:
	}
	panic(fmt.Sprintf("Push<%s> not implemented", i.Type))
}
