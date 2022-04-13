package bytecode

import (
	"eud/parser"
	"fmt"
)

type Symbol struct {
	Type   Type
	Offset uint
}

type SymbolTable struct {
	parent  *SymbolTable
	symbols map[string]Symbol
}

func (s *SymbolTable) Set(name string, symbol Symbol) {
	s.symbols[name] = symbol
}

func (s *SymbolTable) Get(name string) (Symbol, error) {
	for i := range s.symbols {
		if i == name {
			return s.symbols[i], nil
		}
	}
	if s.parent != nil {
		return s.parent.Get(name)
	}
	return Symbol{}, fmt.Errorf("symbol \"%s\" undeclared", name)
}

func (s *SymbolTable) DefinedLocally(name string) bool {
	for i := range s.symbols {
		if i == name {
			return true
		}
	}
	return false
}

func (s *SymbolTable) IncreaseOffset() {
	for i := range s.symbols {
		s.symbols[i] = Symbol{
			Type:   s.symbols[i].Type,
			Offset: s.symbols[i].Offset + 1,
		}
	}
	if s.parent != nil {
		s.parent.IncreaseOffset()
	}
}

func (s *SymbolTable) DecreaseOffset() {
	for i := range s.symbols {
		s.symbols[i] = Symbol{
			Type:   s.symbols[i].Type,
			Offset: s.symbols[i].Offset - 1,
		}
	}
	if s.parent != nil {
		s.parent.DecreaseOffset()
	}
}

type Compiler struct {
	instructions []Instruction
	varId        uint
	symtable     SymbolTable
	globals      map[string]uintptr
	lastType     Type
}

func Compile(ast []parser.BaseStatement) (Program, error) {
	ctx := Compiler{
		instructions: []Instruction{},
		varId:        0,
		symtable: SymbolTable{
			parent:  nil,
			symbols: map[string]Symbol{},
		},
		globals: make(map[string]uintptr),
	}
	if err := compileStatements(&ctx, ast); err != nil {
		return Program{}, err
	}
	return Program{
		Instructions: ctx.instructions,
	}, nil
}

func compileStatements(ctx *Compiler, nodes []parser.BaseStatement) error {
	symtable := ctx.symtable
	ctx.symtable = SymbolTable{
		parent:  &symtable,
		symbols: map[string]Symbol{},
	}
	for i := range nodes {
		if err := compileBaseStatement(ctx, nodes[i]); err != nil {
			return err
		}
	}
	for range ctx.symtable.symbols {
		ctx.symtable.DecreaseOffset()
	}
	ctx.symtable = symtable
	return nil
}

func compileBaseStatement(ctx *Compiler, node parser.BaseStatement) error {
	switch node.StatementType() {
	case parser.TypedInitStatementType:
		return compileTypedInitStatement(ctx, node.(parser.TypedInitStatement))
	case parser.DeclarationStatementType:
		return compileDeclarationStatement(ctx, node.(parser.DeclarationStatement))
	case parser.FuncDefStatementType:
		return compileFuncDefStatement(ctx, node.(parser.FuncDefStatement))
	case parser.WhileStatementType:
		return compileWhileStatementType(ctx, node.(parser.WhileStatement))
	case parser.IfElseStatementType:
		return compileIfElseStatementType(ctx, node.(parser.IfElseStatement))
	case parser.IfStatementType:
		return compileIfStatementType(ctx, node.(parser.IfStatement))
	case parser.ReturnStatementType:
		return compileReturnStatement(ctx, node.(parser.ReturnStatement))
	case parser.ExpressionStatementType:
		return compileExpressionStatement(ctx, node)

	default:
		return fmt.Errorf("unknown or unexpected statement type '%s'", node.StatementType())
	}
}

