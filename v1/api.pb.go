// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package v1

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

type PutResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type PutRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ExpiryPolicy         int64    `protobuf:"varint,3,opt,name=expiry_policy,json=expiryPolicy,proto3" json:"expiry_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PutRequest) GetExpiryPolicy() int64 {
	if m != nil {
		return m.ExpiryPolicy
	}
	return 0
}

type LoadResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (m *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(m, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *LoadResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type LoadRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (m *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(m, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

func (m *LoadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RemoveResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

func (m *RemoveResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type RemoveRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type PingRequest struct {
	Cipher               string   `protobuf:"bytes,1,opt,name=cipher,proto3" json:"cipher,omitempty"`
	Plaintext            string   `protobuf:"bytes,2,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetCipher() string {
	if m != nil {
		return m.Cipher
	}
	return ""
}

func (m *PingRequest) GetPlaintext() string {
	if m != nil {
		return m.Plaintext
	}
	return ""
}

func init() {
	proto.RegisterType((*PutResponse)(nil), "v1.put_response")
	proto.RegisterType((*PutRequest)(nil), "v1.put_request")
	proto.RegisterType((*LoadResponse)(nil), "v1.load_response")
	proto.RegisterType((*LoadRequest)(nil), "v1.load_request")
	proto.RegisterType((*RemoveResponse)(nil), "v1.remove_response")
	proto.RegisterType((*RemoveRequest)(nil), "v1.remove_request")
	proto.RegisterType((*PingRequest)(nil), "v1.ping_request")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x69, 0x8a, 0x8b, 0x1d, 0xbb, 0xab, 0x04, 0x91, 0x1e, 0x44, 0x6a, 0xbc, 0xec, 0x49,
	0x10, 0xf1, 0x1f, 0xf8, 0x03, 0xa4, 0x47, 0x2f, 0x4b, 0x5d, 0x07, 0x0d, 0x8d, 0x9d, 0x31, 0x4d,
	0xcb, 0xf6, 0xdf, 0x8b, 0x49, 0xc8, 0xee, 0xa5, 0xb7, 0xcc, 0x37, 0x2f, 0x2f, 0x2f, 0x0f, 0x8a,
	0x96, 0xf5, 0x23, 0x5b, 0x72, 0x24, 0xc5, 0xf4, 0xa4, 0xee, 0xa0, 0xe4, 0xd1, 0xed, 0x2c, 0x0e,
	0x4c, 0xfd, 0x80, 0x72, 0x03, 0x82, 0xba, 0x2a, 0xab, 0xb3, 0xed, 0x79, 0x23, 0xa8, 0x53, 0xef,
	0x70, 0x11, 0xf6, 0xbf, 0x23, 0x0e, 0x4e, 0x5e, 0x41, 0xde, 0xe1, 0xec, 0xf7, 0x45, 0xf3, 0x7f,
	0x94, 0xd7, 0x70, 0x36, 0xb5, 0x66, 0xc4, 0x4a, 0x78, 0x16, 0x06, 0xf9, 0x00, 0x6b, 0x3c, 0xb0,
	0xb6, 0xf3, 0x8e, 0xc9, 0xe8, 0xfd, 0x5c, 0xe5, 0x75, 0xb6, 0xcd, 0x9b, 0x32, 0xc0, 0x37, 0xcf,
	0xd4, 0x0b, 0xac, 0x0d, 0xb5, 0x9f, 0xc7, 0xc7, 0x93, 0x57, 0x76, 0xea, 0x15, 0x22, 0x89, 0x14,
	0xa9, 0x86, 0x32, 0x5e, 0x5b, 0xc8, 0xa4, 0xee, 0xe1, 0xd2, 0xe2, 0x0f, 0x4d, 0xb8, 0xfc, 0x2f,
	0x05, 0x9b, 0x24, 0x59, 0xb2, 0x79, 0x85, 0x92, 0x75, 0xff, 0x95, 0x14, 0x37, 0xb0, 0xda, 0x6b,
	0xfe, 0x46, 0x1b, 0x45, 0x71, 0x92, 0xb7, 0x50, 0xb0, 0x69, 0x75, 0xef, 0xf0, 0xe0, 0x62, 0x0d,
	0x47, 0xf0, 0xb1, 0xf2, 0x65, 0x3f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0xdb, 0x2d, 0x9d,
	0x79, 0x01, 0x00, 0x00,
}