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

type StatusMsg struct {
	StatusID             uint32   `protobuf:"varint,1,opt,name=statusID,proto3" json:"statusID,omitempty"`
	StatusText           string   `protobuf:"bytes,2,opt,name=statusText,proto3" json:"statusText,omitempty"`
	StatusColor          string   `protobuf:"bytes,3,opt,name=statusColor,proto3" json:"statusColor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusMsg) Reset()         { *m = StatusMsg{} }
func (m *StatusMsg) String() string { return proto.CompactTextString(m) }
func (*StatusMsg) ProtoMessage()    {}
func (*StatusMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{0}
}

func (m *StatusMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMsg.Unmarshal(m, b)
}
func (m *StatusMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMsg.Marshal(b, m, deterministic)
}
func (m *StatusMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMsg.Merge(m, src)
}
func (m *StatusMsg) XXX_Size() int {
	return xxx_messageInfo_StatusMsg.Size(m)
}
func (m *StatusMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMsg proto.InternalMessageInfo

func (m *StatusMsg) GetStatusID() uint32 {
	if m != nil {
		return m.StatusID
	}
	return 0
}

func (m *StatusMsg) GetStatusText() string {
	if m != nil {
		return m.StatusText
	}
	return ""
}

func (m *StatusMsg) GetStatusColor() string {
	if m != nil {
		return m.StatusColor
	}
	return ""
}

type StatusMsgs struct {
	Msgs                 []*StatusMsg `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StatusMsgs) Reset()         { *m = StatusMsgs{} }
func (m *StatusMsgs) String() string { return proto.CompactTextString(m) }
func (*StatusMsgs) ProtoMessage()    {}
func (*StatusMsgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{1}
}

func (m *StatusMsgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMsgs.Unmarshal(m, b)
}
func (m *StatusMsgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMsgs.Marshal(b, m, deterministic)
}
func (m *StatusMsgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMsgs.Merge(m, src)
}
func (m *StatusMsgs) XXX_Size() int {
	return xxx_messageInfo_StatusMsgs.Size(m)
}
func (m *StatusMsgs) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMsgs.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMsgs proto.InternalMessageInfo

func (m *StatusMsgs) GetMsgs() []*StatusMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type IDRequest struct {
	StatusID             uint32   `protobuf:"varint,1,opt,name=statusID,proto3" json:"statusID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{2}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetStatusID() uint32 {
	if m != nil {
		return m.StatusID
	}
	return 0
}

type Query struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{3}
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

func (m *Query) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

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
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{4}
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

func (m *Request) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

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
	return fileDescriptor_b9b7acc5ecfca92c, []int{5}
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
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	X                    uint32   `protobuf:"varint,4,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint32   `protobuf:"varint,5,opt,name=y,proto3" json:"y,omitempty"`
	Angle                uint32   `protobuf:"varint,6,opt,name=angle,proto3" json:"angle,omitempty"`
	Point                uint32   `protobuf:"varint,7,opt,name=point,proto3" json:"point,omitempty"`
	Segment              uint32   `protobuf:"varint,8,opt,name=segment,proto3" json:"segment,omitempty"`
	Battery              uint32   `protobuf:"varint,9,opt,name=battery,proto3" json:"battery,omitempty"`
	Plc                  uint32   `protobuf:"varint,10,opt,name=plc,proto3" json:"plc,omitempty"`
	Status               uint32   `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgvMessage) Reset()         { *m = AgvMessage{} }
func (m *AgvMessage) String() string { return proto.CompactTextString(m) }
func (*AgvMessage) ProtoMessage()    {}
func (*AgvMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{6}
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

func (m *AgvMessage) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *AgvMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgvMessage) GetType() uint32 {
	if m != nil {
		return m.Type
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

func (m *AgvMessage) GetBattery() uint32 {
	if m != nil {
		return m.Battery
	}
	return 0
}

func (m *AgvMessage) GetPlc() uint32 {
	if m != nil {
		return m.Plc
	}
	return 0
}

func (m *AgvMessage) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type AgvReq struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	AgvID                uint32   `protobuf:"varint,2,opt,name=agvID,proto3" json:"agvID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgvReq) Reset()         { *m = AgvReq{} }
func (m *AgvReq) String() string { return proto.CompactTextString(m) }
func (*AgvReq) ProtoMessage()    {}
func (*AgvReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{7}
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

func (m *AgvReq) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *AgvReq) GetAgvID() uint32 {
	if m != nil {
		return m.AgvID
	}
	return 0
}

type AgvsResponse struct {
	Agvs                 []*AgvMessage `protobuf:"bytes,1,rep,name=agvs,proto3" json:"agvs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AgvsResponse) Reset()         { *m = AgvsResponse{} }
func (m *AgvsResponse) String() string { return proto.CompactTextString(m) }
func (*AgvsResponse) ProtoMessage()    {}
func (*AgvsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{8}
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

type QueryError struct {
	Page                 uint32     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32     `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Filter               *ErrFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryError) Reset()         { *m = QueryError{} }
func (m *QueryError) String() string { return proto.CompactTextString(m) }
func (*QueryError) ProtoMessage()    {}
func (*QueryError) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{9}
}

