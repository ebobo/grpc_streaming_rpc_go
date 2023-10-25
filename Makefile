all: lint gen build

lint:
	@buf lint

gen: clean
	@buf generate

clean:
	@rm -rf pkg

init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest

build:
	@go mod tidy
	@cd cmd/server && go build -o ../../bin/grpc_server
	@cd cmd/client && go build -o ../../bin/stream_client

