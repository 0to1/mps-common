// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

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

type SignupReq struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Role                 uint32   `protobuf:"varint,5,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupReq) Reset()         { *m = SignupReq{} }
func (m *SignupReq) String() string { return proto.CompactTextString(m) }
func (*SignupReq) ProtoMessage()    {}
func (*SignupReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *SignupReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupReq.Unmarshal(m, b)
}
func (m *SignupReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupReq.Marshal(b, m, deterministic)
}
func (m *SignupReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupReq.Merge(m, src)
}
func (m *SignupReq) XXX_Size() int {
	return xxx_messageInfo_SignupReq.Size(m)
}
func (m *SignupReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignupReq proto.InternalMessageInfo

func (m *SignupReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SignupReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignupReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignupReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupReq) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

type SignupResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupResp) Reset()         { *m = SignupResp{} }
func (m *SignupResp) String() string { return proto.CompactTextString(m) }
func (*SignupResp) ProtoMessage()    {}
func (*SignupResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *SignupResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupResp.Unmarshal(m, b)
}
func (m *SignupResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupResp.Marshal(b, m, deterministic)
}
func (m *SignupResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupResp.Merge(m, src)
}
func (m *SignupResp) XXX_Size() int {
	return xxx_messageInfo_SignupResp.Size(m)
}
func (m *SignupResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignupResp proto.InternalMessageInfo

type SigninReq struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninReq) Reset()         { *m = SigninReq{} }
func (m *SigninReq) String() string { return proto.CompactTextString(m) }
func (*SigninReq) ProtoMessage()    {}
func (*SigninReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *SigninReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninReq.Unmarshal(m, b)
}
func (m *SigninReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninReq.Marshal(b, m, deterministic)
}
func (m *SigninReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninReq.Merge(m, src)
}
func (m *SigninReq) XXX_Size() int {
	return xxx_messageInfo_SigninReq.Size(m)
}
func (m *SigninReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninReq.DiscardUnknown(m)
}

var xxx_messageInfo_SigninReq proto.InternalMessageInfo

func (m *SigninReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SigninReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResp struct {
	UserID               uint32   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=Token,proto3" json:"Token,omitempty"`
	Role                 uint32   `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResp) Reset()         { *m = UserResp{} }
func (m *UserResp) String() string { return proto.CompactTextString(m) }
func (*UserResp) ProtoMessage()    {}
func (*UserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *UserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResp.Unmarshal(m, b)
}
func (m *UserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResp.Marshal(b, m, deterministic)
}
func (m *UserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResp.Merge(m, src)
}
func (m *UserResp) XXX_Size() int {
	return xxx_messageInfo_UserResp.Size(m)
}
func (m *UserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserResp proto.InternalMessageInfo

func (m *UserResp) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserResp) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResp) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserResp) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

type ResetReq struct {
	UserID               uint32   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Oldpassword          string   `protobuf:"bytes,2,opt,name=oldpassword,proto3" json:"oldpassword,omitempty"`
	Newpassword          string   `protobuf:"bytes,3,opt,name=newpassword,proto3" json:"newpassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetReq) Reset()         { *m = ResetReq{} }
func (m *ResetReq) String() string { return proto.CompactTextString(m) }
func (*ResetReq) ProtoMessage()    {}
func (*ResetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *ResetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetReq.Unmarshal(m, b)
}
func (m *ResetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetReq.Marshal(b, m, deterministic)
}
func (m *ResetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetReq.Merge(m, src)
}
func (m *ResetReq) XXX_Size() int {
	return xxx_messageInfo_ResetReq.Size(m)
}
func (m *ResetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetReq proto.InternalMessageInfo

func (m *ResetReq) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ResetReq) GetOldpassword() string {
	if m != nil {
		return m.Oldpassword
	}
	return ""
}

func (m *ResetReq) GetNewpassword() string {
	if m != nil {
		return m.Newpassword
	}
	return ""
}

type ResetResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetResp) Reset()         { *m = ResetResp{} }
func (m *ResetResp) String() string { return proto.CompactTextString(m) }
func (*ResetResp) ProtoMessage()    {}
func (*ResetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *ResetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetResp.Unmarshal(m, b)
}
func (m *ResetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetResp.Marshal(b, m, deterministic)
}
func (m *ResetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetResp.Merge(m, src)
}
func (m *ResetResp) XXX_Size() int {
	return xxx_messageInfo_ResetResp.Size(m)
}
func (m *ResetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetResp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetResp proto.InternalMessageInfo

type SignoutReq struct {
	UserID               uint32   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutReq) Reset()         { *m = SignoutReq{} }
func (m *SignoutReq) String() string { return proto.CompactTextString(m) }
func (*SignoutReq) ProtoMessage()    {}
func (*SignoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *SignoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutReq.Unmarshal(m, b)
}
func (m *SignoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutReq.Marshal(b, m, deterministic)
}
func (m *SignoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutReq.Merge(m, src)
}
func (m *SignoutReq) XXX_Size() int {
	return xxx_messageInfo_SignoutReq.Size(m)
}
func (m *SignoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutReq proto.InternalMessageInfo

func (m *SignoutReq) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type SignoutResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutResp) Reset()         { *m = SignoutResp{} }
func (m *SignoutResp) String() string { return proto.CompactTextString(m) }
func (*SignoutResp) ProtoMessage()    {}
func (*SignoutResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *SignoutResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutResp.Unmarshal(m, b)
}
func (m *SignoutResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutResp.Marshal(b, m, deterministic)
}
func (m *SignoutResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutResp.Merge(m, src)
}
func (m *SignoutResp) XXX_Size() int {
	return xxx_messageInfo_SignoutResp.Size(m)
}
func (m *SignoutResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutResp proto.InternalMessageInfo

type UpdateReq struct {
	UserID               uint32   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=Token,proto3" json:"Token,omitempty"`
	Role                 uint32   `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
	Routes               string   `protobuf:"bytes,7,opt,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{8}
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

func (m *UpdateReq) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UpdateReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdateReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UpdateReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateReq) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *UpdateReq) GetRoutes() string {
	if m != nil {
		return m.Routes
	}
	return ""
}

type TokenReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenReq) Reset()         { *m = TokenReq{} }
func (m *TokenReq) String() string { return proto.CompactTextString(m) }
func (*TokenReq) ProtoMessage()    {}
func (*TokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{9}
}

func (m *TokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenReq.Unmarshal(m, b)
}
func (m *TokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenReq.Marshal(b, m, deterministic)
}
func (m *TokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenReq.Merge(m, src)
}
func (m *TokenReq) XXX_Size() int {
	return xxx_messageInfo_TokenReq.Size(m)
}
func (m *TokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_TokenReq proto.InternalMessageInfo

func (m *TokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type IdReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdReq) Reset()         { *m = IdReq{} }
func (m *IdReq) String() string { return proto.CompactTextString(m) }
func (*IdReq) ProtoMessage()    {}
func (*IdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{10}
}

func (m *IdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdReq.Unmarshal(m, b)
}
func (m *IdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdReq.Marshal(b, m, deterministic)
}
func (m *IdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdReq.Merge(m, src)
}
func (m *IdReq) XXX_Size() int {
	return xxx_messageInfo_IdReq.Size(m)
}
func (m *IdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IdReq.DiscardUnknown(m)
}

var xxx_messageInfo_IdReq proto.InternalMessageInfo

func (m *IdReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResp) Reset()         { *m = DeleteResp{} }
func (m *DeleteResp) String() string { return proto.CompactTextString(m) }
func (*DeleteResp) ProtoMessage()    {}
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{11}
}

