// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_witness.proto

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

type SoWitness struct {
	Owner                 *prototype.AccountName   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	CreatedTime           *prototype.TimePointSec  `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Url                   string                   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	VoteCount             uint64                   `protobuf:"varint,4,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	LastConfirmedBlockNum uint32                   `protobuf:"varint,5,opt,name=last_confirmed_block_num,json=lastConfirmedBlockNum,proto3" json:"last_confirmed_block_num,omitempty"`
	TotalMissed           uint32                   `protobuf:"varint,6,opt,name=total_missed,json=totalMissed,proto3" json:"total_missed,omitempty"`
	PowWorker             uint32                   `protobuf:"varint,7,opt,name=pow_worker,json=powWorker,proto3" json:"pow_worker,omitempty"`
	SigningKey            *prototype.PublicKeyType `protobuf:"bytes,8,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	LastWork              *prototype.Sha256        `protobuf:"bytes,9,opt,name=last_work,json=lastWork,proto3" json:"last_work,omitempty"`
	RunningVersion        uint32                   `protobuf:"varint,10,opt,name=running_version,json=runningVersion,proto3" json:"running_version,omitempty"`
	LastAslot             uint32                   `protobuf:"varint,11,opt,name=last_aslot,json=lastAslot,proto3" json:"last_aslot,omitempty"`
	ProposedStaminaFree   uint64                   `protobuf:"varint,12,opt,name=proposed_stamina_free,json=proposedStaminaFree,proto3" json:"proposed_stamina_free,omitempty"`
	Active                bool                     `protobuf:"varint,13,opt,name=active,proto3" json:"active,omitempty"`
	TpsExpected           uint64                   `protobuf:"varint,14,opt,name=tps_expected,json=tpsExpected,proto3" json:"tps_expected,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *SoWitness) Reset()         { *m = SoWitness{} }
func (m *SoWitness) String() string { return proto.CompactTextString(m) }
func (*SoWitness) ProtoMessage()    {}
func (*SoWitness) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{0}
}

func (m *SoWitness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoWitness.Unmarshal(m, b)
}
func (m *SoWitness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoWitness.Marshal(b, m, deterministic)
}
func (m *SoWitness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoWitness.Merge(m, src)
}
func (m *SoWitness) XXX_Size() int {
	return xxx_messageInfo_SoWitness.Size(m)
}
func (m *SoWitness) XXX_DiscardUnknown() {
	xxx_messageInfo_SoWitness.DiscardUnknown(m)
}

var xxx_messageInfo_SoWitness proto.InternalMessageInfo

func (m *SoWitness) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *SoWitness) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoWitness) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SoWitness) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

func (m *SoWitness) GetLastConfirmedBlockNum() uint32 {
	if m != nil {
		return m.LastConfirmedBlockNum
	}
	return 0
}

func (m *SoWitness) GetTotalMissed() uint32 {
	if m != nil {
		return m.TotalMissed
	}
	return 0
}

func (m *SoWitness) GetPowWorker() uint32 {
	if m != nil {
		return m.PowWorker
	}
	return 0
}

func (m *SoWitness) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

func (m *SoWitness) GetLastWork() *prototype.Sha256 {
	if m != nil {
		return m.LastWork
	}
	return nil
}

func (m *SoWitness) GetRunningVersion() uint32 {
	if m != nil {
		return m.RunningVersion
	}
	return 0
}

func (m *SoWitness) GetLastAslot() uint32 {
	if m != nil {
		return m.LastAslot
	}
	return 0
}

func (m *SoWitness) GetProposedStaminaFree() uint64 {
	if m != nil {
		return m.ProposedStaminaFree
	}
	return 0
}

func (m *SoWitness) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *SoWitness) GetTpsExpected() uint64 {
	if m != nil {
		return m.TpsExpected
	}
	return 0
}

type SoListWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByOwner) Reset()         { *m = SoListWitnessByOwner{} }
func (m *SoListWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByOwner) ProtoMessage()    {}
func (*SoListWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{1}
}

