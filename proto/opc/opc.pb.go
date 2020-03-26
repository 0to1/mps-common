// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/opc/opc.proto

package go_micro_srv_opc

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

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{0}
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

type SrvReq struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrvReq) Reset()         { *m = SrvReq{} }
func (m *SrvReq) String() string { return proto.CompactTextString(m) }
func (*SrvReq) ProtoMessage()    {}
func (*SrvReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{1}
}

func (m *SrvReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvReq.Unmarshal(m, b)
}
func (m *SrvReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvReq.Marshal(b, m, deterministic)
}
func (m *SrvReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvReq.Merge(m, src)
}
func (m *SrvReq) XXX_Size() int {
	return xxx_messageInfo_SrvReq.Size(m)
}
func (m *SrvReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvReq.DiscardUnknown(m)
}

var xxx_messageInfo_SrvReq proto.InternalMessageInfo

func (m *SrvReq) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

type SrvQuery struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrvQuery) Reset()         { *m = SrvQuery{} }
func (m *SrvQuery) String() string { return proto.CompactTextString(m) }
func (*SrvQuery) ProtoMessage()    {}
func (*SrvQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{2}
}

func (m *SrvQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvQuery.Unmarshal(m, b)
}
func (m *SrvQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvQuery.Marshal(b, m, deterministic)
}
func (m *SrvQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvQuery.Merge(m, src)
}
func (m *SrvQuery) XXX_Size() int {
	return xxx_messageInfo_SrvQuery.Size(m)
}
func (m *SrvQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SrvQuery proto.InternalMessageInfo

func (m *SrvQuery) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *SrvQuery) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SrvQuery) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type FileReq struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	GroupID              uint32   `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	Filebytes            []byte   `protobuf:"bytes,3,opt,name=filebytes,proto3" json:"filebytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileReq) Reset()         { *m = FileReq{} }
func (m *FileReq) String() string { return proto.CompactTextString(m) }
func (*FileReq) ProtoMessage()    {}
func (*FileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{3}
}

func (m *FileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileReq.Unmarshal(m, b)
}
func (m *FileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileReq.Marshal(b, m, deterministic)
}
func (m *FileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileReq.Merge(m, src)
}
func (m *FileReq) XXX_Size() int {
	return xxx_messageInfo_FileReq.Size(m)
}
func (m *FileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FileReq.DiscardUnknown(m)
}

var xxx_messageInfo_FileReq proto.InternalMessageInfo

func (m *FileReq) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *FileReq) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *FileReq) GetFilebytes() []byte {
	if m != nil {
		return m.Filebytes
	}
	return nil
}

