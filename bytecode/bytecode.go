package bytecode

type Program struct {
	Instructions []Instruction
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

type Allocation struct {
	Handle int
}

type Instruction interface{}

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
