package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

func initKgoClient() (*kgo.Client, error) {
	slog.Debug("initializing kgo client")
	seedBrokers := kgo.SeedBrokers(kafkaAddress)
	kgoClient, err := kgo.NewClient(seedBrokers)
	if err != nil {
		return nil, err
	}
	defer kgoClient.Close()
	return kgoClient, nil
}

func initKadmClient(kgoClient *kgo.Client) *kadm.Client {
	slog.Debug("initializing kadm client")
	kadmClient := kadm.NewClient(kgoClient)
	defer kadmClient.Close()
	return kadmClient
}

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