func (m *DeleteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResp.Unmarshal(m, b)
}
func (m *DeleteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResp.Marshal(b, m, deterministic)
}
func (m *DeleteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResp.Merge(m, src)
}
func (m *DeleteResp) XXX_Size() int {
	return xxx_messageInfo_DeleteResp.Size(m)
}
func (m *DeleteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResp proto.InternalMessageInfo

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
	return fileDescriptor_9b283a848145d6b7, []int{12}
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

type UsersResp struct {
	Users                []*UserResp `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Totalpages           uint32      `protobuf:"varint,2,opt,name=totalpages,proto3" json:"totalpages,omitempty"`
	Currpages            uint32      `protobuf:"varint,3,opt,name=currpages,proto3" json:"currpages,omitempty"`
	Totals               uint32      `protobuf:"varint,4,opt,name=totals,proto3" json:"totals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UsersResp) Reset()         { *m = UsersResp{} }
func (m *UsersResp) String() string { return proto.CompactTextString(m) }
func (*UsersResp) ProtoMessage()    {}
func (*UsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{13}
}

func (m *UsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResp.Unmarshal(m, b)
}
func (m *UsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResp.Marshal(b, m, deterministic)
}
func (m *UsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResp.Merge(m, src)
}
func (m *UsersResp) XXX_Size() int {
	return xxx_messageInfo_UsersResp.Size(m)
}
func (m *UsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResp proto.InternalMessageInfo

func (m *UsersResp) GetUsers() []*UserResp {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UsersResp) GetTotalpages() uint32 {
	if m != nil {
		return m.Totalpages
	}
	return 0
}

func (m *UsersResp) GetCurrpages() uint32 {
	if m != nil {
		return m.Currpages
	}
	return 0
}

func (m *UsersResp) GetTotals() uint32 {
	if m != nil {
		return m.Totals
	}
	return 0
}

type ConfigReq struct {
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	UserID               uint32   `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Configs              string   `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigReq) Reset()         { *m = ConfigReq{} }
func (m *ConfigReq) String() string { return proto.CompactTextString(m) }
func (*ConfigReq) ProtoMessage()    {}
func (*ConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{14}
}

func (m *ConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigReq.Unmarshal(m, b)
}
func (m *ConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigReq.Marshal(b, m, deterministic)
}
func (m *ConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigReq.Merge(m, src)
}
func (m *ConfigReq) XXX_Size() int {
	return xxx_messageInfo_ConfigReq.Size(m)
}
func (m *ConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigReq proto.InternalMessageInfo

func (m *ConfigReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConfigReq) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ConfigReq) GetConfigs() string {
	if m != nil {
		return m.Configs
	}
	return ""
}

type ConfigResp struct {
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	UserID               uint32   `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Configs              string   `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigResp) Reset()         { *m = ConfigResp{} }
func (m *ConfigResp) String() string { return proto.CompactTextString(m) }
func (*ConfigResp) ProtoMessage()    {}
func (*ConfigResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{15}
}

func (m *ConfigResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResp.Unmarshal(m, b)
}
func (m *ConfigResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResp.Marshal(b, m, deterministic)
}
func (m *ConfigResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResp.Merge(m, src)
}
func (m *ConfigResp) XXX_Size() int {
	return xxx_messageInfo_ConfigResp.Size(m)
}
func (m *ConfigResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResp.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResp proto.InternalMessageInfo

func (m *ConfigResp) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConfigResp) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ConfigResp) GetConfigs() string {
	if m != nil {
		return m.Configs
	}
	return ""
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
	return fileDescriptor_9b283a848145d6b7, []int{16}
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
	proto.RegisterType((*SignupReq)(nil), "go.micro.srv.user.SignupReq")
	proto.RegisterType((*SignupResp)(nil), "go.micro.srv.user.SignupResp")
	proto.RegisterType((*SigninReq)(nil), "go.micro.srv.user.SigninReq")
	proto.RegisterType((*UserResp)(nil), "go.micro.srv.user.UserResp")
	proto.RegisterType((*ResetReq)(nil), "go.micro.srv.user.ResetReq")
	proto.RegisterType((*ResetResp)(nil), "go.micro.srv.user.ResetResp")
	proto.RegisterType((*SignoutReq)(nil), "go.micro.srv.user.SignoutReq")
	proto.RegisterType((*SignoutResp)(nil), "go.micro.srv.user.SignoutResp")
	proto.RegisterType((*UpdateReq)(nil), "go.micro.srv.user.UpdateReq")
	proto.RegisterType((*TokenReq)(nil), "go.micro.srv.user.TokenReq")
	proto.RegisterType((*IdReq)(nil), "go.micro.srv.user.IdReq")
	proto.RegisterType((*DeleteResp)(nil), "go.micro.srv.user.DeleteResp")
	proto.RegisterType((*Query)(nil), "go.micro.srv.user.Query")
	proto.RegisterType((*UsersResp)(nil), "go.micro.srv.user.UsersResp")
	proto.RegisterType((*ConfigReq)(nil), "go.micro.srv.user.ConfigReq")
	proto.RegisterType((*ConfigResp)(nil), "go.micro.srv.user.ConfigResp")
	proto.RegisterType((*Parameter)(nil), "go.micro.srv.user.Parameter")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x8e, 0x93, 0x38, 0x89, 0x4f, 0x9a, 0xea, 0xff, 0x47, 0x05, 0x4c, 0x28, 0x55, 0x34, 0x62,
	0xd1, 0x55, 0x10, 0xed, 0x9e, 0x05, 0x04, 0xaa, 0x14, 0x5a, 0x15, 0x97, 0xae, 0x91, 0xdb, 0x9c,
	0x04, 0xab, 0x89, 0x67, 0x3a, 0x63, 0xb7, 0xea, 0x53, 0xb0, 0xe1, 0x01, 0x78, 0x0d, 0x5e, 0x80,
	0xc7, 0x42, 0x68, 0x2e, 0x76, 0xdc, 0xd6, 0x71, 0x59, 0x54, 0x62, 0x13, 0xf9, 0x5c, 0xe6, 0x9b,
	0xef, 0x9b, 0x73, 0x51, 0xe0, 0x11, 0x17, 0x2c, 0x61, 0x2f, 0x53, 0x89, 0x42, 0xff, 0x0c, 0xb5,
	0x4d, 0xfe, 0x9f, 0xb1, 0xe1, 0x22, 0x3a, 0x13, 0x6c, 0x28, 0xc5, 0xe5, 0x50, 0x05, 0xe8, 0x37,
	0x07, 0xbc, 0xe3, 0x68, 0x16, 0xa7, 0x3c, 0xc0, 0x0b, 0x32, 0x80, 0x2e, 0xff, 0xca, 0x62, 0x3c,
	0x4c, 0x17, 0xa7, 0x28, 0x7c, 0x67, 0xe0, 0x6c, 0x7b, 0x41, 0xd1, 0x45, 0xfa, 0xd0, 0x51, 0xe7,
	0x0e, 0xc3, 0x05, 0xfa, 0x75, 0x1d, 0xce, 0x6d, 0x15, 0xe3, 0xa1, 0x94, 0x57, 0x4c, 0x4c, 0xfc,
	0x86, 0x89, 0x65, 0x36, 0xd9, 0x00, 0x17, 0x17, 0x61, 0x34, 0xf7, 0x9b, 0x3a, 0x60, 0x0c, 0x42,
	0xa0, 0x29, 0xd8, 0x1c, 0x7d, 0x77, 0xe0, 0x6c, 0xf7, 0x02, 0xfd, 0x4d, 0xd7, 0x00, 0x32, 0x42,
	0x92, 0xd3, 0xb1, 0xa1, 0x17, 0xc5, 0x7f, 0x4d, 0x2f, 0xa7, 0x50, 0xbf, 0x49, 0x81, 0xfe, 0x70,
	0xa0, 0x73, 0x22, 0x51, 0x28, 0x5c, 0xf2, 0x18, 0x5a, 0x8a, 0xf7, 0x78, 0xa4, 0x51, 0x7a, 0x81,
	0xb5, 0x2a, 0xf5, 0xe5, 0x1a, 0x1a, 0x45, 0x0d, 0xb7, 0x48, 0x35, 0xef, 0x92, 0xda, 0x00, 0xf7,
	0x33, 0x3b, 0xc7, 0x58, 0xcb, 0xf4, 0x02, 0x63, 0xe4, 0xda, 0x5b, 0x05, 0xed, 0x53, 0xe8, 0x04,
	0x28, 0x31, 0x51, 0x62, 0x57, 0x31, 0x1c, 0x40, 0x97, 0xcd, 0x27, 0xb7, 0x54, 0x16, 0x5d, 0x2a,
	0x23, 0xc6, 0xab, 0x5b, 0xa5, 0x28, 0xba, 0x68, 0x17, 0x3c, 0x7b, 0x8f, 0xe4, 0xf4, 0x85, 0x79,
	0x70, 0x96, 0x56, 0x5d, 0x4b, 0x7b, 0xd0, 0xcd, 0xb3, 0x24, 0xa7, 0x3f, 0x1d, 0xf0, 0x4e, 0xf8,
	0x24, 0x4c, 0xb0, 0x8a, 0xeb, 0x3f, 0x7d, 0x4d, 0xc5, 0x4a, 0xb0, 0x34, 0x41, 0xe9, 0xb7, 0x75,
	0xaa, 0xb5, 0xe8, 0x00, 0x3a, 0xfa, 0x90, 0x62, 0xbe, 0x01, 0x6e, 0xa2, 0xd1, 0x4c, 0x33, 0x19,
	0x83, 0x3e, 0x01, 0x77, 0x3c, 0x51, 0xe1, 0x75, 0xa8, 0x47, 0x13, 0x2b, 0xaa, 0x1e, 0x4d, 0x54,
	0x73, 0x8e, 0x70, 0x8e, 0x4a, 0xb5, 0xe4, 0xf4, 0x97, 0x03, 0xee, 0xa7, 0x14, 0xc5, 0xb5, 0x82,
	0x99, 0x47, 0x8b, 0x28, 0xb1, 0xa9, 0xc6, 0x50, 0x04, 0xd8, 0x74, 0x2a, 0x31, 0xd1, 0xe2, 0x7b,
	0x81, 0xb5, 0x14, 0x59, 0x1e, 0xce, 0x50, 0x2b, 0xef, 0x05, 0xfa, 0x9b, 0x3c, 0x85, 0x0e, 0x47,
	0xf1, 0x45, 0xfb, 0x9b, 0xda, 0xdf, 0xe6, 0x28, 0x8e, 0x54, 0x68, 0x00, 0xdd, 0x0b, 0x75, 0xcb,
	0x71, 0x22, 0xa2, 0x78, 0x66, 0x75, 0x17, 0x5d, 0xe4, 0xb5, 0xcd, 0x38, 0x0a, 0x45, 0xb8, 0x90,
	0x7e, 0x6b, 0xd0, 0xd8, 0xee, 0xee, 0x6c, 0x0e, 0xef, 0x8c, 0xfb, 0x50, 0x27, 0x60, 0x82, 0x22,
	0x28, 0x1e, 0xa0, 0xdf, 0x55, 0x35, 0x25, 0x0a, 0xa9, 0x67, 0xe3, 0x15, 0xb8, 0x2a, 0x59, 0xfa,
	0x8e, 0xc6, 0x79, 0x56, 0x82, 0x93, 0xcd, 0x51, 0x60, 0x32, 0xc9, 0x16, 0x40, 0xc2, 0x92, 0x70,
	0xae, 0xe8, 0x4b, 0xab, 0xb6, 0xe0, 0x21, 0x9b, 0xe0, 0x9d, 0xa5, 0x42, 0x98, 0xb0, 0x91, 0xbd,
	0x74, 0xa8, 0x77, 0xd2, 0xb9, 0xd2, 0x2a, 0xb7, 0x16, 0x3d, 0x00, 0xef, 0x2d, 0x8b, 0xa7, 0xd1,
	0x6c, 0x59, 0x8a, 0x46, 0x56, 0x8a, 0x42, 0xcf, 0xd5, 0x6f, 0xf4, 0x9c, 0x0f, 0xed, 0x33, 0x7d,
	0x48, 0xda, 0x9a, 0x66, 0x26, 0x3d, 0x04, 0xc8, 0xe0, 0x24, 0x7f, 0x00, 0xbc, 0x5d, 0xf0, 0xf2,
	0xf7, 0x24, 0xff, 0x41, 0xe3, 0x1c, 0xaf, 0x6d, 0x8a, 0xfa, 0x54, 0x3d, 0x71, 0x19, 0xce, 0xd3,
	0xac, 0xf3, 0x8d, 0xb1, 0xf3, 0xbb, 0x05, 0x4d, 0xf5, 0x7a, 0x64, 0x0f, 0x5a, 0x66, 0xcf, 0x91,
	0xb2, 0x42, 0xe5, 0x3b, 0xb9, 0xff, 0xbc, 0x22, 0x2a, 0x39, 0xad, 0x91, 0x77, 0x06, 0x28, 0x8a,
	0x57, 0x02, 0xe9, 0xed, 0xd9, 0xaf, 0xaa, 0x23, 0xad, 0x91, 0x7d, 0x68, 0xdb, 0x01, 0x27, 0xab,
	0xae, 0x34, 0x2b, 0xa2, 0xbf, 0x55, 0x15, 0xd6, 0x58, 0x23, 0x70, 0xf5, 0x7e, 0x21, 0x65, 0x77,
	0x66, 0x1b, 0xae, 0xbf, 0xb9, 0x3a, 0x98, 0x09, 0x33, 0x2b, 0xa6, 0x54, 0x58, 0xbe, 0x7d, 0xee,
	0x13, 0xf6, 0x11, 0xd6, 0xf7, 0x30, 0x51, 0x8e, 0x37, 0xd7, 0x66, 0x59, 0x94, 0x1d, 0xc8, 0x36,
	0xc2, 0x7d, 0x68, 0x63, 0x58, 0x37, 0x1b, 0xc0, 0x00, 0xaa, 0x06, 0x29, 0x39, 0xa0, 0xb7, 0x47,
	0x69, 0xe9, 0x96, 0xeb, 0x83, 0xbc, 0x87, 0x6e, 0x4e, 0xac, 0x12, 0xe7, 0x1e, 0x4a, 0x23, 0xe8,
	0x58, 0x1c, 0x59, 0x0a, 0xa2, 0x57, 0x54, 0xe9, 0x6b, 0xe7, 0x33, 0x4f, 0x6b, 0xe4, 0x03, 0xc0,
	0x71, 0x78, 0x89, 0x66, 0x42, 0x4a, 0x5f, 0x3c, 0x9f, 0xc5, 0x52, 0x61, 0xcb, 0xd1, 0xa2, 0x35,
	0x72, 0x00, 0x6b, 0x46, 0xe8, 0xc3, 0xc0, 0xed, 0x83, 0xb7, 0x87, 0xc9, 0x83, 0x60, 0x9d, 0xb6,
	0xf4, 0x7f, 0xa1, 0xdd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x5a, 0xcd, 0x54, 0x24, 0x09,
	0x00, 0x00,
}
