package main

import (
	"log/slog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func initScheme() *runtime.Scheme {
	slog.Debug("initializing scheme")
	scheme := runtime.NewScheme()
	utilruntime.Must(clientscheme.AddToScheme(scheme))
	groupVersion := schema.GroupVersion{
		Group:   crdGroup,
		Version: kafkaTopicVersion,
	}
	scheme.AddKnownTypes(groupVersion, &KafkaTopic{}, &KafkaTopicList{})
	metav1.AddToGroupVersion(scheme, groupVersion)
	return scheme
}

func main() {
	setupLogging()
	scheme := initScheme()
	cfg := config.GetConfigOrDie()
	mgrOptions := manager.Options{Scheme: scheme}
	slog.Debug("initializing manager")
	mgr, err := manager.New(cfg, mgrOptions)
	if err != nil {
		panic(err)
	}
	slog.Debug("initializing builder")
	client := mgr.GetClient()
	reconciler := &Reconciler{Client: client}
	bldr := builder.ControllerManagedBy(mgr)
	err = bldr.For(&KafkaTopic{}).Complete(reconciler)
	if err != nil {
		panic(err)
	}
	slog.Debug("starting manager")
	err = mgr.Start(signals.SetupSignalHandler())
	if err != nil {
		panic(err)
	}
}
