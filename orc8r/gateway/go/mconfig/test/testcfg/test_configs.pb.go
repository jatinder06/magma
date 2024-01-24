// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_configs.proto

package testcfg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//------------------------------------------------------------------------------
// Service 1 configs
//------------------------------------------------------------------------------
type Service1Config struct {
	Str1                 string   `protobuf:"bytes,1,opt,name=str1,proto3" json:"str1,omitempty"`
	Str2                 string   `protobuf:"bytes,2,opt,name=str2,proto3" json:"str2,omitempty"`
	Uint1                uint32   `protobuf:"varint,3,opt,name=uint1,proto3" json:"uint1,omitempty"`
	Uint2                uint32   `protobuf:"varint,4,opt,name=uint2,proto3" json:"uint2,omitempty"`
	StrArr               []string `protobuf:"bytes,5,rep,name=str_arr,json=strArr,proto3" json:"str_arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service1Config) Reset()         { *m = Service1Config{} }
func (m *Service1Config) String() string { return proto.CompactTextString(m) }
func (*Service1Config) ProtoMessage()    {}
func (*Service1Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ab85a1b6fb8842, []int{0}
}

func (m *Service1Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service1Config.Unmarshal(m, b)
}
func (m *Service1Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service1Config.Marshal(b, m, deterministic)
}
func (m *Service1Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service1Config.Merge(m, src)
}
func (m *Service1Config) XXX_Size() int {
	return xxx_messageInfo_Service1Config.Size(m)
}
func (m *Service1Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Service1Config.DiscardUnknown(m)
}

var xxx_messageInfo_Service1Config proto.InternalMessageInfo

func (m *Service1Config) GetStr1() string {
	if m != nil {
		return m.Str1
	}
	return ""
}

func (m *Service1Config) GetStr2() string {
	if m != nil {
		return m.Str2
	}
	return ""
}

func (m *Service1Config) GetUint1() uint32 {
	if m != nil {
		return m.Uint1
	}
	return 0
}

func (m *Service1Config) GetUint2() uint32 {
	if m != nil {
		return m.Uint2
	}
	return 0
}

func (m *Service1Config) GetStrArr() []string {
	if m != nil {
		return m.StrArr
	}
	return nil
}

//------------------------------------------------------------------------------
// Service 1 configs
//------------------------------------------------------------------------------
type Service2Config struct {
	Str21                string   `protobuf:"bytes,1,opt,name=str21,proto3" json:"str21,omitempty"`
	Str22                string   `protobuf:"bytes,2,opt,name=str22,proto3" json:"str22,omitempty"`
	Uint21               uint32   `protobuf:"varint,3,opt,name=uint21,proto3" json:"uint21,omitempty"`
	Uint22               uint32   `protobuf:"varint,4,opt,name=uint22,proto3" json:"uint22,omitempty"`
	Float1               float32  `protobuf:"fixed32,5,opt,name=float1,proto3" json:"float1,omitempty"`
	Bool1                bool     `protobuf:"varint,6,opt,name=bool1,proto3" json:"bool1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service2Config) Reset()         { *m = Service2Config{} }
func (m *Service2Config) String() string { return proto.CompactTextString(m) }
func (*Service2Config) ProtoMessage()    {}
func (*Service2Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ab85a1b6fb8842, []int{1}
}

func (m *Service2Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service2Config.Unmarshal(m, b)
}
func (m *Service2Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service2Config.Marshal(b, m, deterministic)
}
func (m *Service2Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service2Config.Merge(m, src)
}
func (m *Service2Config) XXX_Size() int {
	return xxx_messageInfo_Service2Config.Size(m)
}
func (m *Service2Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Service2Config.DiscardUnknown(m)
}

var xxx_messageInfo_Service2Config proto.InternalMessageInfo

func (m *Service2Config) GetStr21() string {
	if m != nil {
		return m.Str21
	}
	return ""
}

func (m *Service2Config) GetStr22() string {
	if m != nil {
		return m.Str22
	}
	return ""
}

func (m *Service2Config) GetUint21() uint32 {
	if m != nil {
		return m.Uint21
	}
	return 0
}

func (m *Service2Config) GetUint22() uint32 {
	if m != nil {
		return m.Uint22
	}
	return 0
}

func (m *Service2Config) GetFloat1() float32 {
	if m != nil {
		return m.Float1
	}
	return 0
}

func (m *Service2Config) GetBool1() bool {
	if m != nil {
		return m.Bool1
	}
	return false
}

func init() {
	proto.RegisterType((*Service1Config)(nil), "magma.mconfig.Service1Config")
	proto.RegisterType((*Service2Config)(nil), "magma.mconfig.Service2Config")
}

func init() { proto.RegisterFile("test_configs.proto", fileDescriptor_b5ab85a1b6fb8842) }

var fileDescriptor_b5ab85a1b6fb8842 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe5, 0xb6, 0x09, 0xd4, 0x52, 0x19, 0xac, 0x0a, 0x3c, 0x5a, 0x9d, 0x2c, 0x86, 0x46,
	0x36, 0x0b, 0x2b, 0xf0, 0x06, 0x66, 0x63, 0xa9, 0x5c, 0x2b, 0x89, 0x22, 0x35, 0x31, 0xba, 0xbe,
	0x80, 0x18, 0x78, 0x11, 0x9e, 0x16, 0xc5, 0x3f, 0x61, 0xb1, 0xfc, 0x7d, 0xc3, 0x3d, 0x47, 0x87,
	0x32, 0x6c, 0x03, 0x9e, 0x9c, 0x9f, 0xba, 0xa1, 0x0f, 0xc7, 0x77, 0xf0, 0xe8, 0xd9, 0x6e, 0xb4,
	0xfd, 0x68, 0x8f, 0x63, 0xb2, 0x87, 0x1f, 0x7a, 0xf3, 0xda, 0xc2, 0xe7, 0xe0, 0x5a, 0xf5, 0x12,
	0x0d, 0x63, 0x74, 0x13, 0x10, 0x14, 0x27, 0x82, 0xc8, 0xad, 0x89, 0xff, 0xec, 0x34, 0x5f, 0x2d,
	0x4e, 0xb3, 0x3d, 0xad, 0x3e, 0x86, 0x09, 0x15, 0x5f, 0x0b, 0x22, 0x77, 0x26, 0x41, 0xb1, 0x9a,
	0x6f, 0xfe, 0xad, 0x66, 0x77, 0xf4, 0x2a, 0x20, 0x9c, 0x2c, 0x00, 0xaf, 0xc4, 0x5a, 0x6e, 0x4d,
	0x1d, 0x10, 0x9e, 0x00, 0x0e, 0xbf, 0x64, 0xc9, 0xd7, 0x39, 0x7f, 0x4f, 0xab, 0xf9, 0x7e, 0x29,
	0x90, 0xa0, 0xd8, 0x52, 0x21, 0x01, 0xbb, 0xa5, 0x75, 0x0c, 0x28, 0x25, 0x32, 0x2d, 0xbe, 0xd4,
	0xc8, 0x34, 0xfb, 0xee, 0xe2, 0x2d, 0x2a, 0x5e, 0x09, 0x22, 0x57, 0x26, 0xd3, 0x7c, 0xfd, 0xec,
	0xfd, 0x45, 0xf1, 0x5a, 0x10, 0x79, 0x6d, 0x12, 0x3c, 0xdf, 0xbf, 0xc9, 0x38, 0x56, 0xe3, 0x1d,
	0x3c, 0xba, 0xa6, 0xb7, 0xd8, 0x7e, 0xd9, 0xef, 0x26, 0x4f, 0xd7, 0xcc, 0xe3, 0xc6, 0xc7, 0x75,
	0xfd, 0xb9, 0x8e, 0xeb, 0x3e, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd6, 0x08, 0x87, 0x73,
	0x01, 0x00, 0x00,
}
