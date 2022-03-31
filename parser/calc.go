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
	switch t.tokenType {
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
		return fmt.Sprintf("num{%d}", t.tokenValue)
	case Invalid:
		return "invalid"
	default:
		return "invalid"
	}
}

type Token struct {
	tokenType  TokenType
	tokenValue int
	next       *Token
}

func tokenizeRune(r rune) *Token {
	switch r {
	case '+':
		return &Token{
			tokenType:  Add,
			tokenValue: 0,
		}
	case '-':
		return &Token{
			tokenType:  Sub,
			tokenValue: 0,
		}
	case '*':
		return &Token{
			tokenType:  Mul,
			tokenValue: 0,
		}
	case '/':
		return &Token{
			tokenType:  Div,
			tokenValue: 0,
		}
	case '(':
		return &Token{
			tokenType:  LParen,
			tokenValue: 0,
		}
	case ')':
		return &Token{
			tokenType:  RParen,
			tokenValue: 0,
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return &Token{
			tokenType:  Num,
			tokenValue: int(r) - 48,
		}
	case ' ':
		return &Token{
			tokenType:  Invalid,
			tokenValue: 0,
		}
	default:
		panic(fmt.Sprintf("invalid character %c", r))
	}
}

func tokenizeString(s string) *Token {
	runes := []rune(s)
	firstToken := tokenizeRune(runes[0])
	prevToken := firstToken

	for i := 1; i < len(runes); i++ {
		t := tokenizeRune(runes[i])
		if prevToken.tokenType == Invalid {
			firstToken = t
			prevToken = t
		}
		if t.tokenType != Invalid {
			prevToken.next = t
			prevToken = t
		}
	}

	return firstToken
}

func parseToken(t *Token) int {
	return 0
}

func main() {
	token := tokenizeString("+")
	next := token
	for next != nil {
		fmt.Println(next)
		next = next.next
	}
	parseToken(token)
}
