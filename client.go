package logger

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	config   Config
	logger   *logrus.Logger
	syncInit sync.Once
)

func init() {
	logger = logrus.New()
}

func Init(cfg Config) {
	syncInit.Do(func() {
		config = cfg
		logger = &logrus.Logger{
			Out:          os.Stderr,
			Hooks:        make(logrus.LevelHooks),
			Formatter:    getFormatter(cfg),
			ReportCaller: false,
			Level:        getLevel(cfg),
			ExitFunc:     os.Exit,
		}
	})
}

func getFormatter(cfg Config) logrus.Formatter {
	switch cfg.Format {
	case JSONFormat:
		return &logrus.JSONFormatter{
			TimestampFormat:   time.RFC3339,
			DisableHTMLEscape: true,
			DataKey:           "data",
			PrettyPrint:       cfg.IsBeautify,
		}
	case TextFormat:
		return &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		}
	default:
		return nil
	}
}

func getLevel(cfg Config) logrus.Level {
	switch cfg.Level {
	case PanicLevel:
		return logrus.PanicLevel
	case FatalLevel:
		return logrus.FatalLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case WarnLevel:
		return logrus.WarnLevel
	case InfoLevel:
		return logrus.InfoLevel
	case DebugLevel:
		return logrus.DebugLevel
	case TraceLevel:
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
