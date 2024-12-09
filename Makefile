SERVICE_DIR=cmd/server

PROTO_DIR=./api/proto
SWAGGER_DIR=./api/swagger

setup:
	@chmod +x ./script/setup.sh
	@./script/setup.sh

install:
	@go mod download

clean-mod:
	@go mod tidy


.PHONY: clean-swagger
clean-swagger:
	@rm -rf $(SWAGGER_DIR)/service.swagger.json

.PHONY: protoc
protoc: clean-swagger
	@protoc -I=$(PROTO_DIR) \
			--go_out=$(PROTO_DIR) --go_opt=paths=source_relative \
			--go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=$(PROTO_DIR) --grpc-gateway_opt=paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			-I=$(PROTO_DIR)/protoc-gen-openapiv2/options/ \
			--openapiv2_out=logtostderr=true:$(SWAGGER_DIR) \
			$(PROTO_DIR)/service.proto \

format:
	@find . -name "*.go" -exec gofmt -w {} \;

lint:
	@golangci-lint run ./...

.PHONY: run
run:
	@go run $(SERVICE_DIR)/main.go

