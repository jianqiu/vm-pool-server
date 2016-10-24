// Code generated by protoc-gen-go.
// source: db/models/error.proto
// DO NOT EDIT!

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	db/models/error.proto

It has these top-level messages:
	Error
*/
package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error_Type int32

const (
	Error_UnknownError           Error_Type = 0
	Error_InvalidDomain          Error_Type = 1
	Error_UnkownVersion          Error_Type = 2
	Error_InvalidRecord          Error_Type = 3
	Error_InvalidRequest         Error_Type = 4
	Error_InvalidResponse        Error_Type = 5
	Error_InvalidProtobufMessage Error_Type = 6
	Error_InvalidJSON            Error_Type = 7
	Error_FailedToOpenEnvelope   Error_Type = 8
	Error_InvalidStateTransition Error_Type = 9
	Error_Unauthorized           Error_Type = 10
	Error_ResourceConflict       Error_Type = 11
	Error_ResourceExists         Error_Type = 12
	Error_ResourceNotFound       Error_Type = 13
	Error_RouterError            Error_Type = 14
	Error_SoftLayerAPIError      Error_Type = 15
	Error_GUIDGeneration         Error_Type = 26
	Error_Deserialize            Error_Type = 27
	Error_Deadlock               Error_Type = 28
	Error_Unrecoverable          Error_Type = 29
)

var Error_Type_name = map[int32]string{
	0:  "UnknownError",
	1:  "InvalidDomain",
	2:  "UnkownVersion",
	3:  "InvalidRecord",
	4:  "InvalidRequest",
	5:  "InvalidResponse",
	6:  "InvalidProtobufMessage",
	7:  "InvalidJSON",
	8:  "FailedToOpenEnvelope",
	9:  "InvalidStateTransition",
	10: "Unauthorized",
	11: "ResourceConflict",
	12: "ResourceExists",
	13: "ResourceNotFound",
	14: "RouterError",
	15: "SoftLayerAPIError",
	26: "GUIDGeneration",
	27: "Deserialize",
	28: "Deadlock",
	29: "Unrecoverable",
}
var Error_Type_value = map[string]int32{
	"UnknownError":           0,
	"InvalidDomain":          1,
	"UnkownVersion":          2,
	"InvalidRecord":          3,
	"InvalidRequest":         4,
	"InvalidResponse":        5,
	"InvalidProtobufMessage": 6,
	"InvalidJSON":            7,
	"FailedToOpenEnvelope":   8,
	"InvalidStateTransition": 9,
	"Unauthorized":           10,
	"ResourceConflict":       11,
	"ResourceExists":         12,
	"ResourceNotFound":       13,
	"RouterError":            14,
	"SoftLayerAPIError":      15,
	"GUIDGeneration":         26,
	"Deserialize":            27,
	"Deadlock":               28,
	"Unrecoverable":          29,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}
