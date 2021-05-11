// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/params.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
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

// Params is the parameter set for this module
type Params struct {
	// Deprecated
	LockingPeriod int64 `protobuf:"varint,1,opt,name=locking_period,json=lockingPeriod,proto3" json:"locking_period,omitempty"`
	// MinKeygenThreshold defines the minimum % of stake that must be online
	// to authorize generation of a new key in the system.
	MinKeygenThreshold utils.Threshold `protobuf:"bytes,2,opt,name=min_keygen_threshold,json=minKeygenThreshold,proto3" json:"min_keygen_threshold"`
	// CorruptionThreshold defines the corruption threshold with which
	// we'll run keygen protocol.
	CorruptionThreshold utils.Threshold `protobuf:"bytes,3,opt,name=corruption_threshold,json=corruptionThreshold,proto3" json:"corruption_threshold"`
	// KeyRequirements defines the requirement of each key for each chain
	KeyRequirements []exported.KeyRequirement `protobuf:"bytes,4,rep,name=key_requirements,json=keyRequirements,proto3" json:"key_requirements"`
	// MinBondFractionPerShare defines the % of stake validators have to bond per
	// key share
	MinBondFractionPerShare utils.Threshold `protobuf:"bytes,5,opt,name=min_bond_fraction_per_share,json=minBondFractionPerShare,proto3" json:"min_bond_fraction_per_share"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c9a42e8b26dfec, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "tss.v1beta1.Params")
}

func init() { proto.RegisterFile("tss/v1beta1/params.proto", fileDescriptor_67c9a42e8b26dfec) }

var fileDescriptor_67c9a42e8b26dfec = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0xa1, 0xb8, 0x3e, 0x60, 0xf5, 0x8f, 0xa8, 0xa5, 0x22, 0x57, 0xa5, 0xa8, 0x6a, 0x25,
	0x5f, 0x0a, 0xb5, 0xfb, 0x06, 0x3e, 0xf4, 0x62, 0x29, 0x22, 0x4e, 0x72, 0x89, 0x22, 0x21, 0xfe,
	0x4c, 0x60, 0x05, 0xec, 0x92, 0xd9, 0x25, 0x31, 0x8f, 0x90, 0x5b, 0x1e, 0xcb, 0x47, 0x1f, 0x73,
	0x8a, 0x12, 0xfb, 0x45, 0x22, 0x36, 0xd8, 0x24, 0x37, 0xdf, 0x60, 0xbe, 0xf9, 0x7e, 0xf3, 0xed,
	0x8c, 0x6e, 0x0a, 0xce, 0xdd, 0xeb, 0x49, 0x08, 0x22, 0x98, 0xb8, 0x65, 0x80, 0x41, 0xc1, 0x9d,
	0x12, 0x99, 0x60, 0xc6, 0x40, 0x70, 0xee, 0xb4, 0xca, 0x68, 0x98, 0xb0, 0x84, 0xc9, 0xba, 0xdb,
	0x7c, 0xbd, 0xb4, 0x8c, 0xbe, 0x57, 0x82, 0xe4, 0x9d, 0x5d, 0xa4, 0x08, 0x3c, 0x65, 0x79, 0xdc,
	0xca, 0x76, 0xc3, 0x86, 0x65, 0xc9, 0x50, 0x40, 0xdc, 0x75, 0xd5, 0x25, 0xb4, 0x33, 0x7e, 0xde,
	0x6a, 0x7a, 0xdf, 0x93, 0x43, 0x8d, 0xdf, 0xfa, 0xc7, 0x9c, 0x45, 0x19, 0xa1, 0x89, 0x5f, 0x02,
	0x12, 0x16, 0x9b, 0xaa, 0xad, 0x8e, 0xb5, 0xc5, 0x87, 0xb6, 0xea, 0xc9, 0xa2, 0xe1, 0xe9, 0xc3,
	0x82, 0x50, 0x3f, 0x83, 0x3a, 0x01, 0xea, 0xef, 0x27, 0x9a, 0xef, 0x6c, 0x75, 0x3c, 0x98, 0x9a,
	0x8e, 0x4c, 0xb4, 0x8b, 0xed, 0x9c, 0xee, 0xf4, 0x59, 0x6f, 0xf5, 0xf0, 0x43, 0x59, 0x18, 0x05,
	0xa1, 0x73, 0x69, 0xdd, 0x2b, 0xc6, 0xb1, 0x3e, 0x8c, 0x18, 0x62, 0x55, 0x0a, 0xc2, 0x5e, 0x13,
	0xb5, 0x83, 0x88, 0x5f, 0x3a, 0x6f, 0x87, 0x3c, 0xd3, 0x3f, 0x67, 0x50, 0xfb, 0x08, 0x57, 0x15,
	0x41, 0x28, 0x80, 0x0a, 0x6e, 0xf6, 0x6c, 0x6d, 0x3c, 0x98, 0xfe, 0x72, 0x9a, 0xad, 0xee, 0x76,
	0xb2, 0xa7, 0xce, 0xa1, 0x5e, 0x74, 0xcd, 0x2d, 0xfa, 0x53, 0xf6, 0xa6, 0xca, 0x8d, 0x0b, 0xfd,
	0x5b, 0xf3, 0xf6, 0x90, 0xd1, 0xd8, 0xbf, 0xc4, 0x20, 0x92, 0x81, 0x4b, 0x40, 0x9f, 0xa7, 0x01,
	0x82, 0xf9, 0xfe, 0xa0, 0xc0, 0x5f, 0x0b, 0x42, 0x67, 0x8c, 0xc6, 0xff, 0x5b, 0x80, 0x07, 0x78,
	0xd2, 0xd8, 0x67, 0x47, 0xab, 0x27, 0x4b, 0x59, 0x6d, 0x2c, 0x75, 0xbd, 0xb1, 0xd4, 0xc7, 0x8d,
	0xa5, 0xde, 0x6d, 0x2d, 0x65, 0xbd, 0xb5, 0x94, 0xfb, 0xad, 0xa5, 0x9c, 0xff, 0x4d, 0x88, 0x48,
	0xab, 0xd0, 0x89, 0x58, 0xe1, 0x06, 0x4b, 0xc8, 0x03, 0xa4, 0x20, 0x6e, 0x18, 0x66, 0xed, 0xdf,
	0x9f, 0x88, 0x21, 0xb8, 0x4b, 0xb7, 0x39, 0xb9, 0xbc, 0x70, 0xd8, 0x97, 0x27, 0xfe, 0xf7, 0x1c,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0x33, 0xf0, 0x02, 0x62, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinBondFractionPerShare.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.KeyRequirements) > 0 {
		for iNdEx := len(m.KeyRequirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyRequirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.CorruptionThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.MinKeygenThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LockingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LockingPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LockingPeriod != 0 {
		n += 1 + sovParams(uint64(m.LockingPeriod))
	}
	l = m.MinKeygenThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CorruptionThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.KeyRequirements) > 0 {
		for _, e := range m.KeyRequirements {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.MinBondFractionPerShare.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockingPeriod", wireType)
			}
			m.LockingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinKeygenThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinKeygenThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorruptionThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CorruptionThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyRequirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyRequirements = append(m.KeyRequirements, exported.KeyRequirement{})
			if err := m.KeyRequirements[len(m.KeyRequirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBondFractionPerShare", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBondFractionPerShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)