// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: order_status_notes.proto

package order_notes

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
	OrderStatusNotes_Create_FullMethodName  = "/order_notes.OrderStatusNotes/Create"
	OrderStatusNotes_GetById_FullMethodName = "/order_notes.OrderStatusNotes/GetById"
	OrderStatusNotes_GetAll_FullMethodName  = "/order_notes.OrderStatusNotes/GetAll"
	OrderStatusNotes_Update_FullMethodName  = "/order_notes.OrderStatusNotes/Update"
	OrderStatusNotes_Delete_FullMethodName  = "/order_notes.OrderStatusNotes/Delete"
)

// OrderStatusNotesClient is the client API for OrderStatusNotes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderStatusNotesClient interface {
	Create(ctx context.Context, in *CreateOrderNotes, opts ...grpc.CallOption) (*OrderNotes, error)
	GetById(ctx context.Context, in *OrderNotesPrimaryKey, opts ...grpc.CallOption) (*OrderNotes, error)
	GetAll(ctx context.Context, in *GetListOrderNotesRequest, opts ...grpc.CallOption) (*GetListOrderNotesResponse, error)
	Update(ctx context.Context, in *UpdateOrderNotes, opts ...grpc.CallOption) (*OrderNotes, error)
	Delete(ctx context.Context, in *OrderNotesPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type orderStatusNotesClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderStatusNotesClient(cc grpc.ClientConnInterface) OrderStatusNotesClient {
	return &orderStatusNotesClient{cc}
}

func (c *orderStatusNotesClient) Create(ctx context.Context, in *CreateOrderNotes, opts ...grpc.CallOption) (*OrderNotes, error) {
	out := new(OrderNotes)
	err := c.cc.Invoke(ctx, OrderStatusNotes_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderStatusNotesClient) GetById(ctx context.Context, in *OrderNotesPrimaryKey, opts ...grpc.CallOption) (*OrderNotes, error) {
	out := new(OrderNotes)
	err := c.cc.Invoke(ctx, OrderStatusNotes_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderStatusNotesClient) GetAll(ctx context.Context, in *GetListOrderNotesRequest, opts ...grpc.CallOption) (*GetListOrderNotesResponse, error) {
	out := new(GetListOrderNotesResponse)
	err := c.cc.Invoke(ctx, OrderStatusNotes_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderStatusNotesClient) Update(ctx context.Context, in *UpdateOrderNotes, opts ...grpc.CallOption) (*OrderNotes, error) {
	out := new(OrderNotes)
	err := c.cc.Invoke(ctx, OrderStatusNotes_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderStatusNotesClient) Delete(ctx context.Context, in *OrderNotesPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OrderStatusNotes_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderStatusNotesServer is the server API for OrderStatusNotes service.
// All implementations should embed UnimplementedOrderStatusNotesServer
// for forward compatibility
type OrderStatusNotesServer interface {
	Create(context.Context, *CreateOrderNotes) (*OrderNotes, error)
	GetById(context.Context, *OrderNotesPrimaryKey) (*OrderNotes, error)
	GetAll(context.Context, *GetListOrderNotesRequest) (*GetListOrderNotesResponse, error)
	Update(context.Context, *UpdateOrderNotes) (*OrderNotes, error)
	Delete(context.Context, *OrderNotesPrimaryKey) (*Empty, error)
}

// UnimplementedOrderStatusNotesServer should be embedded to have forward compatible implementations.
type UnimplementedOrderStatusNotesServer struct {
}

func (UnimplementedOrderStatusNotesServer) Create(context.Context, *CreateOrderNotes) (*OrderNotes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderStatusNotesServer) GetById(context.Context, *OrderNotesPrimaryKey) (*OrderNotes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedOrderStatusNotesServer) GetAll(context.Context, *GetListOrderNotesRequest) (*GetListOrderNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedOrderStatusNotesServer) Update(context.Context, *UpdateOrderNotes) (*OrderNotes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderStatusNotesServer) Delete(context.Context, *OrderNotesPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeOrderStatusNotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderStatusNotesServer will
// result in compilation errors.
type UnsafeOrderStatusNotesServer interface {
	mustEmbedUnimplementedOrderStatusNotesServer()
}

func RegisterOrderStatusNotesServer(s grpc.ServiceRegistrar, srv OrderStatusNotesServer) {
	s.RegisterService(&OrderStatusNotes_ServiceDesc, srv)
}

func _OrderStatusNotes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderNotes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderStatusNotesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderStatusNotes_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderStatusNotesServer).Create(ctx, req.(*CreateOrderNotes))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderStatusNotes_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNotesPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderStatusNotesServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderStatusNotes_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderStatusNotesServer).GetById(ctx, req.(*OrderNotesPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderStatusNotes_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOrderNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderStatusNotesServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderStatusNotes_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderStatusNotesServer).GetAll(ctx, req.(*GetListOrderNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderStatusNotes_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderNotes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderStatusNotesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderStatusNotes_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderStatusNotesServer).Update(ctx, req.(*UpdateOrderNotes))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderStatusNotes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNotesPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderStatusNotesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderStatusNotes_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderStatusNotesServer).Delete(ctx, req.(*OrderNotesPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderStatusNotes_ServiceDesc is the grpc.ServiceDesc for OrderStatusNotes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderStatusNotes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_notes.OrderStatusNotes",
	HandlerType: (*OrderStatusNotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderStatusNotes_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _OrderStatusNotes_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _OrderStatusNotes_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderStatusNotes_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrderStatusNotes_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_status_notes.proto",
}
