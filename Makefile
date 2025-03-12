.PHONY: run
run:
	go run cmd/*

.PHONY: build
build:
	go build -o app cmd/*

.PHONY: env
env:
	docker compose up -d

.PHONY: lint
lint:
	golangci-lint run --skip-dirs "go/pkg/mod" --skip-dirs "opt" ./... -c  .golangci.yml  -v

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix --skip-dirs "go/pkg/mod" --skip-dirs "opt" ./... -c  .golangci.yml  -v

.PHONY: test
test:
	go test -v ./...