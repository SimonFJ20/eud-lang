package bytecode

import (
	"fmt"
	"math"
	"os"
)

type RuntimeValue interface {
	Type() Type
	String() string
}

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

func (v U8Value) String() string    { return fmt.Sprintf("U8(%d)", v.Value) }
func (v U16Value) String() string   { return fmt.Sprintf("U16(%d)", v.Value) }
func (v U32Value) String() string   { return fmt.Sprintf("U32(%d)", v.Value) }
func (v U64Value) String() string   { return fmt.Sprintf("U64(%d)", v.Value) }
func (v I8Value) String() string    { return fmt.Sprintf("I8(%d)", v.Value) }
func (v I16Value) String() string   { return fmt.Sprintf("I16(%d)", v.Value) }
func (v I32Value) String() string   { return fmt.Sprintf("I32(%d)", v.Value) }
func (v I64Value) String() string   { return fmt.Sprintf("I64(%d)", v.Value) }
func (v F32Value) String() string   { return fmt.Sprintf("F32(%f)", v.Value) }
func (v F64Value) String() string   { return fmt.Sprintf("F64(%f)", v.Value) }
func (v CharValue) String() string  { return fmt.Sprintf("CHAR(%d)", v.Value) }
func (v UsizeValue) String() string { return fmt.Sprintf("USIZE(%d)", v.Value) }
func (v UptrValue) String() string  { return fmt.Sprintf("UPTR(%d)", v.Value) }

type AllocationEntry struct {
	From uintptr
	To   uintptr
}

type Runtime struct {
	Stack   []RuntimeValue
	Locals  map[uint]RuntimeValue
	Globals map[uintptr]RuntimeValue
	Pc      uintptr
	Sp      uint
	Heap    []RuntimeValue
	Allocs  []AllocationEntry
	Files   []os.File
	Debug   bool
}

func (r *Runtime) String() string {
	stack_string := "["
	first := true
	for i := 0; i < int(r.Sp); i++ {
		if !first {
			stack_string += ", "
		}
		first = false
		stack_string += r.Stack[i].String()
	}
	stack_string += "]"
	return fmt.Sprintf("PC: %d,\tSP: %d,\tSTACK: %s", r.Pc, r.Sp, stack_string)
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
		Stack:   make([]RuntimeValue, 8192),
		Locals:  make(map[uint]RuntimeValue),
		Globals: make(map[uintptr]RuntimeValue),
		Pc:      0,
		Sp:      0,
		Heap:    make([]RuntimeValue, 8192),
		Allocs:  []AllocationEntry{},
		Files:   []os.File{},
		Debug:   p.RunWithDebug || false,
	}
	for ctx.Pc < uintptr(len(p.Instructions)) {
		if ctx.Debug {
			fmt.Printf("  %s\t%s\n", p.Instructions[ctx.Pc].String(), ctx.String())
		}
		runInstruction(&ctx, p.Instructions[ctx.Pc])
		ctx.Pc++
	}
	return ctx
}

func runInstruction(ctx *Runtime, i Instruction) {
	switch i.InstructionType() {
	case AllocateInstruction:
		runAllocate(ctx, i.(Allocate))
	case DeallocateInstruction:
		runDeallocate(ctx, i.(Deallocate))
	case StoreInstruction:
		runStore(ctx, i.(Store))
	case LoadInstruction:
		runLoad(ctx, i.(Load))
	case DeclareLocalInstruction:
		runDeclareLocal(ctx, i.(DeclareLocal))
	case StoreLocalInstruction:
		runStoreLocal(ctx, i.(StoreLocal))
	case LoadLocalInstruction:
		runLoadLocal(ctx, i.(LoadLocal))
	case PushInstruction:
		runPush(ctx, i.(Push))
	case PopInstruction:
		runPop(ctx, i.(Pop))
	case JumpInstruction:
		runJump(ctx, i.(Jump))
	case JumpIfZeroInstruction:
		runJumpIfZero(ctx, i.(JumpIfZero))
	case JumpNotZeroInstruction:
		runJumpNotZero(ctx, i.(JumpNotZero))
	case CallInstruction:
		runCall(ctx, i.(Call))
	case ReturnInstruction:
		runReturn(ctx, i.(Return))
	case NotInstruction:
		runNot(ctx, i.(Not))
	case AddInstruction:
		runAdd(ctx, i)
	case SubtractInstruction:
		runSubtract(ctx, i)
	case MultiplyInstruction:
		runMultiply(ctx, i)
	case DivideInstruction:
		runDivide(ctx, i)
	case ModulusInstruction:
		runModulus(ctx, i)
	case ExponentInstruction:
		runExponent(ctx, i)
	case CmpEqualInstruction:
		runCmpEqual(ctx, i)
	case CmpInequalInstruction:
		runCmpInequal(ctx, i)
	case CmpLTInstruction:
		runCmpLT(ctx, i)
	case CmpGTInstruction:
		runCmpGT(ctx, i)
	case CmpLTEInstruction:
		runCmpLTE(ctx, i)
	case CmpGTEInstruction:
		runCmpGTE(ctx, i)
	case OrInstruction:
		runOr(ctx, i)
	case AndInstruction:
		runAnd(ctx, i)
	case XorInstruction:
		runXor(ctx, i)
	case NorInstruction:
		runNor(ctx, i)
	case NandInstruction:
		runNand(ctx, i)
	case XnorInstruction:
		runXnor(ctx, i)
	case SyscallInstruction:
		runSyscall(ctx, i.(Syscall))
	default:
		panic(fmt.Sprintf("instruction '%s' not implemented", i.InstructionType()))
	}
}

