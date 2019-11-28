// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rack/rack.proto

package go_micro_srv_rack

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

type RackType struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RackType) Reset()         { *m = RackType{} }
func (m *RackType) String() string { return proto.CompactTextString(m) }
func (*RackType) ProtoMessage()    {}
func (*RackType) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{0}
}

func (m *RackType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RackType.Unmarshal(m, b)
}
func (m *RackType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RackType.Marshal(b, m, deterministic)
}
func (m *RackType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RackType.Merge(m, src)
}
func (m *RackType) XXX_Size() int {
	return xxx_messageInfo_RackType.Size(m)
}
func (m *RackType) XXX_DiscardUnknown() {
	xxx_messageInfo_RackType.DiscardUnknown(m)
}

var xxx_messageInfo_RackType proto.InternalMessageInfo

func (m *RackType) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RackType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Cell 储位
type Cell struct {
	// 储位ID
	Id uint32 `protobuf:"varint,21,opt,name=id,proto3" json:"id,omitempty"`
	// 储位编码
	Tag string `protobuf:"bytes,22,opt,name=tag,proto3" json:"tag,omitempty"`
	// 储位名称
	Name string `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// 储位所在货架
	RackID uint32 `protobuf:"varint,24,opt,name=rackID,proto3" json:"rackID,omitempty"`
	Length int32  `protobuf:"varint,25,opt,name=length,proto3" json:"length,omitempty"`
	Width  int32  `protobuf:"varint,26,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,27,opt,name=height,proto3" json:"height,omitempty"`
	// 储位在货架的哪一层
	Layer int32 `protobuf:"varint,28,opt,name=layer,proto3" json:"layer,omitempty"`
	// 储位在货架上的编号
	CellCode int32 `protobuf:"varint,29,opt,name=cellCode,proto3" json:"cellCode,omitempty"`
	// 是否禁用
	IsValid bool `protobuf:"varint,30,opt,name=isValid,proto3" json:"isValid,omitempty"`
	// 存放的物料对于的编号
	MaterialID uint32 `protobuf:"varint,31,opt,name=materialID,proto3" json:"materialID,omitempty"`
	// cell类型id
	Type uint32 `protobuf:"varint,32,opt,name=type,proto3" json:"type,omitempty"`
	//是否被占用
	IsOccupied           bool     `protobuf:"varint,33,opt,name=isOccupied,proto3" json:"isOccupied,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{1}
}

func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (m *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(m, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

func (m *Cell) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cell) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Cell) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cell) GetRackID() uint32 {
	if m != nil {
		return m.RackID
	}
	return 0
}

func (m *Cell) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Cell) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Cell) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Cell) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *Cell) GetCellCode() int32 {
	if m != nil {
		return m.CellCode
	}
	return 0
}

func (m *Cell) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *Cell) GetMaterialID() uint32 {
	if m != nil {
		return m.MaterialID
	}
	return 0
}

func (m *Cell) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Cell) GetIsOccupied() bool {
	if m != nil {
		return m.IsOccupied
	}
	return false
}

