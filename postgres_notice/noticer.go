package postgres_notice

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ljg-cqu/core/logger"
	"github.com/pkg/errors"
)

// TODO: include backoff feature

type Noticer struct {
	logger  *logger.Logger
	pool    *pgxpool.Pool
	channel string
}

// NewNoticer returns  a Noticer that can be used to send notification.
func NewNoticer(log *logger.Logger, pool *pgxpool.Pool, channel string) *Noticer {
	return &Noticer{log, pool, channel}
}

// Notify send a notification to the channel given when call NewNoticer.
func (n *Noticer) Notify(ctx context.Context, payload string) error {
	if err := Notify(ctx, n.pool, n.channel, payload); err != nil {
		n.logger.WithError(err).WithField("channel", n.channel).Errorln("Failed to send a notification")
		return errors.Wrapf(err, "failed to send notification, channel:%s", n.channel)
	}
	return nil
}

// NotifyCh send a notification to a new channel.
func (n *Noticer) NotifyCh(ctx context.Context, channel, payload string) error {
	if err := Notify(ctx, n.pool, channel, payload); err != nil {
		n.logger.WithError(err).WithField("channel", n.channel).Errorln("Failed to send a notification")
		return errors.Wrapf(err, "failed to send notification, channel:%s", n.channel)
	}
	return nil
}

// Notify is a handy function to sends a notification on the channel using PostgreSQL `pg_notify` function.
// Note: this function won't log anything when it gets an error.
func Notify(ctx context.Context, pool *pgxpool.Pool, channel, payload string) error {
	_, err := pool.Exec(ctx, "select pg_notify($1, $2)", channel, payload)
	if err != nil {
		return errors.Wrapf(err, "failed to send notification, channel:%s", channel)
	}
	return nil
}
