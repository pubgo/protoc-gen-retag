.PHONY: install
install:
	go install .

vet:
	@go vet ./...
	gofumpt -l -w -extra .

generate:
	@go generate ./...

lint:
	@golangci-lint run --skip-dirs-use-default --timeout 3m0s
