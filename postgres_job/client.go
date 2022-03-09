package postgres_job

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ljg-cqu/core/logger"
	"github.com/ljg-cqu/core/postgres_job/adapter"
	"github.com/ljg-cqu/core/utils"
	"github.com/pkg/errors"
	gue "github.com/vgarvardt/gue/v3"
	gueAdapter "github.com/vgarvardt/gue/v3/adapter"
	"github.com/vgarvardt/gue/v3/adapter/exponential"
	"github.com/vgarvardt/gue/v3/adapter/pgxv4"
	"log"
	"os"
)

var c *Client

// Client is a Gue client that can add jobs to the queue and remove jobs from
// the queue.
type Client struct {
	ID          string
	Logger      *logger.Logger
	PgxPool     *pgxpool.Pool
	PoolAdapter gueAdapter.ConnPool
	backoff     Backoff
	*gue.Client
}

// NewClient creates a new Client that uses the pgx pool, logger/Logger, and default exponential backoff.
// Additionally, it automatically applies DB migration as defined with SchemaSql const.
func NewClient(pool *pgxpool.Pool, options ...ClientOption) *Client {
	if c != nil {
		c.Logger.WithField("client_id", c.ID).Debugln("Reuse my gue queue client that already existed.")
		return c
	}

	myCLient := Client{
		ID:      utils.MD5NowID(),
		Logger:  logger.New().SetLevelTrace(),
		PgxPool: pool,
		backoff: exponential.Default,
	}
	poolAdapter := pgxv4.NewConnPool(pool)
	myCLient.PoolAdapter = poolAdapter

	// Apply client options for my client
	for _, opt := range options {
		opt(&myCLient)
	}

	// set gue client options
	idOpt := gue.WithClientID(myCLient.ID)
	logOpt := gue.WithClientLogger(adapter.NewGueLogger(myCLient.Logger))
	bkOffOpt := gue.WithClientBackoff(gue.Backoff(myCLient.backoff))

	// Create gue client
	gueClient := gue.NewClient(poolAdapter, idOpt, logOpt, bkOffOpt)
	myCLient.Client = gueClient

	// apply DB migration before Client can work as expected
	_, err := pool.Exec(context.Background(), SchemaSql)
	if err != nil {
		log.Panicf("failed to to do DB migration for gue queue on top of PostgreSQL, error:%+v", err)
		os.Exit(1)
	}

	myCLient.Logger.Debugln("Things went well when applied DB migration for gue queue client.")

	c = &myCLient

	c.Logger.WithField("client_id", c.ID).Debugln("Initialize my new gue queue client.")
	return c
}

// Enqueue adds a job to the queue.
func (c *Client) Enqueue(ctx context.Context, j *gue.Job) error {
	return c.Client.Enqueue(ctx, j)
}

// EnqueueTx adds a job to the queue within the scope of the transaction.
// This allows you to guarantee that an enqueued job will either be committed or
// rolled back atomically with other changes in the course of this transaction.
//
// It is the caller's responsibility to Commit or Rollback the transaction after
// this function is called.
func (c *Client) EnqueueTx(ctx context.Context, j *gue.Job) (gueAdapter.Tx, error) {
	tx, err := c.PoolAdapter.Begin(ctx)
	if err != nil {
		return tx, errors.Wrap(err, "begin transaction failed")
	}

	return tx, c.Client.EnqueueTx(ctx, (*gue.Job)(j), tx)
}

// LockJob attempts to retrieve a Job from the database in the specified queue.
// If a job is found, it will be locked on the transactional level, so other workers
// will be skipping it. If no job is found, nil will be returned instead of an error.
//
// This function cares about the priority first to lock top priority jobs first even if there are available ones that
// should be executed earlier but with the lower priority.
//
// Because Gue uses transaction-level locks, we have to hold the
// same transaction throughout the process of getting a job, working it,
// deleting it, and releasing the lock.
//
// After the Job has been worked, you must call either Done() or Error() on it
// in order to commit transaction to persist Job changes (remove or update it).
func (c *Client) LockJob(ctx context.Context, queue QueueName) (*Job, error) {
	j, err := c.Client.LockJob(ctx, string(queue))
	return (*Job)(j), err
}

// LockJobByID attempts to retrieve a specific Job from the database.
// If the job is found, it will be locked on the transactional level, so other workers
// will be skipping it. If the job is not found, an error will be returned
//
// Because Gue uses transaction-level locks, we have to hold the
// same transaction throughout the process of getting the job, working it,
// deleting it, and releasing the lock.
//
// After the Job has been worked, you must call either Done() or Error() on it
// in order to commit transaction to persist Job changes (remove or update it).
func (c *Client) LockJobByID(ctx context.Context, id int64) (*Job, error) {
	j, err := c.Client.LockJobByID(ctx, id)
	return (*Job)(j), err
}

// LockNextScheduledJob attempts to retrieve the earliest scheduled Job from the database in the specified queue.
// If a job is found, it will be locked on the transactional level, so other workers
// will be skipping it. If no job is found, nil will be returned instead of an error.
//
// This function cares about the scheduled time first to lock earliest to execute jobs first even if there are ones
// with a higher priority scheduled to a later time but already eligible for execution
//
// Because Gue uses transaction-level locks, we have to hold the
// same transaction throughout the process of getting a job, working it,
// deleting it, and releasing the lock.
//
// After the Job has been worked, you must call either Done() or Error() on it
// in order to commit transaction to persist Job changes (remove or update it).
func (c *Client) LockNextScheduledJob(ctx context.Context, queue QueueName) (*Job, error) {
	j, err := c.Client.LockNextScheduledJob(ctx, string(queue))
	return (*Job)(j), err
}
