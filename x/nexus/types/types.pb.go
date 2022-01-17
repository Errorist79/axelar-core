// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nexus/v1beta1/types.proto

package types

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// ChainState represents the state of a registered blockchain
type ChainState struct {
	Chain       exported.Chain                                  `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain"`
	Maintainers []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,rep,name=maintainers,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"maintainers,omitempty"`
	Activated   bool                                            `protobuf:"varint,3,opt,name=activated,proto3" json:"activated,omitempty"`
	Assets      []string                                        `protobuf:"bytes,5,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (m *ChainState) Reset()         { *m = ChainState{} }
func (m *ChainState) String() string { return proto.CompactTextString(m) }
func (*ChainState) ProtoMessage()    {}
func (*ChainState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1651b8508c88d62f, []int{0}
}
func (m *ChainState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainState.Merge(m, src)
}
func (m *ChainState) XXX_Size() int {
	return m.Size()
}
func (m *ChainState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainState.DiscardUnknown(m)
}

var xxx_messageInfo_ChainState proto.InternalMessageInfo

type LinkedAddresses struct {
	DepositAddress   exported.CrossChainAddress `protobuf:"bytes,1,opt,name=deposit_address,json=depositAddress,proto3" json:"deposit_address"`
	RecipientAddress exported.CrossChainAddress `protobuf:"bytes,2,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address"`
}

func (m *LinkedAddresses) Reset()         { *m = LinkedAddresses{} }
func (m *LinkedAddresses) String() string { return proto.CompactTextString(m) }
func (*LinkedAddresses) ProtoMessage()    {}
func (*LinkedAddresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_1651b8508c88d62f, []int{1}
}
func (m *LinkedAddresses) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkedAddresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkedAddresses.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkedAddresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkedAddresses.Merge(m, src)
}
func (m *LinkedAddresses) XXX_Size() int {
	return m.Size()
}
func (m *LinkedAddresses) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkedAddresses.DiscardUnknown(m)
}

var xxx_messageInfo_LinkedAddresses proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainState)(nil), "nexus.v1beta1.ChainState")
	proto.RegisterType((*LinkedAddresses)(nil), "nexus.v1beta1.LinkedAddresses")
}

func init() { proto.RegisterFile("nexus/v1beta1/types.proto", fileDescriptor_1651b8508c88d62f) }

