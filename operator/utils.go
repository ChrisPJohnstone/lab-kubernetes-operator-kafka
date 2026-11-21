package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime"
)

func setupLogging() {
	var level slog.Level
	switch os.Getenv(envarLogLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "ERROR":
		level = slog.LevelError
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	default:
		level = defaultLogLevel
	}
	opts := &slog.HandlerOptions{Level: level}
	handler := slog.NewJSONHandler(os.Stderr, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	controllerruntime.SetLogger(logr.FromSlogHandler(handler))
	slog.Debug(fmt.Sprintf("log level %s", level))
}