type GrpReq struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	GroupID              uint32   `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpReq) Reset()         { *m = GrpReq{} }
func (m *GrpReq) String() string { return proto.CompactTextString(m) }
func (*GrpReq) ProtoMessage()    {}
func (*GrpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{4}
}

func (m *GrpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpReq.Unmarshal(m, b)
}
func (m *GrpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpReq.Marshal(b, m, deterministic)
}
func (m *GrpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpReq.Merge(m, src)
}
func (m *GrpReq) XXX_Size() int {
	return xxx_messageInfo_GrpReq.Size(m)
}
func (m *GrpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpReq.DiscardUnknown(m)
}

var xxx_messageInfo_GrpReq proto.InternalMessageInfo

func (m *GrpReq) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *GrpReq) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

type GrpQuery struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	Page                 uint32   `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32   `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpQuery) Reset()         { *m = GrpQuery{} }
func (m *GrpQuery) String() string { return proto.CompactTextString(m) }
func (*GrpQuery) ProtoMessage()    {}
func (*GrpQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{5}
}

func (m *GrpQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpQuery.Unmarshal(m, b)
}
func (m *GrpQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpQuery.Marshal(b, m, deterministic)
}
func (m *GrpQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpQuery.Merge(m, src)
}
func (m *GrpQuery) XXX_Size() int {
	return xxx_messageInfo_GrpQuery.Size(m)
}
func (m *GrpQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GrpQuery proto.InternalMessageInfo

func (m *GrpQuery) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *GrpQuery) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GrpQuery) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type Group struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	GroupID              uint32   `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{6}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *Group) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Groups struct {
	Group                []*Group `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty"`
	TotalCount           uint32   `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Groups) Reset()         { *m = Groups{} }
func (m *Groups) String() string { return proto.CompactTextString(m) }
func (*Groups) ProtoMessage()    {}
func (*Groups) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{7}
}

func (m *Groups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Groups.Unmarshal(m, b)
}
func (m *Groups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Groups.Marshal(b, m, deterministic)
}
func (m *Groups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Groups.Merge(m, src)
}
func (m *Groups) XXX_Size() int {
	return xxx_messageInfo_Groups.Size(m)
}
func (m *Groups) XXX_DiscardUnknown() {
	xxx_messageInfo_Groups.DiscardUnknown(m)
}

var xxx_messageInfo_Groups proto.InternalMessageInfo

func (m *Groups) GetGroup() []*Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *Groups) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type ItemReq struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	GroupID              uint32   `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	ItemID               uint32   `protobuf:"varint,3,opt,name=itemID,proto3" json:"itemID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemReq) Reset()         { *m = ItemReq{} }
func (m *ItemReq) String() string { return proto.CompactTextString(m) }
func (*ItemReq) ProtoMessage()    {}
func (*ItemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{8}
}

func (m *ItemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemReq.Unmarshal(m, b)
}
func (m *ItemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemReq.Marshal(b, m, deterministic)
}
func (m *ItemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemReq.Merge(m, src)
}
func (m *ItemReq) XXX_Size() int {
	return xxx_messageInfo_ItemReq.Size(m)
}
func (m *ItemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemReq.DiscardUnknown(m)
}

var xxx_messageInfo_ItemReq proto.InternalMessageInfo

func (m *ItemReq) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *ItemReq) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *ItemReq) GetItemID() uint32 {
	if m != nil {
		return m.ItemID
	}
	return 0
}

type ItemValueReq struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	GroupID              uint32   `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	ItemID               uint32   `protobuf:"varint,3,opt,name=itemID,proto3" json:"itemID,omitempty"`
	Value                uint32   `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemValueReq) Reset()         { *m = ItemValueReq{} }
func (m *ItemValueReq) String() string { return proto.CompactTextString(m) }
func (*ItemValueReq) ProtoMessage()    {}
func (*ItemValueReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{9}
}

func (m *ItemValueReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemValueReq.Unmarshal(m, b)
}
func (m *ItemValueReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemValueReq.Marshal(b, m, deterministic)
}
func (m *ItemValueReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemValueReq.Merge(m, src)
}
func (m *ItemValueReq) XXX_Size() int {
	return xxx_messageInfo_ItemValueReq.Size(m)
}
func (m *ItemValueReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemValueReq.DiscardUnknown(m)
}

var xxx_messageInfo_ItemValueReq proto.InternalMessageInfo

func (m *ItemValueReq) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *ItemValueReq) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *ItemValueReq) GetItemID() uint32 {
	if m != nil {
		return m.ItemID
	}
	return 0
}

func (m *ItemValueReq) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ItemValueResp struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Quality              uint32   `protobuf:"varint,2,opt,name=quality,proto3" json:"quality,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemValueResp) Reset()         { *m = ItemValueResp{} }
func (m *ItemValueResp) String() string { return proto.CompactTextString(m) }
func (*ItemValueResp) ProtoMessage()    {}
func (*ItemValueResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{10}
}

func (m *ItemValueResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemValueResp.Unmarshal(m, b)
}
func (m *ItemValueResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemValueResp.Marshal(b, m, deterministic)
}
func (m *ItemValueResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemValueResp.Merge(m, src)
}
func (m *ItemValueResp) XXX_Size() int {
	return xxx_messageInfo_ItemValueResp.Size(m)
}
func (m *ItemValueResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemValueResp.DiscardUnknown(m)
}

var xxx_messageInfo_ItemValueResp proto.InternalMessageInfo

func (m *ItemValueResp) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ItemValueResp) GetQuality() uint32 {
	if m != nil {
		return m.Quality
	}
	return 0
}

