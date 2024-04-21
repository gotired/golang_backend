# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
BINARY_NAME := golang-backend

all: build run

dev:
	$(GOCMD) run cmd/main.go

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd

docker:
	docker compose up --build

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	./$(BINARY_NAME)
