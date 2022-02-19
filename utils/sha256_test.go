package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum256FromString(t *testing.T) {
	c := "hello world\n"
	fmt.Printf("%x", Sum256FromString(c))
}

func TestSum256FromBytes(t *testing.T) {
	c := "hello world\n"
	fmt.Printf("%x", Sum256FromBytes([]byte(c)))
}

func TestSum256FromFile(t *testing.T) {
	path := "./sha256.go"
	r, err := Sum256FromFile(path)
	require.Nil(t, err)
	fmt.Printf("%x", r)
}
