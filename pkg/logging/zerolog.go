package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type zLog struct {
	log zerolog.Logger
}

func GetLogger() *zLog {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Logger()

	return &zLog{log: logger}
}

func (l *zLog) Debug(msg string, params map[string]interface{}) {
	l.log.Debug().Msg(msg)
}

func (l *zLog) Info(msg string, params map[string]interface{}) {
	l.log.Info().Msg(msg)
}

func (l *zLog) Warn(msg string, params map[string]interface{}) {
	l.log.Warn().Msg(msg)
}

func (l *zLog) Error(msg string, params map[string]interface{}) {
	l.log.Error().Msg(msg)
}

func (l *zLog) Fatal(msg string, params map[string]interface{}) {
	l.log.Fatal().Msg(msg)
}
