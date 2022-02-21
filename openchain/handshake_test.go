package openchain

import (
	"fmt"
	"github.com/ljg-cqu/core/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_HandshakeIfNeeded(t *testing.T) {
	client := New(logger.NewForDebugStr(), "../tests/config.json")
	require.NotNil(t, client)
	err := client.HandshakeIfNeeded()
	require.Nil(t, err)
	fmt.Println(client.CurrentToken)
	fmt.Println(client.TokenExpireAt)
}