type ItemQuery struct {
	Page                 uint32      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32      `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Filter               *ItemFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ItemQuery) Reset()         { *m = ItemQuery{} }
func (m *ItemQuery) String() string { return proto.CompactTextString(m) }
func (*ItemQuery) ProtoMessage()    {}
func (*ItemQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{11}
}

func (m *ItemQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemQuery.Unmarshal(m, b)
}
func (m *ItemQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemQuery.Marshal(b, m, deterministic)
}
func (m *ItemQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemQuery.Merge(m, src)
}
func (m *ItemQuery) XXX_Size() int {
	return xxx_messageInfo_ItemQuery.Size(m)
}
func (m *ItemQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ItemQuery proto.InternalMessageInfo

func (m *ItemQuery) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ItemQuery) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *ItemQuery) GetFilter() *ItemFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type ItemFilter struct {
	And                  []*ItemFilter `protobuf:"bytes,1,rep,name=and,proto3" json:"and,omitempty"`
	Or                   []*ItemFilter `protobuf:"bytes,2,rep,name=or,proto3" json:"or,omitempty"`
	GarageIn             []uint32      `protobuf:"varint,3,rep,packed,name=garageIn,proto3" json:"garageIn,omitempty"`
	GroupIn              []uint32      `protobuf:"varint,4,rep,packed,name=groupIn,proto3" json:"groupIn,omitempty"`
	ItemIn               []uint32      `protobuf:"varint,5,rep,packed,name=itemIn,proto3" json:"itemIn,omitempty"`
	NameLike             []string      `protobuf:"bytes,6,rep,name=nameLike,proto3" json:"nameLike,omitempty"`
	AddressLike          []string      `protobuf:"bytes,7,rep,name=addressLike,proto3" json:"addressLike,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ItemFilter) Reset()         { *m = ItemFilter{} }
func (m *ItemFilter) String() string { return proto.CompactTextString(m) }
func (*ItemFilter) ProtoMessage()    {}
func (*ItemFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{12}
}

func (m *ItemFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemFilter.Unmarshal(m, b)
}
func (m *ItemFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemFilter.Marshal(b, m, deterministic)
}
func (m *ItemFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemFilter.Merge(m, src)
}
func (m *ItemFilter) XXX_Size() int {
	return xxx_messageInfo_ItemFilter.Size(m)
}
func (m *ItemFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ItemFilter proto.InternalMessageInfo

func (m *ItemFilter) GetAnd() []*ItemFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *ItemFilter) GetOr() []*ItemFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *ItemFilter) GetGarageIn() []uint32 {
	if m != nil {
		return m.GarageIn
	}
	return nil
}

func (m *ItemFilter) GetGroupIn() []uint32 {
	if m != nil {
		return m.GroupIn
	}
	return nil
}

func (m *ItemFilter) GetItemIn() []uint32 {
	if m != nil {
		return m.ItemIn
	}
	return nil
}

func (m *ItemFilter) GetNameLike() []string {
	if m != nil {
		return m.NameLike
	}
	return nil
}

func (m *ItemFilter) GetAddressLike() []string {
	if m != nil {
		return m.AddressLike
	}
	return nil
}

