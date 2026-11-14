SERVER_DIR := server
<<<<<<< HEAD
=======
CLIENT_DIR := client
>>>>>>> 1fc210190715df134436ed7c0f6d577ca16d696f
BIN_DIR := bin
PORT ?= 4242
GO_SRC := $(shell find $(SERVER_DIR) -name '*.go' ! -name '*_test.go')
GOBIN := $(shell go env GOPATH)/bin
GOLANGCI := $(GOBIN)/golangci-lint
GOLANGCI_VERSION := v2.15.0

all: build

install:
	go -C $(SERVER_DIR) mod download
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(GOBIN) $(GOLANGCI_VERSION)

<<<<<<< HEAD
build: lint
	go -C $(SERVER_DIR) build -o ../$(BIN_DIR)/server
=======
build:
	go -C $(SERVER_DIR) build -o ../$(BIN_DIR)/server
	go -C $(CLIENT_DIR) build -o ../$(BIN_DIR)/client
>>>>>>> 1fc210190715df134436ed7c0f6d577ca16d696f

run-server: build
	./$(BIN_DIR)/server

run-client:
<<<<<<< HEAD
	@echo "run-client: pas encore implémenté" && exit 1
=======
	./$(BIN_DIR)/client
>>>>>>> 1fc210190715df134436ed7c0f6d577ca16d696f

run-client-gui:
	@echo "run-client-gui: pas encore implémenté" && exit 1

lint:
	cd $(SERVER_DIR) && $(GOLANGCI) run ./...
	@! grep -nw goto $(GO_SRC) || (echo "goto interdit" && exit 1)
	@grep -cH '^func ' $(GO_SRC) | awk -F: '$$2 > 5 {print $$1": "$$2" fonctions (max 5)"; bad=1} END {exit bad}'

clean:
	go -C $(SERVER_DIR) clean
	rm -rf $(BIN_DIR)

.PHONY: all install build run-server run-client run-client-gui lint clean
