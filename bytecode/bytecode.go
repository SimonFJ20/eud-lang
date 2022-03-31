package bytecode

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

type Value interface{}

type U8Literal struct {
	Value
	Type
}

type Instruction interface{}

type Declare struct {
	Instruction
	Type
	Identifier uint
}

type Store struct {
	Instruction
	Type
	Identifier Type
}

type Load struct {
	Instruction
	Type
	Identifier uint
}

type Push struct {
	Instruction
	Type
	Value
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
	Destination uint
}

type JumpNotZero struct {
	Instruction
	Destination uint
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
