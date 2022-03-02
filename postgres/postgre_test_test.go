package postgres

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewPgxPoolForTest(t *testing.T) {
	ts, pl := NewPgxPoolForTest()
	defer ts.Stop()
	defer pl.Close()

	var greeting string
	err := pl.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	require.Nil(t, err)

	fmt.Println("connected to: ", ts.PGURL().String())

	after := time.After(time.Second * 10)

	<-after
}
