// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/log/log.proto

package go_micro_srv_log

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

type CfgReq struct {
	LogKeepTime          uint32   `protobuf:"varint,1,opt,name=logKeepTime,proto3" json:"logKeepTime,omitempty"`
	FilenameFormat       string   `protobuf:"bytes,2,opt,name=filenameFormat,proto3" json:"filenameFormat,omitempty"`
	ReportCaller         bool     `protobuf:"varint,3,opt,name=reportCaller,proto3" json:"reportCaller,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CfgReq) Reset()         { *m = CfgReq{} }
func (m *CfgReq) String() string { return proto.CompactTextString(m) }
func (*CfgReq) ProtoMessage()    {}
func (*CfgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{0}
}

func (m *CfgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CfgReq.Unmarshal(m, b)
}
func (m *CfgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CfgReq.Marshal(b, m, deterministic)
}
func (m *CfgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CfgReq.Merge(m, src)
}
func (m *CfgReq) XXX_Size() int {
	return xxx_messageInfo_CfgReq.Size(m)
}
func (m *CfgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CfgReq.DiscardUnknown(m)
}

var xxx_messageInfo_CfgReq proto.InternalMessageInfo

func (m *CfgReq) GetLogKeepTime() uint32 {
	if m != nil {
		return m.LogKeepTime
	}
	return 0
}

func (m *CfgReq) GetFilenameFormat() string {
	if m != nil {
		return m.FilenameFormat
	}
	return ""
}

func (m *CfgReq) GetReportCaller() bool {
	if m != nil {
		return m.ReportCaller
	}
	return false
}

type LogContent struct {
	Time                 string   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	FuncName             string   `protobuf:"bytes,3,opt,name=funcName,proto3" json:"funcName,omitempty"`
	FileName             string   `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Service              string   `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	IpAddr               string   `protobuf:"bytes,6,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	Msg                  string   `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogContent) Reset()         { *m = LogContent{} }
func (m *LogContent) String() string { return proto.CompactTextString(m) }
func (*LogContent) ProtoMessage()    {}
func (*LogContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{1}
}

func (m *LogContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogContent.Unmarshal(m, b)
}
func (m *LogContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogContent.Marshal(b, m, deterministic)
}
func (m *LogContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogContent.Merge(m, src)
}
func (m *LogContent) XXX_Size() int {
	return xxx_messageInfo_LogContent.Size(m)
}
func (m *LogContent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogContent.DiscardUnknown(m)
}

var xxx_messageInfo_LogContent proto.InternalMessageInfo

func (m *LogContent) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *LogContent) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LogContent) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *LogContent) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *LogContent) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *LogContent) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

func (m *LogContent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type Query struct {
	Page                 uint32     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32     `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Filter               *LogFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{4}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Query) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *Query) GetFilter() *LogFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type LogFilter struct {
	And                  []*LogFilter `protobuf:"bytes,1,rep,name=and,proto3" json:"and,omitempty"`
	Or                   []*LogFilter `protobuf:"bytes,2,rep,name=or,proto3" json:"or,omitempty"`
	FromTime             string       `protobuf:"bytes,3,opt,name=fromTime,proto3" json:"fromTime,omitempty"`
	ToTime               string       `protobuf:"bytes,4,opt,name=toTime,proto3" json:"toTime,omitempty"`
	LevelIn              []int32      `protobuf:"varint,5,rep,packed,name=levelIn,proto3" json:"levelIn,omitempty"`
	FileName             string       `protobuf:"bytes,6,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FuncName             string       `protobuf:"bytes,7,opt,name=funcName,proto3" json:"funcName,omitempty"`
	IpAddr               string       `protobuf:"bytes,8,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	LikeMsg              string       `protobuf:"bytes,9,opt,name=likeMsg,proto3" json:"likeMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogFilter) Reset()         { *m = LogFilter{} }
func (m *LogFilter) String() string { return proto.CompactTextString(m) }
func (*LogFilter) ProtoMessage()    {}
func (*LogFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{5}
}

func (m *LogFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogFilter.Unmarshal(m, b)
}
func (m *LogFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogFilter.Marshal(b, m, deterministic)
}
func (m *LogFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogFilter.Merge(m, src)
}
func (m *LogFilter) XXX_Size() int {
	return xxx_messageInfo_LogFilter.Size(m)
}
func (m *LogFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_LogFilter.DiscardUnknown(m)
}

var xxx_messageInfo_LogFilter proto.InternalMessageInfo

func (m *LogFilter) GetAnd() []*LogFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *LogFilter) GetOr() []*LogFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *LogFilter) GetFromTime() string {
	if m != nil {
		return m.FromTime
	}
	return ""
}

