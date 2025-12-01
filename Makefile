.PHONY: build

LINTER_VERSION=v2.6.2

default: unit-tests

get-linter:
	command -v golangci-lint || curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)

## lint: download/install golangci-lint and analyse the source code with the configuration in .golangci.yml
lint: get-linter
	golangci-lint run --timeout=10m

unit-tests: lint
	go fmt ./...
	go test -vet all -shuffle=on ./...


mod:
	go mod vendor -v

tidy:
	go mod tidy -v