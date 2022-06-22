GOBIN = ./bin
GO ?= latest

APP_NAME = proj1

.PHONY: build run test clean

build:
	go build -o $(GOBIN)/$(APP_NAME) ./cmd/$(APP_NAME)/main.go
	@echo "Done building."
	@echo "Run \"$(GOBIN)/$(APP_NAME)\" to launch $(APP_NAME)."

run:
	go run ./cmd/$(APP_NAME)/main.go
	@echo "Done running"

test:
	go test -v ./...
	@echo "Done testing"

clean:
	go clean -cache
	rm -r bin
	@echo "Done cleaning"
