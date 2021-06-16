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
	CorrelationId        string                 `protobuf:"bytes,9,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	CorrelationType      string                 `protobuf:"bytes,10,opt,name=correlation_type,json=correlationType,proto3" json:"correlation_type,omitempty"`
	CorrelationName      string                 `protobuf:"bytes,11,opt,name=correlation_name,json=correlationName,proto3" json:"correlation_name,omitempty"`
	Version              int32                  `protobuf:"varint,12,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SentAt               *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	ReceivedAt           *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
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
	IsSaga               bool                   `protobuf:"varint,5,opt,name=is_saga,json=isSaga,proto3" json:"is_saga,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
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

func (m *Aggregate) GetIsSaga() bool {
	if m != nil {
		return m.IsSaga
	}
	return false
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
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0x4f, 0x8f, 0xd2, 0x5e,
	0x14, 0x4d, 0x19, 0xe8, 0x9f, 0x5b, 0x28, 0xf3, 0x7b, 0xbf, 0x44, 0x1b, 0x5c, 0x88, 0x18, 0x12,
	0x5c, 0x58, 0x12, 0x67, 0x35, 0x71, 0x85, 0x86, 0x05, 0x1b, 0x4d, 0x70, 0xe2, 0xc2, 0x0d, 0x79,
	0xf0, 0xae, 0x9d, 0xa7, 0xb4, 0x8f, 0xb4, 0x0f, 0x22, 0x6e, 0xdd, 0xfb, 0x71, 0xfc, 0x4a, 0xee,
	0xfd, 0x04, 0xa6, 0xb7, 0x7f, 0x86, 0xc1, 0x28, 0x03, 0xae, 0x5c, 0xb5, 0xf7, 0xdc, 0x73, 0xfb,
	0xce, 0xbb, 0xe7, 0x24, 0x85, 0xff, 0xaf, 0x91, 0x0b, 0x4c, 0x86, 0xf9, 0x23, 0x58, 0x25, 0x4a,
	0x2b, 0xe6, 0x61, 0x4a, 0x2f, 0x41, 0x8e, 0x76, 0x1e, 0x86, 0x4a, 0x85, 0x4b, 0x1c, 0x12, 0x38,
	0x5f, 0xbf, 0x1f, 0x6a, 0x19, 0x61, 0xaa, 0x79, 0xb4, 0xca, 0x07, 0x7a, 0xdf, 0xeb, 0x60, 0xbd,
	0x54, 0x51, 0xc4, 0x63, 0xc1, 0xee, 0x81, 0x29, 0x54, 0xc4, 0x65, 0xec, 0x1b, 0x5d, 0x63, 0xe0,
	0x4c, 0x8b, 0x8a, 0x79, 0x50, 0x93, 0xc2, 0xaf, 0x11, 0x56, 0x93, 0x82, 0x31, 0xa8, 0xc7, 0x3c,
	0x42, 0xff, 0x8c, 0x10, 0x7a, 0x67, 0x8f, 0xa0, 0xc9, 0xc3, 0x30, 0xc1, 0x90, 0x6b, 0x9c, 0x49,
	0xe1, 0xd7, 0xa9, 0xe7, 0x56, 0xd8, 0x44, 0xb0, 0x3e, 0x78, 0x37, 0x14, 0xfa, 0x40, 0x83, 0x48,
	0xad, 0x0a, 0x7d, 0x95, 0x7d, 0xe9, 0x29, 0xb0, 0x1b, 0x5a, 0x82, 0x1b, 0x99, 0x4a, 0x15, 0xfb,
	0x66, 0xd7, 0x18, 0x34, 0xa6, 0xff, 0x55, 0x9d, 0x69, 0xd1, 0xb8, 0x4d, 0xc7, 0x4f, 0x32, 0xd5,
	0x32, 0x0e, 0x7d, 0xab, 0x6b, 0x0c, 0xec, 0x1d, 0xfa, 0xb8, 0x68, 0xb0, 0xfb, 0x60, 0xa5, 0x3c,
	0xe4, 0x99, 0x44, 0x3b, 0xbf, 0x64, 0x56, 0x4e, 0x04, 0x7b, 0x00, 0x0e, 0x35, 0x48, 0x98, 0x43,
	0x2d, 0x3b, 0x03, 0x48, 0x53, 0x1f, 0xbc, 0x85, 0x4a, 0x12, 0x5c, 0x72, 0x2d, 0x55, 0x9c, 0x0d,
	0x43, 0x2e, 0x7d, 0x07, 0x9d, 0x08, 0xf6, 0x04, 0xce, 0x77, 0x69, 0x7a, 0xbb, 0x42, 0xdf, 0x25,
	0x62, 0x7b, 0x07, 0xbf, 0xda, 0xae, 0x70, 0x9f, 0x4a, 0xa7, 0x36, 0x7f, 0xa1, 0xd2, 0xe1, 0x3e,
	0x58, 0x1b, 0x4c, 0x68, 0x0b, 0x2d, 0xda, 0x42, 0x59, 0xb2, 0x4b, 0x80, 0x45, 0x82, 0x5c, 0xa3,
	0x98, 0x71, 0xed, 0x7b, 0x5d, 0x63, 0xe0, 0x3e, 0xeb, 0x04, 0xb9, 0xe5, 0x41, 0x69, 0x79, 0x70,
	0x55, 0x5a, 0x3e, 0x75, 0x0a, 0xf6, 0x48, 0xb3, 0x0b, 0xb0, 0x52, 0x8c, 0x75, 0x36, 0xd7, 0x3e,
	0x38, 0x67, 0x66, 0xd4, 0x91, 0x66, 0xcf, 0xc1, 0x4d, 0x70, 0x81, 0x72, 0x93, 0x1f, 0x78, 0x7e,
	0x70, 0x10, 0x4a, 0xfa, 0x48, 0xf7, 0xbe, 0xd5, 0xa1, 0x31, 0xde, 0x60, 0xac, 0xff, 0xcd, 0x9c,
	0xed, 0x04, 0xc7, 0xfa, 0x7d, 0x70, 0xec, 0x83, 0xc1, 0x71, 0xee, 0x1a, 0x1c, 0xb8, 0x7b, 0x70,
	0xdc, 0x83, 0xc1, 0x69, 0xfe, 0x29, 0x38, 0xad, 0x13, 0x83, 0xe3, 0x9d, 0x1a, 0x9c, 0xf6, 0x51,
	0xc1, 0xf9, 0x61, 0x80, 0x33, 0x2a, 0xfd, 0xf8, 0xab, 0xf0, 0x74, 0xc0, 0xae, 0x8c, 0xae, 0xd3,
	0x46, 0xaa, 0x3a, 0xf3, 0x57, 0xa6, 0xb3, 0xcc, 0x38, 0x8a, 0x8b, 0x3d, 0x35, 0x65, 0xfa, 0x86,
	0x87, 0x7c, 0x6f, 0x57, 0xe6, 0x31, 0xbb, 0xba, 0x04, 0x58, 0xaf, 0x44, 0x39, 0x6a, 0x1d, 0x1e,
	0x2d, 0xd8, 0x23, 0xdd, 0xfb, 0x6a, 0x80, 0x39, 0x8e, 0xb5, 0xd4, 0xdb, 0xe2, 0x66, 0x46, 0x75,
	0xb3, 0xdb, 0x82, 0x6a, 0xa7, 0x0b, 0x3a, 0x3b, 0x46, 0xd0, 0x17, 0x03, 0xdc, 0xb7, 0x7c, 0xb9,
	0xc6, 0xd7, 0xf3, 0x0f, 0xb8, 0xd0, 0x7b, 0x2a, 0x8c, 0xd3, 0x55, 0xd4, 0x8e, 0x50, 0xf1, 0xa2,
	0xff, 0xee, 0x71, 0x28, 0xf5, 0xf5, 0x7a, 0x1e, 0x2c, 0x54, 0x34, 0x4c, 0xf8, 0x47, 0x1e, 0xa9,
	0x8d, 0xfc, 0x3c, 0x2c, 0x7e, 0x7b, 0xc5, 0xcf, 0x70, 0x6e, 0x52, 0x75, 0xf1, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xbb, 0x5d, 0xce, 0x24, 0x07, 0x00, 0x00,
}
