// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
	tofnd "github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	exported1 "github.com/axelarnetwork/axelar-core/x/vote/exported"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgAssignNextKey represents a message to assign a new key
type MsgAssignNextKey struct {
	Sender  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Chain   string                                        `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	KeyID   string                                        `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	KeyRole exported.KeyRole                              `protobuf:"varint,4,opt,name=key_role,json=keyRole,proto3,enum=tss.exported.v1beta1.KeyRole" json:"key_role,omitempty"`
}

func (m *MsgAssignNextKey) Reset()         { *m = MsgAssignNextKey{} }
func (m *MsgAssignNextKey) String() string { return proto.CompactTextString(m) }
func (*MsgAssignNextKey) ProtoMessage()    {}
func (*MsgAssignNextKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{0}
}
func (m *MsgAssignNextKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgAssignNextKey.Unmarshal(m, b)
}
func (m *MsgAssignNextKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgAssignNextKey.Marshal(b, m, deterministic)
}
func (m *MsgAssignNextKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAssignNextKey.Merge(m, src)
}
func (m *MsgAssignNextKey) XXX_Size() int {
	return xxx_messageInfo_MsgAssignNextKey.Size(m)
}
func (m *MsgAssignNextKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAssignNextKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAssignNextKey proto.InternalMessageInfo

// MsgDeregister to deregister so that the validator will not participate in any
// future keygen
type MsgDeregister struct {
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
}

func (m *MsgDeregister) Reset()         { *m = MsgDeregister{} }
func (m *MsgDeregister) String() string { return proto.CompactTextString(m) }
func (*MsgDeregister) ProtoMessage()    {}
func (*MsgDeregister) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{1}
}
func (m *MsgDeregister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgDeregister.Unmarshal(m, b)
}
func (m *MsgDeregister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgDeregister.Marshal(b, m, deterministic)
}
func (m *MsgDeregister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeregister.Merge(m, src)
}
func (m *MsgDeregister) XXX_Size() int {
	return xxx_messageInfo_MsgDeregister.Size(m)
}
func (m *MsgDeregister) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeregister.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeregister proto.InternalMessageInfo

// MsgKeygenStart indicate the start of keygen
type MsgKeygenStart struct {
	Sender     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	NewKeyID   string                                        `protobuf:"bytes,2,opt,name=new_key_id,json=newKeyId,proto3" json:"new_key_id,omitempty"`
	SubsetSize int64                                         `protobuf:"varint,3,opt,name=subset_size,json=subsetSize,proto3" json:"subset_size,omitempty"`
}

func (m *MsgKeygenStart) Reset()         { *m = MsgKeygenStart{} }
func (m *MsgKeygenStart) String() string { return proto.CompactTextString(m) }
func (*MsgKeygenStart) ProtoMessage()    {}
func (*MsgKeygenStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{2}
}
func (m *MsgKeygenStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgKeygenStart.Unmarshal(m, b)
}
func (m *MsgKeygenStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgKeygenStart.Marshal(b, m, deterministic)
}
func (m *MsgKeygenStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgKeygenStart.Merge(m, src)
}
func (m *MsgKeygenStart) XXX_Size() int {
	return xxx_messageInfo_MsgKeygenStart.Size(m)
}
func (m *MsgKeygenStart) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgKeygenStart.DiscardUnknown(m)
}

var xxx_messageInfo_MsgKeygenStart proto.InternalMessageInfo

