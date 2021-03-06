// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/log/log.proto

package go_micro_srv_log

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Logslevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logslevel []*Loglevel `protobuf:"bytes,1,rep,name=Logslevel,proto3" json:"Logslevel,omitempty"`
}

func (x *Logslevel) Reset() {
	*x = Logslevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logslevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logslevel) ProtoMessage() {}

func (x *Logslevel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logslevel.ProtoReflect.Descriptor instead.
func (*Logslevel) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{0}
}

func (x *Logslevel) GetLogslevel() []*Loglevel {
	if x != nil {
		return x.Logslevel
	}
	return nil
}

type LogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogsResponse []*LogResponse `protobuf:"bytes,1,rep,name=LogsResponse,proto3" json:"LogsResponse,omitempty"`
}

func (x *LogsResponse) Reset() {
	*x = LogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResponse) ProtoMessage() {}

func (x *LogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResponse.ProtoReflect.Descriptor instead.
func (*LogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogsResponse) GetLogsResponse() []*LogResponse {
	if x != nil {
		return x.LogsResponse
	}
	return nil
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{2}
}

type Loglevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Loglevel) Reset() {
	*x = Loglevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loglevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loglevel) ProtoMessage() {}

func (x *Loglevel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loglevel.ProtoReflect.Descriptor instead.
func (*Loglevel) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{3}
}

func (x *Loglevel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Loglevel) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{4}
}

func (x *LogResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogKeepDays uint32 `protobuf:"varint,1,opt,name=logKeepDays,proto3" json:"logKeepDays,omitempty"`
	LogLevel    uint32 `protobuf:"varint,2,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
}

func (x *LogConfig) Reset() {
	*x = LogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogConfig) ProtoMessage() {}

func (x *LogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogConfig.ProtoReflect.Descriptor instead.
func (*LogConfig) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{5}
}

func (x *LogConfig) GetLogKeepDays() uint32 {
	if x != nil {
		return x.LogKeepDays
	}
	return 0
}

func (x *LogConfig) GetLogLevel() uint32 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

type LogContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Level    string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	FuncName string `protobuf:"bytes,3,opt,name=funcName,proto3" json:"funcName,omitempty"`
	FileName string `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Line     uint32 `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
	Service  string `protobuf:"bytes,6,opt,name=service,proto3" json:"service,omitempty"`
	IpAddr   string `protobuf:"bytes,7,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	Msg      string `protobuf:"bytes,8,opt,name=msg,proto3" json:"msg,omitempty"`
	TaskID   uint32 `protobuf:"varint,9,opt,name=taskID,proto3" json:"taskID,omitempty"`
}

func (x *LogContent) Reset() {
	*x = LogContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogContent) ProtoMessage() {}

func (x *LogContent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogContent.ProtoReflect.Descriptor instead.
func (*LogContent) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{6}
}

func (x *LogContent) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *LogContent) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogContent) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *LogContent) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *LogContent) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *LogContent) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *LogContent) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *LogContent) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LogContent) GetTaskID() uint32 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{7}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{8}
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    uint32     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage uint32     `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Filter  *LogFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	SubPage uint32     `protobuf:"varint,4,opt,name=subPage,proto3" json:"subPage,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{9}
}

func (x *Query) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Query) GetPerPage() uint32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *Query) GetFilter() *LogFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *Query) GetSubPage() uint32 {
	if x != nil {
		return x.SubPage
	}
	return 0
}

type LogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	And       []*LogFilter `protobuf:"bytes,1,rep,name=and,proto3" json:"and,omitempty"`
	Or        []*LogFilter `protobuf:"bytes,2,rep,name=or,proto3" json:"or,omitempty"`
	FromTime  string       `protobuf:"bytes,3,opt,name=fromTime,proto3" json:"fromTime,omitempty"`
	ToTime    string       `protobuf:"bytes,4,opt,name=toTime,proto3" json:"toTime,omitempty"`
	LevelIn   []string     `protobuf:"bytes,5,rep,name=levelIn,proto3" json:"levelIn,omitempty"`
	FuncIn    []string     `protobuf:"bytes,6,rep,name=funcIn,proto3" json:"funcIn,omitempty"`
	FileIn    []string     `protobuf:"bytes,7,rep,name=fileIn,proto3" json:"fileIn,omitempty"`
	LineIn    []uint32     `protobuf:"varint,8,rep,packed,name=lineIn,proto3" json:"lineIn,omitempty"`
	ServiceIn []string     `protobuf:"bytes,9,rep,name=serviceIn,proto3" json:"serviceIn,omitempty"`
	IpIn      []string     `protobuf:"bytes,10,rep,name=ipIn,proto3" json:"ipIn,omitempty"`
	MsgIn     []string     `protobuf:"bytes,11,rep,name=msgIn,proto3" json:"msgIn,omitempty"`
	TaskID    uint32       `protobuf:"varint,12,opt,name=taskID,proto3" json:"taskID,omitempty"`
	PerNum    uint32       `protobuf:"varint,13,opt,name=perNum,proto3" json:"perNum,omitempty"`
	NexNum    uint32       `protobuf:"varint,14,opt,name=nexNum,proto3" json:"nexNum,omitempty"`
}

