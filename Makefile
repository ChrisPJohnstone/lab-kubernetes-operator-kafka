CONTAINER_ENGINE ?= podman

.PHONY: kafka operator cluster clean

all: kafka operator cluster

kafka:
	$(CONTAINER_ENGINE) compose up -d

operator:
	make -C operator image

cluster:
	minikube start

clean:
	minikube delete
	make -C operator clean
	$(CONTAINER_ENGINE) compose down -v
	$(CONTAINER_ENGINE) unshare rm -rf output
