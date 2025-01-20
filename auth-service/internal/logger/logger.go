package logger

import (
	"log/slog"
	"os"
)

var(
	AuthLogger = slog.New(slog.NewTextHandler(os.Stderr, nil))
)