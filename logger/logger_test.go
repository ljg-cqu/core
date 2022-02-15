package logger

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// a simple test
// TODO: more test
func TestLog(t *testing.T) {
	logger := New(WithLFSHook)
	require.NotNil(t, logger)

	for i := 0; i < 5; i++ {
		//logger.WithField("method", "POST").Panic("Panic message") // Written in /tmp/abfpaas-log/panic.log
		//logger.WithField("method", "POST").Fatal("Fatal message") // Written in /tmp/abfpaas-log/fatal.log
		logger.WithField("method", "POST").Error("Error message") // Written in /tmp/abfpaas-log/error.log
		logger.WithField("method", "POST").Warn("Warn message")   // Written in /tmp/abfpaas-log/warn.log
		logger.WithField("method", "POST").Info("Info message")   // Written in /tmp/abfpaas-log/info.log
		logger.WithField("method", "POST").Debug("Debug message") // It is not written to a file (because debug level < minLevel)
		time.Sleep(time.Second * 2)
	}
}
