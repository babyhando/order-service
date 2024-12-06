package logger

import (
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With("trace_id", uuid.NewString())
}
