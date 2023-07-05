package lexer

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Lexer struct {
	pointer int
	chars   []string

	line   int
	column int
	tokens []Token

	eof bool
}

func (s *Lexer) Init(Text string) {
	s.chars = strings.Split(Text, "")
	s.pointer = 0
	s.line = 1
	s.column = 1
	s.tokens = make([]Token, 0)
	s.eof = false
}

func (s *Lexer) at() string {
	if s.pointer >= len(s.chars) {
		s.eof = true
		return " "
	}
	return s.chars[s.pointer]
}

func (s *Lexer) peek(i int) string {
	if s.pointer+i >= len(s.chars) {
		s.eof = true
		return " "
	}
	return s.chars[s.pointer+i]
}

func (s *Lexer) next() {
	if s.pointer+1 >= len(s.chars) {
		s.eof = true
	}
	s.pointer++
}

func (s *Lexer) skippable() {
	switch s.at() {
	case "\n":
		s.line++
		s.column = 0
		s.next()
	case "\r", "\b", "\f", "\t", "\v", "\n\r", " ", "":
		s.next()
	}
}

func (s *Lexer) oneCharTokens() {
	switch s.at() {
	case ":":
		if s.peek(1) == "!" {
			s.tokens = append(s.tokens, s.buildToken(s.at(), BColon))
			s.next()
			s.next()
		}else if s.peek(1) == "^" {
			s.tokens = append(s.tokens, s.buildToken(s.at(), CColon))
			s.next()
			s.next()
		}else {
			s.tokens = append(s.tokens, s.buildToken(s.at(), Colon))
			s.next()
		}
	case ",":
		s.tokens = append(s.tokens, s.buildToken(s.at(), Comma))
		s.next()
	case ";":
		s.tokens = append(s.tokens, s.buildToken(s.at(), SemiColon))
		s.next()
	case "=":
		s.tokens = append(s.tokens, s.buildToken(s.at(), Equals))
		s.next()
	}
}

func (s *Lexer) strings() {
	Value := ""
	if s.at() == "\"" {
		s.next()
		for s.at() != "\"" {
			Value = Value + s.at()
			s.next()
		}
		s.next()
		s.tokens = append(s.tokens, s.buildToken(Value, String))
	}
}

func regx(test, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(test)
}

func (s *Lexer) AlnumTokens() {
	Value := ""
	if regx(s.at(), "[[:alpha:]_]") {
		for regx(s.at(), "[[:alnum:]_]") {
			Value = Value + s.at()
			s.next()
		}
		switch Value {
		case "null":
			s.tokens = append(s.tokens, s.buildToken(Value, Null))
		case "true":
			s.tokens = append(s.tokens, s.buildToken(Value, Boolean))
		case "false":
			s.tokens = append(s.tokens, s.buildToken(Value, Boolean))

		default:
			s.tokens = append(s.tokens, s.buildToken(Value, Identifier))
		}
	}
}

func (s *Lexer) Comments() {
	if s.at()+s.peek(1) == "??" {
		s.next()
		s.next()

		for s.at() != "\n" && !s.eof {
			s.next()
			fmt.Println(s.eof)
		}
	}
}

func (s *Lexer) NumTokens() {
	Value := ""
	if regx(s.at(), "^[0-9]*$") {
		for regx(s.at(), "^[0-9]*$") || s.at() == "." {
			Value = Value + s.at()
			s.next()
		}
		if strings.Count(Value, ".") == 0 {
			s.tokens = append(s.tokens, s.buildToken(Value, Int))
		}else if strings.Count(Value, ".") == 1 && regx(Value, "^([0-9.]|[.])+[0-9]*$"){
			s.tokens = append(s.tokens, s.buildToken(Value, Float))	
		}else {
			fmt.Println("Invalid Float Or Int Value:", Value, "Line:", s.line, "Pos:", s.column)
			os.Exit(1)
		}
	}
}

func (s *Lexer) Tokenize() []Token {
	
	for i := 0; i < len(s.chars); i++ {

		if !s.eof {
			s.column++
			s.skippable()
			s.oneCharTokens()
			s.strings()
			if len(s.tokens) > 0 {
				if s.tokens[len(s.tokens)-1].Value == "%LEXER_EOF%" {
					break;
				}
			}
			s.AlnumTokens()
			s.NumTokens()
			s.Comments()
			
			switch s.at() {
			case "[", "]", "{", "}", "\"", ":", "!", ";", "^", ",", "=", "\n", "\r", "\b", "\f", "\t", "\v", " ":
			default:
				if !regx(s.at(), "[[:alnum:]/@]"){
					fmt.Println("Invalid Character:", s.at(), "Line:", s.line, "Pos:", s.column)
					os.Exit(1)
				}
				
			}
		}
	}
	s.tokens = append(s.tokens, s.buildToken("End Of File", EOF))
	return s.tokens
}