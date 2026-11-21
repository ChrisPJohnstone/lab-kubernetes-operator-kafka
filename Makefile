CONTAINER_ENGINE ?= podman
IMAGE_NAME       ?= localhost/kafka-operator:dev
DEPLOYMENT_NAME  ?= kafka-operator

.PHONY: all
all: kafka deployment

.PHONY: kafka
kafka:
	@if ! $(CONTAINER_ENGINE) container exists kafka; then \
		$(CONTAINER_ENGINE) compose up -d; \
	fi

.PHONY: operator
operator:
	make -C operator image IMAGE_NAME=$(IMAGE_NAME)

.PHONY: cluster
cluster:
	@minikube status || minikube start

.PHONY: deployment
deployment: operator cluster
	$(CONTAINER_ENGINE) save $(IMAGE_NAME) | minikube image load -
	kubectl apply -f manifests/controller.py

.PHONY: clean
clean: clean-kafka clean-cluster clean-operator

.PHONY: clean-kafka
clean-kafka:
	$(CONTAINER_ENGINE) compose down -v --rmi
	$(CONTAINER_ENGINE) unshare rm -rf output

.PHONY: clean-operator
clean-operator:
	make -C operator clean IMAGE_NAME=$(IMAGE_NAME)

.PHONY: clean-cluster
clean-cluster:
	minikube delete

.PHONY: clean-deployment
clean-deployment:
	kubectl delete deployment $(DEPLOYMENT_NAME) --ignore-not-found
	@if minikube image ls | grep -q localhost; then \
		minikube image rm $(IMAGE_NAME); \
	fi
