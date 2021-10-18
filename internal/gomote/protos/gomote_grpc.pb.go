// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// GomoteServiceClient is the client API for GomoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GomoteServiceClient interface {
	// Authenticate provides authentication information without any additonal action.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// CreateInstance creates a gomote instance.
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	// DestroyInstance destroys a gomote instance.
	DestroyInstance(ctx context.Context, in *DestroyInstanceRequest, opts ...grpc.CallOption) (*DestroyInstanceResponse, error)
	// ExecuteCommand executes a command on the gomote instance.
	ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (GomoteService_ExecuteCommandClient, error)
	// InstanceAlive gives the liveness state of a gomote instance.
	InstanceAlive(ctx context.Context, in *InstanceAliveRequest, opts ...grpc.CallOption) (*InstanceAliveResponse, error)
	// ListDirectory lists the contents of a directory on an gomote instance.
	ListDirectory(ctx context.Context, in *ListDirectoryRequest, opts ...grpc.CallOption) (*ListDirectoryResponse, error)
	// ListInstances lists all of the live gomote instances owned by the caller.
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// ReadTGZ tars and zips a dicrectory which exists on the gomote instance.
	ReadTGZ(ctx context.Context, in *ReadTGZRequest, opts ...grpc.CallOption) (GomoteService_ReadTGZClient, error)
	// RemoveDirectory removes a directory from the gomote instance.
	RemoveDirectory(ctx context.Context, in *RemoveDirectoryRequest, opts ...grpc.CallOption) (*RemoveDirectoryResponse, error)
	// RetrieveSSHCredentials retrieves the SSH credentials for the specified gomote instance.
	RetrieveSSHCredentials(ctx context.Context, in *RetrieveSSHCredentialsRequest, opts ...grpc.CallOption) (*RetrieveSSHCredentialsResponse, error)
	// WriteTGZ expands a tar and ziped file onto the file system of a gomote instance.
	WriteTGZ(ctx context.Context, opts ...grpc.CallOption) (GomoteService_WriteTGZClient, error)
}

type gomoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGomoteServiceClient(cc grpc.ClientConnInterface) GomoteServiceClient {
	return &gomoteServiceClient{cc}
}