func compileExpressionStatement(ctx *Compiler, node parser.BaseStatement) error {
	if err := compileBaseExpression(ctx, node.(parser.ExpressionStatement).Expression); err != nil {
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

func compileTypedInitStatement(ctx *Compiler, node parser.TypedInitStatement) error {
	t, err := compileType(ctx, node.DeclType)
	if err != nil {
		return err
	}
	ctx.symtable.IncreaseOffset()
	ctx.symtable.Set(node.Identifier.StringValue, Symbol{Type: t, Offset: 0})
	ctx.instructions = append(ctx.instructions, DeclareLocal{Type: t})
	if err := compileBaseExpression(ctx, node.Value); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, StoreLocal{Type: t, Offset: 0})
	return nil
}

func compileDeclarationStatement(ctx *Compiler, node parser.DeclarationStatement) error {
	t, err := compileType(ctx, node.DeclType)
	if err != nil {
		return err
	}
	ctx.symtable.IncreaseOffset()
	ctx.symtable.Set(node.Identifier.StringValue, Symbol{Type: t, Offset: 0})
	ctx.instructions = append(ctx.instructions, DeclareLocal{Type: t})
	return nil
}

func compileFuncDefStatement(ctx *Compiler, node parser.FuncDefStatement) error {
	start := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, Jump{})
	ctx.globals[node.Identifier.StringValue] = uintptr(start + 2)
	for i := range node.Parameters {
		t, err := compileType(ctx, node.Parameters[i].DeclType)
		if err != nil {
			return err
		}
		ctx.symtable.IncreaseOffset()
		ctx.instructions = append(ctx.instructions, DeclareLocal{Type: t})
		ctx.instructions = append(ctx.instructions, StoreLocal{Type: t, Offset: 0})
		ctx.symtable.Set(node.Parameters[i].Identifier.StringValue, Symbol{Type: t, Offset: 0})
	}
	if err := compileStatements(ctx, node.Body); err != nil {
		return err
	}
	t, err := compileType(ctx, node.ReturnType)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Push{Type: USIZE, Value: 0})
	ctx.instructions = append(ctx.instructions, Return{Type: t})
	ctx.instructions[start] = Push{Type: UPTR, Value: len(ctx.instructions) - start}
	return nil
}

func compileWhileStatementType(ctx *Compiler, node parser.WhileStatement) error {
	condition_start := len(ctx.instructions)
	if err := compileBaseExpression(ctx, node.Condition); err != nil {
		return err
	}
	end_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	if err := compileStatements(ctx, node.Body); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: condition_start})
	ctx.instructions = append(ctx.instructions, Jump{})
	ctx.instructions[end_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileIfElseStatementType(ctx *Compiler, node parser.IfElseStatement) error {
	if err := compileBaseExpression(ctx, node.Condition); err != nil {
		return err
	}
	else_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	if err := compileStatements(ctx, node.Truthy); err != nil {
		return err
	}
	end_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	ctx.instructions[else_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	if err := compileStatements(ctx, node.Falsy); err != nil {
		return err
	}
	ctx.instructions[end_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileIfStatementType(ctx *Compiler, node parser.IfStatement) error {
	if err := compileBaseExpression(ctx, node.Condition); err != nil {
		return err
	}
	end_jpush_index := len(ctx.instructions)
	ctx.instructions = append(ctx.instructions, Push{Type: UPTR, Value: 0})
	ctx.instructions = append(ctx.instructions, JumpIfZero{})
	if err := compileStatements(ctx, node.Body); err != nil {
		return err
	}
	ctx.instructions[end_jpush_index] = Push{Type: UPTR, Value: len(ctx.instructions)}
	return nil
}

func compileReturnStatement(ctx *Compiler, node parser.ReturnStatement) error {
	if err := compileBaseExpression(ctx, node.Value); err != nil {
		return err
	}

	// HACK, type is hardcoded, should be inferred
	ctx.instructions = append(ctx.instructions, Return{Type: I32})
	return nil
}

func compileType(ctx *Compiler, t parser.Token) (Type, error) {
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

func compileBaseExpression(ctx *Compiler, node parser.BaseExpression) error {
	switch node.ExpressionType() {
	case parser.VarAssignExpressionType:
		return compileVarAssignExpression(ctx, node.(parser.VarAssignExpression))
	case parser.NotEqualExpressionType:
		return compileNotEqualExpression(ctx, node.(parser.NotEqualExpression))
	case parser.EqualExpressionType:
		return compileEqualExpression(ctx, node.(parser.EqualExpression))
	case parser.GTEExpressionType:
		return compileGreaterThanOrEqualExpression(ctx, node.(parser.GTEExpression))
	case parser.LTEExpressionType:
		return compileLessThanOrEqualExpression(ctx, node.(parser.LTEExpression))
	case parser.GreaterThanExpressionType:
		return compileGreaterThanExpression(ctx, node.(parser.GreaterThanExpression))
	case parser.LessThanExpressionType:
		return compileLessThanExpression(ctx, node.(parser.LessThanExpression))
	case parser.AddExpressionType:
		return compileAddExpression(ctx, node.(parser.AddExpression))
	case parser.SubExpressionType:
		return compileSubExpression(ctx, node.(parser.SubExpression))
	case parser.MulExpressionType:
		return compileMulExpression(ctx, node.(parser.MulExpression))
	case parser.DivExpressionType:
		return compileDivExpression(ctx, node.(parser.DivExpression))
	case parser.ExpExpressionType:
		return compileExpExpression(ctx, node.(parser.ExpExpression))
	case parser.FuncCallExpressionType:
		return compileFuncCallExpression(ctx, node.(parser.FuncCallExpression))
	case parser.VarAccessExpressionType:
		return compileVarAccessExpression(ctx, node.(parser.VarAccessExpression))
	case parser.IntExpressionType:
		return compileIntLiteral(ctx, node.(parser.IntLiteral))
	default:
		return fmt.Errorf("unknown or unexpected expression type '%s'", node.ExpressionType())
	}
}

func compileVarAssignExpression(ctx *Compiler, node parser.VarAssignExpression) error {
	if err := compileBaseExpression(ctx, node.Value); err != nil {
		return err
	}
	symbol, err := ctx.symtable.Get(node.Identifier.StringValue)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, StoreLocal{Type: symbol.Type, Offset: symbol.Offset})
	ctx.instructions = append(ctx.instructions, LoadLocal{Type: symbol.Type, Offset: symbol.Offset})
	ctx.lastType = symbol.Type
	return nil
}

func compileNotEqualExpression(ctx *Compiler, node parser.NotEqualExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpInequal{Type: I32})
	return nil
}

func compileEqualExpression(ctx *Compiler, node parser.EqualExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpEqual{Type: I32})
	return nil
}

func compileGreaterThanOrEqualExpression(ctx *Compiler, node parser.GTEExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpGTE{Type: I32})
	return nil
}

func compileLessThanOrEqualExpression(ctx *Compiler, node parser.LTEExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpLTE{Type: I32})
	return nil
}

