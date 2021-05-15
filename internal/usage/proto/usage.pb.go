// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/micro/v2/internal/usage/proto/usage.proto

package usage

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

type Usage struct {
	// service name
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// version of service
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// unique service id
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// unix timestamp of report
	Timestamp uint64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// window of report in seconds
	Window uint64 `protobuf:"varint,5,opt,name=window,proto3" json:"window,omitempty"`
	// usage metrics
	Metrics              *Metrics `protobuf:"bytes,6,opt,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Usage) Reset()         { *m = Usage{} }
func (m *Usage) String() string { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()    {}
func (*Usage) Descriptor() ([]byte, []int) {
	return fileDescriptor_usage_7cdccb5f3bba2052, []int{0}
}
func (m *Usage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Usage.Unmarshal(m, b)
}
func (m *Usage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Usage.Marshal(b, m, deterministic)
}
func (dst *Usage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Usage.Merge(dst, src)
}
func (m *Usage) XXX_Size() int {
	return xxx_messageInfo_Usage.Size(m)
}
func (m *Usage) XXX_DiscardUnknown() {
	xxx_messageInfo_Usage.DiscardUnknown(m)
}

var xxx_messageInfo_Usage proto.InternalMessageInfo

func (m *Usage) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Usage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Usage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Usage) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Usage) GetWindow() uint64 {
	if m != nil {
		return m.Window
	}
	return 0
}

func (m *Usage) GetMetrics() *Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Metrics struct {
	// counts such as requests, services, etc
	Count                map[string]uint64 `protobuf:"bytes,1,rep,name=count,proto3" json:"count,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_usage_7cdccb5f3bba2052, []int{1}
}
func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (dst *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(dst, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetCount() map[string]uint64 {
	if m != nil {
		return m.Count
	}
	return nil
}

func init() {
	proto.RegisterType((*Usage)(nil), "Usage")
	proto.RegisterType((*Metrics)(nil), "Metrics")
	proto.RegisterMapType((map[string]uint64)(nil), "Metrics.CountEntry")
}

func init() {
	proto.RegisterFile("github.com/micro/micro/v2/internal/usage/proto/usage.proto", fileDescriptor_usage_7cdccb5f3bba2052)
}

var fileDescriptor_usage_7cdccb5f3bba2052 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbd, 0x6a, 0xc3, 0x30,
	0x14, 0x85, 0x91, 0xe3, 0x9f, 0xe6, 0x06, 0x4a, 0x51, 0x4b, 0x11, 0xa5, 0x83, 0xf1, 0xe4, 0x2e,
	0x36, 0xa4, 0x43, 0x43, 0xd7, 0xd2, 0xb1, 0x8b, 0xa1, 0x0f, 0xe0, 0xc8, 0x22, 0xbd, 0x34, 0x92,
	0x82, 0x24, 0x3b, 0xe4, 0x81, 0xfa, 0x9e, 0xc5, 0x92, 0x4c, 0x16, 0x71, 0xbe, 0xef, 0xdc, 0xe5,
	0x08, 0xde, 0x0e, 0xe8, 0x7e, 0xc6, 0x7d, 0xc3, 0xb5, 0x6c, 0x25, 0x72, 0xa3, 0xe3, 0x8b, 0xca,
	0x09, 0xa3, 0xfa, 0x63, 0x3b, 0xda, 0xfe, 0x20, 0xda, 0x93, 0xd1, 0x4e, 0x87, 0xdc, 0xf8, 0x5c,
	0xfd, 0x11, 0xc8, 0xbe, 0x67, 0xa6, 0x0c, 0x0a, 0x2b, 0xcc, 0x84, 0x5c, 0x30, 0x52, 0x92, 0x7a,
	0xdd, 0x2d, 0x38, 0x37, 0x93, 0x30, 0x16, 0xb5, 0x62, 0x49, 0x68, 0x22, 0xd2, 0x5b, 0x48, 0x70,
	0x60, 0x2b, 0x2f, 0x13, 0x1c, 0xe8, 0x33, 0xac, 0x1d, 0x4a, 0x61, 0x5d, 0x2f, 0x4f, 0x2c, 0x2d,
	0x49, 0x9d, 0x76, 0x57, 0x41, 0x1f, 0x21, 0x3f, 0xa3, 0x1a, 0xf4, 0x99, 0x65, 0xbe, 0x8a, 0x44,
	0x2b, 0x28, 0xa4, 0x70, 0x06, 0xb9, 0x65, 0x79, 0x49, 0xea, 0xcd, 0xf6, 0xa6, 0xf9, 0x0a, 0xdc,
	0x2d, 0x45, 0xa5, 0xa0, 0x88, 0x8e, 0xbe, 0x40, 0xc6, 0xf5, 0xa8, 0x1c, 0x23, 0xe5, 0xaa, 0xde,
	0x6c, 0xef, 0x97, 0xe3, 0xe6, 0x63, 0xb6, 0x9f, 0xca, 0x99, 0x4b, 0x17, 0x2e, 0x9e, 0x76, 0x00,
	0x57, 0x49, 0xef, 0x60, 0xf5, 0x2b, 0x2e, 0x71, 0xdd, 0x1c, 0xe9, 0x03, 0x64, 0x53, 0x7f, 0x1c,
	0x85, 0xdf, 0x95, 0x76, 0x01, 0xde, 0x93, 0x1d, 0xd9, 0xe7, 0xfe, 0x7b, 0x5e, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0xfc, 0xa1, 0x6c, 0x59, 0x01, 0x00, 0x00,
}