func (c *gomoteServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/CreateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) DestroyInstance(ctx context.Context, in *DestroyInstanceRequest, opts ...grpc.CallOption) (*DestroyInstanceResponse, error) {
	out := new(DestroyInstanceResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/DestroyInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (GomoteService_ExecuteCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &GomoteService_ServiceDesc.Streams[0], "/protos.GomoteService/ExecuteCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &gomoteServiceExecuteCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GomoteService_ExecuteCommandClient interface {
	Recv() (*ExecuteCommandResponse, error)
	grpc.ClientStream
}

type gomoteServiceExecuteCommandClient struct {
	grpc.ClientStream
}

func (x *gomoteServiceExecuteCommandClient) Recv() (*ExecuteCommandResponse, error) {
	m := new(ExecuteCommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gomoteServiceClient) InstanceAlive(ctx context.Context, in *InstanceAliveRequest, opts ...grpc.CallOption) (*InstanceAliveResponse, error) {
	out := new(InstanceAliveResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/InstanceAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) ListDirectory(ctx context.Context, in *ListDirectoryRequest, opts ...grpc.CallOption) (*ListDirectoryResponse, error) {
	out := new(ListDirectoryResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/ListDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) ReadTGZ(ctx context.Context, in *ReadTGZRequest, opts ...grpc.CallOption) (GomoteService_ReadTGZClient, error) {
	stream, err := c.cc.NewStream(ctx, &GomoteService_ServiceDesc.Streams[1], "/protos.GomoteService/ReadTGZ", opts...)
	if err != nil {
		return nil, err
	}
	x := &gomoteServiceReadTGZClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GomoteService_ReadTGZClient interface {
	Recv() (*ReadTGZResponse, error)
	grpc.ClientStream
}

type gomoteServiceReadTGZClient struct {
	grpc.ClientStream
}

func (x *gomoteServiceReadTGZClient) Recv() (*ReadTGZResponse, error) {
	m := new(ReadTGZResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gomoteServiceClient) RemoveDirectory(ctx context.Context, in *RemoveDirectoryRequest, opts ...grpc.CallOption) (*RemoveDirectoryResponse, error) {
	out := new(RemoveDirectoryResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/RemoveDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) RetrieveSSHCredentials(ctx context.Context, in *RetrieveSSHCredentialsRequest, opts ...grpc.CallOption) (*RetrieveSSHCredentialsResponse, error) {
	out := new(RetrieveSSHCredentialsResponse)
	err := c.cc.Invoke(ctx, "/protos.GomoteService/RetrieveSSHCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gomoteServiceClient) WriteTGZ(ctx context.Context, opts ...grpc.CallOption) (GomoteService_WriteTGZClient, error) {
	stream, err := c.cc.NewStream(ctx, &GomoteService_ServiceDesc.Streams[2], "/protos.GomoteService/WriteTGZ", opts...)
	if err != nil {
		return nil, err
	}
	x := &gomoteServiceWriteTGZClient{stream}
	return x, nil
}

type GomoteService_WriteTGZClient interface {
	Send(*WriteTGZRequest) error
	CloseAndRecv() (*WriteTGZResponse, error)
	grpc.ClientStream
}

type gomoteServiceWriteTGZClient struct {
	grpc.ClientStream
}

func (x *gomoteServiceWriteTGZClient) Send(m *WriteTGZRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gomoteServiceWriteTGZClient) CloseAndRecv() (*WriteTGZResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteTGZResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GomoteServiceServer is the server API for GomoteService service.
// All implementations must embed UnimplementedGomoteServiceServer
// for forward compatibility
type GomoteServiceServer interface {
	// Authenticate provides authentication information without any additonal action.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// CreateInstance creates a gomote instance.
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	// DestroyInstance destroys a gomote instance.
	DestroyInstance(context.Context, *DestroyInstanceRequest) (*DestroyInstanceResponse, error)
	// ExecuteCommand executes a command on the gomote instance.
	ExecuteCommand(*ExecuteCommandRequest, GomoteService_ExecuteCommandServer) error
	// InstanceAlive gives the liveness state of a gomote instance.
	InstanceAlive(context.Context, *InstanceAliveRequest) (*InstanceAliveResponse, error)
	// ListDirectory lists the contents of a directory on an gomote instance.
	ListDirectory(context.Context, *ListDirectoryRequest) (*ListDirectoryResponse, error)
	// ListInstances lists all of the live gomote instances owned by the caller.
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// ReadTGZ tars and zips a dicrectory which exists on the gomote instance.
	ReadTGZ(*ReadTGZRequest, GomoteService_ReadTGZServer) error
	// RemoveDirectory removes a directory from the gomote instance.
	RemoveDirectory(context.Context, *RemoveDirectoryRequest) (*RemoveDirectoryResponse, error)
	// RetrieveSSHCredentials retrieves the SSH credentials for the specified gomote instance.
	RetrieveSSHCredentials(context.Context, *RetrieveSSHCredentialsRequest) (*RetrieveSSHCredentialsResponse, error)
	// WriteTGZ expands a tar and ziped file onto the file system of a gomote instance.
	WriteTGZ(GomoteService_WriteTGZServer) error
	mustEmbedUnimplementedGomoteServiceServer()
}

// UnimplementedGomoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGomoteServiceServer struct {
}

func (UnimplementedGomoteServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedGomoteServiceServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedGomoteServiceServer) DestroyInstance(context.Context, *DestroyInstanceRequest) (*DestroyInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyInstance not implemented")
}
func (UnimplementedGomoteServiceServer) ExecuteCommand(*ExecuteCommandRequest, GomoteService_ExecuteCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedGomoteServiceServer) InstanceAlive(context.Context, *InstanceAliveRequest) (*InstanceAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstanceAlive not implemented")
}
func (UnimplementedGomoteServiceServer) ListDirectory(context.Context, *ListDirectoryRequest) (*ListDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDirectory not implemented")
}
func (UnimplementedGomoteServiceServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedGomoteServiceServer) ReadTGZ(*ReadTGZRequest, GomoteService_ReadTGZServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadTGZ not implemented")
}
func (UnimplementedGomoteServiceServer) RemoveDirectory(context.Context, *RemoveDirectoryRequest) (*RemoveDirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDirectory not implemented")
}
func (UnimplementedGomoteServiceServer) RetrieveSSHCredentials(context.Context, *RetrieveSSHCredentialsRequest) (*RetrieveSSHCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveSSHCredentials not implemented")
}
func (UnimplementedGomoteServiceServer) WriteTGZ(GomoteService_WriteTGZServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteTGZ not implemented")
}
func (UnimplementedGomoteServiceServer) mustEmbedUnimplementedGomoteServiceServer() {}

// UnsafeGomoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GomoteServiceServer will
// result in compilation errors.
type UnsafeGomoteServiceServer interface {
	mustEmbedUnimplementedGomoteServiceServer()
}

func RegisterGomoteServiceServer(s grpc.ServiceRegistrar, srv GomoteServiceServer) {
	s.RegisterService(&GomoteService_ServiceDesc, srv)
}

func _GomoteService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_DestroyInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).DestroyInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/DestroyInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).DestroyInstance(ctx, req.(*DestroyInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_ExecuteCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteCommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GomoteServiceServer).ExecuteCommand(m, &gomoteServiceExecuteCommandServer{stream})
}

type GomoteService_ExecuteCommandServer interface {
	Send(*ExecuteCommandResponse) error
	grpc.ServerStream
}

type gomoteServiceExecuteCommandServer struct {
	grpc.ServerStream
}

func (x *gomoteServiceExecuteCommandServer) Send(m *ExecuteCommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GomoteService_InstanceAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).InstanceAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/InstanceAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).InstanceAlive(ctx, req.(*InstanceAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_ListDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).ListDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/ListDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).ListDirectory(ctx, req.(*ListDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_ReadTGZ_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadTGZRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GomoteServiceServer).ReadTGZ(m, &gomoteServiceReadTGZServer{stream})
}

type GomoteService_ReadTGZServer interface {
	Send(*ReadTGZResponse) error
	grpc.ServerStream
}

type gomoteServiceReadTGZServer struct {
	grpc.ServerStream
}

func (x *gomoteServiceReadTGZServer) Send(m *ReadTGZResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GomoteService_RemoveDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).RemoveDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/RemoveDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).RemoveDirectory(ctx, req.(*RemoveDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_RetrieveSSHCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveSSHCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GomoteServiceServer).RetrieveSSHCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GomoteService/RetrieveSSHCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GomoteServiceServer).RetrieveSSHCredentials(ctx, req.(*RetrieveSSHCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GomoteService_WriteTGZ_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GomoteServiceServer).WriteTGZ(&gomoteServiceWriteTGZServer{stream})
}

type GomoteService_WriteTGZServer interface {
	SendAndClose(*WriteTGZResponse) error
	Recv() (*WriteTGZRequest, error)
	grpc.ServerStream
}

type gomoteServiceWriteTGZServer struct {
	grpc.ServerStream
}

func (x *gomoteServiceWriteTGZServer) SendAndClose(m *WriteTGZResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gomoteServiceWriteTGZServer) Recv() (*WriteTGZRequest, error) {
	m := new(WriteTGZRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GomoteService_ServiceDesc is the grpc.ServiceDesc for GomoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GomoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GomoteService",
	HandlerType: (*GomoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _GomoteService_Authenticate_Handler,
		},
		{
			MethodName: "CreateInstance",
			Handler:    _GomoteService_CreateInstance_Handler,
		},
		{
			MethodName: "DestroyInstance",
			Handler:    _GomoteService_DestroyInstance_Handler,
		},
		{
			MethodName: "InstanceAlive",
			Handler:    _GomoteService_InstanceAlive_Handler,
		},
		{
			MethodName: "ListDirectory",
			Handler:    _GomoteService_ListDirectory_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _GomoteService_ListInstances_Handler,
		},
		{
			MethodName: "RemoveDirectory",
			Handler:    _GomoteService_RemoveDirectory_Handler,
		},
		{
			MethodName: "RetrieveSSHCredentials",
			Handler:    _GomoteService_RetrieveSSHCredentials_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteCommand",
			Handler:       _GomoteService_ExecuteCommand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReadTGZ",
			Handler:       _GomoteService_ReadTGZ_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WriteTGZ",
			Handler:       _GomoteService_WriteTGZ_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "gomote.proto",
}
