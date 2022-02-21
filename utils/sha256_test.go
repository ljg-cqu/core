package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum256(t *testing.T) {
	c := "hello world\n"
	fmt.Printf("%x", Sum256([]byte(c)))
}

func TestSum256FromFile(t *testing.T) {
	path := "./sha256.go"
	r, err := Sum256FromFile(path)
	require.Nil(t, err)
	fmt.Printf("%x", r)
}
