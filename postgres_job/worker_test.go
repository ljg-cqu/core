package postgres_job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ljg-cqu/core/_errors"
	"github.com/ljg-cqu/core/logger"
	"github.com/ljg-cqu/core/postgres"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/suite"
	gue "github.com/vgarvardt/gue/v3"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

type WorkerTestSuite struct {
	suite.Suite
	db     testserver.TestServer
	pool   *pgxpool.Pool
	logger *logger.Logger
}

func (s *WorkerTestSuite) SetupTest() {
	s.pool = postgres.PgxPool(postgres.TestDBAliConnStr)
	s.logger = logger.New()
}

func (s *WorkerTestSuite) TearDownTest() {
	defer s.pool.Close()
}

func TestWorkerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerTestSuite))
}

// Enqueue adds a job to the queue.
func (s *WorkerTestSuite) TestExample() {
	client := NewClient(s.pool, WithClientLogger(s.logger))

	const printerQueue QueueName = "name_printer"
	const jobTypePrinter JobType = "PrintName"

	// set work map
	type printNameArgs struct {
		Name string `json:"name"`
	}
	printName := func(ctx context.Context, j *gue.Job) error {
		var args printNameArgs
		err := json.Unmarshal(j.Args, &args)
		if err != nil {
			return _errors.NewWithMsgf("failed to unmarshal args:%v", err)
		}
		fmt.Printf("Hello %s!\n", args.Name)

		time.Sleep(1 * time.Second)
		return nil
	}
	wm := WorkMap{jobTypePrinter: printName}

	// set hooks
	finishedJobsLog := func(ctx context.Context, j *gue.Job, err error) {
		if err != nil {
			fmt.Printf("failed to execute a gue job. job_id:%v error:%v", j.ID, err)
			return
		}
		fmt.Printf("Finished a gue job. job_id:%v job_type:%v queue:%v args:%v finished_at:%v",
			j.ID, j.Type, j.Queue, string(j.Args), j.FinishedAt)
	}

	// create a pool w/ 2 workers
	workers := NewWorkerPool(client, wm, 2,
		WithPoolLogger(s.logger),
		WithPoolQueue(printerQueue),
		WithPoolHooksJobDone(finishedJobsLog),
		WithPoolPreserveCompletedJobs(false),
		WithPoolMigrateCompletedJobs(true))

	ctx, shutdown := context.WithCancel(context.Background())

	// 	work jobs in goroutine
	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := workers.Run(gctx)
		if err != nil {
			return _errors.NewWithMsg("gue workers error").Wrap(err)
		}
		return nil
	})

	// enqueue first job
	args, err := json.Marshal(printNameArgs{Name: "Zealy"})
	s.Require().Nil(err)

	gueJob := gue.Job{
		Priority: 5,
		Type:     string(jobTypePrinter),
		Args:     args,
		Queue:    string(printerQueue),
	}

	err = client.Enqueue(context.Background(), &gueJob)
	s.Require().Nil(err)

	sql := `SELECT job_id, priority, run_at, job_type, args, error_count, last_error, queue, created_at, updated_at FROM gue_jobs LIMIT 1`

	var jobSan GueJob
	err = pgxscan.Get(context.Background(), s.pool, &jobSan, sql)
	s.Require().Nil(err)
	utils.PrintlnAsJson("enqueued job:", &jobSan)

	// enqueue second job
	args, err = json.Marshal(printNameArgs{"Gorge"})
	s.Require().Nil(err)

	gueJob = gue.Job{
		Priority: 0,
		Type:     string(jobTypePrinter),
		Args:     args,
		Queue:    string(printerQueue),
		RunAt:    time.Now().UTC().Add(5 * time.Second), // delay 30 seconds
	}
	err = client.Enqueue(context.Background(), &gueJob)
	s.Require().Nil(err)

	err = pgxscan.Get(context.Background(), s.pool, &jobSan, sql)
	s.Require().Nil(err)
	utils.PrintlnAsJson("enqueued job:", &jobSan)

	// send shutdown signal to worker
	<-time.After(60 * time.Second) // wait for while

	shutdown()
	if err := g.Wait(); err != nil {
		s.logger.Fatalf("error encounterd in gue workers:%v", err)
	}
}
