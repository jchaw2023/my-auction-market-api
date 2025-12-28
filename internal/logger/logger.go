package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	log        zerolog.Logger
	gormLogger zerolog.Logger
)

func Init(level string) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = "time"

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    true,
	}

	noColorOutput := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    true,
	}

	var logLevel zerolog.Level
	switch level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	log = zerolog.New(output).With().Timestamp().Logger()
	gormLogger = zerolog.New(noColorOutput).With().Timestamp().Logger()
}

func Debug(format string, v ...interface{}) {
	log.Debug().Msg(fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	log.Info().Msg(fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	log.Warn().Msg(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	log.Error().Msg(fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...interface{}) {
	log.Fatal().Msg(fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	Fatal(format, v...)
}

func GormDebug(format string, v ...interface{}) {
	gormLogger.Debug().Msg(fmt.Sprintf(format, v...))
}

func GormInfo(format string, v ...interface{}) {
	gormLogger.Info().Msg(fmt.Sprintf(format, v...))
}

func GormWarn(format string, v ...interface{}) {
	gormLogger.Warn().Msg(fmt.Sprintf(format, v...))
}

func GormError(format string, v ...interface{}) {
	gormLogger.Error().Msg(fmt.Sprintf(format, v...))
}
