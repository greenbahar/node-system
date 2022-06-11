proto-compile:
	protoc --go-grpc_out=/home/greenbahar/go/src -I=./proto/node --go_out=/home/greenbahar/go/src node.proto

grpc-client:
	go run ./cmd/grpc-client/main.go

grpc-server:
	go run ./cmd/grpc-server/main.go

sync-engine:
	go run ./cmd/syncEngine/main.go