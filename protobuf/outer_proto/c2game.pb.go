// Code generated by protoc-gen-go. DO NOT EDIT.
// source: c2game.proto

package outer_proto

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

//向game服务器注册
type C2GLogin struct {
	Pid                  string   `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GLogin) Reset()         { *m = C2GLogin{} }
func (m *C2GLogin) String() string { return proto.CompactTextString(m) }
func (*C2GLogin) ProtoMessage()    {}
func (*C2GLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc153bb8ab5e36b2, []int{0}
}

func (m *C2GLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GLogin.Unmarshal(m, b)
}
func (m *C2GLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GLogin.Marshal(b, m, deterministic)
}
func (m *C2GLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GLogin.Merge(m, src)
}
func (m *C2GLogin) XXX_Size() int {
	return xxx_messageInfo_C2GLogin.Size(m)
}
func (m *C2GLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GLogin.DiscardUnknown(m)
}

var xxx_messageInfo_C2GLogin proto.InternalMessageInfo

func (m *C2GLogin) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

//向服务器定时发送ping,保证存活
type C2GPing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GPing) Reset()         { *m = C2GPing{} }
func (m *C2GPing) String() string { return proto.CompactTextString(m) }
func (*C2GPing) ProtoMessage()    {}
func (*C2GPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc153bb8ab5e36b2, []int{1}
}

func (m *C2GPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GPing.Unmarshal(m, b)
}
func (m *C2GPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GPing.Marshal(b, m, deterministic)
}
func (m *C2GPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GPing.Merge(m, src)
}
func (m *C2GPing) XXX_Size() int {
	return xxx_messageInfo_C2GPing.Size(m)
}
func (m *C2GPing) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GPing.DiscardUnknown(m)
}

var xxx_messageInfo_C2GPing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*C2GLogin)(nil), "outer_proto.c2g_login")
	proto.RegisterType((*C2GPing)(nil), "outer_proto.c2g_ping")
}

func init() { proto.RegisterFile("c2game.proto", fileDescriptor_dc153bb8ab5e36b2) }

var fileDescriptor_dc153bb8ab5e36b2 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x36, 0x4a, 0x4f,
	0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x2f, 0x2d, 0x49, 0x2d, 0x8a,
	0x07, 0x73, 0x94, 0x64, 0xb9, 0x38, 0x93, 0x8d, 0xd2, 0xe3, 0x73, 0xf2, 0xd3, 0x33, 0xf3, 0x84,
	0x04, 0xb8, 0x98, 0x0b, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x25,
	0x2e, 0x2e, 0x0e, 0x90, 0x74, 0x41, 0x66, 0x5e, 0x7a, 0x12, 0x1b, 0x58, 0x87, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0x2d, 0x20, 0x7e, 0x4e, 0x00, 0x00, 0x00,
}