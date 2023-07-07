
TEST?=$$(go list ./...)
VERSION=2.0.0-alpha.1

default: build

tools:
	go generate -tags tools tools/tools.go

build: fmtcheck
	go install -ldflags="-X github.com/pingidentity/pingctl/cmd/version/version.version=$(VERSION)"

test: fmtcheck
	go test $(TEST) $(TESTARGS) -timeout=5m

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)

vet:
	@echo "==> Running go vet..."
	@go vet ./... ; if [ $$? -ne 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

golangci-lint:
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./...

lint: golangci-lint
	@./scripts/lint-all.sh

gosec:
	@gosec -exclude-generated ./...

devcheck: build vet lint gosec test testacc
	
.PHONY: tools build test testacc depscheck codecheck lint golangci-lint codegen fmtcheck generate securitycheck devcheck