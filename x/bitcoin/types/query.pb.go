// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitcoin/v1beta1/query.proto

package types

import (
	fmt "fmt"
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

// DepositQueryParams describe the parameters used to query for a Bitcoin
// deposit address
type DepositQueryParams struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Chain   string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *DepositQueryParams) Reset()         { *m = DepositQueryParams{} }
func (m *DepositQueryParams) String() string { return proto.CompactTextString(m) }
func (*DepositQueryParams) ProtoMessage()    {}
func (*DepositQueryParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f1bd927442f8a6, []int{0}
}
func (m *DepositQueryParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositQueryParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositQueryParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositQueryParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositQueryParams.Merge(m, src)
}
func (m *DepositQueryParams) XXX_Size() int {
	return m.Size()
}
func (m *DepositQueryParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositQueryParams.DiscardUnknown(m)
}

var xxx_messageInfo_DepositQueryParams proto.InternalMessageInfo

type QueryAddressResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	KeyID   string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (m *QueryAddressResponse) Reset()         { *m = QueryAddressResponse{} }
func (m *QueryAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAddressResponse) ProtoMessage()    {}
func (*QueryAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f1bd927442f8a6, []int{1}
}
func (m *QueryAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAddressResponse.Merge(m, src)
}
func (m *QueryAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAddressResponse proto.InternalMessageInfo

type QueryDepositStatusResponse struct {
	Log    string        `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	Status OutPointState `protobuf:"varint,2,opt,name=status,proto3,enum=bitcoin.v1beta1.OutPointState" json:"status,omitempty"`
}

func (m *QueryDepositStatusResponse) Reset()         { *m = QueryDepositStatusResponse{} }
func (m *QueryDepositStatusResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDepositStatusResponse) ProtoMessage()    {}
func (*QueryDepositStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f1bd927442f8a6, []int{2}
}
func (m *QueryDepositStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositStatusResponse.Merge(m, src)
}
func (m *QueryDepositStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositStatusResponse proto.InternalMessageInfo

type QueryTxResponse struct {
	Tx                   string                         `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Status               TxStatus                       `protobuf:"varint,2,opt,name=status,proto3,enum=bitcoin.v1beta1.TxStatus" json:"status,omitempty"`
	ConfirmationRequired bool                           `protobuf:"varint,3,opt,name=confirmation_required,json=confirmationRequired,proto3" json:"confirmation_required,omitempty"`
	PrevSignedTxHash     string                         `protobuf:"bytes,4,opt,name=prev_signed_tx_hash,json=prevSignedTxHash,proto3" json:"prev_signed_tx_hash,omitempty"`
	AnyoneCanSpendVout   uint32                         `protobuf:"varint,5,opt,name=anyone_can_spend_vout,json=anyoneCanSpendVout,proto3" json:"anyone_can_spend_vout,omitempty"`
	SigningInfos         []*QueryTxResponse_SigningInfo `protobuf:"bytes,6,rep,name=signing_infos,json=signingInfos,proto3" json:"signing_infos,omitempty"`
}

func (m *QueryTxResponse) Reset()         { *m = QueryTxResponse{} }
func (m *QueryTxResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTxResponse) ProtoMessage()    {}
func (*QueryTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f1bd927442f8a6, []int{3}
}
func (m *QueryTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTxResponse.Merge(m, src)
}
func (m *QueryTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTxResponse proto.InternalMessageInfo

type QueryTxResponse_SigningInfo struct {
	RedeemScript string `protobuf:"bytes,1,opt,name=redeem_script,json=redeemScript,proto3" json:"redeem_script,omitempty"`
	Amount       int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *QueryTxResponse_SigningInfo) Reset()         { *m = QueryTxResponse_SigningInfo{} }
func (m *QueryTxResponse_SigningInfo) String() string { return proto.CompactTextString(m) }
func (*QueryTxResponse_SigningInfo) ProtoMessage()    {}
func (*QueryTxResponse_SigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47f1bd927442f8a6, []int{3, 0}
}
func (m *QueryTxResponse_SigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTxResponse_SigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTxResponse_SigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTxResponse_SigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTxResponse_SigningInfo.Merge(m, src)
}
func (m *QueryTxResponse_SigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *QueryTxResponse_SigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTxResponse_SigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTxResponse_SigningInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DepositQueryParams)(nil), "bitcoin.v1beta1.DepositQueryParams")
	proto.RegisterType((*QueryAddressResponse)(nil), "bitcoin.v1beta1.QueryAddressResponse")
	proto.RegisterType((*QueryDepositStatusResponse)(nil), "bitcoin.v1beta1.QueryDepositStatusResponse")
	proto.RegisterType((*QueryTxResponse)(nil), "bitcoin.v1beta1.QueryTxResponse")
	proto.RegisterType((*QueryTxResponse_SigningInfo)(nil), "bitcoin.v1beta1.QueryTxResponse.SigningInfo")
}

