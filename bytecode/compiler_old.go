package bytecode

import (
	"eud/parser"
	"fmt"
)

type Symbol_old struct {
	Type   Type
	Handle uint
}

type SymbolTable_old struct {
	parent  *SymbolTable_old
	symbols map[string]Symbol_old
}

func (s *SymbolTable_old) Set(name string, symbol Symbol_old) {
	s.symbols[name] = symbol
}

func (s *SymbolTable_old) Get(name string) (Symbol_old, error) {
	for i := range s.symbols {
		if i == name {
			return s.symbols[i], nil
		}
	}
	if s.parent != nil {
		return s.parent.Get(name)
	}
	return Symbol_old{}, fmt.Errorf("symbol \"%s\" undeclared", name)
}

func (s *SymbolTable_old) DefinedLocally(name string) bool {
	for i := range s.symbols {
		if i == name {
			return true
		}
	}
	return false
}

type Compiler_old struct {
	instructions []Instruction
	varId        uint
	symtable     SymbolTable_old
	globals      map[string]uintptr
	lastType     Type
}

func (ctx *Compiler_old) nextVarId() uint {
	ctx.varId++
	return ctx.varId - 1
}

func Compile_old(ast []parser.BaseStatement) (Program, error) {
	ctx := Compiler_old{
		instructions: []Instruction{},
		varId:        0,
		symtable: SymbolTable_old{
			parent:  nil,
			symbols: map[string]Symbol_old{},
		},
		globals: make(map[string]uintptr),
	}
	if err := compileStatements_old(&ctx, ast); err != nil {
		return Program{}, err
	}
	return Program{
		Instructions: ctx.instructions,
	}, nil
}

func compileStatements_old(ctx *Compiler_old, nodes []parser.BaseStatement) error {
	symtable := ctx.symtable
	ctx.symtable = SymbolTable_old{
		parent:  &symtable,
		symbols: map[string]Symbol_old{},
	}
	for i := range nodes {
		if err := compileBaseStatement_old(ctx, nodes[i]); err != nil {
			return err
		}
	}
	ctx.symtable = symtable
	return nil
}

func compileBaseStatement_old(ctx *Compiler_old, node parser.BaseStatement) error {
	switch node.StatementType() {
	case parser.TypedInitStatementType:
		return compileTypedInitStatement_old(ctx, node.(parser.TypedInitStatement))
	case parser.DeclarationStatementType:
		return compileDeclarationStatement_old(ctx, node.(parser.DeclarationStatement))
	case parser.FuncDefStatementType:
		return compileFuncDefStatement_old(ctx, node.(parser.FuncDefStatement))
	case parser.WhileStatementType:
		return compileWhileStatementType_old(ctx, node.(parser.WhileStatement))
	case parser.IfElseStatementType:
		return compileIfElseStatementType_old(ctx, node.(parser.IfElseStatement))
	case parser.IfStatementType:
		return compileIfStatementType_old(ctx, node.(parser.IfStatement))
	case parser.ReturnStatementType:
		return compileReturnStatement_old(ctx, node.(parser.ReturnStatement))
	case parser.ExpressionStatementType:
		return compileExpressionStatement_old(ctx, node)

	default:
		return fmt.Errorf("unknown or unexpected statement type '%s'", node.StatementType())
	}
}

func compileExpressionStatement_old(ctx *Compiler_old, node parser.BaseStatement) error {
	if err := compileBaseExpression_old(ctx, node.(parser.ExpressionStatement).Expression); err != nil {
		return err
	}

	// HACK
	//  After an expression we can be fairly certain to have a value on top of the stack.
	// 	after the expression the value is garbage, and needs to be cleaned up.
	//  It would be smart to know the type of this value at compile time,
	//  but that would require a fancy typechecker, which we don't have.
	//  But because of the stack value implementation we can just pop with any type.
	//  Notice. This hack also assumes funccalls and assignments always return a value,
	//  which isn't hard to enforce.
	ctx.instructions = append(ctx.instructions, Pop{Type: ctx.lastType})

	return nil
}

func compileTypedInitStatement_old(ctx *Compiler_old, node parser.TypedInitStatement) error {
	t, err := compileType_old(ctx, node.DeclType)
	if err != nil {
		return err
	}
	handle := ctx.nextVarId()
	ctx.symtable.Set(node.Identifier.StringValue, Symbol_old{Type: t, Handle: handle})
	ctx.instructions = append(ctx.instructions, DeclareLocal{Type: t, Handle: handle})
	if err := compileBaseExpression_old(ctx, node.Value); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, StoreLocal{Type: t, Handle: handle})
	return nil
}

func compileDeclarationStatement_old(ctx *Compiler_old, node parser.DeclarationStatement) error {
	t, err := compileType_old(ctx, node.DeclType)
	if err != nil {
		return err
	}
	handle := ctx.nextVarId()
	ctx.symtable.Set(node.Identifier.StringValue, Symbol_old{Type: t, Handle: handle})
	ctx.instructions = append(ctx.instructions, DeclareLocal{Type: t, Handle: handle})
	return nil
}

