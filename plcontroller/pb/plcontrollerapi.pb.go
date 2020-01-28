// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/NEU-SNS/ReverseTraceroute/plcontroller/pb/plcontrollerapi.proto

package pb

import (
	context "context"
	fmt "fmt"
	datamodel "github.com/NEU-SNS/revtrvp/datamodel"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/NEU-SNS/ReverseTraceroute/plcontroller/pb/plcontrollerapi.proto", fileDescriptor_f8b414d6dde05896)
}

var fileDescriptor_f8b414d6dde05896 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x55, 0xc4, 0xc3, 0x20, 0x04, 0x63, 0x41, 0xec, 0xa9, 0x3c, 0x75, 0x69, 0xd7, 0xea,
	0x10, 0x1d, 0x22, 0x4c, 0x22, 0x10, 0x91, 0x65, 0x2d, 0x0f, 0xdd, 0x66, 0xc6, 0xe7, 0x36, 0xb0,
	0xce, 0x9b, 0x66, 0xde, 0x0a, 0x7d, 0xd4, 0xbe, 0x4d, 0xb8, 0x9a, 0xae, 0x52, 0xd0, 0x1e, 0xf7,
	0xf7, 0x9f, 0xdf, 0xff, 0xbd, 0x7d, 0x6c, 0x98, 0x6a, 0x7a, 0xcf, 0x65, 0xa8, 0x70, 0x11, 0x8d,
	0x9f, 0x5e, 0x2f, 0x27, 0xe3, 0x49, 0x94, 0xc0, 0x12, 0x9c, 0x87, 0x17, 0x27, 0x14, 0x38, 0xcc,
	0x09, 0x22, 0x9b, 0x29, 0x34, 0xe4, 0x30, 0xcb, 0xc0, 0x45, 0x56, 0xee, 0x7d, 0x0b, 0xab, 0x43,
	0xeb, 0x90, 0x90, 0x37, 0xac, 0x0c, 0xee, 0xfe, 0xd5, 0x37, 0x13, 0x24, 0x16, 0x38, 0x83, 0x2c,
	0xb2, 0xda, 0xa4, 0x6b, 0x3d, 0x78, 0xa8, 0xa8, 0xd2, 0x16, 0x6e, 0x0a, 0xfa, 0x15, 0x0b, 0x96,
	0xc2, 0x90, 0x48, 0xc1, 0xa2, 0x36, 0xb4, 0xa9, 0xb8, 0xaf, 0x58, 0xe1, 0x40, 0x79, 0x8b, 0x38,
	0x5f, 0xeb, 0xd7, 0x5f, 0x0d, 0xd6, 0x8e, 0x47, 0x83, 0xed, 0x6d, 0xf8, 0x15, 0x6b, 0xc6, 0xda,
	0xa4, 0x9c, 0x87, 0x5b, 0x27, 0x5c, 0x81, 0xbe, 0x4b, 0x83, 0xa3, 0x03, 0xd6, 0xad, 0x5d, 0xd4,
	0x7b, 0x75, 0x3e, 0x60, 0x6c, 0x37, 0x8a, 0x9f, 0x96, 0x1e, 0xed, 0xf0, 0x4a, 0x3f, 0xf9, 0x35,
	0xd9, 0x94, 0x0c, 0x59, 0x3b, 0x01, 0x05, 0x7a, 0x09, 0x93, 0xd5, 0x7a, 0xbc, 0x53, 0x7a, 0x9c,
	0x80, 0x2a, 0x60, 0x70, 0x5e, 0x82, 0x63, 0x24, 0x3d, 0xff, 0xfc, 0x89, 0x12, 0xf0, 0x16, 0x8d,
	0x87, 0x6e, 0xad, 0x57, 0xe7, 0xb7, 0xac, 0xf5, 0x0c, 0x34, 0x8d, 0x3d, 0x3f, 0x2e, 0x09, 0xd3,
	0x38, 0x81, 0x8f, 0x1c, 0x3c, 0x05, 0x9d, 0x03, 0x4a, 0xb9, 0x33, 0x85, 0x38, 0x62, 0xed, 0xbe,
	0x52, 0x60, 0x29, 0x76, 0x28, 0xc1, 0xef, 0xfd, 0x4b, 0x31, 0x06, 0x66, 0xeb, 0x24, 0x38, 0xfb,
	0x2b, 0xd9, 0x2d, 0xf2, 0xd8, 0x7c, 0x6b, 0x58, 0x29, 0x5b, 0xc5, 0xa1, 0x6f, 0xbe, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x05, 0x2e, 0x0f, 0xb8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PLControllerClient is the client API for PLController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PLControllerClient interface {
	Ping(ctx context.Context, opts ...grpc.CallOption) (PLController_PingClient, error)
	Traceroute(ctx context.Context, opts ...grpc.CallOption) (PLController_TracerouteClient, error)
	ReceiveSpoof(ctx context.Context, in *datamodel.RecSpoof, opts ...grpc.CallOption) (PLController_ReceiveSpoofClient, error)
	GetVPs(ctx context.Context, in *datamodel.VPRequest, opts ...grpc.CallOption) (PLController_GetVPsClient, error)
	AcceptProbes(ctx context.Context, in *datamodel.SpoofedProbes, opts ...grpc.CallOption) (*datamodel.SpoofedProbesResponse, error)
}

