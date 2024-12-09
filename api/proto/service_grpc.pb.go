// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.10
// source: service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CRUDService_CreateItem_FullMethodName = "/service.CRUDService/CreateItem"
	CRUDService_GetItem_FullMethodName    = "/service.CRUDService/GetItem"
	CRUDService_UpdateItem_FullMethodName = "/service.CRUDService/UpdateItem"
	CRUDService_DeleteItem_FullMethodName = "/service.CRUDService/DeleteItem"
)

// CRUDServiceClient is the client API for CRUDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
}

type cRUDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDServiceClient(cc grpc.ClientConnInterface) CRUDServiceClient {
	return &cRUDServiceClient{cc}
}

func (c *cRUDServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, CRUDService_CreateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, CRUDService_GetItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, CRUDService_UpdateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, CRUDService_DeleteItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServiceServer is the server API for CRUDService service.
// All implementations must embed UnimplementedCRUDServiceServer
// for forward compatibility.
type CRUDServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	mustEmbedUnimplementedCRUDServiceServer()
}

// UnimplementedCRUDServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCRUDServiceServer struct{}

func (UnimplementedCRUDServiceServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedCRUDServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedCRUDServiceServer) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedCRUDServiceServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedCRUDServiceServer) mustEmbedUnimplementedCRUDServiceServer() {}
func (UnimplementedCRUDServiceServer) testEmbeddedByValue()                     {}

// UnsafeCRUDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServiceServer will
// result in compilation errors.
type UnsafeCRUDServiceServer interface {
	mustEmbedUnimplementedCRUDServiceServer()
}

func RegisterCRUDServiceServer(s grpc.ServiceRegistrar, srv CRUDServiceServer) {
	// If the following call pancis, it indicates UnimplementedCRUDServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CRUDService_ServiceDesc, srv)
}

func _CRUDService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_GetItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_UpdateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUDService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUDService_DeleteItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUDService_ServiceDesc is the grpc.ServiceDesc for CRUDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.CRUDService",
	HandlerType: (*CRUDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _CRUDService_CreateItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _CRUDService_GetItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _CRUDService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _CRUDService_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
