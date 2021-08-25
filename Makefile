.PHONY: all
all: build run clean

.PHONY: build
build:
	go build ./cmd/foundry/

.PHONY: run
run:
	./foundry gen -f ./test/spec.yaml

.PHONY: clean
clean:
	rm ./foundry
