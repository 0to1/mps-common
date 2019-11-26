// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/point/point.proto

package go_micro_srv_point

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
	return fileDescriptor_3b92d483b7a94ed0, []int{0}
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
	Res                  bool     `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92d483b7a94ed0, []int{1}
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

type FileRequest struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Filebytes            []byte   `protobuf:"bytes,2,opt,name=filebytes,proto3" json:"filebytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92d483b7a94ed0, []int{2}
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
	return fileDescriptor_3b92d483b7a94ed0, []int{3}
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
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92d483b7a94ed0, []int{4}
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

func (m *IdRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Query struct {
	Limit                uint32   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Page                 uint32   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	QueryString          string   `protobuf:"bytes,5,opt,name=queryString,proto3" json:"queryString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92d483b7a94ed0, []int{5}
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

type PointResponse struct {
	PointID              uint32   `protobuf:"varint,1,opt,name=pointID,proto3" json:"pointID,omitempty"`
	X                    uint32   `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint32   `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Angle                uint32   `protobuf:"varint,4,opt,name=angle,proto3" json:"angle,omitempty"`
	Curvature            uint32   `protobuf:"varint,5,opt,name=curvature,proto3" json:"curvature,omitempty"`
	PlcBits              uint32   `protobuf:"varint,6,opt,name=plcBits,proto3" json:"plcBits,omitempty"`
	PlcWord              uint32   `protobuf:"varint,7,opt,name=plcWord,proto3" json:"plcWord,omitempty"`
	PlcWord2             uint32   `protobuf:"varint,8,opt,name=plcWord2,proto3" json:"plcWord2,omitempty"`
	ChangeTime           uint32   `protobuf:"varint,9,opt,name=changeTime,proto3" json:"changeTime,omitempty"`
	Remark               string   `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointResponse) Reset()         { *m = PointResponse{} }
func (m *PointResponse) String() string { return proto.CompactTextString(m) }
func (*PointResponse) ProtoMessage()    {}
func (*PointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92d483b7a94ed0, []int{6}
}

func (m *PointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointResponse.Unmarshal(m, b)
}
func (m *PointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointResponse.Marshal(b, m, deterministic)
}
func (m *PointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointResponse.Merge(m, src)
}
func (m *PointResponse) XXX_Size() int {
	return xxx_messageInfo_PointResponse.Size(m)
}
func (m *PointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PointResponse proto.InternalMessageInfo

func (m *PointResponse) GetPointID() uint32 {
	if m != nil {
		return m.PointID
	}
	return 0
}

func (m *PointResponse) GetX() uint32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *PointResponse) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *PointResponse) GetAngle() uint32 {
	if m != nil {
		return m.Angle
	}
	return 0
}

func (m *PointResponse) GetCurvature() uint32 {
	if m != nil {
		return m.Curvature
	}
	return 0
}

func (m *PointResponse) GetPlcBits() uint32 {
	if m != nil {
		return m.PlcBits
	}
	return 0
}

func (m *PointResponse) GetPlcWord() uint32 {
	if m != nil {
		return m.PlcWord
	}
	return 0
}

func (m *PointResponse) GetPlcWord2() uint32 {
	if m != nil {
		return m.PlcWord2
	}
	return 0
}

func (m *PointResponse) GetChangeTime() uint32 {
	if m != nil {
		return m.ChangeTime
	}
	return 0
}

