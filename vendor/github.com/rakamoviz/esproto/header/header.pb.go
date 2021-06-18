// Code generated by protoc-gen-go. DO NOT EDIT.
// source: header/header.proto

package header

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Command struct {
	Domain               string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Id                   string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AggregateId          string                 `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateName        string                 `protobuf:"bytes,5,opt,name=aggregate_name,json=aggregateName,proto3" json:"aggregate_name,omitempty"`
	AggregateRevision    int32                  `protobuf:"varint,6,opt,name=aggregate_revision,json=aggregateRevision,proto3" json:"aggregate_revision,omitempty"`
	AggregateExisting    bool                   `protobuf:"varint,7,opt,name=aggregate_existing,json=aggregateExisting,proto3" json:"aggregate_existing,omitempty"`
	SagaId               string                 `protobuf:"bytes,8,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	SagaName             string                 `protobuf:"bytes,9,opt,name=saga_name,json=sagaName,proto3" json:"saga_name,omitempty"`
	SagaRevision         int32                  `protobuf:"varint,10,opt,name=saga_revision,json=sagaRevision,proto3" json:"saga_revision,omitempty"`
	CorrelationId        string                 `protobuf:"bytes,11,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	CorrelationType      string                 `protobuf:"bytes,12,opt,name=correlation_type,json=correlationType,proto3" json:"correlation_type,omitempty"`
	CorrelationName      string                 `protobuf:"bytes,13,opt,name=correlation_name,json=correlationName,proto3" json:"correlation_name,omitempty"`
	Version              int32                  `protobuf:"varint,14,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SentAt               *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	ReceivedAt           *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2ea67d34c1bc80, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Command) GetAggregateName() string {
	if m != nil {
		return m.AggregateName
	}
	return ""
}

func (m *Command) GetAggregateRevision() int32 {
	if m != nil {
		return m.AggregateRevision
	}
	return 0
}

func (m *Command) GetAggregateExisting() bool {
	if m != nil {
		return m.AggregateExisting
	}
	return false
}

func (m *Command) GetSagaId() string {
	if m != nil {
		return m.SagaId
	}
	return ""
}

func (m *Command) GetSagaName() string {
	if m != nil {
		return m.SagaName
	}
	return ""
}

func (m *Command) GetSagaRevision() int32 {
	if m != nil {
		return m.SagaRevision
	}
	return 0
}

func (m *Command) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Command) GetCorrelationType() string {
	if m != nil {
		return m.CorrelationType
	}
	return ""
}

func (m *Command) GetCorrelationName() string {
	if m != nil {
		return m.CorrelationName
	}
	return ""
}

func (m *Command) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Command) GetCreatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Command) GetSentAt() *timestamppb.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

func (m *Command) GetReceivedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

type Event struct {
	Domain               string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Id                   string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AggregateId          string                 `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateName        string                 `protobuf:"bytes,5,opt,name=aggregate_name,json=aggregateName,proto3" json:"aggregate_name,omitempty"`
	AggregateRevision    int32                  `protobuf:"varint,6,opt,name=aggregate_revision,json=aggregateRevision,proto3" json:"aggregate_revision,omitempty"`
	SagaId               string                 `protobuf:"bytes,7,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	SagaName             string                 `protobuf:"bytes,8,opt,name=saga_name,json=sagaName,proto3" json:"saga_name,omitempty"`
	SagaRevision         int32                  `protobuf:"varint,9,opt,name=saga_revision,json=sagaRevision,proto3" json:"saga_revision,omitempty"`
	CorrelationId        string                 `protobuf:"bytes,10,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	CorrelationType      string                 `protobuf:"bytes,11,opt,name=correlation_type,json=correlationType,proto3" json:"correlation_type,omitempty"`
	CorrelationName      string                 `protobuf:"bytes,12,opt,name=correlation_name,json=correlationName,proto3" json:"correlation_name,omitempty"`
	Version              int32                  `protobuf:"varint,13,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SentAt               *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	ReceivedAt           *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2ea67d34c1bc80, []int{1}
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

func (m *Event) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateName() string {
	if m != nil {
		return m.AggregateName
	}
	return ""
}

func (m *Event) GetAggregateRevision() int32 {
	if m != nil {
		return m.AggregateRevision
	}
	return 0
}

func (m *Event) GetSagaId() string {
	if m != nil {
		return m.SagaId
	}
	return ""
}

func (m *Event) GetSagaName() string {
	if m != nil {
		return m.SagaName
	}
	return ""
}

func (m *Event) GetSagaRevision() int32 {
	if m != nil {
		return m.SagaRevision
	}
	return 0
}

func (m *Event) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Event) GetCorrelationType() string {
	if m != nil {
		return m.CorrelationType
	}
	return ""
}

func (m *Event) GetCorrelationName() string {
	if m != nil {
		return m.CorrelationName
	}
	return ""
}

func (m *Event) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Event) GetCreatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Event) GetSentAt() *timestamppb.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

func (m *Event) GetReceivedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

