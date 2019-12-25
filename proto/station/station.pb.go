// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/station/station.proto

package go_micro_srv_station

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
	return fileDescriptor_89d528e10793bd7b, []int{0}
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

type GraphqlStations struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlStations) Reset()         { *m = GraphqlStations{} }
func (m *GraphqlStations) String() string { return proto.CompactTextString(m) }
func (*GraphqlStations) ProtoMessage()    {}
func (*GraphqlStations) Descriptor() ([]byte, []int) {
	return fileDescriptor_89d528e10793bd7b, []int{1}
}

func (m *GraphqlStations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlStations.Unmarshal(m, b)
}
func (m *GraphqlStations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlStations.Marshal(b, m, deterministic)
}
func (m *GraphqlStations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlStations.Merge(m, src)
}
func (m *GraphqlStations) XXX_Size() int {
	return xxx_messageInfo_GraphqlStations.Size(m)
}
func (m *GraphqlStations) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlStations.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlStations proto.InternalMessageInfo

func (m *GraphqlStations) GetResponse() string {
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
	return fileDescriptor_89d528e10793bd7b, []int{2}
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
	return fileDescriptor_89d528e10793bd7b, []int{3}
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

type StationResponse struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	StationID            uint32   `protobuf:"varint,2,opt,name=stationID,proto3" json:"stationID,omitempty"`
	PointID              uint32   `protobuf:"varint,3,opt,name=pointID,proto3" json:"pointID,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	SysModes             uint32   `protobuf:"varint,5,opt,name=sysModes,proto3" json:"sysModes,omitempty"`
	CarrierTypes         uint32   `protobuf:"varint,6,opt,name=carrierTypes,proto3" json:"carrierTypes,omitempty"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	StationName          string   `protobuf:"bytes,8,opt,name=stationName,proto3" json:"stationName,omitempty"`
	X                    uint32   `protobuf:"varint,9,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint32   `protobuf:"varint,10,opt,name=y,proto3" json:"y,omitempty"`
	Angle                uint32   `protobuf:"varint,11,opt,name=angle,proto3" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StationResponse) Reset()         { *m = StationResponse{} }
func (m *StationResponse) String() string { return proto.CompactTextString(m) }
func (*StationResponse) ProtoMessage()    {}
func (*StationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89d528e10793bd7b, []int{4}
}

func (m *StationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StationResponse.Unmarshal(m, b)
}
func (m *StationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StationResponse.Marshal(b, m, deterministic)
}
func (m *StationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StationResponse.Merge(m, src)
}
func (m *StationResponse) XXX_Size() int {
	return xxx_messageInfo_StationResponse.Size(m)
}
func (m *StationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StationResponse proto.InternalMessageInfo

func (m *StationResponse) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *StationResponse) GetStationID() uint32 {
	if m != nil {
		return m.StationID
	}
	return 0
}

func (m *StationResponse) GetPointID() uint32 {
	if m != nil {
		return m.PointID
	}
	return 0
}

func (m *StationResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StationResponse) GetSysModes() uint32 {
	if m != nil {
		return m.SysModes
	}
	return 0
}

func (m *StationResponse) GetCarrierTypes() uint32 {
	if m != nil {
		return m.CarrierTypes
	}
	return 0
}

func (m *StationResponse) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *StationResponse) GetStationName() string {
	if m != nil {
		return m.StationName
	}
	return ""
}

func (m *StationResponse) GetX() uint32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *StationResponse) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *StationResponse) GetAngle() uint32 {
	if m != nil {
		return m.Angle
	}
	return 0
}

