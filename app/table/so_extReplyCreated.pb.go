// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_extReplyCreated.proto

package table // import "github.com/coschain/contentos-go/app/table"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototype "github.com/coschain/contentos-go/prototype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoExtReplyCreated struct {
	PostId               uint64                       `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CreatedOrder         *prototype.ReplyCreatedOrder `protobuf:"bytes,2,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoExtReplyCreated) Reset()         { *m = SoExtReplyCreated{} }
func (m *SoExtReplyCreated) String() string { return proto.CompactTextString(m) }
func (*SoExtReplyCreated) ProtoMessage()    {}
func (*SoExtReplyCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_8c7549b7629a26a0, []int{0}
}
func (m *SoExtReplyCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoExtReplyCreated.Unmarshal(m, b)
}
func (m *SoExtReplyCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoExtReplyCreated.Marshal(b, m, deterministic)
}
func (dst *SoExtReplyCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoExtReplyCreated.Merge(dst, src)
}
func (m *SoExtReplyCreated) XXX_Size() int {
	return xxx_messageInfo_SoExtReplyCreated.Size(m)
}
func (m *SoExtReplyCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoExtReplyCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoExtReplyCreated proto.InternalMessageInfo

func (m *SoExtReplyCreated) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *SoExtReplyCreated) GetCreatedOrder() *prototype.ReplyCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

type SoMemExtReplyCreatedByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemExtReplyCreatedByPostId) Reset()         { *m = SoMemExtReplyCreatedByPostId{} }
func (m *SoMemExtReplyCreatedByPostId) String() string { return proto.CompactTextString(m) }
func (*SoMemExtReplyCreatedByPostId) ProtoMessage()    {}
func (*SoMemExtReplyCreatedByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_8c7549b7629a26a0, []int{1}
}
func (m *SoMemExtReplyCreatedByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtReplyCreatedByPostId.Unmarshal(m, b)
}
func (m *SoMemExtReplyCreatedByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtReplyCreatedByPostId.Marshal(b, m, deterministic)
}
func (dst *SoMemExtReplyCreatedByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtReplyCreatedByPostId.Merge(dst, src)
}
func (m *SoMemExtReplyCreatedByPostId) XXX_Size() int {
	return xxx_messageInfo_SoMemExtReplyCreatedByPostId.Size(m)
}
func (m *SoMemExtReplyCreatedByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtReplyCreatedByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtReplyCreatedByPostId proto.InternalMessageInfo

func (m *SoMemExtReplyCreatedByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoMemExtReplyCreatedByCreatedOrder struct {
	CreatedOrder         *prototype.ReplyCreatedOrder `protobuf:"bytes,1,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoMemExtReplyCreatedByCreatedOrder) Reset()         { *m = SoMemExtReplyCreatedByCreatedOrder{} }
func (m *SoMemExtReplyCreatedByCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoMemExtReplyCreatedByCreatedOrder) ProtoMessage()    {}
func (*SoMemExtReplyCreatedByCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_8c7549b7629a26a0, []int{2}
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Unmarshal(m, b)
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Marshal(b, m, deterministic)
}
func (dst *SoMemExtReplyCreatedByCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Merge(dst, src)
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.Size(m)
}
func (m *SoMemExtReplyCreatedByCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemExtReplyCreatedByCreatedOrder proto.InternalMessageInfo

func (m *SoMemExtReplyCreatedByCreatedOrder) GetCreatedOrder() *prototype.ReplyCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

type SoListExtReplyCreatedByCreatedOrder struct {
	CreatedOrder         *prototype.ReplyCreatedOrder `protobuf:"bytes,1,opt,name=created_order,json=createdOrder,proto3" json:"created_order,omitempty"`
	PostId               uint64                       `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SoListExtReplyCreatedByCreatedOrder) Reset()         { *m = SoListExtReplyCreatedByCreatedOrder{} }
func (m *SoListExtReplyCreatedByCreatedOrder) String() string { return proto.CompactTextString(m) }
func (*SoListExtReplyCreatedByCreatedOrder) ProtoMessage()    {}
func (*SoListExtReplyCreatedByCreatedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_8c7549b7629a26a0, []int{3}
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Unmarshal(m, b)
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Marshal(b, m, deterministic)
}
func (dst *SoListExtReplyCreatedByCreatedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Merge(dst, src)
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_Size() int {
	return xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.Size(m)
}
func (m *SoListExtReplyCreatedByCreatedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SoListExtReplyCreatedByCreatedOrder proto.InternalMessageInfo

func (m *SoListExtReplyCreatedByCreatedOrder) GetCreatedOrder() *prototype.ReplyCreatedOrder {
	if m != nil {
		return m.CreatedOrder
	}
	return nil
}

func (m *SoListExtReplyCreatedByCreatedOrder) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoUniqueExtReplyCreatedByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueExtReplyCreatedByPostId) Reset()         { *m = SoUniqueExtReplyCreatedByPostId{} }
func (m *SoUniqueExtReplyCreatedByPostId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueExtReplyCreatedByPostId) ProtoMessage()    {}
func (*SoUniqueExtReplyCreatedByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_so_extReplyCreated_8c7549b7629a26a0, []int{4}
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Unmarshal(m, b)
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Marshal(b, m, deterministic)
}
func (dst *SoUniqueExtReplyCreatedByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Merge(dst, src)
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.Size(m)
}
func (m *SoUniqueExtReplyCreatedByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueExtReplyCreatedByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueExtReplyCreatedByPostId proto.InternalMessageInfo

func (m *SoUniqueExtReplyCreatedByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SoExtReplyCreated)(nil), "table.so_extReplyCreated")
	proto.RegisterType((*SoMemExtReplyCreatedByPostId)(nil), "table.so_mem_extReplyCreated_by_post_id")
	proto.RegisterType((*SoMemExtReplyCreatedByCreatedOrder)(nil), "table.so_mem_extReplyCreated_by_created_order")
	proto.RegisterType((*SoListExtReplyCreatedByCreatedOrder)(nil), "table.so_list_extReplyCreated_by_created_order")
	proto.RegisterType((*SoUniqueExtReplyCreatedByPostId)(nil), "table.so_unique_extReplyCreated_by_post_id")
}

func init() {
	proto.RegisterFile("app/table/so_extReplyCreated.proto", fileDescriptor_so_extReplyCreated_8c7549b7629a26a0)
}

var fileDescriptor_so_extReplyCreated_8c7549b7629a26a0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0xa2, 0x2b, 0x44, 0xbd, 0xf4, 0x62, 0xf1, 0x20, 0x6b, 0x11, 0x2c, 0xa2, 0x0d,
	0xe8, 0x55, 0x10, 0xdc, 0x93, 0x27, 0xa1, 0x47, 0x2f, 0x43, 0x9b, 0x0e, 0xbb, 0x81, 0x36, 0x13,
	0x93, 0x29, 0xd8, 0x37, 0xf0, 0xb1, 0xa5, 0xb5, 0x2c, 0xd6, 0x65, 0x2f, 0x7b, 0xd8, 0xe3, 0x4f,
	0xfe, 0xc9, 0xf7, 0x4d, 0x88, 0x4c, 0x0a, 0xe7, 0x14, 0x17, 0x65, 0x8d, 0x2a, 0x10, 0xe0, 0x17,
	0xe7, 0xe8, 0xea, 0x6e, 0xe9, 0xb1, 0x60, 0xac, 0x32, 0xe7, 0x89, 0x29, 0x3a, 0x1e, 0xce, 0x2f,
	0xe3, 0x21, 0x71, 0xe7, 0x50, 0x35, 0x6d, 0xcd, 0x06, 0xcc, 0x58, 0x48, 0xbc, 0x8c, 0xb6, 0x87,
	0xa3, 0x0b, 0x79, 0xe2, 0x28, 0x30, 0x98, 0x2a, 0x16, 0x0b, 0x91, 0x1e, 0xe5, 0xf3, 0x3e, 0xbe,
	0x55, 0xd1, 0x52, 0x9e, 0xeb, 0xdf, 0x0e, 0x90, 0xaf, 0xd0, 0xc7, 0xb3, 0x85, 0x48, 0x4f, 0x1f,
	0xaf, 0xb2, 0x0d, 0x20, 0xf3, 0xfd, 0x45, 0x30, 0x69, 0xe5, 0x67, 0x63, 0x7c, 0xef, 0x53, 0xf2,
	0x2c, 0xaf, 0x03, 0x41, 0x83, 0xcd, 0x7f, 0x2e, 0x94, 0x1d, 0x8c, 0xdc, 0x9d, 0x0a, 0x89, 0x95,
	0xb7, 0xbb, 0xa7, 0x27, 0xd8, 0x6d, 0x5b, 0xb1, 0x87, 0xed, 0xb7, 0x90, 0x69, 0x20, 0xa8, 0x4d,
	0xe0, 0xc3, 0x10, 0xff, 0xae, 0x3e, 0x9b, 0xac, 0xfe, 0x22, 0x6f, 0x02, 0x41, 0x6b, 0xcd, 0x67,
	0x8b, 0xfb, 0xbc, 0xdd, 0xeb, 0xfd, 0xc7, 0xdd, 0xca, 0xf0, 0xba, 0x2d, 0x33, 0x4d, 0x8d, 0xd2,
	0x14, 0xf4, 0xba, 0x30, 0x56, 0x69, 0xb2, 0x8c, 0x96, 0x29, 0x3c, 0xac, 0x48, 0x6d, 0x7e, 0x55,
	0x39, 0x1f, 0xa4, 0x9f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xa5, 0xd2, 0xb3, 0x69, 0x02,
	0x00, 0x00,
}
