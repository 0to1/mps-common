// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/segment/segment.proto

package go_micro_srv_segment

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

type GraphqlQuery struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlQuery) Reset()         { *m = GraphqlQuery{} }
func (m *GraphqlQuery) String() string { return proto.CompactTextString(m) }
func (*GraphqlQuery) ProtoMessage()    {}
func (*GraphqlQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{0}
}

func (m *GraphqlQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlQuery.Unmarshal(m, b)
}
func (m *GraphqlQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlQuery.Marshal(b, m, deterministic)
}
func (m *GraphqlQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlQuery.Merge(m, src)
}
func (m *GraphqlQuery) XXX_Size() int {
	return xxx_messageInfo_GraphqlQuery.Size(m)
}
func (m *GraphqlQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlQuery proto.InternalMessageInfo

func (m *GraphqlQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type GraphqlSegments struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlSegments) Reset()         { *m = GraphqlSegments{} }
func (m *GraphqlSegments) String() string { return proto.CompactTextString(m) }
func (*GraphqlSegments) ProtoMessage()    {}
func (*GraphqlSegments) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{1}
}

func (m *GraphqlSegments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlSegments.Unmarshal(m, b)
}
func (m *GraphqlSegments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlSegments.Marshal(b, m, deterministic)
}
func (m *GraphqlSegments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlSegments.Merge(m, src)
}
func (m *GraphqlSegments) XXX_Size() int {
	return xxx_messageInfo_GraphqlSegments.Size(m)
}
func (m *GraphqlSegments) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlSegments.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlSegments proto.InternalMessageInfo