type Rack struct {
	// 货架编码
	Id   uint32 `protobuf:"varint,20,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,30,opt,name=name,proto3" json:"name,omitempty"`
	// 货架类型
	Type *RackType `protobuf:"bytes,21,opt,name=type,proto3" json:"type,omitempty"`
	// 是否禁用
	IsValid bool  `protobuf:"varint,22,opt,name=isValid,proto3" json:"isValid,omitempty"`
	Length  int32 `protobuf:"varint,23,opt,name=length,proto3" json:"length,omitempty"`
	Width   int32 `protobuf:"varint,24,opt,name=width,proto3" json:"width,omitempty"`
	Height  int32 `protobuf:"varint,25,opt,name=height,proto3" json:"height,omitempty"`
	// 所在货位
	Racklot int32 `protobuf:"varint,26,opt,name=racklot,proto3" json:"racklot,omitempty"`
	// 货架层数
	Layers int32 `protobuf:"varint,27,opt,name=layers,proto3" json:"layers,omitempty"`
	// 货架状态(空、可用、被占用、满等)
	Status int32 `protobuf:"varint,31,opt,name=status,proto3" json:"status,omitempty"`
	// 拣货面(A、B两面或者只有一面)
	PickingSurface int32  `protobuf:"varint,32,opt,name=pickingSurface,proto3" json:"pickingSurface,omitempty"`
	Description    string `protobuf:"bytes,28,opt,name=description,proto3" json:"description,omitempty"`
	// 货架储位
	Cells                []*Cell  `protobuf:"bytes,29,rep,name=cells,proto3" json:"cells,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rack) Reset()         { *m = Rack{} }
func (m *Rack) String() string { return proto.CompactTextString(m) }
func (*Rack) ProtoMessage()    {}
func (*Rack) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{2}
}

func (m *Rack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rack.Unmarshal(m, b)
}
func (m *Rack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rack.Marshal(b, m, deterministic)
}
func (m *Rack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rack.Merge(m, src)
}
func (m *Rack) XXX_Size() int {
	return xxx_messageInfo_Rack.Size(m)
}
func (m *Rack) XXX_DiscardUnknown() {
	xxx_messageInfo_Rack.DiscardUnknown(m)
}

var xxx_messageInfo_Rack proto.InternalMessageInfo

func (m *Rack) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rack) GetType() *RackType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Rack) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *Rack) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Rack) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Rack) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Rack) GetRacklot() int32 {
	if m != nil {
		return m.Racklot
	}
	return 0
}

func (m *Rack) GetLayers() int32 {
	if m != nil {
		return m.Layers
	}
	return 0
}

func (m *Rack) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Rack) GetPickingSurface() int32 {
	if m != nil {
		return m.PickingSurface
	}
	return 0
}

func (m *Rack) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Rack) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{3}
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

type RackIDReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RackIDReq) Reset()         { *m = RackIDReq{} }
func (m *RackIDReq) String() string { return proto.CompactTextString(m) }
func (*RackIDReq) ProtoMessage()    {}
func (*RackIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{4}
}

func (m *RackIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RackIDReq.Unmarshal(m, b)
}
func (m *RackIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RackIDReq.Marshal(b, m, deterministic)
}
func (m *RackIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RackIDReq.Merge(m, src)
}
func (m *RackIDReq) XXX_Size() int {
	return xxx_messageInfo_RackIDReq.Size(m)
}
func (m *RackIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RackIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_RackIDReq proto.InternalMessageInfo

func (m *RackIDReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateRackReq struct {
	Id                   uint32    `protobuf:"varint,20,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	Type                 *RackType `protobuf:"bytes,22,opt,name=type,proto3" json:"type,omitempty"`
	IsValid              bool      `protobuf:"varint,23,opt,name=isValid,proto3" json:"isValid,omitempty"`
	Length               int32     `protobuf:"varint,24,opt,name=length,proto3" json:"length,omitempty"`
	Width                int32     `protobuf:"varint,25,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32     `protobuf:"varint,26,opt,name=height,proto3" json:"height,omitempty"`
	Layers               int32     `protobuf:"varint,27,opt,name=layers,proto3" json:"layers,omitempty"`
	PickingSurface       int32     `protobuf:"varint,28,opt,name=pickingSurface,proto3" json:"pickingSurface,omitempty"`
	Description          string    `protobuf:"bytes,29,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateRackReq) Reset()         { *m = UpdateRackReq{} }
func (m *UpdateRackReq) String() string { return proto.CompactTextString(m) }
func (*UpdateRackReq) ProtoMessage()    {}
func (*UpdateRackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{5}
}

func (m *UpdateRackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRackReq.Unmarshal(m, b)
}
func (m *UpdateRackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRackReq.Marshal(b, m, deterministic)
}
func (m *UpdateRackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRackReq.Merge(m, src)
}
func (m *UpdateRackReq) XXX_Size() int {
	return xxx_messageInfo_UpdateRackReq.Size(m)
}
func (m *UpdateRackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRackReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRackReq proto.InternalMessageInfo

func (m *UpdateRackReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRackReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRackReq) GetType() *RackType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *UpdateRackReq) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *UpdateRackReq) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *UpdateRackReq) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *UpdateRackReq) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *UpdateRackReq) GetLayers() int32 {
	if m != nil {
		return m.Layers
	}
	return 0
}

func (m *UpdateRackReq) GetPickingSurface() int32 {
	if m != nil {
		return m.PickingSurface
	}
	return 0
}

func (m *UpdateRackReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

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
	return fileDescriptor_80f24b60e0ed6c04, []int{6}
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

type GraphqlRacks struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlRacks) Reset()         { *m = GraphqlRacks{} }
func (m *GraphqlRacks) String() string { return proto.CompactTextString(m) }
func (*GraphqlRacks) ProtoMessage()    {}
func (*GraphqlRacks) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{7}
}

