GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=gbrn

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) cmd/main.go
	./$(BINARY_NAME)
