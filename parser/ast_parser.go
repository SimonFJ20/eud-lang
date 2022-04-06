package parser

import "fmt"

func (p *Parser) makeStatement() []BaseStatement {
	statements := []BaseStatement{}
	for p.tok != nil {
		statements = append(statements, p.makeDeclaration())
	}
	return statements
}

func (p *Parser) makeDeclaration() BaseStatement {
	if p.tok == nil {
		panic("let token is nil")
	}
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
	if p.tok == nil {
		panic("identifier token is nil")
	}
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
	fmt.Println("assignment")
	id := p.tok
	left := p.makeAddition()
	if p.tok == nil {
		return left
	}
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
	fmt.Println("add")
	left := p.makeSubtraction()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == AddToken {
		p.next()
		right := p.makeAddition()
		fmt.Println("returned addition")
		return AddExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeSubtraction() BaseExpression {
	fmt.Println("sub")
	left := p.makeMultiplication()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == SubToken {
		p.next()
		fmt.Println("returned sub")
		right := p.makeSubtraction()
		return SubExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeMultiplication() BaseExpression {
	fmt.Println("mul")
	left := p.makeDivision()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == MulToken {
		p.next()
		fmt.Println("returned mul")
		right := p.makeMultiplication()
		return MulExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeDivision() BaseExpression {
	fmt.Println("div")
	left := p.makeExponentation()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == DivToken {
		p.next()
		fmt.Println("returned div")
		right := p.makeDivision()
		return DivExpression{LeftRightExpression{Left: left, Right: right}}
	} else {
		return left
	}
}

func (p *Parser) makeExponentation() BaseExpression {
	fmt.Println("exp")
	left := p.makeValue()
	if p.tok == nil {
		return left
	}
	if p.tok.Type == ExpToken {
		p.next()
		right := p.makeExponentation()
		fmt.Println("returned exp")
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
		fmt.Printf("%+v\n", expr)
		if p.tok.Type != RParenToken {
			panic("unexpected: tokenType != RParen")
		}
		p.next()
		return expr
	} else if p.tok.Type == IntToken {
		fmt.Println("returned int", p.tok.IntValue)
		return IntLiteral{
			Tok: p.tok,
		}
	} else if token.Type == IdentifierToken {
		fmt.Println("returned id")
		return IntLiteral{
			Tok: token,
		}
	} else {
		return p.makeExpression()
	}
}

func (p *Parser) next() {
	if p.tok.Next == nil {
		// finished parsing
	} else {
		fmt.Printf("%s\n", p.tok)
		p.tok = p.tok.Next
	}
}

func (p *Parser) Parse(t *Token) []BaseStatement {
	p.tok = t
	return p.makeStatement()
}