func runAllocate(ctx *Runtime, i Allocate) {
	amount := ctx.Pop().(UsizeValue).Value
	size := amount * byteSizeOfType(i.Type)
	var addr uintptr
	if len(ctx.Allocs) > 0 {
		addr = ctx.Allocs[len(ctx.Allocs)-1].To + 1
	} else {
		addr = 0
	}
	ctx.Allocs = append(ctx.Allocs, AllocationEntry{
		From: addr,
		To:   addr + uintptr(size),
	})
	ctx.Push(UptrValue{Value: addr})
}

func byteSizeOfType(t Type) uint64 {
	// runtime/platform/target dependent
	switch t {
	case U8:
		return 8
	case U16:
		return 16
	case U32:
		return 32
	case U64:
		return 64
	case I8:
		return 8
	case I16:
		return 16
	case I32:
		return 32
	case I64:
		return 64
	case F32:
		return 32
	case F64:
		return 64
	case CHAR:
		return 8
	case USIZE:
		return 64
	case UPTR:
		return 64
	default:
		panic("unexhaustive")
	}
}

func runDeallocate(ctx *Runtime, i Deallocate) {
	addr := ctx.Pop().(UptrValue).Value
	for i := range ctx.Allocs {
		if ctx.Allocs[i].From >= addr && ctx.Allocs[i].To <= addr {
			ctx.Allocs = append(ctx.Allocs[:i], ctx.Allocs[i+1:]...)
			break
		}
	}
}

