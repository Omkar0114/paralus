// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/system/organization.proto

package systemv3

import (
	context "context"
	v3 "github.com/paralus/paralus/proto/types/systempb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationServiceClient is the client API for OrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationServiceClient interface {
	CreateOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error)
	GetOrganizations(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.OrganizationList, error)
	GetOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error)
	UpdateOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error)
	DeleteOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error)
}

type organizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationServiceClient(cc grpc.ClientConnInterface) OrganizationServiceClient {
	return &organizationServiceClient{cc}
}

func (c *organizationServiceClient) CreateOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error) {
	out := new(v3.Organization)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.system.v3.OrganizationService/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) GetOrganizations(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.OrganizationList, error) {
	out := new(v3.OrganizationList)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.system.v3.OrganizationService/GetOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) GetOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error) {
	out := new(v3.Organization)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.system.v3.OrganizationService/GetOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) UpdateOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error) {
	out := new(v3.Organization)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.system.v3.OrganizationService/UpdateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationServiceClient) DeleteOrganization(ctx context.Context, in *v3.Organization, opts ...grpc.CallOption) (*v3.Organization, error) {
	out := new(v3.Organization)
	err := c.cc.Invoke(ctx, "/paralus.dev.rpc.system.v3.OrganizationService/DeleteOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServiceServer is the server API for OrganizationService service.
// All implementations should embed UnimplementedOrganizationServiceServer
// for forward compatibility
type OrganizationServiceServer interface {
	CreateOrganization(context.Context, *v3.Organization) (*v3.Organization, error)
	GetOrganizations(context.Context, *v3.Organization) (*v3.OrganizationList, error)
	GetOrganization(context.Context, *v3.Organization) (*v3.Organization, error)
	UpdateOrganization(context.Context, *v3.Organization) (*v3.Organization, error)
	DeleteOrganization(context.Context, *v3.Organization) (*v3.Organization, error)
}

// UnimplementedOrganizationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationServiceServer struct {
}

func (UnimplementedOrganizationServiceServer) CreateOrganization(context.Context, *v3.Organization) (*v3.Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) GetOrganizations(context.Context, *v3.Organization) (*v3.OrganizationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizations not implemented")
}
func (UnimplementedOrganizationServiceServer) GetOrganization(context.Context, *v3.Organization) (*v3.Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) UpdateOrganization(context.Context, *v3.Organization) (*v3.Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationServiceServer) DeleteOrganization(context.Context, *v3.Organization) (*v3.Organization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganization not implemented")
}

// UnsafeOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServiceServer will
// result in compilation errors.
type UnsafeOrganizationServiceServer interface {
	mustEmbedUnimplementedOrganizationServiceServer()
}

func RegisterOrganizationServiceServer(s grpc.ServiceRegistrar, srv OrganizationServiceServer) {
	s.RegisterService(&OrganizationService_ServiceDesc, srv)
}

func _OrganizationService_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.system.v3.OrganizationService/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).CreateOrganization(ctx, req.(*v3.Organization))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_GetOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).GetOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.system.v3.OrganizationService/GetOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).GetOrganizations(ctx, req.(*v3.Organization))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_GetOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).GetOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.system.v3.OrganizationService/GetOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).GetOrganization(ctx, req.(*v3.Organization))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.system.v3.OrganizationService/UpdateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).UpdateOrganization(ctx, req.(*v3.Organization))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationService_DeleteOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.Organization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServiceServer).DeleteOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paralus.dev.rpc.system.v3.OrganizationService/DeleteOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServiceServer).DeleteOrganization(ctx, req.(*v3.Organization))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationService_ServiceDesc is the grpc.ServiceDesc for OrganizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "paralus.dev.rpc.system.v3.OrganizationService",
	HandlerType: (*OrganizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrganization",
			Handler:    _OrganizationService_CreateOrganization_Handler,
		},
		{
			MethodName: "GetOrganizations",
			Handler:    _OrganizationService_GetOrganizations_Handler,
		},
		{
			MethodName: "GetOrganization",
			Handler:    _OrganizationService_GetOrganization_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _OrganizationService_UpdateOrganization_Handler,
		},
		{
			MethodName: "DeleteOrganization",
			Handler:    _OrganizationService_DeleteOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/system/organization.proto",
}
