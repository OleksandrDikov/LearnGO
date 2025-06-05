.DEFAULT_GOAL := clean

.PHONY: fmt vet build clean
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
clean: build
	rm -f hello_world
