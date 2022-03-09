package postgres_job

import (
	"github.com/ljg-cqu/core/logger"
	"time"
)

// WorkerPoolOption defines a type that allows to set worker pool properties during the build-time.
type WorkerPoolOption func(pool *WorkerPool)

// WithPoolPollInterval overrides default poll interval with the given value.
// Poll interval is the "sleep" duration if there were no jobs found in the DB.
func WithPoolPollInterval(d time.Duration) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.pollInterval = d
	}
}

// WithPoolPollStrategy overrides default poll strategy with given value
func WithPoolPollStrategy(s PollStrategy) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.pollStrategy = s
	}
}

// WithPoolQueue overrides default worker queue name with the given value.
func WithPoolQueue(queue QueueName) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.Queue = queue
	}
}

// WithPoolID sets worker pool ID for easier identification in logs
func WithPoolID(id string) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.ID = id
	}
}

// WithPoolLogger sets Logger implementation to worker pool
func WithPoolLogger(logger *logger.Logger) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.Logger = logger
	}
}

func WithPoolPreserveCompletedJobs(preserve bool) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.preserveCompletedJobs = preserve
	}
}

func WithPoolMigrateCompletedJobs(migrate bool) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.migrateCompletedJobs = migrate
	}
}

// WithPoolHooksJobLocked calls WithWorkerHooksJobLocked for every worker in the pool.
func WithPoolHooksJobLocked(hooks ...HookFunc) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.hooksJobLocked = hooks
	}
}

// WithPoolHooksUnknownJobType calls WithWorkerHooksUnknownJobType for every worker in the pool.
func WithPoolHooksUnknownJobType(hooks ...HookFunc) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.hooksUnknownJobType = hooks
	}
}

// WithPoolHooksJobDone calls WithWorkerHooksJobDone for every worker in the pool.
func WithPoolHooksJobDone(hooks ...HookFunc) WorkerPoolOption {
	return func(pool *WorkerPool) {
		pool.hooksJobDone = hooks
	}
}
