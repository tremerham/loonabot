// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// RunnerClient is the client API for Runner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	Update(ctx context.Context, opts ...grpc.CallOption) (Runner_UpdateClient, error)
}

type runnerClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerClient(cc grpc.ClientConnInterface) RunnerClient {
	return &runnerClient{cc}
}

func (c *runnerClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/parthpower.runner.v1.Runner/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Update(ctx context.Context, opts ...grpc.CallOption) (Runner_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Runner_ServiceDesc.Streams[0], "/parthpower.runner.v1.Runner/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &runnerUpdateClient{stream}
	return x, nil
}

type Runner_UpdateClient interface {
	Send(*UpdateRequest) error
	Recv() (*UpdateResponse, error)
	grpc.ClientStream
}

type runnerUpdateClient struct {
	grpc.ClientStream
}

func (x *runnerUpdateClient) Send(m *UpdateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runnerUpdateClient) Recv() (*UpdateResponse, error) {
	m := new(UpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RunnerServer is the server API for Runner service.
// All implementations must embed UnimplementedRunnerServer
// for forward compatibility
type RunnerServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
	Update(Runner_UpdateServer) error
	mustEmbedUnimplementedRunnerServer()
}

// UnimplementedRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedRunnerServer struct {
}

func (UnimplementedRunnerServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedRunnerServer) Update(Runner_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRunnerServer) mustEmbedUnimplementedRunnerServer() {}

// UnsafeRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerServer will
// result in compilation errors.
type UnsafeRunnerServer interface {
	mustEmbedUnimplementedRunnerServer()
}

func RegisterRunnerServer(s grpc.ServiceRegistrar, srv RunnerServer) {
	s.RegisterService(&Runner_ServiceDesc, srv)
}

func _Runner_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parthpower.runner.v1.Runner/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RunnerServer).Update(&runnerUpdateServer{stream})
}

type Runner_UpdateServer interface {
	Send(*UpdateResponse) error
	Recv() (*UpdateRequest, error)
	grpc.ServerStream
}

type runnerUpdateServer struct {
	grpc.ServerStream
}

func (x *runnerUpdateServer) Send(m *UpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runnerUpdateServer) Recv() (*UpdateRequest, error) {
	m := new(UpdateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Runner_ServiceDesc is the grpc.ServiceDesc for Runner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Runner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parthpower.runner.v1.Runner",
	HandlerType: (*RunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Runner_Run_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Update",
			Handler:       _Runner_Update_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "update.proto",
}