func compileFuncDefStatement_old(ctx *Compiler_old, node parser.FuncDefStatement) error {
	start := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, Jump{})
	ctx.globals[node.Identifier.StringValue] = uintptr(start + 2)
	for i := range node.Parameters {
		t, err := compileType_old(ctx, node.Parameters[i].DeclType)
		if err != nil {
			return err
		}
		handle := ctx.nextVarId()
		ctx.instructions = append(ctx.instructions, DeclareLocal{Type: t, Handle: handle})
		ctx.instructions = append(ctx.instructions, StoreLocal{Type: t, Handle: handle})
		ctx.symtable.Set(node.Parameters[i].Identifier.StringValue, Symbol_old{Type: t, Handle: handle})
	}
	if err := compileStatements_old(ctx, node.Body); err != nil {
		return err
	}
	t, err := compileType_old(ctx, node.ReturnType)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Push{Type: USIZE, Value: 0})
	ctx.instructions = append(ctx.instructions, Return{Type: t})
	fmt.Printf("bruhrbuhrubhr     %d\n", len(ctx.instructions)-start)
	ctx.instructions[start] = Push{Type: UPTR, Value: len(ctx.instructions) - start}
	return nil
}

func compileWhileStatementType_old(ctx *Compiler_old, node parser.WhileStatement) error {
	condition_start := len(ctx.instructions)
	if err := compileBaseExpression_old(ctx, node.Condition); err != nil {
		return err
	}
	end_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	if err := compileStatements_old(ctx, node.Body); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: condition_start})
	ctx.instructions = append(ctx.instructions, Jump{})
	ctx.instructions[end_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileIfElseStatementType_old(ctx *Compiler_old, node parser.IfElseStatement) error {
	if err := compileBaseExpression_old(ctx, node.Condition); err != nil {
		return err
	}
	else_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	if err := compileStatements_old(ctx, node.Truthy); err != nil {
		return err
	}
	end_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	ctx.instructions[else_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	if err := compileStatements_old(ctx, node.Falsy); err != nil {
		return err
	}
	ctx.instructions[end_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileIfStatementType_old(ctx *Compiler_old, node parser.IfStatement) error {
	if err := compileBaseExpression_old(ctx, node.Condition); err != nil {
		return err
	}
	end_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	if err := compileStatements_old(ctx, node.Body); err != nil {
		return err
	}
	ctx.instructions[end_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileReturnStatement_old(ctx *Compiler_old, node parser.ReturnStatement) error {
	if err := compileBaseExpression_old(ctx, node.Value); err != nil {
		return err
	}

	// HACK, type is hardcoded, should be inferred
	ctx.instructions = append(ctx.instructions, Return{Type: I32})
	return nil
}

func compileType_old(ctx *Compiler_old, t parser.Token) (Type, error) {
	switch t.StringValue {
	case "u8":
		return U8, nil
	case "u16":
		return U16, nil
	case "u32":
		return U32, nil
	case "u64":
		return U64, nil
	case "i8":
		return I8, nil
	case "i16":
		return I16, nil
	case "i32":
		return I32, nil
	case "i64":
		return I64, nil
	case "f32":
		return F32, nil
	case "f64":
		return F64, nil
	case "char":
		return CHAR, nil
	case "usize":
		return USIZE, nil
	case "uptr":
		return UPTR, nil
	default:
		return -1, fmt.Errorf("unknown type '%s'", t.StringValue)
	}
}

func compileBaseExpression_old(ctx *Compiler_old, node parser.BaseExpression) error {
	switch node.ExpressionType() {
	case parser.VarAssignExpressionType:
		return compileVarAssignExpression_old(ctx, node.(parser.VarAssignExpression))
	case parser.NotEqualExpressionType:
		return compileNotEqualExpression_old(ctx, node.(parser.NotEqualExpression))
	case parser.EqualExpressionType:
		return compileEqualExpression_old(ctx, node.(parser.EqualExpression))
	case parser.GTEExpressionType:
		return compileGreaterThanOrEqualExpression_old(ctx, node.(parser.GTEExpression))
	case parser.LTEExpressionType:
		return compileLessThanOrEqualExpression_old(ctx, node.(parser.LTEExpression))
	case parser.GreaterThanExpressionType:
		return compileGreaterThanExpression_old(ctx, node.(parser.GreaterThanExpression))
	case parser.LessThanExpressionType:
		return compileLessThanExpression_old(ctx, node.(parser.LessThanExpression))
	case parser.AddExpressionType:
		return compileAddExpression_old(ctx, node.(parser.AddExpression))
	case parser.SubExpressionType:
		return compileSubExpression_old(ctx, node.(parser.SubExpression))
	case parser.MulExpressionType:
		return compileMulExpression_old(ctx, node.(parser.MulExpression))
	case parser.DivExpressionType:
		return compileDivExpression_old(ctx, node.(parser.DivExpression))
	case parser.ExpExpressionType:
		return compileExpExpression_old(ctx, node.(parser.ExpExpression))
	case parser.FuncCallExpressionType:
		return compileFuncCallExpression_old(ctx, node.(parser.FuncCallExpression))
	case parser.VarAccessExpressionType:
		return compileVarAccessExpression_old(ctx, node.(parser.VarAccessExpression))
	case parser.IntExpressionType:
		return compileIntLiteral_old(ctx, node.(parser.IntLiteral))
	default:
		return fmt.Errorf("unknown or unexpected expression type '%s'", node.ExpressionType())
	}
}

func compileVarAssignExpression_old(ctx *Compiler_old, node parser.VarAssignExpression) error {
	if err := compileBaseExpression_old(ctx, node.Value); err != nil {
		return err
	}
	symbol, err := ctx.symtable.Get(node.Identifier.StringValue)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, StoreLocal{Type: symbol.Type, Handle: symbol.Handle})
	ctx.instructions = append(ctx.instructions, LoadLocal{Type: symbol.Type, Handle: symbol.Handle})
	ctx.lastType = symbol.Type
	return nil
}

func compileNotEqualExpression_old(ctx *Compiler_old, node parser.NotEqualExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpInequal{Type: I32})
	return nil
}

func compileEqualExpression_old(ctx *Compiler_old, node parser.EqualExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpEqual{Type: I32})
	return nil
}

func compileGreaterThanOrEqualExpression_old(ctx *Compiler_old, node parser.GTEExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpGTE{Type: I32})
	return nil
}

func compileLessThanOrEqualExpression_old(ctx *Compiler_old, node parser.LTEExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpLTE{Type: I32})
	return nil
}

func compileGreaterThanExpression_old(ctx *Compiler_old, node parser.GreaterThanExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpGT{Type: I32})
	return nil
}

