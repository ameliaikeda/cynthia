// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/monzo/terrors/proto/error.proto

package terrorsproto

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

type StackFrame struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Line                 int32    `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StackFrame) Reset()         { *m = StackFrame{} }
func (m *StackFrame) String() string { return proto.CompactTextString(m) }
func (*StackFrame) ProtoMessage()    {}
func (*StackFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae33a222c066248f, []int{0}
}

func (m *StackFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFrame.Unmarshal(m, b)
}
func (m *StackFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFrame.Marshal(b, m, deterministic)
}
func (m *StackFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFrame.Merge(m, src)
}
func (m *StackFrame) XXX_Size() int {
	return xxx_messageInfo_StackFrame.Size(m)
}
func (m *StackFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFrame.DiscardUnknown(m)
}

var xxx_messageInfo_StackFrame proto.InternalMessageInfo

func (m *StackFrame) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *StackFrame) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *StackFrame) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type Error struct {
	Code    string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Params  map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Stack   []*StackFrame     `protobuf:"bytes,4,rep,name=stack,proto3" json:"stack,omitempty"`
	// We don't use google.protobuf.BoolValue as it doesn't serialize properly without jsonpb.
	Retryable            *BoolValue `protobuf:"bytes,5,opt,name=retryable,proto3" json:"retryable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae33a222c066248f, []int{1}
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

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Error) GetStack() []*StackFrame {
	if m != nil {
		return m.Stack
	}
	return nil
}

func (m *Error) GetRetryable() *BoolValue {
	if m != nil {
		return m.Retryable
	}
	return nil
}

type BoolValue struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolValue) Reset()         { *m = BoolValue{} }
func (m *BoolValue) String() string { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()    {}
func (*BoolValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae33a222c066248f, []int{2}
}

func (m *BoolValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolValue.Unmarshal(m, b)
}
func (m *BoolValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolValue.Marshal(b, m, deterministic)
}
func (m *BoolValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolValue.Merge(m, src)
}
func (m *BoolValue) XXX_Size() int {
	return xxx_messageInfo_BoolValue.Size(m)
}
func (m *BoolValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolValue.DiscardUnknown(m)
}

var xxx_messageInfo_BoolValue proto.InternalMessageInfo

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*StackFrame)(nil), "StackFrame")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterMapType((map[string]string)(nil), "Error.ParamsEntry")
	proto.RegisterType((*BoolValue)(nil), "BoolValue")
}

func init() {
	proto.RegisterFile("github.com/monzo/terrors/proto/error.proto", fileDescriptor_ae33a222c066248f)
}

var fileDescriptor_ae33a222c066248f = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x41, 0x4f, 0x83, 0x30,
	0x18, 0x4d, 0xc7, 0xc0, 0xf1, 0x61, 0x8c, 0x69, 0x8c, 0x69, 0x76, 0x62, 0x9c, 0xc8, 0x0e, 0x25,
	0x99, 0x17, 0xf5, 0xb8, 0x64, 0x9e, 0x4d, 0x35, 0x1e, 0xbc, 0x15, 0x56, 0x37, 0x32, 0x4a, 0x97,
	0xd2, 0x99, 0xe0, 0x3f, 0xf6, 0x5f, 0x98, 0x96, 0x6e, 0x78, 0xe2, 0x3d, 0xbe, 0xf7, 0xde, 0xf7,
	0xda, 0xc2, 0x72, 0x57, 0x9b, 0xfd, 0xa9, 0xa4, 0x95, 0x92, 0x85, 0x54, 0xed, 0x8f, 0x2a, 0x8c,
	0xd0, 0x5a, 0xe9, 0xae, 0x38, 0x6a, 0x65, 0x54, 0xe1, 0x08, 0x75, 0x38, 0x7b, 0x07, 0x78, 0x33,
	0xbc, 0x3a, 0xbc, 0x68, 0x2e, 0x05, 0x9e, 0xc3, 0xec, 0xab, 0x6e, 0x44, 0xcb, 0xa5, 0x20, 0x28,
	0x45, 0x79, 0xcc, 0x2e, 0x1c, 0x63, 0x98, 0x36, 0x75, 0x2b, 0xc8, 0x24, 0x45, 0x79, 0xc8, 0x1c,
	0xc6, 0xf7, 0x10, 0x49, 0x61, 0xf6, 0x6a, 0x4b, 0x02, 0xa7, 0xf6, 0x2c, 0xfb, 0x45, 0x10, 0x6e,
	0xec, 0x16, 0xeb, 0xaa, 0xd4, 0xf6, 0x9c, 0xe6, 0x30, 0x26, 0x70, 0x25, 0x45, 0xd7, 0xf1, 0xdd,
	0x10, 0x16, 0xb3, 0x33, 0xc5, 0x4b, 0x88, 0x8e, 0x5c, 0x73, 0xd9, 0x91, 0x20, 0x0d, 0xf2, 0x64,
	0x85, 0xa9, 0x4b, 0xa1, 0xaf, 0xee, 0xe7, 0xa6, 0x35, 0xba, 0x67, 0x5e, 0x81, 0x17, 0x10, 0x76,
	0xb6, 0x39, 0x99, 0x3a, 0x69, 0x42, 0xc7, 0x73, 0xb0, 0x61, 0x82, 0x73, 0x88, 0xb5, 0x30, 0xba,
	0xe7, 0x65, 0x23, 0x48, 0x98, 0xa2, 0x3c, 0x59, 0x01, 0x5d, 0x2b, 0xd5, 0x7c, 0xf0, 0xe6, 0x24,
	0xd8, 0x38, 0x9c, 0x3f, 0x41, 0xf2, 0x6f, 0x07, 0xbe, 0x85, 0xe0, 0x20, 0x7a, 0x5f, 0xda, 0x42,
	0x7c, 0x07, 0xe1, 0xb7, 0x35, 0xf9, 0xc6, 0x03, 0x79, 0x9e, 0x3c, 0xa2, 0x6c, 0x01, 0xf1, 0x25,
	0x72, 0x94, 0x59, 0xeb, 0xcc, 0xcb, 0xd6, 0x37, 0x9f, 0xd7, 0xfe, 0x05, 0xdc, 0xa5, 0x97, 0x91,
	0xfb, 0x3c, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x38, 0xbd, 0x96, 0xf5, 0xa9, 0x01, 0x00, 0x00,
}