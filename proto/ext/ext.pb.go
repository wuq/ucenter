// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ext/ext.proto

package go_micro_srv_tenno_ucenter

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64ebe7fa13202e68, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// 个性化配置,按业务需要加参数
type SetRequest struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	BgImg                string   `protobuf:"bytes,2,opt,name=bgImg,proto3" json:"bgImg,omitempty"`
	Notification         bool     `protobuf:"varint,3,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64ebe7fa13202e68, []int{1}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SetRequest) GetBgImg() string {
	if m != nil {
		return m.BgImg
	}
	return ""
}

func (m *SetRequest) GetNotification() bool {
	if m != nil {
		return m.Notification
	}
	return false
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_64ebe7fa13202e68, []int{2}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.tenno.ucenter.GetRequest")
	proto.RegisterType((*SetRequest)(nil), "go.micro.srv.tenno.ucenter.SetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.tenno.ucenter.Response")
}

func init() { proto.RegisterFile("proto/ext/ext.proto", fileDescriptor_64ebe7fa13202e68) }

var fileDescriptor_64ebe7fa13202e68 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0x8d, 0xf5, 0xe7, 0xde, 0xc1, 0x85, 0x44, 0x17, 0xe5, 0x22, 0x52, 0x82, 0x48, 0x57,
	0x11, 0xf4, 0x19, 0xa4, 0x74, 0x9b, 0x6e, 0x74, 0x69, 0xeb, 0x34, 0x64, 0xd1, 0x4c, 0x4d, 0xa6,
	0xd2, 0x77, 0xf3, 0xe5, 0x24, 0x15, 0x11, 0x17, 0x16, 0xef, 0x62, 0xe0, 0x9c, 0x99, 0xe1, 0x1b,
	0xce, 0xc0, 0xc5, 0x18, 0x88, 0xe9, 0x0e, 0x67, 0x4e, 0xa5, 0x17, 0x27, 0x77, 0x96, 0xf4, 0xe0,
	0xba, 0x40, 0x3a, 0x86, 0x77, 0xcd, 0xe8, 0x3d, 0xe9, 0xa9, 0x43, 0xcf, 0x18, 0xd4, 0x35, 0x40,
	0x85, 0x6c, 0xf0, 0x6d, 0xc2, 0xc8, 0xf2, 0x1c, 0xb2, 0xc9, 0xbd, 0xe6, 0xa2, 0x10, 0xe5, 0x91,
	0x49, 0x52, 0x3d, 0x01, 0x34, 0x2b, 0x73, 0x79, 0x09, 0xc7, 0xad, 0xad, 0x07, 0x9b, 0x1f, 0x16,
	0xa2, 0xdc, 0x9a, 0x2f, 0x23, 0x15, 0x9c, 0x79, 0x62, 0xd7, 0xbb, 0xee, 0x85, 0x1d, 0xf9, 0x3c,
	0x2b, 0x44, 0xb9, 0x31, 0xbf, 0x7a, 0xea, 0x0a, 0x36, 0x06, 0xe3, 0x48, 0x3e, 0x62, 0xe2, 0x0e,
	0xd1, 0x2e, 0xdc, 0xad, 0x49, 0xf2, 0xfe, 0x43, 0x40, 0xf6, 0x38, 0xb3, 0x7c, 0x86, 0xd3, 0x0a,
	0xb9, 0xf6, 0x3d, 0xc9, 0x5b, 0xfd, 0x77, 0x0e, 0xfd, 0x13, 0x62, 0x77, 0xb3, 0xb6, 0xf7, 0x7d,
	0x52, 0x1d, 0x24, 0x74, 0xf3, 0x1f, 0x74, 0xb3, 0x37, 0xba, 0x3d, 0x59, 0x1e, 0xff, 0xf0, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x47, 0x98, 0x39, 0x8f, 0x01, 0x00, 0x00,
}