.PHONY: build run clean

build:
	go build -o ai-done main.go

run:
	go run main.go

clean:
	rm -f ai-done
	rm -rf dist/

install:
	go build -o ai-done main.go
	mkdir -p ~/Applications
	mv ai-done ~/Applications/

.DEFAULT_GOAL := build
