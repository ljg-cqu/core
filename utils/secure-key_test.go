package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSecureKey(t *testing.T) {
	key, err := SecureKey(64)
	require.Nil(t, err)
	fmt.Println(key)
}
