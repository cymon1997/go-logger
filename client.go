package logger

import (
	"io"
	"log"
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
			Out:          getOutput(cfg),
			Hooks:        make(logrus.LevelHooks),
			Formatter:    getFormatter(cfg),
			ReportCaller: false,
			Level:        getLevel(cfg),
			ExitFunc:     os.Exit,
		}
	})
}

func getOutput(cfg Config) io.Writer {
	if cfg.OutputFile == "" {
		return os.Stderr
	}
	file, err := os.OpenFile(cfg.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("invalid output file")
	}
	return file
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
