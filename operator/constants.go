package main

import "log/slog"

const (
	envarLogLevel       = "LOG_LEVEL"
	defaultLogLevel     = slog.LevelInfo
	crdGroup            = "lab.com"
	kafkaTopicVersion   = "v1"
	kafkaAddress        = "host.minikube.internal:9092"
	kafkaTopicFinalizer = "lab.com/kafkatopic-finalizer"
)
