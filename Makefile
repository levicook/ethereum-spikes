.DEFAULT_GOAL := build

.PHONY: build
build: test lint install

.PHONY: install
install:
	@go install ./cmd/...

.PHONY: install-tools
install-tools:
	@go get ./...
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint: staticcheck vet

.PHONY: staticcheck
staticcheck:
	@staticcheck ./...

.PHONY: test
test:
	@go test ./...

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: vet
vet:
	@go vet ./...

.PHONY: watch
watch: install-tools
	@modd