func (m *GraphqlRacks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlRacks.Unmarshal(m, b)
}
func (m *GraphqlRacks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlRacks.Marshal(b, m, deterministic)
}
func (m *GraphqlRacks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlRacks.Merge(m, src)
}
func (m *GraphqlRacks) XXX_Size() int {
	return xxx_messageInfo_GraphqlRacks.Size(m)
}
func (m *GraphqlRacks) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlRacks.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlRacks proto.InternalMessageInfo

func (m *GraphqlRacks) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type Query struct {
	Limit                uint32       `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32       `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Page                 uint32       `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32       `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	QueryString          string       `protobuf:"bytes,5,opt,name=queryString,proto3" json:"queryString,omitempty"`
	QueryParams          []*Parameter `protobuf:"bytes,6,rep,name=queryParams,proto3" json:"queryParams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{8}
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

func (m *Query) GetQueryParams() []*Parameter {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

type Racks struct {
	Racks                []*Rack  `protobuf:"bytes,1,rep,name=racks,proto3" json:"racks,omitempty"`
	Totals               uint32   `protobuf:"varint,4,opt,name=totals,proto3" json:"totals,omitempty"`
	Totalpages           uint32   `protobuf:"varint,2,opt,name=totalpages,proto3" json:"totalpages,omitempty"`
	Currpages            uint32   `protobuf:"varint,3,opt,name=currpages,proto3" json:"currpages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Racks) Reset()         { *m = Racks{} }
func (m *Racks) String() string { return proto.CompactTextString(m) }
func (*Racks) ProtoMessage()    {}
func (*Racks) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{9}
}

func (m *Racks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Racks.Unmarshal(m, b)
}
func (m *Racks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Racks.Marshal(b, m, deterministic)
}
func (m *Racks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Racks.Merge(m, src)
}
func (m *Racks) XXX_Size() int {
	return xxx_messageInfo_Racks.Size(m)
}
func (m *Racks) XXX_DiscardUnknown() {
	xxx_messageInfo_Racks.DiscardUnknown(m)
}

var xxx_messageInfo_Racks proto.InternalMessageInfo

func (m *Racks) GetRacks() []*Rack {
	if m != nil {
		return m.Racks
	}
	return nil
}

func (m *Racks) GetTotals() uint32 {
	if m != nil {
		return m.Totals
	}
	return 0
}

func (m *Racks) GetTotalpages() uint32 {
	if m != nil {
		return m.Totalpages
	}
	return 0
}

func (m *Racks) GetCurrpages() uint32 {
	if m != nil {
		return m.Currpages
	}
	return 0
}

type TypeReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeReq) Reset()         { *m = TypeReq{} }
func (m *TypeReq) String() string { return proto.CompactTextString(m) }
func (*TypeReq) ProtoMessage()    {}
func (*TypeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{10}
}

func (m *TypeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeReq.Unmarshal(m, b)
}
func (m *TypeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeReq.Marshal(b, m, deterministic)
}
func (m *TypeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeReq.Merge(m, src)
}
func (m *TypeReq) XXX_Size() int {
	return xxx_messageInfo_TypeReq.Size(m)
}
func (m *TypeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeReq.DiscardUnknown(m)
}

var xxx_messageInfo_TypeReq proto.InternalMessageInfo

func (m *TypeReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TypeReq) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type LayerNumsReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerNumsReq) Reset()         { *m = LayerNumsReq{} }
func (m *LayerNumsReq) String() string { return proto.CompactTextString(m) }
func (*LayerNumsReq) ProtoMessage()    {}
func (*LayerNumsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{11}
}

func (m *LayerNumsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LayerNumsReq.Unmarshal(m, b)
}
func (m *LayerNumsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LayerNumsReq.Marshal(b, m, deterministic)
}
func (m *LayerNumsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LayerNumsReq.Merge(m, src)
}
func (m *LayerNumsReq) XXX_Size() int {
	return xxx_messageInfo_LayerNumsReq.Size(m)
}
func (m *LayerNumsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LayerNumsReq.DiscardUnknown(m)
}

var xxx_messageInfo_LayerNumsReq proto.InternalMessageInfo

