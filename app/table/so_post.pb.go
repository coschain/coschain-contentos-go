// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_post.proto

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

type SoPost struct {
	PostId               uint64                            `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Category             string                            `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Author               *prototype.AccountName            `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Title                string                            `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string                            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Tags                 []string                          `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Created              *prototype.TimePointSec           `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	LastPayout           *prototype.TimePointSec           `protobuf:"bytes,8,opt,name=last_payout,json=lastPayout,proto3" json:"last_payout,omitempty"`
	Depth                uint32                            `protobuf:"varint,9,opt,name=depth,proto3" json:"depth,omitempty"`
	Children             uint32                            `protobuf:"varint,10,opt,name=children,proto3" json:"children,omitempty"`
	RootId               uint64                            `protobuf:"varint,11,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	ParentId             uint64                            `protobuf:"varint,12,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	VoteCnt              uint64                            `protobuf:"varint,13,opt,name=vote_cnt,json=voteCnt,proto3" json:"vote_cnt,omitempty"`
	Beneficiaries        []*prototype.BeneficiaryRouteType `protobuf:"bytes,14,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	CashoutTime          *prototype.TimePointSec           `protobuf:"bytes,15,opt,name=cashout_time,json=cashoutTime,proto3" json:"cashout_time,omitempty"`
	WeightedVp           uint64                            `protobuf:"varint,16,opt,name=weighted_vp,json=weightedVp,proto3" json:"weighted_vp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SoPost) Reset()         { *m = SoPost{} }
func (m *SoPost) String() string { return proto.CompactTextString(m) }
func (*SoPost) ProtoMessage()    {}
func (*SoPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{0}
}

func (m *SoPost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoPost.Unmarshal(m, b)
}
func (m *SoPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoPost.Marshal(b, m, deterministic)
}
func (m *SoPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoPost.Merge(m, src)
}
func (m *SoPost) XXX_Size() int {
	return xxx_messageInfo_SoPost.Size(m)
}
func (m *SoPost) XXX_DiscardUnknown() {
	xxx_messageInfo_SoPost.DiscardUnknown(m)
}

var xxx_messageInfo_SoPost proto.InternalMessageInfo

func (m *SoPost) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *SoPost) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *SoPost) GetAuthor() *prototype.AccountName {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *SoPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SoPost) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SoPost) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SoPost) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SoPost) GetLastPayout() *prototype.TimePointSec {
	if m != nil {
		return m.LastPayout
	}
	return nil
}

func (m *SoPost) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *SoPost) GetChildren() uint32 {
	if m != nil {
		return m.Children
	}
	return 0
}

func (m *SoPost) GetRootId() uint64 {
	if m != nil {
		return m.RootId
	}
	return 0
}

func (m *SoPost) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *SoPost) GetVoteCnt() uint64 {
	if m != nil {
		return m.VoteCnt
	}
	return 0
}

func (m *SoPost) GetBeneficiaries() []*prototype.BeneficiaryRouteType {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

func (m *SoPost) GetCashoutTime() *prototype.TimePointSec {
	if m != nil {
		return m.CashoutTime
	}
	return nil
}

func (m *SoPost) GetWeightedVp() uint64 {
	if m != nil {
		return m.WeightedVp
	}
	return 0
}

type SoListPostByCreated struct {
	Created              *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	PostId               uint64                  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListPostByCreated) Reset()         { *m = SoListPostByCreated{} }
func (m *SoListPostByCreated) String() string { return proto.CompactTextString(m) }
func (*SoListPostByCreated) ProtoMessage()    {}
func (*SoListPostByCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{1}
}

func (m *SoListPostByCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListPostByCreated.Unmarshal(m, b)
}
func (m *SoListPostByCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListPostByCreated.Marshal(b, m, deterministic)
}
func (m *SoListPostByCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListPostByCreated.Merge(m, src)
}
func (m *SoListPostByCreated) XXX_Size() int {
	return xxx_messageInfo_SoListPostByCreated.Size(m)
}
func (m *SoListPostByCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListPostByCreated.DiscardUnknown(m)
}

var xxx_messageInfo_SoListPostByCreated proto.InternalMessageInfo

func (m *SoListPostByCreated) GetCreated() *prototype.TimePointSec {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SoListPostByCreated) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type SoUniquePostByPostId struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniquePostByPostId) Reset()         { *m = SoUniquePostByPostId{} }
func (m *SoUniquePostByPostId) String() string { return proto.CompactTextString(m) }
func (*SoUniquePostByPostId) ProtoMessage()    {}
func (*SoUniquePostByPostId) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaea493edfbbad11, []int{2}
}

