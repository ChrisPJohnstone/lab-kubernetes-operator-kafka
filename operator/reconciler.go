package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
	client.Client
}

func (reconciler *Reconciler) Reconcile(
	ctx context.Context,
	req reconcile.Request,
) (reconcile.Result, error) {
	slog.Debug("starting reconciler")
	var topic KafkaTopic
	err := reconciler.Get(ctx, req.NamespacedName, &topic)
	topicName := topic.Spec.Name
	if err != nil {
		slog.Error(err.Error())
		return reconcile.Result{}, err
	}
	slog.Debug(fmt.Sprintf("Reconciling topic '%s'", topicName))

	seedBrokers := kgo.SeedBrokers(kafkaAddress)
	kgoClient, err := kgo.NewClient(seedBrokers)
	if err != nil {
		slog.Error(err.Error())
		return reconcile.Result{}, err
	}
	defer kgoClient.Close()
	kadmClient := kadm.NewClient(kgoClient)
	slog.Debug(fmt.Sprintf("Creating topic %s", topicName))
	resp, err := kadmClient.CreateTopic(ctx, -1, -1, nil, topicName)
	if err != nil && err != kerr.TopicAlreadyExists {
		slog.Error(err.Error())
		return reconcile.Result{}, err
	}
	slog.Debug(fmt.Sprintf("Created topic succesfully, response: %v", resp))
	return reconcile.Result{}, nil
}
