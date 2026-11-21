package main

import (
	"context"
	"log/slog"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
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
	return reconcile.Result{}, nil
}

func main() {
	setupLogging()
	cfg := config.GetConfigOrDie()
	mgrOptions := manager.Options{}
	slog.Debug("initializing manager")
	mgr, err := manager.New(cfg, mgrOptions)
	if err != nil {
		panic(err)
	}
	client := mgr.GetClient()
	reconciler := &Reconciler{Client: client}
	slog.Debug("initializing builder")
	err = builder.ControllerManagedBy(mgr).
		For(&appsv1.ReplicaSet{}).
		Owns(&corev1.Pod{}).
		Complete(reconciler)
	if err != nil {
		panic(err)
	}
	slog.Debug("starting manager")
	err = mgr.Start(signals.SetupSignalHandler())
	if err != nil {
		panic(err)
	}
}
