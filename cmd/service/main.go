package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/diego-dm-morais/lab-grpc-swagger/internal/service"
	proto "github.com/diego-dm-morais/lab-grpc-swagger/proto/servicing"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// Inicia o servidor gRPC
	grpcServer := grpc.NewServer()
	crudService := &service.CRUDService{}
	proto.RegisterCRUDServiceServer(grpcServer, crudService)

	// Porta para o gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar listener: %v", err)
	}

	go func() {
		log.Println("Servidor gRPC rodando em :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Falha ao iniciar gRPC server: %v", err)
		}
	}()

	// Inicia o servidor REST usando gRPC-Gateway
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := proto.RegisterCRUDServiceHandlerFromEndpoint(context.Background(), mux, ":50051", opts); err != nil {
		log.Fatalf("Falha ao iniciar gRPC-Gateway: %v", err)
	}

	// Serve o gRPC-Gateway
	http.Handle("/", mux)

	// Serve o Swagger UI
	http.Handle("/docs", http.StripPrefix("/docs", http.FileServer(http.Dir("./api/openapi/servicing/swagger-ui"))))

	// Serve o arquivo JSON gerado pelo protoc
	http.Handle("/api/openapi/servicing/service.swagger.json", http.StripPrefix("/api", http.FileServer(http.Dir("./api"))))

	log.Println("Servidor REST rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
