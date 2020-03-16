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
	Line                 uint32   `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
	Service              string   `protobuf:"bytes,6,opt,name=service,proto3" json:"service,omitempty"`
	IpAddr               string   `protobuf:"bytes,7,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	Msg                  string   `protobuf:"bytes,8,opt,name=msg,proto3" json:"msg,omitempty"`
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

func (m *LogContent) GetLine() uint32 {
	if m != nil {
		return m.Line
	}
	return 0
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
	LevelIn              []string     `protobuf:"bytes,5,rep,name=levelIn,proto3" json:"levelIn,omitempty"`
	FuncIn               []string     `protobuf:"bytes,6,rep,name=funcIn,proto3" json:"funcIn,omitempty"`
	FileIn               []string     `protobuf:"bytes,7,rep,name=fileIn,proto3" json:"fileIn,omitempty"`
	LineIn               []uint32     `protobuf:"varint,8,rep,packed,name=lineIn,proto3" json:"lineIn,omitempty"`
	ServiceIn            []string     `protobuf:"bytes,9,rep,name=serviceIn,proto3" json:"serviceIn,omitempty"`
	IpIn                 []string     `protobuf:"bytes,10,rep,name=ipIn,proto3" json:"ipIn,omitempty"`
	MsgIn                []string     `protobuf:"bytes,11,rep,name=msgIn,proto3" json:"msgIn,omitempty"`
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

func (m *LogFilter) GetLevelIn() []string {
	if m != nil {
		return m.LevelIn
	}
	return nil
}

func (m *LogFilter) GetFuncIn() []string {
	if m != nil {
		return m.FuncIn
	}
	return nil
}

func (m *LogFilter) GetFileIn() []string {
	if m != nil {
		return m.FileIn
	}
	return nil
}

func (m *LogFilter) GetLineIn() []uint32 {
	if m != nil {
		return m.LineIn
	}
	return nil
}

func (m *LogFilter) GetServiceIn() []string {
	if m != nil {
		return m.ServiceIn
	}
	return nil
}

func (m *LogFilter) GetIpIn() []string {
	if m != nil {
		return m.IpIn
	}
	return nil
}

func (m *LogFilter) GetMsgIn() []string {
	if m != nil {
		return m.MsgIn
	}
	return nil
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
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0x27, 0x77, 0x4d, 0x72, 0x37, 0x69, 0xa5, 0xac, 0xa2, 0xdb, 0x58, 0xe5, 0xb8, 0x07, 0x09,
	0x88, 0x11, 0xda, 0x77, 0xb1, 0xa4, 0xb4, 0x1c, 0x06, 0xd1, 0xa3, 0xe2, 0xa3, 0x9c, 0xcd, 0x64,
	0x59, 0xdc, 0xdb, 0xbd, 0xee, 0x6d, 0x02, 0x7e, 0x11, 0xbf, 0x91, 0x7e, 0x2e, 0xd9, 0xc9, 0x5e,
	0x5b, 0xdb, 0xda, 0x3e, 0x04, 0xe6, 0xf7, 0x67, 0x77, 0xd8, 0xdf, 0xcc, 0x05, 0x1e, 0x37, 0xd6,
	0x38, 0xf3, 0x56, 0x19, 0xe1, 0x7f, 0x53, 0x42, 0x6c, 0x57, 0x98, 0x69, 0x2d, 0xcf, 0xad, 0x99,
	0xb6, 0x76, 0x3d, 0x55, 0x46, 0xe4, 0x6b, 0x18, 0xcc, 0x96, 0xa2, 0xc4, 0x0b, 0x96, 0xc1, 0x48,
	0x19, 0xf1, 0x01, 0xb1, 0x39, 0x93, 0x35, 0xf2, 0x5e, 0xd6, 0x9b, 0xec, 0x94, 0xd7, 0x29, 0xf6,
	0x0a, 0x1e, 0x2d, 0xa5, 0x42, 0x5d, 0xd5, 0x78, 0x62, 0x6c, 0x5d, 0x39, 0x1e, 0x65, 0xbd, 0x49,
	0x5a, 0xde, 0x60, 0x59, 0x0e, 0xdb, 0x16, 0x1b, 0x63, 0xdd, 0xac, 0x52, 0x0a, 0x2d, 0x8f, 0xb3,
	0xde, 0x24, 0x29, 0xff, 0xe1, 0xf2, 0x3f, 0x3d, 0x80, 0xb9, 0x11, 0x33, 0xa3, 0x1d, 0x6a, 0xc7,
	0x18, 0x6c, 0xb9, 0xae, 0x6b, 0x5a, 0x52, 0xcd, 0x9e, 0x40, 0x5f, 0xe1, 0x1a, 0x55, 0xe8, 0xb2,
	0x01, 0x6c, 0x0c, 0xc9, 0x72, 0xa5, 0xcf, 0x3f, 0x56, 0x35, 0xd2, 0xc5, 0x69, 0x79, 0x89, 0x49,
	0x93, 0x0a, 0x49, 0xdb, 0x0a, 0x5a, 0xc0, 0xbe, 0x83, 0x92, 0x1a, 0x79, 0x9f, 0xde, 0x45, 0x35,
	0xe3, 0x30, 0x6c, 0xd1, 0xae, 0xe5, 0x39, 0xf2, 0x01, 0xd9, 0x3b, 0xc8, 0x9e, 0xc2, 0x40, 0x36,
	0x47, 0x8b, 0x85, 0xe5, 0x43, 0x12, 0x02, 0x62, 0xbb, 0x10, 0xd7, 0xad, 0xe0, 0x09, 0x91, 0xbe,
	0xcc, 0x53, 0x18, 0x96, 0x78, 0xb1, 0xc2, 0xd6, 0xe5, 0x00, 0x49, 0x89, 0x6d, 0x63, 0x74, 0x8b,
	0xf9, 0x0f, 0xe8, 0x7f, 0x5e, 0xa1, 0xfd, 0xe9, 0xfb, 0x36, 0x95, 0xe8, 0xf2, 0xa4, 0x9a, 0xed,
	0x41, 0xd2, 0xa0, 0xfd, 0x46, 0x7c, 0x44, 0xfc, 0xb0, 0x41, 0xfb, 0xc9, 0x4b, 0x87, 0x30, 0x58,
	0x4a, 0xe5, 0x42, 0x6a, 0xa3, 0x83, 0xe7, 0xd3, 0x9b, 0x23, 0x9b, 0xce, 0x8d, 0x38, 0x21, 0x4b,
	0x19, 0xac, 0xf9, 0xef, 0x08, 0xd2, 0x4b, 0x96, 0xbd, 0x81, 0xb8, 0xd2, 0x0b, 0xde, 0xcb, 0xe2,
	0x87, 0xce, 0x7b, 0x1f, 0x7b, 0x0d, 0x91, 0xb1, 0x3c, 0x7a, 0xd8, 0x1d, 0x19, 0x4b, 0x09, 0x5b,
	0x53, 0xd3, 0x86, 0x74, 0xe9, 0x07, 0xec, 0x33, 0x73, 0x86, 0x94, 0x4d, 0xf6, 0x01, 0xf9, 0x94,
	0x69, 0x74, 0x85, 0xe6, 0xfd, 0x2c, 0xf6, 0x29, 0x07, 0xe8, 0x4f, 0xf8, 0xd9, 0x15, 0x9a, 0x0f,
	0x48, 0x08, 0x88, 0x78, 0xa9, 0xb0, 0xd0, 0x7c, 0x18, 0x78, 0x42, 0x9e, 0xf7, 0x73, 0x2b, 0x34,
	0x4f, 0xb2, 0x78, 0xb2, 0x53, 0x06, 0xc4, 0xf6, 0x21, 0x0d, 0x83, 0x2b, 0x34, 0x4f, 0xe9, 0xc8,
	0x15, 0xe1, 0x27, 0x20, 0x9b, 0x42, 0x73, 0x20, 0x81, 0x6a, 0xbf, 0x5b, 0x75, 0x2b, 0x0a, 0xcd,
	0x47, 0x44, 0x6e, 0x40, 0x5e, 0xc3, 0xe8, 0x6a, 0x27, 0x5b, 0xf6, 0x8e, 0xbe, 0x88, 0x0e, 0x86,
	0x40, 0xf7, 0xef, 0x8c, 0x28, 0x98, 0xca, 0xeb, 0x07, 0xd8, 0x4b, 0x80, 0x33, 0xe3, 0x2a, 0x35,
	0x33, 0x2b, 0xed, 0xc2, 0xa0, 0xaf, 0x31, 0x07, 0xbf, 0x22, 0x88, 0xe7, 0x46, 0xb0, 0x63, 0xd8,
	0xfe, 0xd2, 0x2c, 0x2a, 0x87, 0x33, 0xa3, 0x97, 0x52, 0x30, 0x7e, 0xbb, 0xc5, 0xe6, 0x1b, 0x1d,
	0x8f, 0x6f, 0x2b, 0xdd, 0xc6, 0xb1, 0xf7, 0x90, 0x9e, 0xa2, 0x0b, 0x57, 0xec, 0xdd, 0x65, 0xa4,
	0x2d, 0x1d, 0xff, 0xf7, 0x76, 0x76, 0x04, 0xc3, 0x53, 0x74, 0x73, 0x23, 0x5a, 0xf6, 0xec, 0xb6,
	0x89, 0xd6, 0x79, 0xfc, 0xe2, 0xbe, 0xe7, 0xb7, 0xec, 0x18, 0x92, 0xaf, 0x56, 0x3a, 0xf4, 0xcf,
	0xba, 0x37, 0xa9, 0xfb, 0x9e, 0xf2, 0x7d, 0x40, 0xff, 0x56, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x2d, 0xaa, 0x7d, 0x12, 0xc4, 0x04, 0x00, 0x00,
}
