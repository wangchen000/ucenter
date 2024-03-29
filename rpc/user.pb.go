// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package rpc

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

type Register struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	ValidCode            string   `protobuf:"bytes,2,opt,name=validCode,proto3" json:"validCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Register) GetValidCode() string {
	if m != nil {
		return m.ValidCode
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UserDto struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDto) Reset()         { *m = UserDto{} }
func (m *UserDto) String() string { return proto.CompactTextString(m) }
func (*UserDto) ProtoMessage()    {}
func (*UserDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDto.Unmarshal(m, b)
}
func (m *UserDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDto.Marshal(b, m, deterministic)
}
func (m *UserDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDto.Merge(m, src)
}
func (m *UserDto) XXX_Size() int {
	return xxx_messageInfo_UserDto.Size(m)
}
func (m *UserDto) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDto.DiscardUnknown(m)
}

var xxx_messageInfo_UserDto proto.InternalMessageInfo

func (m *UserDto) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserDto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDto) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func init() {
	proto.RegisterType((*Register)(nil), "rpc.Register")
	proto.RegisterType((*User)(nil), "rpc.User")
	proto.RegisterType((*UserDto)(nil), "rpc.UserDto")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x72, 0xe0, 0xe2, 0x08,
	0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x12, 0x12, 0xe3, 0x62, 0xcb, 0xcd, 0x4f, 0xca, 0xcc,
	0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0xcb, 0x12,
	0x73, 0x32, 0x53, 0x9c, 0xf3, 0x53, 0x52, 0x25, 0x98, 0xc0, 0x52, 0x08, 0x01, 0xa5, 0x38, 0x2e,
	0x96, 0xd0, 0xe2, 0xd4, 0x22, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0xa8, 0x4e, 0xa6, 0xcc, 0x14,
	0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0x5c, 0x98, 0x06, 0x30, 0x5b, 0x48, 0x8a, 0x8b, 0xa3, 0x20,
	0xb1, 0xb8, 0xb8, 0x3c, 0xbf, 0x28, 0x45, 0x82, 0x19, 0x2c, 0x0e, 0xe7, 0x23, 0xd9, 0xce, 0x82,
	0x6c, 0xbb, 0x92, 0x2b, 0x17, 0x3b, 0xc8, 0x7c, 0x97, 0x92, 0x7c, 0xa2, 0xac, 0x40, 0x18, 0xc3,
	0x8a, 0x6c, 0x8c, 0x91, 0x05, 0x17, 0x37, 0xc8, 0x98, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0x21, 0x4d, 0x98, 0xbf, 0x8b, 0x2a, 0x85, 0x78, 0xf5, 0x8a, 0x0a, 0x92, 0xf5, 0x60, 0xc1, 0x20,
	0xc5, 0x03, 0xe6, 0x42, 0xed, 0x54, 0x62, 0x48, 0x62, 0x03, 0x07, 0x97, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x50, 0x3e, 0x15, 0x6f, 0x3c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	//用户注册
	Registry(ctx context.Context, in *Register, opts ...grpc.CallOption) (*UserDto, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Registry(ctx context.Context, in *Register, opts ...grpc.CallOption) (*UserDto, error) {
	out := new(UserDto)
	err := c.cc.Invoke(ctx, "/rpc.UserService/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//用户注册
	Registry(context.Context, *Register) (*UserDto, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Registry(ctx context.Context, req *Register) (*UserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registry not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.UserService/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Registry(ctx, req.(*Register))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registry",
			Handler:    _UserService_Registry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
