// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

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

type ReviceBodyResponse struct {
	ConnectReply         *Response_ConnectStruct `protobuf:"bytes,1,opt,name=connectReply,proto3" json:"connectReply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ReviceBodyResponse) Reset()         { *m = ReviceBodyResponse{} }
func (m *ReviceBodyResponse) String() string { return proto.CompactTextString(m) }
func (*ReviceBodyResponse) ProtoMessage()    {}
func (*ReviceBodyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

func (m *ReviceBodyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviceBodyResponse.Unmarshal(m, b)
}
func (m *ReviceBodyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviceBodyResponse.Marshal(b, m, deterministic)
}
func (m *ReviceBodyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviceBodyResponse.Merge(m, src)
}
func (m *ReviceBodyResponse) XXX_Size() int {
	return xxx_messageInfo_ReviceBodyResponse.Size(m)
}
func (m *ReviceBodyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviceBodyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReviceBodyResponse proto.InternalMessageInfo

func (m *ReviceBodyResponse) GetConnectReply() *Response_ConnectStruct {
	if m != nil {
		return m.ConnectReply
	}
	return nil
}

func init() {
	proto.RegisterType((*ReviceBodyResponse)(nil), "internal_proto.ReviceBodyResponse")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x89, 0x07, 0xf3, 0xa5, 0x84, 0xc0, 0x54, 0x7c, 0x71, 0x49, 0x51, 0x69,
	0x72, 0x09, 0x44, 0x8d, 0x52, 0x02, 0x97, 0x50, 0x50, 0x6a, 0x59, 0x66, 0x72, 0xaa, 0x53, 0x7e,
	0x4a, 0x65, 0x10, 0x54, 0xbf, 0x90, 0x17, 0x17, 0x4f, 0x72, 0x7e, 0x5e, 0x5e, 0x6a, 0x72, 0x49,
	0x50, 0x6a, 0x41, 0x4e, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x9a, 0x1e, 0xaa, 0x81,
	0x7a, 0x30, 0xf5, 0xf1, 0xce, 0x10, 0xc5, 0xc1, 0x60, 0x93, 0x83, 0x50, 0xf4, 0x26, 0xb1, 0x81,
	0xd5, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x27, 0x2b, 0x62, 0xb7, 0x9e, 0x00, 0x00, 0x00,
}