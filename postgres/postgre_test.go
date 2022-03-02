package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
	"time"
)

// More: https://github.com/cockroachdb/cockroach-go/tree/master/testserver

func NewPgDBForTest() (testserver.TestServer, *sql.DB) {
	ts, _ := testserver.NewTestServer()
	connStr := ts.PGURL().String()
	db, _ := sql.Open("postgres", connStr)
	TestDBCockroachConnStr = connStr
	return ts, db
}

func NewPgConnForTest() (testserver.TestServer, *pgx.Conn) {
	ts, _ := testserver.NewTestServer()
	connStr := ts.PGURL().String()
	conn, _ := pgx.Connect(context.Background(), connStr)
	TestDBCockroachConnStr = connStr
	return ts, conn
}

func NewPgxPoolForTest() (testserver.TestServer, *pgxpool.Pool) {
	ts, _ := testserver.NewTestServer()
	connStr := ts.PGURL().String()
	dbPool, _ := pgxpool.Connect(context.Background(), connStr)
	TestDBCockroachConnStr = connStr
	return ts, dbPool
}

func TestPgxPool(t *testing.T) {
	pool := PgxPool(TestDBAliConnStr)

	// GueJob corresponds to table gue_jobs as define above
	type GueJob struct {
		JobID   int64  `db:"job_id"`
		JobType string `db:"job_type"`
		Queue   string `db:"queue"`
		Args    []byte `db:"args"`

		Priority int16     `db:"priority"` // use it for OrderByPriority poll strategy
		RunAt    time.Time `db:"run_at"`   // use it for OrderByRunAtPriority poll strategy

		ErrCount int32          `db:"error_count"`
		LastErr  sql.NullString `db:"last_error"`

		CreatedAt  time.Time    `db:"created_at"`
		UpdatedAt  sql.NullTime `db:"updated_at"`  // updated when error occur during job execution
		FinishedAt sql.NullTime `db:"finished_at"` // finished when job execution success
	}

	var j GueJob

	queue := "name_printer"
	now := time.Now().UTC()

	err := pool.QueryRow(context.Background(), `SELECT job_id, queue, priority, run_at, job_type, args, error_count, last_error, created_at, updated_at
FROM gue_jobs
WHERE queue = $1 AND run_at <= $2 AND finished_at IS NULL
ORDER BY priority ASC
LIMIT 1 FOR UPDATE SKIP LOCKED`, queue, now).Scan(
		&j.JobID,
		&j.Queue,
		&j.Priority,
		&j.RunAt,
		&j.JobType,
		(*json.RawMessage)(&j.Args),
		&j.ErrCount,
		&j.LastErr,
		&j.CreatedAt,
		&j.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error:", err)
	}
}
