package lexing

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	input := `print("hello") 42 {x}`

	expected := []Token{
		{Kind: Identifier, val: "print"},
		{Kind: Bracket, val: "("},
		{Kind: String, val: `"hello"`},
		{Kind: Bracket, val: ")"},
		{Kind: Digit, val: "42"},
		{Kind: Bracket, val: "{"},
		{Kind: Identifier, val: "x"},
		{Kind: Bracket, val: "}"},
	}

	tokens := Parser(input)

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokens mismatch.\nGot:      %#v\nExpected: %#v", tokens, expected)
	}
}
