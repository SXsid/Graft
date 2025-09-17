package helper

type Node interface {
	TokenLiteral() string
}
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// let
type LetStatement struct {
	Name  string
	Value Expression
}

// if
type IfStatemetn struct {
	Condition Expression
	ThenBody  []Node
	ElseBody  []Node
}

// varible
type Variable struct {
	Value string
}

// For function calls: print("hello")
type CallExpression struct {
	Function  Expression   // What you're calling
	Arguments []Expression // What you pass to it
}

// For function definitions (if you need them later)
type FunctionStatement struct {
	Name       string
	Parameters []string
	Body       []Node
}
