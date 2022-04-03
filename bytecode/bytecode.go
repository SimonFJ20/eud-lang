package bytecode

import "fmt"

type Program struct {
	Instructions   []Instruction
	Preallocations []AllocationStruct
}

type Type int

const (
	U8 Type = iota
	U16
	U32
	U64
	I8
	I16
	I32
	I64
	F32
	F64
	CHAR
	USIZE
	UPTR
)

type AllocationStruct struct {
	Handle     int
	Components []Allocation
	Pack       bool
}

type Allocation struct {
	Type
	Amount uint
}

type InstructionType int

const (
	AllocateInstruction InstructionType = iota
	StoreInstruction
	LoadInstruction
	DeclareLocalInstruction
	StoreLocalInstruction
	LoadLocalInstruction
	PushInstruction
	PopInstruction
	JumpInstruction
	JumpIfZeroInstruction
	JumpNotZeroInstruction
	NotInstruction
	AddInstruction
	SubtractInstruction
	MultiplyInstruction
	DivideInstruction
	ModulusInstruction
	ExponentInstruction
	CmpEqualInstruction
	CmpInequalInstruction
	CmpLTInstruction
	CmpGTInstruction
	CmpLTEInstruction
	CmpGTEInstruction
	OrInstruction
	AndInstruction
	XorInstruction
	NorInstruction
	NandInstruction
	XnorInstruction
)

type Instruction interface {
	String() string
	InstructionType() InstructionType
}

type Allocate struct {
	Instruction
	Type
}

type Store struct {
	Instruction
	Type
}

type Load struct {
	Instruction
	Type
}

type DeclareLocal struct {
	Instruction
	Type
	Handle uint
}

type StoreLocal struct {
	Instruction
	Type
	Handle uint
}

type LoadLocal struct {
	Instruction
	Type
	Handle uint
}

type Push struct {
	Instruction
	Type
	Value int
}

type Pop struct {
	Instruction
	Type
}

type Jump struct {
	Instruction
}

type JumpIfZero struct {
	Instruction
}

type JumpNotZero struct {
	Instruction
}

type Not struct {
	Instruction
	Type
}

type Add struct {
	Instruction
	Type
}

type Subtract struct {
	Instruction
	Type
}

type Multiply struct {
	Instruction
	Type
}

type Divide struct {
	Instruction
	Type
}

type Modulus struct {
	Instruction
	Type
}

type Exponent struct {
	Instruction
	Type
}

type CmpEqual struct {
	Instruction
	Type
}

type CmpInequal struct {
	Instruction
	Type
}

type CmpLT struct {
	Instruction
	Type
}

type CmpGT struct {
	Instruction
	Type
}

type CmpLTE struct {
	Instruction
	Type
}

type CmpGTE struct {
	Instruction
	Type
}

type Or struct {
	Instruction
	Type
}

type And struct {
	Instruction
	Type
}

type Xor struct {
	Instruction
	Type
}

type Nor struct {
	Instruction
	Type
}

type Nand struct {
	Instruction
	Type
}

type Xnor struct {
	Instruction
	Type
}

func (t Type) String() string {
	switch t {
	case U8:
		return "u8"
	case U16:
		return "u16"
	case U32:
		return "u32"
	case U64:
		return "u64"
	case I8:
		return "i8"
	case I16:
		return "i16"
	case I32:
		return "i32"
	case I64:
		return "i64"
	case F32:
		return "f32"
	case F64:
		return "f64"
	case CHAR:
		return "char"
	case USIZE:
		return "usize"
	case UPTR:
		return "uptr"
	default:
		panic("unknown")
	}
}

func (it InstructionType) String() string {
	switch it {
	case AllocateInstruction:
		return "AllocateInstruction"
	case StoreInstruction:
		return "StoreInstruction"
	case LoadInstruction:
		return "LoadInstruction"
	case DeclareLocalInstruction:
		return "DeclareLocalInstruction"
	case StoreLocalInstruction:
		return "StoreLocalInstruction"
	case LoadLocalInstruction:
		return "LoadLocalInstruction"
	case PushInstruction:
		return "PushInstruction"
	case PopInstruction:
		return "PopInstruction"
	case JumpInstruction:
		return "JumpInstruction"
	case JumpIfZeroInstruction:
		return "JumpIfZeroInstruction"
	case JumpNotZeroInstruction:
		return "JumpNotZeroInstruction"
	case NotInstruction:
		return "NotInstruction"
	case AddInstruction:
		return "AddInstruction"
	case SubtractInstruction:
		return "SubtractInstruction"
	case MultiplyInstruction:
		return "MultiplyInstruction"
	case DivideInstruction:
		return "DivideInstruction"
	case ModulusInstruction:
		return "ModulusInstruction"
	case ExponentInstruction:
		return "ExponentInstruction"
	case CmpEqualInstruction:
		return "CmpEqualInstruction"
	case CmpInequalInstruction:
		return "CmpInequalInstruction"
	case CmpLTInstruction:
		return "CmpLTInstruction"
	case CmpGTInstruction:
		return "CmpGTInstruction"
	case CmpLTEInstruction:
		return "CmpLTEInstruction"
	case CmpGTEInstruction:
		return "CmpGTEInstruction"
	case OrInstruction:
		return "OrInstruction"
	case AndInstruction:
		return "AndInstruction"
	case XorInstruction:
		return "XorInstruction"
	case NorInstruction:
		return "NorInstruction"
	case NandInstruction:
		return "NandInstruction"
	case XnorInstruction:
		return "XnorInstruction"
	default:
		panic("unknown")
	}
}

