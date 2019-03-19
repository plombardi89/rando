SHELL := /usr/bin/env bash

all: fmt test

fmt:
	go fmt ./...

test:
	go test -v ./...