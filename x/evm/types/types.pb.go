// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/v1beta1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// NetworkInfo describes information about a network
type NetworkInfo struct {
	Name string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=id,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"id"`
}

func (m *NetworkInfo) Reset()         { *m = NetworkInfo{} }
func (m *NetworkInfo) String() string { return proto.CompactTextString(m) }
func (*NetworkInfo) ProtoMessage()    {}
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{0}
}
func (m *NetworkInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfo.Merge(m, src)
}
func (m *NetworkInfo) XXX_Size() int {
	return m.Size()
}
func (m *NetworkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfo proto.InternalMessageInfo

// BurnerInfo describes information required to burn token at an burner address
// that is deposited by an user
type BurnerInfo struct {
	TokenAddress Address `protobuf:"bytes,1,opt,name=token_address,json=tokenAddress,proto3,customtype=Address" json:"token_address"`
	Symbol       string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Salt         Hash    `protobuf:"bytes,3,opt,name=salt,proto3,customtype=Hash" json:"salt"`
}

func (m *BurnerInfo) Reset()         { *m = BurnerInfo{} }
func (m *BurnerInfo) String() string { return proto.CompactTextString(m) }
func (*BurnerInfo) ProtoMessage()    {}
func (*BurnerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{1}
}
func (m *BurnerInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnerInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnerInfo.Merge(m, src)
}
func (m *BurnerInfo) XXX_Size() int {
	return m.Size()
}
func (m *BurnerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BurnerInfo proto.InternalMessageInfo

// ERC20Deposit contains information for an ERC20 deposit
type ERC20Deposit struct {
	TxID          Hash                                    `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
	Amount        github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"amount"`
	Symbol        string                                  `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	BurnerAddress Address                                 `protobuf:"bytes,4,opt,name=burner_address,json=burnerAddress,proto3,customtype=Address" json:"burner_address"`
}

func (m *ERC20Deposit) Reset()         { *m = ERC20Deposit{} }
func (m *ERC20Deposit) String() string { return proto.CompactTextString(m) }
func (*ERC20Deposit) ProtoMessage()    {}
func (*ERC20Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{2}
}
func (m *ERC20Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20Deposit.Merge(m, src)
}
func (m *ERC20Deposit) XXX_Size() int {
	return m.Size()
}
func (m *ERC20Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20Deposit proto.InternalMessageInfo

// ERC20TokenDeployment describes information about an ERC20 token
type ERC20TokenDeployment struct {
	Symbol       string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	TokenAddress Address `protobuf:"bytes,2,opt,name=token_address,json=tokenAddress,proto3,customtype=Address" json:"token_address"`
}

func (m *ERC20TokenDeployment) Reset()         { *m = ERC20TokenDeployment{} }
func (m *ERC20TokenDeployment) String() string { return proto.CompactTextString(m) }
func (*ERC20TokenDeployment) ProtoMessage()    {}
func (*ERC20TokenDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{3}
}
func (m *ERC20TokenDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20TokenDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20TokenDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20TokenDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20TokenDeployment.Merge(m, src)
}
func (m *ERC20TokenDeployment) XXX_Size() int {
	return m.Size()
}
func (m *ERC20TokenDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20TokenDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20TokenDeployment proto.InternalMessageInfo

// TransferOwnership contains information for a transfer ownership
type TransferOwnership struct {
	TxID      Hash   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
	NextKeyID string `protobuf:"bytes,3,opt,name=next_key_id,json=nextKeyId,proto3" json:"next_key_id,omitempty"`
}

func (m *TransferOwnership) Reset()         { *m = TransferOwnership{} }
func (m *TransferOwnership) String() string { return proto.CompactTextString(m) }
func (*TransferOwnership) ProtoMessage()    {}
func (*TransferOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2cf809b4baed32, []int{4}
}
func (m *TransferOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferOwnership.Merge(m, src)
}
func (m *TransferOwnership) XXX_Size() int {
	return m.Size()
}
func (m *TransferOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_TransferOwnership proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NetworkInfo)(nil), "evm.v1beta1.NetworkInfo")
	proto.RegisterType((*BurnerInfo)(nil), "evm.v1beta1.BurnerInfo")
	proto.RegisterType((*ERC20Deposit)(nil), "evm.v1beta1.ERC20Deposit")
	proto.RegisterType((*ERC20TokenDeployment)(nil), "evm.v1beta1.ERC20TokenDeployment")
	proto.RegisterType((*TransferOwnership)(nil), "evm.v1beta1.TransferOwnership")
}

