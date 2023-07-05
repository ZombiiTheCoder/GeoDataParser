package glib

import (
	"fmt"
	"os"

	"github.com/ZombiiTheCoder/JsonParser/parser"
)

type geo struct{}

func makeIndent(ident int) string {
	s := ""
	for i := 0; i < ident; i++ {
		s += "	"
	}
	return s
}

func (s *geo) toGeo(ident int, node parser.Stmt, f *os.File) any {

	switch node.GetType() {
	case "Program":
		body := s.toGeo(ident, node.(parser.Program).Body, f)
		f.WriteString("\"%LEXER_EOF%\"")
		return body

	case "Object":
		f.WriteString(":! ")
		for _, v := range node.(parser.Object).Properties {
			f.WriteString(fmt.Sprintf("%v\"%v\"", makeIndent(ident+1), v.Key))
			f.WriteString("=")
			s.toGeo(ident+1, v.Value, f)
			
		}
		f.WriteString(" ;")


		return nil

	case "String":
		f.WriteString(fmt.Sprintf("\"%v\"", node.(parser.String).Value))
		return node.(parser.String).Value
	
	case "Int":
		f.WriteString(fmt.Sprintf("%v", node.(parser.Int).Value))
		return node.(parser.Int).Value

	case "Float":
		f.WriteString(fmt.Sprintf("%v", node.(parser.Float).Value))
		return node.(parser.Float).Value

	case "Null":
		f.WriteString(fmt.Sprintf("%v", node.(parser.Null).Value))
		return node.(parser.Null).Value
	
	case "Boolean":
		f.WriteString(fmt.Sprintf("%v", node.(parser.Boolean).Value))
		return node.(parser.Boolean).Value
	
	case "Array":
		f.WriteString(":^ ")
		for _, v := range node.(parser.Array).Elements {
			s.toGeo(ident+1, v, f)
			f.WriteString(",")
		}
		f.WriteString(" ;")


		return nil

	default:
		fmt.Println("Unexpected Ast Block of type:", node.GetType())
		os.Exit(1)
	}

	return nil
}