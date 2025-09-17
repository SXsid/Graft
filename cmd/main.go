package main

import (
	"fmt"

	"github/SXsid/Graft/internal/lexing"
)

func main() {
	value := `if condition {
    name="aman"
    print("true branch")
} else {
    if name == "sid"{
        print("hello world")
    }
    print("false branch")
}
`
	tokens := lexing.Tokenizer(value)
	for _, token := range tokens {
		fmt.Println(token)
	}
}