type Item struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	GroupID              uint32   `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	ItemID               uint32   `protobuf:"varint,3,opt,name=itemID,proto3" json:"itemID,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Datatype             string   `protobuf:"bytes,6,opt,name=datatype,proto3" json:"datatype,omitempty"`
	RespectDataType      uint32   `protobuf:"varint,7,opt,name=respectDataType,proto3" json:"respectDataType,omitempty"`
	ClientAccess         string   `protobuf:"bytes,8,opt,name=clientAccess,proto3" json:"clientAccess,omitempty"`
	ScanRate             uint32   `protobuf:"varint,9,opt,name=scanRate,proto3" json:"scanRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{13}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *Item) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *Item) GetItemID() uint32 {
	if m != nil {
		return m.ItemID
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Item) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *Item) GetRespectDataType() uint32 {
	if m != nil {
		return m.RespectDataType
	}
	return 0
}

func (m *Item) GetClientAccess() string {
	if m != nil {
		return m.ClientAccess
	}
	return ""
}

func (m *Item) GetScanRate() uint32 {
	if m != nil {
		return m.ScanRate
	}
	return 0
}

type Items struct {
	Item                 []*Item  `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	TotalCount           uint32   `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Items) Reset()         { *m = Items{} }
func (m *Items) String() string { return proto.CompactTextString(m) }
func (*Items) ProtoMessage()    {}
func (*Items) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{14}
}

func (m *Items) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Items.Unmarshal(m, b)
}
func (m *Items) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Items.Marshal(b, m, deterministic)
}
func (m *Items) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Items.Merge(m, src)
}
func (m *Items) XXX_Size() int {
	return xxx_messageInfo_Items.Size(m)
}
func (m *Items) XXX_DiscardUnknown() {
	xxx_messageInfo_Items.DiscardUnknown(m)
}

var xxx_messageInfo_Items proto.InternalMessageInfo

func (m *Items) GetItem() []*Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Items) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type Server struct {
	GarageID             uint32   `protobuf:"varint,1,opt,name=garageID,proto3" json:"garageID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Tag                  string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{15}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetGarageID() uint32 {
	if m != nil {
		return m.GarageID
	}
	return 0
}

