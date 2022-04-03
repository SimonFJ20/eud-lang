package bytecode_test

import (
	"eud-lang/bytecode"
	"testing"
)

func Test(t *testing.T) {
	TestLocals(t)
	TestMath1(t)
	TestMath2(t)
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
			bytecode.Push{Type: bytecode.I32, Value: 3},
			bytecode.Add{Type: bytecode.I32},
			bytecode.Multiply{Type: bytecode.I32},
		},
	})
	result := runtime.Pop().(bytecode.I32Value).Value
	if result != 23 {
		t.Errorf("3 * 4 + 5 != %d", result)
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
	if result != 35 {
		t.Errorf("(3 * 4) + 5 != %d", result)
	}
}