func (m *LayerNumsReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RackIDsReq struct {
	Ids                  []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RackIDsReq) Reset()         { *m = RackIDsReq{} }
func (m *RackIDsReq) String() string { return proto.CompactTextString(m) }
func (*RackIDsReq) ProtoMessage()    {}
func (*RackIDsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{12}
}

func (m *RackIDsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RackIDsReq.Unmarshal(m, b)
}
func (m *RackIDsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RackIDsReq.Marshal(b, m, deterministic)
}
func (m *RackIDsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RackIDsReq.Merge(m, src)
}
func (m *RackIDsReq) XXX_Size() int {
	return xxx_messageInfo_RackIDsReq.Size(m)
}
func (m *RackIDsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RackIDsReq.DiscardUnknown(m)
}

var xxx_messageInfo_RackIDsReq proto.InternalMessageInfo

func (m *RackIDsReq) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type CellsReq struct {
	RackID               uint32   `protobuf:"varint,1,opt,name=rackID,proto3" json:"rackID,omitempty"`
	Cells                []*Cell  `protobuf:"bytes,2,rep,name=cells,proto3" json:"cells,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CellsReq) Reset()         { *m = CellsReq{} }
func (m *CellsReq) String() string { return proto.CompactTextString(m) }
func (*CellsReq) ProtoMessage()    {}
func (*CellsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{13}
}

func (m *CellsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellsReq.Unmarshal(m, b)
}
func (m *CellsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellsReq.Marshal(b, m, deterministic)
}
func (m *CellsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellsReq.Merge(m, src)
}
func (m *CellsReq) XXX_Size() int {
	return xxx_messageInfo_CellsReq.Size(m)
}
func (m *CellsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CellsReq.DiscardUnknown(m)
}

var xxx_messageInfo_CellsReq proto.InternalMessageInfo

func (m *CellsReq) GetRackID() uint32 {
	if m != nil {
		return m.RackID
	}
	return 0
}

func (m *CellsReq) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type CellIDsReq struct {
	RackID               uint32   `protobuf:"varint,1,opt,name=rackID,proto3" json:"rackID,omitempty"`
	CellIDs              []uint32 `protobuf:"varint,2,rep,packed,name=cellIDs,proto3" json:"cellIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CellIDsReq) Reset()         { *m = CellIDsReq{} }
func (m *CellIDsReq) String() string { return proto.CompactTextString(m) }
func (*CellIDsReq) ProtoMessage()    {}
func (*CellIDsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{14}
}

func (m *CellIDsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellIDsReq.Unmarshal(m, b)
}
func (m *CellIDsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellIDsReq.Marshal(b, m, deterministic)
}
func (m *CellIDsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellIDsReq.Merge(m, src)
}
func (m *CellIDsReq) XXX_Size() int {
	return xxx_messageInfo_CellIDsReq.Size(m)
}
func (m *CellIDsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CellIDsReq.DiscardUnknown(m)
}

var xxx_messageInfo_CellIDsReq proto.InternalMessageInfo

func (m *CellIDsReq) GetRackID() uint32 {
	if m != nil {
		return m.RackID
	}
	return 0
}

func (m *CellIDsReq) GetCellIDs() []uint32 {
	if m != nil {
		return m.CellIDs
	}
	return nil
}

type RacklotReq struct {
	RackID               uint32   `protobuf:"varint,1,opt,name=rackID,proto3" json:"rackID,omitempty"`
	RacklotID            uint32   `protobuf:"varint,2,opt,name=racklotID,proto3" json:"racklotID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RacklotReq) Reset()         { *m = RacklotReq{} }
func (m *RacklotReq) String() string { return proto.CompactTextString(m) }
func (*RacklotReq) ProtoMessage()    {}
func (*RacklotReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{15}
}

func (m *RacklotReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RacklotReq.Unmarshal(m, b)
}
func (m *RacklotReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RacklotReq.Marshal(b, m, deterministic)
}
func (m *RacklotReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RacklotReq.Merge(m, src)
}
func (m *RacklotReq) XXX_Size() int {
	return xxx_messageInfo_RacklotReq.Size(m)
}
func (m *RacklotReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RacklotReq.DiscardUnknown(m)
}

var xxx_messageInfo_RacklotReq proto.InternalMessageInfo

func (m *RacklotReq) GetRackID() uint32 {
	if m != nil {
		return m.RackID
	}
	return 0
}

func (m *RacklotReq) GetRacklotID() uint32 {
	if m != nil {
		return m.RacklotID
	}
	return 0
}

type MaterialReq struct {
	RackID               uint32   `protobuf:"varint,1,opt,name=rackID,proto3" json:"rackID,omitempty"`
	CellID               uint32   `protobuf:"varint,2,opt,name=cellID,proto3" json:"cellID,omitempty"`
	MaterialID           uint32   `protobuf:"varint,3,opt,name=materialID,proto3" json:"materialID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaterialReq) Reset()         { *m = MaterialReq{} }
func (m *MaterialReq) String() string { return proto.CompactTextString(m) }
func (*MaterialReq) ProtoMessage()    {}
func (*MaterialReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{16}
}

func (m *MaterialReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaterialReq.Unmarshal(m, b)
}
func (m *MaterialReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaterialReq.Marshal(b, m, deterministic)
}
func (m *MaterialReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaterialReq.Merge(m, src)
}
func (m *MaterialReq) XXX_Size() int {
	return xxx_messageInfo_MaterialReq.Size(m)
}
func (m *MaterialReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MaterialReq.DiscardUnknown(m)
}

var xxx_messageInfo_MaterialReq proto.InternalMessageInfo

func (m *MaterialReq) GetRackID() uint32 {
	if m != nil {
		return m.RackID
	}
	return 0
}

func (m *MaterialReq) GetCellID() uint32 {
	if m != nil {
		return m.CellID
	}
	return 0
}

func (m *MaterialReq) GetMaterialID() uint32 {
	if m != nil {
		return m.MaterialID
	}
	return 0
}

type Parameter struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_80f24b60e0ed6c04, []int{17}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RackType)(nil), "go.micro.srv.rack.RackType")
	proto.RegisterType((*Cell)(nil), "go.micro.srv.rack.Cell")
	proto.RegisterType((*Rack)(nil), "go.micro.srv.rack.Rack")
	proto.RegisterType((*Response)(nil), "go.micro.srv.rack.Response")
	proto.RegisterType((*RackIDReq)(nil), "go.micro.srv.rack.RackIDReq")
	proto.RegisterType((*UpdateRackReq)(nil), "go.micro.srv.rack.UpdateRackReq")
	proto.RegisterType((*GraphqlQuery)(nil), "go.micro.srv.rack.GraphqlQuery")
	proto.RegisterType((*GraphqlRacks)(nil), "go.micro.srv.rack.GraphqlRacks")
	proto.RegisterType((*Query)(nil), "go.micro.srv.rack.Query")
	proto.RegisterType((*Racks)(nil), "go.micro.srv.rack.Racks")
	proto.RegisterType((*TypeReq)(nil), "go.micro.srv.rack.TypeReq")
	proto.RegisterType((*LayerNumsReq)(nil), "go.micro.srv.rack.LayerNumsReq")
	proto.RegisterType((*RackIDsReq)(nil), "go.micro.srv.rack.RackIDsReq")
	proto.RegisterType((*CellsReq)(nil), "go.micro.srv.rack.CellsReq")
	proto.RegisterType((*CellIDsReq)(nil), "go.micro.srv.rack.CellIDsReq")
	proto.RegisterType((*RacklotReq)(nil), "go.micro.srv.rack.RacklotReq")
	proto.RegisterType((*MaterialReq)(nil), "go.micro.srv.rack.MaterialReq")
	proto.RegisterType((*Parameter)(nil), "go.micro.srv.rack.Parameter")
}

