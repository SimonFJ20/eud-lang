package astjson

import (
	"encoding/json"
	"eud/parser"
	"fmt"
	"log"
	"strconv"
)

type IElement interface{ GetType() string }
type Element struct {
	Type string `json:"type"`
	IElement
}

type Position struct {
	Type string `json:"type"`
	IElement
	Row      int    `json:"row"`
	Col      int    `json:"col"`
	Filename string `json:"filename"`
}

type Token struct {
	Type string `json:"type"`
	IElement
	TokenType string `json:"tokenType"`
	Value     string `json:"value"`
}

type IStatementNode interface{ GetType() string }
type StatementNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
}

type IExpressionNode interface{ GetType() string }
type ExpressionNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
}

type TypeNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Token Token `json:"token"`
}

type TypedDeclNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token    `json:"target"`
	ValueType TypeNode `json:"valueType"`
}

type FuncDefNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token            `json:"target"`
	ValueType TypeNode         `json:"valueType"`
	Params    []TypedDeclNode  `json:"params"`
	Body      []IStatementNode `json:"body"`
}

type WhileNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Condition IExpressionNode  `json:"value"`
	Body      []IStatementNode `json:"body"`
}

type IfElseNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Condition IExpressionNode  `json:"value"`
	Truthy    []IStatementNode `json:"truthy"`
	Falsy     []IStatementNode `json:"falsy"`
}

type IfNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Condition IExpressionNode  `json:"value"`
	Body      []IStatementNode `json:"body"`
}

type VarDeclNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target    Token    `json:"target"`
	ValueType TypeNode `json:"valueType"`
}

type AssignNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target Token           `json:"target"`
	Value  IExpressionNode `json:"value"`
}

type NotEqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type EqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type GreaterThanOrEqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type LessThanOrEqualNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type GreaterThanNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type LessThanNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type AddNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type SubNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type MulNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type DivNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type ModNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type ExpNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Left  IExpressionNode `json:"left"`
	Right IExpressionNode `json:"right"`
}

type FuncCallNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Target IExpressionNode   `json:"target"`
	Args   []IExpressionNode `json:"args"`
}

type IntNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Token Token `json:"token"`
}

type VarNode struct {
	Type string `json:"type"`
	IElement
	Filepos Position `json:"fp"`
	IStatementNode
	IExpressionNode
	Token Token `json:"token"`
}

func (e Element) GetType() string                { return e.Type }
func (e Position) GetType() string               { return e.Type }
func (e Token) GetType() string                  { return e.Type }
func (e StatementNode) GetType() string          { return e.Type }
func (e ExpressionNode) GetType() string         { return e.Type }
func (e TypeNode) GetType() string               { return e.Type }
func (e TypedDeclNode) GetType() string          { return e.Type }
func (e FuncDefNode) GetType() string            { return e.Type }
func (e WhileNode) GetType() string              { return e.Type }
func (e IfElseNode) GetType() string             { return e.Type }
func (e IfNode) GetType() string                 { return e.Type }
func (e VarDeclNode) GetType() string            { return e.Type }
func (e AssignNode) GetType() string             { return e.Type }
func (e NotEqualNode) GetType() string           { return e.Type }
func (e EqualNode) GetType() string              { return e.Type }
func (e GreaterThanOrEqualNode) GetType() string { return e.Type }
func (e LessThanOrEqualNode) GetType() string    { return e.Type }
func (e GreaterThanNode) GetType() string        { return e.Type }
func (e LessThanNode) GetType() string           { return e.Type }
func (e AddNode) GetType() string                { return e.Type }
func (e SubNode) GetType() string                { return e.Type }
func (e MulNode) GetType() string                { return e.Type }
func (e DivNode) GetType() string                { return e.Type }
func (e ModNode) GetType() string                { return e.Type }
func (e ExpNode) GetType() string                { return e.Type }
func (e FuncCallNode) GetType() string           { return e.Type }
func (e IntNode) GetType() string                { return e.Type }
func (e VarNode) GetType() string                { return e.Type }

type Object = map[string]interface{}

func Parse(data string) []parser.BaseStatement {
	var raw []Object
	json.Unmarshal([]byte(data), &raw)
	// fmt.Printf("%s", ast)
	elements := []IStatementNode{}
	for i := range raw {
		elements = append(elements, ParseJsonElement(raw[i]))
	}
	return ParseStatements(elements)
}

