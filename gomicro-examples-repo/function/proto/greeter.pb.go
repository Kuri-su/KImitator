// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/function/proto/greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/function/proto/greeter.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package greeter

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/Greeter/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/examples/function/proto/greeter.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/examples/function/proto/greeter.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xcb, 0x0a, 0xc2, 0x30,
	0x10, 0x45, 0xad, 0xf8, 0x1c, 0xac, 0x8b, 0xac, 0x4a, 0x57, 0x92, 0x55, 0x41, 0x48, 0xc0, 0x82,
	0xbf, 0xa0, 0xeb, 0xfe, 0x41, 0x5b, 0xc6, 0x1a, 0x68, 0x32, 0x35, 0x0f, 0xf0, 0xf3, 0x85, 0x51,
	0x44, 0x77, 0x73, 0xe1, 0xce, 0x39, 0x17, 0xce, 0x83, 0x89, 0xf7, 0xd4, 0xa9, 0x9e, 0xac, 0xb6,
	0xa6, 0xf7, 0xa4, 0xf1, 0xd9, 0xda, 0x69, 0xc4, 0xa0, 0x6f, 0xc9, 0xf5, 0xd1, 0x90, 0xd3, 0x93,
	0xa7, 0x48, 0x7a, 0xf0, 0x88, 0x11, 0xbd, 0xe2, 0x24, 0x25, 0xec, 0xae, 0x38, 0x8e, 0xd4, 0xe0,
	0x23, 0x61, 0x88, 0x42, 0xc0, 0xc2, 0xb5, 0x16, 0x8b, 0xec, 0x90, 0x55, 0xdb, 0x86, 0x6f, 0x79,
	0x84, 0xfc, 0xd3, 0x09, 0x13, 0xb9, 0x80, 0xa2, 0x84, 0x0d, 0x53, 0x8c, 0x1b, 0x8a, 0x39, 0x17,
	0xbf, 0xf9, 0x54, 0xc3, 0xfa, 0xf2, 0x36, 0x88, 0x0a, 0x96, 0xfc, 0x27, 0x72, 0xf5, 0xeb, 0x28,
	0xf7, 0xea, 0x0f, 0x27, 0x67, 0xdd, 0x8a, 0xc7, 0xd4, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0x67, 0xb8, 0xdf, 0xc6, 0x00, 0x00, 0x00,
}
