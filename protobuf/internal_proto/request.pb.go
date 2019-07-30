// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package internal_proto

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

type MsgBodyRequest struct {
	Connect              *Request_ConnectStruct `protobuf:"bytes,1,opt,name=connect,proto3" json:"connect,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MsgBodyRequest) Reset()         { *m = MsgBodyRequest{} }
func (m *MsgBodyRequest) String() string { return proto.CompactTextString(m) }
func (*MsgBodyRequest) ProtoMessage()    {}
func (*MsgBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *MsgBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBodyRequest.Unmarshal(m, b)
}
func (m *MsgBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBodyRequest.Marshal(b, m, deterministic)
}
func (m *MsgBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBodyRequest.Merge(m, src)
}
func (m *MsgBodyRequest) XXX_Size() int {
	return xxx_messageInfo_MsgBodyRequest.Size(m)
}
func (m *MsgBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBodyRequest proto.InternalMessageInfo

func (m *MsgBodyRequest) GetConnect() *Request_ConnectStruct {
	if m != nil {
		return m.Connect
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgBodyRequest)(nil), "internal_proto.MsgBodyRequest")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0x89, 0x07, 0xf3, 0xa5, 0x84, 0xc0, 0x54, 0x7c, 0x71, 0x49, 0x51, 0x69, 0x32,
	0x54, 0x8d, 0x52, 0x20, 0x17, 0x9f, 0x6f, 0x71, 0xba, 0x53, 0x7e, 0x4a, 0x65, 0x10, 0x44, 0xaf,
	0x90, 0x3d, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x5e, 0x6a, 0x72, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0xaa, 0x1e, 0xaa, 0x39, 0x7a, 0x50, 0x95, 0xf1, 0xce, 0x10, 0x65, 0xc1, 0x60, 0xf3,
	0x82, 0x60, 0xba, 0x92, 0xd8, 0xc0, 0xaa, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x83,
	0x6d, 0x5a, 0x8e, 0x00, 0x00, 0x00,
}
