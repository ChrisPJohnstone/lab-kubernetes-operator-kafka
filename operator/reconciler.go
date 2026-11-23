package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
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

	slog.Debug("initializing kgo client")
	seedBrokers := kgo.SeedBrokers(kafkaAddress)
	kgoClient, err := kgo.NewClient(seedBrokers)
	if err != nil {
		return reconcile.Result{}, err
	}
	defer kgoClient.Close()
	slog.Debug("initializing kadm client")
	kadmClient := kadm.NewClient(kgoClient)
	defer kadmClient.Close()

	if !topic.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(&topic, kafkaTopicFinalizer) {
			return reconcile.Result{}, nil
		}
		err := deleteKafkaTopic(kadmClient, ctx, topicName)
		if err != nil {
			return reconcile.Result{}, reconciler.Update(ctx, &topic)
		}
		return reconcile.Result{}, nil
	}

	if controllerutil.AddFinalizer(&topic, kafkaTopicFinalizer) {
		return reconcile.Result{}, reconciler.Update(ctx, &topic)
	}
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
