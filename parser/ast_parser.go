package parser

import "fmt"

func (p *Parser) makeStatement() []BaseStatement {
	statements := []BaseStatement{}
	for p.tok.Type != EOFToken && p.tok.Type != RBraceToken {
		statements = append(statements, p.makeFunction())
		p.next()
	}
	return statements
}

func (p *Parser) makeFunction() BaseStatement {
	if p.tok.Type == KeywordToken && p.tok.StringValue == "fn" {
		p.next()
		id := p.tok
		p.next()
		if p.tok.Type != LParenToken {
			panic(fmt.Sprintf("expected l_paren, got %s", p.tok.Type))
		}
		p.next()
		params := p.makeTypedDeclaration()
		p.next()
		if p.tok.Type != RParenToken {
			panic(fmt.Sprintf("expected r_paren, got %s", p.tok.Type))
		}
		p.next()
		if p.tok.Type != ColonToken {
			panic(fmt.Sprintf("expected colon, got %s", p.tok.Type))
		}
		p.next()
		if p.tok.Type != KeywordToken {
			panic("unrecognized type")
		}
		returnType := p.tok
		p.next()
		if p.tok.Type != LBraceToken {
			panic(fmt.Sprintf("expected l_brace, got %s", p.tok.Type))
		}
		p.next()
		body := p.makeStatement()
		if p.tok.Type != RBraceToken {
			panic(fmt.Sprintf("expected r_brace, got %s", p.tok.Type))
		}

		return FuncDefStatement {
			Identifier: *id,
			DeclType: *returnType,
			Parameters: []TypedDeclaration{params},
			Body: body,
		}
	} else {
		return p.makeDeclaration()
	}
}

func (p *Parser) makeDeclaration() BaseStatement {
	if p.tok.Type == KeywordToken && p.tok.StringValue == "let" {
		p.next()
		decl := p.makeTypedDeclaration()
		return DeclarationStatement {
			TypedDeclaration: decl,
		}
	} else {
		return ExpressionStatement{Expression: p.makeExpression()}
	}
}

func (p *Parser) makeTypedDeclaration() TypedDeclaration {
	if p.tok.Type == IdentifierToken {
		idToken := p.tok
		p.next()
		if p.tok.Type != ColonToken {
			panic("let declaration without type")
		}
		p.next()
		if p.tok.Type != KeywordToken {
			panic("unrecognized type")
		}
		idType := p.tok 
		return TypedDeclaration {
			DeclType:   *idType,
			Identifier: *idToken,
		}
	} else {
		panic(fmt.Sprintf("expected identifier, got %s", p.tok.Type))
	}
}

func (p *Parser) makeExpression() BaseExpression {
	return p.makeAssignment()
}

func (p *Parser) makeAssignment() BaseExpression {
	id := p.tok
	left := p.makeAddition()
	if p.tok.Type == AssignmentToken {
		p.next()
		value := p.makeAddition()
		return VarAssignExpression{
			Identifier: *id,
			Value:     value,
		}
	} else {
		return left
	}
}

func (p *Parser) makeAddition() BaseExpression {
	left := p.makeSubtraction()
	if p.tok.Type == AddToken {
		right := p.makeAddition()
		return AddExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeSubtraction() BaseExpression {
	left := p.makeMultiplication()
	if p.tok.Type == SubToken {
		right := p.makeSubtraction()
		return SubExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeMultiplication() BaseExpression {
	left := p.makeDivision()
	if p.tok.Type == MulToken {
		right := p.makeMultiplication()
		return MulExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeDivision() BaseExpression {
	left := p.makeExponentation()
	if p.tok.Type == DivToken {
		right := p.makeDivision()
		return DivExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeExponentation() BaseExpression {
	left := p.makeValue()
	if p.tok.Type == ExpToken {
		right := p.makeExponentation()
		return ExpExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeValue() BaseExpression {
	token := p.tok
	p.next()
	if token.Type == LParenToken {
		expr := p.makeExpression()
		if p.tok.Type != RParenToken {
			panic(fmt.Sprintf("unexpected tokenType, wanted r_paren, got %s", p.tok))
		}
		p.next()
		return expr
	} else if token.Type == IntToken {
		return IntLiteral{
			Tok: token,
		}
	} else if token.Type == IdentifierToken {
		return VarAccessExpression{
			Identifier: *token,
		}
	} else {
		return p.makeExpression()
	}
}

func (p *Parser) next() {
	if p.tok.Next == nil {
		// finished parsing
		fmt.Printf("%s\n",p.tok.Next.Type)
	} else {
		p.tok = p.tok.Next
	}
}

func (p *Parser) Parse(t *Token) []BaseStatement {
	p.tok = t
	return p.makeStatement()
}