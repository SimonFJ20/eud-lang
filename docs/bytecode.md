
# Bytecode

## Instructions

Operands are passed according to the stack, ie. top is 1st argument.
Operands are popped when used.

```go
type Instruction interface {
	String() string
	InstructionType() InstructionType
}
```

### Allocate

Allocate n amount of type on heap (malloc but abstracted).

Operand | Type | Description
---|---|---
1st | `usize` | Amount of `type` should be allocated

Pushes allocation address as uptr.

```go
type Allocate struct {
	Instruction
	Type
}
```

### Store

Move a value of type to somewhere in absolute memory.

Operand | Type | Description
---|---|---
1st | `uptr` | Target address
2nd | `type` | Value to be stored


```go
type Store struct {
	Instruction
	Type
}
```

### Load

Obtain a value of type from somewhere from absolute memory.

Operand | Type | Description
---|---|---
1st | `uptr` | Source address

```go
type Load struct {
	Instruction
	Type
}
```

### DeclareLocal

Reserve space on relative stack memory for value of type.

```go
type DeclareLocal struct {
	Instruction
	Type
	Handle uint
}
```

### StoreLocal

Store value of type in declared relative stack memory.

Operand | Type | Description
---|---|---
1st | `type` | Value to be stored

```go
type StoreLocal struct {
	Instruction
	Type
	Handle uint
}
```

### LoadLocal

Load a value of type from declared relative stack memory onto the stack.

```go
type LoadLocal struct {
	Instruction
	Type
	Handle uint
}
```

### Push

Push a literal value of type onto the stack.

```go
type Push struct {
	Instruction
	Type
	Value int
}
```

### Pop

Pop a value from the stack
```go
type Pop struct {
	Instruction
	Type
}
```

### Jump

Jump to address.

Operand | Type | Description
---|---|---
1st | `uptr` | Target address

```go
type Jump struct {
	Instruction
}
```

### JumpIfZero

Jump to address, if value of type is zero.

Operand | Type | Description
---|---|---
1st | `uptr` | Target address
2nd | `type` | Value

```go
type JumpIfZero struct {
	Instruction
}
```

### JumpNotZero

Jump to address, if value of type is not zero.

Operand | Type | Description
---|---|---
1st | `uptr` | Target address
2nd | `type` | Value

```go
type JumpNotZero struct {
	Instruction
}
```

### Call

Used for jumping to subroutines.

Operand | Type | Description
---|---|---
1st | `uptr` | Subroutine address
2nd | `usize` | Amount of arguments
... | unknown | Arguments to subroutine in reverse order

The instruction will pop all the operands, push the address
of the instruction next to the call instruction, push the arguments
in the right order, push the subroutine address and execute a
jump instruction.

This instruction depends on either one/a combination of the following:
A highlevel stack implementation,
a fancy type checker and/or fancy book keeping and
perfect input (matching argument types).

The type is for typecheckers to save time, although not implemented yet.

```go
type Call struct {
	Instruction
	Type
}
```

### Return

Used to return from subroutine. Mostly a convenience instruction.

Operand | Type | Description
---|---|---
1st | `type` | Value to be returned
2nd | `uptr` | Subroutine address

This will reverse the operands and execute a jump instruction.

```go
type Return struct {
	Instruction
    Type
}
```

### Not

Negates value bitwise and pushes it.

Operand | Type | Description
---|---|---
1st | `type` | Value

```go
type Not struct {
	Instruction
	Type
}
```

### Binary Operation

Run a binary operation, eg. add operation.

Operand | Type | Description
---|---|---
1st | `type` | Right value
2nd | `type` | Left value

```go
type BinaryOperation struct {
	Instruction
	Type
}

type Add = BinaryOperation
type Subtract = BinaryOperation
type Multiply = BinaryOperation
type Divide = BinaryOperation
type Modulus = BinaryOperation
type Exponent = BinaryOperation
type CmpEqual = BinaryOperation
type CmpInequal = BinaryOperation
type CmpLT = BinaryOperation
type CmpGT = BinaryOperation
type CmpLTE = BinaryOperation
type CmpGTE = BinaryOperation
type Or = BinaryOperation
type And = BinaryOperation
type Xor = BinaryOperation
type Nor = BinaryOperation
type Nand = BinaryOperation
type Xnor = BinaryOperation
```

