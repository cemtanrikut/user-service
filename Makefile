APP_NAME := user-service
DOCKER_IMAGE := $(APP_NAME)
DOCKER_TAG := latest

GO_FILES := $(shell find . -name '*.go')
GO_BUILD := go build -o $(APP_NAME) ./cmd/user-service

DOCKER_BUILD := docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: all
all: build

.PHONY: build
build: $(GO_FILES)
	@echo "Building Go binary..."
	$(GO_BUILD)

.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	$(DOCKER_BUILD)

.PHONY: run
run: docker-build
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

.PHONY: fmt
fmt:
	@echo "Checking code format..."
	go fmt ./...

.PHONY: mod
mod:
	@echo "Tidying Go modules..."
	go mod tidy

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
