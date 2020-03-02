// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_vestDelegation.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SoVestDelegation struct {
	Id                   uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FromAccount          *prototype.AccountName `protobuf:"bytes,2,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	ToAccount            *prototype.AccountName `protobuf:"bytes,3,opt,name=to_account,json=toAccount,proto3" json:"to_account,omitempty"`
	Amount               *prototype.Vest        `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedBlock         uint64                 `protobuf:"varint,5,opt,name=created_block,json=createdBlock,proto3" json:"created_block,omitempty"`
	MaturityBlock        uint64                 `protobuf:"varint,6,opt,name=maturity_block,json=maturityBlock,proto3" json:"maturity_block,omitempty"`
	DeliveryBlock        uint64                 `protobuf:"varint,7,opt,name=delivery_block,json=deliveryBlock,proto3" json:"delivery_block,omitempty"`
	Delivering           bool                   `protobuf:"varint,8,opt,name=delivering,proto3" json:"delivering,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoVestDelegation) Reset()         { *m = SoVestDelegation{} }
func (m *SoVestDelegation) String() string { return proto.CompactTextString(m) }
func (*SoVestDelegation) ProtoMessage()    {}
func (*SoVestDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{0}
}

func (m *SoVestDelegation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoVestDelegation.Unmarshal(m, b)
}
func (m *SoVestDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoVestDelegation.Marshal(b, m, deterministic)
}
func (m *SoVestDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoVestDelegation.Merge(m, src)
}
func (m *SoVestDelegation) XXX_Size() int {
	return xxx_messageInfo_SoVestDelegation.Size(m)
}
func (m *SoVestDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_SoVestDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_SoVestDelegation proto.InternalMessageInfo

func (m *SoVestDelegation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SoVestDelegation) GetFromAccount() *prototype.AccountName {
	if m != nil {
		return m.FromAccount
	}
	return nil
}

func (m *SoVestDelegation) GetToAccount() *prototype.AccountName {
	if m != nil {
		return m.ToAccount
	}
	return nil
}

func (m *SoVestDelegation) GetAmount() *prototype.Vest {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *SoVestDelegation) GetCreatedBlock() uint64 {
	if m != nil {
		return m.CreatedBlock
	}
	return 0
}

func (m *SoVestDelegation) GetMaturityBlock() uint64 {
	if m != nil {
		return m.MaturityBlock
	}
	return 0
}

func (m *SoVestDelegation) GetDeliveryBlock() uint64 {
	if m != nil {
		return m.DeliveryBlock
	}
	return 0
}

func (m *SoVestDelegation) GetDelivering() bool {
	if m != nil {
		return m.Delivering
	}
	return false
}

type SoListVestDelegationByFromAccount struct {
	FromAccount          *prototype.AccountName `protobuf:"bytes,1,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	Id                   uint64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListVestDelegationByFromAccount) Reset()         { *m = SoListVestDelegationByFromAccount{} }
func (m *SoListVestDelegationByFromAccount) String() string { return proto.CompactTextString(m) }
func (*SoListVestDelegationByFromAccount) ProtoMessage()    {}
func (*SoListVestDelegationByFromAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{1}
}

func (m *SoListVestDelegationByFromAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListVestDelegationByFromAccount.Unmarshal(m, b)
}
func (m *SoListVestDelegationByFromAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListVestDelegationByFromAccount.Marshal(b, m, deterministic)
}
func (m *SoListVestDelegationByFromAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListVestDelegationByFromAccount.Merge(m, src)
}
func (m *SoListVestDelegationByFromAccount) XXX_Size() int {
	return xxx_messageInfo_SoListVestDelegationByFromAccount.Size(m)
}
func (m *SoListVestDelegationByFromAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListVestDelegationByFromAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListVestDelegationByFromAccount proto.InternalMessageInfo

func (m *SoListVestDelegationByFromAccount) GetFromAccount() *prototype.AccountName {
	if m != nil {
		return m.FromAccount
	}
	return nil
}

func (m *SoListVestDelegationByFromAccount) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SoListVestDelegationByToAccount struct {
	ToAccount            *prototype.AccountName `protobuf:"bytes,1,opt,name=to_account,json=toAccount,proto3" json:"to_account,omitempty"`
	Id                   uint64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListVestDelegationByToAccount) Reset()         { *m = SoListVestDelegationByToAccount{} }
func (m *SoListVestDelegationByToAccount) String() string { return proto.CompactTextString(m) }
func (*SoListVestDelegationByToAccount) ProtoMessage()    {}
func (*SoListVestDelegationByToAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{2}
}

func (m *SoListVestDelegationByToAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListVestDelegationByToAccount.Unmarshal(m, b)
}
func (m *SoListVestDelegationByToAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListVestDelegationByToAccount.Marshal(b, m, deterministic)
}
func (m *SoListVestDelegationByToAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListVestDelegationByToAccount.Merge(m, src)
}
func (m *SoListVestDelegationByToAccount) XXX_Size() int {
	return xxx_messageInfo_SoListVestDelegationByToAccount.Size(m)
}
func (m *SoListVestDelegationByToAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListVestDelegationByToAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListVestDelegationByToAccount proto.InternalMessageInfo

func (m *SoListVestDelegationByToAccount) GetToAccount() *prototype.AccountName {
	if m != nil {
		return m.ToAccount
	}
	return nil
}

func (m *SoListVestDelegationByToAccount) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SoListVestDelegationByMaturityBlock struct {
	MaturityBlock        uint64   `protobuf:"varint,1,opt,name=maturity_block,json=maturityBlock,proto3" json:"maturity_block,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListVestDelegationByMaturityBlock) Reset()         { *m = SoListVestDelegationByMaturityBlock{} }
func (m *SoListVestDelegationByMaturityBlock) String() string { return proto.CompactTextString(m) }
func (*SoListVestDelegationByMaturityBlock) ProtoMessage()    {}
func (*SoListVestDelegationByMaturityBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{3}
}

func (m *SoListVestDelegationByMaturityBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListVestDelegationByMaturityBlock.Unmarshal(m, b)
}
func (m *SoListVestDelegationByMaturityBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListVestDelegationByMaturityBlock.Marshal(b, m, deterministic)
}
func (m *SoListVestDelegationByMaturityBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListVestDelegationByMaturityBlock.Merge(m, src)
}
func (m *SoListVestDelegationByMaturityBlock) XXX_Size() int {
	return xxx_messageInfo_SoListVestDelegationByMaturityBlock.Size(m)
}
func (m *SoListVestDelegationByMaturityBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListVestDelegationByMaturityBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SoListVestDelegationByMaturityBlock proto.InternalMessageInfo

func (m *SoListVestDelegationByMaturityBlock) GetMaturityBlock() uint64 {
	if m != nil {
		return m.MaturityBlock
	}
	return 0
}

func (m *SoListVestDelegationByMaturityBlock) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SoListVestDelegationByDeliveryBlock struct {
	DeliveryBlock        uint64   `protobuf:"varint,1,opt,name=delivery_block,json=deliveryBlock,proto3" json:"delivery_block,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListVestDelegationByDeliveryBlock) Reset()         { *m = SoListVestDelegationByDeliveryBlock{} }
func (m *SoListVestDelegationByDeliveryBlock) String() string { return proto.CompactTextString(m) }
func (*SoListVestDelegationByDeliveryBlock) ProtoMessage()    {}
func (*SoListVestDelegationByDeliveryBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{4}
}

func (m *SoListVestDelegationByDeliveryBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListVestDelegationByDeliveryBlock.Unmarshal(m, b)
}
func (m *SoListVestDelegationByDeliveryBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListVestDelegationByDeliveryBlock.Marshal(b, m, deterministic)
}
func (m *SoListVestDelegationByDeliveryBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListVestDelegationByDeliveryBlock.Merge(m, src)
}
func (m *SoListVestDelegationByDeliveryBlock) XXX_Size() int {
	return xxx_messageInfo_SoListVestDelegationByDeliveryBlock.Size(m)
}
func (m *SoListVestDelegationByDeliveryBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListVestDelegationByDeliveryBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SoListVestDelegationByDeliveryBlock proto.InternalMessageInfo

func (m *SoListVestDelegationByDeliveryBlock) GetDeliveryBlock() uint64 {
	if m != nil {
		return m.DeliveryBlock
	}
	return 0
}

func (m *SoListVestDelegationByDeliveryBlock) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SoListVestDelegationByDelivering struct {
	Delivering           bool     `protobuf:"varint,1,opt,name=delivering,proto3" json:"delivering,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoListVestDelegationByDelivering) Reset()         { *m = SoListVestDelegationByDelivering{} }
func (m *SoListVestDelegationByDelivering) String() string { return proto.CompactTextString(m) }
func (*SoListVestDelegationByDelivering) ProtoMessage()    {}
func (*SoListVestDelegationByDelivering) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{5}
}

func (m *SoListVestDelegationByDelivering) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListVestDelegationByDelivering.Unmarshal(m, b)
}
func (m *SoListVestDelegationByDelivering) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListVestDelegationByDelivering.Marshal(b, m, deterministic)
}
func (m *SoListVestDelegationByDelivering) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListVestDelegationByDelivering.Merge(m, src)
}
func (m *SoListVestDelegationByDelivering) XXX_Size() int {
	return xxx_messageInfo_SoListVestDelegationByDelivering.Size(m)
}
func (m *SoListVestDelegationByDelivering) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListVestDelegationByDelivering.DiscardUnknown(m)
}

var xxx_messageInfo_SoListVestDelegationByDelivering proto.InternalMessageInfo

func (m *SoListVestDelegationByDelivering) GetDelivering() bool {
	if m != nil {
		return m.Delivering
	}
	return false
}

func (m *SoListVestDelegationByDelivering) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SoUniqueVestDelegationById struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueVestDelegationById) Reset()         { *m = SoUniqueVestDelegationById{} }
func (m *SoUniqueVestDelegationById) String() string { return proto.CompactTextString(m) }
func (*SoUniqueVestDelegationById) ProtoMessage()    {}
func (*SoUniqueVestDelegationById) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb4eca8094fbeb0a, []int{6}
}

func (m *SoUniqueVestDelegationById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueVestDelegationById.Unmarshal(m, b)
}
func (m *SoUniqueVestDelegationById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueVestDelegationById.Marshal(b, m, deterministic)
}
func (m *SoUniqueVestDelegationById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueVestDelegationById.Merge(m, src)
}
func (m *SoUniqueVestDelegationById) XXX_Size() int {
	return xxx_messageInfo_SoUniqueVestDelegationById.Size(m)
}
func (m *SoUniqueVestDelegationById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueVestDelegationById.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueVestDelegationById proto.InternalMessageInfo

func (m *SoUniqueVestDelegationById) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*SoVestDelegation)(nil), "table.so_vestDelegation")
	proto.RegisterType((*SoListVestDelegationByFromAccount)(nil), "table.so_list_vestDelegation_by_from_account")
	proto.RegisterType((*SoListVestDelegationByToAccount)(nil), "table.so_list_vestDelegation_by_to_account")
	proto.RegisterType((*SoListVestDelegationByMaturityBlock)(nil), "table.so_list_vestDelegation_by_maturity_block")
	proto.RegisterType((*SoListVestDelegationByDeliveryBlock)(nil), "table.so_list_vestDelegation_by_delivery_block")
	proto.RegisterType((*SoListVestDelegationByDelivering)(nil), "table.so_list_vestDelegation_by_delivering")
	proto.RegisterType((*SoUniqueVestDelegationById)(nil), "table.so_unique_vestDelegation_by_id")
}

func init() { proto.RegisterFile("app/table/so_vestDelegation.proto", fileDescriptor_fb4eca8094fbeb0a) }

var fileDescriptor_fb4eca8094fbeb0a = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0x91, 0x9b, 0xa4, 0xa9, 0xf2, 0x53, 0x6a, 0x0a, 0x35, 0x5d, 0x84, 0x34, 0xfd, 0x33,
	0xa5, 0xb5, 0x4b, 0x0b, 0x5d, 0x74, 0xd7, 0xd0, 0x27, 0xf0, 0x62, 0x16, 0xb3, 0x11, 0xb2, 0xac,
	0x71, 0xc4, 0xd8, 0xba, 0x1e, 0xfb, 0x3a, 0x90, 0xe7, 0x9c, 0x17, 0x1a, 0xfc, 0x9b, 0xc4, 0x71,
	0xc2, 0x64, 0x63, 0xf0, 0xb9, 0xe7, 0xe8, 0x8a, 0xf3, 0x21, 0xfa, 0x81, 0x27, 0x89, 0x8b, 0xdc,
	0x8f, 0xa4, 0x9b, 0x01, 0xdb, 0xca, 0x0c, 0xff, 0xcb, 0x48, 0x86, 0x1c, 0x15, 0x68, 0x27, 0x49,
	0x01, 0xc1, 0x1c, 0x96, 0xe3, 0xf7, 0x6f, 0xcb, 0x3f, 0xdc, 0x25, 0xd2, 0x2d, 0x3e, 0xd5, 0x70,
	0xf5, 0x68, 0xd0, 0x37, 0x27, 0x41, 0x73, 0x4e, 0x0d, 0x15, 0x58, 0x64, 0x49, 0xec, 0x81, 0x67,
	0xa8, 0xc0, 0xfc, 0x4b, 0xa7, 0x77, 0x29, 0xc4, 0x8c, 0x0b, 0x01, 0xb9, 0x46, 0xcb, 0x58, 0x12,
	0x7b, 0xf2, 0xeb, 0x9d, 0xd3, 0x1e, 0xe9, 0xd4, 0x13, 0xa6, 0x79, 0x2c, 0xbd, 0x49, 0x61, 0xfe,
	0x57, 0x29, 0xe6, 0x1f, 0x4a, 0x11, 0xda, 0xe4, 0x8b, 0xcb, 0xc9, 0x57, 0x08, 0x4d, 0xee, 0x2b,
	0x1d, 0xf1, 0xb8, 0xcc, 0x0c, 0xca, 0xcc, 0xeb, 0x83, 0x4c, 0x71, 0x5d, 0xaf, 0x1e, 0x9b, 0x1f,
	0xe9, 0x4c, 0xa4, 0x92, 0xa3, 0x0c, 0x98, 0x1f, 0x81, 0xb8, 0xb7, 0x86, 0xe5, 0xbd, 0xa7, 0xb5,
	0xb8, 0x2e, 0x34, 0xf3, 0x33, 0x9d, 0xc7, 0x1c, 0xf3, 0x54, 0xe1, 0xae, 0x76, 0x8d, 0x4a, 0xd7,
	0xac, 0x51, 0x5b, 0x5b, 0x20, 0x23, 0xb5, 0x95, 0x69, 0x63, 0x7b, 0x59, 0xd9, 0x1a, 0xb5, 0xb2,
	0x2d, 0x28, 0xad, 0x05, 0xa5, 0x43, 0x6b, 0xbc, 0x24, 0xf6, 0xd8, 0x3b, 0x50, 0x56, 0x48, 0xbf,
	0x64, 0xc0, 0x22, 0x95, 0x61, 0xa7, 0x59, 0xe6, 0xef, 0xd8, 0x61, 0x93, 0x27, 0xcd, 0x92, 0x2b,
	0x9a, 0xad, 0x28, 0x19, 0x0d, 0xa5, 0x95, 0xa6, 0x9f, 0xce, 0x6f, 0xdd, 0x33, 0xe8, 0x10, 0x21,
	0xcf, 0x26, 0xd2, 0xdd, 0xc7, 0xa9, 0x7d, 0x7e, 0xdf, 0x71, 0xdb, 0x3d, 0xfd, 0x93, 0xbe, 0xfe,
	0xaf, 0x5a, 0x71, 0x4c, 0xaa, 0x87, 0x1d, 0xe9, 0x63, 0xd7, 0x5d, 0x71, 0x73, 0xa9, 0xb5, 0x3d,
	0xd3, 0x0e, 0x73, 0xd2, 0x65, 0x7e, 0x72, 0xee, 0x4f, 0xba, 0xc8, 0x80, 0xe5, 0x5a, 0x3d, 0xe4,
	0xb2, 0xe7, 0x64, 0x15, 0x74, 0x5f, 0xd9, 0xfa, 0xfb, 0xed, 0xb7, 0x50, 0xe1, 0x26, 0xf7, 0x1d,
	0x01, 0xb1, 0x2b, 0x20, 0x13, 0x1b, 0xae, 0xb4, 0x2b, 0x40, 0xa3, 0xd4, 0x08, 0xd9, 0x8f, 0x10,
	0xdc, 0xf6, 0xb9, 0xfb, 0xa3, 0x12, 0xd8, 0xef, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x0b,
	0x52, 0x27, 0x02, 0x04, 0x00, 0x00,
}
