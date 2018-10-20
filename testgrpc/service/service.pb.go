// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

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

type RectParamIn struct {
	Width                int64    `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               int64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RectParamIn) Reset()         { *m = RectParamIn{} }
func (m *RectParamIn) String() string { return proto.CompactTextString(m) }
func (*RectParamIn) ProtoMessage()    {}
func (*RectParamIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1c16e0c9bc350dc3, []int{0}
}
func (m *RectParamIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RectParamIn.Unmarshal(m, b)
}
func (m *RectParamIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RectParamIn.Marshal(b, m, deterministic)
}
func (dst *RectParamIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RectParamIn.Merge(dst, src)
}
func (m *RectParamIn) XXX_Size() int {
	return xxx_messageInfo_RectParamIn.Size(m)
}
func (m *RectParamIn) XXX_DiscardUnknown() {
	xxx_messageInfo_RectParamIn.DiscardUnknown(m)
}

var xxx_messageInfo_RectParamIn proto.InternalMessageInfo

func (m *RectParamIn) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *RectParamIn) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type RectParamOut struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RectParamOut) Reset()         { *m = RectParamOut{} }
func (m *RectParamOut) String() string { return proto.CompactTextString(m) }
func (*RectParamOut) ProtoMessage()    {}
func (*RectParamOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1c16e0c9bc350dc3, []int{1}
}
func (m *RectParamOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RectParamOut.Unmarshal(m, b)
}
func (m *RectParamOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RectParamOut.Marshal(b, m, deterministic)
}
func (dst *RectParamOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RectParamOut.Merge(dst, src)
}
func (m *RectParamOut) XXX_Size() int {
	return xxx_messageInfo_RectParamOut.Size(m)
}
func (m *RectParamOut) XXX_DiscardUnknown() {
	xxx_messageInfo_RectParamOut.DiscardUnknown(m)
}

var xxx_messageInfo_RectParamOut proto.InternalMessageInfo

func (m *RectParamOut) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AuthParamIn struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TokenId              string   `protobuf:"bytes,2,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthParamIn) Reset()         { *m = AuthParamIn{} }
func (m *AuthParamIn) String() string { return proto.CompactTextString(m) }
func (*AuthParamIn) ProtoMessage()    {}
func (*AuthParamIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1c16e0c9bc350dc3, []int{2}
}
func (m *AuthParamIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthParamIn.Unmarshal(m, b)
}
func (m *AuthParamIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthParamIn.Marshal(b, m, deterministic)
}
func (dst *AuthParamIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthParamIn.Merge(dst, src)
}
func (m *AuthParamIn) XXX_Size() int {
	return xxx_messageInfo_AuthParamIn.Size(m)
}
func (m *AuthParamIn) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthParamIn.DiscardUnknown(m)
}

var xxx_messageInfo_AuthParamIn proto.InternalMessageInfo

func (m *AuthParamIn) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AuthParamIn) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

type AuthParamOut struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthParamOut) Reset()         { *m = AuthParamOut{} }
func (m *AuthParamOut) String() string { return proto.CompactTextString(m) }
func (*AuthParamOut) ProtoMessage()    {}
func (*AuthParamOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_1c16e0c9bc350dc3, []int{3}
}
func (m *AuthParamOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthParamOut.Unmarshal(m, b)
}
func (m *AuthParamOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthParamOut.Marshal(b, m, deterministic)
}
func (dst *AuthParamOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthParamOut.Merge(dst, src)
}
func (m *AuthParamOut) XXX_Size() int {
	return xxx_messageInfo_AuthParamOut.Size(m)
}
func (m *AuthParamOut) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthParamOut.DiscardUnknown(m)
}

var xxx_messageInfo_AuthParamOut proto.InternalMessageInfo

func (m *AuthParamOut) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*RectParamIn)(nil), "service.RectParamIn")
	proto.RegisterType((*RectParamOut)(nil), "service.RectParamOut")
	proto.RegisterType((*AuthParamIn)(nil), "service.AuthParamIn")
	proto.RegisterType((*AuthParamOut)(nil), "service.AuthParamOut")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RectClient is the client API for Rect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RectClient interface {
	Area(ctx context.Context, in *RectParamIn, opts ...grpc.CallOption) (*RectParamOut, error)
	Perimeter(ctx context.Context, in *RectParamIn, opts ...grpc.CallOption) (*RectParamOut, error)
}

