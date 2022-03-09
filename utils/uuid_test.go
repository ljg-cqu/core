package utils

import (
	"fmt"
	"testing"
)

func TestUuid(t *testing.T) {
	u := Uuid()
	fmt.Printf("len:%v, value:%v", len(u), u)

}
