package parser

import "fmt"

type TokenType = int

const (
	Invalid TokenType = iota
	Add
	Sub
	Mul
	Div
	LParen
	RParen
	Num
)

func (t Token) String() string {
	switch t.TokenType {
	case Add:
		return "add"
	case Sub:
		return "sub"
	case Mul:
		return "mul"
	case Div:
		return "div"
	case LParen:
		return "l_paren"
	case RParen:
		return "r_paren"
	case Num:
		return fmt.Sprintf("num{%d}", t.TokenValue)
	case Invalid:
		return "invalid"
	default:
		return "invalid"
	}
}

type Token struct {
	TokenType  TokenType
	TokenValue int
	Next       *Token
}

func tokenizeRune(r rune) *Token {
	switch r {
	case '+':
		return &Token{
			TokenType:  Add,
			TokenValue: 0,
		}
	case '-':
		return &Token{
			TokenType:  Sub,
			TokenValue: 0,
		}
	case '*':
		return &Token{
			TokenType:  Mul,
			TokenValue: 0,
		}
	case '/':
		return &Token{
			TokenType:  Div,
			TokenValue: 0,
		}
	case '(':
		return &Token{
			TokenType:  LParen,
			TokenValue: 0,
		}
	case ')':
		return &Token{
			TokenType:  RParen,
			TokenValue: 0,
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return &Token{
			TokenType:  Num,
			TokenValue: int(r) - 48,
		}
	case ' ':
		return &Token{
			TokenType:  Invalid,
			TokenValue: 0,
		}
	default:
		panic(fmt.Sprintf("invalid character %c", r))
	}
}

func TokenizeString(s string) *Token {
	runes := []rune(s)
	firstToken := TokenizeRune(runes[0])
	prevToken := firstToken

	for i := 1; i < len(runes); i++ {
		t := tokenizeRune(runes[i])
		if prevToken.TokenType == Invalid {
			firstToken = t
			prevToken = t
		}
		if t.TokenType != Invalid {
			prevToken.next = t
			prevToken = t
		}
	}

	return firstToken
}
