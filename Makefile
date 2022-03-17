SHELL := /usr/bin/env bash

GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH)

BINARY_BASENAME=qotm

DOCKER_REPO ?= localhost:31000/tour
TAG ?= latest

.PHONY: all build build.image image.push clean fmt run test.fast

all: clean fmt test.fast build

build: fmt
	$(GOBUILD) ./...
	ln -sf $(BINARY_BASENAME)-$(GOOS)-$(GOARCH) bin/$(BINARY_BASENAME)

run: build
	bin/qotm

build.image:
	docker build \
	-t $(DOCKER_REPO):$(TAG) \
	-f Dockerfile \
	.

image.push: build.image
	docker push \
	$(DOCKER_REPO):$(TAG)

clean:
	rm -rf bin

fmt:
	go fmt ./...

test.fast:
	go test -v ./...
