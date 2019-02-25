## Copyright 2019 Squeeze Authors
##
## Licensed under the Apache License, Version 2.0 (the "License");
## you may not use this file except in compliance with the License.
## You may obtain a copy of the License at
##
##     http://www.apache.org/licenses/LICENSE-2.0
##
## Unless required by applicable law or agreed to in writing, software
## distributed under the License is distributed on an "AS IS" BASIS,
## WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
## See the License for the specific language governing permissions and
## limitations under the License.

BINARY = squeeze
GO_BIN := go
VERSION := "0.1.0"
GIT_TAG ?= $(shell git rev-parse --short HEAD)
VERSION_FLAGS = "-X github.com/agile6v/squeeze/pkg/version.version=${VERSION} \
 		   -X github.com/agile6v/squeeze/pkg/version.gitRevision=$(GIT_TAG)"

export GO111MODULE=on

.PHONY: all
all: protoc build

.PHONY: protoc
protoc:
	@echo "### Generating Go files"
	cd pkg/pb && protoc --go_out=plugins=grpc:. *.proto

.PHONY: build
build:
	@echo "### Building binary"
	$(GO_BIN) build -ldflags ${VERSION_FLAGS} -o $(BINARY)

.PHONY: fmt
fmt:
	@echo "### Formatting project"
	$(GO_BIN) fmt ./...


.PHONY: clean
clean:
	@echo "### Cleaning ..."
	rm -f $(BINARY)