func (Error_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Error struct {
	Type    Error_Type `protobuf:"varint,1,opt,name=type,enum=models.Error_Type" json:"type,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Error)(nil), "models.Error")
	proto.RegisterEnum("models.Error_Type", Error_Type_name, Error_Type_value)
}

func init() { proto.RegisterFile("db/models/error.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0x51, 0x4e, 0x1b, 0x31,
	0x10, 0x86, 0x1b, 0x12, 0x02, 0x0c, 0x21, 0x19, 0x5c, 0xa8, 0x56, 0xb4, 0x95, 0x10, 0x0f, 0x15,
	0x4f, 0x8b, 0xd4, 0x9e, 0xa0, 0xea, 0x06, 0x94, 0xaa, 0x05, 0xb4, 0x49, 0xfa, 0xee, 0x5d, 0x4f,
	0x5a, 0x2b, 0x8e, 0x67, 0x6b, 0x7b, 0x43, 0xc3, 0x5d, 0x7a, 0xc0, 0xde, 0xa2, 0xf2, 0x2e, 0xa0,
	0x3c, 0xfa, 0x9b, 0x7f, 0x46, 0xdf, 0x8c, 0xe1, 0x54, 0x15, 0x57, 0x2b, 0x56, 0x64, 0xfc, 0x15,
	0x39, 0xc7, 0x2e, 0xad, 0x1c, 0x07, 0x16, 0xfd, 0x96, 0x5d, 0xfc, 0xeb, 0xc2, 0xee, 0x38, 0x72,
	0xf1, 0x01, 0x7a, 0x61, 0x53, 0x51, 0xd2, 0x39, 0xef, 0x5c, 0x0e, 0x3f, 0x8a, 0xb4, 0x0d, 0xa4,
	0x4d, 0x31, 0x9d, 0x6d, 0x2a, 0xca, 0x9b, 0xba, 0x48, 0x60, 0x6f, 0x45, 0xde, 0xcb, 0x9f, 0x94,
	0xec, 0x9c, 0x77, 0x2e, 0x0f, 0xf2, 0xe7, 0xe7, 0xc5, 0xdf, 0x2e, 0xf4, 0x62, 0x50, 0x20, 0x0c,
	0xe6, 0x76, 0x69, 0xf9, 0xc1, 0x36, 0xdd, 0xf8, 0x4a, 0x1c, 0xc3, 0xd1, 0xc4, 0xae, 0xa5, 0xd1,
	0x2a, 0xe3, 0x95, 0xd4, 0x16, 0x3b, 0x11, 0xcd, 0xed, 0x92, 0x1f, 0xec, 0x0f, 0x72, 0x5e, 0xb3,
	0xc5, 0x9d, 0xad, 0x54, 0x4e, 0x25, 0x3b, 0x85, 0x5d, 0x21, 0x60, 0xf8, 0x82, 0x7e, 0xd7, 0xe4,
	0x03, 0xf6, 0xc4, 0x6b, 0x18, 0xbd, 0x30, 0x5f, 0xb1, 0xf5, 0x84, 0xbb, 0xe2, 0x0c, 0xde, 0x3c,
	0xc1, 0xfb, 0xb8, 0x60, 0x51, 0x2f, 0xbe, 0xb7, 0x5a, 0xd8, 0x17, 0x23, 0x38, 0x7c, 0xaa, 0x7d,
	0x9d, 0xde, 0xdd, 0xe2, 0x9e, 0x48, 0xe0, 0xe4, 0x5a, 0x6a, 0x43, 0x6a, 0xc6, 0x77, 0x15, 0xd9,
	0xb1, 0x5d, 0x93, 0xe1, 0x8a, 0x70, 0x7f, 0x6b, 0xcc, 0x34, 0xc8, 0x40, 0x33, 0x27, 0xad, 0xd7,
	0x21, 0xea, 0x1d, 0xb4, 0x6b, 0xc9, 0x3a, 0xfc, 0x62, 0xa7, 0x1f, 0x49, 0x21, 0x88, 0x13, 0xc0,
	0x9c, 0x3c, 0xd7, 0xae, 0xa4, 0x2f, 0x6c, 0x17, 0x46, 0x97, 0x01, 0x0f, 0xa3, 0xf3, 0x33, 0x1d,
	0xff, 0xd1, 0x3e, 0x78, 0x1c, 0x6c, 0x27, 0x6f, 0x39, 0x5c, 0x73, 0x6d, 0x15, 0x1e, 0x45, 0xb1,
	0x9c, 0xeb, 0x40, 0xae, 0xbd, 0xd3, 0x50, 0x9c, 0xc2, 0xf1, 0x94, 0x17, 0xe1, 0x9b, 0xdc, 0x90,
	0xfb, 0x7c, 0x3f, 0x69, 0xf1, 0x28, 0x4e, 0xbc, 0x99, 0x4f, 0xb2, 0x1b, 0xb2, 0xe4, 0x64, 0x63,
	0x73, 0x16, 0x7b, 0x33, 0xf2, 0xe4, 0xb4, 0x34, 0xfa, 0x91, 0xf0, 0xad, 0x18, 0xc0, 0x7e, 0x46,
	0x52, 0x19, 0x2e, 0x97, 0xf8, 0xae, 0x3d, 0xaf, 0xa3, 0x92, 0xd7, 0xe4, 0x64, 0x61, 0x08, 0xdf,
	0x17, 0xfd, 0xe6, 0xeb, 0x3f, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x14, 0xce, 0xf7, 0x13,
	0x02, 0x00, 0x00,
}