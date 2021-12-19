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
	PostAdminDiv(ctx context.Context, in *PostAdminDivRequest, opts ...grpc.CallOption) (*PostAdminDivResponse, error)
	GetAdminDiv(ctx context.Context, in *GetAdminDivRequest, opts ...grpc.CallOption) (*GetAdminDivResponse, error)
	PostCitizen(ctx context.Context, in *PostCitizenRequest, opts ...grpc.CallOption) (*PostCitizenResponse, error)
	GetCitizen(ctx context.Context, in *GetCitizenRequest, opts ...grpc.CallOption) (*GetCitizenResponse, error)
	GetOneCitizen(ctx context.Context, in *GetOneCitizenRequest, opts ...grpc.CallOption) (*GetOneCitizenResponse, error)
	GetOneAdminDiv(ctx context.Context, in *GetOneAdminDivRequest, opts ...grpc.CallOption) (*GetOneAdminDivResponse, error)
	PutOneCitizen(ctx context.Context, in *PutOneCitizenRequest, opts ...grpc.CallOption) (*PutOneCitizenResponse, error)
	PutOneAdminDiv(ctx context.Context, in *PutOneAdminDivRequest, opts ...grpc.CallOption) (*PutOneAdminDivResponse, error)
	PostUserIssue(ctx context.Context, in *PostUserIssueRequest, opts ...grpc.CallOption) (*PostUserIssueResponse, error)
	PostAuthPassword(ctx context.Context, in *PostAuthPasswordRequest, opts ...grpc.CallOption) (*PostAuthPasswordResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	PostUserBan(ctx context.Context, in *PostUserActiveRequest, opts ...grpc.CallOption) (*PostUserActiveResponse, error)
	PostUserUnban(ctx context.Context, in *PostUserActiveRequest, opts ...grpc.CallOption) (*PostUserActiveResponse, error)
	DeleteCitizen(ctx context.Context, in *DeleteCitizenRequest, opts ...grpc.CallOption) (*DeleteCitizenResponse, error)
	GetAdminDivOptions(ctx context.Context, in *GetAdminDivOptionsRequest, opts ...grpc.CallOption) (*GetAdminDivOptionsResponse, error)
	GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error)
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

