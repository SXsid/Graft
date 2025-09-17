package helper

func IsLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func IsBracket(r rune) bool {
	return r == '{' || r == '}' || r == '[' || r == ']' || r == '(' || r == ')'
}

func IsQuote(r rune) bool {
	return r == '"' || r == '\''
}

func IsOperator(r rune) bool {
	return r == '='
}
