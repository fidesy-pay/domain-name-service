// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: api/domain-name-service/domain-name-service.proto

package domain_name_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DomainNameService_GetAddress_FullMethodName    = "/domain_name_service.DomainNameService/GetAddress"
	DomainNameService_UpdateAddress_FullMethodName = "/domain_name_service.DomainNameService/UpdateAddress"
)

// DomainNameServiceClient is the client API for DomainNameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainNameServiceClient interface {
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error)
}

type domainNameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainNameServiceClient(cc grpc.ClientConnInterface) DomainNameServiceClient {
	return &domainNameServiceClient{cc}
}

func (c *domainNameServiceClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, DomainNameService_GetAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainNameServiceClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error) {
	out := new(UpdateAddressResponse)
	err := c.cc.Invoke(ctx, DomainNameService_UpdateAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainNameServiceServer is the server API for DomainNameService service.
// All implementations must embed UnimplementedDomainNameServiceServer
// for forward compatibility
type DomainNameServiceServer interface {
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error)
	mustEmbedUnimplementedDomainNameServiceServer()
}

// UnimplementedDomainNameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDomainNameServiceServer struct {
}

func (UnimplementedDomainNameServiceServer) GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedDomainNameServiceServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedDomainNameServiceServer) mustEmbedUnimplementedDomainNameServiceServer() {}

// UnsafeDomainNameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainNameServiceServer will
// result in compilation errors.
type UnsafeDomainNameServiceServer interface {
	mustEmbedUnimplementedDomainNameServiceServer()
}

func RegisterDomainNameServiceServer(s grpc.ServiceRegistrar, srv DomainNameServiceServer) {
	s.RegisterService(&DomainNameService_ServiceDesc, srv)
}

func _DomainNameService_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainNameServiceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainNameService_GetAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainNameServiceServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainNameService_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainNameServiceServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainNameService_UpdateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainNameServiceServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainNameService_ServiceDesc is the grpc.ServiceDesc for DomainNameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainNameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain_name_service.DomainNameService",
	HandlerType: (*DomainNameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddress",
			Handler:    _DomainNameService_GetAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _DomainNameService_UpdateAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/domain-name-service/domain-name-service.proto",
}