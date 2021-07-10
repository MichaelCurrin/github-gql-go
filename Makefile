export GH_TOKEN

OUT_DIR = build
COMPILED := $(OUT_DIR)/ghgql

.PHONY: $(OUT_DIR)


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
	$(MAKE) run > export.json



build:
	go build -v -o $(COMPILED) main.go


global:
	go install
