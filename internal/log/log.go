package log

import (
	"io"
	"log/slog"
)

// New creates a new slog logger with a destination and a minimum log level.
func New(dest io.Writer, minLevel slog.Leveler) *slog.Logger {
	handler := slog.NewTextHandler(dest, &slog.HandlerOptions{
		AddSource: false,
		Level:     minLevel,
	})

	return slog.New(handler)
}
