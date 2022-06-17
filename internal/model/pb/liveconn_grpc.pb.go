// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: liveconn.proto

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

// ConnectedLiveServiceClient is the client API for ConnectedLiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectedLiveServiceClient interface {
	PushReportBatch(ctx context.Context, in *ReportBatchRequest, opts ...grpc.CallOption) (*ReportBatchACK, error)
	PushMatrixBatch(ctx context.Context, in *MatrixBatchRequest, opts ...grpc.CallOption) (*MatrixBatchACK, error)
	GetMatrixBatch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type connectedLiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectedLiveServiceClient(cc grpc.ClientConnInterface) ConnectedLiveServiceClient {
	return &connectedLiveServiceClient{cc}
}

func (c *connectedLiveServiceClient) PushReportBatch(ctx context.Context, in *ReportBatchRequest, opts ...grpc.CallOption) (*ReportBatchACK, error) {
	out := new(ReportBatchACK)
	err := c.cc.Invoke(ctx, "/ConnectedLiveService/PushReportBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectedLiveServiceClient) PushMatrixBatch(ctx context.Context, in *MatrixBatchRequest, opts ...grpc.CallOption) (*MatrixBatchACK, error) {
	out := new(MatrixBatchACK)
	err := c.cc.Invoke(ctx, "/ConnectedLiveService/PushMatrixBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectedLiveServiceClient) GetMatrixBatch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ConnectedLiveService/GetMatrixBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectedLiveServiceServer is the server API for ConnectedLiveService service.
// All implementations must embed UnimplementedConnectedLiveServiceServer
// for forward compatibility
type ConnectedLiveServiceServer interface {
	PushReportBatch(context.Context, *ReportBatchRequest) (*ReportBatchACK, error)
	PushMatrixBatch(context.Context, *MatrixBatchRequest) (*MatrixBatchACK, error)
	GetMatrixBatch(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedConnectedLiveServiceServer()
}

// UnimplementedConnectedLiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectedLiveServiceServer struct {
}

func (UnimplementedConnectedLiveServiceServer) PushReportBatch(context.Context, *ReportBatchRequest) (*ReportBatchACK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushReportBatch not implemented")
}
func (UnimplementedConnectedLiveServiceServer) PushMatrixBatch(context.Context, *MatrixBatchRequest) (*MatrixBatchACK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMatrixBatch not implemented")
}
func (UnimplementedConnectedLiveServiceServer) GetMatrixBatch(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatrixBatch not implemented")
}
func (UnimplementedConnectedLiveServiceServer) mustEmbedUnimplementedConnectedLiveServiceServer() {}

// UnsafeConnectedLiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectedLiveServiceServer will
// result in compilation errors.
type UnsafeConnectedLiveServiceServer interface {
	mustEmbedUnimplementedConnectedLiveServiceServer()
}

func RegisterConnectedLiveServiceServer(s grpc.ServiceRegistrar, srv ConnectedLiveServiceServer) {
	s.RegisterService(&ConnectedLiveService_ServiceDesc, srv)
}

func _ConnectedLiveService_PushReportBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectedLiveServiceServer).PushReportBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConnectedLiveService/PushReportBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectedLiveServiceServer).PushReportBatch(ctx, req.(*ReportBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectedLiveService_PushMatrixBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatrixBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectedLiveServiceServer).PushMatrixBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConnectedLiveService/PushMatrixBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectedLiveServiceServer).PushMatrixBatch(ctx, req.(*MatrixBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectedLiveService_GetMatrixBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectedLiveServiceServer).GetMatrixBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConnectedLiveService/GetMatrixBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectedLiveServiceServer).GetMatrixBatch(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectedLiveService_ServiceDesc is the grpc.ServiceDesc for ConnectedLiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectedLiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ConnectedLiveService",
	HandlerType: (*ConnectedLiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushReportBatch",
			Handler:    _ConnectedLiveService_PushReportBatch_Handler,
		},
		{
			MethodName: "PushMatrixBatch",
			Handler:    _ConnectedLiveService_PushMatrixBatch_Handler,
		},
		{
			MethodName: "GetMatrixBatch",
			Handler:    _ConnectedLiveService_GetMatrixBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liveconn.proto",
}
