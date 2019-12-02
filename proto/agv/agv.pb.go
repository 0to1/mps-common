// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/agv/agv.proto

package go_micro_srv_agv

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

type Query struct {
	Page                 uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{0}
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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{1}
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
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{2}
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

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AgvMessage struct {
	AgvID                uint32   `protobuf:"varint,2,opt,name=agvID,proto3" json:"agvID,omitempty"`
	X                    uint32   `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint32   `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
	Angle                uint32   `protobuf:"varint,5,opt,name=angle,proto3" json:"angle,omitempty"`
	Point                uint32   `protobuf:"varint,6,opt,name=point,proto3" json:"point,omitempty"`
	Segment              uint32   `protobuf:"varint,7,opt,name=segment,proto3" json:"segment,omitempty"`
	Type                 uint32   `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	Battery              uint32   `protobuf:"varint,9,opt,name=battery,proto3" json:"battery,omitempty"`
	Res                  bool     `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgvMessage) Reset()         { *m = AgvMessage{} }
func (m *AgvMessage) String() string { return proto.CompactTextString(m) }
func (*AgvMessage) ProtoMessage()    {}
func (*AgvMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{3}
}

func (m *AgvMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgvMessage.Unmarshal(m, b)
}
func (m *AgvMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgvMessage.Marshal(b, m, deterministic)
}
func (m *AgvMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgvMessage.Merge(m, src)
}
func (m *AgvMessage) XXX_Size() int {
	return xxx_messageInfo_AgvMessage.Size(m)
}
func (m *AgvMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgvMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AgvMessage proto.InternalMessageInfo

func (m *AgvMessage) GetAgvID() uint32 {
	if m != nil {
		return m.AgvID
	}
	return 0
}

func (m *AgvMessage) GetX() uint32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *AgvMessage) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *AgvMessage) GetAngle() uint32 {
	if m != nil {
		return m.Angle
	}
	return 0
}

func (m *AgvMessage) GetPoint() uint32 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *AgvMessage) GetSegment() uint32 {
	if m != nil {
		return m.Segment
	}
	return 0
}

func (m *AgvMessage) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AgvMessage) GetBattery() uint32 {
	if m != nil {
		return m.Battery
	}
	return 0
}

func (m *AgvMessage) GetRes() bool {
	if m != nil {
		return m.Res
	}
	return false
}

