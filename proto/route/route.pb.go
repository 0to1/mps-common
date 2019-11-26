// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/route/route.proto

package go_micro_srv_route

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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{0}
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
	return fileDescriptor_ecb1e1187677b1b4, []int{1}
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

type StationsRequest struct {
	StartStn             uint32   `protobuf:"varint,1,opt,name=startStn,proto3" json:"startStn,omitempty"`
	Stations             []uint32 `protobuf:"varint,2,rep,packed,name=stations,proto3" json:"stations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StationsRequest) Reset()         { *m = StationsRequest{} }
func (m *StationsRequest) String() string { return proto.CompactTextString(m) }
func (*StationsRequest) ProtoMessage()    {}
func (*StationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{2}
}

func (m *StationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StationsRequest.Unmarshal(m, b)
}
func (m *StationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StationsRequest.Marshal(b, m, deterministic)
}
func (m *StationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StationsRequest.Merge(m, src)
}
func (m *StationsRequest) XXX_Size() int {
	return xxx_messageInfo_StationsRequest.Size(m)
}
func (m *StationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StationsRequest proto.InternalMessageInfo

func (m *StationsRequest) GetStartStn() uint32 {
	if m != nil {
		return m.StartStn
	}
	return 0
}

func (m *StationsRequest) GetStations() []uint32 {
	if m != nil {
		return m.Stations
	}
	return nil
}

type GetStationListRequest struct {
	Val1                 uint32   `protobuf:"varint,1,opt,name=val1,proto3" json:"val1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStationListRequest) Reset()         { *m = GetStationListRequest{} }
func (m *GetStationListRequest) String() string { return proto.CompactTextString(m) }
func (*GetStationListRequest) ProtoMessage()    {}
func (*GetStationListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{3}
}

func (m *GetStationListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStationListRequest.Unmarshal(m, b)
}
func (m *GetStationListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStationListRequest.Marshal(b, m, deterministic)
}
func (m *GetStationListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStationListRequest.Merge(m, src)
}
func (m *GetStationListRequest) XXX_Size() int {
	return xxx_messageInfo_GetStationListRequest.Size(m)
}
func (m *GetStationListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStationListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStationListRequest proto.InternalMessageInfo

func (m *GetStationListRequest) GetVal1() uint32 {
	if m != nil {
		return m.Val1
	}
	return 0
}

type StationWeight struct {
	StationID            uint32   `protobuf:"varint,1,opt,name=stationID,proto3" json:"stationID,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StationWeight) Reset()         { *m = StationWeight{} }
func (m *StationWeight) String() string { return proto.CompactTextString(m) }
func (*StationWeight) ProtoMessage()    {}
func (*StationWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{4}
}

func (m *StationWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StationWeight.Unmarshal(m, b)
}
func (m *StationWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StationWeight.Marshal(b, m, deterministic)
}
func (m *StationWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StationWeight.Merge(m, src)
}
func (m *StationWeight) XXX_Size() int {
	return xxx_messageInfo_StationWeight.Size(m)
}
func (m *StationWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_StationWeight.DiscardUnknown(m)
}

var xxx_messageInfo_StationWeight proto.InternalMessageInfo

func (m *StationWeight) GetStationID() uint32 {
	if m != nil {
		return m.StationID
	}
	return 0
}

func (m *StationWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type GetStationListResponse struct {
	Res                  bool             `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	Stations             []*StationWeight `protobuf:"bytes,2,rep,name=stations,proto3" json:"stations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetStationListResponse) Reset()         { *m = GetStationListResponse{} }
func (m *GetStationListResponse) String() string { return proto.CompactTextString(m) }
func (*GetStationListResponse) ProtoMessage()    {}
func (*GetStationListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{5}
}

func (m *GetStationListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStationListResponse.Unmarshal(m, b)
}
func (m *GetStationListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStationListResponse.Marshal(b, m, deterministic)
}
func (m *GetStationListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStationListResponse.Merge(m, src)
}
func (m *GetStationListResponse) XXX_Size() int {
	return xxx_messageInfo_GetStationListResponse.Size(m)
}
func (m *GetStationListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStationListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStationListResponse proto.InternalMessageInfo

func (m *GetStationListResponse) GetRes() bool {
	if m != nil {
		return m.Res
	}
	return false
}

func (m *GetStationListResponse) GetStations() []*StationWeight {
	if m != nil {
		return m.Stations
	}
	return nil
}

type GetRouteRequest struct {
	Val1                 uint32   `protobuf:"varint,1,opt,name=val1,proto3" json:"val1,omitempty"`
	Val2                 uint32   `protobuf:"varint,2,opt,name=val2,proto3" json:"val2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRouteRequest) Reset()         { *m = GetRouteRequest{} }
func (m *GetRouteRequest) String() string { return proto.CompactTextString(m) }
func (*GetRouteRequest) ProtoMessage()    {}
func (*GetRouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{6}
}

func (m *GetRouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRouteRequest.Unmarshal(m, b)
}
func (m *GetRouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRouteRequest.Marshal(b, m, deterministic)
}
func (m *GetRouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRouteRequest.Merge(m, src)
}
func (m *GetRouteRequest) XXX_Size() int {
	return xxx_messageInfo_GetRouteRequest.Size(m)
}
func (m *GetRouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRouteRequest proto.InternalMessageInfo

func (m *GetRouteRequest) GetVal1() uint32 {
	if m != nil {
		return m.Val1
	}
	return 0
}

func (m *GetRouteRequest) GetVal2() uint32 {
	if m != nil {
		return m.Val2
	}
	return 0
}

type GetRouteResponse struct {
	Res                  bool     `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Segments             []uint32 `protobuf:"varint,3,rep,packed,name=segments,proto3" json:"segments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRouteResponse) Reset()         { *m = GetRouteResponse{} }
func (m *GetRouteResponse) String() string { return proto.CompactTextString(m) }
func (*GetRouteResponse) ProtoMessage()    {}
func (*GetRouteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e1187677b1b4, []int{7}
}

func (m *GetRouteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRouteResponse.Unmarshal(m, b)
}
func (m *GetRouteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRouteResponse.Marshal(b, m, deterministic)
}
func (m *GetRouteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRouteResponse.Merge(m, src)
}
func (m *GetRouteResponse) XXX_Size() int {
	return xxx_messageInfo_GetRouteResponse.Size(m)
}
func (m *GetRouteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRouteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRouteResponse proto.InternalMessageInfo

func (m *GetRouteResponse) GetRes() bool {
	if m != nil {
		return m.Res
	}
	return false
}

func (m *GetRouteResponse) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *GetRouteResponse) GetSegments() []uint32 {
	if m != nil {
		return m.Segments
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.route.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.route.Response")
	proto.RegisterType((*StationsRequest)(nil), "go.micro.srv.route.StationsRequest")
	proto.RegisterType((*GetStationListRequest)(nil), "go.micro.srv.route.GetStationListRequest")
	proto.RegisterType((*StationWeight)(nil), "go.micro.srv.route.StationWeight")
	proto.RegisterType((*GetStationListResponse)(nil), "go.micro.srv.route.GetStationListResponse")
	proto.RegisterType((*GetRouteRequest)(nil), "go.micro.srv.route.GetRouteRequest")
	proto.RegisterType((*GetRouteResponse)(nil), "go.micro.srv.route.GetRouteResponse")
}

func init() { proto.RegisterFile("proto/route/route.proto", fileDescriptor_ecb1e1187677b1b4) }

var fileDescriptor_ecb1e1187677b1b4 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x4f, 0xe2, 0x40,
	0x14, 0x5d, 0x3e, 0x96, 0x8f, 0xbb, 0x21, 0xb0, 0x93, 0x5d, 0xb6, 0x61, 0x79, 0x60, 0x67, 0x79,
	0xd8, 0x8f, 0xa4, 0x46, 0x7c, 0xf2, 0xc1, 0x37, 0x95, 0x90, 0x18, 0x43, 0x5a, 0x8c, 0xbe, 0x56,
	0x98, 0x94, 0x26, 0xb6, 0x53, 0x67, 0x2e, 0xf8, 0xa3, 0xfd, 0x13, 0xa6, 0xc3, 0xad, 0x60, 0x29,
	0x44, 0x13, 0x7d, 0x69, 0xe6, 0xce, 0x3d, 0x73, 0xce, 0x3d, 0xa7, 0xd3, 0xc2, 0x8f, 0x58, 0x49,
	0x94, 0x07, 0x4a, 0x2e, 0x50, 0xac, 0x9e, 0xb6, 0xd9, 0x61, 0xcc, 0x97, 0x76, 0x18, 0x4c, 0x95,
	0xb4, 0xb5, 0x5a, 0xda, 0xa6, 0xc3, 0xeb, 0x50, 0x75, 0xc4, 0xfd, 0x42, 0x68, 0xe4, 0x7d, 0xa8,
	0x39, 0x42, 0xc7, 0x32, 0xd2, 0x82, 0x59, 0x50, 0x0d, 0x85, 0xd6, 0x9e, 0x2f, 0xac, 0x42, 0xaf,
	0xf0, 0xa7, 0xee, 0xa4, 0x25, 0x1f, 0x41, 0xd3, 0x45, 0x0f, 0x03, 0x19, 0x69, 0x3a, 0xc8, 0x3a,
	0x50, 0xd3, 0xe8, 0x29, 0x74, 0x31, 0x32, 0xe8, 0x86, 0xf3, 0x5c, 0x53, 0xcf, 0xc0, 0xad, 0x62,
	0xaf, 0x44, 0x3d, 0x53, 0xf3, 0xff, 0xf0, 0x7d, 0x28, 0x90, 0xd8, 0x2e, 0x02, 0x8d, 0x29, 0x21,
	0x83, 0xf2, 0xd2, 0xbb, 0x3b, 0x24, 0x32, 0xb3, 0xe6, 0x67, 0xd0, 0x20, 0xe4, 0xb5, 0x08, 0xfc,
	0x39, 0xb2, 0x2e, 0xd4, 0x89, 0x69, 0x74, 0x4a, 0xc8, 0xf5, 0x06, 0x6b, 0x43, 0xe5, 0xc1, 0xe0,
	0xac, 0xa2, 0x69, 0x51, 0xc5, 0x03, 0x68, 0x67, 0x35, 0xc9, 0x72, 0x0b, 0x4a, 0x4a, 0x68, 0xc3,
	0x54, 0x73, 0x92, 0x25, 0x3b, 0xc9, 0xcc, 0xfe, 0x65, 0xf0, 0xcb, 0xde, 0x8e, 0xd0, 0x7e, 0x31,
	0xd6, 0x86, 0xbd, 0x63, 0x68, 0x0e, 0x05, 0x3a, 0x09, 0x66, 0x8f, 0x31, 0xda, 0x1b, 0xd0, 0x9c,
	0x66, 0xcd, 0x6f, 0xa0, 0xb5, 0x3e, 0xba, 0x73, 0xbe, 0x1d, 0x1e, 0x4d, 0xe6, 0xc2, 0x0f, 0x45,
	0x84, 0xda, 0x2a, 0x51, 0xe6, 0x54, 0x0f, 0x1e, 0xcb, 0xf0, 0xd9, 0xf0, 0xb2, 0x29, 0x7c, 0x4b,
	0x35, 0xce, 0x95, 0x0c, 0x5d, 0x8c, 0x26, 0x32, 0x79, 0x63, 0xbf, 0xf3, 0x3c, 0x66, 0x8c, 0x74,
	0xfa, 0xfb, 0x41, 0xab, 0x91, 0xf9, 0xa7, 0xac, 0xc8, 0x38, 0xc2, 0x89, 0x1c, 0x47, 0xf8, 0xbe,
	0x22, 0x71, 0xf6, 0x1e, 0x91, 0x14, 0xfb, 0xbb, 0x83, 0x60, 0xfb, 0xca, 0x75, 0xfe, 0xbd, 0x06,
	0xba, 0x5f, 0x31, 0x09, 0xef, 0xc3, 0x14, 0x67, 0xd0, 0x74, 0xa5, 0xda, 0x6c, 0xe6, 0x67, 0x98,
	0xf9, 0x36, 0xdf, 0xa8, 0x72, 0x09, 0x5f, 0xaf, 0xe2, 0x99, 0x87, 0xc2, 0x44, 0xec, 0x0a, 0x4f,
	0x4d, 0xe7, 0xec, 0x67, 0x1e, 0x45, 0xca, 0xdf, 0xcd, 0x6f, 0xae, 0x18, 0x6f, 0x2b, 0xe6, 0xc7,
	0x73, 0xf4, 0x14, 0x00, 0x00, 0xff, 0xff, 0xae, 0x6e, 0x70, 0xc2, 0x93, 0x04, 0x00, 0x00,
}