func init() { proto.RegisterFile("bitcoin/v1beta1/query.proto", fileDescriptor_47f1bd927442f8a6) }

var fileDescriptor_47f1bd927442f8a6 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x95, 0x16, 0xe6, 0xfd, 0x94, 0xe9, 0x50, 0x28, 0x52, 0xa8, 0xca, 0xa5, 0x07,
	0x96, 0xa8, 0x1b, 0xe2, 0xce, 0xd8, 0x81, 0xc2, 0x81, 0xcd, 0xad, 0x38, 0x70, 0x89, 0xdc, 0xe4,
	0x35, 0xb5, 0xba, 0xda, 0x99, 0xed, 0x94, 0xe4, 0x5f, 0xe0, 0xc4, 0x9f, 0xb5, 0xe3, 0x8e, 0x9c,
	0x10, 0xb4, 0xff, 0x08, 0x8a, 0xe3, 0xb1, 0xa9, 0x95, 0xb8, 0xf9, 0xf9, 0xf3, 0xf5, 0x27, 0x2f,
	0x4f, 0x0f, 0xbd, 0x18, 0x33, 0x1d, 0x09, 0xc6, 0x83, 0x45, 0x7f, 0x0c, 0x9a, 0xf6, 0x83, 0xeb,
	0x0c, 0x64, 0xe1, 0xa7, 0x52, 0x68, 0x81, 0x0f, 0x2c, 0xf4, 0x2d, 0x6c, 0xb7, 0x12, 0x91, 0x08,
	0xc3, 0x82, 0xf2, 0x54, 0xc5, 0xda, 0x1b, 0x0e, 0x5d, 0xa4, 0xa0, 0x2a, 0xd8, 0x3d, 0x47, 0xf8,
	0x1c, 0x52, 0xa1, 0x98, 0xbe, 0x2c, 0xcd, 0x17, 0x54, 0xd2, 0xb9, 0xc2, 0x2e, 0x7a, 0x4c, 0xe3,
	0x58, 0x82, 0x52, 0xae, 0xd3, 0x71, 0x7a, 0xdb, 0xe4, 0xae, 0xc4, 0x2d, 0xd4, 0x88, 0xa6, 0x94,
	0x71, 0x77, 0xcb, 0xdc, 0x57, 0x45, 0x97, 0xa0, 0x96, 0x79, 0xfe, 0xae, 0x4a, 0x11, 0x50, 0xa9,
	0xe0, 0x0a, 0xfe, 0xe3, 0xe9, 0xa0, 0xe6, 0x0c, 0x8a, 0x90, 0xc5, 0x95, 0xe8, 0x6c, 0x7b, 0xf9,
	0xeb, 0x65, 0xe3, 0x13, 0x14, 0x83, 0x73, 0xd2, 0x98, 0x41, 0x31, 0x88, 0xbb, 0x13, 0xd4, 0x36,
	0x4e, 0xdb, 0xde, 0x50, 0x53, 0x9d, 0xdd, 0x9b, 0x0f, 0x51, 0xfd, 0x4a, 0x24, 0xd6, 0x5a, 0x1e,
	0xf1, 0x5b, 0xd4, 0x54, 0x26, 0x63, 0x8c, 0xfb, 0x27, 0x9e, 0xbf, 0x36, 0x1e, 0xff, 0x73, 0xa6,
	0x2f, 0x04, 0xe3, 0x46, 0x05, 0xc4, 0xa6, 0xbb, 0xdf, 0xeb, 0xe8, 0xc0, 0x7c, 0x68, 0x94, 0xff,
	0xb3, 0xef, 0xa3, 0x2d, 0x9d, 0x5b, 0xf9, 0x96, 0xce, 0x71, 0x7f, 0xcd, 0xfd, 0x7c, 0xc3, 0x3d,
	0xca, 0x6d, 0x83, 0x36, 0x88, 0x4f, 0xd1, 0x51, 0x24, 0xf8, 0x84, 0xc9, 0x39, 0xd5, 0x4c, 0xf0,
	0x50, 0xc2, 0x75, 0xc6, 0x24, 0xc4, 0x6e, 0xbd, 0xe3, 0xf4, 0x9e, 0x90, 0xd6, 0x43, 0x48, 0x2c,
	0xc3, 0xc7, 0xe8, 0x69, 0x2a, 0x61, 0x11, 0x2a, 0x96, 0x70, 0x88, 0x43, 0x9d, 0x87, 0x53, 0xaa,
	0xa6, 0xee, 0x23, 0xd3, 0xc8, 0x61, 0x89, 0x86, 0x86, 0x8c, 0xf2, 0x0f, 0x54, 0x4d, 0x71, 0x1f,
	0x1d, 0x51, 0x5e, 0x08, 0x0e, 0x61, 0x44, 0x79, 0xa8, 0x52, 0xe0, 0x71, 0xb8, 0x10, 0x99, 0x76,
	0x1b, 0x1d, 0xa7, 0xb7, 0x47, 0x70, 0x05, 0xdf, 0x53, 0x3e, 0x2c, 0xd1, 0x17, 0x91, 0x69, 0x7c,
	0x89, 0xf6, 0x4a, 0x39, 0xe3, 0x49, 0xc8, 0xf8, 0x44, 0x28, 0xb7, 0xd9, 0xa9, 0xf7, 0x76, 0x4e,
	0x5e, 0x6f, 0xfc, 0xd0, 0xda, 0x48, 0xfc, 0x61, 0xf5, 0x6a, 0xc0, 0x27, 0x82, 0xec, 0xaa, 0xfb,
	0x42, 0xb5, 0x3f, 0xa2, 0x9d, 0x07, 0x10, 0xbf, 0x42, 0x7b, 0x12, 0x62, 0x80, 0x79, 0xa8, 0x22,
	0xc9, 0x52, 0x6d, 0xc7, 0xb8, 0x5b, 0x5d, 0x0e, 0xcd, 0x1d, 0x7e, 0x86, 0x9a, 0x74, 0x2e, 0x32,
	0xae, 0xcd, 0x40, 0xeb, 0xc4, 0x56, 0x67, 0xe4, 0xe6, 0x8f, 0x57, 0xbb, 0x59, 0x7a, 0xce, 0xed,
	0xd2, 0x73, 0x7e, 0x2f, 0x3d, 0xe7, 0xc7, 0xca, 0xab, 0xdd, 0xae, 0xbc, 0xda, 0xcf, 0x95, 0x57,
	0xfb, 0xfa, 0x26, 0x61, 0x7a, 0x9a, 0x8d, 0xfd, 0x48, 0xcc, 0x03, 0x9a, 0xc3, 0x15, 0x95, 0x1c,
	0xf4, 0x37, 0x21, 0x67, 0xb6, 0x3a, 0x8e, 0x84, 0x84, 0x20, 0x0f, 0xee, 0x16, 0xde, 0x2c, 0xfa,
	0xb8, 0x69, 0x36, 0xfd, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xa2, 0x5d, 0xc7, 0x4c,
	0x03, 0x00, 0x00,
}

