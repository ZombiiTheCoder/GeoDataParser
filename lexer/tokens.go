package lexer

type Types string

const (
	Colon     Types = ":"
	BColon    Types = ":!"
	CColon    Types = ":^"
	SemiColon Types = ";"
	Equals    Types = "="
	Comma     Types = ","

	Int        Types = "Int"
	Float      Types = "Float"
	Boolean    Types = "Boolean"
	String     Types = "String"
	Identifier Types = "Identifier"
	Null       Types = "Null"
	EOF        Types = "EOF"
)

type Token struct {
	Value  string
	Type   Types
	Line   int
	Column int
}

func (s *Lexer) buildToken(Value string, Type Types) Token {
	return Token{Value, Type, s.line, s.column}
}