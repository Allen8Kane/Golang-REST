.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	 

.DEFAULT_GOAL := build