func (m *SoListWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByOwner.Unmarshal(m, b)
}
func (m *SoListWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoListWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByOwner.Merge(m, src)
}
func (m *SoListWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByOwner.Size(m)
}
func (m *SoListWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByOwner proto.InternalMessageInfo

func (m *SoListWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoListWitnessByVoteCount struct {
	VoteCount            uint64                 `protobuf:"varint,1,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	Owner                *prototype.AccountName `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByVoteCount) Reset()         { *m = SoListWitnessByVoteCount{} }
func (m *SoListWitnessByVoteCount) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByVoteCount) ProtoMessage()    {}
func (*SoListWitnessByVoteCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{2}
}

func (m *SoListWitnessByVoteCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByVoteCount.Unmarshal(m, b)
}
func (m *SoListWitnessByVoteCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByVoteCount.Marshal(b, m, deterministic)
}
func (m *SoListWitnessByVoteCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByVoteCount.Merge(m, src)
}
func (m *SoListWitnessByVoteCount) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByVoteCount.Size(m)
}
func (m *SoListWitnessByVoteCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByVoteCount.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByVoteCount proto.InternalMessageInfo

func (m *SoListWitnessByVoteCount) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

func (m *SoListWitnessByVoteCount) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoUniqueWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoUniqueWitnessByOwner) Reset()         { *m = SoUniqueWitnessByOwner{} }
func (m *SoUniqueWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoUniqueWitnessByOwner) ProtoMessage()    {}
func (*SoUniqueWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{3}
}

func (m *SoUniqueWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Unmarshal(m, b)
}
func (m *SoUniqueWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoUniqueWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueWitnessByOwner.Merge(m, src)
}
func (m *SoUniqueWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Size(m)
}
func (m *SoUniqueWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueWitnessByOwner proto.InternalMessageInfo

func (m *SoUniqueWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*SoWitness)(nil), "table.so_witness")
	proto.RegisterType((*SoListWitnessByOwner)(nil), "table.so_list_witness_by_owner")
	proto.RegisterType((*SoListWitnessByVoteCount)(nil), "table.so_list_witness_by_vote_count")
	proto.RegisterType((*SoUniqueWitnessByOwner)(nil), "table.so_unique_witness_by_owner")
}

func init() { proto.RegisterFile("app/table/so_witness.proto", fileDescriptor_00097e516fc05425) }

