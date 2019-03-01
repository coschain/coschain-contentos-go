// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extDailyTrx.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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

type SoExtDailyTrx struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Count                uint64                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoExtDailyTrx) Reset()         { *m = SoExtDailyTrx{} }
func (m *SoExtDailyTrx) String() string { return proto.CompactTextString(m) }
func (*SoExtDailyTrx) ProtoMessage()    {}
func (*SoExtDailyTrx) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{0}
}

func (m *SoExtDailyTrx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtDailyTrx.Unmarshal(m, b)
}
func (m *SoExtDailyTrx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtDailyTrx.Marshal(b, m, deterministic)
}
func (m *SoExtDailyTrx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtDailyTrx.Merge(m, src)
}
func (m *SoExtDailyTrx) XXX_Size() int {
	return xxx_messageInfo_SoExtDailyTrx.Size(m)
}
func (m *SoExtDailyTrx) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtDailyTrx.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtDailyTrx proto.InternalMessageInfo

func (m *SoExtDailyTrx) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *SoExtDailyTrx) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SoMemExtDailyTrxByDate struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemExtDailyTrxByDate) Reset()         { *m = SoMemExtDailyTrxByDate{} }
func (m *SoMemExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoMemExtDailyTrxByDate) ProtoMessage()    {}
func (*SoMemExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{1}
}

func (m *SoMemExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoMemExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoMemExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtDailyTrxByDate.Merge(m, src)
}
func (m *SoMemExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoMemExtDailyTrxByDate.Size(m)
}
func (m *SoMemExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoMemExtDailyTrxByDate) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

type SoMemExtDailyTrxByCount struct {
	Count                uint64   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtDailyTrxByCount) Reset()         { *m = SoMemExtDailyTrxByCount{} }
func (m *SoMemExtDailyTrxByCount) String() string { return proto.CompactTextString(m) }
func (*SoMemExtDailyTrxByCount) ProtoMessage()    {}
func (*SoMemExtDailyTrxByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{2}
}

func (m *SoMemExtDailyTrxByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtDailyTrxByCount.Unmarshal(m, b)
}
func (m *SoMemExtDailyTrxByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtDailyTrxByCount.Marshal(b, m, deterministic)
}
func (m *SoMemExtDailyTrxByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtDailyTrxByCount.Merge(m, src)
}
func (m *SoMemExtDailyTrxByCount) XXX_Size() int {
	return xxx_messageInfo_SoMemExtDailyTrxByCount.Size(m)
}
func (m *SoMemExtDailyTrxByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtDailyTrxByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtDailyTrxByCount proto.InternalMessageInfo

func (m *SoMemExtDailyTrxByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SoListExtDailyTrxByDate struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtDailyTrxByDate) Reset()         { *m = SoListExtDailyTrxByDate{} }
func (m *SoListExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoListExtDailyTrxByDate) ProtoMessage()    {}
func (*SoListExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{3}
}

func (m *SoListExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoListExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoListExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtDailyTrxByDate.Merge(m, src)
}
func (m *SoListExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoListExtDailyTrxByDate.Size(m)
}
func (m *SoListExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoListExtDailyTrxByDate) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

type SoListExtDailyTrxByCount struct {
	Count                uint64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Date                 *prototype.TimePointSec `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListExtDailyTrxByCount) Reset()         { *m = SoListExtDailyTrxByCount{} }
func (m *SoListExtDailyTrxByCount) String() string { return proto.CompactTextString(m) }
func (*SoListExtDailyTrxByCount) ProtoMessage()    {}
func (*SoListExtDailyTrxByCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{4}
}

func (m *SoListExtDailyTrxByCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Unmarshal(m, b)
}
func (m *SoListExtDailyTrxByCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Marshal(b, m, deterministic)
}
func (m *SoListExtDailyTrxByCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtDailyTrxByCount.Merge(m, src)
}
func (m *SoListExtDailyTrxByCount) XXX_Size() int {
	return xxx_messageInfo_SoListExtDailyTrxByCount.Size(m)
}
func (m *SoListExtDailyTrxByCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtDailyTrxByCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtDailyTrxByCount proto.InternalMessageInfo

func (m *SoListExtDailyTrxByCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SoListExtDailyTrxByCount) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

type SoUniqueExtDailyTrxByDate struct {
	Date                 *prototype.TimePointSec `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoUniqueExtDailyTrxByDate) Reset()         { *m = SoUniqueExtDailyTrxByDate{} }
func (m *SoUniqueExtDailyTrxByDate) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtDailyTrxByDate) ProtoMessage()    {}
func (*SoUniqueExtDailyTrxByDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71b07e7e9721030, []int{5}
}

func (m *SoUniqueExtDailyTrxByDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Unmarshal(m, b)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Marshal(b, m, deterministic)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtDailyTrxByDate.Merge(m, src)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtDailyTrxByDate.Size(m)
}
func (m *SoUniqueExtDailyTrxByDate) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtDailyTrxByDate.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtDailyTrxByDate proto.InternalMessageInfo

func (m *SoUniqueExtDailyTrxByDate) GetDate() *prototype.TimePointSec {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterType((*SoExtDailyTrx)(nil), "table.so_extDailyTrx")
	proto.RegisterType((*SoMemExtDailyTrxByDate)(nil), "table.so_mem_extDailyTrx_by_date")
	proto.RegisterType((*SoMemExtDailyTrxByCount)(nil), "table.so_mem_extDailyTrx_by_count")
	proto.RegisterType((*SoListExtDailyTrxByDate)(nil), "table.so_list_extDailyTrx_by_date")
	proto.RegisterType((*SoListExtDailyTrxByCount)(nil), "table.so_list_extDailyTrx_by_count")
	proto.RegisterType((*SoUniqueExtDailyTrxByDate)(nil), "table.so_unique_extDailyTrx_by_date")
}

func init() { proto.RegisterFile("app/table/so_extDailyTrx.proto", fileDescriptor_a71b07e7e9721030) }

var fileDescriptor_a71b07e7e9721030 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xd2, 0x7a, 0x18, 0xc1, 0xc3, 0xd2, 0x43, 0xad, 0x1f, 0x94, 0x3d, 0xc8, 0x5e,
	0xba, 0x01, 0xfb, 0x0f, 0xc4, 0x9b, 0xe2, 0xa1, 0xe8, 0xc5, 0x4b, 0x48, 0xe2, 0xd0, 0x06, 0x36,
	0x99, 0xd8, 0x4c, 0xa0, 0xfb, 0xef, 0xc5, 0x2c, 0xd4, 0x0a, 0x2a, 0x16, 0x7a, 0x09, 0x0c, 0x93,
	0xf7, 0x7d, 0x9e, 0xc3, 0xc0, 0xb5, 0x0a, 0x41, 0xb0, 0xd2, 0x2d, 0x8a, 0x48, 0x12, 0xb7, 0x7c,
	0xaf, 0x6c, 0xdb, 0x3d, 0x6f, 0xb6, 0x4d, 0xd8, 0x10, 0x53, 0x39, 0xca, 0xbb, 0xe9, 0x38, 0x4f,
	0xdc, 0x05, 0x14, 0x9f, 0x4f, 0xbf, 0xac, 0x5e, 0xe0, 0xec, 0x7b, 0xa8, 0x9c, 0xc3, 0xf0, 0x4d,
	0x31, 0x4e, 0x8a, 0x59, 0x51, 0x9f, 0xde, 0x9e, 0x37, 0xbb, 0x58, 0xc3, 0xd6, 0xa1, 0x0c, 0x64,
	0x3d, 0xcb, 0x88, 0x66, 0x99, 0xbf, 0x95, 0x63, 0x18, 0x19, 0x4a, 0x9e, 0x27, 0x83, 0x59, 0x51,
	0x0f, 0x97, 0xfd, 0x50, 0x3d, 0xc0, 0x34, 0x92, 0x74, 0xe8, 0xf6, 0xab, 0xa5, 0xee, 0x64, 0xce,
	0x1c, 0x86, 0xa8, 0x16, 0x70, 0xf1, 0x73, 0x59, 0x66, 0x7d, 0x19, 0x14, 0xfb, 0x06, 0x8f, 0x39,
	0xd4, 0xda, 0xc8, 0xc7, 0x50, 0x30, 0x70, 0xf9, 0x4b, 0xdb, 0x1f, 0x0e, 0x3b, 0xc8, 0xe0, 0x7f,
	0x90, 0x27, 0xb8, 0x8a, 0x24, 0x93, 0xb7, 0xef, 0x09, 0x8f, 0x20, 0x7d, 0x57, 0xbf, 0xde, 0xac,
	0x2c, 0xaf, 0x93, 0x6e, 0x0c, 0x39, 0x61, 0x28, 0x9a, 0xb5, 0xb2, 0x5e, 0x18, 0xf2, 0x8c, 0x9e,
	0x29, 0xce, 0x57, 0xd4, 0xdf, 0x8d, 0x3e, 0xc9, 0x4d, 0x8b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0x10, 0x6a, 0x60, 0x4b, 0x02, 0x00, 0x00,
}