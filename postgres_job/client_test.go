package postgres_job

import (
	"context"
	"encoding/json"
	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ljg-cqu/core/postgres"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/suite"
	gue "github.com/vgarvardt/gue/v3"
	"testing"
	"time"
)

type ClientTestSuite struct {
	suite.Suite
	pool    *pgxpool.Pool
	client  *Client
	queue   QueueName
	jobType JobType
}

func (s *ClientTestSuite) SetupTest() {
	s.pool = postgres.PgxPool(postgres.TestDBAliConnStr)
	s.client = NewClient(s.pool)
	s.queue = "name_printer"
	s.jobType = "PrintName"
}

func (s *ClientTestSuite) TearDownTest() {
	defer s.pool.Close()
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

// Enqueue adds a job to the queue.
func (s *ClientTestSuite) TestEnqueue() {
	// Args must be in valid JSON format
	type printNameArgs struct {
		Name string `json:"name"`
	}

	// Enqueue gue jobs
	for i := 0; i < 10; i++ {
		args, _ := json.Marshal(&printNameArgs{Name: gofakeit.Name()}) // Note: if you marshal by msgpack, it will fail.
		job := &gue.Job{
			Queue:    string(s.queue),
			Priority: int16(i),
			RunAt:    time.Now().Add(time.Duration(i*10) * time.Second),
			Type:     string(s.jobType),
			Args:     args,
		}

		err := s.client.Enqueue(context.Background(), job)
		s.Require().Nil(err)
		time.Sleep(1 * time.Second)
	}

	sql := `SELECT job_id, priority, job_type, queue, args, created_at, run_at, finished_at, error_count, last_error, updated_at FROM gue_jobs`

	// Query gue jobs
	var gueJobs []GueJob
	err := pgxscan.Select(context.Background(), s.pool, &gueJobs, sql)
	s.Require().Nil(err)

	for _, gueJob := range gueJobs {
		utils.PrintlnAsJson("enqueued job:", gueJob)
	}
}

// EnqueueTx adds a job to the queue within the scope of the transaction.
// This allows you to guarantee that an enqueued job will either be committed or
// rolled back atomically with other changes in the course of this transaction.
//
// It is the caller's responsibility to Commit or Rollback the transaction after
// this function is called.
func (s *ClientTestSuite) TestEnqueueTx() {
	// Args must be in valid JSON format
	type printNameArgs struct {
		Name string `json:"name"`
	}

	// Enqueue gue jobs
	for i := 0; i < 10; i++ {
		args, _ := json.Marshal(&printNameArgs{Name: gofakeit.Name()}) // Note: if you marshal by msgpack, it will fail.
		job := &gue.Job{
			Queue:    string(s.queue),
			Priority: int16(i),
			RunAt:    time.Now().Add(time.Duration(i*10) * time.Second),
			Type:     string(s.jobType),
			Args:     args,
		}

		ctx := context.Background()

		tx, err := s.client.EnqueueTx(ctx, job)
		s.Nil(err)
		if err != nil {
			s.Require().Nil(tx.Rollback(ctx))
		}
		s.Require().Nil(tx.Commit(ctx))

		time.Sleep(1 * time.Second)
	}

	sql := `SELECT job_id, priority, job_type, queue, args, created_at, run_at, finished_at, error_count, last_error, updated_at FROM gue_jobs`

	// Query gue jobs
	var gueJobs []GueJob
	err := pgxscan.Select(context.Background(), s.pool, &gueJobs, sql)
	s.Require().Nil(err)

	for _, gueJob := range gueJobs {
		utils.PrintlnAsJson("enqueued job in transaction:", gueJob)
	}
}
