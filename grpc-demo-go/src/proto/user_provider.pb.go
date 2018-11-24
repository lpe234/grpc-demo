// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_provider.proto

package user_provider

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

// 请求
type UserIdRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_provider_83aa1d764e9c5417, []int{0}
}
func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (dst *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(dst, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 响应
type UserVoReplay struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserVoReplay) Reset()         { *m = UserVoReplay{} }
func (m *UserVoReplay) String() string { return proto.CompactTextString(m) }
func (*UserVoReplay) ProtoMessage()    {}
func (*UserVoReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_provider_83aa1d764e9c5417, []int{1}
}
func (m *UserVoReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserVoReplay.Unmarshal(m, b)
}
func (m *UserVoReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserVoReplay.Marshal(b, m, deterministic)
}
func (dst *UserVoReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserVoReplay.Merge(dst, src)
}
func (m *UserVoReplay) XXX_Size() int {
	return xxx_messageInfo_UserVoReplay.Size(m)
}
func (m *UserVoReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_UserVoReplay.DiscardUnknown(m)
}

var xxx_messageInfo_UserVoReplay proto.InternalMessageInfo

func (m *UserVoReplay) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserVoReplay) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*UserIdRequest)(nil), "UserIdRequest")
	proto.RegisterType((*UserVoReplay)(nil), "UserVoReplay")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserProviderClient is the client API for UserProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserProviderClient interface {
	// 根据用户id获取用户信息的服务(具体服务/函数)
	GetByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserVoReplay, error)
}

type userProviderClient struct {
	cc *grpc.ClientConn
}

func NewUserProviderClient(cc *grpc.ClientConn) UserProviderClient {
	return &userProviderClient{cc}
}

func (c *userProviderClient) GetByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserVoReplay, error) {
	out := new(UserVoReplay)
	err := c.cc.Invoke(ctx, "/UserProvider/getByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProviderServer is the server API for UserProvider service.
type UserProviderServer interface {
	// 根据用户id获取用户信息的服务(具体服务/函数)
	GetByUserId(context.Context, *UserIdRequest) (*UserVoReplay, error)
}

func RegisterUserProviderServer(s *grpc.Server, srv UserProviderServer) {
	s.RegisterService(&_UserProvider_serviceDesc, srv)
}

func _UserProvider_GetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProviderServer).GetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProvider/GetByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProviderServer).GetByUserId(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserProvider",
	HandlerType: (*UserProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getByUserId",
			Handler:    _UserProvider_GetByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_provider.proto",
}

func init() { proto.RegisterFile("user_provider.proto", fileDescriptor_user_provider_83aa1d764e9c5417) }

var fileDescriptor_user_provider_83aa1d764e9c5417 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x28, 0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0x92, 0xe7, 0xe2, 0x0d, 0x2d, 0x4e, 0x2d, 0xf2, 0x4c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c,
	0x51, 0xb2, 0xe2, 0xe2, 0x01, 0x29, 0x08, 0xcb, 0x0f, 0x4a, 0x2d, 0xc8, 0x49, 0xac, 0x44, 0x97,
	0x17, 0x92, 0xe2, 0xe2, 0x00, 0x99, 0x9b, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe7, 0x1b, 0xd9, 0x41, 0xf4, 0x06, 0x40, 0xad, 0x14, 0xd2, 0xe3, 0xe2, 0x4e, 0x4f,
	0x2d, 0x71, 0xaa, 0x84, 0xd8, 0x28, 0xc4, 0xa7, 0x87, 0x62, 0xb5, 0x14, 0xaf, 0x1e, 0xb2, 0x4d,
	0x4a, 0x0c, 0x4e, 0x72, 0x5c, 0x32, 0xc9, 0x79, 0x7a, 0x39, 0x05, 0xa9, 0x46, 0xc6, 0x26, 0x7a,
	0xe9, 0x45, 0x05, 0xc9, 0x60, 0x22, 0x25, 0x35, 0x37, 0x1f, 0xcc, 0x48, 0x62, 0x03, 0xfb, 0xc1,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x9e, 0xff, 0x93, 0xda, 0x00, 0x00, 0x00,
}