var fileDescriptor_00097e516fc05425 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6f, 0x13, 0x31,
	0x10, 0x85, 0xb5, 0x6d, 0x13, 0x12, 0x6f, 0x5a, 0xc0, 0x50, 0x30, 0x91, 0x2a, 0x85, 0x5c, 0x88,
	0x10, 0xcd, 0x4a, 0x41, 0xc0, 0x01, 0x2e, 0xb4, 0x02, 0x09, 0x55, 0x70, 0x58, 0x10, 0x48, 0x5c,
	0x2c, 0xaf, 0x33, 0x4d, 0xac, 0xec, 0x7a, 0x8c, 0xed, 0x4d, 0xc8, 0x4f, 0xe4, 0x5f, 0x21, 0x7b,
	0xd3, 0x34, 0xaa, 0x38, 0x80, 0xb8, 0x44, 0xf1, 0xf7, 0x66, 0xe6, 0x8d, 0xad, 0xb7, 0xa4, 0x2f,
	0x8c, 0xc9, 0xbc, 0x28, 0x4a, 0xc8, 0x1c, 0xf2, 0x95, 0xf2, 0x1a, 0x9c, 0x1b, 0x1b, 0x8b, 0x1e,
	0x69, 0x2b, 0xf2, 0xfe, 0xfd, 0x78, 0xf2, 0x6b, 0x03, 0x59, 0xf8, 0x69, 0xc4, 0xe1, 0xaf, 0x03,
	0x42, 0xae, 0x3b, 0xe8, 0x29, 0x69, 0xe1, 0x4a, 0x83, 0x65, 0xc9, 0x20, 0x19, 0xa5, 0x93, 0x87,
	0xe3, 0x6d, 0xd3, 0x58, 0x48, 0x89, 0xb5, 0xf6, 0x5c, 0x8b, 0x0a, 0xf2, 0xa6, 0x8a, 0xbe, 0x21,
	0x3d, 0x69, 0x41, 0x78, 0x98, 0x72, 0xaf, 0x2a, 0x60, 0x7b, 0xb1, 0xeb, 0xd1, 0x4e, 0x57, 0xc0,
	0xdc, 0xa0, 0xd2, 0x9e, 0x3b, 0x90, 0x79, 0xba, 0x29, 0xff, 0xa2, 0x2a, 0xa0, 0x77, 0xc8, 0x7e,
	0x6d, 0x4b, 0xb6, 0x3f, 0x48, 0x46, 0xdd, 0x3c, 0xfc, 0xa5, 0x27, 0x84, 0x2c, 0xd1, 0x03, 0x8f,
	0x4e, 0xec, 0x60, 0x90, 0x8c, 0x0e, 0xf2, 0x6e, 0x20, 0xe7, 0x01, 0xd0, 0x57, 0x84, 0x95, 0xc2,
	0x79, 0x2e, 0x51, 0x5f, 0x2a, 0x5b, 0xc1, 0x94, 0x17, 0x25, 0xca, 0x05, 0xd7, 0x75, 0xc5, 0x5a,
	0x83, 0x64, 0x74, 0x98, 0x1f, 0x07, 0xfd, 0xfc, 0x4a, 0x3e, 0x0b, 0xea, 0xa7, 0xba, 0xa2, 0x8f,
	0x49, 0xcf, 0xa3, 0x17, 0x25, 0xaf, 0x94, 0x73, 0x30, 0x65, 0xed, 0x58, 0x9c, 0x46, 0xf6, 0x31,
	0xa2, 0x60, 0x6d, 0x70, 0xc5, 0x57, 0x68, 0x17, 0x60, 0xd9, 0xad, 0x58, 0xd0, 0x35, 0xb8, 0xfa,
	0x16, 0x01, 0x7d, 0x4d, 0x52, 0xa7, 0x66, 0x5a, 0xe9, 0x19, 0x5f, 0xc0, 0x9a, 0x75, 0xe2, 0x45,
	0xfb, 0x3b, 0x17, 0x35, 0x75, 0x51, 0x2a, 0x19, 0x44, 0x1e, 0xce, 0x39, 0xd9, 0x94, 0x5f, 0xc0,
	0x9a, 0x8e, 0x49, 0x37, 0xee, 0x1d, 0x86, 0xb3, 0x6e, 0x6c, 0xbd, 0xbb, 0xd3, 0xea, 0xe6, 0x62,
	0xf2, 0xe2, 0x65, 0xde, 0x09, 0x35, 0xc1, 0x8e, 0x3e, 0x21, 0xb7, 0x6d, 0xad, 0xa3, 0xd9, 0x12,
	0xac, 0x53, 0xa8, 0x19, 0x89, 0x0b, 0x1d, 0x6d, 0xf0, 0xd7, 0x86, 0x86, 0xa5, 0xe3, 0x60, 0xe1,
	0x4a, 0xf4, 0x2c, 0x6d, 0x96, 0x0e, 0xe4, 0x6d, 0x00, 0x74, 0x42, 0x8e, 0x8d, 0x45, 0x83, 0x0e,
	0xa6, 0xdc, 0x79, 0x51, 0x29, 0x2d, 0xf8, 0xa5, 0x05, 0x60, 0xbd, 0xf8, 0xb2, 0xf7, 0xae, 0xc4,
	0xcf, 0x8d, 0xf6, 0xde, 0x02, 0xd0, 0x07, 0xa4, 0x2d, 0xa4, 0x57, 0x4b, 0x60, 0x87, 0x83, 0x64,
	0xd4, 0xc9, 0x37, 0xa7, 0xf8, 0x84, 0xc6, 0x71, 0xf8, 0x69, 0x40, 0x7a, 0x98, 0xb2, 0xa3, 0x38,
	0x22, 0xf5, 0xc6, 0xbd, 0xdb, 0xa0, 0xe1, 0x07, 0xc2, 0x1c, 0xf2, 0x52, 0x85, 0x9b, 0x36, 0x79,
	0xe2, 0xc5, 0x9a, 0x37, 0x49, 0xf9, 0xb7, 0x60, 0x0d, 0x2b, 0x72, 0xf2, 0x87, 0x51, 0xd7, 0xd9,
	0xb8, 0x91, 0x94, 0xe4, 0x66, 0x52, 0xb6, 0x76, 0x7b, 0x7f, 0x65, 0x77, 0x41, 0xfa, 0x0e, 0x79,
	0xad, 0xd5, 0x8f, 0x1a, 0xfe, 0x77, 0xf7, 0xb3, 0x67, 0xdf, 0x9f, 0xce, 0x94, 0x9f, 0xd7, 0xc5,
	0x58, 0x62, 0x95, 0x49, 0x74, 0x72, 0x2e, 0x94, 0xce, 0x24, 0x6a, 0x0f, 0xda, 0xa3, 0x3b, 0x9d,
	0x61, 0xb6, 0xfd, 0x5c, 0x8b, 0x76, 0x1c, 0xf6, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13,
	0x29, 0xa8, 0x6e, 0xc2, 0x03, 0x00, 0x00,
}