type MsgRotateKey struct {
	Sender     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Chain      string                                        `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	SubsetSize int64                                         `protobuf:"varint,3,opt,name=subset_size,json=subsetSize,proto3" json:"subset_size,omitempty"`
	KeyRole    exported.KeyRole                              `protobuf:"varint,4,opt,name=key_role,json=keyRole,proto3,enum=tss.exported.v1beta1.KeyRole" json:"key_role,omitempty"`
}

func (m *MsgRotateKey) Reset()         { *m = MsgRotateKey{} }
func (m *MsgRotateKey) String() string { return proto.CompactTextString(m) }
func (*MsgRotateKey) ProtoMessage()    {}
func (*MsgRotateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{3}
}
func (m *MsgRotateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRotateKey.Unmarshal(m, b)
}
func (m *MsgRotateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRotateKey.Marshal(b, m, deterministic)
}
func (m *MsgRotateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRotateKey.Merge(m, src)
}
func (m *MsgRotateKey) XXX_Size() int {
	return xxx_messageInfo_MsgRotateKey.Size(m)
}
func (m *MsgRotateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRotateKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRotateKey proto.InternalMessageInfo

// MsgKeygenTraffic protocol message
type MsgKeygenTraffic struct {
	Sender    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	SessionID string                                        `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Payload   *tofnd.TrafficOut                             `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgKeygenTraffic) Reset()         { *m = MsgKeygenTraffic{} }
func (m *MsgKeygenTraffic) String() string { return proto.CompactTextString(m) }
func (*MsgKeygenTraffic) ProtoMessage()    {}
func (*MsgKeygenTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{4}
}
func (m *MsgKeygenTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgKeygenTraffic.Unmarshal(m, b)
}
func (m *MsgKeygenTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgKeygenTraffic.Marshal(b, m, deterministic)
}
func (m *MsgKeygenTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgKeygenTraffic.Merge(m, src)
}
func (m *MsgKeygenTraffic) XXX_Size() int {
	return xxx_messageInfo_MsgKeygenTraffic.Size(m)
}
func (m *MsgKeygenTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgKeygenTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_MsgKeygenTraffic proto.InternalMessageInfo

// MsgSignTraffic protocol message
type MsgSignTraffic struct {
	Sender    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	SessionID string                                        `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Payload   *tofnd.TrafficOut                             `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgSignTraffic) Reset()         { *m = MsgSignTraffic{} }
func (m *MsgSignTraffic) String() string { return proto.CompactTextString(m) }
func (*MsgSignTraffic) ProtoMessage()    {}
func (*MsgSignTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{5}
}
func (m *MsgSignTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSignTraffic.Unmarshal(m, b)
}
func (m *MsgSignTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSignTraffic.Marshal(b, m, deterministic)
}
func (m *MsgSignTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignTraffic.Merge(m, src)
}
func (m *MsgSignTraffic) XXX_Size() int {
	return xxx_messageInfo_MsgSignTraffic.Size(m)
}
func (m *MsgSignTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignTraffic proto.InternalMessageInfo

// MsgVotePubKey represents the message to vote on a public key
type MsgVotePubKey struct {
	Sender   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	PollMeta exported1.PollMeta                            `protobuf:"bytes,2,opt,name=poll_meta,json=pollMeta,proto3" json:"poll_meta"`
	// need to vote on the bytes instead of ecdsa.PublicKey, otherwise we lose the
	// elliptic curve information
	PubKeyBytes []byte `protobuf:"bytes,3,opt,name=pub_key_bytes,json=pubKeyBytes,proto3" json:"pub_key_bytes,omitempty"`
}

func (m *MsgVotePubKey) Reset()         { *m = MsgVotePubKey{} }
func (m *MsgVotePubKey) String() string { return proto.CompactTextString(m) }
func (*MsgVotePubKey) ProtoMessage()    {}
func (*MsgVotePubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{6}
}
func (m *MsgVotePubKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgVotePubKey.Unmarshal(m, b)
}
func (m *MsgVotePubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgVotePubKey.Marshal(b, m, deterministic)
}
func (m *MsgVotePubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVotePubKey.Merge(m, src)
}
func (m *MsgVotePubKey) XXX_Size() int {
	return xxx_messageInfo_MsgVotePubKey.Size(m)
}
func (m *MsgVotePubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVotePubKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVotePubKey proto.InternalMessageInfo

// MsgVoteSig represents a message to vote for a signature
type MsgVoteSig struct {
	Sender   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	PollMeta exported1.PollMeta                            `protobuf:"bytes,2,opt,name=poll_meta,json=pollMeta,proto3" json:"poll_meta"`
	// need to vote on the bytes instead of r, s, because Go cannot deserialize
	// private fields using reflection (so *big.Int does not work)
	SigBytes []byte `protobuf:"bytes,3,opt,name=sig_bytes,json=sigBytes,proto3" json:"sig_bytes,omitempty"`
}

func (m *MsgVoteSig) Reset()         { *m = MsgVoteSig{} }
func (m *MsgVoteSig) String() string { return proto.CompactTextString(m) }
func (*MsgVoteSig) ProtoMessage()    {}
func (*MsgVoteSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d13e1023e3ffaf, []int{7}
}
func (m *MsgVoteSig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgVoteSig.Unmarshal(m, b)
}
func (m *MsgVoteSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgVoteSig.Marshal(b, m, deterministic)
}
func (m *MsgVoteSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteSig.Merge(m, src)
}
func (m *MsgVoteSig) XXX_Size() int {
	return xxx_messageInfo_MsgVoteSig.Size(m)
}
func (m *MsgVoteSig) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteSig.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteSig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAssignNextKey)(nil), "tss.v1beta1.MsgAssignNextKey")
	proto.RegisterType((*MsgDeregister)(nil), "tss.v1beta1.MsgDeregister")
	proto.RegisterType((*MsgKeygenStart)(nil), "tss.v1beta1.MsgKeygenStart")
	proto.RegisterType((*MsgRotateKey)(nil), "tss.v1beta1.MsgRotateKey")
	proto.RegisterType((*MsgKeygenTraffic)(nil), "tss.v1beta1.MsgKeygenTraffic")
	proto.RegisterType((*MsgSignTraffic)(nil), "tss.v1beta1.MsgSignTraffic")
	proto.RegisterType((*MsgVotePubKey)(nil), "tss.v1beta1.MsgVotePubKey")
	proto.RegisterType((*MsgVoteSig)(nil), "tss.v1beta1.MsgVoteSig")
}

func init() { proto.RegisterFile("tss/v1beta1/tx.proto", fileDescriptor_58d13e1023e3ffaf) }

var fileDescriptor_58d13e1023e3ffaf = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xed, 0x5f, 0xff, 0xc5, 0x9b, 0xb4, 0xfa, 0xc9, 0xea, 0x21, 0x2a, 0xaa, 0x1d, 0x7c,
	0xaa, 0x10, 0x4d, 0x68, 0x39, 0xc0, 0xb5, 0x56, 0x2f, 0x51, 0x94, 0x52, 0x39, 0x88, 0x43, 0x2f,
	0x91, 0xff, 0x4c, 0xb7, 0xab, 0xb8, 0x5e, 0x6b, 0x67, 0xd3, 0xc6, 0x7d, 0x0a, 0x1e, 0x83, 0x47,
	0xe0, 0x8c, 0x38, 0x54, 0x5c, 0xe8, 0x91, 0x53, 0x04, 0xe9, 0x5b, 0x70, 0x42, 0xf6, 0x6e, 0xd4,
	0x02, 0x12, 0x1c, 0x5a, 0x21, 0x38, 0x65, 0x67, 0x76, 0x32, 0xfa, 0x7e, 0xbe, 0x3b, 0xeb, 0x25,
	0xeb, 0x12, 0xb1, 0x73, 0xb6, 0x13, 0x81, 0x0c, 0x77, 0x3a, 0x72, 0xd2, 0xce, 0x05, 0x97, 0xdc,
	0xae, 0x4b, 0xc4, 0xb6, 0xce, 0x6e, 0xac, 0x53, 0x4e, 0x79, 0x95, 0xef, 0x94, 0x2b, 0x55, 0xb2,
	0xd1, 0x2a, 0xff, 0x08, 0x93, 0x9c, 0x0b, 0x09, 0xc9, 0x4d, 0x87, 0x22, 0x07, 0xd4, 0x15, 0x9b,
	0x65, 0x85, 0xe4, 0xc7, 0xd9, 0xad, 0xed, 0x32, 0xd2, 0xdb, 0x0f, 0xcf, 0xb8, 0x84, 0x5f, 0x76,
	0xf0, 0x3e, 0x9a, 0xe4, 0xff, 0x3e, 0xd2, 0x3d, 0x44, 0x46, 0xb3, 0x03, 0x98, 0xc8, 0x1e, 0x14,
	0x76, 0x97, 0x2c, 0x23, 0x64, 0x09, 0x88, 0xa6, 0xd9, 0x32, 0xb7, 0x1a, 0xfe, 0xce, 0xd7, 0xa9,
	0xbb, 0x4d, 0x99, 0x3c, 0x19, 0x47, 0xed, 0x98, 0x9f, 0x76, 0x62, 0x8e, 0xa7, 0x1c, 0xf5, 0xcf,
	0x36, 0x26, 0x23, 0xdd, 0x72, 0x2f, 0x8e, 0xf7, 0x92, 0x44, 0x00, 0x62, 0xa0, 0x1b, 0xd8, 0xeb,
	0x64, 0x29, 0x3e, 0x09, 0x59, 0xd6, 0xfc, 0xaf, 0x65, 0x6e, 0x59, 0x81, 0x0a, 0xec, 0x16, 0x59,
	0x1e, 0x41, 0x31, 0x64, 0x49, 0x73, 0xa1, 0x4c, 0xfb, 0xd6, 0x6c, 0xea, 0x2e, 0xf5, 0xa0, 0xe8,
	0xee, 0x07, 0x4b, 0x23, 0x28, 0xba, 0x89, 0xfd, 0x9c, 0xd4, 0xca, 0x0a, 0xc1, 0x53, 0x68, 0x2e,
	0xb6, 0xcc, 0xad, 0xb5, 0xdd, 0xcd, 0x76, 0xe9, 0xd8, 0x1c, 0x66, 0x6e, 0x5d, 0xbb, 0x07, 0x45,
	0xc0, 0x53, 0x08, 0x56, 0x46, 0x6a, 0xe1, 0x1d, 0x91, 0xd5, 0x3e, 0xd2, 0x7d, 0x10, 0x40, 0x19,
	0x4a, 0x10, 0x3f, 0xd0, 0x58, 0x77, 0xa0, 0xf1, 0xde, 0x98, 0x64, 0xad, 0x8f, 0xb4, 0x07, 0x05,
	0x85, 0x6c, 0x20, 0x43, 0x21, 0xef, 0xb1, 0xbb, 0xfd, 0x88, 0x90, 0x0c, 0xce, 0x87, 0xda, 0x99,
	0xca, 0x30, 0xbf, 0x31, 0x9b, 0xba, 0xb5, 0x03, 0x38, 0x57, 0xe6, 0xd4, 0x32, 0xb5, 0x4a, 0x6c,
	0x97, 0xd4, 0x71, 0x1c, 0x21, 0xc8, 0x21, 0xb2, 0x0b, 0xa8, 0x6c, 0x5c, 0x08, 0x88, 0x4a, 0x0d,
	0xd8, 0x05, 0x78, 0x1f, 0x4c, 0xd2, 0xe8, 0x23, 0x0d, 0xb8, 0x0c, 0x25, 0xfc, 0x91, 0x43, 0xfd,
	0x9d, 0xa4, 0x3b, 0x9c, 0xe9, 0x7b, 0x35, 0xa5, 0xca, 0xf7, 0x97, 0x22, 0x3c, 0x3e, 0x66, 0xf1,
	0x7d, 0x02, 0x3d, 0x26, 0x04, 0x01, 0x91, 0xf1, 0xec, 0xc6, 0xf9, 0xd5, 0xd9, 0xd4, 0xb5, 0x06,
	0x2a, 0xdb, 0xdd, 0x0f, 0x2c, 0x5d, 0xd0, 0x4d, 0xec, 0x67, 0x64, 0x25, 0x0f, 0x8b, 0x94, 0x87,
	0x6a, 0x7c, 0xeb, 0x1a, 0x43, 0xdd, 0xbc, 0x39, 0x83, 0x56, 0xf9, 0x62, 0x2c, 0x83, 0x79, 0xb5,
	0xf7, 0x4e, 0x8d, 0xcf, 0x80, 0xd1, 0x7f, 0x1b, 0xa2, 0xbc, 0x60, 0xaf, 0xb8, 0x84, 0xc3, 0x71,
	0x74, 0xcf, 0x93, 0xe5, 0x13, 0x2b, 0xe7, 0x69, 0x3a, 0x3c, 0x05, 0x19, 0x56, 0x08, 0xf5, 0x5d,
	0xb7, 0x5d, 0x7e, 0xc5, 0x7e, 0x1e, 0x92, 0x43, 0x9e, 0xa6, 0x7d, 0x90, 0xa1, 0xbf, 0x78, 0x39,
	0x75, 0x8d, 0xa0, 0x96, 0xeb, 0xd8, 0xf6, 0xc8, 0x6a, 0x3e, 0x8e, 0xaa, 0x6b, 0x14, 0x15, 0x12,
	0xb0, 0xe2, 0x6b, 0x04, 0xf5, 0xbc, 0x52, 0xeb, 0x97, 0x29, 0xef, 0xad, 0x49, 0x88, 0x86, 0x18,
	0x30, 0xfa, 0xb7, 0x11, 0x3c, 0x20, 0x16, 0x32, 0xfa, 0x9d, 0xfa, 0x1a, 0x32, 0x5a, 0x49, 0xf7,
	0x0f, 0x2e, 0xbf, 0x38, 0xc6, 0xe5, 0xcc, 0x31, 0xae, 0x66, 0x8e, 0xf1, 0x79, 0xe6, 0x18, 0xaf,
	0xaf, 0x1d, 0xe3, 0xea, 0xda, 0x31, 0x3e, 0x5d, 0x3b, 0xc6, 0xd1, 0x93, 0x5b, 0xaa, 0xc3, 0x09,
	0xa4, 0xa1, 0xc8, 0x40, 0x9e, 0x73, 0x31, 0xd2, 0xd1, 0x76, 0xcc, 0x05, 0x74, 0x26, 0x9d, 0xea,
	0xe1, 0x28, 0x19, 0xa2, 0xe5, 0xea, 0x21, 0x78, 0xfa, 0x2d, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x8d,
	0x06, 0x5c, 0xa7, 0x06, 0x00, 0x00,
}
