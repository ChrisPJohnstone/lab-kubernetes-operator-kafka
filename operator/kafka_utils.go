package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kadm"
)

func createKafkaTopic(
	kadmClient *kadm.Client,
	ctx context.Context,
	name string,
	partitions int32,
	replicationFactor int16,
) error {
	slog.Debug(fmt.Sprintf("Creating topic %s", name))
	resp, err := kadmClient.CreateTopic(
		ctx,
		partitions,
		replicationFactor,
		nil,
		name,
	)
	slog.Debug(fmt.Sprintf("Create topic response: %v", resp))
	if err != nil {
		return err
	}
	return nil
}

func deleteKafkaTopic(
	kadmClient *kadm.Client,
	ctx context.Context,
	name string,
) error {
	slog.Debug(fmt.Sprintf("Deleting topic %s", name))
	resp, err := kadmClient.DeleteTopic(ctx, name)
	slog.Debug(fmt.Sprintf("Delete topic response: %v", resp))
	if err != nil {
		return err
	}
	return nil
}
