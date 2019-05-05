// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: events.proto

package events

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/buptbill220/protobuf/gogoproto"
	proto "github.com/buptbill220/protobuf/proto"
	issue312 "github.com/buptbill220/protobuf/test/issue312"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Subtype struct {
	State                *issue312.TaskState `protobuf:"varint,4,opt,name=state,enum=issue312.TaskState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Subtype) Reset()         { *m = Subtype{} }
func (m *Subtype) String() string { return proto.CompactTextString(m) }
func (*Subtype) ProtoMessage()    {}
func (*Subtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}
func (m *Subtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subtype.Unmarshal(m, b)
}
func (m *Subtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subtype.Marshal(b, m, deterministic)
}
func (m *Subtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subtype.Merge(m, src)
}
func (m *Subtype) XXX_Size() int {
	return xxx_messageInfo_Subtype.Size(m)
}
func (m *Subtype) XXX_DiscardUnknown() {
	xxx_messageInfo_Subtype.DiscardUnknown(m)
}

var xxx_messageInfo_Subtype proto.InternalMessageInfo

func (m *Subtype) GetState() issue312.TaskState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return issue312.TaskState_TASK_STAGING
}

func init() {
	proto.RegisterType((*Subtype)(nil), "issue312.events.Subtype")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x4b, 0xcd,
	0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0x2c, 0x2e, 0x2e, 0x4d, 0x35,
	0x36, 0x34, 0xd2, 0x83, 0x08, 0x4b, 0x99, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0xd5, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0xa4, 0x16,
	0x97, 0xe8, 0xc3, 0x94, 0xc3, 0x19, 0x10, 0x73, 0xa4, 0x74, 0x71, 0x6a, 0x03, 0xf1, 0xc0, 0x1c,
	0x30, 0x0b, 0xa2, 0x5c, 0xc9, 0x84, 0x8b, 0x3d, 0xb8, 0x34, 0xa9, 0xa4, 0xb2, 0x20, 0x55, 0x48,
	0x93, 0x8b, 0xb5, 0xb8, 0x24, 0xb1, 0x24, 0x55, 0x82, 0x45, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x58,
	0x0f, 0x6e, 0x72, 0x48, 0x62, 0x71, 0x76, 0x30, 0x48, 0x2a, 0x08, 0xa2, 0xc2, 0x49, 0xe2, 0xc3,
	0x43, 0x39, 0xc6, 0x1f, 0x0f, 0xe5, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc, 0xf1, 0x48, 0x8e, 0x31,
	0x8a, 0x0d, 0xe2, 0x6a, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x76, 0xdc, 0x82, 0xd5, 0x00,
	0x00, 0x00,
}

func (this *Subtype) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Subtype)
	if !ok {
		that2, ok := that.(Subtype)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.State != nil && that1.State != nil {
		if *this.State != *that1.State {
			return false
		}
	} else if this.State != nil {
		return false
	} else if that1.State != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Subtype) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&events.Subtype{")
	if this.State != nil {
		s = append(s, "State: "+valueToGoStringEvents(this.State, "issue312.TaskState")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEvents(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedSubtype(r randyEvents, easy bool) *Subtype {
	this := &Subtype{}
	if r.Intn(10) != 0 {
		v1 := issue312.TaskState([]int32{6, 0, 1}[r.Intn(3)])
		this.State = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEvents(r, 5)
	}
	return this
}

type randyEvents interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEvents(r randyEvents) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEvents(r randyEvents) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneEvents(r)
	}
	return string(tmps)
}
func randUnrecognizedEvents(r randyEvents, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEvents(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEvents(dAtA []byte, r randyEvents, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEvents(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateEvents(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateEvents(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEvents(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEvents(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEvents(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEvents(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
