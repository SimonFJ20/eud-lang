package bytecode_test

import (
	"eud/bytecode"
	"eud/parser"
	"fmt"
	"testing"
)

func TestFunctions(t *testing.T) {
	program, err := bytecode.Compile([]parser.BaseStatement{
		parser.FuncDefStatement{
			Identifier: parser.Token{
				Type: parser.IdentifierToken, StringValue: "sum", Next: nil,
			},
			ReturnType: parser.Token{
				Type: parser.KeywordToken, StringValue: "i32", Next: nil,
			},
			Parameters: []parser.TypedDeclaration{
				{
					DeclType: parser.Token{
						Type: parser.KeywordToken, StringValue: "i32", Next: nil,
					},
					Identifier: parser.Token{
						Type: parser.IdentifierToken, StringValue: "a", Next: nil,
					},
				},
				{
					DeclType: parser.Token{
						Type: parser.KeywordToken, StringValue: "i32", Next: nil,
					},
					Identifier: parser.Token{
						Type: parser.IdentifierToken, StringValue: "b", Next: nil,
					},
				},
			},
			Body: []parser.BaseStatement{
				parser.ReturnStatement{
					Value: parser.AddExpression{
						LeftRightExpression: parser.LeftRightExpression{
							Left: parser.VarAccessExpression{
								Identifier: parser.Token{
									Type: parser.IdentifierToken, StringValue: "a", Next: nil,
								},
							},
							Right: parser.VarAccessExpression{
								Identifier: parser.Token{
									Type: parser.IdentifierToken, StringValue: "b", Next: nil,
								},
							},
						},
					},
				},
			},
		},
		parser.DeclarationStatement{
			TypedDeclaration: parser.TypedDeclaration{
				DeclType: parser.Token{
					Type: parser.KeywordToken, StringValue: "i32", Next: nil,
				},
				Identifier: parser.Token{
					Type: parser.IdentifierToken, StringValue: "result", Next: nil,
				},
			},
		},
		parser.ExpressionStatement{
			Expression: parser.VarAssignExpression{
				Identifier: parser.Token{
					Type: parser.IdentifierToken, StringValue: "result", Next: nil,
				},
				Value: parser.FuncCallExpression{
					Identifier: parser.VarAccessExpression{
						Identifier: parser.Token{
							Type: parser.IdentifierToken, StringValue: "sum", Next: nil,
						},
					},
					Arguments: []parser.BaseExpression{
						parser.IntLiteral{
							Tok: &parser.Token{
								Type: parser.IntToken, IntValue: 5, StringValue: "5", Next: nil,
							},
						},
						parser.IntLiteral{
							Tok: &parser.Token{
								Type: parser.IntToken, IntValue: 3, StringValue: "3", Next: nil,
							},
						},
					},
				},
			},
		},
	})
	for i := range program.Instructions {
		fmt.Printf("    %d\t%s\n", i, program.Instructions[i].String())
	}
	fmt.Println("")
	if err != nil {
		t.Error(err)
	}
	program.RunWithDebug = true
	program.Instructions = append(program.Instructions, bytecode.LoadLocal{Type: bytecode.I32, Handle: 2})
	runtime := bytecode.Run(program)
	result := runtime.Stack[0].(bytecode.I32Value).Value
	if result != 8 {
		t.Errorf("unexpected result %d", result)
	}
}

func TestVariables(t *testing.T) {
	program, err := bytecode.Compile([]parser.BaseStatement{
		parser.DeclarationStatement{
			TypedDeclaration: parser.TypedDeclaration{
				DeclType: parser.Token{
					Type: parser.KeywordToken, StringValue: "i32", Next: nil,
				},
				Identifier: parser.Token{
					Type: parser.IdentifierToken, StringValue: "a", Next: nil,
				},
			},
		},
		parser.ExpressionStatement{
			Expression: parser.VarAssignExpression{
				Identifier: parser.Token{
					Type: parser.IdentifierToken, StringValue: "a", Next: nil,
				},
				Value: parser.IntLiteral{
					Tok: &parser.Token{
						Type: parser.IntToken, IntValue: 60, StringValue: "60", Next: nil,
					},
				},
			},
		},
		parser.ExpressionStatement{
			Expression: parser.VarAssignExpression{
				Identifier: parser.Token{
					Type: parser.IdentifierToken, StringValue: "a", Next: nil,
				},
				Value: parser.AddExpression{
					LeftRightExpression: parser.LeftRightExpression{
						Left: parser.VarAccessExpression{
							Identifier: parser.Token{
								Type: parser.IdentifierToken, StringValue: "a", Next: nil,
							},
						},
						Right: parser.IntLiteral{
							Tok: &parser.Token{
								Type: parser.IntToken, IntValue: 5, StringValue: "5", Next: nil,
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	program.RunWithDebug = true
	program.Instructions = append(program.Instructions, bytecode.LoadLocal{Type: bytecode.I32, Handle: 0})
	runtime := bytecode.Run(program)
	result := runtime.Stack[0].(bytecode.I32Value).Value
	if result != 65 {
		t.Errorf("unexpected result %d", result)
	}
}

func TestMath(t *testing.T) {
	program, err := bytecode.Compile([]parser.BaseStatement{
		parser.ExpressionStatement{
			Expression: parser.MulExpression{
				LeftRightExpression: parser.LeftRightExpression{
					Left: parser.AddExpression{
						LeftRightExpression: parser.LeftRightExpression{
							Left: parser.IntLiteral{
								Tok: &parser.Token{
									Type: parser.IntToken, IntValue: 3, StringValue: "3", Next: nil,
								},
							},
							Right: parser.IntLiteral{
								Tok: &parser.Token{
									Type: parser.IntToken, IntValue: 4, StringValue: "4", Next: nil,
								},
							},
						},
					},
					Right: parser.IntLiteral{
						Tok: &parser.Token{
							Type: parser.IntToken, IntValue: 5, StringValue: "5", Next: nil,
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	program.RunWithDebug = true
	bytecode.Run(program)
}
