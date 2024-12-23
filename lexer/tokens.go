package lexer

import "fmt"

type Kind int

const (
	EOF Kind = iota
	LineBreak
	LineContinuation
	Identifier
	Number
	String
	LParen
	RParen
	Concat
	Comma

	// Operators
	Add
	Subtract
	Multiply
	Divide
	DivideInt
	Modulus
	Exponent
	Equal
	NotEqual
	GreaterThan
	GreaterThanOrEqual
	LessThan
	LessThanOrEqual
	And
	Or
	Not
	Xor
	// TODO: Eqv
	// TODO: Imp

	// Keywords
	Dim
	As
	Const
	Private
	Public
	ReDim
	OptionExplicit
	Type
	EndType
	Enum
	EndEnum
	Declare
	Sub
	Function
	ByVal
	ByRef
	Call
	Lib
	Alias
	DoEvents

	// Control Flow
	If
	Then
	Else
	ElseIf
	EndIf
	Select
	Case
	EndSelect
	For
	To
	Step
	While
	Wend
	Do
	Loop
	Until
	GoTo
	With
	EndWith

	// Data Types
	BooleanType
	ByteType
	IntegerType
	LongType
	SingleType
	DoubleType
	StringType

	// Special
	Dot
	FileNumber

	// TODO: Event + RaiseEvent keywords
	// TODO: AddressOf Lib Alias, Array LBound, UBound
	// TODO: Recordset, OpenDatabase, CloseDatabase, Field, Fields
	// TODO: File handling: Open, Close, Get, Put, Input, Print, Write, LineInput, EOF, FreeFile, Seek, FileAttr, FileCopy, Kill, Lock, Unlock
)

func TokenKindString(kind Kind) string {
	switch kind {
	case EOF:
		return "EOF"
	case LineBreak:
		return "LineBreak"
	case LineContinuation:
		return "LineContinuation"
	case Identifier:
		return "Identifier"
	case Number:
		return "Number"
	case String:
		return "String"
	case LParen:
		return "LParen"
	case RParen:
		return "RParen"
	case Concat:
		return "Concat"
	case Comma:
		return "Comma"
	case Add:
		return "Add"
	case Subtract:
		return "Subtract"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case DivideInt:
		return "DivideInt"
	case Modulus:
		return "Modulus"
	case Exponent:
		return "Exponent"
	case Equal:
		return "Equal"
	case NotEqual:
		return "NotEqual"
	case GreaterThan:
		return "GreaterThan"
	case GreaterThanOrEqual:
		return "GreaterThanOrEqual"
	case LessThan:
		return "LessThan"
	case LessThanOrEqual:
		return "LessThanOrEqual"
	case And:
		return "And"
	case Or:
		return "Or"
	case Not:
		return "Not"
	case Xor:
		return "Xor"
	case Dim:
		return "Dim"
	case As:
		return "As"
	case Const:
		return "Const"
	case Private:
		return "Private"
	case Public:
		return "Public"
	case ReDim:
		return "ReDim"
	case OptionExplicit:
		return "OptionExplicit"
	case Type:
		return "Type"
	case EndType:
		return "EndType"
	case Enum:
		return "Enum"
	case EndEnum:
		return "EndEnum"
	case Declare:
		return "Declare"
	case Sub:
		return "Sub"
	case Function:
		return "Function"
	case ByVal:
		return "ByVal"
	case ByRef:
		return "ByRef"
	case Call:
		return "Call"
	case Lib:
		return "Lib"
	case Alias:
		return "Alias"
	case DoEvents:
		return "DoEvents"
	case If:
		return "If"
	case Then:
		return "Then"
	case Else:
		return "Else"
	case ElseIf:
		return "ElseIf"
	case EndIf:
		return "EndIf"
	case Select:
		return "Select"
	case Case:
		return "Case"
	case EndSelect:
		return "EndSelect"
	case For:
		return "For"
	case To:
		return "To"
	case Step:
		return "Step"
	case While:
		return "While"
	case Wend:
		return "Wend"
	case Do:
		return "Do"
	case Loop:
		return "Loop"
	case Until:
		return "Until"
	case GoTo:
		return "GoTo"
	case With:
		return "With"
	case EndWith:
		return "EndWith"
	case BooleanType:
		return "Boolean"
	case ByteType:
		return "Byte"
	case IntegerType:
		return "Integer"
	case LongType:
		return "Long"
	case SingleType:
		return "Single"
	case DoubleType:
		return "Double"
	case StringType:
		return "String"
	case Dot:
		return "Dot"
	case FileNumber:
		return "FileNumber"
	default:
		return "Unknown"
	}
}

type Token struct {
	Value string
	Kind  Kind
}

func NewToken(kind Kind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

func (t Token) isOneOf(kinds ...Kind) bool {
	for _, kind := range kinds {
		if t.Kind == kind {
			return true
		}
	}
	return false
}

func (t Token) IsOperator() bool {
	return t.isOneOf(
		Add, Subtract,
		Multiply, Divide, DivideInt,
		Modulus, Exponent,
		Equal, NotEqual,
		GreaterThan,
		GreaterThanOrEqual,
		LessThan,
		LessThanOrEqual,
		And, Or, Not, Xor)
}

func (t Token) IsArithmeticOperator() bool {
	return t.isOneOf(Add, Subtract, Multiply, Divide, DivideInt, Modulus, Exponent)
}

func (t Token) IsComparisonOperator() bool {
	return t.isOneOf(Equal, NotEqual, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual)
}

func (t Token) IsLogicalOperator() bool {
	return t.isOneOf(And, Or, Not, Xor)
}

func (t Token) IsControlFlow() bool {
	return t.isOneOf(
		If, Then, Else, ElseIf, EndIf,
		Select, Case, EndSelect,
		For, To, Step, While, Wend,
		Do, Loop, Until, GoTo, With, EndWith)
}

func (t Token) IsKeyword() bool {
	return t.isOneOf(
		Dim, As, Const, Private, Public,
		ReDim, OptionExplicit, Type, EndType,
		Enum, EndEnum, Declare, Sub, Function,
		ByVal, ByRef, Call,
		If, Then, Else, ElseIf, EndIf,
		Select, Case, EndSelect,
		For, To, Step, While, Wend,
		Do, Loop, Until, GoTo, With, EndWith)
}

func (t Token) IsDataType() bool {
	return t.isOneOf(
		BooleanType,
		ByteType,
		IntegerType,
		LongType,
		SingleType,
		DoubleType,
		StringType)
}

func (t Token) String() string {
	if t.isOneOf(Identifier, String, Number) {
		return fmt.Sprintf("%s (%s)", TokenKindString(t.Kind), t.Value)
	} else {
		return TokenKindString(t.Kind)
	}
}
