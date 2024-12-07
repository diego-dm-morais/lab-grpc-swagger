# struct
```
project/
    ├── cmd/                     # Entrypoints para os binários (ex: main.go)
    │   └── service/
    │       └── main.go          # Arquivo principal para iniciar o serviço
    ├── internal/                # Código privado ao módulo (não deve ser importado fora)
    │   ├── service/             # Implementação da lógica de negócio
    │   │   ├── handlers.go
    │   │   ├── repository.go
    │   │   └── service.go
    │   └── utils/               # Utilitários internos (helpers, validações, etc.)
    ├── pkg/                     # Pacotes compartilháveis fora do módulo
    ├── api/                     # Definições de API (protocol buffers, OpenAPI, etc.)
    │   └── proto/
    │       ├── service.proto    # Arquivo proto para definir gRPC
    │       ├── service.pb.go    # Gerado automaticamente (gRPC)
    │       └── service.pb.gw.go # Gerado automaticamente (gRPC-Gateway)
    ├── configs/                 # Configurações (ex: arquivos .yaml, .env)
    │   └── config.yaml
    ├── docs/                    # Documentação do projeto (Swagger, etc.)
    ├── migrations/              # Scripts de migração de banco de dados
    ├── build/                   # Arquivos de CI/CD, Dockerfiles, etc.
    │   └── Dockerfile
    ├── scripts/                 # Scripts de automação (bash, make, etc.)
    ├── go.mod                   # Dependências do módulo
    ├── go.sum                   # Hashes de dependências
```

# Command Golang
## Add module
```
go get {MODULE}@{VERSION} 
```

## Synchronization
```
go mod tidy
```

## Download module
```
go mod download
```
## Update module if exist
```
go get -u {MODULE}@{VERSION} 
```

