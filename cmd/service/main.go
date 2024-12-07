package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	proto "github.com/diego-dm-morais/lab-grpc-swagger/proto/servicing"
	"github.com/diego-dm-morais/lab-grpc-swagger/internal/service"
)

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

	// Serve a documentação Swagger
	http.Handle("/swagger/", http.StripPrefix("/swagger", http.FileServer(http.Dir("./"))))
	log.Println("Servidor REST rodando em :8080")
	http.ListenAndServe(":8080", mux)
}