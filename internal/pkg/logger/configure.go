package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/penguin-statistics/livehouse/internal/config"
)

func Configure(conf *config.Config) {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	_ = os.Mkdir("logs", os.ModePerm)

	zerologLogLevelMap := map[string]zerolog.Level{
		"debug": zerolog.DebugLevel,
		"info":  zerolog.InfoLevel,
		"warn":  zerolog.WarnLevel,
		"error": zerolog.ErrorLevel,
		"fatal": zerolog.FatalLevel,
		"panic": zerolog.PanicLevel,
		"trace": zerolog.TraceLevel,
	}

	level, ok := zerologLogLevelMap[conf.LogLevel]
	if !ok {
		level = zerolog.InfoLevel
	}

	var writer io.Writer

	if conf.LogJsonStdout {
		writer = os.Stdout
	} else {
		writer = zerolog.MultiLevelWriter(
			&lumberjack.Logger{
				Filename: "logs/app.log",
				MaxSize:  100, // megabytes
				MaxAge:   90,  // days
				Compress: true,
			},
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339Nano,
			},
		)
	}

	log.Logger = zerolog.New(writer).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(level)
}