func runStore(ctx *Runtime, i Store) {
	addr := ctx.Pop().(UptrValue).Value
	allocated := false
	for i := range ctx.Allocs {
		if addr >= ctx.Allocs[i].From && addr <= ctx.Allocs[i].To {
			allocated = true
			break
		}
	}
	if !allocated {
		print("Segmentation fault")
		os.Exit(1)
	}
	ctx.Heap[addr] = ctx.Pop()
}
func runLoad(ctx *Runtime, i Load) {
	addr := ctx.Pop().(UptrValue).Value
	allocated := false
	for i := range ctx.Allocs {
		if addr >= ctx.Allocs[i].From && addr <= ctx.Allocs[i].To {
			allocated = true
			break
		}
	}
	if !allocated {
		print("Segmentation fault")
		os.Exit(1)
	}
	ctx.Push(ctx.Heap[addr])
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

func runJump(ctx *Runtime, i Jump) {
	ctx.Pc = ctx.Pop().(UptrValue).Value - 1 // compensate for iterating ctx.Pc++
}

func getIntValue(v RuntimeValue) int {
	switch v.Type() {
	case U8:
		return int(v.(U8Value).Value)
	case U16:
		return int(v.(U16Value).Value)
	case U32:
		return int(v.(U32Value).Value)
	case U64:
		return int(v.(U64Value).Value)
	case I8:
		return int(v.(I8Value).Value)
	case I16:
		return int(v.(I16Value).Value)
	case I32:
		return int(v.(I32Value).Value)
	case I64:
		return int(v.(I64Value).Value)
	case CHAR:
		return int(v.(CharValue).Value)
	case USIZE:
		return int(v.(UsizeValue).Value)
	case UPTR:
		return int(v.(UptrValue).Value)
	}
	panic("unreachable")
}

func runJumpIfZero(ctx *Runtime, i JumpIfZero) {
	addr := ctx.Pop().(UptrValue).Value
	condition := getIntValue(ctx.Pop())
	if condition == 0 {
		ctx.Pc = addr - 1 // compensate for iterating ctx.Pc++
	}
}

func runJumpNotZero(ctx *Runtime, i JumpNotZero) {
	addr := ctx.Pop().(UptrValue).Value
	condition := getIntValue(ctx.Pop())
	if condition != 0 {
		ctx.Pc = addr - 1 // compensate for iterating ctx.Pc++
	}
}

func runCall(ctx *Runtime, i Call) {
	addr := ctx.Pop().(UptrValue).Value
	argc := ctx.Pop().(UsizeValue).Value
	argv := []RuntimeValue{}
	for i := 0; i < int(argc); i++ {
		argv = append(argv, ctx.Pop())
	}
	ctx.Push(UptrValue{Value: ctx.Pc + 1})
	for i := range argv {
		ctx.Push(argv[i])
	}
	ctx.Pc = addr - 1
}

func runReturn(ctx *Runtime, i Return) {
	value := ctx.Pop()
	addr := ctx.Pop().(UptrValue).Value
	ctx.Push(value)
	ctx.Pc = addr - 1
}

func runPush(ctx *Runtime, i Push) {
	switch i.Type {
	case U8:
		ctx.Push(U8Value{Value: uint8(i.Value)})
	case U16:
		ctx.Push(U16Value{Value: uint16(i.Value)})
	case U32:
		ctx.Push(U32Value{Value: uint32(i.Value)})
	case U64:
		ctx.Push(U64Value{Value: uint64(i.Value)})
	case I16:
		ctx.Push(I16Value{Value: int16(i.Value)})
	case I32:
		ctx.Push(I32Value{Value: int32(i.Value)})
	case I64:
		ctx.Push(I64Value{Value: int64(i.Value)})
	case F32:
		ctx.Push(F32Value{Value: float32(i.Value)})
	case F64:
		ctx.Push(F64Value{Value: float64(i.Value)})
	case CHAR:
		ctx.Push(CharValue{Value: int8(i.Value)})
	case USIZE:
		ctx.Push(UsizeValue{Value: uint64(i.Value)})
	case UPTR:
		ctx.Push(UptrValue{Value: uintptr(i.Value)})
	default:
		panic(fmt.Sprintf("Push<%s> not implemented", i.Type))
	}
}

func runPop(ctx *Runtime, i Pop) {
	ctx.Pop()
}

func runNot(ctx *Runtime, i Not) {
	a := ctx.Pop()
	switch i.Type {
	case U8:
		ctx.Push(U8Value{Value: ^a.(U8Value).Value})
	case U16:
		ctx.Push(U16Value{Value: ^a.(U16Value).Value})
	case U32:
		ctx.Push(U32Value{Value: ^a.(U32Value).Value})
	case U64:
		ctx.Push(U64Value{Value: ^a.(U64Value).Value})
	case I8:
		ctx.Push(I8Value{Value: ^a.(I8Value).Value})
	case I16:
		ctx.Push(I16Value{Value: ^a.(I16Value).Value})
	case I32:
		ctx.Push(I32Value{Value: ^a.(I32Value).Value})
	case I64:
		ctx.Push(I64Value{Value: ^a.(I64Value).Value})
	case CHAR:
		ctx.Push(CharValue{Value: ^a.(CharValue).Value})
	case USIZE:
		ctx.Push(UsizeValue{Value: ^a.(UsizeValue).Value})
	case UPTR:
		ctx.Push(UptrValue{Value: ^a.(UptrValue).Value})
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
	case U16:
		b := ctx.Pop().(U16Value).Value
		a := ctx.Pop().(U16Value).Value
		ctx.Push(U16Value{Value: u16Op(a, b)})
	case U32:
		b := ctx.Pop().(U32Value).Value
		a := ctx.Pop().(U32Value).Value
		ctx.Push(U32Value{Value: u32Op(a, b)})
	case U64:
		b := ctx.Pop().(U64Value).Value
		a := ctx.Pop().(U64Value).Value
		ctx.Push(U64Value{Value: u64Op(a, b)})
	case I8:
		b := ctx.Pop().(I8Value).Value
		a := ctx.Pop().(I8Value).Value
		ctx.Push(I8Value{Value: i8Op(a, b)})
		return
	case I16:
		b := ctx.Pop().(I16Value).Value
		a := ctx.Pop().(I16Value).Value
		ctx.Push(I16Value{Value: i16Op(a, b)})
	case I32:
		b := ctx.Pop().(I32Value).Value
		a := ctx.Pop().(I32Value).Value
		ctx.Push(I32Value{Value: i32Op(a, b)})
	case I64:
		b := ctx.Pop().(I64Value).Value
		a := ctx.Pop().(I64Value).Value
		ctx.Push(I64Value{Value: i64Op(a, b)})
	case CHAR:
		b := ctx.Pop().(CharValue).Value
		a := ctx.Pop().(CharValue).Value
		ctx.Push(CharValue{Value: charOp(a, b)})
	case USIZE:
		b := ctx.Pop().(UsizeValue).Value
		a := ctx.Pop().(UsizeValue).Value
		ctx.Push(UsizeValue{Value: usizeOp(a, b)})
	case UPTR:
		b := ctx.Pop().(UptrValue).Value
		a := ctx.Pop().(UptrValue).Value
		ctx.Push(UptrValue{Value: uptrOp(a, b)})
	}
}

func runSyscall(ctx *Runtime, i Syscall) {
	id := ctx.Pop().(UsizeValue).Value
	switch id {
	case 1000:
		ctx.Push(UptrValue{Value: ctx.Pc})
	default:
		panic(fmt.Sprintf("no syscall with id %d", id))
	}
}
