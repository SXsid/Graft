package parser

import "github/SXsid/Graft/internal/lexing"

var keywords = map[string]bool{
	"if":    true,
	"else":  true,
	"let":   true,
	"for":   true,
	"in":    true,
	"print": true,
}

type AST struct {
	Type string `json:"type"`
	Body []any  `json:"body"`
}

func NewAST() *AST {
	return &AST{
		Type: "program",
		Body: make([]any, 0),
	}
}

func (ast *AST) Parse(token []lexing.Token) {
}
