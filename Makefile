.PHONY: build run clean test vet fmt lint help

APP         = hurrah
VERSION     = $(shell git describe --tags --abbrev=0)
GO          = go
GO_BUILD    = $(GO) build
GO_RUN      = $(GO) run
GO_FORMAT   = $(GO) fmt
GOFMT       = gofmt
GO_LIST     = $(GO) list
GO_TEST     = $(GO) test -v
GO_TOOL     = $(GO) tool
GO_VET      = $(GO) vet
GO_DEP      = $(GO) mod
GOOS        = ""
GOARCH      = ""
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT))
GO_LDFLAGS  = -ldflags '-X github.com/nao1215/hurrah/config.Version=${VERSION}'

build:  ## Build binary
	env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) $(GO_LDFLAGS) -o $(APP) cmd/hurrah/main.go

run: ## Run game
	env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_RUN) $(GO_LDFLAGS) cmd/hurrah/main.go

clean: ## Clean project
	-rm -rf $(APP) cover.out cover.html dist

test: ## Start test
	env GOOS=$(GOOS) $(GO_TEST) -cover $(GO_PKGROOT) -coverprofile=cover.out
	$(GO_TOOL) cover -html=cover.out -o cover.html

vet: ## Start go vet
	$(GO_VET) $(GO_PACKAGES)

fmt: ## Format go source code 
	$(GO_FORMAT) $(GO_PKGROOT)

lint: ## Lint go source code
	golangci-lint run --config .golangci.yml

.DEFAULT_GOAL := help
help:  
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'
