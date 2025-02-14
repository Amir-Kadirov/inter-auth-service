// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: seller-auth.proto

package auth_service

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
	SellerAuth_SellerLoginByPassword_FullMethodName       = "/auth_service.SellerAuth/SellerLoginByPassword"
	SellerAuth_SellerGmailCheck_FullMethodName            = "/auth_service.SellerAuth/SellerGmailCheck"
	SellerAuth_SellerRegisterByMail_FullMethodName        = "/auth_service.SellerAuth/SellerRegisterByMail"
	SellerAuth_SellerRegisterByMailConfirm_FullMethodName = "/auth_service.SellerAuth/SellerRegisterByMailConfirm"
	SellerAuth_SellerCreate_FullMethodName                = "/auth_service.SellerAuth/SellerCreate"
	SellerAuth_SellerLoginByGmail_FullMethodName          = "/auth_service.SellerAuth/SellerLoginByGmail"
	SellerAuth_SellerLoginByGmailComfirm_FullMethodName   = "/auth_service.SellerAuth/SellerLoginByGmailComfirm"
	SellerAuth_SellerUpdatePassword_FullMethodName        = "/auth_service.SellerAuth/SellerUpdatePassword"
	SellerAuth_SellerResetPassword_FullMethodName         = "/auth_service.SellerAuth/SellerResetPassword"
	SellerAuth_SellerResetPasswordConfirm_FullMethodName  = "/auth_service.SellerAuth/SellerResetPasswordConfirm"
)

// SellerAuthClient is the client API for SellerAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SellerAuthClient interface {
	SellerLoginByPassword(ctx context.Context, in *SellerLoginRequest, opts ...grpc.CallOption) (*SellerLoginResponse, error)
	SellerGmailCheck(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerGmailCheckResponse, error)
	SellerRegisterByMail(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerEmpty, error)
	SellerRegisterByMailConfirm(ctx context.Context, in *SellerRConfirm, opts ...grpc.CallOption) (*SellerEmpty, error)
	SellerCreate(ctx context.Context, in *SellerCreateRequest, opts ...grpc.CallOption) (*SellerEmpty, error)
	SellerLoginByGmail(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerEmpty, error)
	SellerLoginByGmailComfirm(ctx context.Context, in *SellerLoginByGmailRequest, opts ...grpc.CallOption) (*SellerLoginResponse, error)
	SellerUpdatePassword(ctx context.Context, in *SellerCreateRequest, opts ...grpc.CallOption) (*SellerEmpty, error)
	SellerResetPassword(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerEmpty, error)
	SellerResetPasswordConfirm(ctx context.Context, in *SellerPasswordConfirm, opts ...grpc.CallOption) (*SellerEmpty, error)
}

type sellerAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewSellerAuthClient(cc grpc.ClientConnInterface) SellerAuthClient {
	return &sellerAuthClient{cc}
}

