SERVICE_DIR=cmd/service

setup:
	@chmod +x ./script/setup.sh
	@./script/setup.sh


.PHONY: protoc
protoc:
	@protoc -I=./api/proto \
			--go_out=./api/proto --go_opt=paths=source_relative \
			--go-grpc_out=./api/proto --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=./api/proto --grpc-gateway_opt=paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			./api/proto/service.proto


.PHONY: clean-swagger
clean-swagger:
	@rm -f ./api/swagger/service.swagger.json

swagger: clean-swagger
	@protoc -I=./api/proto \
			--plugin=protoc-gen-swagger=$(GOPATH)/bin/protoc-gen-swagger \
			--openapiv2_out ./api/swagger \
			./api/proto/service.proto

clean-mod:
	@go mod tidy

install:
	@go mod download
	

.PHONY: run
run:
	@go run $(SERVICE_DIR)/main.go


format:
	@find . -name "*.go" -exec gofmt -w {} \;
