cat:
	cat Makefile

run:
	go run ./cmd/jsonpath

build:
	go build ./cmd/jsonpath

install:
	go install ./cmd/jsonpath