type rectClient struct {
	cc *grpc.ClientConn
}

func NewRectClient(cc *grpc.ClientConn) RectClient {
	return &rectClient{cc}
}

func (c *rectClient) Area(ctx context.Context, in *RectParamIn, opts ...grpc.CallOption) (*RectParamOut, error) {
	out := new(RectParamOut)
	err := c.cc.Invoke(ctx, "/service.Rect/Area", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rectClient) Perimeter(ctx context.Context, in *RectParamIn, opts ...grpc.CallOption) (*RectParamOut, error) {
	out := new(RectParamOut)
	err := c.cc.Invoke(ctx, "/service.Rect/Perimeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RectServer is the server API for Rect service.
type RectServer interface {
	Area(context.Context, *RectParamIn) (*RectParamOut, error)
	Perimeter(context.Context, *RectParamIn) (*RectParamOut, error)
}

func RegisterRectServer(s *grpc.Server, srv RectServer) {
	s.RegisterService(&_Rect_serviceDesc, srv)
}

func _Rect_Area_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RectParamIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RectServer).Area(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Rect/Area",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RectServer).Area(ctx, req.(*RectParamIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rect_Perimeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RectParamIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RectServer).Perimeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Rect/Perimeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RectServer).Perimeter(ctx, req.(*RectParamIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Rect",
	HandlerType: (*RectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Area",
			Handler:    _Rect_Area_Handler,
		},
		{
			MethodName: "Perimeter",
			Handler:    _Rect_Perimeter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Auth(ctx context.Context, in *AuthParamIn, opts ...grpc.CallOption) (*AuthParamOut, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Auth(ctx context.Context, in *AuthParamIn, opts ...grpc.CallOption) (*AuthParamOut, error) {
	out := new(AuthParamOut)
	err := c.cc.Invoke(ctx, "/service.Auth/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Auth(context.Context, *AuthParamIn) (*AuthParamOut, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthParamIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Auth/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Auth(ctx, req.(*AuthParamIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Auth_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_1c16e0c9bc350dc3) }

var fileDescriptor_service_1c16e0c9bc350dc3 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xac, 0xb9,
	0xb8, 0x83, 0x52, 0x93, 0x4b, 0x02, 0x12, 0x8b, 0x12, 0x73, 0x3d, 0xf3, 0x84, 0x44, 0xb8, 0x58,
	0xcb, 0x33, 0x53, 0x4a, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x21, 0x31,
	0x2e, 0xb6, 0x8c, 0xd4, 0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x26, 0xb0, 0x30, 0x94, 0xa7, 0xa4, 0xc6,
	0xc5, 0x03, 0xd7, 0xec, 0x5f, 0x5a, 0x02, 0x52, 0x57, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0xd5,
	0x0e, 0xe5, 0x29, 0x59, 0x72, 0x71, 0x3b, 0x96, 0x96, 0x64, 0xc0, 0x2c, 0x11, 0xe0, 0x62, 0x2e,
	0xcd, 0x4c, 0x81, 0xaa, 0x01, 0x31, 0x85, 0x24, 0xb8, 0xd8, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x3c,
	0x53, 0xc0, 0x36, 0x70, 0x06, 0xc1, 0xb8, 0x20, 0x2b, 0xe0, 0x5a, 0x31, 0xad, 0xe0, 0x84, 0x59,
	0x61, 0x54, 0xc9, 0xc5, 0x02, 0x72, 0x8a, 0x90, 0x29, 0x17, 0x8b, 0x63, 0x51, 0x6a, 0xa2, 0x90,
	0x88, 0x1e, 0xcc, 0xc3, 0x48, 0xde, 0x93, 0x12, 0xc5, 0x14, 0xf5, 0x2f, 0x2d, 0x51, 0x62, 0x10,
	0xb2, 0xe2, 0xe2, 0x0c, 0x48, 0x2d, 0xca, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0x22, 0x51, 0xaf, 0x91,
	0x2d, 0x17, 0x0b, 0xc8, 0x89, 0x60, 0xab, 0x41, 0x34, 0x42, 0x3b, 0x92, 0xa7, 0x91, 0xb4, 0x23,
	0xfb, 0x47, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x23, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x92, 0xd9, 0x8b, 0xa2, 0x01, 0x00, 0x00,
}
