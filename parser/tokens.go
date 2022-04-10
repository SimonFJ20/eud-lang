package parser

import "fmt"

type TokenType int

const (
	InvalidToken TokenType = iota
	AddToken
	SubToken
	MulToken
	DivToken
	ExpToken
	IntToken
	ColonToken
	AssignmentToken
	ParameterSeperatorToken
	LParenToken
	RParenToken
	LBraceToken
	RBraceToken
	RuneToken
	WordToken
	IdentifierToken
	KeywordToken
	EOFToken
)

func (t TokenType) String() string {
	switch t {
	case AddToken:
		return "add"
	case SubToken:
		return "sub"
	case MulToken:
		return "mul"
	case DivToken:
		return "div"
	case ExpToken:
		return "exp"
	case LParenToken:
		return "l_paren"
	case RParenToken:
		return "r_paren"
	case LBraceToken:
		return "l_brace"
	case RBraceToken:
		return "r_brace"
	case IntToken:
		return "int"
	case RuneToken:
		return "rune"
	case ParameterSeperatorToken:
		return "param_sep"
	case WordToken:
		return "word"
	case AssignmentToken:
		return "assign"
	case ColonToken:
		return "colon"
	case IdentifierToken:
		return "identifier"
	case KeywordToken:
		return "keyword"
	case EOFToken:
		return "eof"
	case InvalidToken:
		return "invalid"
	default:
		return "invalid"
	}
}

func (t Token) String() string {
	switch t.Type {
	case IntToken:
		return fmt.Sprintf("%s{%d}", t.Type, t.IntValue)
	case RuneToken:
		return fmt.Sprintf("%s{%d|'%c'}", t.Type, t.RuneValue, t.RuneValue)
	case WordToken:
		return fmt.Sprintf("%s{%s}", t.Type, t.StringValue)
	case IdentifierToken:
		return fmt.Sprintf("%s{%s}", t.Type, t.StringValue)
	case KeywordToken:
		return fmt.Sprintf("%s{%s}", t.Type, t.StringValue)
	default:
		return fmt.Sprint(t.Type)
	}
}

type Token struct {
	Type TokenType
	Next *Token
	Prev *Token

	// union type
	IntValue    int
	StringValue string
	RuneValue   rune
}
