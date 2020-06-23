GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=gbrn

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go
run:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go
	./$(BINARY_NAME)
test:
	$(GOTEST) ./internal/util -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
