package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kerr"
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
	if err != nil {
		return reconcile.Result{}, err
	}
	topicName := topic.Spec.Name
	slog.Debug(fmt.Sprintf("Reconciling topic '%s'", topicName))

	kgoClient, err := initKgoClient()
	if err != nil {
		return reconcile.Result{}, err
	}
	kadmClient := initKadmClient(kgoClient)

	err = createKafkaTopic(
		kadmClient,
		ctx,
		topicName,
		topic.Spec.Partitions,
		topic.Spec.ReplicationFactor,
	)
	if err == kerr.TopicAlreadyExists {
		slog.Debug(fmt.Sprintf("Topic '%s' already exists", topicName))
		return reconcile.Result{}, nil
	}
	if err != nil {
		return reconcile.Result{}, err
	}
	slog.Debug("Created topic succesfully")
	return reconcile.Result{}, nil
}
