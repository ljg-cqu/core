package postgres_job

import (
	"github.com/ljg-cqu/core/logger"
)

// ClientOption defines a type that allows to set client properties during the build-time.
type ClientOption func(*Client)

// WithClientLogger sets Logger implementation to client
func WithClientLogger(logger *logger.Logger) ClientOption {
	return func(client *Client) {
		client.Logger = logger
	}
}

// WithClientID sets client ID for easier identification in logs
func WithClientID(id string) ClientOption {
	return func(client *Client) {
		client.ID = id
	}
}

// WithClientBackoff sets backoff implementation that will be applied to errored jobs
// within current client session.
func WithClientBackoff(backoff Backoff) ClientOption {
	return func(client *Client) {
		client.backoff = backoff
	}
}
