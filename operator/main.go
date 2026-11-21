package main

import (
	"fmt"
	"log/slog"
	"os"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
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
	slog.Debug(fmt.Sprintf("Log level %s", level))
}

func main() {
	setupLogging()
	cfg := config.GetConfigOrDie()
	mgrOptions := manager.Options{}
	mgr, err := manager.New(cfg, mgrOptions)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Manager: %s\n", mgr)
}