func (c *sellerAuthClient) SellerLoginByPassword(ctx context.Context, in *SellerLoginRequest, opts ...grpc.CallOption) (*SellerLoginResponse, error) {
	out := new(SellerLoginResponse)
	err := c.cc.Invoke(ctx, SellerAuth_SellerLoginByPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerGmailCheck(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerGmailCheckResponse, error) {
	out := new(SellerGmailCheckResponse)
	err := c.cc.Invoke(ctx, SellerAuth_SellerGmailCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerRegisterByMail(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerRegisterByMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerRegisterByMailConfirm(ctx context.Context, in *SellerRConfirm, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerRegisterByMailConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerCreate(ctx context.Context, in *SellerCreateRequest, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerLoginByGmail(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerLoginByGmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerLoginByGmailComfirm(ctx context.Context, in *SellerLoginByGmailRequest, opts ...grpc.CallOption) (*SellerLoginResponse, error) {
	out := new(SellerLoginResponse)
	err := c.cc.Invoke(ctx, SellerAuth_SellerLoginByGmailComfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerUpdatePassword(ctx context.Context, in *SellerCreateRequest, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerUpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerResetPassword(ctx context.Context, in *SellerGmailCheckRequest, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerAuthClient) SellerResetPasswordConfirm(ctx context.Context, in *SellerPasswordConfirm, opts ...grpc.CallOption) (*SellerEmpty, error) {
	out := new(SellerEmpty)
	err := c.cc.Invoke(ctx, SellerAuth_SellerResetPasswordConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SellerAuthServer is the server API for SellerAuth service.
// All implementations should embed UnimplementedSellerAuthServer
// for forward compatibility
type SellerAuthServer interface {
	SellerLoginByPassword(context.Context, *SellerLoginRequest) (*SellerLoginResponse, error)
	SellerGmailCheck(context.Context, *SellerGmailCheckRequest) (*SellerGmailCheckResponse, error)
	SellerRegisterByMail(context.Context, *SellerGmailCheckRequest) (*SellerEmpty, error)
	SellerRegisterByMailConfirm(context.Context, *SellerRConfirm) (*SellerEmpty, error)
	SellerCreate(context.Context, *SellerCreateRequest) (*SellerEmpty, error)
	SellerLoginByGmail(context.Context, *SellerGmailCheckRequest) (*SellerEmpty, error)
	SellerLoginByGmailComfirm(context.Context, *SellerLoginByGmailRequest) (*SellerLoginResponse, error)
	SellerUpdatePassword(context.Context, *SellerCreateRequest) (*SellerEmpty, error)
	SellerResetPassword(context.Context, *SellerGmailCheckRequest) (*SellerEmpty, error)
	SellerResetPasswordConfirm(context.Context, *SellerPasswordConfirm) (*SellerEmpty, error)
}

// UnimplementedSellerAuthServer should be embedded to have forward compatible implementations.
type UnimplementedSellerAuthServer struct {
}

func (UnimplementedSellerAuthServer) SellerLoginByPassword(context.Context, *SellerLoginRequest) (*SellerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerLoginByPassword not implemented")
}
func (UnimplementedSellerAuthServer) SellerGmailCheck(context.Context, *SellerGmailCheckRequest) (*SellerGmailCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerGmailCheck not implemented")
}
func (UnimplementedSellerAuthServer) SellerRegisterByMail(context.Context, *SellerGmailCheckRequest) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerRegisterByMail not implemented")
}
func (UnimplementedSellerAuthServer) SellerRegisterByMailConfirm(context.Context, *SellerRConfirm) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerRegisterByMailConfirm not implemented")
}
func (UnimplementedSellerAuthServer) SellerCreate(context.Context, *SellerCreateRequest) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerCreate not implemented")
}
func (UnimplementedSellerAuthServer) SellerLoginByGmail(context.Context, *SellerGmailCheckRequest) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerLoginByGmail not implemented")
}
func (UnimplementedSellerAuthServer) SellerLoginByGmailComfirm(context.Context, *SellerLoginByGmailRequest) (*SellerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerLoginByGmailComfirm not implemented")
}
func (UnimplementedSellerAuthServer) SellerUpdatePassword(context.Context, *SellerCreateRequest) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerUpdatePassword not implemented")
}
func (UnimplementedSellerAuthServer) SellerResetPassword(context.Context, *SellerGmailCheckRequest) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerResetPassword not implemented")
}
func (UnimplementedSellerAuthServer) SellerResetPasswordConfirm(context.Context, *SellerPasswordConfirm) (*SellerEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellerResetPasswordConfirm not implemented")
}

// UnsafeSellerAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SellerAuthServer will
// result in compilation errors.
type UnsafeSellerAuthServer interface {
	mustEmbedUnimplementedSellerAuthServer()
}

func RegisterSellerAuthServer(s grpc.ServiceRegistrar, srv SellerAuthServer) {
	s.RegisterService(&SellerAuth_ServiceDesc, srv)
}

func _SellerAuth_SellerLoginByPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerLoginByPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerLoginByPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerLoginByPassword(ctx, req.(*SellerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerGmailCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerGmailCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerGmailCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerGmailCheck(ctx, req.(*SellerGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerRegisterByMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerRegisterByMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerRegisterByMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerRegisterByMail(ctx, req.(*SellerGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerRegisterByMailConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerRConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerRegisterByMailConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerRegisterByMailConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerRegisterByMailConfirm(ctx, req.(*SellerRConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerCreate(ctx, req.(*SellerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerLoginByGmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerLoginByGmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerLoginByGmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerLoginByGmail(ctx, req.(*SellerGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerLoginByGmailComfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerLoginByGmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerLoginByGmailComfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerLoginByGmailComfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerLoginByGmailComfirm(ctx, req.(*SellerLoginByGmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerUpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerUpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerUpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerUpdatePassword(ctx, req.(*SellerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerGmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerResetPassword(ctx, req.(*SellerGmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerAuth_SellerResetPasswordConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellerPasswordConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerAuthServer).SellerResetPasswordConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerAuth_SellerResetPasswordConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerAuthServer).SellerResetPasswordConfirm(ctx, req.(*SellerPasswordConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

// SellerAuth_ServiceDesc is the grpc.ServiceDesc for SellerAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SellerAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.SellerAuth",
	HandlerType: (*SellerAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SellerLoginByPassword",
			Handler:    _SellerAuth_SellerLoginByPassword_Handler,
		},
		{
			MethodName: "SellerGmailCheck",
			Handler:    _SellerAuth_SellerGmailCheck_Handler,
		},
		{
			MethodName: "SellerRegisterByMail",
			Handler:    _SellerAuth_SellerRegisterByMail_Handler,
		},
		{
			MethodName: "SellerRegisterByMailConfirm",
			Handler:    _SellerAuth_SellerRegisterByMailConfirm_Handler,
		},
		{
			MethodName: "SellerCreate",
			Handler:    _SellerAuth_SellerCreate_Handler,
		},
		{
			MethodName: "SellerLoginByGmail",
			Handler:    _SellerAuth_SellerLoginByGmail_Handler,
		},
		{
			MethodName: "SellerLoginByGmailComfirm",
			Handler:    _SellerAuth_SellerLoginByGmailComfirm_Handler,
		},
		{
			MethodName: "SellerUpdatePassword",
			Handler:    _SellerAuth_SellerUpdatePassword_Handler,
		},
		{
			MethodName: "SellerResetPassword",
			Handler:    _SellerAuth_SellerResetPassword_Handler,
		},
		{
			MethodName: "SellerResetPasswordConfirm",
			Handler:    _SellerAuth_SellerResetPasswordConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seller-auth.proto",
}
