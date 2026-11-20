BINARY           ?= kafka-operator
CONTAINER_ENGINE ?= podman
IMAGE_NAME       ?= localhost/kafka-operator:dev

.DEFAULT_GOAL := $(BINARY)
.PHONY: all image clean lint

all: $(BINARY) image

$(BINARY): $(shell find . -type f -name '*.go')
	go build -o $(BINARY) .

image: $(shell find . -type f -name '*.go') Containerfile
	$(CONTAINER_ENGINE) build . -t $(IMAGE_NAME)

clean:
	rm -f $(BINARY)
	$(CONTAINER_ENGINE) rmi -i $(IMAGE_NAME)

lint:
	$(CONTAINER_ENGINE) run --rm -v "$(PWD):/app:Z" -w /app \
		docker.io/golangci/golangci-lint:v2.14.0 \
		golangci-lint run