func (m *QueryError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryError.Unmarshal(m, b)
}
func (m *QueryError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryError.Marshal(b, m, deterministic)
}
func (m *QueryError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryError.Merge(m, src)
}
func (m *QueryError) XXX_Size() int {
	return xxx_messageInfo_QueryError.Size(m)
}
func (m *QueryError) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryError.DiscardUnknown(m)
}

var xxx_messageInfo_QueryError proto.InternalMessageInfo

func (m *QueryError) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *QueryError) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *QueryError) GetFilter() *ErrFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type ErrFilter struct {
	And                  []*ErrFilter `protobuf:"bytes,1,rep,name=and,proto3" json:"and,omitempty"`
	Or                   []*ErrFilter `protobuf:"bytes,2,rep,name=or,proto3" json:"or,omitempty"`
	FromTime             string       `protobuf:"bytes,3,opt,name=fromTime,proto3" json:"fromTime,omitempty"`
	ToTime               string       `protobuf:"bytes,4,opt,name=toTime,proto3" json:"toTime,omitempty"`
	GarageIn             []uint32     `protobuf:"varint,5,rep,packed,name=garageIn,proto3" json:"garageIn,omitempty"`
	AgvIn                []uint32     `protobuf:"varint,6,rep,packed,name=agvIn,proto3" json:"agvIn,omitempty"`
	ErrorIn              []uint32     `protobuf:"varint,7,rep,packed,name=errorIn,proto3" json:"errorIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ErrFilter) Reset()         { *m = ErrFilter{} }
func (m *ErrFilter) String() string { return proto.CompactTextString(m) }
func (*ErrFilter) ProtoMessage()    {}
func (*ErrFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{10}
}

func (m *ErrFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrFilter.Unmarshal(m, b)
}
func (m *ErrFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrFilter.Marshal(b, m, deterministic)
}
func (m *ErrFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrFilter.Merge(m, src)
}
func (m *ErrFilter) XXX_Size() int {
	return xxx_messageInfo_ErrFilter.Size(m)
}
func (m *ErrFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ErrFilter proto.InternalMessageInfo

func (m *ErrFilter) GetAnd() []*ErrFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *ErrFilter) GetOr() []*ErrFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *ErrFilter) GetFromTime() string {
	if m != nil {
		return m.FromTime
	}
	return ""
}

func (m *ErrFilter) GetToTime() string {
	if m != nil {
		return m.ToTime
	}
	return ""
}

func (m *ErrFilter) GetGarageIn() []uint32 {
	if m != nil {
		return m.GarageIn
	}
	return nil
}

func (m *ErrFilter) GetAgvIn() []uint32 {
	if m != nil {
		return m.AgvIn
	}
	return nil
}

func (m *ErrFilter) GetErrorIn() []uint32 {
	if m != nil {
		return m.ErrorIn
	}
	return nil
}

type AgvError struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	AgvID                uint32   `protobuf:"varint,2,opt,name=agvID,proto3" json:"agvID,omitempty"`
	ErrorCode            uint32   `protobuf:"varint,3,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	PointID              uint32   `protobuf:"varint,4,opt,name=pointID,proto3" json:"pointID,omitempty"`
	StartTime            string   `protobuf:"bytes,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              string   `protobuf:"bytes,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgvError) Reset()         { *m = AgvError{} }
func (m *AgvError) String() string { return proto.CompactTextString(m) }
func (*AgvError) ProtoMessage()    {}
func (*AgvError) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{11}
}

func (m *AgvError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgvError.Unmarshal(m, b)
}
func (m *AgvError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgvError.Marshal(b, m, deterministic)
}
func (m *AgvError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgvError.Merge(m, src)
}
func (m *AgvError) XXX_Size() int {
	return xxx_messageInfo_AgvError.Size(m)
}
func (m *AgvError) XXX_DiscardUnknown() {
	xxx_messageInfo_AgvError.DiscardUnknown(m)
}

