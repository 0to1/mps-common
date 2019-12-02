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

type FileRequest struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Filebytes            []byte   `protobuf:"bytes,3,opt,name=filebytes,proto3" json:"filebytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{0}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *FileRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *FileRequest) GetFilebytes() []byte {
	if m != nil {
		return m.Filebytes
	}
	return nil
}

type FileResponse struct {
	Res                  bool     `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponse) Reset()         { *m = FileResponse{} }
func (m *FileResponse) String() string { return proto.CompactTextString(m) }
func (*FileResponse) ProtoMessage()    {}
func (*FileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50bee624c47b409f, []int{1}
}

func (m *FileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponse.Unmarshal(m, b)
}
func (m *FileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponse.Marshal(b, m, deterministic)
}
func (m *FileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponse.Merge(m, src)
}
func (m *FileResponse) XXX_Size() int {
	return xxx_messageInfo_FileResponse.Size(m)
}
func (m *FileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponse proto.InternalMessageInfo

func (m *FileResponse) GetRes() bool {
	if m != nil {
		return m.Res
	}
	return false
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
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,5,opt,name=perPage,proto3" json:"perPage,omitempty"`
	QueryString          string   `protobuf:"bytes,6,opt,name=queryString,proto3" json:"queryString,omitempty"`
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

func (m *Query) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Query) GetOffset() uint32 {
	if m != nil {
		return m.Offset
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

func (m *Query) GetQueryString() string {
	if m != nil {
		return m.QueryString
	}
	return ""
}

type SegmentResponse struct {
	SegmentID            uint32   `protobuf:"varint,1,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	StartPoint           uint32   `protobuf:"varint,2,opt,name=startPoint,proto3" json:"startPoint,omitempty"`
	EndPoint             uint32   `protobuf:"varint,3,opt,name=endPoint,proto3" json:"endPoint,omitempty"`
	SegmentTemplate      string   `protobuf:"bytes,4,opt,name=segmentTemplate,proto3" json:"segmentTemplate,omitempty"`
	Vehicle              string   `protobuf:"bytes,5,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	TravelTime           uint32   `protobuf:"varint,6,opt,name=travelTime,proto3" json:"travelTime,omitempty"`
	StartsAsStartPoint   uint32   `protobuf:"varint,7,opt,name=startsAsStartPoint,proto3" json:"startsAsStartPoint,omitempty"`
	EndsAsEndPoint       uint32   `protobuf:"varint,8,opt,name=endsAsEndPoint,proto3" json:"endsAsEndPoint,omitempty"`
	Direction            uint32   `protobuf:"varint,9,opt,name=direction,proto3" json:"direction,omitempty"`
	PlcBits              uint32   `protobuf:"varint,10,opt,name=plcBits,proto3" json:"plcBits,omitempty"`
	PlcWord              uint32   `protobuf:"varint,11,opt,name=plcWord,proto3" json:"plcWord,omitempty"`
	PlcWord2             uint32   `protobuf:"varint,12,opt,name=plcWord2,proto3" json:"plcWord2,omitempty"`
	AddWeight            uint32   `protobuf:"varint,13,opt,name=addWeight,proto3" json:"addWeight,omitempty"`
	Length               uint32   `protobuf:"varint,14,opt,name=length,proto3" json:"length,omitempty"`
	CarrierTypes         uint32   `protobuf:"varint,15,opt,name=carrierTypes,proto3" json:"carrierTypes,omitempty"`
	Remark               string   `protobuf:"bytes,16,opt,name=remark,proto3" json:"remark,omitempty"`
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
	Currpages            uint32             `protobuf:"varint,3,opt,name=currpages,proto3" json:"currpages,omitempty"`
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

func (m *SegmentsResponse) GetCurrpages() uint32 {
	if m != nil {
		return m.Currpages
	}
	return 0
}

func init() {
	proto.RegisterType((*FileRequest)(nil), "go.micro.srv.segment.FileRequest")
	proto.RegisterType((*FileResponse)(nil), "go.micro.srv.segment.FileResponse")
	proto.RegisterType((*IdRequest)(nil), "go.micro.srv.segment.IdRequest")
	proto.RegisterType((*Query)(nil), "go.micro.srv.segment.Query")
	proto.RegisterType((*SegmentResponse)(nil), "go.micro.srv.segment.SegmentResponse")
	proto.RegisterType((*SegmentsResponse)(nil), "go.micro.srv.segment.SegmentsResponse")
}

func init() { proto.RegisterFile("proto/segment/segment.proto", fileDescriptor_50bee624c47b409f) }

var fileDescriptor_50bee624c47b409f = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4b, 0x6b, 0xdb, 0x40,
	0x10, 0x8e, 0x9c, 0x97, 0x35, 0x7e, 0x85, 0x25, 0x94, 0xc5, 0x09, 0xad, 0x2b, 0x68, 0xf0, 0x49,
	0x85, 0xf4, 0xd0, 0x73, 0x42, 0xda, 0x92, 0x4b, 0x49, 0x65, 0x43, 0x20, 0xb7, 0x8d, 0x34, 0x51,
	0x96, 0xea, 0x95, 0xdd, 0xb5, 0xc1, 0xbf, 0xa0, 0x3f, 0xa0, 0x87, 0x9e, 0xfb, 0x4f, 0xcb, 0x3e,
	0x24, 0xb9, 0xc6, 0x98, 0x9c, 0xb4, 0xdf, 0x37, 0x9a, 0x99, 0x6f, 0x1e, 0xbb, 0x70, 0x56, 0x89,
	0x52, 0x95, 0x1f, 0x25, 0xa6, 0x39, 0x16, 0xaa, 0xfe, 0x86, 0x86, 0x25, 0xa7, 0x69, 0x19, 0xe6,
	0x3c, 0x16, 0x65, 0x28, 0xc5, 0x32, 0x74, 0xb6, 0x20, 0x86, 0xde, 0x57, 0x9e, 0x61, 0x84, 0x2f,
	0x0b, 0x94, 0x8a, 0x8c, 0xa1, 0x9b, 0x32, 0xc1, 0x52, 0xbc, 0xbd, 0xa1, 0xde, 0xc4, 0x9b, 0x0e,
	0xa2, 0x06, 0x6b, 0xdb, 0x13, 0xcf, 0xf0, 0x3b, 0xcb, 0x91, 0x76, 0x26, 0xde, 0xd4, 0x8f, 0x1a,
	0x4c, 0xce, 0xc1, 0xd7, 0xe7, 0xc7, 0x95, 0x42, 0x49, 0xf7, 0x27, 0xde, 0xb4, 0x1f, 0xb5, 0x44,
	0x30, 0x81, 0xbe, 0x4d, 0x22, 0xab, 0xb2, 0x90, 0x48, 0x4e, 0x60, 0x5f, 0xa0, 0x34, 0x09, 0xba,
	0x91, 0x3e, 0x06, 0x9f, 0xc1, 0xbf, 0x4d, 0x5e, 0x23, 0x62, 0x08, 0x1d, 0x9e, 0x98, 0xf4, 0x83,
	0xa8, 0xc3, 0x93, 0xe0, 0xaf, 0x07, 0x87, 0x3f, 0x16, 0x28, 0x56, 0x3b, 0xbd, 0x4e, 0xe1, 0x30,
	0xe3, 0x39, 0x57, 0xce, 0xd1, 0x02, 0xf2, 0x06, 0x8e, 0xca, 0xa7, 0x27, 0x89, 0xca, 0x28, 0x1e,
	0x44, 0x0e, 0x11, 0x02, 0x07, 0x15, 0x4b, 0x91, 0x1e, 0x18, 0xd6, 0x9c, 0x09, 0x85, 0xe3, 0x0a,
	0xc5, 0x9d, 0xa6, 0x0f, 0x0d, 0x5d, 0x43, 0x32, 0x81, 0xde, 0x8b, 0x16, 0x30, 0x53, 0x82, 0x17,
	0x29, 0x3d, 0x32, 0x9d, 0x59, 0xa7, 0x82, 0x5f, 0x07, 0x30, 0x9a, 0xd9, 0x7e, 0x37, 0x2d, 0x38,
	0x07, 0xdf, 0x8d, 0xa0, 0x91, 0xdb, 0x12, 0xe4, 0x2d, 0x80, 0x54, 0x4c, 0xa8, 0xbb, 0x92, 0x17,
	0xb5, 0xe8, 0x35, 0x46, 0xd7, 0x8a, 0x45, 0x62, 0xad, 0x56, 0x7b, 0x83, 0xc9, 0x14, 0x46, 0x2e,
	0xd0, 0x1c, 0xf3, 0x2a, 0x63, 0xca, 0x16, 0xe2, 0x47, 0x9b, 0xb4, 0xae, 0x69, 0x89, 0xcf, 0x3c,
	0xce, 0x6c, 0x4d, 0x7e, 0x54, 0x43, 0x9d, 0x5f, 0x09, 0xb6, 0xc4, 0x6c, 0xce, 0x73, 0x34, 0x25,
	0x0d, 0xa2, 0x35, 0x86, 0x84, 0x40, 0x8c, 0x1a, 0x79, 0x25, 0x67, 0xad, 0xce, 0x63, 0xf3, 0xdf,
	0x16, 0x0b, 0xb9, 0x80, 0x21, 0x16, 0x89, 0xbc, 0x92, 0x5f, 0x6a, 0xd5, 0x5d, 0xf3, 0xef, 0x06,
	0xab, 0xbb, 0x92, 0x70, 0x81, 0xb1, 0xe2, 0x65, 0x41, 0x7d, 0xdb, 0x95, 0x86, 0x30, 0x33, 0xc8,
	0xe2, 0x6b, 0xae, 0x24, 0x05, 0x37, 0x03, 0x0b, 0x9d, 0xe5, 0xbe, 0x14, 0x09, 0xed, 0x35, 0x16,
	0x0d, 0x75, 0xa7, 0xdc, 0xf1, 0x92, 0xf6, 0x6d, 0xa7, 0x6a, 0xac, 0xb3, 0xb1, 0x24, 0xb9, 0x47,
	0x9e, 0x3e, 0x2b, 0x3a, 0xb0, 0xd9, 0x1a, 0x42, 0x6f, 0x47, 0x86, 0x45, 0xaa, 0x9e, 0xe9, 0xd0,
	0x6e, 0x87, 0x45, 0x24, 0x80, 0x7e, 0xcc, 0x84, 0xe0, 0x28, 0xe6, 0xab, 0x0a, 0x25, 0x1d, 0x19,
	0xeb, 0x7f, 0x9c, 0xf6, 0x15, 0x98, 0x33, 0xf1, 0x93, 0x9e, 0x98, 0xc6, 0x3a, 0x14, 0xfc, 0xf6,
	0xe0, 0xc4, 0x6d, 0x82, 0x6c, 0x56, 0xe1, 0x0a, 0xba, 0x6e, 0x32, 0xfa, 0x4a, 0xec, 0x4f, 0x7b,
	0x97, 0x1f, 0xc2, 0x6d, 0x77, 0x35, 0xdc, 0xd8, 0xa1, 0xa8, 0x71, 0x33, 0xf3, 0x2a, 0x15, 0xcb,
	0xf4, 0xaa, 0xca, 0x7a, 0x5f, 0x5a, 0x46, 0x57, 0x1a, 0x2f, 0x84, 0xb0, 0x66, 0xbb, 0x30, 0x2d,
	0x71, 0xf9, 0xa7, 0x03, 0xc7, 0x2e, 0x36, 0x79, 0x80, 0xd1, 0x8c, 0x2d, 0xd1, 0x41, 0x7d, 0x6b,
	0xc9, 0xfb, 0xed, 0x6a, 0xd6, 0x9e, 0x8d, 0x71, 0xb0, 0xeb, 0x17, 0xab, 0x36, 0xd8, 0x23, 0x0f,
	0x30, 0xfc, 0x86, 0xca, 0x85, 0xbe, 0x5e, 0xdd, 0xde, 0x90, 0x77, 0xdb, 0xfd, 0x9a, 0xa7, 0x60,
	0xfc, 0xba, 0x4e, 0x04, 0x7b, 0x64, 0x0e, 0xbd, 0x36, 0xb6, 0x24, 0x67, 0xdb, 0xfd, 0xcc, 0x4b,
	0x31, 0xbe, 0xd8, 0x19, 0x54, 0xb6, 0x51, 0x1f, 0x8f, 0xcc, 0xd3, 0xf9, 0xe9, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0x78, 0x92, 0xe1, 0x59, 0x05, 0x00, 0x00,
}
