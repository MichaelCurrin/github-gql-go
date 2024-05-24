SHELL = /bin/bash

JSON_DIR = json_dumps
BUILD_DIR = build
COMPILED := $(BUILD_DIR)/ghgql

export GH_TOKEN


.PHONY: $(BUILD_DIR)


default: install

all: install fmt build


h help:
	@grep '^[a-z]' Makefile


install:
	go get ./...

upgrade:
	go get -u ./...

tidy:
	go mod tidy


fmt:
	go fmt ./...


usage:
	go run main.go -h


.check-env:
	@if [ ! -f .env ]; then \
		echo "Error: `.env` file not found. Please create it."; \
	fi

run: .check-env
	@source .env \
		&& go run main.go

# TODO: Toggle type based on args to avoid mixing in one JSON file.
run-write:
	$(MAKE) run > $(JSON_DIR)/export.json



build:
	go build -v -o $(COMPILED) main.go

global:
	go install