func ParseJsonElement(raw Object) IElement {
	switch raw["type"] {
	case "Position":
		return Position{
			Type:     raw["type"].(string),
			Row:      int(raw["row"].(float64)),
			Col:      int(raw["col"].(float64)),
			Filename: raw["filename"].(string),
		}
	case "Token":
		return Token{
			Type:      raw["type"].(string),
			TokenType: raw["tokenType"].(string),
			Value:     raw["value"].(string),
		}
	case "FuncDefNode":
		params := []TypedDeclNode{}
		for i := range raw["params"].([]interface{}) {
			params = append(params, ParseJsonElement(raw["params"].([]interface{})[i].(Object)).(TypedDeclNode))
		}
		body := []IStatementNode{}
		for i := range raw["body"].([]interface{}) {
			body = append(body, ParseJsonElement(raw["body"].([]interface{})[i].(Object)).(IStatementNode))
		}
		return FuncDefNode{
			Type:      raw["type"].(string),
			Filepos:   ParseJsonElement(raw["fp"].(Object)).(Position),
			Target:    ParseJsonElement(raw["target"].(Object)).(Token),
			ValueType: ParseJsonElement(raw["valueType"].(Object)).(TypeNode),
			Params:    params,
			Body:      body,
		}
	case "WhileNode":
		body := []IStatementNode{}
		for i := range raw["body"].([]interface{}) {
			body = append(body, ParseJsonElement(raw["body"].([]interface{})[i].(Object)).(IStatementNode))
		}
		return IfNode{
			Type:      raw["type"].(string),
			Filepos:   ParseJsonElement(raw["fp"].(Object)).(Position),
			Condition: ParseJsonElement(raw["condition"].(Object)).(IExpressionNode),
			Body:      body,
		}
	case "IfElseNode":
		body_truthy := []IStatementNode{}
		for i := range raw["body"].([]interface{}) {
			body_truthy = append(body_truthy, ParseJsonElement(raw["body"].([]interface{})[i].(Object)).(IStatementNode))
		}
		body_falsy := []IStatementNode{}
		for i := range raw["body"].([]interface{}) {
			body_falsy = append(body_falsy, ParseJsonElement(raw["body"].([]interface{})[i].(Object)).(IStatementNode))
		}
		return IfElseNode{
			Type:      raw["type"].(string),
			Filepos:   ParseJsonElement(raw["fp"].(Object)).(Position),
			Condition: ParseJsonElement(raw["condition"].(Object)).(IExpressionNode),
			Truthy:    body_truthy,
			Falsy:     body_falsy,
		}
	case "IfNode":
		body := []IStatementNode{}
		for i := range raw["body"].([]interface{}) {
			body = append(body, ParseJsonElement(raw["body"].([]interface{})[i].(Object)).(IStatementNode))
		}
		return IfNode{
			Type:      raw["type"].(string),
			Filepos:   ParseJsonElement(raw["fp"].(Object)).(Position),
			Condition: ParseJsonElement(raw["condition"].(Object)).(IExpressionNode),
			Body:      body,
		}
	case "VarDeclNode":
		return VarDeclNode{
			Type:      raw["type"].(string),
			Filepos:   ParseJsonElement(raw["fp"].(Object)).(Position),
			Target:    ParseJsonElement(raw["target"].(Object)).(Token),
			ValueType: ParseJsonElement(raw["valueType"].(Object)).(TypeNode),
		}
	case "TypedDeclNode":
		return TypedDeclNode{
			Target:    ParseJsonElement(raw["target"].(Object)).(Token),
			ValueType: ParseJsonElement(raw["valueType"].(Object)).(TypeNode),
		}
	case "TypeNode":
		return TypeNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Token:   ParseJsonElement(raw["token"].(Object)).(Token),
		}
	case "AssignNode":
		return AssignNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Target:  ParseJsonElement(raw["target"].(Object)).(Token),
			Value:   ParseJsonElement(raw["value"].(Object)).(IExpressionNode),
		}
	case "NotEqualNode":
		return NotEqualNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "EqualNode":
		return EqualNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "GreaterThanOrEqualNode":
		return GreaterThanOrEqualNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "LessThanOrEqualNode":
		return LessThanOrEqualNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "GreaterThanNode":
		return GreaterThanNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "LessThanNode":
		return LessThanNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "AddNode":
		return AddNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "SubNode":
		return SubNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "MulNode":
		return MulNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "DivNode":
		return DivNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "ExpNode":
		return ExpNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Left:    ParseJsonElement(raw["left"].(Object)).(IExpressionNode),
			Right:   ParseJsonElement(raw["right"].(Object)).(IExpressionNode),
		}
	case "FuncCallNode":
		args := []IExpressionNode{}
		for i := range raw["args"].([]interface{}) {
			args = append(args, ParseJsonElement(raw["args"].([]interface{})[i].(Object)).(IExpressionNode))
		}
		return FuncCallNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Target:  ParseJsonElement(raw["target"].(Object)).(IExpressionNode),
			Args:    args,
		}
	case "VarNode":
		return VarNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Token:   ParseJsonElement(raw["token"].(Object)).(Token),
		}
	case "IntNode":
		return IntNode{
			Type:    raw["type"].(string),
			Filepos: ParseJsonElement(raw["fp"].(Object)).(Position),
			Token:   ParseJsonElement(raw["token"].(Object)).(Token),
		}
	default:
		log.Fatalf("json element type '%s' not implemented", raw["type"])
	}
	panic("unreachable")
}

