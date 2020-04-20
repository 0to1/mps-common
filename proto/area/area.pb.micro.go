// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/area/area.proto

package go_micro_srv_area

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AreaService service

type AreaService interface {
	// 添加货位类型
	AddAreaType(ctx context.Context, in *AreaType, opts ...client.CallOption) (*Response, error)
	// 修改货位类型
	UpdateAreaType(ctx context.Context, in *UpdateAreaTypeReq, opts ...client.CallOption) (*Response, error)
	// 删除货位类型
	DeleteAreaType(ctx context.Context, in *AreaType, opts ...client.CallOption) (*Response, error)
	// 获取货位类型
	GetAreaType(ctx context.Context, in *AreaTypeQuery, opts ...client.CallOption) (*AreaTypes, error)
	// 获取货位类型
	GetOneAreaType(ctx context.Context, in *AreaType, opts ...client.CallOption) (*AreaType, error)
	// 添加区域
	AddArea(ctx context.Context, in *Area, opts ...client.CallOption) (*Response, error)
	// 修改区域
	UpdateArea(ctx context.Context, in *UpdateAreaReq, opts ...client.CallOption) (*Response, error)
	// 删除区域
	DeleteArea(ctx context.Context, in *AreaIDReq, opts ...client.CallOption) (*Response, error)
	// 根据ID获取指定区域信息
	GetOneArea(ctx context.Context, in *AreaIDReq, opts ...client.CallOption) (*Area, error)
	// 根据查询条件获取区域
	GetAreas(ctx context.Context, in *AreaQuery, opts ...client.CallOption) (*Areas, error)
	//对应区域设置货位组
	AddRacklots(ctx context.Context, in *RacklotsReq, opts ...client.CallOption) (*Response, error)
	AddChildAreas(ctx context.Context, in *ChildsReq, opts ...client.CallOption) (*Response, error)
	// rpc UpdateArea(UpdateAreaReq) returns (Response);
	// 设置区域类型
	SetAreaType(ctx context.Context, in *TypeReq, opts ...client.CallOption) (*Response, error)
	// 设置区域的任务容量
	SetOrderCapacity(ctx context.Context, in *CapReq, opts ...client.CallOption) (*Response, error)
	// 设置是否允许存车
	Inbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error)
	// 设置是否允许取车
	Outbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error)
	// 设置是否启用
	SetValid(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error)
	//设置父区域
	SetParentArea(ctx context.Context, in *ParentReq, opts ...client.CallOption) (*Response, error)
	//增加属性
	SetProperties(ctx context.Context, in *PropertiesReq, opts ...client.CallOption) (*Response, error)
}

type areaService struct {
	c    client.Client
	name string
}

func NewAreaService(name string, c client.Client) AreaService {
	return &areaService{
		c:    c,
		name: name,
	}
}