func (m *Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Server) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Server) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Servers struct {
	Server               []*Server `protobuf:"bytes,1,rep,name=server,proto3" json:"server,omitempty"`
	TotalCount           uint32    `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Servers) Reset()         { *m = Servers{} }
func (m *Servers) String() string { return proto.CompactTextString(m) }
func (*Servers) ProtoMessage()    {}
func (*Servers) Descriptor() ([]byte, []int) {
	return fileDescriptor_9947f1df27adff2f, []int{16}
}

func (m *Servers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Servers.Unmarshal(m, b)
}
func (m *Servers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Servers.Marshal(b, m, deterministic)
}
func (m *Servers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Servers.Merge(m, src)
}
func (m *Servers) XXX_Size() int {
	return xxx_messageInfo_Servers.Size(m)
}
func (m *Servers) XXX_DiscardUnknown() {
	xxx_messageInfo_Servers.DiscardUnknown(m)
}

var xxx_messageInfo_Servers proto.InternalMessageInfo

func (m *Servers) GetServer() []*Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *Servers) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Response)(nil), "go.micro.srv.opc.Response")
	proto.RegisterType((*SrvReq)(nil), "go.micro.srv.opc.SrvReq")
	proto.RegisterType((*SrvQuery)(nil), "go.micro.srv.opc.SrvQuery")
	proto.RegisterType((*FileReq)(nil), "go.micro.srv.opc.FileReq")
	proto.RegisterType((*GrpReq)(nil), "go.micro.srv.opc.GrpReq")
	proto.RegisterType((*GrpQuery)(nil), "go.micro.srv.opc.GrpQuery")
	proto.RegisterType((*Group)(nil), "go.micro.srv.opc.Group")
	proto.RegisterType((*Groups)(nil), "go.micro.srv.opc.Groups")
	proto.RegisterType((*ItemReq)(nil), "go.micro.srv.opc.ItemReq")
	proto.RegisterType((*ItemValueReq)(nil), "go.micro.srv.opc.ItemValueReq")
	proto.RegisterType((*ItemValueResp)(nil), "go.micro.srv.opc.ItemValueResp")
	proto.RegisterType((*ItemQuery)(nil), "go.micro.srv.opc.ItemQuery")
	proto.RegisterType((*ItemFilter)(nil), "go.micro.srv.opc.ItemFilter")
	proto.RegisterType((*Item)(nil), "go.micro.srv.opc.Item")
	proto.RegisterType((*Items)(nil), "go.micro.srv.opc.Items")
	proto.RegisterType((*Server)(nil), "go.micro.srv.opc.Server")
	proto.RegisterType((*Servers)(nil), "go.micro.srv.opc.Servers")
}

func init() { proto.RegisterFile("proto/opc/opc.proto", fileDescriptor_9947f1df27adff2f) }

var fileDescriptor_9947f1df27adff2f = []byte{
	// 846 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x80, 0x15, 0x27, 0xb1, 0xe3, 0xb3, 0x59, 0xa8, 0x06, 0xd4, 0x4e, 0x4d, 0x55, 0x22, 0x8b,
	0x8b, 0x08, 0x41, 0x40, 0x0b, 0x97, 0xc0, 0x36, 0x34, 0xda, 0x28, 0x12, 0x12, 0xe0, 0x14, 0x56,
	0x02, 0x71, 0x31, 0xb5, 0xa7, 0x91, 0x85, 0x63, 0xcf, 0xce, 0x4c, 0x22, 0xe5, 0x0d, 0x78, 0x31,
	0x5e, 0x89, 0x6b, 0x34, 0x3f, 0x4e, 0xdc, 0xae, 0xed, 0x54, 0xc9, 0x5e, 0x58, 0x9a, 0x73, 0xe6,
	0xf8, 0x3b, 0x7f, 0xf3, 0x07, 0x1f, 0x31, 0x5e, 0xc8, 0xe2, 0xab, 0x82, 0xc5, 0xea, 0x9b, 0x68,
	0x09, 0x3d, 0x5a, 0x15, 0x93, 0x75, 0x1a, 0xf3, 0x62, 0x22, 0xf8, 0x76, 0x52, 0xb0, 0x38, 0x04,
	0x18, 0x44, 0x54, 0xb0, 0x22, 0x17, 0x34, 0xfc, 0x0c, 0xdc, 0x25, 0xdf, 0x46, 0xf4, 0x0e, 0x05,
	0x30, 0x58, 0x11, 0x4e, 0x56, 0x74, 0x31, 0xc3, 0x9d, 0x51, 0x67, 0x7c, 0x19, 0xed, 0xe5, 0xf0,
	0x15, 0x0c, 0x96, 0x7c, 0xfb, 0xeb, 0x86, 0xf2, 0x5d, 0x9b, 0x1d, 0x42, 0xd0, 0x63, 0x64, 0x45,
	0xb1, 0xa3, 0xf5, 0x7a, 0x8c, 0x30, 0x78, 0x8c, 0xf2, 0x5f, 0x94, 0xba, 0xab, 0xd5, 0xa5, 0x18,
	0xfe, 0x05, 0xde, 0x4d, 0x9a, 0xd1, 0x23, 0xce, 0x15, 0x60, 0xc5, 0x8b, 0x0d, 0x5b, 0xcc, 0x2c,
	0xb7, 0x14, 0xd1, 0x33, 0xf0, 0xdf, 0xa4, 0x19, 0x7d, 0xbd, 0x93, 0x54, 0x68, 0xf8, 0x30, 0x3a,
	0x28, 0xc2, 0x1f, 0xc0, 0x9d, 0x73, 0x76, 0x32, 0x5d, 0x25, 0x3d, 0xe7, 0xec, 0xa1, 0x93, 0x8e,
	0xa1, 0x3f, 0x57, 0x0e, 0x4e, 0x4c, 0x19, 0x41, 0x2f, 0x27, 0x6b, 0x43, 0xf5, 0x23, 0x3d, 0x46,
	0x8f, 0xa0, 0x2b, 0xc9, 0x0a, 0xf7, 0xb4, 0x4a, 0x0d, 0xc3, 0x5b, 0x95, 0x7a, 0xb1, 0x61, 0x02,
	0x7d, 0x09, 0x7d, 0xfd, 0x2b, 0xee, 0x8c, 0xba, 0xe3, 0x8b, 0xab, 0x27, 0x93, 0x77, 0x57, 0xc3,
	0x44, 0x1b, 0x46, 0xc6, 0x0a, 0x3d, 0x07, 0x90, 0x85, 0x24, 0xd9, 0xcb, 0x62, 0x93, 0x4b, 0xeb,
	0xbb, 0xa2, 0x09, 0x6f, 0xc1, 0x5b, 0x48, 0xba, 0x3e, 0xbd, 0x65, 0x8f, 0xc1, 0x4d, 0x25, 0x5d,
	0x2f, 0x66, 0xb6, 0x2e, 0x56, 0x0a, 0x39, 0x0c, 0x15, 0xf8, 0x77, 0x92, 0x6d, 0xe8, 0x83, 0xd3,
	0xd1, 0xc7, 0xd0, 0xdf, 0x2a, 0xb2, 0xae, 0xd1, 0x65, 0x64, 0x84, 0xf0, 0x1a, 0x2e, 0x2b, 0x3e,
	0x05, 0x3b, 0x98, 0x75, 0x2a, 0x66, 0xca, 0xdd, 0xdd, 0x86, 0x64, 0xa9, 0xdc, 0x95, 0xee, 0xac,
	0x18, 0x16, 0xe0, 0x2b, 0x80, 0x59, 0x22, 0xe5, 0x32, 0xe8, 0xd4, 0x2f, 0x03, 0xe7, 0xad, 0x65,
	0x80, 0xbe, 0x05, 0xf7, 0x4d, 0x9a, 0x49, 0xca, 0x75, 0xa4, 0x17, 0x57, 0xcf, 0xee, 0x37, 0x46,
	0xa1, 0x6f, 0xb4, 0x4d, 0x64, 0x6d, 0xc3, 0xff, 0x3a, 0x00, 0x07, 0x35, 0x9a, 0x40, 0x97, 0xe4,
	0x89, 0x6d, 0x6d, 0x3b, 0x41, 0x19, 0xa2, 0x2f, 0xc0, 0x29, 0x38, 0x76, 0xde, 0xc3, 0xdc, 0x29,
	0x78, 0xa5, 0x05, 0x39, 0xee, 0x8e, 0xba, 0x95, 0x16, 0xe4, 0x87, 0x16, 0xe4, 0xb8, 0xa7, 0xa7,
	0x4a, 0x71, 0xdf, 0x82, 0x1c, 0xf7, 0xf5, 0x84, 0x95, 0x14, 0x4d, 0x2d, 0xd6, 0x9f, 0xd2, 0xbf,
	0x29, 0x76, 0x47, 0xdd, 0xb1, 0x1f, 0xed, 0x65, 0x34, 0x82, 0x0b, 0x92, 0x24, 0x9c, 0x0a, 0xa1,
	0xa7, 0x3d, 0x3d, 0x5d, 0x55, 0x85, 0xff, 0x38, 0xd0, 0x53, 0xe1, 0x3d, 0xf0, 0xba, 0x28, 0x77,
	0x53, 0xaf, 0xb2, 0x9b, 0x30, 0x78, 0xd6, 0x33, 0xee, 0x6b, 0x75, 0x29, 0x2a, 0xdf, 0x09, 0x91,
	0x44, 0xee, 0x98, 0x4a, 0x41, 0x4d, 0xed, 0x65, 0x34, 0x86, 0x0f, 0x39, 0x15, 0x8c, 0xc6, 0x72,
	0x46, 0x24, 0x79, 0xa5, 0x4c, 0x3c, 0xed, 0xea, 0x5d, 0x35, 0x0a, 0x61, 0x18, 0x67, 0x29, 0xcd,
	0xe5, 0x34, 0x8e, 0x95, 0x93, 0x81, 0x26, 0xbd, 0xa5, 0x53, 0x9e, 0x44, 0x4c, 0xf2, 0x88, 0x48,
	0x8a, 0x7d, 0x93, 0x65, 0x29, 0x87, 0x4b, 0xe8, 0xab, 0x4a, 0x08, 0xf4, 0x39, 0xf4, 0x54, 0x1a,
	0xb6, 0xfd, 0x8f, 0xeb, 0xfb, 0x19, 0x69, 0x9b, 0xa3, 0xfb, 0xfa, 0x0f, 0x70, 0x97, 0x94, 0x6f,
	0x29, 0x3f, 0x76, 0xd2, 0xe9, 0x72, 0x39, 0x95, 0x72, 0x7d, 0x00, 0x4e, 0xca, 0xec, 0x71, 0xe4,
	0xa4, 0xac, 0xe6, 0x30, 0xfa, 0x13, 0x3c, 0xc3, 0x16, 0xe8, 0x6b, 0x70, 0x85, 0x1e, 0xda, 0xa0,
	0xf1, 0xfd, 0xa0, 0x8d, 0x69, 0x64, 0xed, 0x8e, 0x05, 0x7e, 0xf5, 0xaf, 0x0f, 0xf0, 0x33, 0x8b,
	0xd5, 0x5f, 0x69, 0x4c, 0xd1, 0x14, 0xfc, 0x69, 0x92, 0xd8, 0x54, 0x1a, 0xe9, 0x41, 0x70, 0x7f,
	0xa6, 0xbc, 0x11, 0xd1, 0x0c, 0x86, 0xbf, 0xb1, 0x84, 0x48, 0x7a, 0x2e, 0x65, 0x46, 0x33, 0x7a,
	0x26, 0xe5, 0x1a, 0xfc, 0x39, 0x95, 0x2d, 0x08, 0x7d, 0x75, 0x07, 0x8d, 0x70, 0xf4, 0x12, 0x60,
	0x0f, 0x10, 0x28, 0xa8, 0x25, 0xe8, 0xe3, 0x2b, 0x78, 0xda, 0xc4, 0x10, 0xe8, 0x7b, 0x18, 0xcc,
	0xa9, 0x34, 0xb7, 0x16, 0xae, 0xbb, 0x40, 0xd4, 0x25, 0x1b, 0x34, 0x5d, 0x2d, 0xaa, 0x27, 0xe5,
	0xef, 0xb5, 0x21, 0x94, 0x97, 0x6c, 0x80, 0x1b, 0x08, 0x02, 0x5d, 0xc3, 0x60, 0x9a, 0x24, 0x06,
	0xd7, 0xe4, 0xa7, 0xb5, 0x90, 0x3f, 0xc2, 0x85, 0x69, 0xea, 0x79, 0x0c, 0xd3, 0xd2, 0x33, 0x18,
	0xdf, 0x81, 0x37, 0xa7, 0x52, 0x9f, 0x64, 0x4f, 0x1b, 0x36, 0x2c, 0xbd, 0x0b, 0x1a, 0xf6, 0x32,
	0x7a, 0xa1, 0x1b, 0x61, 0x76, 0xff, 0x27, 0xf5, 0x36, 0xa6, 0x92, 0x4f, 0xea, 0x27, 0x55, 0x2b,
	0xbd, 0x69, 0x92, 0x68, 0x58, 0x83, 0x93, 0xd6, 0xf0, 0x5f, 0x00, 0x98, 0x32, 0x9e, 0x43, 0x30,
	0x45, 0x3c, 0x99, 0x30, 0x07, 0xff, 0x96, 0xa7, 0x16, 0xf0, 0xbc, 0x1e, 0x50, 0x3e, 0x23, 0x5a,
	0x41, 0x37, 0xea, 0x19, 0x4c, 0x92, 0x63, 0xcd, 0xf8, 0xb4, 0xd5, 0x85, 0x60, 0x6a, 0x8f, 0x99,
	0xa2, 0xa8, 0xc7, 0x6c, 0x1d, 0xc9, 0x3e, 0x72, 0xdb, 0x82, 0x79, 0xed, 0xea, 0xc7, 0xfa, 0x37,
	0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x22, 0x25, 0xa8, 0x3c, 0xc3, 0x0b, 0x00, 0x00,
}
