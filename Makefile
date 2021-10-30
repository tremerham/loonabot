GO ?= $(shell which go)
GO_FLAGS = -tags netgo -ldflags '-w -s'
GO_SRC = $(shell find -name '*.go')
EMBED_SRC = $(shell find cmd/loona/static)
PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip

all: loonabot

.PHONY: protoc
protoc:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$(PROTOC_ZIP)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)
	sudo chmod a+x /usr/local/bin/protoc

gen: cmd/runner/api/update.proto
	$(GO) list -f '{{ join .Imports "\n" }}' tools.go | xargs $(GO) get
	$(GO) generate ./...

grpc: server.go client.go $(GO_SRC) gen
	CGO_ENABLED=0 $(GO) build -v $(GO_FLAGS) -tags server server.go
	CGO_ENABLED=0 $(GO) build -v $(GO_FLAGS) -tags client client.go
	rm -rf ./cert/*.pem ./cert/*.srl

loonabot: $(GO_SRC)
	CGO_ENABLED=0 $(GO) build -v $(GO_FLAGS) -o $@

.PHONY: clean
deploy: loona
	./deploy.sh

.PHONY: clean
clean:
	rm -rf loona