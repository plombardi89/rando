SHELL := /usr/bin/env bash

all: fmt

fmt:
	go fmt ./...

test:
	go test -v ./...