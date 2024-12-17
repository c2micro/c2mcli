BIN_DIR=$(PWD)/bin
C2M_DIR=$(PWD)/cmd/c2mcli
CC=gcc
CXX=g++
VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "0.0.0")
BUILD=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags="-s -w -X github.com/c2micro/c2mcli/internal/version.gitCommit=${BUILD} -X github.com/c2micro/c2mcli/internal/version.gitVersion=${VERSION}"

.PHONY: c2mcli
c2mcli:
	@mkdir -p ${BIN_DIR}
	@echo "Building operator cli..."
	CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build ${LDFLAGS} -o ${BIN_DIR}/c2mcli ${C2M_DIR}
	@strip bin/c2mcli

.PHONY: go-sync
go-sync:
	@go mod tidy && go mod vendor

.PHONY: dep-shared
dep-shared:
	@echo "Update shared components..."
	@export GOPRIVATE="github.com/c2micro" && go get -u github.com/c2micro/c2mshr/ && go mod tidy && go mod vendor

.PHONY: dep-mlan
dep-mlan:
	@echo "Update mlan components..."
	@export GOPRIVATE="github.com/c2micro" && go get -u github.com/c2micro/mlan/ && go mod tidy && go mod vendor