type StationsResponse struct {
	Stations             []*StationResponse `protobuf:"bytes,1,rep,name=stations,proto3" json:"stations,omitempty"`
	Totalpages           uint32             `protobuf:"varint,2,opt,name=totalpages,proto3" json:"totalpages,omitempty"`
	TotalNumber          uint32             `protobuf:"varint,3,opt,name=totalNumber,proto3" json:"totalNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StationsResponse) Reset()         { *m = StationsResponse{} }
func (m *StationsResponse) String() string { return proto.CompactTextString(m) }
func (*StationsResponse) ProtoMessage()    {}
func (*StationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89d528e10793bd7b, []int{5}
}

func (m *StationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StationsResponse.Unmarshal(m, b)
}
func (m *StationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StationsResponse.Marshal(b, m, deterministic)
}
func (m *StationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StationsResponse.Merge(m, src)
}
func (m *StationsResponse) XXX_Size() int {
	return xxx_messageInfo_StationsResponse.Size(m)
}
func (m *StationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StationsResponse proto.InternalMessageInfo

func (m *StationsResponse) GetStations() []*StationResponse {
	if m != nil {
		return m.Stations
	}
	return nil
}

func (m *StationsResponse) GetTotalpages() uint32 {
	if m != nil {
		return m.Totalpages
	}
	return 0
}

func (m *StationsResponse) GetTotalNumber() uint32 {
	if m != nil {
		return m.TotalNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*GraphqlQuery)(nil), "go.micro.srv.station.GraphqlQuery")
	proto.RegisterType((*GraphqlStations)(nil), "go.micro.srv.station.GraphqlStations")
	proto.RegisterType((*IdRequest)(nil), "go.micro.srv.station.IdRequest")
	proto.RegisterType((*Query)(nil), "go.micro.srv.station.Query")
	proto.RegisterType((*StationResponse)(nil), "go.micro.srv.station.StationResponse")
	proto.RegisterType((*StationsResponse)(nil), "go.micro.srv.station.StationsResponse")
}

func init() { proto.RegisterFile("proto/station/station.proto", fileDescriptor_89d528e10793bd7b) }

var fileDescriptor_89d528e10793bd7b = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xdd, 0xe6, 0x6b, 0x12, 0x5a, 0xb4, 0x8a, 0xd0, 0x2a, 0x45, 0x10, 0xad, 0x00, 0xf5,
	0x82, 0x91, 0xca, 0x81, 0x33, 0x55, 0xa4, 0xca, 0x07, 0x2a, 0x6a, 0x7a, 0xe2, 0xb6, 0x4d, 0x46,
	0xc6, 0x22, 0xc9, 0x3a, 0xbb, 0x1b, 0x54, 0xff, 0x12, 0xae, 0x5c, 0xf8, 0x9f, 0xc8, 0xe3, 0x89,
	0x6b, 0x22, 0x13, 0x71, 0xf2, 0xbe, 0xb7, 0x33, 0x6f, 0x9e, 0x67, 0x67, 0xe0, 0x3c, 0xb7, 0xc6,
	0x9b, 0x77, 0xce, 0x6b, 0x9f, 0x99, 0xf5, 0xee, 0x1b, 0x11, 0x2b, 0xc6, 0xa9, 0x89, 0x56, 0xd9,
	0xdc, 0x9a, 0xc8, 0xd9, 0x1f, 0x11, 0xdf, 0xa9, 0x57, 0x30, 0xba, 0xb6, 0x3a, 0xff, 0xb6, 0x59,
	0xde, 0x6e, 0xd1, 0x16, 0x62, 0x0c, 0x9d, 0x4d, 0x79, 0x90, 0xc1, 0x34, 0xb8, 0x18, 0x24, 0x15,
	0x50, 0x6f, 0xe1, 0x8c, 0xa3, 0xbe, 0x54, 0x79, 0x4e, 0x4c, 0xa0, 0x6f, 0xd1, 0xe5, 0x66, 0xed,
	0x90, 0x63, 0x6b, 0xac, 0x3e, 0xc0, 0x20, 0x5e, 0x24, 0xb8, 0xd9, 0xa2, 0xf3, 0x65, 0x60, 0xaa,
	0xad, 0x4e, 0x31, 0x9e, 0x51, 0xe0, 0x93, 0xa4, 0xc6, 0xe2, 0x14, 0xc2, 0x6c, 0x21, 0x43, 0x62,
	0xc3, 0x6c, 0xa1, 0x6e, 0xa1, 0x53, 0xd9, 0x38, 0x94, 0x24, 0xe0, 0x24, 0xd7, 0x29, 0x72, 0x1a,
	0x9d, 0x85, 0x84, 0x5e, 0x8e, 0xf6, 0x73, 0x49, 0x1f, 0x13, 0xbd, 0x83, 0xea, 0x77, 0x08, 0x67,
	0x6c, 0x3a, 0x61, 0x7f, 0x07, 0xd5, 0x9f, 0xc3, 0x80, 0x7b, 0x13, 0xcf, 0xb8, 0xc4, 0x23, 0x41,
	0x75, 0x4c, 0xb6, 0xf6, 0xf1, 0xac, 0xae, 0x53, 0xc1, 0xd2, 0x95, 0x2f, 0x72, 0x94, 0x27, 0xd4,
	0x0b, 0x3a, 0x97, 0x75, 0x5c, 0xe1, 0x3e, 0x99, 0x05, 0x3a, 0xd9, 0xa9, 0xea, 0xec, 0xb0, 0x50,
	0x30, 0x9a, 0x6b, 0x6b, 0x33, 0xb4, 0x77, 0x45, 0x8e, 0x4e, 0x76, 0xe9, 0xfe, 0x2f, 0x4e, 0x3c,
	0x83, 0xae, 0xc5, 0x95, 0xb6, 0xdf, 0x65, 0x8f, 0x54, 0x19, 0x89, 0x29, 0x0c, 0xd9, 0xd2, 0x8d,
	0x5e, 0xa1, 0xec, 0xd3, 0x65, 0x93, 0x12, 0x23, 0x08, 0x1e, 0xe4, 0x80, 0x24, 0x83, 0x87, 0x12,
	0x15, 0x12, 0x2a, 0x44, 0x4f, 0xac, 0xd7, 0xe9, 0x12, 0xe5, 0x90, 0x98, 0x0a, 0xa8, 0x9f, 0x01,
	0x3c, 0xdd, 0x3d, 0x6e, 0xdd, 0xa8, 0x8f, 0xd0, 0x67, 0x55, 0x27, 0x83, 0xe9, 0xf1, 0xc5, 0xf0,
	0xf2, 0x75, 0xd4, 0x36, 0x46, 0xd1, 0x5e, 0x87, 0x93, 0x3a, 0x4d, 0xbc, 0x00, 0xf0, 0xc6, 0xeb,
	0x65, 0xf9, 0x4c, 0x8e, 0x1b, 0xda, 0x60, 0xca, 0x7f, 0x21, 0x74, 0xb3, 0x5d, 0xdd, 0xa3, 0xe5,
	0xae, 0x36, 0xa9, 0xcb, 0x5f, 0x21, 0xf4, 0x58, 0x5f, 0x7c, 0x85, 0xd3, 0x6b, 0xf4, 0x8c, 0xae,
	0x8a, 0x78, 0x26, 0x5e, 0xb6, 0x1b, 0xaa, 0xe7, 0x6f, 0xf2, 0x7f, 0x8e, 0xd5, 0x91, 0xb8, 0x83,
	0xe1, 0xa3, 0xb6, 0x13, 0xe7, 0xed, 0x79, 0x34, 0x9f, 0x93, 0x37, 0x07, 0x45, 0x5d, 0x43, 0x75,
	0x0e, 0xe3, 0x86, 0xea, 0x55, 0xc1, 0x7b, 0x24, 0x54, 0xbb, 0x42, 0x73, 0x19, 0xff, 0x65, 0x7d,
	0x6f, 0x15, 0xd5, 0xd1, 0x7d, 0x97, 0x56, 0xfc, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16,
	0xf1, 0x83, 0x6d, 0x01, 0x04, 0x00, 0x00,
}
