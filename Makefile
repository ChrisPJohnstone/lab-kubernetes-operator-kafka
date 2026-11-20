CONTAINER_ENGINE ?= podman

.PHONY: kafka cluster destroy

all: kafka cluster

kafka:
	$(CONTAINER_ENGINE) compose up -d

cluster:
	minikube start

destroy:
	minikube delete
	$(CONTAINER_ENGINE) compose down -v
	$(CONTAINER_ENGINE) unshare rm -rf output
