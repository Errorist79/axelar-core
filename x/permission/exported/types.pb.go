// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
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

type Role int32

const (
	ROLE_UNSPECIFIED      Role = 0
	ROLE_ACCESS_CONTROL   Role = 1
	ROLE_CHAIN_MANAGEMENT Role = 2
)

var Role_name = map[int32]string{
	0: "ROLE_UNSPECIFIED",
	1: "ROLE_ACCESS_CONTROL",
	2: "ROLE_CHAIN_MANAGEMENT",
}

var Role_value = map[string]int32{
	"ROLE_UNSPECIFIED":      0,
	"ROLE_ACCESS_CONTROL":   1,
	"ROLE_CHAIN_MANAGEMENT": 2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_122b0697c9fb828e, []int{0}
}

func init() {
	proto.RegisterEnum("permission.exported.v1beta1.Role", Role_name, Role_value)
}

func init() {
	proto.RegisterFile("permission/exported/v1beta1/types.proto", fileDescriptor_122b0697c9fb828e)
}

var fileDescriptor_122b0697c9fb828e = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x48, 0x2d, 0xca,
	0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x4f, 0xad, 0x28, 0xc8, 0x2f, 0x2a, 0x49, 0x4d, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x46, 0x28, 0xd4, 0x83, 0x29, 0xd4, 0x83, 0x2a, 0x94, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd3, 0x07, 0xb1, 0x20, 0x5a, 0xb4, 0xa2, 0xb8, 0x58, 0x82, 0xf2,
	0x73, 0x52, 0x85, 0x44, 0xb8, 0x04, 0x82, 0xfc, 0x7d, 0x5c, 0xe3, 0x43, 0xfd, 0x82, 0x03, 0x5c,
	0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0xc4, 0xb9, 0x84, 0xc1, 0xa2, 0x8e, 0xce,
	0xce, 0xae, 0xc1, 0xc1, 0xf1, 0xce, 0xfe, 0x7e, 0x21, 0x41, 0xfe, 0x3e, 0x02, 0x8c, 0x42, 0x92,
	0x5c, 0xa2, 0x60, 0x09, 0x67, 0x0f, 0x47, 0x4f, 0xbf, 0x78, 0x5f, 0x47, 0x3f, 0x47, 0x77, 0x57,
	0x5f, 0x57, 0xbf, 0x10, 0x01, 0x26, 0x29, 0x8e, 0x8e, 0xc5, 0x72, 0x0c, 0x2b, 0x96, 0xc8, 0x31,
	0x3a, 0x45, 0x9c, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x56, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x89, 0x15, 0xa9,
	0x39, 0x89, 0x45, 0x79, 0xa9, 0x25, 0xe5, 0xf9, 0x45, 0xd9, 0x50, 0x9e, 0x6e, 0x72, 0x7e, 0x51,
	0xaa, 0x7e, 0x85, 0x3e, 0x16, 0xcf, 0x27, 0xb1, 0x81, 0x1d, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x0f, 0xa8, 0xf9, 0x1a, 0x01, 0x00, 0x00,
}
