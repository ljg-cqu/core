// Package logger provides wrapper of logrus.Log
// with additional features supported such as customized hooks

package logger

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/ljg-cqu/core/utils"
	"github.com/orandin/lumberjackrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"path"
	"runtime"

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

var (
	DefaultLogLevel           logrus.Level     = logrus.DebugLevel
	DefaultFormatter          logrus.Formatter = exampleFormatter
	DefaultEnableReportCaller                  = true
)

// powerful-logrus-formatter. get fileName, log's line number and the latest function's name when print log;
// Sava log to files. More: https://github.com/zput/zxcTool
var exampleFormatter = &zt_formatter.ZtFormatter{
	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		filename := path.Base(f.File)
		return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
	},
	Formatter: nested.Formatter{
		//HideKeys: true,
		FieldsOrder: []string{"component", "category"},
	},
}

var logger *Logger

var logLevels = map[string]logrus.Level{
	"PanicLevel": logrus.PanicLevel,
	"FatalLevel": logrus.FatalLevel,
	"ErrorLevel": logrus.ErrorLevel,
	"WarnLevel":  logrus.WarnLevel,
	"InfoLevel":  logrus.InfoLevel,
	"DebugLevel": logrus.DebugLevel,
	"TraceLevel": logrus.TraceLevel,
}

type Logger struct {
	ID string
	*logrus.Logger
}

// TODO: to be implemented
func NewByConfig(config string) *Logger {
	return nil
}

// New returns Log with given Hook(s), which defaults to DebugLevel
// if the logLevelEnv isn't specified appropriately.
func New(hkoptions ...Option) *Logger {
	if logger != nil {
		logger.WithField("logger_id", logger.ID).Debugln("Reuse logger")
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
		logLevel = DefaultLogLevel
	}

	log := logrus.New()
	log.SetLevel(logLevel)
	log.SetFormatter(DefaultFormatter)
	log.SetReportCaller(DefaultEnableReportCaller)

	if hkoptions != nil {
		for _, hk := range hkoptions {
			hk(logger)
		}
	}

	logger = &Logger{ID: utils.MD5NowID(), Logger: log}
	logger.WithField("logger_id", logger.ID).Debugln("Apply new logger")
	return logger
}

func (l *Logger) SetLevelTrace() *Logger {
	l.SetLevel(logrus.TraceLevel)
	return l
}

func (l *Logger) SetLevelDebug() *Logger {
	l.SetLevel(logrus.DebugLevel)
	return l
}

func (l *Logger) SetLevelInfo() *Logger {
	l.SetLevel(logrus.InfoLevel)
	return l
}

func (l *Logger) SetFormatterJSON() *Logger {
	l.SetFormatter(&logrus.JSONFormatter{})
	return l
}

func (l *Logger) SetFormatterText() *Logger {
	l.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	return l
}

func (l *Logger) SetLFSHook(logDir string) {
	WithLFSHook(logDir)(l)
}

type Option func(logger *Logger)

func WithLogLevel(lvl logrus.Level) Option {
	return func(logger *Logger) {
		logger.SetLevel(lvl)
	}
}

func WithFormatter(fmtr logrus.Formatter) Option {
	return func(logger *Logger) {
		logger.SetFormatter(fmtr)
	}
}

// WithLFSHook adds hook for logging to the local filesystem,
// which defaults "/tmp/abfpaas-log/" if the logDirEnv isn't specified appropriately,
// with logrotate and a file per log level from InfoLevel
func WithLFSHook(logDir string) Option {
	return func(logger *Logger) {
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