func (m *DepositQueryParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositQueryParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositQueryParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SigningInfos) > 0 {
		for iNdEx := len(m.SigningInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SigningInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.AnyoneCanSpendVout != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.AnyoneCanSpendVout))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PrevSignedTxHash) > 0 {
		i -= len(m.PrevSignedTxHash)
		copy(dAtA[i:], m.PrevSignedTxHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PrevSignedTxHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.ConfirmationRequired {
		i--
		if m.ConfirmationRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTxResponse_SigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTxResponse_SigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTxResponse_SigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RedeemScript) > 0 {
		i -= len(m.RedeemScript)
		copy(dAtA[i:], m.RedeemScript)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RedeemScript)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositQueryParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDepositStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	return n
}

func (m *QueryTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	if m.ConfirmationRequired {
		n += 2
	}
	l = len(m.PrevSignedTxHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.AnyoneCanSpendVout != 0 {
		n += 1 + sovQuery(uint64(m.AnyoneCanSpendVout))
	}
	if len(m.SigningInfos) > 0 {
		for _, e := range m.SigningInfos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryTxResponse_SigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RedeemScript)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovQuery(uint64(m.Amount))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositQueryParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: DepositQueryParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositQueryParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDepositStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDepositStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OutPointState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TxStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ConfirmationRequired = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevSignedTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevSignedTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnyoneCanSpendVout", wireType)
			}
			m.AnyoneCanSpendVout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AnyoneCanSpendVout |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigningInfos = append(m.SigningInfos, &QueryTxResponse_SigningInfo{})
			if err := m.SigningInfos[len(m.SigningInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTxResponse_SigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemScript", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedeemScript = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