func (m *PointResponse) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type PointsResponse struct {
	Points               []*PointResponse `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Totalpages           uint32           `protobuf:"varint,2,opt,name=totalpages,proto3" json:"totalpages,omitempty"`
	Currpages            uint32           `protobuf:"varint,3,opt,name=currpages,proto3" json:"currpages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PointsResponse) Reset()         { *m = PointsResponse{} }
func (m *PointsResponse) String() string { return proto.CompactTextString(m) }
func (*PointsResponse) ProtoMessage()    {}
func (*PointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b92d483b7a94ed0, []int{7}
}

func (m *PointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointsResponse.Unmarshal(m, b)
}
func (m *PointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointsResponse.Marshal(b, m, deterministic)
}
func (m *PointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointsResponse.Merge(m, src)
}
func (m *PointsResponse) XXX_Size() int {
	return xxx_messageInfo_PointsResponse.Size(m)
}
func (m *PointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PointsResponse proto.InternalMessageInfo

func (m *PointsResponse) GetPoints() []*PointResponse {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *PointsResponse) GetTotalpages() uint32 {
	if m != nil {
		return m.Totalpages
	}
	return 0
}

func (m *PointsResponse) GetCurrpages() uint32 {
	if m != nil {
		return m.Currpages
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.point.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.point.Response")
	proto.RegisterType((*FileRequest)(nil), "go.micro.srv.point.FileRequest")
	proto.RegisterType((*FileResponse)(nil), "go.micro.srv.point.FileResponse")
	proto.RegisterType((*IdRequest)(nil), "go.micro.srv.point.IdRequest")
	proto.RegisterType((*Query)(nil), "go.micro.srv.point.Query")
	proto.RegisterType((*PointResponse)(nil), "go.micro.srv.point.PointResponse")
	proto.RegisterType((*PointsResponse)(nil), "go.micro.srv.point.PointsResponse")
}

func init() { proto.RegisterFile("proto/point/point.proto", fileDescriptor_3b92d483b7a94ed0) }

var fileDescriptor_3b92d483b7a94ed0 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0xda, 0xf5, 0x47, 0xde, 0xda, 0x81, 0x2c, 0x04, 0x5e, 0x57, 0xa0, 0xf8, 0xd4, 0x53,
	0x90, 0xca, 0x89, 0x23, 0x53, 0xc5, 0x54, 0x09, 0x4d, 0xc3, 0x9b, 0xc4, 0x11, 0x65, 0xed, 0x6b,
	0xb1, 0x48, 0x9b, 0xcc, 0x76, 0xab, 0xf5, 0xc8, 0x05, 0x89, 0xbf, 0x8e, 0x7f, 0x09, 0xf9, 0xc5,
	0xc9, 0x82, 0x48, 0xc5, 0x25, 0xf2, 0xf7, 0x7d, 0x7e, 0xbf, 0x3e, 0x3f, 0x05, 0x5e, 0x64, 0x3a,
	0xb5, 0xe9, 0xdb, 0x2c, 0x55, 0x1b, 0x9b, 0x7f, 0x23, 0x62, 0x18, 0x5b, 0xa5, 0xd1, 0x5a, 0xcd,
	0x75, 0x1a, 0x19, 0xbd, 0x8b, 0x48, 0x11, 0x21, 0x74, 0x24, 0xde, 0x6f, 0xd1, 0x58, 0x31, 0x84,
	0xae, 0x44, 0x93, 0xa5, 0x1b, 0x83, 0xec, 0x29, 0x34, 0x35, 0x1a, 0x1e, 0x8c, 0x82, 0x71, 0x57,
	0xba, 0xa3, 0xb8, 0x84, 0x93, 0x8f, 0x2a, 0x41, 0x7f, 0x99, 0x0d, 0xa0, 0xbb, 0x54, 0x09, 0x5e,
	0xc5, 0x6b, 0xa4, 0x5b, 0xa1, 0x2c, 0x31, 0x1b, 0x42, 0xe8, 0xce, 0x77, 0x7b, 0x8b, 0x86, 0x37,
	0x46, 0xc1, 0xb8, 0x27, 0x1f, 0x09, 0x31, 0x82, 0x5e, 0x9e, 0xe8, 0x60, 0xa9, 0x73, 0x08, 0x67,
	0x8b, 0xa2, 0xd0, 0x29, 0x34, 0xd4, 0x82, 0xd4, 0xbe, 0x6c, 0xa8, 0x85, 0xf8, 0x19, 0x40, 0xeb,
	0xf3, 0x16, 0xf5, 0x9e, 0x3d, 0x83, 0x56, 0xa2, 0xd6, 0xca, 0x7a, 0x31, 0x07, 0xec, 0x39, 0xb4,
	0xd3, 0xe5, 0xd2, 0xa0, 0xa5, 0xca, 0x7d, 0xe9, 0x11, 0x63, 0x70, 0x9c, 0xc5, 0x2b, 0xe4, 0x4d,
	0x62, 0xe9, 0xcc, 0xce, 0xa0, 0x9b, 0xa1, 0xfe, 0x4a, 0xfc, 0x31, 0xf1, 0x9d, 0x0c, 0xf5, 0xb5,
	0x93, 0x46, 0x70, 0x72, 0xef, 0xaa, 0xdc, 0x58, 0xad, 0x36, 0x2b, 0xde, 0xa2, 0x11, 0xab, 0x94,
	0xf8, 0xd1, 0x80, 0xfe, 0xb5, 0xf3, 0xb0, 0x9c, 0x84, 0x43, 0x87, 0x4c, 0x9d, 0x4d, 0x7d, 0x4b,
	0x05, 0x64, 0x3d, 0x08, 0x1e, 0x7c, 0x3f, 0xc1, 0x83, 0x43, 0x7b, 0xdf, 0x47, 0x40, 0x63, 0xc4,
	0x9b, 0x55, 0x52, 0x74, 0x90, 0x03, 0xe7, 0xe1, 0x7c, 0xab, 0x77, 0xb1, 0xdd, 0x6a, 0xa4, 0xea,
	0x7d, 0xf9, 0x48, 0x50, 0xa5, 0x64, 0x7e, 0xa1, 0xac, 0xe1, 0x6d, 0x5f, 0x29, 0x87, 0x5e, 0xf9,
	0x92, 0xea, 0x05, 0xef, 0x94, 0x8a, 0x83, 0xee, 0xc5, 0xfc, 0x71, 0xc2, 0xbb, 0x24, 0x95, 0x98,
	0xbd, 0x02, 0x98, 0x7f, 0x8b, 0x37, 0x2b, 0xbc, 0x55, 0x6b, 0xe4, 0x21, 0xa9, 0x15, 0xc6, 0x99,
	0xaa, 0x71, 0x1d, 0xeb, 0xef, 0x1c, 0xc8, 0x08, 0x8f, 0xc4, 0xaf, 0x00, 0x4e, 0xc9, 0x03, 0x53,
	0x9a, 0xf0, 0x1e, 0xda, 0x34, 0xb5, 0x7b, 0xd1, 0xe6, 0xf8, 0x64, 0xf2, 0x26, 0xfa, 0x77, 0xeb,
	0xa2, 0xbf, 0x7c, 0x93, 0x3e, 0xc0, 0x75, 0x61, 0x53, 0x1b, 0x27, 0xee, 0x3d, 0x8c, 0xb7, 0xab,
	0xc2, 0x78, 0x4f, 0x74, 0x2e, 0x37, 0x4b, 0x4f, 0x72, 0x62, 0xf2, 0xbb, 0x01, 0x2d, 0xca, 0xcb,
	0x6e, 0xa1, 0x7f, 0x13, 0xef, 0x90, 0x80, 0x5b, 0x35, 0xf6, 0xba, 0xae, 0x87, 0xca, 0x36, 0x0f,
	0x46, 0x87, 0x2f, 0xe4, 0x3d, 0x8a, 0x23, 0x26, 0xa1, 0x77, 0x89, 0x96, 0x92, 0x5e, 0xec, 0x67,
	0x53, 0xf6, 0xb2, 0x2e, 0xa6, 0xdc, 0xdb, 0xc1, 0xff, 0xe7, 0x16, 0x47, 0xec, 0x13, 0x84, 0x45,
	0x4e, 0xc3, 0xce, 0xea, 0x22, 0x68, 0xd5, 0x07, 0xe2, 0x60, 0x32, 0x53, 0xc9, 0x76, 0x05, 0x4f,
	0xa6, 0x98, 0xa0, 0xc5, 0x0f, 0x49, 0xe2, 0x73, 0x9e, 0xd7, 0x05, 0x16, 0x2d, 0x0e, 0xeb, 0xc5,
	0x22, 0xdf, 0x5d, 0x9b, 0x7e, 0x1b, 0xef, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x96, 0xd3,
	0xd7, 0x51, 0x04, 0x00, 0x00,
}