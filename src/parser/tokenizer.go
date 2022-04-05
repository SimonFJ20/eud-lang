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
	RuneToken
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
		return fmt.Sprintf("%s{%d}", t.Type, t.IntValue)
	case RuneToken:
		return fmt.Sprintf("%s{%d}", t.Type, t.RuneValue)
	default:
		return fmt.Sprint(t.Type)
	}
}

type Token struct {
	Type TokenType
	Next *Token

	// union type
	IntValue    int
	StringValue string
	RuneValue   rune
}

func tokenizeRune(r rune) *Token {
	switch r {
	case '+':
		return &Token{
			Type: AddToken,
		}
	case '-':
		return &Token{
			Type: SubToken,
		}
	case '*':
		return &Token{
			Type: MulToken,
		}
	case '^':
		return &Token{
			Type: ExpToken,
		}
	case '/':
		return &Token{
			Type: DivToken,
		}
	case '(':
		return &Token{
			Type: LParenToken,
		}
	case ')':
		return &Token{
			Type: RParenToken,
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return &Token{
			Type:     IntToken,
			IntValue: int(r) - 48,
		}
	case ' ', '\n':
		return &Token{
			Type: InvalidToken,
		}
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':

		return &Token{
			Type:      RuneToken,
			RuneValue: r,
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
