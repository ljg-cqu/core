// Package logger provides wrapper of logrus.Log
// with additional features supported such as customized hooks

package logger

import (
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	logLevelEnv = "ABFPAAS_LOG_LEVEL"
	logDirEnv   = "ABFPAAS_LOG_DIR"

	logFileMaxSize    = 2
	logFileMaxBackups = 1
	logFileMaxAge     = 1
	logFileCompress   = false
	logFileLocalTime  = true
)

var logger *logrus.Logger

var logLevels = map[string]logrus.Level{
	"PanicLevel": logrus.PanicLevel,
	"FatalLevel": logrus.FatalLevel,
	"ErrorLevel": logrus.ErrorLevel,
	"WarnLevel":  logrus.WarnLevel,
	"InfoLevel":  logrus.InfoLevel,
	"DebugLevel": logrus.DebugLevel,
	"TraceLevel": logrus.TraceLevel,
}

// New returns Log with given Hook(s), which defaults to InfoLevel
// if the logLevelEnv isn't specified appropriately.
func New(hkoptions ...Option) *logrus.Logger {
	if logger != nil {
		return logger
	}

	var logLevel logrus.Level
	if level, exits := os.LookupEnv(logLevelEnv); exits && level != "" {
		if value, ok := logLevels[level]; !ok {
			println("Unsupported logger level value: ", value)
			os.Exit(1)
		} else {
			logLevel = value
		}
	} else {
		logLevel = logrus.InfoLevel
	}

	logger = logrus.New()
	logger.SetLevel(logLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	if hkoptions != nil {
		for _, hk := range hkoptions {
			hk(logger)
		}
	}

	return logger
}

func NewForDebugJSON() *logrus.Logger {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

func NewForDebugStr() *logrus.Logger {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	return logger
}

type Option func(logger *logrus.Logger)

func WithLogLevel(lvl logrus.Level) Option {
	return func(logger *logrus.Logger) {
		logger.SetLevel(lvl)
	}
}

func WithFormatter(fmtr logrus.Formatter) Option {
	return func(logger *logrus.Logger) {
		logger.SetFormatter(fmtr)
	}
}

// WithLFSHook adds hook for logging to the local filesystem,
// which defaults "/tmp/abfpaas-log/" if the logDirEnv isn't specified appropriately,
// with logrotate and a file per log level from InfoLevel
func WithLFSHook(logDir string) Option {
	return func(logger *logrus.Logger) {
		if logDir == "" {
			logDir = "/tmp/abfpaas-log/"

			dirEnv, exists := os.LookupEnv(logDirEnv)
			if exists && dirEnv != "" {
				logDir = dirEnv
			}
		}

		fileInfo, err := os.Stat(logDir)
		if err != nil {
			println("Path error: ", logDir)
			os.Exit(1)
		}
		if !fileInfo.IsDir() {
			println(logDir, " is not a directory")
			os.Exit(1)
		}

		logDir = strings.TrimRight(logDir, "/")

		hook, err := lumberjackrus.NewHook(
			&lumberjackrus.LogFile{
				Filename:   logDir + "/general.log",
				MaxSize:    logFileMaxSize,
				MaxBackups: logFileMaxBackups,
				MaxAge:     logFileMaxAge,
				Compress:   logFileCompress,
				LocalTime:  logFileLocalTime,
			},
			logrus.InfoLevel,
			&logrus.JSONFormatter{},
			&lumberjackrus.LogFileOpts{
				logrus.PanicLevel: &lumberjackrus.LogFile{
					Filename:   logDir + "/panic.log",
					MaxSize:    logFileMaxSize,
					MaxBackups: logFileMaxBackups,
					MaxAge:     logFileMaxAge,
					Compress:   logFileCompress,
					LocalTime:  logFileLocalTime,
				},
				logrus.FatalLevel: &lumberjackrus.LogFile{
					Filename:   logDir + "/fatal.log",
					MaxSize:    logFileMaxSize,
					MaxBackups: logFileMaxBackups,
					MaxAge:     logFileMaxAge,
					Compress:   logFileCompress,
					LocalTime:  logFileLocalTime,
				},
				logrus.ErrorLevel: &lumberjackrus.LogFile{
					Filename:   logDir + "/error.log",
					MaxSize:    logFileMaxSize,
					MaxBackups: logFileMaxBackups,
					MaxAge:     logFileMaxAge,
					Compress:   logFileCompress,
					LocalTime:  logFileLocalTime,
				},
				logrus.WarnLevel: &lumberjackrus.LogFile{
					Filename:   logDir + "/warn.log",
					MaxSize:    logFileMaxSize,
					MaxBackups: logFileMaxBackups,
					MaxAge:     logFileMaxAge,
					Compress:   logFileCompress,
					LocalTime:  logFileLocalTime,
				},
				logrus.InfoLevel: &lumberjackrus.LogFile{
					Filename:   logDir + "/info.log",
					MaxSize:    logFileMaxSize,
					MaxBackups: logFileMaxBackups,
					MaxAge:     logFileMaxAge,
					Compress:   logFileCompress,
					LocalTime:  logFileLocalTime,
				},
			},
		)

		if err != nil {
			println("Failed to enable hook, error: ", err.Error())
			os.Exit(1)
		}
		logger.AddHook(hook)
	}
}
