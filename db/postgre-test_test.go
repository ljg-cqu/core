package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPgxPoolForTest(t *testing.T) {
	ts, pl := NewPgxPoolForTest()
	defer ts.Stop()
	defer pl.Close()

	var greeting string
	err := pl.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	require.Nil(t, err)
	require.Equal(t, "Hello, world!", greeting)
}
