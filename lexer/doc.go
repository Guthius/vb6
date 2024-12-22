// Package lexer is designed to process VB6 source code by breaking it into tokens for further processing.
//
// # Example Usage
//
// Given the following VB6 source code:
//
//	Dim x As Integer
//	x = 10 + 20
//	If x > 15 Then
//	   	Print "x is large"
//	End If
//
// The lexer can be used to tokenize the source code as follows:
//
//	source := "Dim x As Integer\nx = 10 + 20\nIf x > 15 Then\n    Print \"x is large\"\nEnd If"
//	tokens := lexer.Tokenize(source)
//	for _, token := range tokens {
//	    fmt.Println(token)
//	}
//
// The output will be:
//
//	Dim
//	Identifier (x)
//	As
//	Integer
//	LineBreak
//	Identifier (x)
//	Equal
//	Number (10)
//	Add
//	Number (20)
//	LineBreak
//	If
//	Identifier (x)
//	GreaterThan
//	Number (15)
//	Then
//	LineBreak
//	Identifier (Print)
//	String ("x is large")
//	LineBreak
//	EndIf
//	EOF
//
// # Future Enhancements
//
//   - Add support for Eqv and Imp keywords.
//   - Improve error handling for invalid characters.
package lexer
