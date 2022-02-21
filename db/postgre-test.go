package db

import (
	"context"
	"database/sql"
	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// More: https://github.com/cockroachdb/cockroach-go/tree/master/testserver

func NewPgDBForTest() (testserver.TestServer, *sql.DB) {
	ts, _ := testserver.NewTestServer()

	db, _ := sql.Open("postgres", ts.PGURL().String())
	return ts, db
}

func NewPgConnForTest() (testserver.TestServer, *pgx.Conn) {
	ts, _ := testserver.NewTestServer()

	conn, _ := pgx.Connect(context.Background(), ts.PGURL().String())
	return ts, conn
}

func NewPgxPoolForTest() (testserver.TestServer, *pgxpool.Pool) {
	ts, _ := testserver.NewTestServer()

	dbpool, _ := pgxpool.Connect(context.Background(), ts.PGURL().String())
	return ts, dbpool
}
