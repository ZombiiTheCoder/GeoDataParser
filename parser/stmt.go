package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ZombiiTheCoder/GeoDataParser/lexer"
)

func (s *Parser) parseArray() Stmt {
	s.next()
	elements := make([]Stmt, 0)
	for s.at().Type != lexer.SemiColon {
		elements = append(elements, s.parseValue())
		if s.at().Type == lexer.Comma {
			s.next()
		}
	}
	if s.at().Type == lexer.SemiColon {
		s.expect(lexer.SemiColon, "Expected ; For End Of Object")
	}
	return Array{
		Type: "Array",
		Elements: elements,
	}
}

func (s *Parser) parseObject(a bool) Stmt {
	s.next()
	properties := make([]Property, 0)
	for s.at().Type != lexer.SemiColon {
		var key string
		sat := false
		var value Stmt
		switch s.tokenPeek(1).Type {
		case lexer.BColon:
			sat = true
			f := s.parseObject(true)
			properties = append(properties, f.(Object).Properties...)
		case lexer.Colon:
			if s.at().Type == lexer.String {
				key = s.expect(lexer.String, "Expected String For Object").Value
			}else {
				key = s.expect(lexer.Identifier, "Expected Identifier For Object").Value
			}
			value = s.parseObject(false)
		case lexer.CColon:
			f := s.parseArray()
			properties = append(properties, f.(Object).Properties...)
		case lexer.Equals:
			if s.at().Type == lexer.String {
				key = s.expect(lexer.String, "Expected String For Object").Value
			}else {
				key = s.expect(lexer.Identifier, "Expected Identifier For Object").Value
			}
			s.expect(lexer.Equals, "Expected = For Variable")
			value = s.parseValue()
		}
		
		if !sat{
			properties = append(properties, Property{Type: "Property", Key: key, Value: value})
		}
	}
	if s.at().Type == lexer.SemiColon {
		s.expect(lexer.SemiColon, "Expected ; For End Of Object")
		if a {
			var k string
			if s.at().Type == lexer.String {
				k = s.expect(lexer.String, "Key Expected After Semicolon").Value
			}else {
				k = s.expect(lexer.Identifier, "Key Expected After Semicolon").Value
			}
			return Object{
				Type: "Object",
				Properties: []Property{
					{
						Type: "Property",
						Key: k,
						Value: Object{
							Type: "Object",
							Properties: properties,
						},
					},
				},
			}
		}
	}
	return Object{
		Type: "Object",
		Properties: properties,
	}

}

func (s *Parser) parseValue() Stmt {

	switch s.at().Type {
		case lexer.EOF:
			return Null{Type: "Null", Value: nil}
		case lexer.Int:
			q, _ := strconv.ParseInt(s.next().Value, 10, 64)
			return Int{Type: "Int", Value: q}
		case lexer.Float:
			q, _ := strconv.ParseFloat(s.next().Value, 64)
			return Float{Type: "Float", Value: q}
		case lexer.Null:
			s.next()
			return Null{Type: "Null", Value: nil}
		case lexer.Boolean:
			q := false
			if s.next().Value == "true" {
				q = true
			}
			return Boolean{Type: "Boolean", Value: q}
		case lexer.Colon, lexer.CColon, lexer.BColon, lexer.Identifier, lexer.String:
			return s.parseStmt()
		default:
			fmt.Println("Invalid Token Found During Parsing", s.at(), "Line:", s.at().Line, "Pos:", s.at().Column)
			fmt.Println(s.tokens)
			os.Exit(1)
			return Null{Type: "Null", Value: nil}
		}
}

func (s *Parser) parseStmt() Stmt {
	switch s.at().Type {
	case lexer.EOF:
		return Null{Type: "Null", Value: nil}
	case lexer.Colon:
		return s.parseObject(false)
	case lexer.BColon:
		return s.parseObject(false)
	case lexer.CColon:
		return s.parseArray()
	case lexer.Identifier:
		switch s.tokenPeek(1).Type {
		case lexer.Equals:
			k := s.next().Value
			s.next()
			v := s.parseValue()
			return Object{
				Type: "Object",
				Properties: []Property{{Type:"Property", Key: k, Value: v}},
			}
		case lexer.Colon:
			k := s.next().Value
			v := s.parseObject(false)
			return Object{
				Type: "Object",
				Properties: []Property{{Type:"Property", Key: k, Value: v}},
			}
		default:
			return Identifier{Type: "Identifier", Value: s.next().Value}
		}
	case lexer.String:
		switch s.tokenPeek(1).Type {
		case lexer.Equals:
			k := s.next().Value
			s.next()
			v := s.parseValue()
			return Object{
				Type: "Object",
				Properties: []Property{{Type:"Property", Key: k, Value: v}},
			}
		case lexer.Colon:
			k := s.next().Value
			v := s.parseObject(false)
			return Object{
				Type: "Object",
				Properties: []Property{{Type:"Property", Key: k, Value: v}},
			}
		default:
			return String{Type: "String", Value: s.next().Value}
		}
	default:
		fmt.Println("Invalid Token Found During Parsing", s.at(), "Line:", s.at().Line, "Pos:", s.at().Column)
		fmt.Println(s.tokens)
		os.Exit(1)
		return Null{Type: "Null", Value: nil}
	}
}