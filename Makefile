all: build

build:
	mkdir -p bin/
	go build -o bin/ddpd ./cmd/ddpd/
	go build -o bin/ddpc ./cmd/ddpc/
