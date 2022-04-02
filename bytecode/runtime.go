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
type I8Value struct{ Value int8 }
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
func (v I8Value) Type() Type    { return I8 }
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

func runInstruction(ctx *Runtime, i Instruction) {
	switch i.InstructionType() {
	case AddInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a + b },
			func(a, b uint16) uint16 { return a + b },
			func(a, b uint32) uint32 { return a + b },
			func(a, b uint64) uint64 { return a + b },
			func(a, b int8) int8 { return a + b },
			func(a, b int16) int16 { return a + b },
			func(a, b int32) int32 { return a + b },
			func(a, b int64) int64 { return a + b },
			func(a, b int8) int8 { return a + b },
			func(a, b uint64) uint64 { return a + b },
			func(a, b uintptr) uintptr { return a + b },
		)
		return
	case SubtractInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a - b },
			func(a, b uint16) uint16 { return a - b },
			func(a, b uint32) uint32 { return a - b },
			func(a, b uint64) uint64 { return a - b },
			func(a, b int8) int8 { return a - b },
			func(a, b int16) int16 { return a - b },
			func(a, b int32) int32 { return a - b },
			func(a, b int64) int64 { return a - b },
			func(a, b int8) int8 { return a - b },
			func(a, b uint64) uint64 { return a - b },
			func(a, b uintptr) uintptr { return a - b },
		)
		return
	case MultiplyInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a * b },
			func(a, b uint16) uint16 { return a * b },
			func(a, b uint32) uint32 { return a * b },
			func(a, b uint64) uint64 { return a * b },
			func(a, b int8) int8 { return a * b },
			func(a, b int16) int16 { return a * b },
			func(a, b int32) int32 { return a * b },
			func(a, b int64) int64 { return a * b },
			func(a, b int8) int8 { return a * b },
			func(a, b uint64) uint64 { return a * b },
			func(a, b uintptr) uintptr { return a * b },
		)
		return
	case DivideInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a / b },
			func(a, b uint16) uint16 { return a / b },
			func(a, b uint32) uint32 { return a / b },
			func(a, b uint64) uint64 { return a / b },
			func(a, b int8) int8 { return a / b },
			func(a, b int16) int16 { return a / b },
			func(a, b int32) int32 { return a / b },
			func(a, b int64) int64 { return a / b },
			func(a, b int8) int8 { return a / b },
			func(a, b uint64) uint64 { return a / b },
			func(a, b uintptr) uintptr { return a / b },
		)
		return
	case ModulesInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a % b },
			func(a, b uint16) uint16 { return a % b },
			func(a, b uint32) uint32 { return a % b },
			func(a, b uint64) uint64 { return a % b },
			func(a, b int8) int8 { return a % b },
			func(a, b int16) int16 { return a % b },
			func(a, b int32) int32 { return a % b },
			func(a, b int64) int64 { return a % b },
			func(a, b int8) int8 { return a % b },
			func(a, b uint64) uint64 { return a % b },
			func(a, b uintptr) uintptr { return a % b },
		)
		return
	case ExponentInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return uint8(math.Pow(float64(a), float64(b))) },
			func(a, b uint16) uint16 { return uint16(math.Pow(float64(a), float64(b))) },
			func(a, b uint32) uint32 { return uint32(math.Pow(float64(a), float64(b))) },
			func(a, b uint64) uint64 { return uint64(math.Pow(float64(a), float64(b))) },
			func(a, b int8) int8 { return int8(math.Pow(float64(a), float64(b))) },
			func(a, b int16) int16 { return int16(math.Pow(float64(a), float64(b))) },
			func(a, b int32) int32 { return int32(math.Pow(float64(a), float64(b))) },
			func(a, b int64) int64 { return int64(math.Pow(float64(a), float64(b))) },
			func(a, b int8) int8 { return int8(math.Pow(float64(a), float64(b))) },
			func(a, b uint64) uint64 { return uint64(math.Pow(float64(a), float64(b))) },
			func(a, b uintptr) uintptr { return uintptr(math.Pow(float64(a), float64(b))) },
		)
		return
	case CmpEqualInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint16) uint16 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint32) uint32 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int16) int16 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int32) int32 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int64) int64 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uintptr) uintptr {
				if a == b {
					return 1
				} else {
					return 0
				}
			},
		)
		return
	case CmpInequalInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint16) uint16 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint32) uint32 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int16) int16 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int32) int32 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int64) int64 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uintptr) uintptr {
				if a != b {
					return 1
				} else {
					return 0
				}
			},
		)
		return
	case CmpLTInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint16) uint16 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint32) uint32 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int16) int16 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int32) int32 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int64) int64 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uintptr) uintptr {
				if a < b {
					return 1
				} else {
					return 0
				}
			},
		)
		return
	case CmpGTInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint16) uint16 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint32) uint32 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int16) int16 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int32) int32 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int64) int64 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uintptr) uintptr {
				if a > b {
					return 1
				} else {
					return 0
				}
			},
		)
		return
	case CmpLTEInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint16) uint16 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint32) uint32 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int16) int16 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int32) int32 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int64) int64 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uintptr) uintptr {
				if a <= b {
					return 1
				} else {
					return 0
				}
			},
		)
		return
	case CmpGTEInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint16) uint16 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint32) uint32 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int16) int16 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int32) int32 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int64) int64 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b int8) int8 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uint64) uint64 {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
			func(a, b uintptr) uintptr {
				if a >= b {
					return 1
				} else {
					return 0
				}
			},
		)
		return
	case OrInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a | b },
			func(a, b uint16) uint16 { return a | b },
			func(a, b uint32) uint32 { return a | b },
			func(a, b uint64) uint64 { return a | b },
			func(a, b int8) int8 { return a | b },
			func(a, b int16) int16 { return a | b },
			func(a, b int32) int32 { return a | b },
			func(a, b int64) int64 { return a | b },
			func(a, b int8) int8 { return a | b },
			func(a, b uint64) uint64 { return a | b },
			func(a, b uintptr) uintptr { return a | b },
		)
		return
	case AndInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a & b },
			func(a, b uint16) uint16 { return a & b },
			func(a, b uint32) uint32 { return a & b },
			func(a, b uint64) uint64 { return a & b },
			func(a, b int8) int8 { return a & b },
			func(a, b int16) int16 { return a & b },
			func(a, b int32) int32 { return a & b },
			func(a, b int64) int64 { return a & b },
			func(a, b int8) int8 { return a & b },
			func(a, b uint64) uint64 { return a & b },
			func(a, b uintptr) uintptr { return a & b },
		)
		return
	case XorInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return a ^ b },
			func(a, b uint16) uint16 { return a ^ b },
			func(a, b uint32) uint32 { return a ^ b },
			func(a, b uint64) uint64 { return a ^ b },
			func(a, b int8) int8 { return a ^ b },
			func(a, b int16) int16 { return a ^ b },
			func(a, b int32) int32 { return a ^ b },
			func(a, b int64) int64 { return a ^ b },
			func(a, b int8) int8 { return a ^ b },
			func(a, b uint64) uint64 { return a ^ b },
			func(a, b uintptr) uintptr { return a ^ b },
		)
		return
	case NorInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return ^(a | b) },
			func(a, b uint16) uint16 { return ^(a | b) },
			func(a, b uint32) uint32 { return ^(a | b) },
			func(a, b uint64) uint64 { return ^(a | b) },
			func(a, b int8) int8 { return ^(a | b) },
			func(a, b int16) int16 { return ^(a | b) },
			func(a, b int32) int32 { return ^(a | b) },
			func(a, b int64) int64 { return ^(a | b) },
			func(a, b int8) int8 { return ^(a | b) },
			func(a, b uint64) uint64 { return ^(a | b) },
			func(a, b uintptr) uintptr { return ^(a | b) },
		)
		return
	case NandInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return ^(a & b) },
			func(a, b uint16) uint16 { return ^(a & b) },
			func(a, b uint32) uint32 { return ^(a & b) },
			func(a, b uint64) uint64 { return ^(a & b) },
			func(a, b int8) int8 { return ^(a & b) },
			func(a, b int16) int16 { return ^(a & b) },
			func(a, b int32) int32 { return ^(a & b) },
			func(a, b int64) int64 { return ^(a & b) },
			func(a, b int8) int8 { return ^(a & b) },
			func(a, b uint64) uint64 { return ^(a & b) },
			func(a, b uintptr) uintptr { return ^(a & b) },
		)
		return
	case XnorInstruction:
		runBinaryOperationInstruction(
			ctx, i.(Add).Type,
			func(a, b uint8) uint8 { return ^(a ^ b) },
			func(a, b uint16) uint16 { return ^(a ^ b) },
			func(a, b uint32) uint32 { return ^(a ^ b) },
			func(a, b uint64) uint64 { return ^(a ^ b) },
			func(a, b int8) int8 { return ^(a ^ b) },
			func(a, b int16) int16 { return ^(a ^ b) },
			func(a, b int32) int32 { return ^(a ^ b) },
			func(a, b int64) int64 { return ^(a ^ b) },
			func(a, b int8) int8 { return ^(a ^ b) },
			func(a, b uint64) uint64 { return ^(a ^ b) },
			func(a, b uintptr) uintptr { return ^(a ^ b) },
		)
		return
	case PushInstruction:
		runPush(ctx, i.(Push))
		return
	default:
		panic(fmt.Sprintf("instruction '%s' not implemented", i.InstructionType()))
	}
}

