package service

import (
	"context"

	proto "github.com/diego-dm-morais/lab-grpc-swagger/api/proto"
)

type CRUDService struct {
	proto.UnimplementedCRUDServiceServer
}

func (s *CRUDService) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	// Lógica para criar um item
	return &proto.CreateItemResponse{Id: "12345"}, nil
}
func (s *CRUDService) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	// Lógica para obter um item
	return &proto.GetItemResponse{Id: req.Id, Name: "Sample Item"}, nil
}
func (s *CRUDService) UpdateItem(ctx context.Context, req *proto.UpdateItemRequest) (*proto.UpdateItemResponse, error) {
	// Lógica para atualizar um item
	return &proto.UpdateItemResponse{Id: req.Id}, nil
}
func (s *CRUDService) DeleteItem(ctx context.Context, req *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	// Lógica para deletar um item
	return &proto.DeleteItemResponse{Status: "Deleted"}, nil
}