type AgvReq struct {
	AgvID                uint32   `protobuf:"varint,1,opt,name=agvID,proto3" json:"agvID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgvReq) Reset()         { *m = AgvReq{} }
func (m *AgvReq) String() string { return proto.CompactTextString(m) }
func (*AgvReq) ProtoMessage()    {}
func (*AgvReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{4}
}

func (m *AgvReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgvReq.Unmarshal(m, b)
}
func (m *AgvReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgvReq.Marshal(b, m, deterministic)
}
func (m *AgvReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgvReq.Merge(m, src)
}
func (m *AgvReq) XXX_Size() int {
	return xxx_messageInfo_AgvReq.Size(m)
}
func (m *AgvReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AgvReq.DiscardUnknown(m)
}

var xxx_messageInfo_AgvReq proto.InternalMessageInfo

func (m *AgvReq) GetAgvID() uint32 {
	if m != nil {
		return m.AgvID
	}
	return 0
}

type AgvsResponse struct {
	Agvs                 []*AgvMessage `protobuf:"bytes,1,rep,name=agvs,proto3" json:"agvs,omitempty"`
	Totalpages           uint32        `protobuf:"varint,2,opt,name=totalpages,proto3" json:"totalpages,omitempty"`
	Currpages            uint32        `protobuf:"varint,3,opt,name=currpages,proto3" json:"currpages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AgvsResponse) Reset()         { *m = AgvsResponse{} }
func (m *AgvsResponse) String() string { return proto.CompactTextString(m) }
func (*AgvsResponse) ProtoMessage()    {}
func (*AgvsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{5}
}

func (m *AgvsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgvsResponse.Unmarshal(m, b)
}
func (m *AgvsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgvsResponse.Marshal(b, m, deterministic)
}
func (m *AgvsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgvsResponse.Merge(m, src)
}
func (m *AgvsResponse) XXX_Size() int {
	return xxx_messageInfo_AgvsResponse.Size(m)
}
func (m *AgvsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AgvsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AgvsResponse proto.InternalMessageInfo

func (m *AgvsResponse) GetAgvs() []*AgvMessage {
	if m != nil {
		return m.Agvs
	}
	return nil
}

func (m *AgvsResponse) GetTotalpages() uint32 {
	if m != nil {
		return m.Totalpages
	}
	return 0
}

func (m *AgvsResponse) GetCurrpages() uint32 {
	if m != nil {
		return m.Currpages
	}
	return 0
}

func init() {
	proto.RegisterType((*Query)(nil), "go.micro.srv.agv.Query")
	proto.RegisterType((*Request)(nil), "go.micro.srv.agv.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.agv.Response")
	proto.RegisterType((*AgvMessage)(nil), "go.micro.srv.agv.AgvMessage")
	proto.RegisterType((*AgvReq)(nil), "go.micro.srv.agv.AgvReq")
	proto.RegisterType((*AgvsResponse)(nil), "go.micro.srv.agv.AgvsResponse")
}

func init() { proto.RegisterFile("proto/agv/agv.proto", fileDescriptor_b9b7acc5ecfca92c) }

var fileDescriptor_b9b7acc5ecfca92c = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0xaa, 0xda, 0x30,
	0x18, 0xc7, 0xa9, 0x55, 0xdb, 0x7e, 0x3a, 0x90, 0x6c, 0xb0, 0x4c, 0x44, 0x46, 0xd9, 0x85, 0x57,
	0xdd, 0x70, 0xec, 0x01, 0xea, 0x84, 0xe1, 0xc5, 0x60, 0xcb, 0x9e, 0xa0, 0xca, 0x47, 0x10, 0xb4,
	0xa9, 0x49, 0x0c, 0xf6, 0x66, 0xaf, 0x71, 0x5e, 0xe8, 0x3c, 0xd8, 0x21, 0x49, 0xb5, 0x72, 0x14,
	0xcf, 0x45, 0x21, 0xbf, 0x7f, 0xbe, 0xfc, 0xfb, 0x7d, 0xff, 0xa6, 0xf0, 0xbe, 0x92, 0x42, 0x8b,
	0xaf, 0x05, 0x37, 0xf6, 0xc9, 0x1c, 0x91, 0x11, 0x17, 0xd9, 0x7e, 0xbb, 0x91, 0x22, 0x53, 0xd2,
	0x64, 0x05, 0x37, 0xe9, 0x0f, 0xe8, 0xfd, 0x3d, 0xa2, 0xac, 0x09, 0x81, 0x6e, 0x55, 0x70, 0xa4,
	0xc1, 0xe7, 0x60, 0xf6, 0x8e, 0xb9, 0x35, 0xa1, 0x10, 0x55, 0x28, 0xff, 0x58, 0xb9, 0xe3, 0xe4,
	0x33, 0xa6, 0x09, 0x44, 0x0c, 0x0f, 0x47, 0x54, 0x3a, 0xfd, 0x02, 0x31, 0x43, 0x55, 0x89, 0x52,
	0xb9, 0x03, 0x7b, 0x54, 0xea, 0xec, 0x93, 0xb0, 0x33, 0xa6, 0xcf, 0x01, 0x40, 0xce, 0xcd, 0x6f,
	0x8f, 0xe4, 0x03, 0xf4, 0x0a, 0x6e, 0x56, 0xcb, 0xc6, 0xd7, 0x03, 0x19, 0x42, 0x70, 0xa2, 0xa1,
	0x53, 0x82, 0x93, 0xa5, 0x9a, 0x76, 0x3d, 0xd5, 0xee, 0x44, 0xc9, 0x77, 0x48, 0x7b, 0xcd, 0x09,
	0x0b, 0x56, 0xad, 0xc4, 0xb6, 0xd4, 0xb4, 0xef, 0x55, 0x07, 0xb6, 0x0d, 0x85, 0x7c, 0x8f, 0xa5,
	0xa6, 0x91, 0xef, 0xbb, 0x41, 0x3b, 0xa5, 0xae, 0x2b, 0xa4, 0xb1, 0x9f, 0xd2, 0xae, 0x6d, 0xf5,
	0xba, 0xd0, 0x1a, 0x65, 0x4d, 0x13, 0x5f, 0xdd, 0x20, 0x19, 0x41, 0x28, 0x51, 0xb9, 0x51, 0x62,
	0x66, 0x97, 0xe9, 0x14, 0xfa, 0x39, 0x37, 0x0c, 0x0f, 0xed, 0x04, 0xc1, 0xd5, 0x04, 0xe9, 0x7f,
	0x18, 0xe6, 0xdc, 0xa8, 0x4b, 0x20, 0xdf, 0xa0, 0x5b, 0x70, 0x63, 0x2d, 0xc2, 0xd9, 0x60, 0x3e,
	0xc9, 0x5e, 0xe7, 0x9f, 0xb5, 0x99, 0x30, 0x57, 0x49, 0xa6, 0x00, 0x5a, 0xe8, 0x62, 0x67, 0x3f,
	0x80, 0x6a, 0xe2, 0xb9, 0x52, 0xc8, 0x04, 0x92, 0xcd, 0x51, 0x4a, 0xbf, 0xed, 0xb3, 0x6a, 0x85,
	0xf9, 0x53, 0x07, 0xc2, 0x9c, 0x1b, 0xb2, 0x04, 0xf8, 0x85, 0x3a, 0xe7, 0x66, 0x51, 0xaf, 0x96,
	0x84, 0xde, 0x7d, 0x2f, 0xc3, 0xc3, 0xf8, 0x61, 0x47, 0x64, 0x01, 0x91, 0x77, 0x51, 0xe4, 0xe3,
	0x6d, 0xa1, 0xbb, 0x37, 0xe3, 0xe9, 0x5d, 0x87, 0x36, 0x81, 0x9f, 0x30, 0xf8, 0xa7, 0x45, 0xf5,
	0x76, 0x2b, 0xe3, 0xdb, 0x9d, 0x8b, 0x49, 0x0e, 0x71, 0x63, 0xa2, 0xc8, 0xa7, 0x7b, 0x75, 0xee,
	0x2a, 0x3e, 0xb2, 0x58, 0xf7, 0xdd, 0x1f, 0xf0, 0xfd, 0x25, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xf7,
	0xf3, 0x30, 0x18, 0x03, 0x00, 0x00,
}
