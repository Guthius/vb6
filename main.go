package main

import (
	"os"

	"github.com/guthius/vb6/lexer"
	"github.com/guthius/vb6/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("examples/declarations.bas")
	source := string(bytes)

	tokens := lexer.Tokenize(source)
	ast := parser.Parse(tokens)

	litter.Dump(ast)
}