func init() { proto.RegisterFile("proto/rack/rack.proto", fileDescriptor_80f24b60e0ed6c04) }

var fileDescriptor_80f24b60e0ed6c04 = []byte{
	// 1052 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xef, 0x6e, 0xe2, 0x46,
	0x10, 0x17, 0x10, 0x02, 0x8c, 0x43, 0x74, 0xb7, 0xba, 0x90, 0x0d, 0x01, 0x8e, 0x5a, 0x55, 0x85,
	0x2a, 0x85, 0x93, 0x72, 0xdf, 0x53, 0x5d, 0x42, 0x45, 0xe9, 0xf5, 0x7a, 0x77, 0x4e, 0xd3, 0x6f,
	0x55, 0xe5, 0xd8, 0x1b, 0x58, 0xc5, 0xd8, 0x8e, 0x77, 0x49, 0xc5, 0x3b, 0xdc, 0x73, 0xf4, 0x21,
	0xfa, 0x00, 0x7d, 0xac, 0xaa, 0xda, 0x3f, 0xfe, 0x43, 0xb0, 0x39, 0x94, 0xe4, 0x4b, 0xb4, 0x33,
	0xfb, 0xdb, 0xd9, 0x9d, 0xf9, 0xcd, 0xfc, 0x1c, 0xe0, 0x20, 0x8c, 0x02, 0x1e, 0xbc, 0x89, 0x6c,
	0xe7, 0x56, 0xfe, 0x19, 0x4a, 0x1b, 0xbd, 0x9c, 0x06, 0xc3, 0x39, 0x75, 0xa2, 0x60, 0xc8, 0xa2,
	0xfb, 0xa1, 0xd8, 0x30, 0x87, 0x50, 0xb7, 0x6c, 0xe7, 0xf6, 0xb7, 0x65, 0x48, 0xd0, 0x3e, 0x94,
	0xa9, 0x8b, 0x4b, 0xfd, 0xd2, 0xa0, 0x69, 0x95, 0xa9, 0x8b, 0x10, 0xec, 0xf8, 0xf6, 0x9c, 0xe0,
	0x72, 0xbf, 0x34, 0x68, 0x58, 0x72, 0x6d, 0xfe, 0x53, 0x86, 0x9d, 0x0b, 0xe2, 0x79, 0x1a, 0x7c,
	0x90, 0x80, 0x5f, 0x40, 0x85, 0xdb, 0x53, 0xdc, 0x92, 0x58, 0xb1, 0x4c, 0x8e, 0x1f, 0xa6, 0xc7,
	0x51, 0x0b, 0x76, 0xc5, 0xb5, 0x93, 0x11, 0xc6, 0xf2, 0xa4, 0xb6, 0x84, 0xdf, 0x23, 0xfe, 0x94,
	0xcf, 0xf0, 0x51, 0xbf, 0x34, 0xa8, 0x5a, 0xda, 0x42, 0xaf, 0xa0, 0xfa, 0x17, 0x75, 0xf9, 0x0c,
	0xb7, 0xa5, 0x5b, 0x19, 0x02, 0x3d, 0x23, 0x74, 0x3a, 0xe3, 0xf8, 0x58, 0xa1, 0x95, 0x25, 0xd0,
	0x9e, 0xbd, 0x24, 0x11, 0xee, 0x28, 0xb4, 0x34, 0x50, 0x1b, 0xea, 0x0e, 0xf1, 0xbc, 0x8b, 0xc0,
	0x25, 0xb8, 0x2b, 0x37, 0x12, 0x1b, 0x61, 0xa8, 0x51, 0xf6, 0xbb, 0xed, 0x51, 0x17, 0xf7, 0xfa,
	0xa5, 0x41, 0xdd, 0x8a, 0x4d, 0xd4, 0x03, 0x98, 0xdb, 0x9c, 0x44, 0xd4, 0xf6, 0x26, 0x23, 0xfc,
	0x5a, 0xbe, 0x36, 0xe3, 0x11, 0xd9, 0xf1, 0x65, 0x48, 0x70, 0x5f, 0xee, 0xc8, 0xb5, 0x38, 0x43,
	0xd9, 0x47, 0xc7, 0x59, 0x84, 0x94, 0xb8, 0xf8, 0x1b, 0x19, 0x30, 0xe3, 0x31, 0xff, 0x2b, 0xc3,
	0x8e, 0xa8, 0xb6, 0x2e, 0xde, 0xab, 0xb5, 0x4a, 0xf7, 0x32, 0xa5, 0x7a, 0xa3, 0x2f, 0x10, 0x25,
	0x36, 0x4e, 0x8f, 0x87, 0x6b, 0xdc, 0x0d, 0x63, 0xe2, 0xf4, 0xed, 0x99, 0x5c, 0x5a, 0xab, 0xb9,
	0xa4, 0xd5, 0x3d, 0xcc, 0xaf, 0x2e, 0xce, 0xaf, 0xee, 0xd1, 0x4a, 0x75, 0x31, 0xd4, 0xc4, 0xb5,
	0x5e, 0xc0, 0x35, 0x1b, 0xb1, 0x29, 0xe3, 0x8b, 0x52, 0xb3, 0x98, 0x0f, 0x65, 0x09, 0x3f, 0xe3,
	0x36, 0x5f, 0x30, 0x59, 0xbf, 0xaa, 0xa5, 0x2d, 0xf4, 0x1d, 0xec, 0x87, 0xd4, 0xb9, 0xa5, 0xfe,
	0xf4, 0x72, 0x11, 0xdd, 0xd8, 0x8e, 0xaa, 0x62, 0xd5, 0x7a, 0xe0, 0x45, 0x7d, 0x30, 0x5c, 0xc2,
	0x9c, 0x88, 0x86, 0x9c, 0x06, 0xbe, 0x64, 0xb5, 0x61, 0x65, 0x5d, 0xe8, 0x04, 0xaa, 0x82, 0x4b,
	0x86, 0xbb, 0xfd, 0xca, 0xc0, 0x38, 0x3d, 0xcc, 0xa9, 0x92, 0xe8, 0x56, 0x4b, 0xa1, 0x4c, 0x80,
	0xba, 0x45, 0x58, 0x18, 0xf8, 0x8c, 0x98, 0xc7, 0xd0, 0xb0, 0x64, 0xf3, 0x59, 0xe4, 0xee, 0x61,
	0xeb, 0x9b, 0x7f, 0x97, 0xa1, 0x79, 0x15, 0xba, 0x36, 0x27, 0x02, 0x93, 0x22, 0xd6, 0x29, 0x3b,
	0xc8, 0xa1, 0xac, 0xf5, 0x08, 0xca, 0x0e, 0x8b, 0x28, 0xc3, 0xf9, 0x94, 0x1d, 0xe5, 0x53, 0xd6,
	0x5e, 0xa1, 0xac, 0x88, 0x98, 0x75, 0x02, 0x3a, 0xdb, 0x10, 0xd0, 0x5d, 0x23, 0xc0, 0xfc, 0x16,
	0xf6, 0xc6, 0x91, 0x1d, 0xce, 0xee, 0xbc, 0xcf, 0x0b, 0x12, 0x2d, 0xc5, 0xfb, 0xee, 0xc4, 0x42,
	0xd6, 0xb2, 0x61, 0x29, 0xc3, 0xfc, 0x3e, 0x41, 0x89, 0x02, 0x30, 0x31, 0x92, 0x91, 0xe6, 0x41,
	0x03, 0x13, 0xdb, 0xfc, 0xb7, 0x04, 0xd5, 0x24, 0x96, 0x47, 0xe7, 0x94, 0x6b, 0x5e, 0x94, 0x21,
	0x72, 0x0a, 0x6e, 0x6e, 0x18, 0xe1, 0x52, 0x97, 0x9a, 0x96, 0xb6, 0x04, 0x21, 0xa1, 0x3d, 0x25,
	0xb8, 0xa2, 0x06, 0x52, 0xac, 0xd1, 0x11, 0xd4, 0x43, 0x12, 0xfd, 0x29, 0xfd, 0x3b, 0xd2, 0x5f,
	0x0b, 0x49, 0xf4, 0x49, 0x6c, 0xf5, 0xc1, 0x90, 0x6f, 0xbb, 0xe4, 0x11, 0xf5, 0xa7, 0xb8, 0xaa,
	0x52, 0xcb, 0xb8, 0xd0, 0x99, 0x46, 0x7c, 0xb2, 0x23, 0x7b, 0xce, 0xf0, 0xae, 0xec, 0xb0, 0x4e,
	0x0e, 0xa9, 0x12, 0x40, 0x38, 0x89, 0xac, 0xec, 0x01, 0xf3, 0x4b, 0x09, 0xaa, 0x2a, 0xdd, 0x13,
	0xa8, 0x0a, 0x20, 0xc3, 0xa5, 0xc2, 0x2e, 0x95, 0x6d, 0xa6, 0x50, 0x22, 0x43, 0x1e, 0x70, 0xdb,
	0x63, 0xfa, 0xcd, 0xda, 0x12, 0xf2, 0x22, 0x57, 0x22, 0x1d, 0xa6, 0xb3, 0xcf, 0x78, 0x50, 0x07,
	0x1a, 0xce, 0x22, 0x8a, 0xd4, 0xb6, 0x2a, 0x43, 0xea, 0x30, 0x4f, 0xa0, 0x26, 0x3b, 0x6f, 0xbd,
	0xdb, 0x13, 0x2d, 0x2b, 0xa7, 0x5a, 0x66, 0xf6, 0x60, 0xef, 0x17, 0xd1, 0x2c, 0xbf, 0x2e, 0xe6,
	0x2c, 0x6f, 0x42, 0x7a, 0x00, 0x6a, 0x7c, 0xe4, 0xee, 0x0b, 0xa8, 0x50, 0x57, 0xe5, 0xd7, 0xb4,
	0xc4, 0xd2, 0xfc, 0x0c, 0x75, 0x31, 0x79, 0x72, 0x37, 0x55, 0xfd, 0xd2, 0x8a, 0xea, 0x27, 0xd3,
	0x5b, 0xde, 0x6a, 0x7a, 0xcf, 0x00, 0x84, 0xa9, 0xaf, 0x2c, 0x0a, 0x8a, 0xa1, 0xe6, 0x28, 0x94,
	0x0c, 0xdb, 0xb4, 0x62, 0xd3, 0x3c, 0x57, 0x4f, 0xf6, 0x02, 0xbe, 0xe9, 0x7c, 0x07, 0x1a, 0x5a,
	0xd7, 0x26, 0x23, 0x5d, 0x91, 0xd4, 0x61, 0xfe, 0x01, 0xc6, 0x07, 0xfd, 0x11, 0xd8, 0x14, 0xa4,
	0x05, 0xbb, 0xea, 0xd6, 0xb8, 0x49, 0x95, 0xf5, 0xe0, 0xab, 0x52, 0x79, 0xf8, 0x55, 0x31, 0xdf,
	0x42, 0x23, 0xe9, 0x26, 0x51, 0xd4, 0x5b, 0x12, 0x4f, 0x92, 0x58, 0x8a, 0x89, 0xb8, 0xb7, 0xbd,
	0x45, 0xfc, 0x49, 0x56, 0xc6, 0xe9, 0x17, 0x03, 0x0c, 0x91, 0xd8, 0x25, 0x89, 0xee, 0xa9, 0x43,
	0xd0, 0x0f, 0x50, 0x7b, 0xe7, 0xba, 0xf2, 0x43, 0x53, 0xd4, 0x6a, 0xed, 0x5c, 0x71, 0xd2, 0x23,
	0x88, 0xc6, 0x00, 0x23, 0xe2, 0x11, 0x25, 0x7e, 0xa8, 0x53, 0x10, 0x43, 0x2a, 0xe7, 0xe6, 0x40,
	0xef, 0x01, 0x52, 0x15, 0x45, 0xfd, 0x1c, 0xe8, 0x8a, 0xc8, 0x6e, 0x0e, 0x76, 0x01, 0x30, 0x26,
	0xfc, 0xa3, 0xbf, 0xcd, 0xab, 0x8a, 0xf2, 0x46, 0x67, 0x50, 0x1f, 0x13, 0xae, 0xc6, 0x12, 0xe7,
	0x80, 0xa4, 0xf2, 0xb4, 0x71, 0xc1, 0x71, 0x86, 0xae, 0xe0, 0x65, 0x7c, 0xfe, 0x7c, 0xa9, 0x35,
	0x0d, 0xbd, 0xce, 0x81, 0x67, 0x55, 0xb1, 0xbd, 0x01, 0xa0, 0xc2, 0x8e, 0x60, 0x3f, 0x0d, 0x2b,
	0xff, 0x19, 0x7b, 0xcc, 0xe3, 0x26, 0x60, 0xfc, 0xe8, 0xdb, 0xd7, 0x1e, 0x51, 0x66, 0xb7, 0xb0,
	0x44, 0xec, 0xab, 0xc5, 0xfe, 0x19, 0xf6, 0x46, 0x94, 0x3d, 0x4f, 0xac, 0x31, 0x40, 0xfa, 0xac,
	0xa7, 0xb4, 0xd3, 0x4f, 0x60, 0x64, 0x1e, 0xf5, 0x94, 0x48, 0x13, 0x30, 0xce, 0xa9, 0xef, 0x6a,
	0x39, 0x28, 0xcc, 0x4e, 0x49, 0xc5, 0xd7, 0x7a, 0xbc, 0x79, 0xe5, 0x5f, 0x3f, 0x53, 0xb0, 0x11,
	0xd4, 0xdf, 0xb9, 0xae, 0x14, 0x4e, 0x74, 0x5c, 0x20, 0x87, 0x6c, 0x9b, 0xec, 0x2c, 0x32, 0x0f,
	0xee, 0x89, 0x0a, 0xd4, 0x2d, 0x08, 0xb4, 0x0d, 0x77, 0x49, 0x4b, 0x3d, 0x3d, 0x54, 0xda, 0x52,
	0x4f, 0x8f, 0xf5, 0x1e, 0xf6, 0x04, 0x7f, 0xb1, 0x14, 0xa3, 0x5e, 0x0e, 0x38, 0xa3, 0xd3, 0x9b,
	0x83, 0x7d, 0x80, 0x7d, 0xc5, 0xe0, 0xb3, 0x84, 0xbb, 0xde, 0x95, 0x3f, 0xb6, 0xde, 0xfe, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0x4e, 0x2d, 0xa7, 0xe6, 0x85, 0x0d, 0x00, 0x00,
}