func (c *apiClient) PostAdminDiv(ctx context.Context, in *PostAdminDivRequest, opts ...grpc.CallOption) (*PostAdminDivResponse, error) {
	out := new(PostAdminDivResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostAdminDiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetAdminDiv(ctx context.Context, in *GetAdminDivRequest, opts ...grpc.CallOption) (*GetAdminDivResponse, error) {
	out := new(GetAdminDivResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetAdminDiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostCitizen(ctx context.Context, in *PostCitizenRequest, opts ...grpc.CallOption) (*PostCitizenResponse, error) {
	out := new(PostCitizenResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostCitizen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetCitizen(ctx context.Context, in *GetCitizenRequest, opts ...grpc.CallOption) (*GetCitizenResponse, error) {
	out := new(GetCitizenResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetCitizen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOneCitizen(ctx context.Context, in *GetOneCitizenRequest, opts ...grpc.CallOption) (*GetOneCitizenResponse, error) {
	out := new(GetOneCitizenResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetOneCitizen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOneAdminDiv(ctx context.Context, in *GetOneAdminDivRequest, opts ...grpc.CallOption) (*GetOneAdminDivResponse, error) {
	out := new(GetOneAdminDivResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetOneAdminDiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PutOneCitizen(ctx context.Context, in *PutOneCitizenRequest, opts ...grpc.CallOption) (*PutOneCitizenResponse, error) {
	out := new(PutOneCitizenResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PutOneCitizen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PutOneAdminDiv(ctx context.Context, in *PutOneAdminDivRequest, opts ...grpc.CallOption) (*PutOneAdminDivResponse, error) {
	out := new(PutOneAdminDivResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PutOneAdminDiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostUserIssue(ctx context.Context, in *PostUserIssueRequest, opts ...grpc.CallOption) (*PostUserIssueResponse, error) {
	out := new(PostUserIssueResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostUserIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostAuthPassword(ctx context.Context, in *PostAuthPasswordRequest, opts ...grpc.CallOption) (*PostAuthPasswordResponse, error) {
	out := new(PostAuthPasswordResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostAuthPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostUserBan(ctx context.Context, in *PostUserActiveRequest, opts ...grpc.CallOption) (*PostUserActiveResponse, error) {
	out := new(PostUserActiveResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostUserBan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostUserUnban(ctx context.Context, in *PostUserActiveRequest, opts ...grpc.CallOption) (*PostUserActiveResponse, error) {
	out := new(PostUserActiveResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/PostUserUnban", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteCitizen(ctx context.Context, in *DeleteCitizenRequest, opts ...grpc.CallOption) (*DeleteCitizenResponse, error) {
	out := new(DeleteCitizenResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/DeleteCitizen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetAdminDivOptions(ctx context.Context, in *GetAdminDivOptionsRequest, opts ...grpc.CallOption) (*GetAdminDivOptionsResponse, error) {
	out := new(GetAdminDivOptionsResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetAdminDivOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error) {
	out := new(GetAuthResponse)
	err := c.cc.Invoke(ctx, "/citizenv.Api/GetAuth", in, out, opts...)
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
	PostAdminDiv(context.Context, *PostAdminDivRequest) (*PostAdminDivResponse, error)
	GetAdminDiv(context.Context, *GetAdminDivRequest) (*GetAdminDivResponse, error)
	PostCitizen(context.Context, *PostCitizenRequest) (*PostCitizenResponse, error)
	GetCitizen(context.Context, *GetCitizenRequest) (*GetCitizenResponse, error)
	GetOneCitizen(context.Context, *GetOneCitizenRequest) (*GetOneCitizenResponse, error)
	GetOneAdminDiv(context.Context, *GetOneAdminDivRequest) (*GetOneAdminDivResponse, error)
	PutOneCitizen(context.Context, *PutOneCitizenRequest) (*PutOneCitizenResponse, error)
	PutOneAdminDiv(context.Context, *PutOneAdminDivRequest) (*PutOneAdminDivResponse, error)
	PostUserIssue(context.Context, *PostUserIssueRequest) (*PostUserIssueResponse, error)
	PostAuthPassword(context.Context, *PostAuthPasswordRequest) (*PostAuthPasswordResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	PostUserBan(context.Context, *PostUserActiveRequest) (*PostUserActiveResponse, error)
	PostUserUnban(context.Context, *PostUserActiveRequest) (*PostUserActiveResponse, error)
	DeleteCitizen(context.Context, *DeleteCitizenRequest) (*DeleteCitizenResponse, error)
	GetAdminDivOptions(context.Context, *GetAdminDivOptionsRequest) (*GetAdminDivOptionsResponse, error)
	GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error)
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
func (UnimplementedApiServer) PostAdminDiv(context.Context, *PostAdminDivRequest) (*PostAdminDivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAdminDiv not implemented")
}
func (UnimplementedApiServer) GetAdminDiv(context.Context, *GetAdminDivRequest) (*GetAdminDivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminDiv not implemented")
}
func (UnimplementedApiServer) PostCitizen(context.Context, *PostCitizenRequest) (*PostCitizenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCitizen not implemented")
}
func (UnimplementedApiServer) GetCitizen(context.Context, *GetCitizenRequest) (*GetCitizenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCitizen not implemented")
}
func (UnimplementedApiServer) GetOneCitizen(context.Context, *GetOneCitizenRequest) (*GetOneCitizenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneCitizen not implemented")
}
func (UnimplementedApiServer) GetOneAdminDiv(context.Context, *GetOneAdminDivRequest) (*GetOneAdminDivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneAdminDiv not implemented")
}
func (UnimplementedApiServer) PutOneCitizen(context.Context, *PutOneCitizenRequest) (*PutOneCitizenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutOneCitizen not implemented")
}
func (UnimplementedApiServer) PutOneAdminDiv(context.Context, *PutOneAdminDivRequest) (*PutOneAdminDivResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutOneAdminDiv not implemented")
}
func (UnimplementedApiServer) PostUserIssue(context.Context, *PostUserIssueRequest) (*PostUserIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserIssue not implemented")
}
func (UnimplementedApiServer) PostAuthPassword(context.Context, *PostAuthPasswordRequest) (*PostAuthPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAuthPassword not implemented")
}
func (UnimplementedApiServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedApiServer) PostUserBan(context.Context, *PostUserActiveRequest) (*PostUserActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserBan not implemented")
}
func (UnimplementedApiServer) PostUserUnban(context.Context, *PostUserActiveRequest) (*PostUserActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserUnban not implemented")
}
func (UnimplementedApiServer) DeleteCitizen(context.Context, *DeleteCitizenRequest) (*DeleteCitizenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCitizen not implemented")
}
func (UnimplementedApiServer) GetAdminDivOptions(context.Context, *GetAdminDivOptionsRequest) (*GetAdminDivOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminDivOptions not implemented")
}
func (UnimplementedApiServer) GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
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

func _Api_PostAdminDiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAdminDivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostAdminDiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostAdminDiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostAdminDiv(ctx, req.(*PostAdminDivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetAdminDiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminDivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetAdminDiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetAdminDiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetAdminDiv(ctx, req.(*GetAdminDivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostCitizen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCitizenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostCitizen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostCitizen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostCitizen(ctx, req.(*PostCitizenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetCitizen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCitizenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetCitizen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetCitizen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetCitizen(ctx, req.(*GetCitizenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOneCitizen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneCitizenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOneCitizen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetOneCitizen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOneCitizen(ctx, req.(*GetOneCitizenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOneAdminDiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneAdminDivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOneAdminDiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetOneAdminDiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOneAdminDiv(ctx, req.(*GetOneAdminDivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PutOneCitizen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutOneCitizenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PutOneCitizen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PutOneCitizen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PutOneCitizen(ctx, req.(*PutOneCitizenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PutOneAdminDiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutOneAdminDivRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PutOneAdminDiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PutOneAdminDiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PutOneAdminDiv(ctx, req.(*PutOneAdminDivRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostUserIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostUserIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostUserIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostUserIssue(ctx, req.(*PostUserIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostAuthPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAuthPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostAuthPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostAuthPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostAuthPassword(ctx, req.(*PostAuthPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostUserBan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostUserBan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostUserBan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostUserBan(ctx, req.(*PostUserActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostUserUnban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostUserUnban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/PostUserUnban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostUserUnban(ctx, req.(*PostUserActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteCitizen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCitizenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteCitizen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/DeleteCitizen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteCitizen(ctx, req.(*DeleteCitizenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetAdminDivOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminDivOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetAdminDivOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetAdminDivOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetAdminDivOptions(ctx, req.(*GetAdminDivOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citizenv.Api/GetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetAuth(ctx, req.(*GetAuthRequest))
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
		{
			MethodName: "PostAdminDiv",
			Handler:    _Api_PostAdminDiv_Handler,
		},
		{
			MethodName: "GetAdminDiv",
			Handler:    _Api_GetAdminDiv_Handler,
		},
		{
			MethodName: "PostCitizen",
			Handler:    _Api_PostCitizen_Handler,
		},
		{
			MethodName: "GetCitizen",
			Handler:    _Api_GetCitizen_Handler,
		},
		{
			MethodName: "GetOneCitizen",
			Handler:    _Api_GetOneCitizen_Handler,
		},
		{
			MethodName: "GetOneAdminDiv",
			Handler:    _Api_GetOneAdminDiv_Handler,
		},
		{
			MethodName: "PutOneCitizen",
			Handler:    _Api_PutOneCitizen_Handler,
		},
		{
			MethodName: "PutOneAdminDiv",
			Handler:    _Api_PutOneAdminDiv_Handler,
		},
		{
			MethodName: "PostUserIssue",
			Handler:    _Api_PostUserIssue_Handler,
		},
		{
			MethodName: "PostAuthPassword",
			Handler:    _Api_PostAuthPassword_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Api_GetUsers_Handler,
		},
		{
			MethodName: "PostUserBan",
			Handler:    _Api_PostUserBan_Handler,
		},
		{
			MethodName: "PostUserUnban",
			Handler:    _Api_PostUserUnban_Handler,
		},
		{
			MethodName: "DeleteCitizen",
			Handler:    _Api_DeleteCitizen_Handler,
		},
		{
			MethodName: "GetAdminDivOptions",
			Handler:    _Api_GetAdminDivOptions_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _Api_GetAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