func (c *areaService) AddAreaType(ctx context.Context, in *AreaType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.AddAreaType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) UpdateAreaType(ctx context.Context, in *UpdateAreaTypeReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.UpdateAreaType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) DeleteAreaType(ctx context.Context, in *AreaType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.DeleteAreaType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) GetAreaType(ctx context.Context, in *AreaTypeQuery, opts ...client.CallOption) (*AreaTypes, error) {
	req := c.c.NewRequest(c.name, "AreaService.GetAreaType", in)
	out := new(AreaTypes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) GetOneAreaType(ctx context.Context, in *AreaType, opts ...client.CallOption) (*AreaType, error) {
	req := c.c.NewRequest(c.name, "AreaService.GetOneAreaType", in)
	out := new(AreaType)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) AddArea(ctx context.Context, in *Area, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.AddArea", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) UpdateArea(ctx context.Context, in *UpdateAreaReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.UpdateArea", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) DeleteArea(ctx context.Context, in *AreaIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.DeleteArea", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) GetOneArea(ctx context.Context, in *AreaIDReq, opts ...client.CallOption) (*Area, error) {
	req := c.c.NewRequest(c.name, "AreaService.GetOneArea", in)
	out := new(Area)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) GetAreas(ctx context.Context, in *AreaQuery, opts ...client.CallOption) (*Areas, error) {
	req := c.c.NewRequest(c.name, "AreaService.GetAreas", in)
	out := new(Areas)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) AddRacklots(ctx context.Context, in *RacklotsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.AddRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) AddChildAreas(ctx context.Context, in *ChildsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.AddChildAreas", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) SetAreaType(ctx context.Context, in *TypeReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.SetAreaType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) SetOrderCapacity(ctx context.Context, in *CapReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.SetOrderCapacity", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) Inbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.Inbound", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) Outbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.Outbound", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) SetValid(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.SetValid", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) SetParentArea(ctx context.Context, in *ParentReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.SetParentArea", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) SetProperties(ctx context.Context, in *PropertiesReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AreaService.SetProperties", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AreaService service

type AreaServiceHandler interface {
	// 添加货位类型
	AddAreaType(context.Context, *AreaType, *Response) error
	// 修改货位类型
	UpdateAreaType(context.Context, *UpdateAreaTypeReq, *Response) error
	// 删除货位类型
	DeleteAreaType(context.Context, *AreaType, *Response) error
	// 获取货位类型
	GetAreaType(context.Context, *AreaTypeQuery, *AreaTypes) error
	// 获取货位类型
	GetOneAreaType(context.Context, *AreaType, *AreaType) error
	// 添加区域
	AddArea(context.Context, *Area, *Response) error
	// 修改区域
	UpdateArea(context.Context, *UpdateAreaReq, *Response) error
	// 删除区域
	DeleteArea(context.Context, *AreaIDReq, *Response) error
	// 根据ID获取指定区域信息
	GetOneArea(context.Context, *AreaIDReq, *Area) error
	// 根据查询条件获取区域
	GetAreas(context.Context, *AreaQuery, *Areas) error
	//对应区域设置货位组
	AddRacklots(context.Context, *RacklotsReq, *Response) error
	AddChildAreas(context.Context, *ChildsReq, *Response) error
	// rpc UpdateArea(UpdateAreaReq) returns (Response);
	// 设置区域类型
	SetAreaType(context.Context, *TypeReq, *Response) error
	// 设置区域的任务容量
	SetOrderCapacity(context.Context, *CapReq, *Response) error
	// 设置是否允许存车
	Inbound(context.Context, *FlagReq, *Response) error
	// 设置是否允许取车
	Outbound(context.Context, *FlagReq, *Response) error
	// 设置是否启用
	SetValid(context.Context, *FlagReq, *Response) error
	//设置父区域
	SetParentArea(context.Context, *ParentReq, *Response) error
	//增加属性
	SetProperties(context.Context, *PropertiesReq, *Response) error
}

func RegisterAreaServiceHandler(s server.Server, hdlr AreaServiceHandler, opts ...server.HandlerOption) error {
	type areaService interface {
		AddAreaType(ctx context.Context, in *AreaType, out *Response) error
		UpdateAreaType(ctx context.Context, in *UpdateAreaTypeReq, out *Response) error
		DeleteAreaType(ctx context.Context, in *AreaType, out *Response) error
		GetAreaType(ctx context.Context, in *AreaTypeQuery, out *AreaTypes) error
		GetOneAreaType(ctx context.Context, in *AreaType, out *AreaType) error
		AddArea(ctx context.Context, in *Area, out *Response) error
		UpdateArea(ctx context.Context, in *UpdateAreaReq, out *Response) error
		DeleteArea(ctx context.Context, in *AreaIDReq, out *Response) error
		GetOneArea(ctx context.Context, in *AreaIDReq, out *Area) error
		GetAreas(ctx context.Context, in *AreaQuery, out *Areas) error
		AddRacklots(ctx context.Context, in *RacklotsReq, out *Response) error
		AddChildAreas(ctx context.Context, in *ChildsReq, out *Response) error
		SetAreaType(ctx context.Context, in *TypeReq, out *Response) error
		SetOrderCapacity(ctx context.Context, in *CapReq, out *Response) error
		Inbound(ctx context.Context, in *FlagReq, out *Response) error
		Outbound(ctx context.Context, in *FlagReq, out *Response) error
		SetValid(ctx context.Context, in *FlagReq, out *Response) error
		SetParentArea(ctx context.Context, in *ParentReq, out *Response) error
		SetProperties(ctx context.Context, in *PropertiesReq, out *Response) error
	}
	type AreaService struct {
		areaService
	}
	h := &areaServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AreaService{h}, opts...))
}

