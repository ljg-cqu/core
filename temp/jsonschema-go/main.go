package main

import (
	"encoding/json"
	"fmt"
	jsonschema "github.com/swaggest/jsonschema-go"
	"log"
)

func main() {
	type MyStruct struct {
		Amount float64 `json:"amount" minimum:"10.5" example:"20.6" required:"true"`
		Abc    string  `json:"abc" pattern:"[abc]"`
	}

	reflector := jsonschema.Reflector{}

	schema, err := reflector.Reflect(MyStruct{})
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.MarshalIndent(schema, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(j))

	// Output:
	// {
	//  "required": [
	//   "amount"
	//  ],
	//  "properties": {
	//   "abc": {
	//    "pattern": "[abc]",
	//    "type": "string"
	//   },
	//   "amount": {
	//    "examples": [
	//     20.6
	//    ],
	//    "minimum": 10.5,
	//    "type": "number"
	//   }
	//  },
	//  "type": "object"
	// }
}
