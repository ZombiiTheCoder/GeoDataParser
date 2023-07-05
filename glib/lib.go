package glib

import (
	"fmt"
	"os"

	GLexer "github.com/ZombiiTheCoder/GeoDataParser/lexer"
	GMapper "github.com/ZombiiTheCoder/GeoDataParser/mapper"
	GParser "github.com/ZombiiTheCoder/GeoDataParser/parser"

	JLexer "github.com/ZombiiTheCoder/JsonParser/lexer"
	JParser "github.com/ZombiiTheCoder/JsonParser/parser"
)

func MapGeoData(text string) any {

	lex := GLexer.Lexer{}
	lex.Init(text)
	Tokens := lex.Tokenize()
	
	par := GParser.Parser{}
	par.Init(Tokens)
	Ast := par.Parse()

	mapp := GMapper.Mapper{}
	return mapp.Eval(Ast)

}

func JsonToGeoData(text string) string {
	lex := JLexer.Lexer{}
	lex.Init(text)
	Tokens := lex.Tokenize()

	par := JParser.Parser{}
	par.Init(Tokens)
	Ast := par.Parse()
	f, _ := os.OpenFile("wosdaosdfjnasdfasdfasdf.geo_data", os.O_CREATE, 0644)
	geoc := geo{}

	fmt.Println(geoc.toGeo(0, Ast, f))
	f.Close()
	q, _ := os.ReadFile("wosdaosdfjnasdfasdfasdf.geo_data")
	os.Remove("wosdaosdfjnasdfasdfasdf.geo_data")
	return string(q)
}

func Access(j map[string]any, text string) map[string]any {

	if j[text] == nil {
		fmt.Printf("Value \"%v\" Does Not Exist In Map Or Value in map is not a nested map[string]any\n", text)
		return j
	}

	return j[text].(map[string]any)

}