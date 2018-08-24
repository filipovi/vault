// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/generator.proto

package generator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NewPasswordRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passphrase           string   `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	Length               int32    `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	Counter              int32    `protobuf:"varint,5,opt,name=counter,proto3" json:"counter,omitempty"`
	Scope                string   `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPasswordRequest) Reset()         { *m = NewPasswordRequest{} }
func (m *NewPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*NewPasswordRequest) ProtoMessage()    {}
func (*NewPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_generator_2f0bef73dc81de94, []int{0}
}
func (m *NewPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPasswordRequest.Unmarshal(m, b)
}
func (m *NewPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *NewPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPasswordRequest.Merge(dst, src)
}
func (m *NewPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_NewPasswordRequest.Size(m)
}
func (m *NewPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewPasswordRequest proto.InternalMessageInfo

func (m *NewPasswordRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewPasswordRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *NewPasswordRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *NewPasswordRequest) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *NewPasswordRequest) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *NewPasswordRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

type NewPasswordResponse struct {
	Password             string   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPasswordResponse) Reset()         { *m = NewPasswordResponse{} }
func (m *NewPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*NewPasswordResponse) ProtoMessage()    {}
func (*NewPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_generator_2f0bef73dc81de94, []int{1}
}
func (m *NewPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPasswordResponse.Unmarshal(m, b)
}
func (m *NewPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPasswordResponse.Marshal(b, m, deterministic)
}
func (dst *NewPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPasswordResponse.Merge(dst, src)
}
func (m *NewPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_NewPasswordResponse.Size(m)
}
func (m *NewPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewPasswordResponse proto.InternalMessageInfo

func (m *NewPasswordResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*NewPasswordRequest)(nil), "NewPasswordRequest")
	proto.RegisterType((*NewPasswordResponse)(nil), "NewPasswordResponse")
}

func init() { proto.RegisterFile("proto/generator.proto", fileDescriptor_generator_2f0bef73dc81de94) }

var fileDescriptor_generator_2f0bef73dc81de94 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4e, 0xc4, 0x30,
	0x0c, 0x45, 0x09, 0x4c, 0x3b, 0x8c, 0xd9, 0x79, 0x06, 0x64, 0xcd, 0x02, 0x8d, 0xba, 0xea, 0xaa,
	0x08, 0xd8, 0x71, 0x81, 0xee, 0x10, 0xea, 0x0d, 0x42, 0xb1, 0x5a, 0x24, 0x48, 0x42, 0x9c, 0xd2,
	0x3b, 0x71, 0x4a, 0x84, 0x69, 0x51, 0x11, 0xec, 0xf2, 0xfe, 0xcf, 0x4f, 0xfc, 0x0d, 0xe7, 0x21,
	0xfa, 0xe4, 0xaf, 0x3a, 0x76, 0x1c, 0x6d, 0xf2, 0xb1, 0x52, 0x2e, 0x3e, 0x0c, 0xe0, 0x3d, 0x8f,
	0x0f, 0x56, 0x64, 0xf4, 0xf1, 0xa9, 0xe1, 0xb7, 0x81, 0x25, 0x21, 0xc2, 0xca, 0xd9, 0x57, 0x26,
	0x73, 0x30, 0xe5, 0xa6, 0xd1, 0x33, 0x5e, 0x02, 0x04, 0x2b, 0x12, 0xfa, 0x68, 0x85, 0xe9, 0x58,
	0x9d, 0x85, 0x82, 0x04, 0x6b, 0xe1, 0xf8, 0xfe, 0xdc, 0x32, 0x9d, 0xa8, 0x39, 0x23, 0x5e, 0x40,
	0xfe, 0xc2, 0xae, 0x4b, 0x3d, 0xad, 0x0e, 0xa6, 0xcc, 0x9a, 0x89, 0xbe, 0x12, 0xad, 0x1f, 0x5c,
	0xe2, 0x48, 0x99, 0x1a, 0x33, 0xe2, 0x0e, 0x32, 0x69, 0x7d, 0x60, 0xca, 0xf5, 0xa5, 0x6f, 0x28,
	0xae, 0x61, 0xfb, 0x6b, 0x56, 0x09, 0xde, 0x09, 0xe3, 0x1e, 0x4e, 0xc3, 0xa4, 0xd1, 0x5a, 0xef,
	0xff, 0xf0, 0x4d, 0x0d, 0x9b, 0x7a, 0xae, 0x8c, 0x77, 0x70, 0xb6, 0xc8, 0xe3, 0xb6, 0xfa, 0xdb,
	0x7c, 0xbf, 0xab, 0xfe, 0xf9, 0xa2, 0x38, 0x7a, 0xcc, 0x75, 0x5f, 0xb7, 0x9f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0xe8, 0xc0, 0x26, 0x48, 0x01, 0x00, 0x00,
}