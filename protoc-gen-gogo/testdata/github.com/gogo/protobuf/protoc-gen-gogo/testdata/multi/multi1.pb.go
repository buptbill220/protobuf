// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multi/multi1.proto

/*
Package multitest is a generated protocol buffer package.

It is generated from these files:
	multi/multi1.proto

It has these top-level messages:
	Multi1
*/
package multitest

import proto "github.com/buptbill220/protobuf/proto"
import fmt "fmt"
import math "math"
import multitest1 "github.com/buptbill220/protobuf/protoc-gen-gogo/testdata/multi"
import multitest2 "github.com/buptbill220/protobuf/protoc-gen-gogo/testdata/multi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Multi1 struct {
	Multi2           *multitest1.Multi2         `protobuf:"bytes,1,req,name=multi2" json:"multi2,omitempty"`
	Color            *multitest1.Multi2_Color   `protobuf:"varint,2,opt,name=color,enum=multitest.Multi2_Color" json:"color,omitempty"`
	HatType          *multitest2.Multi3_HatType `protobuf:"varint,3,opt,name=hat_type,enum=multitest.Multi3_HatType" json:"hat_type,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Multi1) Reset()                    { *m = Multi1{} }
func (m *Multi1) String() string            { return proto.CompactTextString(m) }
func (*Multi1) ProtoMessage()               {}
func (*Multi1) Descriptor() ([]byte, []int) { return fileDescriptorMulti1, []int{0} }

func (m *Multi1) GetMulti2() *multitest1.Multi2 {
	if m != nil {
		return m.Multi2
	}
	return nil
}

func (m *Multi1) GetColor() multitest1.Multi2_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return multitest1.Multi2_BLUE
}

func (m *Multi1) GetHatType() multitest2.Multi3_HatType {
	if m != nil && m.HatType != nil {
		return *m.HatType
	}
	return multitest2.Multi3_FEDORA
}

func init() {
	proto.RegisterType((*Multi1)(nil), "multitest.Multi1")
}

func init() { proto.RegisterFile("multi/multi1.proto", fileDescriptorMulti1) }

var fileDescriptorMulti1 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x2d, 0xcd, 0x29,
	0xc9, 0xd4, 0x07, 0x93, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x60, 0x5e, 0x49,
	0x6a, 0x71, 0x89, 0x14, 0xb2, 0xb4, 0x11, 0x44, 0x1a, 0x45, 0xcc, 0x18, 0x22, 0xa6, 0xd4, 0xc0,
	0xc8, 0xc5, 0xe6, 0x0b, 0x36, 0x43, 0x48, 0x91, 0x8b, 0x0d, 0xa2, 0x5c, 0x82, 0x51, 0x81, 0x49,
	0x83, 0xdb, 0x48, 0x50, 0x0f, 0x6e, 0x9c, 0x1e, 0x58, 0x89, 0x91, 0x90, 0x1a, 0x17, 0x6b, 0x72,
	0x7e, 0x4e, 0x7e, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x38, 0x86, 0x0a, 0x3d, 0x67,
	0x90, 0xb4, 0x90, 0x36, 0x17, 0x47, 0x46, 0x62, 0x49, 0x7c, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x33,
	0x58, 0xa9, 0x24, 0xba, 0x52, 0x63, 0x3d, 0x8f, 0xc4, 0x92, 0x90, 0xca, 0x82, 0x54, 0x27, 0xe7,
	0x28, 0xc7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf4, 0xfc, 0xf4,
	0x7c, 0x7d, 0xb0, 0xdb, 0x92, 0x4a, 0xd3, 0x20, 0x8c, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0x5d, 0xb0,
	0x04, 0x48, 0x7f, 0x4a, 0x62, 0x49, 0x22, 0xc4, 0x13, 0xd6, 0x70, 0x33, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0x12, 0x3b, 0xd5, 0x0f, 0x01, 0x00, 0x00,
}
