package main

import (
	"log/slog"
	"os"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(l)
	l.Info("Application was Initialized")
}
