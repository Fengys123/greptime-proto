// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/frontend_heartbeat.proto

package meta

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InvalidateCache_Level int32

const (
	InvalidateCache_CATALOG InvalidateCache_Level = 0
	InvalidateCache_SCHEMA  InvalidateCache_Level = 1
	InvalidateCache_TABLE   InvalidateCache_Level = 2
)

// Enum value maps for InvalidateCache_Level.
var (
	InvalidateCache_Level_name = map[int32]string{
		0: "CATALOG",
		1: "SCHEMA",
		2: "TABLE",
	}
	InvalidateCache_Level_value = map[string]int32{
		"CATALOG": 0,
		"SCHEMA":  1,
		"TABLE":   2,
	}
)

func (x InvalidateCache_Level) Enum() *InvalidateCache_Level {
	p := new(InvalidateCache_Level)
	*p = x
	return p
}

func (x InvalidateCache_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvalidateCache_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_meta_frontend_heartbeat_proto_enumTypes[0].Descriptor()
}

func (InvalidateCache_Level) Type() protoreflect.EnumType {
	return &file_greptime_v1_meta_frontend_heartbeat_proto_enumTypes[0]
}

func (x InvalidateCache_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvalidateCache_Level.Descriptor instead.
func (InvalidateCache_Level) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP(), []int{4, 0}
}

type FrontendHeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Self peer
	Peer *Peer `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
	// Actually reported time interval
	ReportInterval *TimeInterval `protobuf:"bytes,3,opt,name=report_interval,json=reportInterval,proto3" json:"report_interval,omitempty"`
	// Node stat
	Stat *FrontendStat `protobuf:"bytes,4,opt,name=stat,proto3" json:"stat,omitempty"`
}

func (x *FrontendHeartbeatRequest) Reset() {
	*x = FrontendHeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendHeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendHeartbeatRequest) ProtoMessage() {}

func (x *FrontendHeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendHeartbeatRequest.ProtoReflect.Descriptor instead.
func (*FrontendHeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *FrontendHeartbeatRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FrontendHeartbeatRequest) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *FrontendHeartbeatRequest) GetReportInterval() *TimeInterval {
	if x != nil {
		return x.ReportInterval
	}
	return nil
}

func (x *FrontendHeartbeatRequest) GetStat() *FrontendStat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type FrontendStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rcus     int64   `protobuf:"varint,1,opt,name=rcus,proto3" json:"rcus,omitempty"`
	Wcus     int64   `protobuf:"varint,2,opt,name=wcus,proto3" json:"wcus,omitempty"`
	CpuUsage float64 `protobuf:"fixed64,3,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	Load     float64 `protobuf:"fixed64,4,opt,name=load,proto3" json:"load,omitempty"`
}

func (x *FrontendStat) Reset() {
	*x = FrontendStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendStat) ProtoMessage() {}

func (x *FrontendStat) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendStat.ProtoReflect.Descriptor instead.
func (*FrontendStat) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP(), []int{1}
}

func (x *FrontendStat) GetRcus() int64 {
	if x != nil {
		return x.Rcus
	}
	return 0
}

func (x *FrontendStat) GetWcus() int64 {
	if x != nil {
		return x.Wcus
	}
	return 0
}

func (x *FrontendStat) GetCpuUsage() float64 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *FrontendStat) GetLoad() float64 {
	if x != nil {
		return x.Load
	}
	return 0
}

type FrontendHeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header          *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	InstructionList []*Instruction  `protobuf:"bytes,2,rep,name=instruction_list,json=instructionList,proto3" json:"instruction_list,omitempty"`
}

func (x *FrontendHeartbeatResponse) Reset() {
	*x = FrontendHeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendHeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendHeartbeatResponse) ProtoMessage() {}

func (x *FrontendHeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendHeartbeatResponse.ProtoReflect.Descriptor instead.
func (*FrontendHeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP(), []int{2}
}

func (x *FrontendHeartbeatResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FrontendHeartbeatResponse) GetInstructionList() []*Instruction {
	if x != nil {
		return x.InstructionList
	}
	return nil
}

// The instruction sent to frontend from meta.
type Instruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*Instruction_InvalidCache
	Payload isInstruction_Payload `protobuf_oneof:"payload"`
}

func (x *Instruction) Reset() {
	*x = Instruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruction) ProtoMessage() {}

func (x *Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruction.ProtoReflect.Descriptor instead.
func (*Instruction) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP(), []int{3}
}

func (m *Instruction) GetPayload() isInstruction_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Instruction) GetInvalidCache() *InvalidateCache {
	if x, ok := x.GetPayload().(*Instruction_InvalidCache); ok {
		return x.InvalidCache
	}
	return nil
}

type isInstruction_Payload interface {
	isInstruction_Payload()
}

type Instruction_InvalidCache struct {
	InvalidCache *InvalidateCache `protobuf:"bytes,1,opt,name=invalid_cache,json=invalidCache,proto3,oneof"`
}

func (*Instruction_InvalidCache) isInstruction_Payload() {}

