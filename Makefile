.PHONY: dev build install image test deps clean

CGO_ENABLED=0
COMMIT=$(shell git rev-parse --short HEAD)

all: dev

dev: build
	@./irccat -v

build:
	@go build \
		-tags "netgo static_build" -installsuffix netgo \
		-ldflags "-w -X $(shell go list)/.Commit=$(COMMIT)" \
		.

install: build
	@go install

image:
	@docker build -t prologic/irccat .

test:
	@go test -v -cover -race .

clean:
	@git clean -f -d -X
