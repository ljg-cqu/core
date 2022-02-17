package utils

import (
	"encoding/json"
	"fmt"
)

func PrintlnAsJson(pre string, v any) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(pre)
	fmt.Println(string(bytes))
}