// An instruction to invalidate the catalog cache in frontend.
//
// When level = CATALOG, frontend needs to invalidate all caches under the catalog_name.
// When level = SCHEMA, frontend needs to invalidate all caches under the schema_name.
// When level = TABLE, frontend needs to invalidate all caches under the table_name.
type InvalidateCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level       InvalidateCache_Level `protobuf:"varint,1,opt,name=level,proto3,enum=greptime.v1.meta.InvalidateCache_Level" json:"level,omitempty"`
	CatalogName string                `protobuf:"bytes,2,opt,name=catalog_name,json=catalogName,proto3" json:"catalog_name,omitempty"`
	SchemaName  string                `protobuf:"bytes,3,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	TableName   string                `protobuf:"bytes,4,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *InvalidateCache) Reset() {
	*x = InvalidateCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidateCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateCache) ProtoMessage() {}

func (x *InvalidateCache) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateCache.ProtoReflect.Descriptor instead.
func (*InvalidateCache) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP(), []int{4}
}

func (x *InvalidateCache) GetLevel() InvalidateCache_Level {
	if x != nil {
		return x.Level
	}
	return InvalidateCache_CATALOG
}

func (x *InvalidateCache) GetCatalogName() string {
	if x != nil {
		return x.CatalogName
	}
	return ""
}

func (x *InvalidateCache) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *InvalidateCache) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

var File_greptime_v1_meta_frontend_heartbeat_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_frontend_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1d, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a,
	0x18, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x47,
	0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x22, 0x67, 0x0a, 0x0c, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x63, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x63, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x77, 0x63, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x77,
	0x63, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x19, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x10,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x48,
	0x00, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x3d,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x2b, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x54,
	0x41, 0x4c, 0x4f, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x32, 0x87, 0x01,
	0x0a, 0x11, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x12, 0x72, 0x0a, 0x11, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_frontend_heartbeat_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_frontend_heartbeat_proto_rawDescData = file_greptime_v1_meta_frontend_heartbeat_proto_rawDesc
)

func file_greptime_v1_meta_frontend_heartbeat_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_frontend_heartbeat_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_frontend_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_frontend_heartbeat_proto_rawDescData)
	})
	return file_greptime_v1_meta_frontend_heartbeat_proto_rawDescData
}

var file_greptime_v1_meta_frontend_heartbeat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_greptime_v1_meta_frontend_heartbeat_proto_goTypes = []interface{}{
	(InvalidateCache_Level)(0),        // 0: greptime.v1.meta.InvalidateCache.Level
	(*FrontendHeartbeatRequest)(nil),  // 1: greptime.v1.meta.FrontendHeartbeatRequest
	(*FrontendStat)(nil),              // 2: greptime.v1.meta.FrontendStat
	(*FrontendHeartbeatResponse)(nil), // 3: greptime.v1.meta.FrontendHeartbeatResponse
	(*Instruction)(nil),               // 4: greptime.v1.meta.Instruction
	(*InvalidateCache)(nil),           // 5: greptime.v1.meta.InvalidateCache
	(*RequestHeader)(nil),             // 6: greptime.v1.meta.RequestHeader
	(*Peer)(nil),                      // 7: greptime.v1.meta.Peer
	(*TimeInterval)(nil),              // 8: greptime.v1.meta.TimeInterval
	(*ResponseHeader)(nil),            // 9: greptime.v1.meta.ResponseHeader
}
var file_greptime_v1_meta_frontend_heartbeat_proto_depIdxs = []int32{
	6, // 0: greptime.v1.meta.FrontendHeartbeatRequest.header:type_name -> greptime.v1.meta.RequestHeader
	7, // 1: greptime.v1.meta.FrontendHeartbeatRequest.peer:type_name -> greptime.v1.meta.Peer
	8, // 2: greptime.v1.meta.FrontendHeartbeatRequest.report_interval:type_name -> greptime.v1.meta.TimeInterval
	2, // 3: greptime.v1.meta.FrontendHeartbeatRequest.stat:type_name -> greptime.v1.meta.FrontendStat
	9, // 4: greptime.v1.meta.FrontendHeartbeatResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	4, // 5: greptime.v1.meta.FrontendHeartbeatResponse.instruction_list:type_name -> greptime.v1.meta.Instruction
	5, // 6: greptime.v1.meta.Instruction.invalid_cache:type_name -> greptime.v1.meta.InvalidateCache
	0, // 7: greptime.v1.meta.InvalidateCache.level:type_name -> greptime.v1.meta.InvalidateCache.Level
	1, // 8: greptime.v1.meta.FrontendHeartbeat.FrontendHeartbeat:input_type -> greptime.v1.meta.FrontendHeartbeatRequest
	3, // 9: greptime.v1.meta.FrontendHeartbeat.FrontendHeartbeat:output_type -> greptime.v1.meta.FrontendHeartbeatResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_frontend_heartbeat_proto_init() }
func file_greptime_v1_meta_frontend_heartbeat_proto_init() {
	if File_greptime_v1_meta_frontend_heartbeat_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendHeartbeatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendStat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendHeartbeatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidateCache); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Instruction_InvalidCache)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greptime_v1_meta_frontend_heartbeat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_frontend_heartbeat_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_frontend_heartbeat_proto_depIdxs,
		EnumInfos:         file_greptime_v1_meta_frontend_heartbeat_proto_enumTypes,
		MessageInfos:      file_greptime_v1_meta_frontend_heartbeat_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_frontend_heartbeat_proto = out.File
	file_greptime_v1_meta_frontend_heartbeat_proto_rawDesc = nil
	file_greptime_v1_meta_frontend_heartbeat_proto_goTypes = nil
	file_greptime_v1_meta_frontend_heartbeat_proto_depIdxs = nil
}
