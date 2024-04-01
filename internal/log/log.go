package log

import (
	"io"
	"log/slog"
)

func New(dest io.Writer, minLevel slog.Leveler) *slog.Logger {
	handler := slog.NewTextHandler(dest, &slog.HandlerOptions{
		AddSource: false,
		Level:     minLevel,
	})

	return slog.New(handler)
}
