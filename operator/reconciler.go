package main

import (
	"context"
	"fmt"
	"log/slog"

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
	slog.Debug(fmt.Sprintf("Got topic '%s'", topic.Spec.Name))
	return reconcile.Result{}, nil
}
