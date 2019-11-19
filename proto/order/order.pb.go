// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order/order.proto

package go_micro_srv_agv_order

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

type UpdateReq struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	ParaNo               uint32   `protobuf:"varint,2,opt,name=paraNo,proto3" json:"paraNo,omitempty"`
	Value                uint32   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{0}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *UpdateReq) GetParaNo() uint32 {
	if m != nil {
		return m.ParaNo
	}
	return 0
}

func (m *UpdateReq) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type CancelReq struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelReq) Reset()         { *m = CancelReq{} }
func (m *CancelReq) String() string { return proto.CompactTextString(m) }
func (*CancelReq) ProtoMessage()    {}
func (*CancelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{1}
}

func (m *CancelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelReq.Unmarshal(m, b)
}
func (m *CancelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelReq.Marshal(b, m, deterministic)
}
func (m *CancelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelReq.Merge(m, src)
}
func (m *CancelReq) XXX_Size() int {
	return xxx_messageInfo_CancelReq.Size(m)
}
func (m *CancelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelReq.DiscardUnknown(m)
}

var xxx_messageInfo_CancelReq proto.InternalMessageInfo

func (m *CancelReq) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type Response struct {
	Res                  bool     `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{2}
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

func (m *Response) GetRes() bool {
	if m != nil {
		return m.Res
	}
	return false
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Task struct {
	TsId                 uint32   `protobuf:"varint,1,opt,name=tsId,proto3" json:"tsId,omitempty"`
	Priority             uint32   `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Ikey                 uint32   `protobuf:"varint,3,opt,name=ikey,proto3" json:"ikey,omitempty"`
	Params               []int32  `protobuf:"varint,4,rep,packed,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{3}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTsId() uint32 {
	if m != nil {
		return m.TsId
	}
	return 0
}

func (m *Task) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Task) GetIkey() uint32 {
	if m != nil {
		return m.Ikey
	}
	return 0
}

func (m *Task) GetParams() []int32 {
	if m != nil {
		return m.Params
	}
	return nil
}

type QBMessage struct {
	TaskID               uint32   `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	TsID                 uint32   `protobuf:"varint,2,opt,name=tsID,proto3" json:"tsID,omitempty"`
	Priority             uint32   `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Code                 uint32   `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Ikey                 uint32   `protobuf:"varint,5,opt,name=ikey,proto3" json:"ikey,omitempty"`
	Params               []uint32 `protobuf:"varint,6,rep,packed,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QBMessage) Reset()         { *m = QBMessage{} }
func (m *QBMessage) String() string { return proto.CompactTextString(m) }
func (*QBMessage) ProtoMessage()    {}
func (*QBMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{4}
}

func (m *QBMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QBMessage.Unmarshal(m, b)
}
func (m *QBMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QBMessage.Marshal(b, m, deterministic)
}
func (m *QBMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QBMessage.Merge(m, src)
}
func (m *QBMessage) XXX_Size() int {
	return xxx_messageInfo_QBMessage.Size(m)
}
func (m *QBMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QBMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QBMessage proto.InternalMessageInfo

func (m *QBMessage) GetTaskID() uint32 {
	if m != nil {
		return m.TaskID
	}
	return 0
}

func (m *QBMessage) GetTsID() uint32 {
	if m != nil {
		return m.TsID
	}
	return 0
}

func (m *QBMessage) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *QBMessage) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QBMessage) GetIkey() uint32 {
	if m != nil {
		return m.Ikey
	}
	return 0
}

func (m *QBMessage) GetParams() []uint32 {
	if m != nil {
		return m.Params
	}
	return nil
}

type MMessage struct {
	TaskID               uint32   `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Function             uint32   `protobuf:"varint,3,opt,name=function,proto3" json:"function,omitempty"`
	ParNo                uint32   `protobuf:"varint,4,opt,name=parNo,proto3" json:"parNo,omitempty"`
	Params               []uint32 `protobuf:"varint,6,rep,packed,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MMessage) Reset()         { *m = MMessage{} }
func (m *MMessage) String() string { return proto.CompactTextString(m) }
func (*MMessage) ProtoMessage()    {}
func (*MMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_986e030a471601a2, []int{5}
}

func (m *MMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MMessage.Unmarshal(m, b)
}
func (m *MMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MMessage.Marshal(b, m, deterministic)
}
func (m *MMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MMessage.Merge(m, src)
}
func (m *MMessage) XXX_Size() int {
	return xxx_messageInfo_MMessage.Size(m)
}
func (m *MMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MMessage proto.InternalMessageInfo

func (m *MMessage) GetTaskID() uint32 {
	if m != nil {
		return m.TaskID
	}
	return 0
}

func (m *MMessage) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MMessage) GetFunction() uint32 {
	if m != nil {
		return m.Function
	}
	return 0
}

func (m *MMessage) GetParNo() uint32 {
	if m != nil {
		return m.ParNo
	}
	return 0
}

func (m *MMessage) GetParams() []uint32 {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateReq)(nil), "go.micro.srv.agv.order.UpdateReq")
	proto.RegisterType((*CancelReq)(nil), "go.micro.srv.agv.order.CancelReq")
	proto.RegisterType((*Response)(nil), "go.micro.srv.agv.order.Response")
	proto.RegisterType((*Task)(nil), "go.micro.srv.agv.order.Task")
	proto.RegisterType((*QBMessage)(nil), "go.micro.srv.agv.order.QBMessage")
	proto.RegisterType((*MMessage)(nil), "go.micro.srv.agv.order.MMessage")
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor_986e030a471601a2) }

var fileDescriptor_986e030a471601a2 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x6e, 0xe2, 0x30,
	0x14, 0x55, 0x48, 0x02, 0xc9, 0xd5, 0x20, 0x8d, 0xac, 0x11, 0x13, 0xa1, 0x59, 0x84, 0xac, 0x58,
	0x65, 0xa4, 0x19, 0x69, 0x3e, 0x60, 0xca, 0x86, 0x05, 0x20, 0x52, 0xfa, 0x01, 0x26, 0x71, 0x51,
	0x04, 0x89, 0x53, 0x3b, 0xd0, 0xb2, 0xed, 0x2f, 0xf4, 0x0f, 0xfa, 0xa5, 0x95, 0x1f, 0x71, 0x5b,
	0x95, 0xb4, 0xd9, 0xa0, 0x7b, 0x7c, 0x8f, 0xcf, 0x3d, 0xf8, 0x9e, 0xc0, 0xcf, 0x8a, 0xd1, 0x9a,
	0xfe, 0xa6, 0x2c, 0x23, 0x4c, 0xfd, 0xc6, 0xf2, 0x04, 0x8d, 0x76, 0x34, 0x2e, 0xf2, 0x94, 0xd1,
	0x98, 0xb3, 0x53, 0x8c, 0x77, 0xa7, 0x58, 0x76, 0xa3, 0x15, 0xf8, 0x37, 0x55, 0x86, 0x6b, 0x92,
	0x90, 0x3b, 0xf4, 0x03, 0xdc, 0xbc, 0xcc, 0xc8, 0x43, 0x60, 0x85, 0xd6, 0x74, 0x98, 0x28, 0x80,
	0x46, 0xd0, 0xaf, 0x30, 0xc3, 0x4b, 0x1a, 0xf4, 0xe4, 0xb1, 0x46, 0x82, 0x7d, 0xc2, 0x87, 0x23,
	0x09, 0x6c, 0xc5, 0x96, 0x20, 0x9a, 0x80, 0x7f, 0x85, 0xcb, 0x94, 0x1c, 0x5a, 0x05, 0xa3, 0x7f,
	0xe0, 0x25, 0x84, 0x57, 0xb4, 0xe4, 0x04, 0x7d, 0x07, 0x9b, 0x11, 0x2e, 0xfb, 0x5e, 0x22, 0x4a,
	0x14, 0xc0, 0xa0, 0x20, 0x9c, 0xe3, 0x1d, 0x91, 0xf3, 0xfc, 0xa4, 0x81, 0xd1, 0x16, 0x9c, 0x0d,
	0xe6, 0x7b, 0x84, 0xc0, 0xa9, 0xf9, 0x3c, 0xd3, 0xa2, 0xb2, 0x46, 0x63, 0xf0, 0x2a, 0x96, 0x53,
	0x96, 0xd7, 0x67, 0x6d, 0xd3, 0x60, 0xc1, 0xcf, 0xf7, 0xe4, 0xac, 0x7d, 0xca, 0xba, 0xf9, 0x53,
	0x05, 0x0f, 0x9c, 0xd0, 0x9e, 0xba, 0x89, 0x46, 0xd1, 0x93, 0x05, 0xfe, 0xfa, 0xff, 0x42, 0x4d,
	0x14, 0xac, 0x1a, 0xf3, 0xfd, 0x7c, 0xa6, 0x67, 0x69, 0xa4, 0x1d, 0xcc, 0xf4, 0x24, 0x59, 0xbf,
	0x73, 0x60, 0x7f, 0x74, 0x90, 0xd2, 0x8c, 0x04, 0x8e, 0xe2, 0x8b, 0xda, 0xb8, 0x72, 0x2f, 0xba,
	0xea, 0x87, 0x76, 0xf3, 0xd4, 0x05, 0x8f, 0x1e, 0x2d, 0xf0, 0x16, 0x5f, 0x99, 0x32, 0x8f, 0xdd,
	0x7b, 0xbb, 0xbd, 0x31, 0x78, 0xb7, 0xc7, 0x32, 0xad, 0x73, 0x5a, 0x36, 0xb6, 0x1a, 0x2c, 0x6e,
	0x54, 0x98, 0x2d, 0xa9, 0xf6, 0xa5, 0x40, 0x9b, 0x89, 0x3f, 0xcf, 0x36, 0xb8, 0x2b, 0x11, 0x1a,
	0x34, 0x87, 0xc1, 0x92, 0xdc, 0xcb, 0x5d, 0xfc, 0x8a, 0x2f, 0x07, 0x2b, 0x16, 0xdd, 0x71, 0xd8,
	0xd6, 0x35, 0xfb, 0x5f, 0x03, 0xa8, 0xb8, 0x48, 0xb5, 0x49, 0x1b, 0xdf, 0x44, 0xaa, 0x9b, 0xa4,
	0x8a, 0xf4, 0xe7, 0x92, 0x26, 0xf6, 0x1d, 0x24, 0x37, 0x30, 0xbc, 0x26, 0x65, 0xf6, 0x1a, 0x8c,
	0x56, 0x55, 0x43, 0xe9, 0xa0, 0x9a, 0xc0, 0x37, 0xa1, 0x6a, 0x16, 0xdb, 0x7a, 0x63, 0xd1, 0x59,
	0x73, 0xdb, 0x97, 0x9f, 0xfb, 0xdf, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x96, 0x64, 0x3b,
	0x09, 0x04, 0x00, 0x00,
}
