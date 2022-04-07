package bytecode_test

import (
	"eud/bytecode"
	"testing"
)

func TestHeap(t *testing.T) {
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.DeclareLocal{Type: bytecode.UPTR, Handle: 0},
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.Push{Type: bytecode.USIZE, Value: 1},
			bytecode.Allocate{Type: bytecode.I32},
			bytecode.StoreLocal{Type: bytecode.UPTR, Handle: 0},
			bytecode.LoadLocal{Type: bytecode.UPTR, Handle: 0},
			bytecode.Store{Type: bytecode.I32},
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.LoadLocal{Type: bytecode.UPTR, Handle: 0},
			bytecode.Load{Type: bytecode.I32},
			bytecode.Add{Type: bytecode.I32},
			bytecode.LoadLocal{Type: bytecode.UPTR, Handle: 0},
			bytecode.Deallocate{Type: bytecode.I32},
		},
		RunWithDebug: true})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 8 {
		t.Errorf("unexpected result %d", result)
	}
}

func TestSubroutine(t *testing.T) {
	/*
		func sum(a: i32, b: i32): i32
			let sum: i32 = a + b
			return sum
		end
		let x: i32 = 5
		let y: i32 = 3
		let s: i32 = sum(x, y)

			Push uptr start
			Jump
		sum:
			Add i32
			Return i32
		start:
			Push i32 5
			Push i32 3
			Push usize 2
			Push uptr sum
			Call
	*/
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.Push{Type: bytecode.UPTR, Value: 4}, // 4 = start
			bytecode.Jump{},
			bytecode.Add{Type: bytecode.I32},
			bytecode.Return{Type: bytecode.I32},
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.Push{Type: bytecode.USIZE, Value: 2},
			bytecode.Push{Type: bytecode.UPTR, Value: 2}, // 2 = sum
			bytecode.Call{},
		},
		RunWithDebug: true})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 8 {
		t.Errorf("unexpected result %d", result)
	}
}

func TestJumpFunction(t *testing.T) {
	/*
		func sum(a: i32, b: i32): i32
			let sum: i32 = a + b
			return sum
		end
		let x: i32 = 5
		let y: i32 = 3
		let s: i32 = sum(x, y)

			Push uptr start
			Jump
		sum:
			DeclareLocal uptr 3
			DeclareLocal i32 4
			DeclareLocal i32 5
			DeclareLocal i32 6
			StoreLocal i32 5
			StoreLocal i32 4
			StoreLocal uptr 3
			LoadLocal i32 4
			LoadLocal i32 5
			Add i32
			StoreLocal i32 6
			LoadLocal i32 6
			LoadLocal uptr 3
			Jump
		start:
			DeclareLocal i32 0
			DeclareLocal i32 1
			DeclareLocal i32 2
			Push i32 5
			StoreLocal i32 0
			Push i32 3
			StoreLocal i32 1
			Push uptr program_counter
			Load uptr
			LoadLocal i32 1
			LoadLocal i32 0
			Push uptr sum
			Jump
			StoreLocal i32 2
			LoadLocal i32 2
	*/
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.Push{Type: bytecode.UPTR, Value: 16}, // 16 = start
			bytecode.Jump{},
			bytecode.DeclareLocal{Type: bytecode.UPTR, Handle: 3},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 4},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 5},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 6},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 5},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 4},
			bytecode.StoreLocal{Type: bytecode.UPTR, Handle: 3},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 4},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 5},
			bytecode.Add{Type: bytecode.I32},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 6},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 6},
			bytecode.LoadLocal{Type: bytecode.UPTR, Handle: 3},
			bytecode.Jump{},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 0},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 1},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 2},
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 0},
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Push{Type: bytecode.USIZE, Value: 1000}, // program_counter
			bytecode.Syscall{},
			bytecode.Push{Type: bytecode.UPTR, Value: 7},
			bytecode.Add{Type: bytecode.UPTR},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 1},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 0},
			bytecode.Push{Type: bytecode.UPTR, Value: 2}, // 2 = sum
			bytecode.Jump{},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 2},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 2},
		},
	})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 8 {
		t.Errorf("unexpected result %d", result)
	}
}

func TestJump(t *testing.T) {
	/*
		let a: int = 0
		if (a == 0)
			a = 4
		else
			a = 5
		end

			DeclareLocal i32 1
			Push i32 0
			StoreLocal i32 1
			LoadLocal i32 1
			Push uptr .else
			JumpIfZero
			Push i32 4
			StoreLocal i32 1
			Push uptr .end
			Jump
		.else:
			Push i32 5
			StoreLocal i32 1
		.end:
			LoadLocal 1
	*/
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Push{Type: bytecode.I32, Value: 0},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 1},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Push{Type: bytecode.UPTR, Value: 10},
			bytecode.JumpIfZero{},
			bytecode.Push{Type: bytecode.I32, Value: 4},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Push{Type: bytecode.UPTR, Value: 12},
			bytecode.Jump{},
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 1},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 1},
		},
	})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 5 {
		t.Errorf("unexpected result %d", result)
	}
}

func TestLocals(t *testing.T) {
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 0},
			bytecode.DeclareLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Push{Type: bytecode.I32, Value: 4},
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 0},
			bytecode.StoreLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 0},
			bytecode.Add{Type: bytecode.I32},
			bytecode.LoadLocal{Type: bytecode.I32, Handle: 1},
			bytecode.Add{Type: bytecode.I32},
		},
	})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 12 {
		t.Errorf("3 + 4 + 5 != %d", result)
	}
}

func TestMath1(t *testing.T) {
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.Push{Type: bytecode.I32, Value: 4},
			bytecode.Multiply{Type: bytecode.I32},
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.Add{Type: bytecode.I32},
		},
	})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 3+4*5 {
		t.Errorf("3 + 4 * 5 != %d", result)
	}
}

func TestMath2(t *testing.T) {
	runtime := bytecode.Run(bytecode.Program{
		Instructions: []bytecode.Instruction{
			bytecode.Push{Type: bytecode.I32, Value: 4},
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.Add{Type: bytecode.I32},
			bytecode.Push{Type: bytecode.I32, Value: 5},
			bytecode.Multiply{Type: bytecode.I32},
		},
	})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != (3+4)*5 {
		t.Errorf("(3 * 4) + 5 != %d", result)
	}
}
