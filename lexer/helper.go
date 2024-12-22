package lexer

import (
	"unicode"
)

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isLetter(c byte) bool {
	return unicode.IsLetter(rune(c))
}

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func isLetterOrDigitOrUnderscore(c byte) bool {
	return isLetter(c) || isDigit(c) || c == '_'
}