func (m *LogFilter) GetToTime() string {
	if m != nil {
		return m.ToTime
	}
	return ""
}

func (m *LogFilter) GetLevelIn() []int32 {
	if m != nil {
		return m.LevelIn
	}
	return nil
}

func (m *LogFilter) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *LogFilter) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *LogFilter) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

func (m *LogFilter) GetLikeMsg() string {
	if m != nil {
		return m.LikeMsg
	}
	return ""
}

type LogContents struct {
	LogContents          []*LogContent `protobuf:"bytes,1,rep,name=logContents,proto3" json:"logContents,omitempty"`
	TotalCount           uint32        `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LogContents) Reset()         { *m = LogContents{} }
func (m *LogContents) String() string { return proto.CompactTextString(m) }
func (*LogContents) ProtoMessage()    {}
func (*LogContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_c672f68bc74af5b1, []int{6}
}

func (m *LogContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogContents.Unmarshal(m, b)
}
func (m *LogContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogContents.Marshal(b, m, deterministic)
}
func (m *LogContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogContents.Merge(m, src)
}
func (m *LogContents) XXX_Size() int {
	return xxx_messageInfo_LogContents.Size(m)
}
func (m *LogContents) XXX_DiscardUnknown() {
	xxx_messageInfo_LogContents.DiscardUnknown(m)
}

var xxx_messageInfo_LogContents proto.InternalMessageInfo

func (m *LogContents) GetLogContents() []*LogContent {
	if m != nil {
		return m.LogContents
	}
	return nil
}

func (m *LogContents) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*CfgReq)(nil), "go.micro.srv.log.CfgReq")
	proto.RegisterType((*LogContent)(nil), "go.micro.srv.log.LogContent")
	proto.RegisterType((*Request)(nil), "go.micro.srv.log.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.log.Response")
	proto.RegisterType((*Query)(nil), "go.micro.srv.log.Query")
	proto.RegisterType((*LogFilter)(nil), "go.micro.srv.log.LogFilter")
	proto.RegisterType((*LogContents)(nil), "go.micro.srv.log.LogContents")
}

func init() { proto.RegisterFile("proto/log/log.proto", fileDescriptor_c672f68bc74af5b1) }

var fileDescriptor_c672f68bc74af5b1 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0x49, 0x37, 0x1f, 0xb7, 0x55, 0xca, 0x55, 0x74, 0xba, 0x7e, 0x10, 0xf2, 0x20, 0x0b,
	0x62, 0x84, 0xf6, 0x5d, 0x2c, 0x5b, 0x5a, 0xc4, 0x55, 0x34, 0x54, 0x7c, 0x94, 0xd8, 0xbd, 0x3b,
	0x84, 0x26, 0xb9, 0xe9, 0x64, 0x76, 0xc1, 0x3f, 0xe2, 0xab, 0xff, 0xc1, 0x5f, 0x28, 0x99, 0x9d,
	0xb4, 0xdb, 0xdd, 0xba, 0x7d, 0x58, 0x98, 0x73, 0xce, 0x9d, 0xb9, 0xdc, 0x73, 0xee, 0x06, 0x1e,
	0x35, 0x8a, 0x35, 0xbf, 0x2d, 0x59, 0x76, 0xbf, 0xd4, 0x20, 0xdc, 0x97, 0x9c, 0x56, 0xc5, 0x85,
	0xe2, 0xb4, 0x55, 0x8b, 0xb4, 0x64, 0x99, 0x2c, 0xc0, 0x1f, 0xcf, 0x64, 0x46, 0x57, 0x18, 0xc3,
	0x6e, 0xc9, 0xf2, 0x23, 0x51, 0x73, 0x5e, 0x54, 0x24, 0x9c, 0xd8, 0x19, 0x3d, 0xc8, 0x56, 0x29,
	0x7c, 0x05, 0x0f, 0x67, 0x45, 0x49, 0x75, 0x5e, 0xd1, 0x29, 0xab, 0x2a, 0xd7, 0xc2, 0x8d, 0x9d,
	0x51, 0x94, 0xad, 0xb1, 0x98, 0xc0, 0x9e, 0xa2, 0x86, 0x95, 0x1e, 0xe7, 0x65, 0x49, 0x4a, 0x78,
	0xb1, 0x33, 0x0a, 0xb3, 0x5b, 0x5c, 0xf2, 0xd7, 0x01, 0x98, 0xb0, 0x1c, 0x73, 0xad, 0xa9, 0xd6,
	0x88, 0xb0, 0xa3, 0xfb, 0xae, 0x51, 0x66, 0xce, 0xf8, 0x18, 0x06, 0x25, 0x2d, 0xa8, 0xb4, 0x5d,
	0x96, 0x00, 0x87, 0x10, 0xce, 0xe6, 0xf5, 0xc5, 0xe7, 0xbc, 0x22, 0xf3, 0x70, 0x94, 0x5d, 0x63,
	0xa3, 0x15, 0x25, 0x19, 0x6d, 0xc7, 0x6a, 0x16, 0xa3, 0x80, 0xa0, 0x25, 0xb5, 0x28, 0x2e, 0x48,
	0x0c, 0x8c, 0xd4, 0x43, 0x7c, 0x02, 0x7e, 0xd1, 0x1c, 0x4f, 0xa7, 0x4a, 0xf8, 0x46, 0xb0, 0x08,
	0xf7, 0xc1, 0xab, 0x5a, 0x29, 0x02, 0x43, 0x76, 0xc7, 0x24, 0x82, 0x20, 0xa3, 0xab, 0x39, 0xb5,
	0x3a, 0x01, 0x08, 0x33, 0x6a, 0x1b, 0xae, 0x5b, 0x4a, 0x2e, 0x61, 0xf0, 0x75, 0x4e, 0xea, 0x57,
	0x37, 0x45, 0x93, 0xcb, 0xde, 0x3b, 0x73, 0xc6, 0x03, 0x08, 0x1b, 0x52, 0x3f, 0x0c, 0xef, 0x1a,
	0x3e, 0x68, 0x48, 0x7d, 0xe9, 0xa4, 0x23, 0xf0, 0x67, 0x45, 0xa9, 0xad, 0x43, 0xbb, 0x87, 0xcf,
	0xd2, 0xf5, 0x78, 0xd2, 0x09, 0xcb, 0x53, 0x53, 0x92, 0xd9, 0xd2, 0xe4, 0x8f, 0x0b, 0xd1, 0x35,
	0x8b, 0x6f, 0xc0, 0xcb, 0xeb, 0xa9, 0x70, 0x62, 0xef, 0xbe, 0xfb, 0x5d, 0x1d, 0xbe, 0x06, 0x97,
	0x95, 0x70, 0xef, 0xaf, 0x76, 0x59, 0x19, 0x37, 0x15, 0x57, 0x66, 0x1b, 0x7a, 0xa7, 0x2d, 0xee,
	0x3c, 0xd3, 0x6c, 0x94, 0xa5, 0xcf, 0x16, 0x75, 0x2e, 0x9b, 0x98, 0x3e, 0xd4, 0x62, 0x10, 0x7b,
	0xa3, 0x41, 0xd6, 0xc3, 0x5b, 0xd9, 0xf8, 0x6b, 0xd9, 0xac, 0x66, 0x1a, 0xac, 0x65, 0x7a, 0x93,
	0x4e, 0x78, 0x2b, 0x9d, 0xae, 0x53, 0x71, 0x49, 0x9f, 0x5a, 0x29, 0xa2, 0x65, 0x9e, 0x16, 0x26,
	0x15, 0xec, 0xde, 0x6c, 0x56, 0x8b, 0xef, 0xcc, 0x5e, 0xf7, 0xd0, 0x5a, 0xf5, 0xfc, 0xce, 0xe1,
	0x6d, 0x51, 0xb6, 0x7a, 0x01, 0x5f, 0x02, 0x9c, 0xb3, 0xce, 0xcb, 0x31, 0xcf, 0x6b, 0x6d, 0x23,
	0x5c, 0x61, 0x0e, 0x7f, 0xbb, 0xe0, 0x4d, 0x58, 0xe2, 0x09, 0xec, 0x7d, 0x6b, 0xa6, 0xb9, 0xa6,
	0x31, 0xd7, 0xb3, 0x42, 0xa2, 0xd8, 0x6c, 0xb1, 0xfc, 0xa7, 0x0d, 0x87, 0x9b, 0x4a, 0xbf, 0x4b,
	0xf8, 0x1e, 0xa2, 0x33, 0xd2, 0xf6, 0x89, 0x83, 0xbb, 0x0a, 0xcd, 0xfe, 0x0d, 0xff, 0xfb, 0x3a,
	0x1e, 0x43, 0x70, 0x46, 0x7a, 0xc2, 0xb2, 0xc5, 0xa7, 0x9b, 0x45, 0x66, 0x51, 0x87, 0x2f, 0xb6,
	0x8d, 0xdf, 0xe2, 0x09, 0x84, 0xdf, 0x55, 0xa1, 0xa9, 0x1b, 0x6b, 0xab, 0x53, 0xdb, 0x46, 0xf9,
	0xe9, 0x9b, 0x6f, 0xce, 0xd1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x18, 0x44, 0x4d, 0x8a,
	0x04, 0x00, 0x00,
}
