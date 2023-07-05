package main

import (
	"fmt"
	"os"

	"github.com/ZombiiTheCoder/GeoDataParser/glib"
)

func main (){

	z, _ := os.ReadFile("./test/NodeMongoDB.json")
	m := glib.MapGeoData(glib.JsonToGeoData(string(z))) // returns Map of type map[string]any from json source file

	fmt.Println(m)

	// m.body accesses the contents of the geodata file
	// glib.Access changes type any to map[string]any to allow you to look through the map and see values
	// glib.Access takes a map[string]any for the tree and a string for the key 
}