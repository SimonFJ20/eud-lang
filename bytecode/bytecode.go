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

type Instruction interface {
	String() string
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

type Modules struct {
	Instruction
	Type
}

type Exponent struct {
	Instruction
	Type
}

type JumpIfZero struct {
	Instruction
}

type JumpNotZero struct {
	Instruction
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

type Not struct {
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
		return "invalid"
	}
}

func (node Allocate) String() string {
	return "Allocate"
}
func (node Store) String() string {
	return "Store"
}
func (node Load) String() string {
	return "Load"
}
func (node DeclareLocal) String() string {
	return "DeclareLocal"
}
func (node StoreLocal) String() string {
	return "StoreLocal"
}
func (node LoadLocal) String() string {
	return "LoadLocal"
}
func (node Push) String() string {
	return fmt.Sprintf("Push<%s> %d", node.Type, node.Value)
}

// return fmt.Sprintf("Push<%s>", node.Type)

func (node Pop) String() string {
	return fmt.Sprintf("Pop<%s>", node.Type)
}
func (node Add) String() string {
	return fmt.Sprintf("Add<%s>", node.Type)
}
func (node Subtract) String() string {
	return fmt.Sprintf("Subtract<%s>", node.Type)
}
func (node Multiply) String() string {
	return fmt.Sprintf("Multiply<%s>", node.Type)
}
func (node Divide) String() string {
	return fmt.Sprintf("Divide<%s>", node.Type)
}
func (node Modules) String() string {
	return fmt.Sprintf("Modules<%s>", node.Type)
}
func (node Exponent) String() string {
	return fmt.Sprintf("Exponent<%s>", node.Type)
}
func (node JumpIfZero) String() string {
	return "JumpIfZero"
}
func (node JumpNotZero) String() string {
	return "JumpNotZero"
}
func (node CmpEqual) String() string {
	return fmt.Sprintf("CmpEqual<%s>", node.Type)
}
func (node CmpInequal) String() string {
	return fmt.Sprintf("CmpInequal<%s>", node.Type)
}
func (node CmpLT) String() string {
	return fmt.Sprintf("CmpLT<%s>", node.Type)
}
func (node CmpGT) String() string {
	return fmt.Sprintf("CmpGT<%s>", node.Type)
}
func (node CmpLTE) String() string {
	return fmt.Sprintf("CmpLTE<%s>", node.Type)
}
func (node CmpGTE) String() string {
	return fmt.Sprintf("CmpGTE<%s>", node.Type)
}
func (node Not) String() string {
	return fmt.Sprintf("Not<%s>", node.Type)
}
func (node Or) String() string {
	return fmt.Sprintf("Or<%s>", node.Type)
}
func (node And) String() string {
	return fmt.Sprintf("And<%s>", node.Type)
}
func (node Xor) String() string {
	return fmt.Sprintf("Xor<%s>", node.Type)
}
func (node Nor) String() string {
	return fmt.Sprintf("Nor<%s>", node.Type)
}
func (node Nand) String() string {
	return fmt.Sprintf("Nand<%s>", node.Type)
}
func (node Xnor) String() string {
	return fmt.Sprintf("Xnor<%s>", node.Type)
}