func init() { proto.RegisterFile("evm/v1beta1/types.proto", fileDescriptor_af2cf809b4baed32) }

var fileDescriptor_af2cf809b4baed32 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0xab, 0xd3, 0x4e,
	0x10, 0xc7, 0x93, 0xbe, 0xfc, 0xfa, 0xa3, 0xdb, 0x54, 0x31, 0x14, 0x2d, 0x1e, 0xd2, 0x92, 0x83,
	0x3e, 0x0f, 0x6d, 0xde, 0x53, 0xf1, 0x28, 0x18, 0x2b, 0x1a, 0x84, 0x0a, 0xa1, 0x5e, 0xbc, 0x94,
	0x4d, 0x77, 0x5e, 0x1b, 0xda, 0xdd, 0x0d, 0xbb, 0xdb, 0xbe, 0x04, 0xfc, 0x23, 0xfc, 0xb3, 0x7a,
	0x7c, 0x78, 0x12, 0x0f, 0x45, 0xd3, 0x7f, 0x44, 0xba, 0x5d, 0xa4, 0x45, 0x41, 0x3d, 0x65, 0x66,
	0x77, 0x66, 0xbe, 0x33, 0x9f, 0xcc, 0xa2, 0x7b, 0xb0, 0xa6, 0xe1, 0xfa, 0x32, 0x05, 0x85, 0x2f,
	0x43, 0x55, 0xe6, 0x20, 0x07, 0xb9, 0xe0, 0x8a, 0x7b, 0x4d, 0x58, 0xd3, 0x81, 0xb9, 0xb8, 0xdf,
	0x9e, 0xf1, 0x19, 0xd7, 0xe7, 0xe1, 0xde, 0x3a, 0x84, 0x04, 0x18, 0x35, 0x47, 0xa0, 0xae, 0xb9,
	0x58, 0xc4, 0xec, 0x8a, 0x7b, 0x1e, 0x72, 0x18, 0xa6, 0xd0, 0xb1, 0x7b, 0xf6, 0x79, 0x23, 0xd1,
	0xb6, 0xf7, 0x1c, 0xd5, 0x32, 0xd2, 0xa9, 0xf5, 0xec, 0x73, 0x37, 0x1a, 0x6c, 0xb6, 0x5d, 0xeb,
	0xeb, 0xb6, 0xfb, 0x60, 0x96, 0xa9, 0xf9, 0x2a, 0x1d, 0x4c, 0x39, 0x0d, 0xa7, 0x5c, 0x52, 0x2e,
	0xcd, 0xa7, 0x2f, 0xc9, 0xc2, 0xf4, 0x10, 0x33, 0x95, 0xd4, 0x32, 0x12, 0x7c, 0x44, 0x28, 0x5a,
	0x09, 0x06, 0x42, 0x2b, 0x3c, 0x45, 0x2d, 0xc5, 0x17, 0xc0, 0x26, 0x98, 0x10, 0x01, 0x52, 0x6a,
	0x29, 0x37, 0xba, 0x6d, 0x0a, 0xff, 0xff, 0xe2, 0x70, 0x9c, 0xb8, 0x3a, 0xca, 0x78, 0xde, 0x5d,
	0x54, 0x97, 0x25, 0x4d, 0xf9, 0x52, 0xf7, 0xd1, 0x48, 0x8c, 0xe7, 0xf5, 0x90, 0x23, 0xf1, 0x52,
	0x75, 0xce, 0x74, 0x11, 0xd7, 0x14, 0x71, 0xde, 0x60, 0x39, 0x4f, 0xf4, 0x4d, 0xf0, 0xd9, 0x46,
	0xee, 0xab, 0xe4, 0xe5, 0xe3, 0x8b, 0x21, 0xe4, 0x5c, 0x66, 0xca, 0x7b, 0x84, 0xfe, 0x53, 0xc5,
	0x24, 0x23, 0x46, 0xb8, 0x7d, 0x9c, 0x53, 0x6d, 0xbb, 0xce, 0xb8, 0x88, 0x87, 0x89, 0xa3, 0x8a,
	0x98, 0x78, 0xaf, 0x51, 0x1d, 0x53, 0xbe, 0x62, 0xca, 0x4c, 0x1f, 0x9a, 0xd8, 0x87, 0x7f, 0x31,
	0xfd, 0xfb, 0x8c, 0xa9, 0xc4, 0xa4, 0x1f, 0xb5, 0x7f, 0x76, 0xd2, 0xfe, 0x33, 0x74, 0x2b, 0xd5,
	0x68, 0x7e, 0xd2, 0x70, 0x7e, 0x4f, 0xa3, 0x75, 0x08, 0x33, 0x6e, 0x40, 0x50, 0x5b, 0xcf, 0x34,
	0xde, 0x33, 0x1a, 0x42, 0xbe, 0xe4, 0x25, 0x85, 0x13, 0x1d, 0xfb, 0x44, 0xe7, 0x17, 0xe8, 0x9a,
	0xe2, 0x1f, 0xa0, 0x07, 0x14, 0xdd, 0x19, 0x0b, 0xcc, 0xe4, 0x15, 0x88, 0x77, 0xd7, 0x0c, 0x84,
	0x9c, 0x67, 0xf9, 0xbf, 0xe0, 0xeb, 0xa3, 0x26, 0x83, 0x42, 0x4d, 0x16, 0x50, 0xee, 0x13, 0xf4,
	0xe8, 0x51, 0xab, 0xda, 0x76, 0x1b, 0x23, 0x28, 0xd4, 0x5b, 0x28, 0xe3, 0x61, 0xd2, 0x60, 0xc6,
	0x24, 0xd1, 0x68, 0xf3, 0xdd, 0xb7, 0x36, 0x95, 0x6f, 0xdf, 0x54, 0xbe, 0xfd, 0xad, 0xf2, 0xed,
	0x4f, 0x3b, 0xdf, 0xba, 0xd9, 0xf9, 0xd6, 0x97, 0x9d, 0x6f, 0x7d, 0xb8, 0x38, 0x62, 0x8e, 0x0b,
	0x58, 0x62, 0xc1, 0x0e, 0x7b, 0x6b, 0xbc, 0xfe, 0x94, 0x0b, 0x08, 0x8b, 0x70, 0xff, 0x16, 0xf4,
	0x1f, 0x48, 0xeb, 0x7a, 0xc3, 0x9f, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xe5, 0xa6, 0xfe,
	0x1f, 0x03, 0x00, 0x00,
}

func (m *NetworkInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Id.Size()
		i -= size
		if _, err := m.Id.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BurnerInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnerInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnerInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Salt.Size()
		i -= size
		if _, err := m.Salt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.TokenAddress.Size()
		i -= size
		if _, err := m.TokenAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ERC20Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BurnerAddress.Size()
		i -= size
		if _, err := m.BurnerAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.TxID.Size()
		i -= size
		if _, err := m.TxID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ERC20TokenDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20TokenDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20TokenDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenAddress.Size()
		i -= size
		if _, err := m.TokenAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransferOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferOwnership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferOwnership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NextKeyID) > 0 {
		i -= len(m.NextKeyID)
		copy(dAtA[i:], m.NextKeyID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextKeyID)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.TxID.Size()
		i -= size
		if _, err := m.TxID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetworkInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Id.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *BurnerInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Salt.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *ERC20Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TxID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.BurnerAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *ERC20TokenDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.TokenAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *TransferOwnership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TxID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.NextKeyID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BurnerInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BurnerInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnerInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Salt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ERC20Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ERC20Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnerAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ERC20TokenDeployment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ERC20TokenDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20TokenDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TransferOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransferOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
