# Json Parser

A json parser/mapper written in go to help you read and use json data whenever you want

## Use

```go

package main

import (
	"fmt"
	"github.com/ZombiiTheCoder/GeoDataParser/glib"
)

func main (){
	m := glib.MapGeoData(`
	{
		"dependencies": {
		"ejs": "^3.1.6",
		"express": "^4.17.3",
		"mongoose": "^6.2.10",
		"morgan": "^1.10.0",
		"nodemon": "^2.0.15"
		}
	}
	`) // returns Map of type map[string]any from json source file

	// m.body accesses the contents of the json file
	// JsonParser.Access changes type any to map[string]any to allow you to look through the map and see values
	// JsonParser.Access takes a map[string]any for the tree and a string for the key 
	body := jlib.Access(m, "body")

	fmt.Println(jlib.Access(body, "d"))
}


```

### Valid Ways of representation in .geo_data

##### Alone Variables
```
	p = 1
```
##### Variable Objects
```
p = : 
	a = "b"
;
```
##### Objects With Name
```
p:
	a = "b"
;
```

##### Arrays
```
:^
	1, 2, 3, 4, 5, 6, 7
;
```
##### Variable With Named Object
```
p = p:
	a = "b"
;
```
##### Objects With SemiColon Before Name
```
:
	a = "b"
; e
```

##### Unnamed Objects
```
:!
	a = "b"
;
```

### Supported Types

String 
   `""`

Bool
    `true | false`

Null
   `null`

Int
   `0, 1, 2...`

Float
  `0.0, 0.1, 0.2...`