func (n Allocate) InstructionType() InstructionType     { return AllocateInstruction }
func (n Store) InstructionType() InstructionType        { return StoreInstruction }
func (n Load) InstructionType() InstructionType         { return LoadInstruction }
func (n DeclareLocal) InstructionType() InstructionType { return DeclareLocalInstruction }
func (n StoreLocal) InstructionType() InstructionType   { return StoreLocalInstruction }
func (n LoadLocal) InstructionType() InstructionType    { return LoadLocalInstruction }
func (n Push) InstructionType() InstructionType         { return PushInstruction }
func (n Pop) InstructionType() InstructionType          { return PopInstruction }
func (n Jump) InstructionType() InstructionType         { return JumpInstruction }
func (n JumpIfZero) InstructionType() InstructionType   { return JumpIfZeroInstruction }
func (n JumpNotZero) InstructionType() InstructionType  { return JumpNotZeroInstruction }
func (n Not) InstructionType() InstructionType          { return NotInstruction }
func (n Add) InstructionType() InstructionType          { return AddInstruction }
func (n Subtract) InstructionType() InstructionType     { return SubtractInstruction }
func (n Multiply) InstructionType() InstructionType     { return MultiplyInstruction }
func (n Divide) InstructionType() InstructionType       { return DivideInstruction }
func (n Modulus) InstructionType() InstructionType      { return ModulusInstruction }
func (n Exponent) InstructionType() InstructionType     { return ExponentInstruction }
func (n CmpEqual) InstructionType() InstructionType     { return CmpEqualInstruction }
func (n CmpInequal) InstructionType() InstructionType   { return CmpInequalInstruction }
func (n CmpLT) InstructionType() InstructionType        { return CmpLTInstruction }
func (n CmpGT) InstructionType() InstructionType        { return CmpGTInstruction }
func (n CmpLTE) InstructionType() InstructionType       { return CmpLTEInstruction }
func (n CmpGTE) InstructionType() InstructionType       { return CmpGTEInstruction }
func (n Or) InstructionType() InstructionType           { return OrInstruction }
func (n And) InstructionType() InstructionType          { return AndInstruction }
func (n Xor) InstructionType() InstructionType          { return XorInstruction }
func (n Nor) InstructionType() InstructionType          { return NorInstruction }
func (n Nand) InstructionType() InstructionType         { return NandInstruction }
func (n Xnor) InstructionType() InstructionType         { return XnorInstruction }

func (node Allocate) String() string     { return "Allocate" }
func (node Store) String() string        { return "Store" }
func (node Load) String() string         { return "Load" }
func (node DeclareLocal) String() string { return "DeclareLocal" }
func (node StoreLocal) String() string   { return "StoreLocal" }
func (node LoadLocal) String() string    { return "LoadLocal" }
func (node Push) String() string         { return fmt.Sprintf("Push<%s> %d", node.Type, node.Value) }
func (node Pop) String() string          { return fmt.Sprintf("Pop<%s>", node.Type) }
func (node Jump) String() string         { return "Jump" }
func (node JumpIfZero) String() string   { return "JumpIfZero" }
func (node JumpNotZero) String() string  { return "JumpNotZero" }
func (node Not) String() string          { return fmt.Sprintf("Not<%s>", node.Type) }
func (node Add) String() string          { return fmt.Sprintf("Add<%s>", node.Type) }
func (node Subtract) String() string     { return fmt.Sprintf("Subtract<%s>", node.Type) }
func (node Multiply) String() string     { return fmt.Sprintf("Multiply<%s>", node.Type) }
func (node Divide) String() string       { return fmt.Sprintf("Divide<%s>", node.Type) }
func (node Modulus) String() string      { return fmt.Sprintf("Modulus<%s>", node.Type) }
func (node Exponent) String() string     { return fmt.Sprintf("Exponent<%s>", node.Type) }
func (node CmpEqual) String() string     { return fmt.Sprintf("CmpEqual<%s>", node.Type) }
func (node CmpInequal) String() string   { return fmt.Sprintf("CmpInequal<%s>", node.Type) }
func (node CmpLT) String() string        { return fmt.Sprintf("CmpLT<%s>", node.Type) }
func (node CmpGT) String() string        { return fmt.Sprintf("CmpGT<%s>", node.Type) }
func (node CmpLTE) String() string       { return fmt.Sprintf("CmpLTE<%s>", node.Type) }
func (node CmpGTE) String() string       { return fmt.Sprintf("CmpGTE<%s>", node.Type) }
func (node Or) String() string           { return fmt.Sprintf("Or<%s>", node.Type) }
func (node And) String() string          { return fmt.Sprintf("And<%s>", node.Type) }
func (node Xor) String() string          { return fmt.Sprintf("Xor<%s>", node.Type) }
func (node Nor) String() string          { return fmt.Sprintf("Nor<%s>", node.Type) }
func (node Nand) String() string         { return fmt.Sprintf("Nand<%s>", node.Type) }
func (node Xnor) String() string         { return fmt.Sprintf("Xnor<%s>", node.Type) }