type pLControllerClient struct {
	cc *grpc.ClientConn
}

func NewPLControllerClient(cc *grpc.ClientConn) PLControllerClient {
	return &pLControllerClient{cc}
}

func (c *pLControllerClient) Ping(ctx context.Context, opts ...grpc.CallOption) (PLController_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PLController_serviceDesc.Streams[0], "/pb.PLController/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerPingClient{stream}
	return x, nil
}

type PLController_PingClient interface {
	Send(*datamodel.PingArg) error
	Recv() (*datamodel.Ping, error)
	grpc.ClientStream
}

type pLControllerPingClient struct {
	grpc.ClientStream
}

func (x *pLControllerPingClient) Send(m *datamodel.PingArg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pLControllerPingClient) Recv() (*datamodel.Ping, error) {
	m := new(datamodel.Ping)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) Traceroute(ctx context.Context, opts ...grpc.CallOption) (PLController_TracerouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PLController_serviceDesc.Streams[1], "/pb.PLController/Traceroute", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerTracerouteClient{stream}
	return x, nil
}

type PLController_TracerouteClient interface {
	Send(*datamodel.TracerouteArg) error
	Recv() (*datamodel.Traceroute, error)
	grpc.ClientStream
}

type pLControllerTracerouteClient struct {
	grpc.ClientStream
}

