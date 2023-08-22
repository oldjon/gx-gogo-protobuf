// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_a_1/m1.proto

package test_a_1

import (
	"database/sql/driver"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type E1 int32

const (
	E1_E1_ZERO E1 = 0
)

var E1_name = map[int32]string{
	0: "E1_ZERO",
}

var E1_value = map[string]int32{
	"E1_ZERO": 0,
}

func (x E1) String() string {
	return proto.EnumName(E1_name, int32(x))
}

func (E1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c1091de3fa870a14, []int{0}
}

type M1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M1) Reset()         { *m = M1{} }
func (m *M1) String() string { return proto.CompactTextString(m) }
func (*M1) ProtoMessage()    {}
func (*M1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1091de3fa870a14, []int{0}
}
func (m *M1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M1.Unmarshal(m, b)
}
func (m *M1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M1.Marshal(b, m, deterministic)
}
func (m *M1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M1.Merge(m, src)
}
func (m *M1) XXX_Size() int {
	return xxx_messageInfo_M1.Size(m)
}
func (m *M1) XXX_DiscardUnknown() {
	xxx_messageInfo_M1.DiscardUnknown(m)
}
func (m *M1) Scan(src interface{}) error {
	if p == nil {
		return errors.New("M1 is nil")
	}
	var source []byte
	switch v := src.(type) {
	case []byte:
		source = src.([]byte)
	case nil:
		return nil
	default:
		return fmt.Errorf("incompatible type for M1 %T", v)
	}
	return m.Unmarshal(source)
}
func (m *M1) Value() (driver.Value, error) {
	return m.Marshal()
}

var xxx_messageInfo_M1 proto.InternalMessageInfo

type M1_1 struct {
	M1                   *M1      `protobuf:"bytes,1,opt,name=m1,proto3" json:"m1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M1_1) Reset()         { *m = M1_1{} }
func (m *M1_1) String() string { return proto.CompactTextString(m) }
func (*M1_1) ProtoMessage()    {}
func (*M1_1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1091de3fa870a14, []int{1}
}
func (m *M1_1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M1_1.Unmarshal(m, b)
}
func (m *M1_1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M1_1.Marshal(b, m, deterministic)
}
func (m *M1_1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M1_1.Merge(m, src)
}
func (m *M1_1) XXX_Size() int {
	return xxx_messageInfo_M1_1.Size(m)
}
func (m *M1_1) XXX_DiscardUnknown() {
	xxx_messageInfo_M1_1.DiscardUnknown(m)
}
func (m *M1_1) Scan(src interface{}) error {
	if p == nil {
		return errors.New("M1_1 is nil")
	}
	var source []byte
	switch v := src.(type) {
	case []byte:
		source = src.([]byte)
	case nil:
		return nil
	default:
		return fmt.Errorf("incompatible type for M1_1 %T", v)
	}
	return m.Unmarshal(source)
}
func (m *M1_1) Value() (driver.Value, error) {
	return m.Marshal()
}

var xxx_messageInfo_M1_1 proto.InternalMessageInfo

func (m *M1_1) GetM1() *M1 {
	if m != nil {
		return m.M1
	}
	return nil
}

func init() {
	proto.RegisterEnum("test.a.E1", E1_name, E1_value)
	proto.RegisterType((*M1)(nil), "test.a.M1")
	proto.RegisterType((*M1_1)(nil), "test.a.M1_1")
}

func init() { proto.RegisterFile("imports/test_a_1/m1.proto", fileDescriptor_c1091de3fa870a14) }

var fileDescriptor_c1091de3fa870a14 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8c, 0x37, 0xd4, 0xcf, 0x35, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x09, 0xe9, 0x25, 0x2a, 0xb1, 0x70, 0x31, 0xf9,
	0x1a, 0x2a, 0x29, 0x71, 0xb1, 0xf8, 0x1a, 0xc6, 0x1b, 0x0a, 0x49, 0x71, 0x31, 0xe5, 0x1a, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xe9, 0x41, 0x94, 0xe8, 0xf9, 0x1a, 0x06, 0x31, 0xe5,
	0x1a, 0x6a, 0x09, 0x72, 0x31, 0xb9, 0x1a, 0x0a, 0x71, 0x73, 0xb1, 0xbb, 0x1a, 0xc6, 0x47, 0xb9,
	0x06, 0xf9, 0x0b, 0x30, 0x38, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x4d, 0x4f, 0x2a, 0x4d, 0x83, 0x30, 0x92,
	0x75, 0xd3, 0x53, 0xf3, 0x74, 0xc1, 0x12, 0x20, 0xc3, 0x52, 0x12, 0x4b, 0x12, 0xf5, 0xd1, 0xdd,
	0x94, 0xc4, 0x06, 0x56, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xd5, 0x3e, 0x41, 0xae,
	0x00, 0x00, 0x00,
}
