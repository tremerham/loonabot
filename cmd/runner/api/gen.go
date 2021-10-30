package api

//go:generate protoc -I. --go_out=module=github.com/parthpower/loonabot/cmd/runner/api:. --go-grpc_out=module=github.com/parthpower/loonabot/cmd/runner/api:. update.proto
