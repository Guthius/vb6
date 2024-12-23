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
		{"declare", "testcases/06_declare.bas"},
		{"assignments", "testcases/07_assignments.bas"},
		{"functions", "testcases/08_functions.bas"},
		{"if", "testcases/09_if.bas"},
		{"if", "testcases/10_if.bas"},
		{"exit function", "testcases/11_exit_function.bas"},
		{"for", "testcases/12_for.bas"},
		{"sub", "testcases/13_sub.bas"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bytes, _ := os.ReadFile(c.filename)
			tokens := lexer.Tokenize(string(bytes))

			Parse(tokens)
		})
	}
}
