// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	PostRegister(ctx context.Context, in *PostRegisterRequest, opts ...grpc.CallOption) (*PostRegisterResponse, error)
	PostLogin(ctx context.Context, in *PostLoginRequest, opts ...grpc.CallOption) (*PostLoginResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) PostRegister(ctx context.Context, in *PostRegisterRequest, opts ...grpc.CallOption) (*PostRegisterResponse, error) {
	out := new(PostRegisterResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostLogin(ctx context.Context, in *PostLoginRequest, opts ...grpc.CallOption) (*PostLoginResponse, error) {
	out := new(PostLoginResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	PostRegister(context.Context, *PostRegisterRequest) (*PostRegisterResponse, error)
	PostLogin(context.Context, *PostLoginRequest) (*PostLoginResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) PostRegister(context.Context, *PostRegisterRequest) (*PostRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRegister not implemented")
}
func (UnimplementedApiServer) PostLogin(context.Context, *PostLoginRequest) (*PostLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostLogin not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_PostRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostRegister(ctx, req.(*PostRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostLogin(ctx, req.(*PostLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "citizenv.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostRegister",
			Handler:    _Api_PostRegister_Handler,
		},
		{
			MethodName: "PostLogin",
			Handler:    _Api_PostLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
