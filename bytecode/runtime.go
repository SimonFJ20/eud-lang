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
	Stack  []RuntimeValue
	Locals map[uint]RuntimeValue
	Pc     uint
	Sp     uint
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
		Stack:  make([]RuntimeValue, 8192),
		Locals: make(map[uint]RuntimeValue),
		Pc:     0,
		Sp:     0,
	}
	for ctx.Pc < uint(len(p.Instructions)) {
		runInstruction(&ctx, p.Instructions[ctx.Pc])
		ctx.Pc++
	}
	return ctx
}

func runInstruction(ctx *Runtime, i Instruction) {
	switch i.InstructionType() {
	// case AllocateInstruction:
	// case StoreInstruction:
	// case LoadInstruction:
	case DeclareLocalInstruction:
		runDeclareLocal(ctx, i.(DeclareLocal))
		return
	case StoreLocalInstruction:
		runStoreLocal(ctx, i.(StoreLocal))
		return
	case LoadLocalInstruction:
		runLoadLocal(ctx, i.(LoadLocal))
		return
	case PushInstruction:
		runPush(ctx, i.(Push))
		return
	case PopInstruction:
		runPop(ctx, i.(Pop))
		return
	case JumpIfZeroInstruction:
		runJumpIfZero(ctx, i.(JumpIfZero))
		return
	case JumpNotZeroInstruction:
		runJumpNotZero(ctx, i.(JumpNotZero))
		return
	case NotInstruction:
		runNot(ctx, i.(Not))
		return
	case AddInstruction:
		runAdd(ctx, i)
		return
	case SubtractInstruction:
		runSubtract(ctx, i)
		return
	case MultiplyInstruction:
		runMultiply(ctx, i)
		return
	case DivideInstruction:
		runDivide(ctx, i)
		return
	case ModulusInstruction:
		runModulus(ctx, i)
		return
	case ExponentInstruction:
		runExponent(ctx, i)
		return
	case CmpEqualInstruction:
		runCmpEqual(ctx, i)
		return
	case CmpInequalInstruction:
		runCmpInequal(ctx, i)
		return
	case CmpLTInstruction:
		runCmpLT(ctx, i)
		return
	case CmpGTInstruction:
		runCmpGT(ctx, i)
		return
	case CmpLTEInstruction:
		runCmpLTE(ctx, i)
		return
	case CmpGTEInstruction:
		runCmpGTE(ctx, i)
		return
	case OrInstruction:
		runOr(ctx, i)
		return
	case AndInstruction:
		runAnd(ctx, i)
		return
	case XorInstruction:
		runXor(ctx, i)
		return
	case NorInstruction:
		runNor(ctx, i)
		return
	case NandInstruction:
		runNand(ctx, i)
		return
	case XnorInstruction:
		runXnor(ctx, i)
		return
	default:
		panic(fmt.Sprintf("instruction '%s' not implemented", i.InstructionType()))
	}
}

func runDeclareLocal(ctx *Runtime, i DeclareLocal) {
	switch i.Type {
	case U8:
		ctx.Locals[i.Handle] = U8Value{}
	case U16:
		ctx.Locals[i.Handle] = U16Value{}
	case U32:
		ctx.Locals[i.Handle] = U32Value{}
	case U64:
		ctx.Locals[i.Handle] = U64Value{}
	case I8:
		ctx.Locals[i.Handle] = I8Value{}
	case I16:
		ctx.Locals[i.Handle] = I16Value{}
	case I32:
		ctx.Locals[i.Handle] = I32Value{}
	case I64:
		ctx.Locals[i.Handle] = I64Value{}
	case CHAR:
		ctx.Locals[i.Handle] = CharValue{}
	case USIZE:
		ctx.Locals[i.Handle] = UsizeValue{}
	case UPTR:
		ctx.Locals[i.Handle] = UptrValue{}
	}
}

func runStoreLocal(ctx *Runtime, i StoreLocal) {
	ctx.Locals[i.Handle] = ctx.Pop()
}

func runLoadLocal(ctx *Runtime, i LoadLocal) {
	ctx.Push(ctx.Locals[i.Handle])
}

func runJumpIfZero(ctx *Runtime, i JumpIfZero) {
	addr := ctx.Pop().(UptrValue).Value
	condition := ctx.Pop().(U64Value).Value
	if condition == 0 {
		ctx.Pc = uint(addr) - 1 // compensate for iterating ctx.Pc++
	}
}

func runJumpNotZero(ctx *Runtime, i JumpNotZero) {
	addr := ctx.Pop().(UptrValue).Value
	condition := ctx.Pop().(U64Value).Value
	if condition != 0 {
		ctx.Pc = uint(addr) - 1 // compensate for iterating ctx.Pc++
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

func runPop(ctx *Runtime, i Pop) {
	ctx.Pop()
}

func runNot(ctx *Runtime, i Not) {
	a := ctx.Pop()
	switch i.Type {
	case U8:
		ctx.Push(U8Value{Value: ^a.(U8Value).Value})
		return
	case U16:
		ctx.Push(U16Value{Value: ^a.(U16Value).Value})
		return
	case U32:
		ctx.Push(U32Value{Value: ^a.(U32Value).Value})
		return
	case U64:
		ctx.Push(U64Value{Value: ^a.(U64Value).Value})
		return
	case I8:
		ctx.Push(I8Value{Value: ^a.(I8Value).Value})
		return
	case I16:
		ctx.Push(I16Value{Value: ^a.(I16Value).Value})
		return
	case I32:
		ctx.Push(I32Value{Value: ^a.(I32Value).Value})
		return
	case I64:
		ctx.Push(I64Value{Value: ^a.(I64Value).Value})
		return
	case CHAR:
		ctx.Push(CharValue{Value: ^a.(CharValue).Value})
		return
	case USIZE:
		ctx.Push(UsizeValue{Value: ^a.(UsizeValue).Value})
		return
	case UPTR:
		ctx.Push(UptrValue{Value: ^a.(UptrValue).Value})
		return
	}
}

func runAdd(ctx *Runtime, i Instruction) {
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
}

func runSubtract(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Subtract).Type,
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
}

func runMultiply(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Multiply).Type,
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
}

func runDivide(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Divide).Type,
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
}

func runModulus(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Modulus).Type,
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
}

func runExponent(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Exponent).Type,
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
}

func runCmpEqual(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(CmpEqual).Type,
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
}

func runCmpInequal(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(CmpInequal).Type,
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
}

func runCmpLT(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(CmpLT).Type,
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
}

func runCmpGT(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(CmpGT).Type,
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
}

func runCmpLTE(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(CmpLTE).Type,
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
}

func runCmpGTE(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(CmpGTE).Type,
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
}

func runOr(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Or).Type,
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
}

func runAnd(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(And).Type,
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
}

func runXor(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Xor).Type,
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
}

func runNor(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Nor).Type,
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
}

func runNand(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Nand).Type,
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
}

func runXnor(ctx *Runtime, i Instruction) {
	runBinaryOperationInstruction(
		ctx, i.(Xnor).Type,
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
