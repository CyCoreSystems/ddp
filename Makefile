all: build

build:
	go build ./cmd/ddpd/
	go build ./cmd/ddpc/
