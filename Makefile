.DEFAULT_GOAL := build

.PHONY: build
build: go-get go-test lint go-install

.PHONY: go-get
go-get:
	@go get ./...

.PHONY: install
go-install:
	@go install ./cmd/...

.PHONY: go-test
go-test:
	@go test ./...

.PHONY: go-vet
go-vet:
	@go vet ./...

.PHONY: install-tools
install-tools: go-get
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint: go-vet
	@staticcheck ./...

.PHONY: watch
watch: install-tools
	@modd
