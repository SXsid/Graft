package main

import (
	"fmt"

	"github/SXsid/Graft/internal/lexing"
)

func main() {
	tokens := lexing.Parser(`print("hello") 42 {x}`)
	fmt.Println(tokens)
}
