package openchain

import (
	"fmt"
	"github.com/ljg-cqu/core/logger"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_QueryAccount(t *testing.T) {
	client := New(logger.NewForDebugStr(), "../tests/config.json")
	client.SetDebug(true)
	require.NotNil(t, client)
	account, err := client.QueryAccount("ff89d3d4569342028b4767e779e55994") // TODO: account type re def
	require.Nil(t, err)
	fmt.Println(account)
}

func TestClient_CreateAccount(t *testing.T) {
	client := New(logger.NewForDebugStr(), "../tests/config.json")
	client.SetDebug(true)
	require.NotNil(t, client)

	id := "Zealy" + utils.NowTimestamp13()

	account, err := client.CreateAccount(id) // TODO: account type re def
	require.Nil(t, err)
	fmt.Println(account)
}
