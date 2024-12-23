package lexer

import (
	"fmt"
)

type lexer struct {
	Tokens []Token
	Source string
	Pos    int
}

func newLexer(source string) *lexer {
	return &lexer{
		Tokens: make([]Token, 0),
		Source: source,
		Pos:    0,
	}
}

func Tokenize(source string) []Token {
	l := newLexer(source)
	l.tokenize()
	return l.Tokens
}

func (l *lexer) add(kind Kind, value string) {
	// If the previous token is a line continuation token, we remove it.
	if len(l.Tokens) > 0 && l.Tokens[len(l.Tokens)-1].Kind == LineContinuation {
		l.Tokens = l.Tokens[:len(l.Tokens)-1]

		// If the current token is a line break token, we consume it and don't it to the list.
		if kind == LineBreak {
			return
		}
	}

	l.Tokens = append(l.Tokens, Token{
		Kind:  kind,
		Value: value,
	})
}

type tokenSpec struct {
	String string
	Kind   Kind
}

// tokenSpecs is a list of all the tokens that the lexer can recognize.
// Tokens must be ordered from the longest to the shortest. This is important
// because the lexer will try to match the longest token first.
//
// For example, if the lexer
// finds a '>' character, it will first check if it can match it with the '>='
// token before matching it with the '>' token.
var tokenSpecs = []tokenSpec{
	{String: "Option Explicit", Kind: OptionExplicit},
	{String: "End Select", Kind: EndSelect},
	{String: "DoEvents", Kind: DoEvents},
	{String: "End Enum", Kind: EndEnum},
	{String: "End Type", Kind: EndType},
	{String: "End With", Kind: EndWith},
	{String: "Function", Kind: Function},
	{String: "Boolean", Kind: BooleanType},
	{String: "Declare", Kind: Declare},
	{String: "Integer", Kind: IntegerType},
	{String: "Private", Kind: Private},
	{String: "Double", Kind: DoubleType},
	{String: "ElseIf", Kind: ElseIf},
	{String: "End If", Kind: EndIf},
	{String: "Public", Kind: Public},
	{String: "Select", Kind: Select},
	{String: "Single", Kind: SingleType},
	{String: "String", Kind: StringType},
	{String: "Alias", Kind: Alias},
	{String: "ByRef", Kind: ByRef},
	{String: "ByVal", Kind: ByVal},
	{String: "Const", Kind: Const},
	{String: "ReDim", Kind: ReDim},
	{String: "Until", Kind: Until},
	{String: "While", Kind: While},
	{String: "Byte", Kind: ByteType},
	{String: "Call", Kind: Call},
	{String: "Case", Kind: Case},
	{String: "Else", Kind: Else},
	{String: "Enum", Kind: Enum},
	{String: "GoTo", Kind: GoTo},
	{String: "Long", Kind: LongType},
	{String: "Loop", Kind: Loop},
	{String: "Step", Kind: Step},
	{String: "Then", Kind: Then},
	{String: "Type", Kind: Type},
	{String: "Wend", Kind: Wend},
	{String: "With", Kind: With},
	{String: "And", Kind: And},
	{String: "Dim", Kind: Dim},
	{String: "For", Kind: For},
	{String: "Lib", Kind: Lib},
	{String: "Mod", Kind: Modulus},
	{String: "Not", Kind: Not},
	{String: "Sub", Kind: Sub},
	{String: "Xor", Kind: Xor},
	{String: "As", Kind: As},
	{String: "Do", Kind: Do},
	{String: "If", Kind: If},
	{String: "Or", Kind: Or},
	{String: "To", Kind: To},
	{String: "<>", Kind: NotEqual},
	{String: ">=", Kind: GreaterThanOrEqual},
	{String: "<=", Kind: LessThanOrEqual},
	{String: "\n", Kind: LineBreak},
	{String: ".", Kind: Dot},
	{String: "_", Kind: LineContinuation},
	{String: ",", Kind: Comma},
	{String: "+", Kind: Add},
	{String: "-", Kind: Subtract},
	{String: "*", Kind: Multiply},
	{String: "/", Kind: Divide},
	{String: "\\", Kind: DivideInt},
	{String: "^", Kind: Exponent},
	{String: "&", Kind: Concat},
	{String: "(", Kind: LParen},
	{String: ")", Kind: RParen},
	{String: "=", Kind: Equal},
	{String: ">", Kind: GreaterThan},
	{String: "<", Kind: LessThan},
	{String: "#", Kind: FileNumber},
}

func (lex *lexer) tokenize() {
	for lex.Pos < len(lex.Source) {
		// Rem indicates a comment, but it can only be used at the beginning of a line.
		// So it must either be the first token in the source, or it must be preceded by a line break.
		if len(lex.Source)-lex.Pos >= 3 && lex.Source[lex.Pos:lex.Pos+3] == "Rem" && (len(lex.Tokens) == 0 || lex.Tokens[len(lex.Tokens)-1].Kind == LineBreak) {
			lex.skipComment(3)
			continue
		}

		c := lex.Source[lex.Pos]
		tokenFound := false
		for _, tokenSpec := range tokenSpecs {
			tokenLen := len(tokenSpec.String)
			if tokenLen > len(lex.Source)-lex.Pos {
				continue
			}
			if lex.Source[lex.Pos:lex.Pos+tokenLen] == tokenSpec.String {
				lex.add(tokenSpec.Kind, tokenSpec.String)
				lex.Pos += tokenLen
				tokenFound = true
				break
			}
		}

		if tokenFound {
			continue
		}

		if lex.Pos < len(lex.Source) && isWhitespace(lex.Source[lex.Pos]) {
			lex.Pos++
			continue
		}

		switch c {
		case '\'':
			lex.skipComment(1)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lex.tokenizeNumber()
		case '"':
			lex.tokenizeString()
		default:
			lex.tokenizeIdentifier()
		}
	}
	lex.add(EOF, "")
}

func (l *lexer) skipComment(skip int) {
	l.Pos += skip
	for l.Pos < len(l.Source) && l.Source[l.Pos] != '\n' && l.Source[l.Pos] != '\r' {
		l.Pos++
	}
}

func (l *lexer) tokenizeNumber() {
	start := l.Pos
	for l.Pos < len(l.Source) && isDigit(l.Source[l.Pos]) {
		l.Pos++
	}
	l.add(Number, l.Source[start:l.Pos])
}

func (l *lexer) tokenizeString() {
	l.Pos++
	start := l.Pos
	for l.Pos < len(l.Source) && l.Source[l.Pos] != '"' {
		l.Pos++
	}
	l.add(String, l.Source[start:l.Pos])
	l.Pos++
}

func (l *lexer) tokenizeIdentifier() {
	start := l.Pos
	if !isLetter(l.Source[l.Pos]) {
		panic(fmt.Sprintf("unexpected character '%c' in identifier near %s", l.Source[l.Pos], l.Source[l.Pos:l.Pos+10]))
	}
	for l.Pos < len(l.Source) && isLetterOrDigitOrUnderscore(l.Source[l.Pos]) {
		l.Pos++
	}
	l.add(Identifier, l.Source[start:l.Pos])
}
