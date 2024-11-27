// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: modules/printer/proto/print_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VirtualPrinter_SubmitPrintJob_FullMethodName    = "/printer.VirtualPrinter/SubmitPrintJob"
	VirtualPrinter_GetJobStatus_FullMethodName      = "/printer.VirtualPrinter/GetJobStatus"
	VirtualPrinter_CancelPrintJob_FullMethodName    = "/printer.VirtualPrinter/CancelPrintJob"
	VirtualPrinter_MonitorPrintJob_FullMethodName   = "/printer.VirtualPrinter/MonitorPrintJob"
	VirtualPrinter_ListPrintJobs_FullMethodName     = "/printer.VirtualPrinter/ListPrintJobs"
	VirtualPrinter_ViewPrinterStatus_FullMethodName = "/printer.VirtualPrinter/ViewPrinterStatus"
)

// VirtualPrinterClient is the client API for VirtualPrinter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VirtualPrinterClient interface {
	SubmitPrintJob(ctx context.Context, in *PrintDocument, opts ...grpc.CallOption) (*PrintJob, error)
	GetJobStatus(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (*PrintJob, error)
	CancelPrintJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*PrintJob, error)
	MonitorPrintJob(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrintJob], error)
	ListPrintJobs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPrintJobsResponse, error)
	ViewPrinterStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PrinterStatus, error)
}

type virtualPrinterClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtualPrinterClient(cc grpc.ClientConnInterface) VirtualPrinterClient {
	return &virtualPrinterClient{cc}
}

func (c *virtualPrinterClient) SubmitPrintJob(ctx context.Context, in *PrintDocument, opts ...grpc.CallOption) (*PrintJob, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrintJob)
	err := c.cc.Invoke(ctx, VirtualPrinter_SubmitPrintJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrinterClient) GetJobStatus(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (*PrintJob, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrintJob)
	err := c.cc.Invoke(ctx, VirtualPrinter_GetJobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrinterClient) CancelPrintJob(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*PrintJob, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrintJob)
	err := c.cc.Invoke(ctx, VirtualPrinter_CancelPrintJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrinterClient) MonitorPrintJob(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrintJob], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VirtualPrinter_ServiceDesc.Streams[0], VirtualPrinter_MonitorPrintJob_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetJobStatusRequest, PrintJob]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VirtualPrinter_MonitorPrintJobClient = grpc.ServerStreamingClient[PrintJob]

func (c *virtualPrinterClient) ListPrintJobs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListPrintJobsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPrintJobsResponse)
	err := c.cc.Invoke(ctx, VirtualPrinter_ListPrintJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrinterClient) ViewPrinterStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PrinterStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrinterStatus)
	err := c.cc.Invoke(ctx, VirtualPrinter_ViewPrinterStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualPrinterServer is the server API for VirtualPrinter service.
// All implementations must embed UnimplementedVirtualPrinterServer
// for forward compatibility.
type VirtualPrinterServer interface {
	SubmitPrintJob(context.Context, *PrintDocument) (*PrintJob, error)
	GetJobStatus(context.Context, *GetJobStatusRequest) (*PrintJob, error)
	CancelPrintJob(context.Context, *CancelJobRequest) (*PrintJob, error)
	MonitorPrintJob(*GetJobStatusRequest, grpc.ServerStreamingServer[PrintJob]) error
	ListPrintJobs(context.Context, *Empty) (*ListPrintJobsResponse, error)
	ViewPrinterStatus(context.Context, *Empty) (*PrinterStatus, error)
	mustEmbedUnimplementedVirtualPrinterServer()
}

// UnimplementedVirtualPrinterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVirtualPrinterServer struct{}

func (UnimplementedVirtualPrinterServer) SubmitPrintJob(context.Context, *PrintDocument) (*PrintJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPrintJob not implemented")
}
func (UnimplementedVirtualPrinterServer) GetJobStatus(context.Context, *GetJobStatusRequest) (*PrintJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobStatus not implemented")
}
func (UnimplementedVirtualPrinterServer) CancelPrintJob(context.Context, *CancelJobRequest) (*PrintJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPrintJob not implemented")
}
func (UnimplementedVirtualPrinterServer) MonitorPrintJob(*GetJobStatusRequest, grpc.ServerStreamingServer[PrintJob]) error {
	return status.Errorf(codes.Unimplemented, "method MonitorPrintJob not implemented")
}
func (UnimplementedVirtualPrinterServer) ListPrintJobs(context.Context, *Empty) (*ListPrintJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrintJobs not implemented")
}
func (UnimplementedVirtualPrinterServer) ViewPrinterStatus(context.Context, *Empty) (*PrinterStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewPrinterStatus not implemented")
}
func (UnimplementedVirtualPrinterServer) mustEmbedUnimplementedVirtualPrinterServer() {}
func (UnimplementedVirtualPrinterServer) testEmbeddedByValue()                        {}

// UnsafeVirtualPrinterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VirtualPrinterServer will
// result in compilation errors.
type UnsafeVirtualPrinterServer interface {
	mustEmbedUnimplementedVirtualPrinterServer()
}

func RegisterVirtualPrinterServer(s grpc.ServiceRegistrar, srv VirtualPrinterServer) {
	// If the following call pancis, it indicates UnimplementedVirtualPrinterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VirtualPrinter_ServiceDesc, srv)
}

func _VirtualPrinter_SubmitPrintJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintDocument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrinterServer).SubmitPrintJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualPrinter_SubmitPrintJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrinterServer).SubmitPrintJob(ctx, req.(*PrintDocument))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrinter_GetJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrinterServer).GetJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualPrinter_GetJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrinterServer).GetJobStatus(ctx, req.(*GetJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrinter_CancelPrintJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrinterServer).CancelPrintJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualPrinter_CancelPrintJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrinterServer).CancelPrintJob(ctx, req.(*CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrinter_MonitorPrintJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetJobStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VirtualPrinterServer).MonitorPrintJob(m, &grpc.GenericServerStream[GetJobStatusRequest, PrintJob]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VirtualPrinter_MonitorPrintJobServer = grpc.ServerStreamingServer[PrintJob]

func _VirtualPrinter_ListPrintJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrinterServer).ListPrintJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualPrinter_ListPrintJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrinterServer).ListPrintJobs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrinter_ViewPrinterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrinterServer).ViewPrinterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualPrinter_ViewPrinterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrinterServer).ViewPrinterStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VirtualPrinter_ServiceDesc is the grpc.ServiceDesc for VirtualPrinter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VirtualPrinter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "printer.VirtualPrinter",
	HandlerType: (*VirtualPrinterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPrintJob",
			Handler:    _VirtualPrinter_SubmitPrintJob_Handler,
		},
		{
			MethodName: "GetJobStatus",
			Handler:    _VirtualPrinter_GetJobStatus_Handler,
		},
		{
			MethodName: "CancelPrintJob",
			Handler:    _VirtualPrinter_CancelPrintJob_Handler,
		},
		{
			MethodName: "ListPrintJobs",
			Handler:    _VirtualPrinter_ListPrintJobs_Handler,
		},
		{
			MethodName: "ViewPrinterStatus",
			Handler:    _VirtualPrinter_ViewPrinterStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorPrintJob",
			Handler:       _VirtualPrinter_MonitorPrintJob_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "modules/printer/proto/print_service.proto",
}