func (x *pLControllerTracerouteClient) Send(m *datamodel.TracerouteArg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pLControllerTracerouteClient) Recv() (*datamodel.Traceroute, error) {
	m := new(datamodel.Traceroute)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) ReceiveSpoof(ctx context.Context, in *datamodel.RecSpoof, opts ...grpc.CallOption) (PLController_ReceiveSpoofClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PLController_serviceDesc.Streams[2], "/pb.PLController/ReceiveSpoof", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerReceiveSpoofClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PLController_ReceiveSpoofClient interface {
	Recv() (*datamodel.NotifyRecSpoofResponse, error)
	grpc.ClientStream
}

type pLControllerReceiveSpoofClient struct {
	grpc.ClientStream
}

func (x *pLControllerReceiveSpoofClient) Recv() (*datamodel.NotifyRecSpoofResponse, error) {
	m := new(datamodel.NotifyRecSpoofResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) GetVPs(ctx context.Context, in *datamodel.VPRequest, opts ...grpc.CallOption) (PLController_GetVPsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PLController_serviceDesc.Streams[3], "/pb.PLController/GetVPs", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerGetVPsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PLController_GetVPsClient interface {
	Recv() (*datamodel.VPReturn, error)
	grpc.ClientStream
}

type pLControllerGetVPsClient struct {
	grpc.ClientStream
}

func (x *pLControllerGetVPsClient) Recv() (*datamodel.VPReturn, error) {
	m := new(datamodel.VPReturn)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) AcceptProbes(ctx context.Context, in *datamodel.SpoofedProbes, opts ...grpc.CallOption) (*datamodel.SpoofedProbesResponse, error) {
	out := new(datamodel.SpoofedProbesResponse)
	err := c.cc.Invoke(ctx, "/pb.PLController/AcceptProbes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PLControllerServer is the server API for PLController service.
type PLControllerServer interface {
	Ping(PLController_PingServer) error
	Traceroute(PLController_TracerouteServer) error
	ReceiveSpoof(*datamodel.RecSpoof, PLController_ReceiveSpoofServer) error
	GetVPs(*datamodel.VPRequest, PLController_GetVPsServer) error
	AcceptProbes(context.Context, *datamodel.SpoofedProbes) (*datamodel.SpoofedProbesResponse, error)
}

// UnimplementedPLControllerServer can be embedded to have forward compatible implementations.
type UnimplementedPLControllerServer struct {
}

func (*UnimplementedPLControllerServer) Ping(srv PLController_PingServer) error {
	return status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedPLControllerServer) Traceroute(srv PLController_TracerouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Traceroute not implemented")
}
func (*UnimplementedPLControllerServer) ReceiveSpoof(req *datamodel.RecSpoof, srv PLController_ReceiveSpoofServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveSpoof not implemented")
}
func (*UnimplementedPLControllerServer) GetVPs(req *datamodel.VPRequest, srv PLController_GetVPsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVPs not implemented")
}
func (*UnimplementedPLControllerServer) AcceptProbes(ctx context.Context, req *datamodel.SpoofedProbes) (*datamodel.SpoofedProbesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptProbes not implemented")
}

func RegisterPLControllerServer(s *grpc.Server, srv PLControllerServer) {
	s.RegisterService(&_PLController_serviceDesc, srv)
}

func _PLController_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PLControllerServer).Ping(&pLControllerPingServer{stream})
}

type PLController_PingServer interface {
	Send(*datamodel.Ping) error
	Recv() (*datamodel.PingArg, error)
	grpc.ServerStream
}

type pLControllerPingServer struct {
	grpc.ServerStream
}

func (x *pLControllerPingServer) Send(m *datamodel.Ping) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pLControllerPingServer) Recv() (*datamodel.PingArg, error) {
	m := new(datamodel.PingArg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PLController_Traceroute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PLControllerServer).Traceroute(&pLControllerTracerouteServer{stream})
}

type PLController_TracerouteServer interface {
	Send(*datamodel.Traceroute) error
	Recv() (*datamodel.TracerouteArg, error)
	grpc.ServerStream
}

type pLControllerTracerouteServer struct {
	grpc.ServerStream
}

func (x *pLControllerTracerouteServer) Send(m *datamodel.Traceroute) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pLControllerTracerouteServer) Recv() (*datamodel.TracerouteArg, error) {
	m := new(datamodel.TracerouteArg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PLController_ReceiveSpoof_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(datamodel.RecSpoof)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PLControllerServer).ReceiveSpoof(m, &pLControllerReceiveSpoofServer{stream})
}

type PLController_ReceiveSpoofServer interface {
	Send(*datamodel.NotifyRecSpoofResponse) error
	grpc.ServerStream
}

type pLControllerReceiveSpoofServer struct {
	grpc.ServerStream
}

func (x *pLControllerReceiveSpoofServer) Send(m *datamodel.NotifyRecSpoofResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PLController_GetVPs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(datamodel.VPRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PLControllerServer).GetVPs(m, &pLControllerGetVPsServer{stream})
}

type PLController_GetVPsServer interface {
	Send(*datamodel.VPReturn) error
	grpc.ServerStream
}

type pLControllerGetVPsServer struct {
	grpc.ServerStream
}

func (x *pLControllerGetVPsServer) Send(m *datamodel.VPReturn) error {
	return x.ServerStream.SendMsg(m)
}

func _PLController_AcceptProbes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(datamodel.SpoofedProbes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PLControllerServer).AcceptProbes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PLController/AcceptProbes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PLControllerServer).AcceptProbes(ctx, req.(*datamodel.SpoofedProbes))
	}
	return interceptor(ctx, in, info, handler)
}

var _PLController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PLController",
	HandlerType: (*PLControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptProbes",
			Handler:    _PLController_AcceptProbes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ping",
			Handler:       _PLController_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Traceroute",
			Handler:       _PLController_Traceroute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveSpoof",
			Handler:       _PLController_ReceiveSpoof_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVPs",
			Handler:       _PLController_GetVPs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/NEU-SNS/ReverseTraceroute/plcontroller/pb/plcontrollerapi.proto",
}
