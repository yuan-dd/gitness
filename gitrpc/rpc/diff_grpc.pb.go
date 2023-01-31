// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: diff.proto

package rpc

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

// DiffServiceClient is the client API for DiffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiffServiceClient interface {
	RawDiff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (DiffService_RawDiffClient, error)
	DiffShortStat(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffShortStatResponse, error)
}

type diffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiffServiceClient(cc grpc.ClientConnInterface) DiffServiceClient {
	return &diffServiceClient{cc}
}

func (c *diffServiceClient) RawDiff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (DiffService_RawDiffClient, error) {
	stream, err := c.cc.NewStream(ctx, &DiffService_ServiceDesc.Streams[0], "/rpc.DiffService/RawDiff", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffServiceRawDiffClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiffService_RawDiffClient interface {
	Recv() (*RawDiffResponse, error)
	grpc.ClientStream
}

type diffServiceRawDiffClient struct {
	grpc.ClientStream
}

func (x *diffServiceRawDiffClient) Recv() (*RawDiffResponse, error) {
	m := new(RawDiffResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diffServiceClient) DiffShortStat(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffShortStatResponse, error) {
	out := new(DiffShortStatResponse)
	err := c.cc.Invoke(ctx, "/rpc.DiffService/DiffShortStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServiceServer is the server API for DiffService service.
// All implementations must embed UnimplementedDiffServiceServer
// for forward compatibility
type DiffServiceServer interface {
	RawDiff(*DiffRequest, DiffService_RawDiffServer) error
	DiffShortStat(context.Context, *DiffRequest) (*DiffShortStatResponse, error)
	mustEmbedUnimplementedDiffServiceServer()
}

// UnimplementedDiffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiffServiceServer struct {
}

func (UnimplementedDiffServiceServer) RawDiff(*DiffRequest, DiffService_RawDiffServer) error {
	return status.Errorf(codes.Unimplemented, "method RawDiff not implemented")
}
func (UnimplementedDiffServiceServer) DiffShortStat(context.Context, *DiffRequest) (*DiffShortStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiffShortStat not implemented")
}
func (UnimplementedDiffServiceServer) mustEmbedUnimplementedDiffServiceServer() {}

// UnsafeDiffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiffServiceServer will
// result in compilation errors.
type UnsafeDiffServiceServer interface {
	mustEmbedUnimplementedDiffServiceServer()
}

func RegisterDiffServiceServer(s grpc.ServiceRegistrar, srv DiffServiceServer) {
	s.RegisterService(&DiffService_ServiceDesc, srv)
}

func _DiffService_RawDiff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServiceServer).RawDiff(m, &diffServiceRawDiffServer{stream})
}

type DiffService_RawDiffServer interface {
	Send(*RawDiffResponse) error
	grpc.ServerStream
}

type diffServiceRawDiffServer struct {
	grpc.ServerStream
}

func (x *diffServiceRawDiffServer) Send(m *RawDiffResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DiffService_DiffShortStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).DiffShortStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DiffService/DiffShortStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).DiffShortStat(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiffService_ServiceDesc is the grpc.ServiceDesc for DiffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DiffService",
	HandlerType: (*DiffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiffShortStat",
			Handler:    _DiffService_DiffShortStat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RawDiff",
			Handler:       _DiffService_RawDiff_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "diff.proto",
}
