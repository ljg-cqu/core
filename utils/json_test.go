package utils

import (
	"fmt"
	"testing"
)

func TestPrintlnAsJson(t *testing.T) {
	type Private struct {
		Age      int
		Password string
	}

	type User struct {
		Name     string
		Tel      string
		PrivInfo Private
	}

	var me = User{"Zealy", "15215013094", Private{28, "1339382dsds"}}
	PrintlnAsJson("User me:", me)

	fmt.Println("-------compare Println------")
	fmt.Println(me)

	fmt.Println("-------compare Printf-------")
	fmt.Printf("%+v\n", me)

	type DepositContent struct {
		Content string `json:"content"`
	}

	var dc = &DepositContent{"This is a secret string"}
	PrintlnAsJson("", dc)
}
