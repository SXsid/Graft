package lexing

import "github/SXsid/Graft/internal/helper"

type Kind string

const (
	Bracket    Kind = "bracket"
	Digit      Kind = "digit"
	String     Kind = "stirng"
	Identifier Kind = "ident"
)

type Token struct {
	Kind Kind
	val  string
}

func Parser(data string) []Token {
	var token []Token
	runes := []rune(data)
	// we need to control the flow of cursor
	for i := 0; i < len(runes); {
		r := runes[i]
		switch {
		case helper.IsBracket(r):
			token = append(token, Token{
				Bracket,
				string(r),
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
				val:  string(runes[start:i]),
			})
		case helper.IsQuote(r):
			start := i
			i++ // skip the quteo
			// we need to match the exact quopte
			quote := r
			for i < len(runes) && runes[i] != quote {
				i++
			}
			// count the quote too
			i++
			token = append(token, Token{
				Kind: String,
				val:  string(runes[start:i]),
			})
		case helper.IsLetter(r):
			start := i
			for i < len(runes) && (helper.IsLetter(runes[i]) || helper.IsDigit(runes[i])) {
				i++
			}
			token = append(token, Token{
				Kind: Identifier,
				val:  string(runes[start:i]),
			})
		default:
			i++

		}

	}

	return token
}