func (m *GraphqlSegments) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type IdRequest struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{2}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *IdRequest) GetId() uint32 {
	if m != nil {
		return m.Id
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
	return fileDescriptor_50bee624c47b409f, []int{3}
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

type SegmentResponse struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	SegmentID            uint32   `protobuf:"varint,2,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	StartPoint           uint32   `protobuf:"varint,3,opt,name=startPoint,proto3" json:"startPoint,omitempty"`
	EndPoint             uint32   `protobuf:"varint,4,opt,name=endPoint,proto3" json:"endPoint,omitempty"`
	SegmentTemplate      string   `protobuf:"bytes,5,opt,name=segmentTemplate,proto3" json:"segmentTemplate,omitempty"`
	Vehicle              string   `protobuf:"bytes,6,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	TravelTime           uint32   `protobuf:"varint,7,opt,name=travelTime,proto3" json:"travelTime,omitempty"`
	StartsAsStartPoint   uint32   `protobuf:"varint,8,opt,name=startsAsStartPoint,proto3" json:"startsAsStartPoint,omitempty"`
	EndsAsEndPoint       uint32   `protobuf:"varint,9,opt,name=endsAsEndPoint,proto3" json:"endsAsEndPoint,omitempty"`
	Direction            uint32   `protobuf:"varint,10,opt,name=direction,proto3" json:"direction,omitempty"`
	PlcBits              uint32   `protobuf:"varint,11,opt,name=plcBits,proto3" json:"plcBits,omitempty"`
	PlcWord              uint32   `protobuf:"varint,12,opt,name=plcWord,proto3" json:"plcWord,omitempty"`
	PlcWord2             uint32   `protobuf:"varint,13,opt,name=plcWord2,proto3" json:"plcWord2,omitempty"`
	AddWeight            uint32   `protobuf:"varint,14,opt,name=addWeight,proto3" json:"addWeight,omitempty"`
	Length               uint32   `protobuf:"varint,15,opt,name=length,proto3" json:"length,omitempty"`
	CarrierTypes         uint32   `protobuf:"varint,16,opt,name=carrierTypes,proto3" json:"carrierTypes,omitempty"`
	Remark               string   `protobuf:"bytes,17,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentResponse) Reset()         { *m = SegmentResponse{} }
func (m *SegmentResponse) String() string { return proto.CompactTextString(m) }
func (*SegmentResponse) ProtoMessage()    {}
func (*SegmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{4}
}

func (m *SegmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentResponse.Unmarshal(m, b)
}
func (m *SegmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentResponse.Marshal(b, m, deterministic)
}
func (m *SegmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentResponse.Merge(m, src)
}
func (m *SegmentResponse) XXX_Size() int {
	return xxx_messageInfo_SegmentResponse.Size(m)
}
func (m *SegmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentResponse proto.InternalMessageInfo

func (m *SegmentResponse) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *SegmentResponse) GetSegmentID() uint32 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *SegmentResponse) GetStartPoint() uint32 {
	if m != nil {
		return m.StartPoint
	}
	return 0
}

func (m *SegmentResponse) GetEndPoint() uint32 {
	if m != nil {
		return m.EndPoint
	}
	return 0
}

func (m *SegmentResponse) GetSegmentTemplate() string {
	if m != nil {
		return m.SegmentTemplate
	}
	return ""
}

func (m *SegmentResponse) GetVehicle() string {
	if m != nil {
		return m.Vehicle
	}
	return ""
}

func (m *SegmentResponse) GetTravelTime() uint32 {
	if m != nil {
		return m.TravelTime
	}
	return 0
}

func (m *SegmentResponse) GetStartsAsStartPoint() uint32 {
	if m != nil {
		return m.StartsAsStartPoint
	}
	return 0
}

func (m *SegmentResponse) GetEndsAsEndPoint() uint32 {
	if m != nil {
		return m.EndsAsEndPoint
	}
	return 0
}

func (m *SegmentResponse) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *SegmentResponse) GetPlcBits() uint32 {
	if m != nil {
		return m.PlcBits
	}
	return 0
}

func (m *SegmentResponse) GetPlcWord() uint32 {
	if m != nil {
		return m.PlcWord
	}
	return 0
}

func (m *SegmentResponse) GetPlcWord2() uint32 {
	if m != nil {
		return m.PlcWord2
	}
	return 0
}

func (m *SegmentResponse) GetAddWeight() uint32 {
	if m != nil {
		return m.AddWeight
	}
	return 0
}

func (m *SegmentResponse) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *SegmentResponse) GetCarrierTypes() uint32 {
	if m != nil {
		return m.CarrierTypes
	}
	return 0
}

func (m *SegmentResponse) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type SegmentsResponse struct {
	Segments             []*SegmentResponse `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"`
	Totalpages           uint32             `protobuf:"varint,2,opt,name=totalpages,proto3" json:"totalpages,omitempty"`
	TotalNumber          uint32             `protobuf:"varint,3,opt,name=totalNumber,proto3" json:"totalNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SegmentsResponse) Reset()         { *m = SegmentsResponse{} }
func (m *SegmentsResponse) String() string { return proto.CompactTextString(m) }
func (*SegmentsResponse) ProtoMessage()    {}
func (*SegmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{5}
}

func (m *SegmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentsResponse.Unmarshal(m, b)
}
func (m *SegmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentsResponse.Marshal(b, m, deterministic)
}
func (m *SegmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentsResponse.Merge(m, src)
}
func (m *SegmentsResponse) XXX_Size() int {
	return xxx_messageInfo_SegmentsResponse.Size(m)
}
func (m *SegmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentsResponse proto.InternalMessageInfo

func (m *SegmentsResponse) GetSegments() []*SegmentResponse {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *SegmentsResponse) GetTotalpages() uint32 {
	if m != nil {
		return m.Totalpages
	}
	return 0
}

func (m *SegmentsResponse) GetTotalNumber() uint32 {
	if m != nil {
		return m.TotalNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*GraphqlQuery)(nil), "go.micro.srv.segment.GraphqlQuery")
	proto.RegisterType((*GraphqlSegments)(nil), "go.micro.srv.segment.GraphqlSegments")
	proto.RegisterType((*IdRequest)(nil), "go.micro.srv.segment.IdRequest")
	proto.RegisterType((*Query)(nil), "go.micro.srv.segment.Query")
	proto.RegisterType((*SegmentResponse)(nil), "go.micro.srv.segment.SegmentResponse")
	proto.RegisterType((*SegmentsResponse)(nil), "go.micro.srv.segment.SegmentsResponse")
}

func init() { proto.RegisterFile("proto/segment/segment.proto", fileDescriptor_50bee624c47b409f) }

var fileDescriptor_50bee624c47b409f = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x8b, 0x1a, 0x41,
	0x10, 0x5d, 0xdd, 0xf5, 0xab, 0x74, 0x75, 0xd3, 0x48, 0x68, 0xdc, 0x90, 0xc8, 0x90, 0x2c, 0x5e,
	0x32, 0x01, 0x73, 0xc8, 0x59, 0x31, 0x2c, 0x5e, 0xc2, 0xae, 0x2b, 0x2c, 0xe4, 0xd6, 0x3b, 0x53,
	0x8c, 0x4d, 0xe6, 0xcb, 0xee, 0x56, 0xf0, 0x97, 0xe4, 0x1a, 0xc8, 0x1f, 0x0d, 0xdd, 0xd3, 0x33,
	0x4e, 0x64, 0x22, 0x39, 0xd9, 0xef, 0x55, 0xd5, 0xab, 0xf2, 0x4d, 0x75, 0xc3, 0x6d, 0x2a, 0x12,
	0x95, 0x7c, 0x92, 0x18, 0x44, 0x18, 0xab, 0xfc, 0xd7, 0x35, 0x2c, 0x19, 0x06, 0x89, 0x1b, 0x71,
	0x4f, 0x24, 0xae, 0x14, 0x7b, 0xd7, 0xc6, 0x9c, 0xf7, 0xd0, 0xbb, 0x17, 0x2c, 0xdd, 0x6c, 0xc3,
	0xc7, 0x1d, 0x8a, 0x03, 0x19, 0x42, 0x63, 0xab, 0x0f, 0xb4, 0x36, 0xae, 0x4d, 0x3a, 0xab, 0x0c,
	0x38, 0x1f, 0x61, 0x60, 0xb3, 0x9e, 0xb2, 0x3a, 0x49, 0x46, 0xd0, 0x16, 0x28, 0xd3, 0x24, 0x96,
	0x68, 0x73, 0x0b, 0xec, 0x7c, 0x81, 0xce, 0xd2, 0x5f, 0xe1, 0x76, 0x87, 0x52, 0xe9, 0xc4, 0x80,
	0x09, 0x16, 0xe0, 0x72, 0x61, 0x12, 0xaf, 0x57, 0x05, 0x26, 0x7d, 0xa8, 0x73, 0x9f, 0xd6, 0x0d,
	0x5b, 0xe7, 0xbe, 0xf3, 0x08, 0x8d, 0x6c, 0x8c, 0x73, 0x45, 0x04, 0xae, 0x52, 0x16, 0xa0, 0x2d,
	0x33, 0x67, 0x42, 0xa1, 0x95, 0xa2, 0x78, 0xd0, 0xf4, 0xa5, 0xa1, 0x73, 0xe8, 0xfc, 0xbe, 0x82,
	0x81, 0x1d, 0x7a, 0x65, 0xe7, 0x3b, 0xab, 0xfe, 0x06, 0x3a, 0xd6, 0x9b, 0xe5, 0xc2, 0xb6, 0x38,
	0x12, 0xe4, 0x2d, 0x80, 0x54, 0x4c, 0xa8, 0x87, 0x84, 0xc7, 0xca, 0xb6, 0x2a, 0x31, 0x5a, 0x19,
	0x63, 0x3f, 0x8b, 0x5e, 0x65, 0xca, 0x39, 0x26, 0x13, 0x18, 0x58, 0xa1, 0x35, 0x46, 0x69, 0xc8,
	0x14, 0xd2, 0x86, 0x31, 0xee, 0x94, 0xd6, 0xff, 0x66, 0x8f, 0x1b, 0xee, 0x85, 0x48, 0x9b, 0x26,
	0x23, 0x87, 0xba, 0xbf, 0x12, 0x6c, 0x8f, 0xe1, 0x9a, 0x47, 0x48, 0x5b, 0x59, 0xff, 0x23, 0x43,
	0x5c, 0x20, 0x66, 0x1a, 0x39, 0x93, 0x4f, 0xc7, 0x39, 0xdb, 0x26, 0xaf, 0x22, 0x42, 0xee, 0xa0,
	0x8f, 0xb1, 0x2f, 0x67, 0xf2, 0x6b, 0x3e, 0x75, 0xc7, 0xe4, 0x9e, 0xb0, 0xda, 0x15, 0x9f, 0x0b,
	0xf4, 0x14, 0x4f, 0x62, 0x0a, 0x99, 0x2b, 0x05, 0x61, 0xdc, 0x0f, 0xbd, 0x39, 0x57, 0x92, 0x76,
	0xad, 0xfb, 0x19, 0xb4, 0x91, 0xe7, 0x44, 0xf8, 0xb4, 0x57, 0x44, 0x34, 0xd4, 0x4e, 0xd9, 0xe3,
	0x94, 0x5e, 0x67, 0x4e, 0xe5, 0x58, 0x77, 0x63, 0xbe, 0xff, 0x8c, 0x3c, 0xd8, 0x28, 0xda, 0xcf,
	0xba, 0x15, 0x04, 0x79, 0x0d, 0xcd, 0x10, 0xe3, 0x40, 0x6d, 0xe8, 0xc0, 0x84, 0x2c, 0x22, 0x0e,
	0xf4, 0x3c, 0x26, 0x04, 0x47, 0xb1, 0x3e, 0xa4, 0x28, 0xe9, 0x8d, 0x89, 0xfe, 0xc5, 0xe9, 0x5a,
	0x81, 0x11, 0x13, 0x3f, 0xe8, 0x2b, 0x63, 0xac, 0x45, 0xce, 0xcf, 0x1a, 0xdc, 0xe4, 0xab, 0x5d,
	0xac, 0xc9, 0x0c, 0xda, 0xf6, 0xcb, 0x48, 0x5a, 0x1b, 0x5f, 0x4e, 0xba, 0xd3, 0x0f, 0x6e, 0xd5,
	0x25, 0x72, 0x4f, 0xf6, 0x6b, 0x55, 0x94, 0x99, 0xef, 0x95, 0x28, 0x16, 0xea, 0x25, 0x95, 0x76,
	0x9d, 0x4a, 0x0c, 0x19, 0x43, 0xd7, 0xa0, 0x6f, 0xbb, 0xe8, 0x05, 0x85, 0x5d, 0xa8, 0x32, 0x35,
	0xfd, 0x55, 0x87, 0x96, 0xd5, 0x27, 0xdf, 0xa1, 0x7f, 0x8f, 0xca, 0xa2, 0xf9, 0x61, 0xb9, 0x20,
	0xef, 0xaa, 0x07, 0x2a, 0x6e, 0xdf, 0xe8, 0xff, 0x26, 0x76, 0x2e, 0xc8, 0x1a, 0xba, 0x47, 0x6d,
	0x49, 0x6e, 0xab, 0xeb, 0xcc, 0xed, 0x1c, 0xdd, 0x9d, 0x15, 0x95, 0x25, 0x55, 0x0f, 0x86, 0x25,
	0xd5, 0xf9, 0xc1, 0xbe, 0x22, 0xc4, 0xa9, 0x56, 0x28, 0x3f, 0x45, 0xff, 0x1a, 0xfd, 0xe4, 0x21,
	0x72, 0x2e, 0x5e, 0x9a, 0xe6, 0x81, 0xfb, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xac, 0xae,
	0x7a, 0xff, 0x04, 0x00, 0x00,
}
