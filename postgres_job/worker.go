package postgres_job

import (
	"context"
	"github.com/ljg-cqu/core/logger"
	"github.com/ljg-cqu/core/postgres_job/adapter"
	"github.com/ljg-cqu/core/utils"
	gue "github.com/vgarvardt/gue/v3"
	"time"
)

// PollStrategy determines how the DB is queried for the next job to work on
type PollStrategy string

var (
	DefaultPoolSize     = 2
	DefaultPollInterval = 5 * time.Second
)

const (
	DefaultQueueName QueueName = "default_queue"

	// PriorityPollStrategy cares about the priority first to lock top priority jobs first even if there are available
	//ones that should be executed earlier but with lower priority.
	PriorityPollStrategy PollStrategy = "OrderByPriority"
	// RunAtPollStrategy cares about the scheduled time first to lock earliest to execute jobs first even if there
	// are ones with a higher priority scheduled to a later time but already eligible for execution
	RunAtPollStrategy PollStrategy = "OrderByRunAtPriority"
)

// By default, jobs are deleted from the gue_jobs table after being performed. If you would prefer to leave them in the
// table for analytics or debugging purposes, you can pass the PreserveCompletedJobs option to NewWorkerPool.
var PreserveCompletedJobs = false

// By default, jobs are migrate to the gue_jobs_finished table after being performed.
// You can pass the MigratedCompletedJobs option to gue.NewWorkerPool to have your own control.
var MigrateCompletedJobs = true

type QueueName string
type JobType string

// WorkFunc is a function that performs a Job. If an error is returned, the job
// is re-enqueued with exponential backoff.
type WorkFunc func(ctx context.Context, j *gue.Job) error

// HookFunc is a function that may react to a Job lifecycle events. All the callbacks are being executed synchronously,
// so be careful with the long-running locking operations. Hooks do not return an error, therefore they can not and
// must not be used to affect the Job execution flow, e.g. cancel it - this is the WorkFunc responsibility.
// Modifying Job fields and calling any methods that are modifying its state within hooks may lead to undefined
// behaviour. Please never do this.
//
// Depending on the event err parameter may be empty or not - check the event description for its meaning.
type HookFunc func(ctx context.Context, j *gue.Job, err error)

// WorkMap is a map of Job names to WorkFuncs that are used to perform Jobs of a
// given type.
type WorkMap map[JobType]WorkFunc

type WorkerPool struct {
	ID                    string
	Size                  int
	Logger                *logger.Logger
	Client                *Client
	WorkMap               WorkMap
	Queue                 QueueName
	pollInterval          time.Duration
	pollStrategy          PollStrategy
	preserveCompletedJobs bool
	migrateCompletedJobs  bool

	hooksJobLocked      []HookFunc
	hooksUnknownJobType []HookFunc
	hooksJobDone        []HookFunc

	*gue.WorkerPool
}

// NewWorkerPool creates a new WorkerPool with count workers using the Client c,
// configured with default poll interval of 5 seconds and default poll strategy of priority poll strategy.
//
// PriorityPollStrategy cares about the priority first to lock top priority jobs first even if there are available
//ones that should be executed earlier but with lower priority.
func NewWorkerPool(c *Client, wm WorkMap, size int, options ...WorkerPoolOption) *WorkerPool {
	var poolSize = size
	if size <= 0 {
		poolSize = DefaultPoolSize
	}
	myPool := &WorkerPool{
		ID:                    utils.MD5NowID(),
		Size:                  poolSize,
		Logger:                logger.New().SetLevelTrace(),
		Client:                c,
		WorkMap:               wm,
		Queue:                 DefaultQueueName,
		pollInterval:          DefaultPollInterval,
		pollStrategy:          PriorityPollStrategy,
		preserveCompletedJobs: PreserveCompletedJobs,
		migrateCompletedJobs:  MigrateCompletedJobs,
	}

	// Apply my worker pool options update
	for _, opt := range options {
		opt(myPool)
	}

	gueWm := make(gue.WorkMap)
	for jobType, workFun := range myPool.WorkMap {
		gueWm[string(jobType)] = gue.WorkFunc(workFun)
	}

	idOpt := gue.WithPoolID(myPool.ID)
	logOpt := gue.WithPoolLogger(adapter.NewGueLogger(myPool.Logger))
	presvOpt := gue.WithPoolPreserveCompletedJobs(myPool.preserveCompletedJobs)
	migrOpt := gue.WithPoolMigrateCompletedJobs(myPool.migrateCompletedJobs)
	queOpt := gue.WithPoolQueue(string(myPool.Queue))
	pollInt := gue.WithPoolPollInterval(myPool.pollInterval)
	pollStr := gue.WithPoolPollStrategy(gue.PollStrategy(myPool.pollStrategy))

	lockedHooks := make([]gue.HookFunc, 0)
	for _, hk := range myPool.hooksJobLocked {
		lockedHooks = append(lockedHooks, gue.HookFunc(hk))
	}
	lockedHooks_ := gue.WithPoolHooksJobLocked(lockedHooks...)

	unknownTypeHooks := make([]gue.HookFunc, 0)
	for _, hk := range myPool.hooksJobLocked {
		unknownTypeHooks = append(unknownTypeHooks, gue.HookFunc(hk))
	}
	unknownTypeHooks_ := gue.WithPoolHooksUnknownJobType(lockedHooks...)

	doneHooks := make([]gue.HookFunc, 0)
	for _, hk := range myPool.hooksJobDone {
		doneHooks = append(doneHooks, gue.HookFunc(hk))
	}
	doneHooks_ := gue.WithPoolHooksJobDone(doneHooks...)

	guePool := gue.NewWorkerPool(myPool.Client.Client, gueWm, myPool.Size,
		idOpt, logOpt, presvOpt, migrOpt, queOpt, pollInt, pollStr, lockedHooks_, unknownTypeHooks_, doneHooks_)
	myPool.WorkerPool = guePool

	return myPool
}

// Run runs all the Workers in the WorkerPool in own goroutines.
// Run blocks until all workers exit. Use context cancellation for
// shutdown.
func (w *WorkerPool) Run(ctx context.Context) error {
	return w.WorkerPool.Run(ctx)
}
