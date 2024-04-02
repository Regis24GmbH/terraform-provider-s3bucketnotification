LIST_ALL := $(shell go list ./... | grep -v vendor | grep -v mocks)
# Force using Go Modules and always read the dependencies from the `vendor` folder.
# Also force building for hiveregis.
export GOFLAGS=-mod=vendor -tags=hiveregis

all: lint test

.PHONY: lint
lint: ## Lint the files
	@go fmt ${LIST_ALL}
	@golangci-lint version
	@golangci-lint run

.PHONY: test
test: ## Run unit tests
	@go test -short -count 1 -v ./...

.PHONY: race
race: ## Run data race detector
	@go test -race -short -count 1 -v ./...

.PHONY: coverage
coverage: ## Generate coverage report
	@go-acc ./...
	@go tool cover -func=coverage.txt

.PHONY: coverhtml
coverhtml: ## Generate coverage report as HTML
	@go-acc ./...
	@go tool cover -html=coverage.txt -o coverage.html

.PHONY: build
build: ## Build all binary files based on directory `./cmd/`
	@GOARCH=amd64 GOOS=darwin go build -tags lambda.norpc -o ./bin/darwin/amd64/terraform-provider-s3bucketnotification ; \
    GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o ./bin/linux/amd64/terraform-provider-s3bucketnotification ; \
    GOARCH=amd64 GOOS=windows go build -tags lambda.norpc -o ./bin/windows/amd64/terraform-provider-s3bucketnotification ;

.PHONY: install_locally
install_locally: ## Move all built files to the local terraform installation
	 @mkdir -p ~/.terraform.d/plugins/darwin_amd64 ; \
 	 mkdir -p ~/.terraform.d/plugins/linux_amd64 ; \
 	 mkdir -p ~/.terraform.d/plugins/windows_amd64 ; \
     mv ./bin/darwin/amd64/terraform-provider-s3bucketnotification ~/.terraform.d/plugins/darwin_amd64/; \
     mv ./bin/linux/amd64/terraform-provider-s3bucketnotification ~/.terraform.d/plugins/linux_amd64/; \
     mv ./bin/windows/amd64/terraform-provider-s3bucketnotification ~/.terraform.d/plugins/windows_amd64/;


.PHONY: upgrade
upgrade: ## Upgrade the dependencies
	@go get -u -t ./...
	@go mod tidy
	@go mod vendor

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
