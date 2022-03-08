package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

//The *pgx.Conn returned by pgx.Connect() represents a single connection
//and is not concurrency safe. This is entirely appropriate for a simple
//command line example such as above. However, for many uses, such as a
//web application server, concurrency is required.
//	To use a connection pool replace the import github.com/jackc/pgx/v4 with
//github.com/jackc/pgx/v4/pgxpool and connect with pgxpool.Connect() instead of pgx.Connect().

func PgxPool(connstr string) *pgxpool.Pool {
	pgxCfg, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		fmt.Printf("Failed to parse Postgres connect string %q, error:%v", connstr, err)
		os.Exit(1)
	}
	pgxPool, err := pgxpool.ConnectConfig(context.Background(), pgxCfg)
	if err != nil {
		fmt.Printf("Failed to establish pgxpool, connect string %q, error:%v", connstr, err)
		os.Exit(1)
	}

	return pgxPool
}
