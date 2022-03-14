package openchain

import (
	"fmt"
	"github.com/ljg-cqu/core/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	client := New(logger.New(), "../tests/configs.json")
	require.NotNil(t, client)
	fmt.Println(client.CurrentToken)
}
