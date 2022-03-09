package adapter

import (
	"github.com/ljg-cqu/core/logger"
	"github.com/sirupsen/logrus"
	"github.com/vgarvardt/gue/v3/adapter"
)

// gueAdapter represents logger adapter for https://github.com/vgarvardt/gue/blob/master/adapter/logger.go
type gueAdapter struct {
	l *logger.Logger
	logrus.Logger
}

// NewGueLogger instantiates new adapter.Logger using sirupsen/logrus
func NewGueLogger(log *logger.Logger) adapter.Logger {
	var l = &gueAdapter{
		l: log,
	}
	return l
}

// Debug implements Logger.Debug for sirupsen/logrus logger
func (l *gueAdapter) Debug(msg string, fields ...adapter.Field) {
	fieldsl := logrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Debug(msg)
}

// Info implements Logger.Debug for sirupsen/logrus logger
func (l *gueAdapter) Info(msg string, fields ...adapter.Field) {
	fieldsl := logrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Info(msg)
}

// Error implements Logger.Debug for sirupsen/logrus logger
func (l *gueAdapter) Error(msg string, fields ...adapter.Field) {
	fieldsl := logrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Error(msg)
}

// With implements nested logrus for sirupsen/logrus logger
func (l *gueAdapter) With(fields ...adapter.Field) adapter.Logger {
	fieldsl := logrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl)
	return l
}
