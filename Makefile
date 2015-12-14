all: build

build:
	mkdir -p bin/
	go build -o bin/ddp-server ./cmd/ddp-server/
	go build -o bin/ddp-client ./cmd/ddp-client/
