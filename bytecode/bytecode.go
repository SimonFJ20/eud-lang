package bytecode

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
	return "Push"
}
func (node Pop) String() string {
	return "Pop"
}
func (node Add) String() string {
	return "Add"
}
func (node Subtract) String() string {
	return "Subtract"
}
func (node Multiply) String() string {
	return "Multiply"
}
func (node Divide) String() string {
	return "Divide"
}
func (node Modules) String() string {
	return "Modules"
}
func (node Exponent) String() string {
	return "Exponent"
}
func (node JumpIfZero) String() string {
	return "JumpIfZero"
}
func (node JumpNotZero) String() string {
	return "JumpNotZero"
}
func (node CmpEqual) String() string {
	return "CmpEqual"
}
func (node CmpInequal) String() string {
	return "CmpInequal"
}
func (node CmpLT) String() string {
	return "CmpLT"
}
func (node CmpGT) String() string {
	return "CmpGT"
}
func (node CmpLTE) String() string {
	return "CmpLTE"
}
func (node CmpGTE) String() string {
	return "CmpGTE"
}
func (node Not) String() string {
	return "Not"
}
func (node Or) String() string {
	return "Or"
}
func (node And) String() string {
	return "And"
}
func (node Xor) String() string {
	return "Xor"
}
func (node Nor) String() string {
	return "Nor"
}
func (node Nand) String() string {
	return "Nand"
}
func (node Xnor) String() string {
	return "Xnor"
}
