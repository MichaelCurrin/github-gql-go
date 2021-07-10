export GH_TOKEN

JSON_DIR = json_dumps
BUILD_DIR = build
COMPILED := $(BUILD_DIR)/ghgql

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

run:
	@source .env \
		&& go run main.go

run-write:
	$(MAKE) run > $(JSON_DIR)/export.json



build:
	go build -v -o $(COMPILED) main.go


global:
	go install
