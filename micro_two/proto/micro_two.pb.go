// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/micro_two.proto

package micro_two

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

type CallRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{0}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{1}
}

func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (m *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(m, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ClientStreamRequest struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStreamRequest) Reset()         { *m = ClientStreamRequest{} }
func (m *ClientStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStreamRequest) ProtoMessage()    {}
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{2}
}

func (m *ClientStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStreamRequest.Unmarshal(m, b)
}
func (m *ClientStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStreamRequest.Marshal(b, m, deterministic)
}
func (m *ClientStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStreamRequest.Merge(m, src)
}
func (m *ClientStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStreamRequest.Size(m)
}
func (m *ClientStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStreamRequest proto.InternalMessageInfo

func (m *ClientStreamRequest) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type ClientStreamResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStreamResponse) Reset()         { *m = ClientStreamResponse{} }
func (m *ClientStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStreamResponse) ProtoMessage()    {}
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{3}
}

func (m *ClientStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStreamResponse.Unmarshal(m, b)
}
func (m *ClientStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStreamResponse.Marshal(b, m, deterministic)
}
func (m *ClientStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStreamResponse.Merge(m, src)
}
func (m *ClientStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStreamResponse.Size(m)
}
func (m *ClientStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStreamResponse proto.InternalMessageInfo

func (m *ClientStreamResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ServerStreamRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamRequest) Reset()         { *m = ServerStreamRequest{} }
func (m *ServerStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ServerStreamRequest) ProtoMessage()    {}
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{4}
}

func (m *ServerStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamRequest.Unmarshal(m, b)
}
func (m *ServerStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamRequest.Marshal(b, m, deterministic)
}
func (m *ServerStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamRequest.Merge(m, src)
}
func (m *ServerStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ServerStreamRequest.Size(m)
}
func (m *ServerStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamRequest proto.InternalMessageInfo

func (m *ServerStreamRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ServerStreamResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamResponse) Reset()         { *m = ServerStreamResponse{} }
func (m *ServerStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ServerStreamResponse) ProtoMessage()    {}
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{5}
}

func (m *ServerStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamResponse.Unmarshal(m, b)
}
func (m *ServerStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamResponse.Marshal(b, m, deterministic)
}
func (m *ServerStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamResponse.Merge(m, src)
}
func (m *ServerStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ServerStreamResponse.Size(m)
}
func (m *ServerStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamResponse proto.InternalMessageInfo

func (m *ServerStreamResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type BidiStreamRequest struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidiStreamRequest) Reset()         { *m = BidiStreamRequest{} }
func (m *BidiStreamRequest) String() string { return proto.CompactTextString(m) }
func (*BidiStreamRequest) ProtoMessage()    {}
func (*BidiStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{6}
}

func (m *BidiStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidiStreamRequest.Unmarshal(m, b)
}
func (m *BidiStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidiStreamRequest.Marshal(b, m, deterministic)
}
func (m *BidiStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidiStreamRequest.Merge(m, src)
}
func (m *BidiStreamRequest) XXX_Size() int {
	return xxx_messageInfo_BidiStreamRequest.Size(m)
}
func (m *BidiStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidiStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidiStreamRequest proto.InternalMessageInfo

func (m *BidiStreamRequest) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type BidiStreamResponse struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidiStreamResponse) Reset()         { *m = BidiStreamResponse{} }
func (m *BidiStreamResponse) String() string { return proto.CompactTextString(m) }
func (*BidiStreamResponse) ProtoMessage()    {}
func (*BidiStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a4d5d6538f1b68f, []int{7}
}

func (m *BidiStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidiStreamResponse.Unmarshal(m, b)
}
func (m *BidiStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidiStreamResponse.Marshal(b, m, deterministic)
}
func (m *BidiStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidiStreamResponse.Merge(m, src)
}
func (m *BidiStreamResponse) XXX_Size() int {
	return xxx_messageInfo_BidiStreamResponse.Size(m)
}
func (m *BidiStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BidiStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BidiStreamResponse proto.InternalMessageInfo

func (m *BidiStreamResponse) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*CallRequest)(nil), "micro_two.CallRequest")
	proto.RegisterType((*CallResponse)(nil), "micro_two.CallResponse")
	proto.RegisterType((*ClientStreamRequest)(nil), "micro_two.ClientStreamRequest")
	proto.RegisterType((*ClientStreamResponse)(nil), "micro_two.ClientStreamResponse")
	proto.RegisterType((*ServerStreamRequest)(nil), "micro_two.ServerStreamRequest")
	proto.RegisterType((*ServerStreamResponse)(nil), "micro_two.ServerStreamResponse")
	proto.RegisterType((*BidiStreamRequest)(nil), "micro_two.BidiStreamRequest")
	proto.RegisterType((*BidiStreamResponse)(nil), "micro_two.BidiStreamResponse")
}

