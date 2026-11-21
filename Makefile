CONTAINER_ENGINE ?= podman
IMAGE_NAME       ?= localhost/kafka-operator:dev
DEPLOYMENT_NAME  ?= kafka-operator

.PHONY: all
all: kafka operator

.PHONY: kafka
kafka:
	@if ! $(CONTAINER_ENGINE) container exists kafka; then \
		$(CONTAINER_ENGINE) compose up -d; \
	fi

.PHONY: operator-image
operator-image:
	make -C operator image IMAGE_NAME=$(IMAGE_NAME)

.PHONY: cluster
cluster:
	@minikube status || minikube start

.PHONY: operator
operator: operator-image cluster
	@if minikube image ls | grep -q localhost; then \
		minikube image rm $(IMAGE_NAME); \
    fi
	$(CONTAINER_ENGINE) save $(IMAGE_NAME) | minikube image load -
	kubectl apply -f manifests/controller.yaml

.PHONY: clean
clean: clean-kafka clean-operator-image clean-cluster

.PHONY: clean-kafka
clean-kafka:
	@if $(CONTAINER_ENGINE) container exists kafka; then \
        $(CONTAINER_ENGINE) compose down -v --rmi; \
        $(CONTAINER_ENGINE) unshare rm -rf output; \
    fi

.PHONY: clean-operator-image
clean-operator-image:
	make -C operator clean IMAGE_NAME=$(IMAGE_NAME)

.PHONY: clean-cluster
clean-cluster:
	minikube delete

.PHONY: clean-operator
clean-operator:
	kubectl delete deployment $(DEPLOYMENT_NAME) --ignore-not-found
	@if minikube image ls | grep -q localhost; then \
		minikube image rm $(IMAGE_NAME); \
	fi
