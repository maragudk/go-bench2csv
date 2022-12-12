.PHONY: benchmark
benchmark:
	go test -cpu 1,2,4,8 -bench .

.PHONY: build
build:
	go build ./cmd/bench2csv

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: install
install:
	go install ./cmd/bench2csv

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

