// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/metrics/error.proto

package metrics

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Error_Type int32

const (
	Error_UNSET       Error_Type = 0
	Error_COLLECTION  Error_Type = 1
	Error_REPLAY      Error_Type = 2
	Error_STORAGE     Error_Type = 3
	Error_DESTINATION Error_Type = 4
	Error_SCHEMA      Error_Type = 5
	Error_SEARCH      Error_Type = 6
	Error_TASK        Error_Type = 7
)

var Error_Type_name = map[int32]string{
	0: "UNSET",
	1: "COLLECTION",
	2: "REPLAY",
	3: "STORAGE",
	4: "DESTINATION",
	5: "SCHEMA",
	6: "SEARCH",
	7: "TASK",
}

var Error_Type_value = map[string]int32{
	"UNSET":       0,
	"COLLECTION":  1,
	"REPLAY":      2,
	"STORAGE":     3,
	"DESTINATION": 4,
	"SCHEMA":      5,
	"SEARCH":      6,
	"TASK":        7,
}

func (x Error_Type) String() string {
	return proto.EnumName(Error_Type_name, int32(x))
}

func (Error_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f84d81ae4efc9ce4, []int{0, 0}
}

// Used for internal error reporting using mlib library and metrics service.
type Error struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               string            `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Type                 Error_Type        `protobuf:"varint,3,opt,name=type,proto3,enum=metrics.Error_Type" json:"type,omitempty"`
	Source               string            `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Value                string            `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timestamp            int64             `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84d81ae4efc9ce4, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Error) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_UNSET
}

func (m *Error) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Error) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Error) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Error) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("metrics.Error_Type", Error_Type_name, Error_Type_value)
	proto.RegisterType((*Error)(nil), "metrics.Error")
	proto.RegisterMapType((map[string]string)(nil), "metrics.Error.MetadataEntry")
}

func init() { proto.RegisterFile("events/metrics/error.proto", fileDescriptor_f84d81ae4efc9ce4) }

var fileDescriptor_f84d81ae4efc9ce4 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x0b, 0xd3, 0x40,
	0x10, 0xc5, 0xcd, 0xff, 0x76, 0x8a, 0x75, 0x59, 0x45, 0x43, 0xe9, 0x21, 0xf4, 0x62, 0x4e, 0x09,
	0x54, 0x91, 0xa2, 0xa7, 0x18, 0x17, 0x5b, 0xec, 0x1f, 0x49, 0xe2, 0x41, 0x2f, 0xb2, 0x49, 0x96,
	0x36, 0xd8, 0x6d, 0xc2, 0x66, 0x53, 0xc8, 0x67, 0xf0, 0x4b, 0x4b, 0xd2, 0x60, 0xad, 0xb7, 0x79,
	0x6f, 0xde, 0xfc, 0x98, 0x61, 0x60, 0xc6, 0xae, 0xec, 0x22, 0x6b, 0x9f, 0x33, 0x29, 0x8a, 0xac,
	0xf6, 0x99, 0x10, 0xa5, 0xf0, 0x2a, 0x51, 0xca, 0x12, 0x5b, 0x83, 0xb9, 0xf8, 0xad, 0x81, 0x41,
	0xba, 0x06, 0x9e, 0x82, 0x5a, 0xe4, 0xb6, 0xe2, 0x28, 0xee, 0x38, 0x52, 0x8b, 0x1c, 0xbf, 0x02,
	0x4b, 0x32, 0xca, 0x7f, 0x16, 0xb9, 0xad, 0xf6, 0xa6, 0xd9, 0xc9, 0x4d, 0x8e, 0x5f, 0x83, 0x2e,
	0xdb, 0x8a, 0xd9, 0x9a, 0xa3, 0xb8, 0xd3, 0xe5, 0x73, 0x6f, 0x40, 0x79, 0x3d, 0xc6, 0x4b, 0xda,
	0x8a, 0x45, 0x7d, 0x00, 0xbf, 0x04, 0xb3, 0x2e, 0x1b, 0x91, 0x31, 0x5b, 0xbf, 0x01, 0x6e, 0x0a,
	0xbf, 0x00, 0xe3, 0x4a, 0xcf, 0x0d, 0xb3, 0x8d, 0xde, 0xbe, 0x09, 0xbc, 0x82, 0x11, 0x67, 0x92,
	0xe6, 0x54, 0x52, 0xdb, 0x74, 0x34, 0x77, 0xb2, 0x9c, 0xff, 0x87, 0xde, 0x0d, 0x6d, 0x72, 0x91,
	0xa2, 0x8d, 0xfe, 0xa6, 0xf1, 0x1c, 0xc6, 0xb2, 0xe0, 0xac, 0x96, 0x94, 0x57, 0xb6, 0xe5, 0x28,
	0xae, 0x16, 0xdd, 0x8d, 0xd9, 0x07, 0x78, 0xfa, 0x30, 0x88, 0x11, 0x68, 0xbf, 0x58, 0x3b, 0x5c,
	0xda, 0x95, 0xf7, 0x85, 0xd4, 0x7f, 0x16, 0x7a, 0xaf, 0xae, 0x94, 0x05, 0x07, 0xbd, 0x3b, 0x08,
	0x8f, 0xc1, 0xf8, 0xb6, 0x8f, 0x49, 0x82, 0x9e, 0xe0, 0x29, 0x40, 0x78, 0xd8, 0x6e, 0x49, 0x98,
	0x6c, 0x0e, 0x7b, 0xa4, 0x60, 0x00, 0x33, 0x22, 0x5f, 0xb7, 0xc1, 0x77, 0xa4, 0xe2, 0x09, 0x58,
	0x71, 0x72, 0x88, 0x82, 0xcf, 0x04, 0x69, 0xf8, 0x19, 0x4c, 0x3e, 0x91, 0x38, 0xd9, 0xec, 0x83,
	0x3e, 0xa9, 0x77, 0xc9, 0x38, 0x5c, 0x93, 0x5d, 0x80, 0x8c, 0xbe, 0x26, 0x41, 0x14, 0xae, 0x91,
	0x89, 0x47, 0xa0, 0x27, 0x41, 0xfc, 0x05, 0x59, 0x1f, 0xdf, 0xfd, 0x78, 0x7b, 0x2c, 0xe4, 0xa9,
	0x49, 0xbd, 0xac, 0xe4, 0x7e, 0x4a, 0x65, 0x76, 0xca, 0x4a, 0x51, 0xf9, 0x75, 0x76, 0x62, 0x9c,
	0xd6, 0x7e, 0xda, 0x14, 0xe7, 0xdc, 0x3f, 0x96, 0xfe, 0xe3, 0x6b, 0x53, 0xb3, 0xff, 0xea, 0x9b,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x8a, 0xb7, 0x18, 0xf3, 0x01, 0x00, 0x00,
}
