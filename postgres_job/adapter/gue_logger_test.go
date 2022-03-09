package adapter

import (
	"errors"
	"github.com/ljg-cqu/core/logger"
	"github.com/vgarvardt/gue/v3/adapter"
	"testing"
)

func TestNew(t *testing.T) {
	l := logger.New()

	ll := NewGueLogger(l)
	err := errors.New("something went wrong")

	ll.Debug("debug-1", adapter.F("debug-key", "debug-val"))
	ll.Info("info-1", adapter.F("info-key", "info-val"))
	ll.Error("error-1", adapter.F("error-key", "error-val"))
	ll.Error("error-2", adapter.Err(err))

	lll := ll.With(adapter.F("nested-key", "nested-val"))
	lll.Info("info-2", adapter.F("info-key-2", "info-val-2"))
}
