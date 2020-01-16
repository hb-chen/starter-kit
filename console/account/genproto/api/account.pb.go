// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/account.proto

package go_micro_api_console_account

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	srv "github.com/micro-in-cn/starter-kit/console/account/genproto/srv"
	_ "github.com/micro/go-micro/api/proto"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f8c2a44ebe43c87, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Code                 int64              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string             `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Data                 *srv.LoginResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f8c2a44ebe43c87, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Response) GetData() *srv.LoginResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "go.micro.api.console.account.LoginRequest")
	proto.RegisterType((*Response)(nil), "go.micro.api.console.account.Response")
}

func init() { proto.RegisterFile("api/account.proto", fileDescriptor_3f8c2a44ebe43c87) }

var fileDescriptor_3f8c2a44ebe43c87 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0xc7, 0x7f, 0xfd, 0xad, 0xd6, 0x2d, 0x0a, 0xba, 0x20, 0x32, 0x86, 0x87, 0x32, 0x51, 0x8a,
	0xd2, 0x14, 0xe6, 0xc1, 0xf3, 0x76, 0x13, 0x76, 0xca, 0x59, 0x0f, 0x59, 0x1a, 0x6b, 0x60, 0xcb,
	0xab, 0x49, 0x5a, 0xdd, 0xff, 0xe7, 0x5f, 0xe5, 0x49, 0x9a, 0x74, 0x75, 0x22, 0x0c, 0x6f, 0xdf,
	0xf6, 0x7d, 0xde, 0xa7, 0xaf, 0xef, 0xa1, 0x21, 0x2b, 0x65, 0xc6, 0x38, 0x87, 0x4a, 0x59, 0x52,
	0x6a, 0xb0, 0x80, 0x2f, 0x0a, 0x20, 0x6b, 0xc9, 0x35, 0x10, 0x56, 0x4a, 0xc2, 0x41, 0x19, 0x58,
	0x09, 0xd2, 0x32, 0xe3, 0xb4, 0x90, 0xf6, 0xa5, 0x5a, 0x12, 0x0e, 0xeb, 0xcc, 0x51, 0x59, 0x01,
	0xa9, 0x0f, 0x8d, 0xca, 0x29, 0x9a, 0xe4, 0x65, 0xe3, 0xd9, 0x0e, 0x2e, 0x54, 0x0d, 0x9b, 0x52,
	0xc3, 0xfb, 0xc6, 0x63, 0x3c, 0x2d, 0x84, 0x4a, 0x6b, 0xb6, 0x92, 0x39, 0xb3, 0x22, 0xfb, 0x15,
	0x5a, 0xc5, 0xd0, 0xe8, 0xfa, 0xe7, 0x88, 0x93, 0x47, 0x74, 0xbc, 0x80, 0x42, 0x2a, 0x2a, 0x5e,
	0x2b, 0x61, 0x2c, 0xbe, 0x42, 0xfd, 0xca, 0x08, 0xad, 0xd8, 0x5a, 0x8c, 0x82, 0x38, 0x48, 0x06,
	0xf3, 0xc1, 0xe7, 0x3c, 0xd2, 0x61, 0x1c, 0x26, 0x31, 0xed, 0x4a, 0x0d, 0x56, 0x32, 0x63, 0xde,
	0x40, 0xe7, 0xa3, 0xff, 0xbb, 0x58, 0x94, 0x9c, 0xd1, 0xae, 0x34, 0x01, 0xd4, 0xa7, 0xc2, 0x94,
	0xa0, 0x8c, 0xc0, 0x18, 0x85, 0x1c, 0x72, 0x6f, 0xed, 0x51, 0x97, 0xf1, 0x39, 0x8a, 0x72, 0x61,
	0x99, 0x5c, 0x79, 0x09, 0x6d, 0x9f, 0xf0, 0x3d, 0x0a, 0x73, 0x66, 0xd9, 0xa8, 0x17, 0x07, 0xc9,
	0xd1, 0xf4, 0x92, 0x74, 0x7b, 0x34, 0xba, 0xde, 0xee, 0x8f, 0xb4, 0x73, 0x7b, 0x3d, 0x75, 0x0d,
	0xd3, 0x8f, 0x00, 0x1d, 0xce, 0x7c, 0x1d, 0x3f, 0xa1, 0x03, 0x87, 0xe0, 0x1b, 0xb2, 0xef, 0x0e,
	0x64, 0xf7, 0xff, 0xc7, 0xd7, 0xfb, 0xd9, 0xed, 0xe7, 0x26, 0xff, 0x70, 0x8a, 0xa2, 0x05, 0x14,
	0x50, 0x59, 0x7c, 0xd2, 0xf4, 0x34, 0xf4, 0x56, 0x72, 0xfa, 0xfd, 0xa2, 0xc3, 0x6f, 0x51, 0xf8,
	0xa0, 0x9e, 0xe1, 0x4f, 0xf0, 0x32, 0x72, 0xc7, 0xb9, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x4d,
	0x9d, 0x72, 0x28, 0x54, 0x02, 0x00, 0x00,
}
