package parser

import "fmt"

func (p *Parser) makeStatement() []BaseStatement {
	statements := []BaseStatement{}
	for p.tok.Type != EOFToken {
		statements = append(statements, p.makeFunction())
		if p.tok.Type == RBraceToken {
			break
		}
		p.next()
	}
	fmt.Printf("%s\n", p.tok)
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
		p.next()

		return FuncDefStatement{
			Identifier: *id,
			ReturnType: *returnType,
			Parameters: []TypedDeclaration{params},
			Body:       body,
		}
	} else {
		return p.makeDeclaration()
	}
}

func (p *Parser) makeDeclaration() BaseStatement {
	if p.tok.Type == KeywordToken && p.tok.StringValue == "let" {
		p.next()
		decl := p.makeTypedDeclaration()
		return DeclarationStatement{
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
		return TypedDeclaration{
			DeclType:   *idType,
			Identifier: *idToken,
		}
	} else {
		panic(fmt.Sprintf("expected identifier, got %s", p.tok.Type))
	}
}

func (p *Parser) makeExpression() BaseExpression {
	fmt.Printf("expr, %s\n", p.tok)
	return p.makeAssignment()
}

func (p *Parser) makeAssignment() BaseExpression {
	fmt.Printf("assign, %s\n", p.tok)
	id := p.tok
	left := p.makeAddition()
	if p.tok.Type == AssignmentToken {
		value := p.makeAddition()
		return VarAssignExpression{
			Identifier: *id,
			Value:      value,
		}
	} else {
		return left
	}
}

func (p *Parser) makeAddition() BaseExpression {
	fmt.Printf("add, %s\n", p.tok)
	left := p.makeSubtraction()
	if p.tok.Type == AddToken {
		right := p.makeAddition()
		return AddExpression{Left: left, Right: right}
	} else {
		return left
	}
}

func (p *Parser) makeSubtraction() BaseExpression {
	fmt.Printf("sub, %s\n", p.tok)
	left := p.makeMultiplication()
	if p.tok.Type == SubToken {
		right := p.makeSubtraction()
		return SubExpression{Left: left, Right: right}
	} else {
		return left
	}
}

func (p *Parser) makeMultiplication() BaseExpression {
	fmt.Printf("mul, %s\n", p.tok)
	left := p.makeDivision()
	if p.tok.Type == MulToken {
		right := p.makeMultiplication()
		return MulExpression{Left: left, Right: right}
	} else {
		return left
	}
}

func (p *Parser) makeDivision() BaseExpression {
	fmt.Printf("div, %s\n", p.tok)
	left := p.makeExponentation()
	if p.tok.Type == DivToken {
		right := p.makeDivision()
		return DivExpression{Left: left, Right: right}
	} else {
		return left
	}
}

func (p *Parser) makeExponentation() BaseExpression {
	fmt.Printf("exp, %s\n", p.tok)
	left := p.makeValue()
	if p.tok.Type == ExpToken {
		right := p.makeExponentation()
		return ExpExpression{Left: left, Right: right}
	} else {
		return left
	}
}

func (p *Parser) makeValue() BaseExpression {
	fmt.Printf("val, %s\n", p.tok)
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
	} else {
		p.tok = p.tok.Next
	}
}

func (p *Parser) prev() {
	if p.tok.Prev == nil {
		// unable
	} else {
		p.tok = p.tok.Prev
	}
}

func (p *Parser) Parse(t *Token) []BaseStatement {
	p.tok = t
	return p.makeStatement()
}
