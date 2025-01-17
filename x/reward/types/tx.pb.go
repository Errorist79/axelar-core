// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reward/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type RefundMsgRequest struct {
	Sender       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	InnerMessage *types.Any                                    `protobuf:"bytes,2,opt,name=inner_message,json=innerMessage,proto3" json:"inner_message,omitempty"`
}

func (m *RefundMsgRequest) Reset()         { *m = RefundMsgRequest{} }
func (m *RefundMsgRequest) String() string { return proto.CompactTextString(m) }
func (*RefundMsgRequest) ProtoMessage()    {}
func (*RefundMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a40a70d6e1fd0ef, []int{0}
}
func (m *RefundMsgRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefundMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefundMsgRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefundMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundMsgRequest.Merge(m, src)
}
func (m *RefundMsgRequest) XXX_Size() int {
	return m.Size()
}
func (m *RefundMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefundMsgRequest proto.InternalMessageInfo

type RefundMsgResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Log  string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *RefundMsgResponse) Reset()         { *m = RefundMsgResponse{} }
func (m *RefundMsgResponse) String() string { return proto.CompactTextString(m) }
func (*RefundMsgResponse) ProtoMessage()    {}
func (*RefundMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a40a70d6e1fd0ef, []int{1}
}
func (m *RefundMsgResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefundMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefundMsgResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefundMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundMsgResponse.Merge(m, src)
}
func (m *RefundMsgResponse) XXX_Size() int {
	return m.Size()
}
func (m *RefundMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefundMsgResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RefundMsgRequest)(nil), "reward.v1beta1.RefundMsgRequest")
	proto.RegisterType((*RefundMsgResponse)(nil), "reward.v1beta1.RefundMsgResponse")
}

func init() { proto.RegisterFile("reward/v1beta1/tx.proto", fileDescriptor_3a40a70d6e1fd0ef) }

var fileDescriptor_3a40a70d6e1fd0ef = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x31, 0x6e, 0x22, 0x31,
	0x14, 0x86, 0x67, 0x76, 0x57, 0x48, 0xeb, 0x65, 0x11, 0x19, 0x21, 0x05, 0x28, 0x26, 0x88, 0x8a,
	0x66, 0x6c, 0x11, 0xaa, 0x94, 0xd0, 0x51, 0x50, 0x64, 0xca, 0x34, 0xc8, 0x33, 0x7e, 0x38, 0x88,
	0xc1, 0x26, 0xb6, 0x27, 0xc0, 0x2d, 0x72, 0x8c, 0x1c, 0x20, 0x87, 0x40, 0xa9, 0x28, 0x53, 0x45,
	0x09, 0xdc, 0x22, 0x55, 0x84, 0xed, 0x48, 0x54, 0xfe, 0x9f, 0x3f, 0xbf, 0xe7, 0xff, 0xd7, 0x43,
	0x97, 0x0a, 0xd6, 0x54, 0x31, 0xf2, 0xd8, 0xcf, 0xc0, 0xd0, 0x3e, 0x31, 0x1b, 0xbc, 0x52, 0xd2,
	0xc8, 0xa8, 0xe6, 0x00, 0xf6, 0xa0, 0xdd, 0xe2, 0x52, 0xf2, 0x02, 0x88, 0xa5, 0x59, 0x39, 0x23,
	0x54, 0x6c, 0xdd, 0xd3, 0x76, 0x83, 0x4b, 0x2e, 0xad, 0x24, 0x27, 0xe5, 0x6f, 0x5b, 0xb9, 0xd4,
	0x4b, 0xa9, 0xa7, 0x0e, 0xb8, 0xc2, 0xa1, 0xee, 0x73, 0x88, 0xea, 0x29, 0xcc, 0x4a, 0xc1, 0x26,
	0x9a, 0xa7, 0xf0, 0x50, 0x82, 0x36, 0xd1, 0x18, 0x55, 0x34, 0x08, 0x06, 0xaa, 0x19, 0x76, 0xc2,
	0x5e, 0x75, 0xd4, 0xff, 0x7a, 0xbf, 0x4a, 0xf8, 0xdc, 0xdc, 0x97, 0x19, 0xce, 0xe5, 0xd2, 0x4f,
	0xf0, 0x47, 0xa2, 0xd9, 0x82, 0x98, 0xed, 0x0a, 0x34, 0x1e, 0xe6, 0xf9, 0x90, 0x31, 0x05, 0x5a,
	0xa7, 0x7e, 0x40, 0x34, 0x46, 0xff, 0xe7, 0x42, 0x80, 0x9a, 0x2e, 0x41, 0x6b, 0xca, 0xa1, 0xf9,
	0xab, 0x13, 0xf6, 0xfe, 0x5d, 0x37, 0xb0, 0xcb, 0x80, 0x7f, 0x32, 0xe0, 0xa1, 0xd8, 0x8e, 0x6a,
	0xaf, 0x2f, 0x09, 0x72, 0x6e, 0x68, 0x56, 0x40, 0x5a, 0xb5, 0xad, 0x13, 0xd7, 0xd9, 0xbd, 0x41,
	0x17, 0x67, 0x4e, 0xf5, 0x4a, 0x0a, 0x0d, 0x51, 0x84, 0xfe, 0x30, 0x6a, 0xa8, 0x33, 0x9a, 0x5a,
	0x1d, 0xd5, 0xd1, 0xef, 0x42, 0x72, 0xfb, 0xd3, 0xdf, 0xf4, 0x24, 0x47, 0xb7, 0xbb, 0xcf, 0x38,
	0xd8, 0x1d, 0xe2, 0x70, 0x7f, 0x88, 0xc3, 0x8f, 0x43, 0x1c, 0x3e, 0x1d, 0xe3, 0x60, 0x7f, 0x8c,
	0x83, 0xb7, 0x63, 0x1c, 0xdc, 0x0d, 0xce, 0xa2, 0xd1, 0x0d, 0x14, 0x54, 0x09, 0x30, 0x6b, 0xa9,
	0x16, 0xbe, 0x4a, 0x72, 0xa9, 0x80, 0x6c, 0x88, 0xdf, 0x8f, 0xcd, 0x9a, 0x55, 0xac, 0xf3, 0xc1,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x3e, 0xd0, 0x99, 0xb6, 0x01, 0x00, 0x00,
}

func (m *RefundMsgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefundMsgRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefundMsgRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InnerMessage != nil {
		{
			size, err := m.InnerMessage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RefundMsgResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefundMsgResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefundMsgResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RefundMsgRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InnerMessage != nil {
		l = m.InnerMessage.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *RefundMsgResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RefundMsgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RefundMsgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefundMsgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InnerMessage == nil {
				m.InnerMessage = &types.Any{}
			}
			if err := m.InnerMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *RefundMsgResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RefundMsgResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefundMsgResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
