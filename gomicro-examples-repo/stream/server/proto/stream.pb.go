// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/stream/server/proto/stream.proto

/*
Package stream is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/stream/server/proto/stream.proto

It has these top-level messages:
	Request
	Response
*/
package stream

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Response struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Streamer service

type StreamerClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Streamer_StreamClient, error)
	ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Streamer_ServerStreamClient, error)
}

type streamerClient struct {
	cc *grpc.ClientConn
}

func NewStreamerClient(cc *grpc.ClientConn) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Streamer_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Streamer_serviceDesc.Streams[0], c.cc, "/Streamer/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerStreamClient{stream}
	return x, nil
}

type Streamer_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamerStreamClient struct {
	grpc.ClientStream
}

func (x *streamerStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerClient) ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Streamer_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Streamer_serviceDesc.Streams[1], c.cc, "/Streamer/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streamer_ServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamerServerStreamClient struct {
	grpc.ClientStream
}

func (x *streamerServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Streamer service

type StreamerServer interface {
	Stream(Streamer_StreamServer) error
	ServerStream(*Request, Streamer_ServerStreamServer) error
}

func RegisterStreamerServer(s *grpc.Server, srv StreamerServer) {
	s.RegisterService(&_Streamer_serviceDesc, srv)
}

func _Streamer_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamerServer).Stream(&streamerStreamServer{stream})
}

type Streamer_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type streamerStreamServer struct {
	grpc.ServerStream
}

func (x *streamerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamerStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Streamer_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).ServerStream(m, &streamerServerStreamServer{stream})
}

type Streamer_ServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type streamerServerStreamServer struct {
	grpc.ServerStream
}

func (x *streamerServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Streamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Streamer_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _Streamer_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/micro/examples/stream/server/proto/stream.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/examples/stream/server/proto/stream.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2f, 0x4e, 0x2d,
	0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x8a, 0xe9, 0x81, 0x39, 0x4a, 0xf2,
	0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0xc9, 0xf9, 0xa5,
	0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x92, 0x02, 0x17, 0x47, 0x50,
	0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x76, 0x15, 0x46, 0x11, 0x5c, 0x1c, 0xc1, 0x60, 0x23,
	0x53, 0x8b, 0x84, 0x94, 0xb9, 0xd8, 0x20, 0x6c, 0x21, 0x0e, 0x3d, 0xa8, 0xb9, 0x52, 0x9c, 0x7a,
	0x30, 0x03, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x85, 0xd4, 0xb9, 0x78, 0x82, 0xc1, 0x0e, 0xc2,
	0xab, 0xd4, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x46, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16,
	0x4a, 0x84, 0x3f, 0xe1, 0x00, 0x00, 0x00,
}
