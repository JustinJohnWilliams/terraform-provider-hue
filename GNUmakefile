TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

build: fmtcheck generate
	go install

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

tools:
	@echo "==> installing required tooling..."
	go install github.com/bflad/tfproviderdocs@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/katbyte/terrafmt@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH || $$GOPATH)/bin v1.32.0

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

tflint:
	./scripts/run-tflint.sh

docscheck:
	@echo "==> Checking documentation spelling..."
	@misspell -error -source=text -i hdinsight,exportfs docs/
	@misspell -error -source text CHANGELOG.md
	@echo "==> Checking documentation for errors..."
	@tfproviderdocs check -require-resource-subcategory \
		-allowed-resource-subcategories-file docs/allowed-subcategories.txt

tffmtfix:
	@echo "==> Fixing acceptance test terraform blocks code with terrafmt..."
	@find hue | egrep "_test.go" | sort | while read f; do terrafmt fmt -f $$f; done
	@echo "==> Fixing docs terraform blocks code with terrafmt..."
	@find docs | egrep ".md" | sort | while read f; do terrafmt fmt $$f; done

.PHONY: build fmtcheck test testacc lint generate gencheck depscheck docscheck tflint
