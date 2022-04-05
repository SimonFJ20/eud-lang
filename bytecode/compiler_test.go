package bytecode_test

import (
	"eud/bytecode"
	"eud/parser"
	"testing"
)

func TestVariables(t *testing.T) {
	program, err := bytecode.Compile([]parser.BaseStatement{
		parser.DeclarationStatement{
			TypedDeclaration: parser.TypedDeclaration{
				DeclType: parser.Token{
					Type: parser.KeywordToken, Value: 0, Text: "i32", Next: nil,
				},
				Identifier: parser.Token{
					Type: parser.IdentifierToken, Value: 0, Text: "a", Next: nil,
				},
			},
		},
		parser.ExpressionStatement{
			Expression: parser.VarAssignExpression{
				Identifier: parser.Token{
					Type: parser.IdentifierToken, Value: 0, Text: "a", Next: nil,
				},
				Value: parser.IntLiteral{
					Tok: &parser.Token{
						Type: parser.IntToken, Value: 60, Text: "60", Next: nil,
					},
				},
			},
		},
		parser.ExpressionStatement{
			Expression: parser.VarAssignExpression{
				Identifier: parser.Token{
					Type: parser.IdentifierToken, Value: 0, Text: "a", Next: nil,
				},
				Value: parser.AddExpression{
					LeftRightExpression: parser.LeftRightExpression{
						Left: parser.VarAccessExpression{
							Identifier: parser.Token{
								Type: parser.IdentifierToken, Value: 0, Text: "a", Next: nil,
							},
						},
						Right: parser.IntLiteral{
							Tok: &parser.Token{
								Type: parser.IntToken, Value: 5, Text: "5", Next: nil,
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
									Type: parser.IntToken, Value: 3, Text: "3", Next: nil,
								},
							},
							Right: parser.IntLiteral{
								Tok: &parser.Token{
									Type: parser.IntToken, Value: 4, Text: "4", Next: nil,
								},
							},
						},
					},
					Right: parser.IntLiteral{
						Tok: &parser.Token{
							Type: parser.IntToken, Value: 5, Text: "5", Next: nil,
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
