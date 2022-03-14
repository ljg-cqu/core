package openchain

import (
	"fmt"
	"github.com/ljg-cqu/core/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_HandshakeIfNeeded(t *testing.T) {
	client := New(logger.New(), "../tests/configs.json")
	require.NotNil(t, client)
	err := client.HandshakeIfNeeded()
	require.Nil(t, err)
	fmt.Println(client.CurrentToken)
	fmt.Println(client.TokenExpireAt)
}
