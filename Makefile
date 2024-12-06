setup:
	@echo "Iniciando a configuração do ambiente..."
	@chmod +x ./setup.sh
	@./setup.sh
	@echo "Configuração do ambiente concluída!


.PHONY: proto-buf

clean-module:
	@go mod tidy
	
install:
	@go mod download


proto-buf:
	@buf generate