// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: kvStore.proto

package kvStore

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

// KvStoreServiceClient is the client API for KvStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KvStoreServiceClient interface {
	Get(ctx context.Context, in *GetArgs, opts ...grpc.CallOption) (*GetReply, error)
	PutAppend(ctx context.Context, in *PutAppendArgs, opts ...grpc.CallOption) (*PutAppendReply, error)
}

type kvStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKvStoreServiceClient(cc grpc.ClientConnInterface) KvStoreServiceClient {
	return &kvStoreServiceClient{cc}
}

func (c *kvStoreServiceClient) Get(ctx context.Context, in *GetArgs, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/kvStore.KvStoreService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvStoreServiceClient) PutAppend(ctx context.Context, in *PutAppendArgs, opts ...grpc.CallOption) (*PutAppendReply, error) {
	out := new(PutAppendReply)
	err := c.cc.Invoke(ctx, "/kvStore.KvStoreService/PutAppend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KvStoreServiceServer is the server API for KvStoreService service.
// All implementations must embed UnimplementedKvStoreServiceServer
// for forward compatibility
type KvStoreServiceServer interface {
	Get(context.Context, *GetArgs) (*GetReply, error)
	PutAppend(context.Context, *PutAppendArgs) (*PutAppendReply, error)
	mustEmbedUnimplementedKvStoreServiceServer()
}

// UnimplementedKvStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKvStoreServiceServer struct {
}

func (UnimplementedKvStoreServiceServer) Get(context.Context, *GetArgs) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKvStoreServiceServer) PutAppend(context.Context, *PutAppendArgs) (*PutAppendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAppend not implemented")
}
func (UnimplementedKvStoreServiceServer) mustEmbedUnimplementedKvStoreServiceServer() {}

// UnsafeKvStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KvStoreServiceServer will
// result in compilation errors.
type UnsafeKvStoreServiceServer interface {
	mustEmbedUnimplementedKvStoreServiceServer()
}

func RegisterKvStoreServiceServer(s grpc.ServiceRegistrar, srv KvStoreServiceServer) {
	s.RegisterService(&KvStoreService_ServiceDesc, srv)
}

func _KvStoreService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvStore.KvStoreService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServiceServer).Get(ctx, req.(*GetArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KvStoreService_PutAppend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAppendArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServiceServer).PutAppend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvStore.KvStoreService/PutAppend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServiceServer).PutAppend(ctx, req.(*PutAppendArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// KvStoreService_ServiceDesc is the grpc.ServiceDesc for KvStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KvStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvStore.KvStoreService",
	HandlerType: (*KvStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KvStoreService_Get_Handler,
		},
		{
			MethodName: "PutAppend",
			Handler:    _KvStoreService_PutAppend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvStore.proto",
}
