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
	LParenToken
	RParenToken
	IdentifierToken
	KeywordToken
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
	case IntToken:
		return "int"
	case InvalidToken:
		return "invalid"
	default:
		return "invalid"
	}
}

func (t Token) String() string {
	switch t.Type {
	case IntToken:
		return fmt.Sprintf("%s{%d}", t.Type, t.Value)
	default:
		return fmt.Sprint(t.Type)
	}
}

type Token struct {
	Type  TokenType
	Value int
	Text  string
	Next  *Token
}

func tokenizeRune(r rune) *Token {
	switch r {
	case '+':
		return &Token{
			Type:  AddToken,
			Value: 0,
		}
	case '-':
		return &Token{
			Type:  SubToken,
			Value: 0,
		}
	case '*':
		return &Token{
			Type:  MulToken,
			Value: 0,
		}
	case '^':
		return &Token{
			Type:  ExpToken,
			Value: 0,
		}
	case '/':
		return &Token{
			Type:  DivToken,
			Value: 0,
		}
	case '(':
		return &Token{
			Type:  LParenToken,
			Value: 0,
		}
	case ')':
		return &Token{
			Type:  RParenToken,
			Value: 0,
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return &Token{
			Type:  IntToken,
			Value: int(r) - 48,
		}
	case ' ', '\n':
		return &Token{
			Type:  InvalidToken,
			Value: 0,
		}
	default:
		panic(fmt.Sprintf("invalid character %c", r))
	}
}

func TokenizeString(s string) *Token {
	runes := []rune(s)
	firstToken := tokenizeRune(runes[0])
	prevToken := firstToken

	for i := 1; i < len(runes); i++ {
		t := tokenizeRune(runes[i])
		if prevToken.Type == InvalidToken {
			firstToken = t
			prevToken = t
		}
		if t.Type != InvalidToken {
			prevToken.Next = t
			prevToken = t
		}
	}

	return firstToken
}
