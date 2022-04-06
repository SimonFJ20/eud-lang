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
	LParenToken
	RParenToken
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
	case IntToken:
		return "int"
	case RuneToken:
		return "rune"
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
		return fmt.Sprintf("%s{%d|%c}", t.Type, t.RuneValue, t.RuneValue)
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
	case '=':
		return &Token{
			Type: AssignmentToken,
		}
	case ':':
		return &Token{
			Type: ColonToken,
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return &Token{
			Type:     IntToken,
			IntValue: int(r) - 48,
		}
	case ' ', '\n', '\r', '\t':
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
		panic(fmt.Sprintf("invalid character {%d|%c}", r, r))
	}
}

func linkTokens(tokens []*Token) {
	firstToken := tokens[0]
	prevToken := firstToken

	for i := 1; i < len(tokens); i++ {
		t := tokens[i]
		if prevToken.Type == InvalidToken {
			firstToken = t
			prevToken = t
		}
		if t.Type != InvalidToken {
			prevToken.Next = t
			prevToken = t
		}
	}
}

func combineTokens(tokens []*Token) {
	previousToken := tokens[0]
	for i := 1; i < len(tokens); i++ {
		if previousToken.Type == RuneToken {
			previousToken.StringValue = string(previousToken.RuneValue)
			previousToken.Type = WordToken
		}
		if previousToken.Type == tokens[i].Type {
			if tokens[i].Type == IntToken {
				tokens[i].IntValue += previousToken.IntValue * 10
				previousToken.Type = InvalidToken
			}
		} else if tokens[i].Type == RuneToken && previousToken.Type == WordToken {
			tokens[i].StringValue = previousToken.StringValue + string(tokens[i].RuneValue)
			tokens[i].Type = WordToken
			previousToken.Type = InvalidToken
		} else if tokens[i].Type == IntToken && previousToken.Type == WordToken {
			tokens[i].StringValue = previousToken.StringValue + string(tokens[i].IntValue+48)
			tokens[i].Type = WordToken
			previousToken.Type = InvalidToken
		}

		previousToken = tokens[i]
	}
}

func wordTokenToIdentifierOrKeywordToken(token *Token) {
	keywords := listOfKeywords()
	for i := 0; i < len(keywords); i++ {
		if token.StringValue == keywords[i] {
			token.Type = KeywordToken
			return
		}
	}
	token.Type = IdentifierToken
}

func wordTokensToIdentifierOrKeywordTokens(tokens []*Token) {
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == WordToken {
			wordTokenToIdentifierOrKeywordToken(tokens[i])
		}
	}
}

func filterInvalidTokens(tokens *[]*Token) {
	n := 0
	for i := 0; i < len(*tokens); i++ {
		if (*tokens)[i].Type != InvalidToken {
			(*tokens)[n] = (*tokens)[i]
			n++
		}
	}
	*tokens = (*tokens)[:n]
}

func TokenizeString(s string) *Token {
	runes := []rune(s)
	tokens := make([]*Token, len(runes) + 1)
	for i := 0; i < len(runes); i++ {
		tokens[i] = tokenizeRune(runes[i])
	}
	tokens[len(runes)] = &Token {
		Type: EOFToken,
	}
	combineTokens(tokens)
	linkTokens(tokens)
	wordTokensToIdentifierOrKeywordTokens(tokens)
	filterInvalidTokens(&tokens)
	for _, str := range tokens {
		fmt.Printf("%s\n", str)
	}

	return tokens[0]
}
