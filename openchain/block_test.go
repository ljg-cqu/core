package openchain

import (
	"fmt"
	"github.com/ljg-cqu/core/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_QueryLastBlock(t *testing.T) {
	client := New(logger.NewForDebugStr(), "../tests/config.json")
	client.SetDebug(true)
	require.NotNil(t, client)
	account, err := client.QueryLastBlock()
	require.Nil(t, err)
	fmt.Println(account)
}
