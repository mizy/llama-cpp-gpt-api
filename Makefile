DOCKER_REGISTRY ?= localhost:5000
DOCKER_REPO ?= ${DOCKER_REGISTRY}/vesoft
IMAGE_TAG ?= latest

export GO111MODULE := on
GOOS := $(if $(GOOS),$(GOOS),linux)
GOARCH := $(if $(GOARCH),$(GOARCH),amd64)
GOENV  := CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH)
GO     := $(GOENV) go
GO_BUILD := $(GO) build -trimpath
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GitHASH := $(shell git describe --no-match --always --dirty)
GITREF := $(shell git rev-parse --abbrev-ref HEAD)

REPO := github.com/mizy/llama-cpp-gpt-api

all: build

gen: gen-api

gen-api:
	rm -rf internal/handler
	$(GOBIN)/goctl api format -dir ./restapi
	$(GOBIN)/goctl api go --api ./restapi/index.api --dir .  --home ./.goctl

go-generate: $(GOBIN)/mockgen
	go generate ./...

check: tidy fmt vet imports lint

tidy:
	go mod tidy

fmt: $(GOBIN)/gofumpt
	# go fmt ./...
	$(GOBIN)/gofumpt -w -l ./

vet:
	go vet ./...

imports: $(GOBIN)/goimports $(GOBIN)/impi
	$(GOBIN)/impi --local github.com/vesoft-inc --scheme stdLocalThirdParty \
	    --skip handler/*.go --skip model/*.go \
	    -ignore-generated ./... \
	    || exit 1

lint: $(GOBIN)/golangci-lint
	$(GOBIN)/golangci-lint run

build:
	$(GO_BUILD) -ldflags '$(LDFLAGS)' -o bin/llama-cpp-gpt-api gpt.go

run:
	go build -ldflags '$(LDFLAGS)' -o bin/llama-cpp-gpt-api gpt.go
	bin/llama-cpp-gpt-api

test:
	go test ./...
 
install-goctl:
	GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@v1.4.0

tools: $(GOBIN)/goimports \
	$(GOBIN)/impi \
	$(GOBIN)/gofumpt \
	$(GOBIN)/golangci-lint \
	$(GOBIN)/controller-gen \
	$(GOBIN)/mockgen \
	$(GOBIN)/goctl

$(GOBIN)/goimports:
	go install golang.org/x/tools/cmd/goimports@v0.1.12

$(GOBIN)/impi:
	go install github.com/pavius/impi/cmd/impi@v0.0.3

$(GOBIN)/gofumpt:
	go install mvdan.cc/gofumpt@v0.3.1

$(GOBIN)/golangci-lint:
	@[ -f $(GOBIN)/golangci-lint ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.42.0 ;\
	}

$(GOBIN)/mockgen:
	go install github.com/golang/mock/mockgen@v1.6.0