var fileDescriptor_1651b8508c88d62f = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0xaf, 0xd2, 0x40,
	0x14, 0xc5, 0x3b, 0x0f, 0x78, 0x79, 0x6f, 0x9e, 0x0a, 0x36, 0xc6, 0x54, 0xa2, 0x43, 0xc3, 0xaa,
	0x2e, 0x68, 0x03, 0xae, 0x5c, 0x8a, 0x3b, 0xe3, 0xc2, 0x94, 0xc4, 0x18, 0x63, 0x62, 0xa6, 0xed,
	0x0d, 0x4c, 0x80, 0x99, 0x66, 0xee, 0x80, 0xf5, 0x5b, 0xf8, 0xb1, 0x70, 0xc7, 0x92, 0x15, 0x51,
	0xf8, 0x16, 0xae, 0x0c, 0xed, 0xf0, 0x67, 0xa1, 0x9b, 0xb7, 0x6a, 0xef, 0x39, 0x9d, 0x5f, 0xcf,
	0xb9, 0x2d, 0x7d, 0x26, 0xa1, 0x58, 0x60, 0xb4, 0xec, 0x27, 0x60, 0x78, 0x3f, 0x32, 0xdf, 0x73,
	0xc0, 0x30, 0xd7, 0xca, 0x28, 0xf7, 0x61, 0x69, 0x85, 0xd6, 0x6a, 0x3f, 0x19, 0xab, 0xb1, 0x2a,
	0x9d, 0xe8, 0x70, 0x57, 0x3d, 0xd4, 0x66, 0xa9, 0xc2, 0xb9, 0xc2, 0x28, 0xe1, 0x08, 0x27, 0x4a,
	0xaa, 0x84, 0xb4, 0x7e, 0xb7, 0xe2, 0x43, 0x91, 0x2b, 0x6d, 0x20, 0xfb, 0xd7, 0x8b, 0xba, 0x1b,
	0x42, 0xe9, 0xdb, 0x09, 0x17, 0x72, 0x64, 0xb8, 0x01, 0xf7, 0x35, 0x6d, 0xa4, 0x87, 0xc9, 0x23,
	0x3e, 0x09, 0xee, 0x06, 0x2f, 0xc2, 0x2a, 0xc7, 0x11, 0x71, 0x0c, 0x14, 0x96, 0x47, 0x86, 0xf5,
	0xd5, 0xb6, 0xe3, 0xc4, 0xd5, 0x09, 0x77, 0x44, 0xef, 0xe6, 0x5c, 0x48, 0xc3, 0x85, 0x04, 0x8d,
	0xde, 0x95, 0x5f, 0x0b, 0x1e, 0x0c, 0xfb, 0x7f, 0xb6, 0x9d, 0xde, 0x58, 0x98, 0xc9, 0x22, 0x09,
	0x53, 0x35, 0x8f, 0x6c, 0xe2, 0xea, 0xd2, 0xc3, 0x6c, 0x6a, 0xc3, 0x7c, 0xe4, 0xb3, 0x37, 0x59,
	0xa6, 0x01, 0x31, 0xbe, 0xa4, 0xb8, 0xcf, 0xe9, 0x2d, 0x4f, 0x8d, 0x58, 0x72, 0x03, 0x99, 0x57,
	0xf3, 0x49, 0x70, 0x13, 0x9f, 0x05, 0xf7, 0x29, 0xbd, 0xe6, 0x88, 0x60, 0xd0, 0x6b, 0xf8, 0xb5,
	0xe0, 0x36, 0xb6, 0xd3, 0xbb, 0xfa, 0x4d, 0xbd, 0xd5, 0xe8, 0xfe, 0x24, 0xb4, 0xf9, 0x5e, 0xc8,
	0x29, 0x64, 0x16, 0x0d, 0xe8, 0x7e, 0xa2, 0xcd, 0x0c, 0x72, 0x85, 0xc2, 0x7c, 0xe5, 0x95, 0x68,
	0x9b, 0xbe, 0xfc, 0x6f, 0x53, 0xad, 0x10, 0xcb, 0xba, 0x96, 0x62, 0x5b, 0x3f, 0xb2, 0x1c, 0xab,
	0xba, 0x5f, 0xe8, 0x63, 0x0d, 0xa9, 0xc8, 0x05, 0xc8, 0x33, 0xfb, 0xea, 0x7e, 0xec, 0xd6, 0x89,
	0x74, 0xd4, 0x3f, 0xac, 0x7e, 0x33, 0x67, 0xb5, 0x63, 0x64, 0xbd, 0x63, 0xe4, 0xd7, 0x8e, 0x91,
	0x1f, 0x7b, 0xe6, 0xac, 0xf7, 0xcc, 0xd9, 0xec, 0x99, 0xf3, 0x79, 0x70, 0xb1, 0x61, 0x5e, 0xc0,
	0x8c, 0x6b, 0x09, 0xe6, 0x9b, 0xd2, 0x53, 0x3b, 0xf5, 0x52, 0xa5, 0x21, 0x2a, 0xa2, 0xea, 0x7f,
	0x28, 0x37, 0x9e, 0x5c, 0x97, 0xdf, 0xff, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x7a,
	0x52, 0x02, 0x85, 0x02, 0x00, 0x00,
}

func (m *ChainState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Assets[iNdEx])
			copy(dAtA[i:], m.Assets[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Assets[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Activated {
		i--
		if m.Activated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Maintainers) > 0 {
		for iNdEx := len(m.Maintainers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Maintainers[iNdEx])
			copy(dAtA[i:], m.Maintainers[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Maintainers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LinkedAddresses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkedAddresses) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkedAddresses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RecipientAddress.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.DepositAddress.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
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
func (m *ChainState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Chain.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Maintainers) > 0 {
		for _, b := range m.Maintainers {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Activated {
		n += 2
	}
	if len(m.Assets) > 0 {
		for _, s := range m.Assets {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *LinkedAddresses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DepositAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.RecipientAddress.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ChainState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maintainers", wireType)
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
			m.Maintainers = append(m.Maintainers, make([]byte, postIndex-iNdEx))
			copy(m.Maintainers[len(m.Maintainers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Activated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.Activated = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
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
			m.Assets = append(m.Assets, string(dAtA[iNdEx:postIndex]))
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
func (m *LinkedAddresses) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LinkedAddresses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkedAddresses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RecipientAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
