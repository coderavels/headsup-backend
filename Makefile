GOCMD=go
GOBUILD=$(GOCMD) build

BINARY_NAME=headsupbackend
ZIP_NAME=headsupbackend

build:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME) main.go

build-lambda:
	mkdir -p bin/
	zip bin/$(ZIP_NAME).zip $(BINARY_NAME)
	rm $(BINARY_NAME)