func compileGreaterThanExpression(ctx *Compiler, node parser.GreaterThanExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpGT{Type: I32})
	return nil
}

func compileLessThanExpression(ctx *Compiler, node parser.LessThanExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, CmpLT{Type: I32})
	return nil
}

func compileAddExpression(ctx *Compiler, node parser.AddExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Add{Type: I32})
	return nil
}

func compileSubExpression(ctx *Compiler, node parser.SubExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Subtract{Type: I32})
	return nil
}

func compileMulExpression(ctx *Compiler, node parser.MulExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Multiply{Type: I32})
	return nil
}

func compileDivExpression(ctx *Compiler, node parser.DivExpression) error {
	if err := compileBaseExpression(ctx, node.Left); err != nil {
		return err
	}
	if err := compileBaseExpression(ctx, node.Right); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Divide{Type: I32})
	return nil
}

func compileExpExpression(ctx *Compiler, node parser.ExpExpression) error {
	var err error = nil
	err = compileBaseExpression(ctx, node.Left)
	if err != nil {
		return err
	}
	err = compileBaseExpression(ctx, node.Right)
	if err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Exponent{Type: I32})
	return nil
}

func compileFuncCallExpression(ctx *Compiler, node parser.FuncCallExpression) error {
	for i := range node.Arguments {
		if err := compileBaseExpression(ctx, node.Arguments[i]); err != nil {
			return err
		}
	}

	ctx.instructions = append(ctx.instructions, Push{Type: USIZE, Value: len(node.Arguments)})
	if err := compileBaseExpression(ctx, node.Identifier); err != nil {
		return err
	}
	ctx.instructions = append(ctx.instructions, Call{Type: UPTR}) // type is omittable
	return nil
}

func compileVarAccessExpression(ctx *Compiler, node parser.VarAccessExpression) error {
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
	ctx.instructions = append(ctx.instructions, LoadLocal{Type: symbol.Type, Offset: symbol.Offset})
	ctx.lastType = symbol.Type
	return nil
}

func compileIntLiteral(ctx *Compiler, node parser.IntLiteral) error {
	ctx.instructions = append(ctx.instructions, Push{Type: I32, Value: node.Tok.IntValue})
	return nil
}
