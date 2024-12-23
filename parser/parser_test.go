package parser

import (
	"os"
	"testing"

	"github.com/guthius/vb6/lexer"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		filename string
	}{
		{"expr", "testcases/01_expr.bas"},
		{"declarations", "testcases/02_declarations.bas"},
		{"types", "testcases/03_types.bas"},
		{"call expr", "testcases/04_call_expr.bas"},
		{"call", "testcases/05_call.bas"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bytes, _ := os.ReadFile(c.filename)
			tokens := lexer.Tokenize(string(bytes))

			Parse(tokens)
		})
	}
}
