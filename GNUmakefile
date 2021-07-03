TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

build: fmtcheck generate
	go install

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

generate:
	go generate ./...

gencheck:
	@echo "==> Generating..."
	@make generate
	@echo "==> Comparing generated code to committed code..."
	@git diff --compact-summary --exit-code -- ./ || \
    		(echo; echo "Unexpected difference in generated code. Run 'go generate' to update the generated code and commit."; exit 1)

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --compact-summary --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)

test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

lint:
	@golangci-lint run ./...

.PHONY: build fmtcheck test testacc lint generate gencheck depscheck