func compileLessThanExpression_old(ctx *Compiler_old, node parser.LessThanExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpLT{Type: I32})
	return nil
}

func compileAddExpression_old(ctx *Compiler_old, node parser.AddExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Add{Type: I32})
	return nil
}

func compileSubExpression_old(ctx *Compiler_old, node parser.SubExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Subtract{Type: I32})
	return nil
}

func compileMulExpression_old(ctx *Compiler_old, node parser.MulExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Multiply{Type: I32})
	return nil
}

func compileDivExpression_old(ctx *Compiler_old, node parser.DivExpression) error {
	if err := compileBaseExpression_old(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression_old(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Divide{Type: I32})
	return nil
}

func compileExpExpression_old(ctx *Compiler_old, node parser.ExpExpression) error {
	var err error = nil
	err = compileBaseExpression_old(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileBaseExpression_old(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Exponent{Type: I32})
	return nil
}

func compileFuncCallExpression_old(ctx *Compiler_old, node parser.FuncCallExpression) error {
	for i := range node.Arguments {
		if err := compileBaseExpression_old(ctx, node.Arguments[i]); err != nil {
			return err
		}
	}

	// HACK
	if node.Identifier.ExpressionType() == parser.VarAccessExpressionType &&
		node.Identifier.(parser.VarAccessExpression).Identifier.StringValue == "syscall" {
		ctx.instructions = append(ctx.instructions, I32ToUsize{})
		ctx.instructions = append(ctx.instructions, Syscall{})
		ctx.instructions = append(ctx.instructions, Push{Type: I8, Value: 0})
		return nil
	}

	ctx.instructions = append(ctx.instructions, Push{Type: USIZE, Value: len(node.Arguments)})
	if err := compileBaseExpression_old(ctx, node.Identifier); err != nil {
		return err
	}
	// ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 1000000})
	// ctx.instructions = append(ctx.instructions, CmpGTE{Type: UPTR})
	// inotsyscall := len(ctx.instructions)
	// ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	// ctx.instructions = append(ctx.instructions, JumpIfZero{})
	// ctx.instructions = append(ctx.instructions, Syscall{})
	// isyscall := len(ctx.instructions)
	// ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	// ctx.instructions = append(ctx.instructions, Jump{})
	// ctx.instructions[inotsyscall] = Push{Type: UPTR, Value: len(ctx.instructions)}
	ctx.instructions = append(ctx.instructions, Call{Type: UPTR}) // type is omittable
	// ctx.instructions[isyscall] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileVarAccessExpression_old(ctx *Compiler_old, node parser.VarAccessExpression) error {
	for i := range ctx.globals {
		if i == node.Identifier.StringValue {
			ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: int(ctx.globals[i])})
			return nil
		}
	}
	symbol, err := ctx.symtable.Get(node.Identifier.StringValue)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, LoadLocal{Type: symbol.Type, Handle: symbol.Handle})
	ctx.lastType = symbol.Type
	return nil
}

func compileIntLiteral_old(ctx *Compiler_old, node parser.IntLiteral) error {
	ctx.instructions = append(ctx.instructions, Push{Type: I32, Value: node.Tok.IntValue})
	return nil
}