func runBinaryOperationInstruction(
	ctx *Runtime,
	t Type,
	u8Op func(uint8, uint8) uint8,
	u16Op func(uint16, uint16) uint16,
	u32Op func(uint32, uint32) uint32,
	u64Op func(uint64, uint64) uint64,
	i8Op func(int8, int8) int8,
	i16Op func(int16, int16) int16,
	i32Op func(int32, int32) int32,
	i64Op func(int64, int64) int64,
	charOp func(int8, int8) int8,
	usizeOp func(uint64, uint64) uint64,
	uptrOp func(uintptr, uintptr) uintptr,
) {
	switch t {
	case U8:
		b := ctx.Pop().(U8Value).Value
		a := ctx.Pop().(U8Value).Value
		ctx.Push(U8Value{Value: u8Op(a, b)})
		return
	case U16:
		b := ctx.Pop().(U16Value).Value
		a := ctx.Pop().(U16Value).Value
		ctx.Push(U16Value{Value: u16Op(a, b)})
		return
	case U32:
		b := ctx.Pop().(U32Value).Value
		a := ctx.Pop().(U32Value).Value
		ctx.Push(U32Value{Value: u32Op(a, b)})
		return
	case U64:
		b := ctx.Pop().(U64Value).Value
		a := ctx.Pop().(U64Value).Value
		ctx.Push(U64Value{Value: u64Op(a, b)})
		return
	case I8:
		b := ctx.Pop().(I8Value).Value
		a := ctx.Pop().(I8Value).Value
		ctx.Push(I8Value{Value: i8Op(a, b)})
		return
	case I16:
		b := ctx.Pop().(I16Value).Value
		a := ctx.Pop().(I16Value).Value
		ctx.Push(I16Value{Value: i16Op(a, b)})
		return
	case I32:
		b := ctx.Pop().(I32Value).Value
		a := ctx.Pop().(I32Value).Value
		ctx.Push(I32Value{Value: i32Op(a, b)})
		return
	case I64:
		b := ctx.Pop().(I64Value).Value
		a := ctx.Pop().(I64Value).Value
		ctx.Push(I64Value{Value: i64Op(a, b)})
		return
	case CHAR:
		b := ctx.Pop().(CharValue).Value
		a := ctx.Pop().(CharValue).Value
		ctx.Push(CharValue{Value: charOp(a, b)})
		return
	case USIZE:
		b := ctx.Pop().(UsizeValue).Value
		a := ctx.Pop().(UsizeValue).Value
		ctx.Push(UsizeValue{Value: usizeOp(a, b)})
		return
	case UPTR:
		b := ctx.Pop().(UptrValue).Value
		a := ctx.Pop().(UptrValue).Value
		ctx.Push(UptrValue{Value: uptrOp(a, b)})
		return
	}
}

func runPush(ctx *Runtime, i Push) {
	switch i.Type {
	case U8:
		ctx.Push(U8Value{Value: uint8(i.Value)})
		return
	case U16:
		ctx.Push(U16Value{Value: uint16(i.Value)})
		return
	case U32:
		ctx.Push(U32Value{Value: uint32(i.Value)})
		return
	case U64:
		ctx.Push(U64Value{Value: uint64(i.Value)})
		return
	case I16:
		ctx.Push(I16Value{Value: int16(i.Value)})
		return
	case I32:
		ctx.Push(I32Value{Value: int32(i.Value)})
		return
	case I64:
		ctx.Push(I64Value{Value: int64(i.Value)})
		return
	case F32:
		ctx.Push(F32Value{Value: float32(i.Value)})
		return
	case F64:
		ctx.Push(F64Value{Value: float64(i.Value)})
		return
	case CHAR:
		ctx.Push(CharValue{Value: int8(i.Value)})
		return
	case USIZE:
		ctx.Push(UsizeValue{Value: uint64(i.Value)})
		return
	case UPTR:
		ctx.Push(UptrValue{Value: uintptr(i.Value)})
		return
	}
	panic(fmt.Sprintf("Push<%s> not implemented", i.Type))
}
