// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event-service/pkg/pb/service.proto

package eventService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	EventId              string               `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Checksum             string               `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	EntityType           string               `protobuf:"bytes,4,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	EntityId             string               `protobuf:"bytes,5,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Data                 []byte               `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_64242f18e16eef93, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Event) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Event) GetEntityType() string {
	if m != nil {
		return m.EntityType
	}
	return ""
}

func (m *Event) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TestMessage struct {
	Message              string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_64242f18e16eef93, []int{1}
}

func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (m *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(m, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TestMessage) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*TestMessage)(nil), "TestMessage")
}

func init() { proto.RegisterFile("event-service/pkg/pb/service.proto", fileDescriptor_64242f18e16eef93) }

var fileDescriptor_64242f18e16eef93 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4e, 0xf4, 0x30,
	0x10, 0xc4, 0x95, 0xfb, 0xee, 0x5f, 0x36, 0xa7, 0xaf, 0x70, 0x65, 0x42, 0x71, 0x51, 0xaa, 0x34,
	0xc4, 0x12, 0x54, 0x94, 0x20, 0x51, 0x5c, 0x41, 0x13, 0x52, 0xd1, 0x44, 0x4e, 0xbc, 0x04, 0xeb,
	0xc8, 0xc5, 0x8a, 0xf7, 0x4e, 0xca, 0x33, 0xf2, 0x52, 0x08, 0x3b, 0xa1, 0xa0, 0xa3, 0xdb, 0xdf,
	0x7a, 0x34, 0xb3, 0x1e, 0x48, 0xf1, 0x82, 0x27, 0xba, 0xb1, 0x38, 0x5c, 0x74, 0x83, 0xc2, 0x1c,
	0x5b, 0x61, 0x6a, 0x31, 0x61, 0x6e, 0x86, 0x9e, 0xfa, 0x78, 0xdf, 0xf6, 0x7d, 0xfb, 0x81, 0xc2,
	0x51, 0x7d, 0x7e, 0x13, 0xa4, 0x3b, 0xb4, 0x24, 0x3b, 0xe3, 0x05, 0xe9, 0x67, 0x00, 0xab, 0xa7,
	0x6f, 0x1f, 0x76, 0x05, 0x5b, 0x67, 0x58, 0x69, 0xc5, 0x83, 0x24, 0xc8, 0xc2, 0x62, 0xe3, 0xf8,
	0xa0, 0x58, 0x0c, 0xdb, 0xe6, 0x1d, 0x9b, 0xa3, 0x3d, 0x77, 0x7c, 0xe1, 0x9e, 0x7e, 0x98, 0xdd,
	0x03, 0x34, 0x03, 0x4a, 0x42, 0x55, 0x49, 0xe2, 0xff, 0x92, 0x20, 0x8b, 0x6e, 0xe3, 0xdc, 0xc7,
	0xe6, 0x73, 0x6c, 0x5e, 0xce, 0xb1, 0x45, 0x38, 0xa9, 0x1f, 0x88, 0xed, 0x21, 0xc2, 0x13, 0x69,
	0x1a, 0x2b, 0x1a, 0x0d, 0xf2, 0xa5, 0x73, 0x06, 0xbf, 0x2a, 0x47, 0x83, 0xec, 0x1a, 0xc2, 0x49,
	0xa0, 0x15, 0x5f, 0xf9, 0x60, 0xbf, 0x38, 0x28, 0xc6, 0x60, 0xa9, 0x24, 0x49, 0xbe, 0x4e, 0x82,
	0x6c, 0x57, 0xb8, 0x39, 0xad, 0x21, 0x2a, 0xd1, 0xd2, 0x33, 0x5a, 0x2b, 0x5b, 0x64, 0x1c, 0x36,
	0x9d, 0x1f, 0xe7, 0x1f, 0x4d, 0xf8, 0xeb, 0xea, 0xc5, 0x1f, 0xae, 0x7e, 0xfc, 0xff, 0xba, 0x73,
	0xbd, 0xbc, 0xf8, 0xa2, 0xeb, 0xb5, 0x93, 0xdf, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xde, 0x3d,
	0x13, 0x3d, 0x8f, 0x01, 0x00, 0x00,
}
