package lexing

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	input := `print("hello") 42 {x} @`

	expected := []Token{
		{Kind: Identifier, Val: "print"},
		{Kind: Bracket, Val: "("},
		{Kind: String, Val: "hello"},
		{Kind: Bracket, Val: ")"},
		{Kind: Digit, Val: "42"},
		{Kind: Bracket, Val: "{"},
		{Kind: Identifier, Val: "x"},
		{Kind: Bracket, Val: "}"},
		{Kind: Error, Val: "@", Pos: 22},
	}

	tokens := Tokenizer(input)

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Tokens mismatch.\nGot:      %#v\nExpected: %#v", tokens, expected)
	}
}
