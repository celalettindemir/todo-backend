all: build

test: ## Run unittests
	@go test ./controllers ./repository ./server ./service
pact: ## Run unittests
	@go test ./pact
coverage: ## Generate global code coverage report
	./tools/coverage.sh;
build: ## Build the binary file
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .