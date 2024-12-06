setup:
	@chmod +x ./setup.sh
	@./setup.sh

.PHONY: proto-buf

clean-module:
	@go mod tidy
	
install:
	@go mod download


proto-buf:
	@buf generate