type areaServiceHandler struct {
	AreaServiceHandler
}

func (h *areaServiceHandler) AddAreaType(ctx context.Context, in *AreaType, out *Response) error {
	return h.AreaServiceHandler.AddAreaType(ctx, in, out)
}

func (h *areaServiceHandler) UpdateAreaType(ctx context.Context, in *UpdateAreaTypeReq, out *Response) error {
	return h.AreaServiceHandler.UpdateAreaType(ctx, in, out)
}

func (h *areaServiceHandler) DeleteAreaType(ctx context.Context, in *AreaType, out *Response) error {
	return h.AreaServiceHandler.DeleteAreaType(ctx, in, out)
}

func (h *areaServiceHandler) GetAreaType(ctx context.Context, in *AreaTypeQuery, out *AreaTypes) error {
	return h.AreaServiceHandler.GetAreaType(ctx, in, out)
}

func (h *areaServiceHandler) GetOneAreaType(ctx context.Context, in *AreaType, out *AreaType) error {
	return h.AreaServiceHandler.GetOneAreaType(ctx, in, out)
}

func (h *areaServiceHandler) AddArea(ctx context.Context, in *Area, out *Response) error {
	return h.AreaServiceHandler.AddArea(ctx, in, out)
}

func (h *areaServiceHandler) UpdateArea(ctx context.Context, in *UpdateAreaReq, out *Response) error {
	return h.AreaServiceHandler.UpdateArea(ctx, in, out)
}

func (h *areaServiceHandler) DeleteArea(ctx context.Context, in *AreaIDReq, out *Response) error {
	return h.AreaServiceHandler.DeleteArea(ctx, in, out)
}

func (h *areaServiceHandler) GetOneArea(ctx context.Context, in *AreaIDReq, out *Area) error {
	return h.AreaServiceHandler.GetOneArea(ctx, in, out)
}

func (h *areaServiceHandler) GetAreas(ctx context.Context, in *AreaQuery, out *Areas) error {
	return h.AreaServiceHandler.GetAreas(ctx, in, out)
}

func (h *areaServiceHandler) AddRacklots(ctx context.Context, in *RacklotsReq, out *Response) error {
	return h.AreaServiceHandler.AddRacklots(ctx, in, out)
}

func (h *areaServiceHandler) AddChildAreas(ctx context.Context, in *ChildsReq, out *Response) error {
	return h.AreaServiceHandler.AddChildAreas(ctx, in, out)
}

func (h *areaServiceHandler) SetAreaType(ctx context.Context, in *TypeReq, out *Response) error {
	return h.AreaServiceHandler.SetAreaType(ctx, in, out)
}

func (h *areaServiceHandler) SetOrderCapacity(ctx context.Context, in *CapReq, out *Response) error {
	return h.AreaServiceHandler.SetOrderCapacity(ctx, in, out)
}

func (h *areaServiceHandler) Inbound(ctx context.Context, in *FlagReq, out *Response) error {
	return h.AreaServiceHandler.Inbound(ctx, in, out)
}

func (h *areaServiceHandler) Outbound(ctx context.Context, in *FlagReq, out *Response) error {
	return h.AreaServiceHandler.Outbound(ctx, in, out)
}

func (h *areaServiceHandler) SetValid(ctx context.Context, in *FlagReq, out *Response) error {
	return h.AreaServiceHandler.SetValid(ctx, in, out)
}

func (h *areaServiceHandler) SetParentArea(ctx context.Context, in *ParentReq, out *Response) error {
	return h.AreaServiceHandler.SetParentArea(ctx, in, out)
}

func (h *areaServiceHandler) SetProperties(ctx context.Context, in *PropertiesReq, out *Response) error {
	return h.AreaServiceHandler.SetProperties(ctx, in, out)
}