var xxx_messageInfo_AgvError proto.InternalMessageInfo

func (m *AgvError) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *AgvError) GetAgvID() uint32 {
	if m != nil {
		return m.AgvID
	}
	return 0
}

func (m *AgvError) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *AgvError) GetPointID() uint32 {
	if m != nil {
		return m.PointID
	}
	return 0
}

func (m *AgvError) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *AgvError) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type AgvErrors struct {
	AgvError             []*AgvError `protobuf:"bytes,1,rep,name=agvError,proto3" json:"agvError,omitempty"`
	TotalCount           uint32      `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AgvErrors) Reset()         { *m = AgvErrors{} }
func (m *AgvErrors) String() string { return proto.CompactTextString(m) }
func (*AgvErrors) ProtoMessage()    {}
func (*AgvErrors) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9b7acc5ecfca92c, []int{12}
}

func (m *AgvErrors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgvErrors.Unmarshal(m, b)
}
func (m *AgvErrors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgvErrors.Marshal(b, m, deterministic)
}
func (m *AgvErrors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgvErrors.Merge(m, src)
}
func (m *AgvErrors) XXX_Size() int {
	return xxx_messageInfo_AgvErrors.Size(m)
}
func (m *AgvErrors) XXX_DiscardUnknown() {
	xxx_messageInfo_AgvErrors.DiscardUnknown(m)
}

var xxx_messageInfo_AgvErrors proto.InternalMessageInfo

func (m *AgvErrors) GetAgvError() []*AgvError {
	if m != nil {
		return m.AgvError
	}
	return nil
}

func (m *AgvErrors) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*StatusMsg)(nil), "go.micro.srv.agv.StatusMsg")
	proto.RegisterType((*StatusMsgs)(nil), "go.micro.srv.agv.StatusMsgs")
	proto.RegisterType((*IDRequest)(nil), "go.micro.srv.agv.IDRequest")
	proto.RegisterType((*Query)(nil), "go.micro.srv.agv.Query")
	proto.RegisterType((*Request)(nil), "go.micro.srv.agv.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.agv.Response")
	proto.RegisterType((*AgvMessage)(nil), "go.micro.srv.agv.AgvMessage")
	proto.RegisterType((*AgvReq)(nil), "go.micro.srv.agv.AgvReq")
	proto.RegisterType((*AgvsResponse)(nil), "go.micro.srv.agv.AgvsResponse")
	proto.RegisterType((*QueryError)(nil), "go.micro.srv.agv.QueryError")
	proto.RegisterType((*ErrFilter)(nil), "go.micro.srv.agv.ErrFilter")
	proto.RegisterType((*AgvError)(nil), "go.micro.srv.agv.AgvError")
	proto.RegisterType((*AgvErrors)(nil), "go.micro.srv.agv.AgvErrors")
}

func init() { proto.RegisterFile("proto/agv/agv.proto", fileDescriptor_b9b7acc5ecfca92c) }

var fileDescriptor_b9b7acc5ecfca92c = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6a, 0x23, 0x37,
	0x18, 0x65, 0xfc, 0xef, 0xcf, 0x49, 0x1b, 0xd4, 0xd2, 0x2a, 0x4e, 0x08, 0x66, 0x68, 0x69, 0xa0,
	0xd4, 0x29, 0x09, 0xf4, 0xa2, 0xb0, 0xb0, 0x8e, 0x9d, 0x0d, 0x26, 0x04, 0x36, 0x93, 0xec, 0xf5,
	0x32, 0x89, 0x15, 0x31, 0x60, 0x8f, 0x26, 0x92, 0x3c, 0xc4, 0xef, 0xb4, 0x6f, 0xb2, 0x0f, 0xb2,
	0xf7, 0xfb, 0x04, 0x8b, 0x3e, 0x69, 0x7e, 0xc8, 0x8e, 0x27, 0xec, 0x85, 0x41, 0xe7, 0xe8, 0xf8,
	0xd3, 0x91, 0xbe, 0xa3, 0x11, 0xfc, 0x92, 0x48, 0xa1, 0xc5, 0x49, 0xc8, 0x53, 0xf3, 0x1b, 0x23,
	0x22, 0x7b, 0x5c, 0x8c, 0x57, 0xd1, 0x83, 0x14, 0x63, 0x25, 0xd3, 0x71, 0xc8, 0x53, 0x3f, 0x82,
	0xfe, 0xad, 0x0e, 0xf5, 0x5a, 0x5d, 0x2b, 0x4e, 0x86, 0xd0, 0x53, 0x08, 0xe6, 0x33, 0xea, 0x8d,
	0xbc, 0xe3, 0xdd, 0x20, 0xc7, 0xe4, 0x08, 0xc0, 0x8e, 0xef, 0xd8, 0xb3, 0xa6, 0x8d, 0x91, 0x77,
	0xdc, 0x0f, 0x4a, 0x0c, 0x19, 0xc1, 0xc0, 0xa2, 0xa9, 0x58, 0x0a, 0x49, 0x9b, 0x28, 0x28, 0x53,
	0xfe, 0x1b, 0x80, 0x7c, 0x29, 0x45, 0x4e, 0xa0, 0xb5, 0x52, 0x5c, 0x51, 0x6f, 0xd4, 0x3c, 0x1e,
	0x9c, 0x1e, 0x8c, 0x5f, 0x3a, 0x1b, 0xe7, 0xda, 0x00, 0x85, 0xfe, 0x5f, 0xd0, 0x9f, 0xcf, 0x02,
	0xf6, 0xb4, 0x66, 0x4a, 0xd7, 0x39, 0xf5, 0x6f, 0xa0, 0x7d, 0xb3, 0x66, 0x72, 0x63, 0x44, 0x3c,
	0x94, 0x21, 0x67, 0x85, 0x28, 0xc3, 0x84, 0x40, 0x2b, 0x09, 0x39, 0xc3, 0x8d, 0xec, 0x06, 0x38,
	0x26, 0x14, 0xba, 0x09, 0x93, 0xef, 0x0d, 0xdd, 0x44, 0x3a, 0x83, 0xfe, 0x9f, 0xd0, 0x2d, 0xad,
	0xbc, 0xad, 0xa8, 0xff, 0x07, 0xf4, 0x02, 0xa6, 0x12, 0x11, 0x2b, 0x2c, 0xb6, 0x62, 0x4a, 0x99,
	0x62, 0x1e, 0x9e, 0x45, 0x06, 0xfd, 0xaf, 0x1e, 0xc0, 0x84, 0xa7, 0xd7, 0x16, 0xd6, 0xba, 0xfc,
	0x09, 0x1a, 0xd1, 0xc2, 0x79, 0x6c, 0x44, 0x0b, 0xe3, 0x5a, 0x6f, 0x92, 0xcc, 0x1e, 0x8e, 0xc9,
	0x0e, 0x78, 0xcf, 0xb4, 0x85, 0x84, 0xf7, 0x6c, 0xd0, 0x86, 0xb6, 0x2d, 0xda, 0x90, 0x5f, 0xa1,
	0x1d, 0xc6, 0x7c, 0xc9, 0x68, 0x07, 0x19, 0x0b, 0x0c, 0x9b, 0x88, 0x28, 0xd6, 0xb4, 0x6b, 0x59,
	0x04, 0xc6, 0xb0, 0x62, 0x7c, 0xc5, 0x62, 0x4d, 0x7b, 0x76, 0xf7, 0x0e, 0x9a, 0x99, 0xfb, 0x50,
	0x6b, 0x26, 0x37, 0xb4, 0x6f, 0x67, 0x1c, 0x24, 0x7b, 0xd0, 0x4c, 0x96, 0x0f, 0x14, 0x90, 0x35,
	0x43, 0xf2, 0x1b, 0x74, 0x6c, 0x23, 0xe8, 0x00, 0x49, 0x87, 0xfc, 0xff, 0xa1, 0x33, 0xe1, 0x69,
	0xc0, 0x9e, 0x6a, 0xf7, 0x6b, 0xfc, 0xf2, 0x74, 0x3e, 0x73, 0x5b, 0xb6, 0xc0, 0x7f, 0x0b, 0x3b,
	0x13, 0x9e, 0xaa, 0xfc, 0x68, 0xff, 0x85, 0x56, 0xc8, 0xd3, 0x2c, 0x3a, 0x87, 0xdf, 0x47, 0xa7,
	0x38, 0xdd, 0x00, 0x95, 0x7e, 0x02, 0x80, 0x91, 0xb8, 0x90, 0x52, 0xc8, 0xbc, 0xf7, 0x5e, 0xa9,
	0xf7, 0xfb, 0xd0, 0x4b, 0x98, 0xfc, 0x58, 0xca, 0x44, 0xd6, 0x7c, 0x72, 0x06, 0x9d, 0xc7, 0x68,
	0xa9, 0x99, 0x0d, 0x75, 0x65, 0x56, 0x2f, 0xa4, 0x7c, 0x87, 0x92, 0xc0, 0x49, 0xfd, 0x2f, 0x1e,
	0xf4, 0x73, 0x96, 0xfc, 0x03, 0xcd, 0x30, 0x5e, 0x6c, 0xcf, 0x7a, 0xf1, 0x7f, 0xa3, 0x23, 0x7f,
	0x43, 0x43, 0x48, 0xda, 0x78, 0x5d, 0xdd, 0x10, 0xd2, 0x9c, 0xe7, 0xa3, 0x14, 0xab, 0xbb, 0x68,
	0xc5, 0xdc, 0xad, 0xcb, 0xb1, 0xe9, 0x86, 0x16, 0x38, 0xd3, 0xc2, 0x19, 0x87, 0x4a, 0x3d, 0x88,
	0x69, 0x7b, 0xd4, 0x2c, 0xf5, 0x20, 0xce, 0x7a, 0x10, 0xd3, 0x0e, 0x4e, 0x58, 0x60, 0x32, 0xc0,
	0xcc, 0xe1, 0xcd, 0x63, 0xda, 0x45, 0x3e, 0x83, 0xfe, 0x27, 0x0f, 0x7a, 0x13, 0x9e, 0xda, 0xa3,
	0xfd, 0xe1, 0xe6, 0x92, 0x43, 0xe8, 0x63, 0xa5, 0xa9, 0x58, 0x64, 0xb9, 0x2e, 0x08, 0xbc, 0x92,
	0x26, 0x9d, 0xf3, 0x99, 0x8b, 0x78, 0x06, 0xcd, 0xff, 0x94, 0x0e, 0xa5, 0xc6, 0xdd, 0xb5, 0x71,
	0x77, 0x05, 0x81, 0x76, 0xe3, 0x05, 0xce, 0x75, 0xec, 0xed, 0x73, 0xd0, 0x7f, 0x80, 0x7e, 0xe6,
	0x56, 0x91, 0xff, 0xa0, 0x17, 0x3a, 0xe0, 0x9a, 0x33, 0xac, 0x4c, 0x13, 0x2a, 0x82, 0x5c, 0x6b,
	0x3e, 0x86, 0x77, 0x42, 0x87, 0xcb, 0xa9, 0x58, 0xc7, 0xda, 0xed, 0xa7, 0xc4, 0x9c, 0x7e, 0x6e,
	0x41, 0x73, 0xc2, 0x53, 0x32, 0x87, 0x9d, 0x4b, 0xa6, 0x8b, 0xf5, 0x2a, 0xb2, 0x5a, 0xe4, 0x72,
	0x78, 0xb0, 0x7d, 0x6d, 0x45, 0x66, 0x00, 0xb6, 0xd4, 0xf9, 0x66, 0x3e, 0x23, 0xb4, 0x52, 0x1a,
	0xb0, 0xa7, 0x61, 0xed, 0x75, 0x20, 0xe7, 0xd0, 0xb5, 0x55, 0x14, 0xf9, 0x7d, 0x8b, 0x97, 0xe1,
	0x51, 0x65, 0x85, 0xe2, 0xfa, 0x4d, 0x61, 0x70, 0xab, 0x45, 0xf2, 0xba, 0x95, 0x8a, 0xb3, 0xcc,
	0x8b, 0x4c, 0xa0, 0xe7, 0x8a, 0x28, 0xb2, 0x5f, 0xa5, 0xc3, 0xaf, 0x6d, 0x6d, 0x89, 0x6b, 0x20,
	0x1f, 0x92, 0x45, 0xa8, 0xd9, 0xad, 0x7b, 0x64, 0xe2, 0xc7, 0x88, 0x93, 0xba, 0x97, 0xa4, 0xb6,
	0xdc, 0x15, 0xec, 0x5d, 0x32, 0x5d, 0xae, 0x55, 0xeb, 0xec, 0xb0, 0x66, 0x1d, 0x45, 0xae, 0xe0,
	0xe7, 0x17, 0xc5, 0xaa, 0x8c, 0xe5, 0xef, 0xd9, 0xb0, 0xce, 0xf5, 0x7d, 0x07, 0x1f, 0xef, 0xb3,
	0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xeb, 0x1b, 0x95, 0xd3, 0x07, 0x00, 0x00,
}
