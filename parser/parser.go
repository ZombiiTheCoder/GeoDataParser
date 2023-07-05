package parser

import (
	"fmt"
	"os"

	"github.com/ZombiiTheCoder/GeoDataParser/lexer"
)

type Parser struct {
	pointer int
	tokens  []lexer.Token
}

func (s *Parser) at() lexer.Token {
	if (s.pointer) >= len(s.tokens) {
		return lexer.Token{Value: "Unexpected End Of File", Type: lexer.EOF}
	}
	return s.tokens[s.pointer]
}

func (s *Parser) next() lexer.Token {
	if !(s.pointer+1 >= len(s.tokens)) {
		s.pointer++
	}
	return s.tokens[s.pointer-1]
}

func (s *Parser) tokenPeek(i int) lexer.Token {
	if s.pointer+i >= len(s.tokens) {
		return s.tokens[len(s.tokens)-1]
	}
	return s.tokens[s.pointer+i]
}
func (s *Parser) expect(Type lexer.Types, err string) lexer.Token {
	if !(s.pointer+1 >= len(s.tokens)) {
		s.pointer++
	}
	last := s.tokens[s.pointer-1]
	if last.Type != Type {
		fmt.Println(err, "\nUnexpected Token", last, "Expecting Token Type",Type, "Line:", last.Line, "Pos:", last.Column)
		os.Exit(1)
	}
	return last
}

func (s *Parser) Init(tokens []lexer.Token) {
	s.tokens = append(s.tokens, tokens...)
	s.pointer = 0
}

func (s *Parser) Parse() Stmt {

	return Program{Type: "Program", Body:s.parseStmt()}
}