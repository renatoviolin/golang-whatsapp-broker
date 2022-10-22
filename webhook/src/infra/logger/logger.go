package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLog() {
	writer := io.MultiWriter(os.Stdout)
	Log = zerolog.New(writer).With().Timestamp().Logger()
}

func Error(from string, method string, msg string) {
	Log.Error().Str("from", from).Str("method", method).Str("message", msg).Send()
}

func Info(from string, method string, msg string) {
	Log.Info().Str("from", from).Str("method", method).Str("message", msg).Send()
}

func Fatal(from string, method string, msg string) {
	Log.Fatal().Str("from", from).Str("method", method).Str("message", msg).Send()
}

func Panic(from string, method string, msg string) {
	Log.Panic().Str("from", from).Str("method", method).Str("message", msg).Send()
}
