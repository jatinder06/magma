// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/hello.proto

package protos

import (
	context "context"
	fmt "fmt"
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

type HelloRequest struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	// A grpc_err_code is an unsigned 32-bit error code as defined in the gRPC
	// spec.
	//	OK Code = 0
	//	Canceled Code = 1
	//	Unknown Code = 2
	//	InvalidArgument Code = 3
	//	DeadlineExceeded Code = 4
	//	NotFound Code = 5
	//	AlreadyExists Code = 6
	//	PermissionDenied Code = 7
	//	ResourceExhausted Code = 8
	//	FailedPrecondition Code = 9
	//	Aborted Code = 10
	//	OutOfRange Code = 11
	//	Unimplemented Code = 12
	//	Internal Code = 13
	//	Unavailable Code = 14
	//	DataLoss Code = 15
	//	Unauthenticated Code = 16
	// See https://github.com/grpc/grpc-go/blob/master/codes/codes.go for details
	// This field lets user request for a grpc err code, and expect server to
	// successfully send back this err code. If something goes wrong in the path,
	// the error code will likely come back different.
	GrpcErrCode          uint32   `protobuf:"varint,2,opt,name=grpc_err_code,json=grpcErrCode,proto3" json:"grpc_err_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6edc8f794450209b, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *HelloRequest) GetGrpcErrCode() uint32 {
	if m != nil {
		return m.GrpcErrCode
	}
	return 0
}

type HelloReply struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6edc8f794450209b, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "magma.feg.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "magma.feg.HelloReply")
}

func init() { proto.RegisterFile("feg/protos/hello.proto", fileDescriptor_6edc8f794450209b) }

var fileDescriptor_6edc8f794450209b = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4b, 0x4d, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0x73, 0x84,
	0x38, 0x73, 0x13, 0xd3, 0x73, 0x13, 0xf5, 0xd2, 0x52, 0xd3, 0x95, 0xfc, 0xb8, 0x78, 0x3c, 0x40,
	0x32, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0xe9, 0x45, 0xa9, 0xa9,
	0x25, 0x99, 0x79, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x12, 0x17,
	0x6f, 0x7a, 0x51, 0x41, 0x72, 0x7c, 0x6a, 0x51, 0x51, 0x7c, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x6f, 0x10, 0x37, 0x48, 0xd0, 0xb5, 0xa8, 0xc8, 0x39, 0x3f, 0x25, 0x55, 0x49,
	0x83, 0x8b, 0x0b, 0x6a, 0x5e, 0x41, 0x4e, 0x25, 0x3e, 0xd3, 0x8c, 0x5c, 0xb9, 0x58, 0xc1, 0x2a,
	0x85, 0x6c, 0xb8, 0x38, 0x82, 0x13, 0x2b, 0x21, 0x6c, 0x71, 0x3d, 0xb8, 0xd3, 0xf4, 0x90, 0xdd,
	0x25, 0x25, 0x8a, 0x29, 0x51, 0x90, 0x53, 0xa9, 0xc4, 0xe0, 0x24, 0x1d, 0x25, 0x09, 0x96, 0xd1,
	0x07, 0xf9, 0x35, 0x39, 0x27, 0xbf, 0x34, 0x45, 0x3f, 0x3d, 0x1f, 0xea, 0xe9, 0x24, 0x36, 0x30,
	0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xd7, 0x70, 0xda, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	// Sample rpc for Hello service
	//
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/magma.feg.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	// Sample rpc for Hello service
	//
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/hello.proto",
}