type Aggregate struct {
	Domain               string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Id                   string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Revision             int32                  `protobuf:"varint,4,opt,name=revision,proto3" json:"revision,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Aggregate) Reset()         { *m = Aggregate{} }
func (m *Aggregate) String() string { return proto.CompactTextString(m) }
func (*Aggregate) ProtoMessage()    {}
func (*Aggregate) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2ea67d34c1bc80, []int{2}
}

func (m *Aggregate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Aggregate.Unmarshal(m, b)
}
func (m *Aggregate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Aggregate.Marshal(b, m, deterministic)
}
func (m *Aggregate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aggregate.Merge(m, src)
}
func (m *Aggregate) XXX_Size() int {
	return xxx_messageInfo_Aggregate.Size(m)
}
func (m *Aggregate) XXX_DiscardUnknown() {
	xxx_messageInfo_Aggregate.DiscardUnknown(m)
}

var xxx_messageInfo_Aggregate proto.InternalMessageInfo

func (m *Aggregate) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Aggregate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Aggregate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Aggregate) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *Aggregate) GetCreatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Aggregate) GetUpdatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Entity struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2ea67d34c1bc80, []int{3}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entity) GetCreatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Entity) GetUpdatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ValueObject struct {
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ValueObject) Reset()         { *m = ValueObject{} }
func (m *ValueObject) String() string { return proto.CompactTextString(m) }
func (*ValueObject) ProtoMessage()    {}
func (*ValueObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2ea67d34c1bc80, []int{4}
}

func (m *ValueObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueObject.Unmarshal(m, b)
}
func (m *ValueObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueObject.Marshal(b, m, deterministic)
}
func (m *ValueObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueObject.Merge(m, src)
}
func (m *ValueObject) XXX_Size() int {
	return xxx_messageInfo_ValueObject.Size(m)
}
func (m *ValueObject) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueObject.DiscardUnknown(m)
}

var xxx_messageInfo_ValueObject proto.InternalMessageInfo

func (m *ValueObject) GetCreatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ValueObject) GetUpdatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "esproto.header.Command")
	proto.RegisterType((*Event)(nil), "esproto.header.Event")
	proto.RegisterType((*Aggregate)(nil), "esproto.header.Aggregate")
	proto.RegisterType((*Entity)(nil), "esproto.header.Entity")
	proto.RegisterType((*ValueObject)(nil), "esproto.header.ValueObject")
}

func init() {
	proto.RegisterFile("header/header.proto", fileDescriptor_aa2ea67d34c1bc80)
}

var fileDescriptor_aa2ea67d34c1bc80 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0x4f, 0x8f, 0xd2, 0x50,
	0x14, 0xc5, 0x53, 0x06, 0x5a, 0x7a, 0x0b, 0x85, 0x79, 0x26, 0xda, 0xe0, 0x42, 0x64, 0x42, 0x82,
	0x0b, 0x4b, 0xe2, 0xac, 0x26, 0xae, 0xd0, 0xb0, 0x60, 0xa3, 0x09, 0x99, 0xb8, 0x70, 0x33, 0x79,
	0xf0, 0xae, 0x9d, 0xa7, 0xb4, 0x8f, 0xb4, 0x0f, 0x22, 0x6e, 0xdd, 0x1a, 0xbf, 0x9f, 0xf1, 0xcb,
	0x98, 0xde, 0xfe, 0x19, 0x06, 0xa3, 0xfc, 0x71, 0xe5, 0xaa, 0xbd, 0xe7, 0x9e, 0xdb, 0x9e, 0xfb,
	0xfa, 0x4b, 0x0a, 0x0f, 0x6e, 0x91, 0x0b, 0x8c, 0x87, 0xd9, 0xc5, 0x5f, 0xc6, 0x4a, 0x2b, 0xe6,
	0x62, 0x42, 0x37, 0x7e, 0xa6, 0x76, 0x9e, 0x04, 0x4a, 0x05, 0x0b, 0x1c, 0x92, 0x38, 0x5b, 0x7d,
	0x18, 0x6a, 0x19, 0x62, 0xa2, 0x79, 0xb8, 0xcc, 0x06, 0x7a, 0xdf, 0x6a, 0x60, 0xbd, 0x56, 0x61,
	0xc8, 0x23, 0xc1, 0x1e, 0x82, 0x29, 0x54, 0xc8, 0x65, 0xe4, 0x19, 0x5d, 0x63, 0x60, 0x4f, 0xf3,
	0x8a, 0xb9, 0x50, 0x91, 0xc2, 0xab, 0x90, 0x56, 0x91, 0x82, 0x31, 0xa8, 0x46, 0x3c, 0x44, 0xef,
	0x8c, 0x14, 0xba, 0x67, 0x4f, 0xa1, 0xc1, 0x83, 0x20, 0xc6, 0x80, 0x6b, 0xbc, 0x91, 0xc2, 0xab,
	0x52, 0xcf, 0x29, 0xb5, 0x89, 0x60, 0x7d, 0x70, 0xef, 0x2c, 0xf4, 0x80, 0x1a, 0x99, 0x9a, 0xa5,
	0xfa, 0x26, 0x7d, 0xd2, 0x73, 0x60, 0x77, 0xb6, 0x18, 0xd7, 0x32, 0x91, 0x2a, 0xf2, 0xcc, 0xae,
	0x31, 0xa8, 0x4d, 0xcf, 0xcb, 0xce, 0x34, 0x6f, 0xdc, 0xb7, 0xe3, 0x67, 0x99, 0x68, 0x19, 0x05,
	0x9e, 0xd5, 0x35, 0x06, 0xf5, 0x2d, 0xfb, 0x38, 0x6f, 0xb0, 0x47, 0x60, 0x25, 0x3c, 0xe0, 0x69,
	0xc4, 0x7a, 0xb6, 0x64, 0x5a, 0x4e, 0x04, 0x7b, 0x0c, 0x36, 0x35, 0x28, 0x98, 0x4d, 0xad, 0x7a,
	0x2a, 0x50, 0xa6, 0x0b, 0x68, 0x52, 0xb3, 0x8c, 0x03, 0x14, 0xa7, 0x91, 0x8a, 0x65, 0x92, 0x3e,
	0xb8, 0x73, 0x15, 0xc7, 0xb8, 0xe0, 0x5a, 0xaa, 0x28, 0x7d, 0x83, 0x93, 0xed, 0xb7, 0xa5, 0x4e,
	0x04, 0x7b, 0x06, 0xed, 0x6d, 0x9b, 0xde, 0x2c, 0xd1, 0x6b, 0x90, 0xb1, 0xb5, 0xa5, 0x5f, 0x6f,
	0x96, 0xb8, 0x6b, 0xa5, 0x68, 0xcd, 0xdf, 0xac, 0x94, 0xd0, 0x03, 0x6b, 0x8d, 0x31, 0x65, 0x73,
	0x29, 0x5b, 0x51, 0xb2, 0x2b, 0x80, 0x79, 0x8c, 0x5c, 0xa3, 0xb8, 0xe1, 0xda, 0x6b, 0x75, 0x8d,
	0x81, 0xf3, 0xa2, 0xe3, 0x67, 0x5c, 0xf8, 0x05, 0x17, 0xfe, 0x75, 0xc1, 0xc5, 0xd4, 0xce, 0xdd,
	0x23, 0xcd, 0x2e, 0xc1, 0x4a, 0x30, 0xd2, 0xe9, 0x5c, 0x7b, 0xef, 0x9c, 0x99, 0x5a, 0x47, 0x9a,
	0xbd, 0x04, 0x27, 0xc6, 0x39, 0xca, 0x75, 0xf6, 0xc2, 0xf3, 0xbd, 0x83, 0x50, 0xd8, 0x47, 0xba,
	0xf7, 0xb3, 0x0a, 0xb5, 0xf1, 0x1a, 0x23, 0xfd, 0x7f, 0xc2, 0xb8, 0x45, 0x97, 0xf5, 0x67, 0xba,
	0xea, 0xfb, 0xe8, 0xb2, 0x0f, 0xa2, 0x0b, 0x0e, 0xa5, 0xcb, 0x39, 0x9c, 0xae, 0xc6, 0x5e, 0xba,
	0x9a, 0x7f, 0xa3, 0xcb, 0x3d, 0x91, 0xae, 0xd6, 0xa9, 0x74, 0xb5, 0x8f, 0xa2, 0xeb, 0x87, 0x01,
	0xf6, 0xa8, 0xf8, 0x68, 0xff, 0x44, 0x58, 0x07, 0xea, 0xe5, 0xd7, 0xaa, 0xd2, 0x89, 0x94, 0xf5,
	0xce, 0x91, 0xd4, 0x8e, 0x39, 0x92, 0x2b, 0x80, 0xd5, 0x52, 0x14, 0xa3, 0xe6, 0xfe, 0xd1, 0xdc,
	0x3d, 0xd2, 0xbd, 0xef, 0x06, 0x98, 0xe3, 0x48, 0x4b, 0xbd, 0xc9, 0x17, 0x30, 0xca, 0x05, 0xee,
	0x07, 0xaa, 0x9c, 0x1e, 0xe8, 0xec, 0x98, 0x40, 0x5f, 0x0d, 0x70, 0xde, 0xf1, 0xc5, 0x0a, 0xdf,
	0xce, 0x3e, 0xe2, 0x5c, 0xef, 0xa4, 0x30, 0x4e, 0x4f, 0x51, 0x39, 0x22, 0xc5, 0xab, 0xfe, 0xfb,
	0x8b, 0x40, 0xea, 0xdb, 0xd5, 0xcc, 0x9f, 0xab, 0x70, 0x18, 0xf3, 0x4f, 0x3c, 0x54, 0x6b, 0xf9,
	0x65, 0x98, 0xff, 0x27, 0xf3, 0xbf, 0xe7, 0xcc, 0xa4, 0xea, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x73, 0x78, 0x21, 0xf7, 0x55, 0x07, 0x00, 0x00,
}
