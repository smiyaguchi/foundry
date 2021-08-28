.PHONY: all
all: build run

.PHONY: build
build:
	go build ./cmd/foundry/

.PHONY: run
run:
	./foundry gen -f ./test/spec.yaml
	./foundry gen -f ./test/spec.yaml -n 3