func (x *LogFilter) Reset() {
	*x = LogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFilter) ProtoMessage() {}

func (x *LogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFilter.ProtoReflect.Descriptor instead.
func (*LogFilter) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{10}
}

func (x *LogFilter) GetAnd() []*LogFilter {
	if x != nil {
		return x.And
	}
	return nil
}

func (x *LogFilter) GetOr() []*LogFilter {
	if x != nil {
		return x.Or
	}
	return nil
}

func (x *LogFilter) GetFromTime() string {
	if x != nil {
		return x.FromTime
	}
	return ""
}

func (x *LogFilter) GetToTime() string {
	if x != nil {
		return x.ToTime
	}
	return ""
}

func (x *LogFilter) GetLevelIn() []string {
	if x != nil {
		return x.LevelIn
	}
	return nil
}

func (x *LogFilter) GetFuncIn() []string {
	if x != nil {
		return x.FuncIn
	}
	return nil
}

func (x *LogFilter) GetFileIn() []string {
	if x != nil {
		return x.FileIn
	}
	return nil
}

func (x *LogFilter) GetLineIn() []uint32 {
	if x != nil {
		return x.LineIn
	}
	return nil
}

func (x *LogFilter) GetServiceIn() []string {
	if x != nil {
		return x.ServiceIn
	}
	return nil
}

func (x *LogFilter) GetIpIn() []string {
	if x != nil {
		return x.IpIn
	}
	return nil
}

func (x *LogFilter) GetMsgIn() []string {
	if x != nil {
		return x.MsgIn
	}
	return nil
}

func (x *LogFilter) GetTaskID() uint32 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *LogFilter) GetPerNum() uint32 {
	if x != nil {
		return x.PerNum
	}
	return 0
}

func (x *LogFilter) GetNexNum() uint32 {
	if x != nil {
		return x.NexNum
	}
	return 0
}

type LogContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogContents []*LogContent `protobuf:"bytes,1,rep,name=logContents,proto3" json:"logContents,omitempty"`
	TotalCount  uint32        `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
}

func (x *LogContents) Reset() {
	*x = LogContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_log_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogContents) ProtoMessage() {}

func (x *LogContents) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_log_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogContents.ProtoReflect.Descriptor instead.
func (*LogContents) Descriptor() ([]byte, []int) {
	return file_proto_log_log_proto_rawDescGZIP(), []int{11}
}

func (x *LogContents) GetLogContents() []*LogContent {
	if x != nil {
		return x.LogContents
	}
	return nil
}

func (x *LogContents) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_proto_log_log_proto protoreflect.FileDescriptor

var file_proto_log_log_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x22, 0x45, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x51,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x34, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x4b, 0x65, 0x65, 0x70,
	0x44, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x4b,
	0x65, 0x65, 0x70, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0xde, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x50,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x75, 0x62, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x8d, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x03, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x03, 0x61, 0x6e, 0x64, 0x12,
	0x2b, 0x0a, 0x02, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x02, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x75,
	0x6e, 0x63, 0x49, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x75, 0x6e, 0x63,
	0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69,
	0x6e, 0x65, 0x49, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65,
	0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x49, 0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x70, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x6e, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x70, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65,
	0x78, 0x4e, 0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x4e,
	0x75, 0x6d, 0x22, 0x6d, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x90, 0x05, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x19, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1d, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x4c, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x48, 0x0a, 0x0b, 0x53, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x47,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4c, 0x6f, 0x67, 0x12,
	0x17, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f,
	0x67, 0x3b, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x73, 0x72, 0x76, 0x5f, 0x6c,
	0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_log_log_proto_rawDescOnce sync.Once
	file_proto_log_log_proto_rawDescData = file_proto_log_log_proto_rawDesc
)

func file_proto_log_log_proto_rawDescGZIP() []byte {
	file_proto_log_log_proto_rawDescOnce.Do(func() {
		file_proto_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_log_log_proto_rawDescData)
	})
	return file_proto_log_log_proto_rawDescData
}

var file_proto_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_log_log_proto_goTypes = []interface{}{
	(*Logslevel)(nil),    // 0: go.micro.srv.log.Logslevel
	(*LogsResponse)(nil), // 1: go.micro.srv.log.LogsResponse
	(*LogRequest)(nil),   // 2: go.micro.srv.log.LogRequest
	(*Loglevel)(nil),     // 3: go.micro.srv.log.Loglevel
	(*LogResponse)(nil),  // 4: go.micro.srv.log.LogResponse
	(*LogConfig)(nil),    // 5: go.micro.srv.log.LogConfig
	(*LogContent)(nil),   // 6: go.micro.srv.log.LogContent
	(*Request)(nil),      // 7: go.micro.srv.log.Request
	(*Response)(nil),     // 8: go.micro.srv.log.Response
	(*Query)(nil),        // 9: go.micro.srv.log.Query
	(*LogFilter)(nil),    // 10: go.micro.srv.log.LogFilter
	(*LogContents)(nil),  // 11: go.micro.srv.log.LogContents
}
var file_proto_log_log_proto_depIdxs = []int32{
	3,  // 0: go.micro.srv.log.Logslevel.Logslevel:type_name -> go.micro.srv.log.Loglevel
	4,  // 1: go.micro.srv.log.LogsResponse.LogsResponse:type_name -> go.micro.srv.log.LogResponse
	10, // 2: go.micro.srv.log.Query.filter:type_name -> go.micro.srv.log.LogFilter
	10, // 3: go.micro.srv.log.LogFilter.and:type_name -> go.micro.srv.log.LogFilter
	10, // 4: go.micro.srv.log.LogFilter.or:type_name -> go.micro.srv.log.LogFilter
	6,  // 5: go.micro.srv.log.LogContents.logContents:type_name -> go.micro.srv.log.LogContent
	5,  // 6: go.micro.srv.log.Log.UpdateConfig:input_type -> go.micro.srv.log.LogConfig
	7,  // 7: go.micro.srv.log.Log.GetConfig:input_type -> go.micro.srv.log.Request
	9,  // 8: go.micro.srv.log.Log.GetLogs:input_type -> go.micro.srv.log.Query
	6,  // 9: go.micro.srv.log.Log.WriteLog:input_type -> go.micro.srv.log.LogContent
	2,  // 10: go.micro.srv.log.Log.GetLogLevel:input_type -> go.micro.srv.log.LogRequest
	3,  // 11: go.micro.srv.log.Log.SetLogLevel:input_type -> go.micro.srv.log.Loglevel
	2,  // 12: go.micro.srv.log.Log.GetLogsLevel:input_type -> go.micro.srv.log.LogRequest
	9,  // 13: go.micro.srv.log.Log.GetContextLog:input_type -> go.micro.srv.log.Query
	0,  // 14: go.micro.srv.log.Log.SetLogsLevel:input_type -> go.micro.srv.log.Logslevel
	8,  // 15: go.micro.srv.log.Log.UpdateConfig:output_type -> go.micro.srv.log.Response
	5,  // 16: go.micro.srv.log.Log.GetConfig:output_type -> go.micro.srv.log.LogConfig
	11, // 17: go.micro.srv.log.Log.GetLogs:output_type -> go.micro.srv.log.LogContents
	8,  // 18: go.micro.srv.log.Log.WriteLog:output_type -> go.micro.srv.log.Response
	3,  // 19: go.micro.srv.log.Log.GetLogLevel:output_type -> go.micro.srv.log.Loglevel
	4,  // 20: go.micro.srv.log.Log.SetLogLevel:output_type -> go.micro.srv.log.LogResponse
	0,  // 21: go.micro.srv.log.Log.GetLogsLevel:output_type -> go.micro.srv.log.Logslevel
	11, // 22: go.micro.srv.log.Log.GetContextLog:output_type -> go.micro.srv.log.LogContents
	1,  // 23: go.micro.srv.log.Log.SetLogsLevel:output_type -> go.micro.srv.log.LogsResponse
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_log_log_proto_init() }
func file_proto_log_log_proto_init() {
	if File_proto_log_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_log_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logslevel); i {
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
		file_proto_log_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsResponse); i {
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
		file_proto_log_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_proto_log_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loglevel); i {
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
		file_proto_log_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
		file_proto_log_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogConfig); i {
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
		file_proto_log_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogContent); i {
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
		file_proto_log_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_log_log_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_log_log_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_proto_log_log_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogFilter); i {
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
		file_proto_log_log_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogContents); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_log_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_log_log_proto_goTypes,
		DependencyIndexes: file_proto_log_log_proto_depIdxs,
		MessageInfos:      file_proto_log_log_proto_msgTypes,
	}.Build()
	File_proto_log_log_proto = out.File
	file_proto_log_log_proto_rawDesc = nil
	file_proto_log_log_proto_goTypes = nil
	file_proto_log_log_proto_depIdxs = nil
}
