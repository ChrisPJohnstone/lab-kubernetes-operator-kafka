CONTAINER_ENGINE ?= podman
IMAGE_NAME       ?= localhost/kafka-operator:dev
DEPLOYMENT_NAME  ?= kafka-operator

.PHONY: kafka operator cluster deployment clean

all: kafka deployment

kafka:
	$(CONTAINER_ENGINE) compose up -d

operator:
	make -C operator image IMAGE_NAME=$(IMAGE_NAME)

cluster:
	minikube start

deployment: operator cluster
	$(CONTAINER_ENGINE) save $(IMAGE_NAME) | minikube image load -
	kubectl create deployment $(DEPLOYMENT_NAME) --image $(IMAGE_NAME)

clean:
	minikube delete
	make -C operator clean IMAGE_NAME=$(IMAGE_NAME)
	$(CONTAINER_ENGINE) compose down -v --rmi
	$(CONTAINER_ENGINE) unshare rm -rf output
