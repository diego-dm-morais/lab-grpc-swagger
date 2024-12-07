SERVICE_DIR=cmd/service

setup:
	@chmod +x ./script/setup.sh
	@./script/setup.sh




.PHONY: protoc
protoc:
	@protoc -I=./proto \
			--go_out=./proto --go_opt=paths=source_relative \
			--go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=./proto --grpc-gateway_opt=paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			./proto/servicing/service.proto


.PHONY: clean-swagger
clean-swagger:
	@rm -f ./api/openapi/servicing/service.swagger.json

swagger: clean-swagger
	@protoc -I=./proto \
			--plugin=protoc-gen-swagger=$(GOPATH)/bin/protoc-gen-swagger \
			--openapiv2_out ./api/openapi \
			./proto/servicing/service.proto

clean-mod:
	@go mod tidy

install:
	@go mod download
	

.PHONY: run
run:
	@go run $(SERVICE_DIR)/main.go