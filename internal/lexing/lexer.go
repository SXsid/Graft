package lexing

import "github/SXsid/Graft/internal/helper"

type Kind string

const (
	Bracket    Kind = "bracket"
	Digit      Kind = "digit"
	String     Kind = "stirng"
	Identifier Kind = "ident"
	Error      Kind = "error"
	Assignment Kind = "assignment"
	Equate     Kind = "equate"
)

type Token struct {
	Kind Kind
	Val  string
	Pos  int
}

func Tokenizer(data string) []Token {
	var token []Token
	runes := []rune(data)
	// we need to control the flow of cursor
	for i := 0; i < len(runes); {
		r := runes[i]
		switch {
		case helper.IsBracket(r):
			token = append(token, Token{
				Kind: Bracket,
				Val:  string(r),
			})
			i++
		case helper.IsWhitespace(r):
			i++
		case helper.IsDigit(r):
			// accumate all sequnectial digits
			start := i
			for i < len(runes) && helper.IsDigit(runes[i]) {
				i++
			}
			token = append(token, Token{
				Kind: Digit,
				Val:  string(runes[start:i]),
			})
		case helper.IsQuote(r):
			i++ // skip the quteo
			start := i
			// we need to match the exact quopte
			quote := r
			for i < len(runes) && runes[i] != quote {
				i++
			}
			token = append(token, Token{
				Kind: String,
				Val:  string(runes[start:i]),
			})
			i++
		case helper.IsLetter(r):
			start := i
			for i < len(runes) && (helper.IsLetter(runes[i]) || helper.IsDigit(runes[i])) {
				i++
			}
			token = append(token, Token{
				Kind: Identifier,
				Val:  string(runes[start:i]),
			})
		case helper.IsOperator(r):

			if i+1 < len(runes) && helper.IsOperator(runes[i+1]) {
				token = append(token, Token{
					Kind: Equate,
					Val:  string(runes[i : i+2]),
				})
				i += 2
			} else {
				i++
				token = append(token, Token{
					Kind: Assignment,
					Val:  string(r),
				})
			}
		default:
			token = append(token, Token{
				Kind: Error,
				Val:  string(r),
				Pos:  i,
			})
			i++

		}

	}

	return token
}