func ParseStatements(elements []IStatementNode) []parser.BaseStatement {
	res := []parser.BaseStatement{}
	for i := range elements {
		res = append(res, ParseStatement(elements[i]))
	}
	return res
}

func ParseStatement(element IStatementNode) parser.BaseStatement {
	switch element.GetType() {
	case "FuncDefNode":
		n := element.(FuncDefNode)
		params := ParseTypedDeclarations(n.Params)
		body := ParseStatements(n.Body)
		return parser.FuncDefStatement{
			Identifier: n.Target.Convert(),
			ReturnType: n.ValueType.Token.Convert(),
			Parameters: params,
			Body:       body,
		}
	case "WhileNode":
		n := element.(WhileNode)
		condition := ParseBaseExpression(n.Condition)
		body := ParseStatements(n.Body)
		return parser.WhileStatement{
			Condition: condition,
			Body:      body,
		}
	case "IfElseNode":
		n := element.(IfElseNode)
		condition := ParseBaseExpression(n.Condition)
		body_truthy := ParseStatements(n.Truthy)
		body_falsy := ParseStatements(n.Falsy)
		return parser.IfElseStatement{
			Condition: condition,
			Truthy:    body_truthy,
			Falsy:     body_falsy,
		}
	case "IfNode":
		n := element.(IfNode)
		condition := ParseBaseExpression(n.Condition)
		body := ParseStatements(n.Body)
		return parser.IfStatement{
			Condition: condition,
			Body:      body,
		}
	case "VarDeclNode":
		n := element.(VarDeclNode)
		return parser.DeclarationStatement{
			TypedDeclaration: ParseTypedDeclaration(TypedDeclNode{
				Type:      n.Type,
				Filepos:   n.Filepos,
				Target:    n.Target,
				ValueType: n.ValueType,
			}),
		}
	case "AssignNode":
		n := element.(AssignNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "NotEqualNode":
		n := element.(NotEqualNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "EqualNode":
		n := element.(EqualNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "GreaterThanOrEqualNode":
		n := element.(GreaterThanOrEqualNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "LessThanOrEqualNode":
		n := element.(LessThanOrEqualNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "GreaterThanNode":
		n := element.(GreaterThanNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "LessThanNode":
		n := element.(LessThanNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "AddNode":
		n := element.(AddNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "SubNode":
		n := element.(SubNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "MulNode":
		n := element.(MulNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "DivNode":
		n := element.(DivNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "ModNode":
		n := element.(ModNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "ExpNode":
		n := element.(ExpNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "FuncCallNode":
		n := element.(FuncCallNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "IntNode":
		n := element.(IntNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	case "VarNode":
		n := element.(VarNode)
		return parser.ExpressionStatement{
			Expression: ParseBaseExpression(n),
		}
	default:
		panic(fmt.Sprintf("'%s' unexpected", element.(IElement).GetType()))
	}

}

func ParseTypedDeclarations(elements []TypedDeclNode) []parser.TypedDeclaration {
	res := []parser.TypedDeclaration{}
	for i := range elements {
		res = append(res, ParseTypedDeclaration(elements[i]))
	}
	return res
}

func ParseTypedDeclaration(elements TypedDeclNode) parser.TypedDeclaration {
	return parser.TypedDeclaration{
		Identifier: elements.Target.Convert(),
		DeclType:   elements.ValueType.Token.Convert(),
	}
}

func ParseBaseExpressions(elements []IExpressionNode) []parser.BaseExpression {
	res := []parser.BaseExpression{}
	for i := range elements {
		res = append(res, ParseBaseExpression(elements[i]))
	}
	return res
}

func ParseBaseExpression(element IExpressionNode) parser.BaseExpression {
	switch element.GetType() {
	case "AssignNode":
		n := element.(AssignNode)
		return parser.VarAssignExpression{
			Identifier: n.Target.Convert(),
			Value:      ParseBaseExpression(n.Value),
		}
	case "NotEqualNode":
		n := element.(NotEqualNode)
		return parser.NotEqualExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "EqualNode":
		n := element.(EqualNode)
		return parser.EqualExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "GreaterThanOrEqualNode":
		n := element.(GreaterThanOrEqualNode)
		return parser.GreaterThanOrEqualExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "LessThanOrEqualNode":
		n := element.(LessThanOrEqualNode)
		return parser.LessThanOrEqualExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "GreaterThanNode":
		n := element.(GreaterThanNode)
		return parser.GreaterThanExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "LessThanNode":
		n := element.(LessThanNode)
		return parser.LessThanExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "AddNode":
		n := element.(AddNode)
		return parser.AddExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "SubNode":
		n := element.(SubNode)
		return parser.SubExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "MulNode":
		n := element.(MulNode)
		return parser.MulExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "DivNode":
		n := element.(DivNode)
		return parser.DivExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "ExpNode":
		n := element.(ExpNode)
		return parser.ExpExpression{
			// LeftRightExpression: parser.LeftRightExpression{
			Left:  ParseBaseExpression(n.Left),
			Right: ParseBaseExpression(n.Right),
			// },
		}
	case "FuncCallNode":
		n := element.(FuncCallNode)
		return parser.FuncCallExpression{
			Identifier: ParseBaseExpression(n.Target),
			Arguments:  ParseBaseExpressions(n.Args),
		}
	case "IntNode":
		n := element.(IntNode)
		t := n.Token.Convert()
		return parser.IntLiteral{
			Tok: &t,
		}
	case "VarNode":
		n := element.(VarNode)
		return parser.VarAccessExpression{
			Identifier: n.Token.Convert(),
		}
	default:
		panic(fmt.Sprintf("'%s' unexpected", element.(IElement).GetType()))
	}
}

func convertTokenType(tt string) parser.TokenType {
	switch tt {
	case "EOF":
		return parser.EOFToken
	case "IDENTIFIER":
		return parser.IdentifierToken
	case "KEYWORD":
		return parser.KeywordToken
	case "INT":
		return parser.IntToken
	case "LPAREN":
		return parser.LParenToken
	case "RPAREN":
		return parser.RParenToken
	case "LBRACE":
		return parser.LBraceToken
	case "RBRACE":
		return parser.RBraceToken
	case "ADD_OP":
		return parser.AddToken
	case "SUB_OP":
		return parser.SubToken
	case "MUL_OP":
		return parser.MulToken
	case "DIV_OP":
		return parser.DivToken
	case "EXP_OP":
		return parser.ExpToken
	case "ASGN_OP":
		return parser.AssignmentToken
	case "COLON":
		return parser.ColonToken
	case "COMMA":
		return parser.ParameterSeperatorToken
	default:
		panic(fmt.Sprintf("backend: '%s' not implemented", tt))
	}
}

func (t *Token) Convert() parser.Token {
	intValue, _ := strconv.Atoi(t.Value)
	return parser.Token{
		Type:        convertTokenType(t.TokenType),
		Next:        nil,
		Prev:        nil,
		IntValue:    intValue,
		StringValue: t.Value,
	}
}