func init() { proto.RegisterFile("proto/micro_two.proto", fileDescriptor_7a4d5d6538f1b68f) }

var fileDescriptor_7a4d5d6538f1b68f = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xbb, 0x30,
	0x18, 0x80, 0xc7, 0x6f, 0xfb, 0x2d, 0xe1, 0x75, 0x07, 0xf7, 0x82, 0xd3, 0x10, 0xff, 0xcc, 0x9e,
	0x96, 0x38, 0xd9, 0xa2, 0x27, 0xe3, 0x6d, 0x3b, 0x1b, 0x13, 0xb8, 0x79, 0x31, 0x38, 0x1b, 0x43,
	0x04, 0x3a, 0xdb, 0x4e, 0x3f, 0x9f, 0xdf, 0xcc, 0x00, 0x65, 0xb4, 0x01, 0xa2, 0xb7, 0xbe, 0x6f,
	0x9f, 0x3d, 0x69, 0x9e, 0x01, 0x47, 0x5b, 0xce, 0x24, 0x5b, 0xa4, 0xf1, 0x86, 0xb3, 0x67, 0xf9,
	0xc5, 0xfc, 0x62, 0x46, 0x7b, 0xbf, 0x20, 0x97, 0x70, 0xb0, 0x8e, 0x92, 0x24, 0xa0, 0x1f, 0x3b,
	0x2a, 0x24, 0x22, 0x0c, 0xb2, 0x28, 0xa5, 0x27, 0xd6, 0xd4, 0x9a, 0xd9, 0x41, 0x71, 0x26, 0x53,
	0x18, 0x95, 0x88, 0xd8, 0xb2, 0x4c, 0x50, 0x3c, 0x84, 0x7e, 0x2a, 0xde, 0x14, 0x92, 0x1f, 0xc9,
	0x35, 0x38, 0xeb, 0x24, 0xa6, 0x99, 0x0c, 0x25, 0xa7, 0x51, 0x5a, 0xc9, 0x26, 0x30, 0x14, 0x92,
	0xb3, 0xf7, 0x52, 0xd7, 0x0f, 0xd4, 0x44, 0xe6, 0xe0, 0x9a, 0xb8, 0x12, 0xbb, 0xf0, 0x7f, 0xc3,
	0x76, 0x99, 0x54, 0x78, 0x39, 0x90, 0x2b, 0x70, 0x42, 0xca, 0x3f, 0x29, 0x37, 0xe5, 0xed, 0xf0,
	0x1c, 0x5c, 0x13, 0xfe, 0x45, 0x3d, 0x5e, 0xc5, 0xaf, 0xf1, 0x5f, 0x5f, 0x8d, 0x3a, 0xac, 0xc4,
	0x1d, 0xf4, 0xcd, 0xf7, 0x3f, 0xb0, 0x1f, 0xaa, 0xca, 0x78, 0x07, 0x83, 0x3c, 0x21, 0x4e, 0xfc,
	0xfa, 0xaf, 0xd0, 0xb2, 0x7b, 0xc7, 0x8d, 0x7d, 0xa9, 0x27, 0x3d, 0x0c, 0x61, 0xa4, 0xc7, 0xc2,
	0x73, 0x1d, 0x6d, 0x46, 0xf7, 0x2e, 0x3a, 0xef, 0x2b, 0xe5, 0xcc, 0xca, 0xa5, 0x7a, 0x26, 0x43,
	0xda, 0x12, 0xdb, 0x90, 0xb6, 0xf5, 0x25, 0xbd, 0xa5, 0x85, 0x8f, 0x00, 0x75, 0x20, 0x3c, 0xd5,
	0x7e, 0xd2, 0x88, 0xec, 0x9d, 0x75, 0xdc, 0xd6, 0x6f, 0x5c, 0x5a, 0x2b, 0xe7, 0x69, 0xec, 0x2f,
	0x8a, 0x2f, 0xf6, 0x7e, 0x4f, 0xbf, 0x0c, 0x8b, 0xc5, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0xf9, 0x94, 0x98, 0xdb, 0x02, 0x00, 0x00,
}
