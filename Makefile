# Make commands
.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy

.PHONY: generate
generate:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	getweather/getweather.proto

.PHONY: serve
serve:
	go run server/server.go

.PHONY: client
client:
	go run client/client.go