package main

import (
	"fmt"
	"os"

	"github.com/guthius/vb6/lexer"
)

func main() {
	bytes, _ := os.ReadFile("examples/modDatabase.bas")
	source := string(bytes)

	tokens := lexer.Tokenize(source)
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}
}
