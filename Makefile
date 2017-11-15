.PHONY: dev build image test deps clean

CGO_ENABLED=0
COMMIT=`git rev-parse --short HEAD`
APP=irccat
PACKAGE=
REPO?=prologic/$(APP)
TAG?=latest
BUILD?=-dev

all: dev

dev: build
	@./$(APP)

deps:
	@go get ./...

build: clean deps
	@echo " -> Building $(TAG)$(BUILD)"
	@go build -tags "netgo static_build" -installsuffix netgo \
		-ldflags "-w -X github.com/$(REPO)/${PACKAGE}.GitCommit=$(COMMIT) -X github.com/$(REPO)/${PACKAGE}.Build=$(BUILD)" .
	@echo "Built $$(./$(APP) -v)"

image:
	@docker build --build-arg TAG=$(TAG) --build-arg BUILD=$(BUILD) -t $(REPO):$(TAG) .
	@echo "Image created: $(REPO):$(TAG)"

test:
	@go test -v -cover -race $(TEST_ARGS)

clean:
	@rm -rf $(APP)