func (m *SoUniquePostByPostId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniquePostByPostId.Unmarshal(m, b)
}
func (m *SoUniquePostByPostId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniquePostByPostId.Marshal(b, m, deterministic)
}
func (m *SoUniquePostByPostId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniquePostByPostId.Merge(m, src)
}
func (m *SoUniquePostByPostId) XXX_Size() int {
	return xxx_messageInfo_SoUniquePostByPostId.Size(m)
}
func (m *SoUniquePostByPostId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniquePostByPostId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniquePostByPostId proto.InternalMessageInfo

func (m *SoUniquePostByPostId) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SoPost)(nil), "table.so_post")
	proto.RegisterType((*SoListPostByCreated)(nil), "table.so_list_post_by_created")
	proto.RegisterType((*SoUniquePostByPostId)(nil), "table.so_unique_post_by_post_id")
}

func init() { proto.RegisterFile("app/table/so_post.proto", fileDescriptor_aaea493edfbbad11) }

var fileDescriptor_aaea493edfbbad11 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xb6, 0xff, 0xce, 0x16, 0x90, 0xb5, 0x52, 0xdd, 0xe5, 0x40, 0xe8, 0x01, 0xe5, 0x42,
	0x23, 0xed, 0x72, 0x42, 0x9c, 0xe0, 0x80, 0x7a, 0x43, 0x11, 0xe2, 0xc0, 0xc5, 0x72, 0x9c, 0x21,
	0xb1, 0xd4, 0x7a, 0x42, 0x3c, 0x59, 0x94, 0x47, 0xe4, 0xad, 0x90, 0x9d, 0xb6, 0x94, 0x03, 0xda,
	0xbd, 0x24, 0xfe, 0x7e, 0x9c, 0xb1, 0xbf, 0x99, 0xb0, 0x95, 0x6a, 0x9a, 0x8c, 0x54, 0xb1, 0x87,
	0xcc, 0xa1, 0x6c, 0xd0, 0xd1, 0xb6, 0x69, 0x91, 0x90, 0x4f, 0x02, 0x79, 0x7b, 0x13, 0x10, 0xf5,
	0x0d, 0x64, 0xfe, 0x31, 0x88, 0x9b, 0xdf, 0x63, 0x36, 0x3b, 0xda, 0xf9, 0x8a, 0xcd, 0xfc, 0x5b,
	0x9a, 0x52, 0x44, 0x49, 0x94, 0x8e, 0xf3, 0xa9, 0x87, 0xbb, 0x92, 0xdf, 0xb2, 0xb9, 0x56, 0x04,
	0x15, 0xb6, 0xbd, 0xb8, 0x4a, 0xa2, 0x74, 0x91, 0x9f, 0x31, 0xcf, 0xd8, 0x54, 0x75, 0x54, 0x63,
	0x2b, 0x46, 0x49, 0x94, 0xc6, 0x77, 0xab, 0xed, 0xb9, 0xce, 0x56, 0x69, 0x8d, 0x9d, 0x25, 0x69,
	0xd5, 0x01, 0xf2, 0xa3, 0x8d, 0xdf, 0xb0, 0x09, 0x19, 0xda, 0x83, 0x18, 0x87, 0x2f, 0x0d, 0x80,
	0x73, 0x36, 0x2e, 0xb0, 0xec, 0xc5, 0x24, 0x90, 0x61, 0xed, 0x39, 0x52, 0x95, 0x13, 0xd3, 0x64,
	0xe4, 0x39, 0xbf, 0xe6, 0xf7, 0x6c, 0xa6, 0x5b, 0x50, 0x04, 0xa5, 0x98, 0x85, 0x7a, 0xeb, 0x8b,
	0x7a, 0x64, 0x0e, 0x20, 0x1b, 0x34, 0x96, 0xa4, 0x03, 0x9d, 0x9f, 0x9c, 0xfc, 0x3d, 0x8b, 0xf7,
	0xca, 0x91, 0x6c, 0x54, 0x8f, 0x1d, 0x89, 0xf9, 0x63, 0x1b, 0x99, 0x77, 0x7f, 0x09, 0x66, 0x7f,
	0xdc, 0x12, 0x1a, 0xaa, 0xc5, 0x22, 0x89, 0xd2, 0x65, 0x3e, 0x80, 0x90, 0x48, 0x6d, 0xf6, 0x65,
	0x0b, 0x56, 0xb0, 0x20, 0x9c, 0xb1, 0x8f, 0xb1, 0x45, 0x0c, 0x31, 0xc6, 0x43, 0x8c, 0x1e, 0xee,
	0x4a, 0xfe, 0x92, 0x2d, 0x1a, 0xd5, 0x82, 0x0d, 0xd2, 0x75, 0x90, 0xe6, 0x03, 0xb1, 0x2b, 0xf9,
	0x9a, 0xcd, 0x1f, 0x90, 0x40, 0x6a, 0x4b, 0x62, 0x19, 0xb4, 0x99, 0xc7, 0x9f, 0x2c, 0xf1, 0xcf,
	0x6c, 0x59, 0x80, 0x85, 0x1f, 0x46, 0x1b, 0xd5, 0x1a, 0x70, 0xe2, 0x59, 0x32, 0x4a, 0xe3, 0xbb,
	0xd7, 0x17, 0x17, 0xf8, 0xab, 0xf7, 0xb2, 0xc5, 0x8e, 0x40, 0x7a, 0x3a, 0xff, 0x77, 0x1f, 0xff,
	0xc0, 0xae, 0xb5, 0x72, 0x35, 0x76, 0x24, 0xfd, 0x8d, 0xc5, 0xf3, 0xc7, 0x82, 0x88, 0x8f, 0xf6,
	0xaf, 0xe6, 0x00, 0xfc, 0x15, 0x8b, 0x7f, 0x81, 0xa9, 0x6a, 0x82, 0x52, 0x3e, 0x34, 0xe2, 0x45,
	0x38, 0x24, 0x3b, 0x51, 0xdf, 0x9a, 0x4d, 0xc5, 0x56, 0x0e, 0xe5, 0xde, 0xf8, 0xa4, 0xfd, 0x1c,
	0x15, 0xbd, 0x3c, 0x75, 0xe0, 0xa2, 0x6d, 0xd1, 0x93, 0xdb, 0x76, 0x31, 0x8f, 0x57, 0x97, 0xf3,
	0xb8, 0x79, 0xc7, 0xd6, 0x0e, 0x65, 0x67, 0xcd, 0xcf, 0x0e, 0xce, 0xa5, 0x8e, 0xd6, 0xff, 0x4e,
	0xf1, 0xc7, 0xf4, 0xfb, 0x9b, 0xca, 0x50, 0xdd, 0x15, 0x5b, 0x8d, 0x87, 0x4c, 0xa3, 0xd3, 0xb5,
	0x32, 0x36, 0xd3, 0x68, 0x09, 0x2c, 0xa1, 0x7b, 0x5b, 0xe1, 0xf0, 0xff, 0x14, 0xd3, 0x70, 0xb6,
	0xfb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x3f, 0x77, 0x5e, 0x53, 0x03, 0x00, 0x00,
}
