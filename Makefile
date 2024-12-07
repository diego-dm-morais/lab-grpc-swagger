SERVICE_DIR=cmd/service

setup:
	@chmod +x ./script/setup.sh
	@./script/setup.sh

.PHONY: protoc
protoc:
	@protoc -I ./proto \
			--go_out ./proto --go_opt paths=source_relative \
			--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
			--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
			./proto/servicing/service.proto

install:
	@go mod download

clean-mod:
	@go mod tidy
	

.PHONY: run
run:
	@go run $(SERVICE_DIR)/main.go