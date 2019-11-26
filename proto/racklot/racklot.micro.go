// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/racklot/racklot.proto

package go_micro_srv_racklot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for RacklotService service

type RacklotService interface {
	// 添加货位类型
	AddRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*AddResp, error)
	// 修改货位类型
	UpdateRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*UpdateResp, error)
	// 删除货位类型
	DeleteRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*DeleteResp, error)
	// 获取货位类型
	GetRacklotType(ctx context.Context, in *Query, opts ...client.CallOption) (*RacklotTypes, error)
	// 添加货位
	AddRacklot(ctx context.Context, in *Racklot, opts ...client.CallOption) (*AddResp, error)
	// 导入货位
	BatchAddRacklots(ctx context.Context, in *Racklots, opts ...client.CallOption) (*Response, error)
	// 删除货位
	DeleteRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*DeleteResp, error)
	// 更新货位基础信息
	UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, opts ...client.CallOption) (*UpdateResp, error)
	// 根据查询条件获取货位
	GetRacklots(ctx context.Context, in *Query, opts ...client.CallOption) (*Racklots, error)
	// 根据查询条件获取货位
	GetRacklotsByType(ctx context.Context, in *Query, opts ...client.CallOption) (*Racklots, error)
	// 根据ID获取货位信息
	GetRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Racklot, error)
	// 设置重列组
	SetMultipleGroup(ctx context.Context, in *MultipleGroup, opts ...client.CallOption) (*Response, error)
	// 设置货位组
	SetRacklotGroup(ctx context.Context, in *RacklotGroup, opts ...client.CallOption) (*Response, error)
	// 为货位绑定/解绑货架
	BindRack(ctx context.Context, in *RackIDReq, opts ...client.CallOption) (*Response, error)
	UnbindRack(ctx context.Context, in *RackIDReq, opts ...client.CallOption) (*Response, error)
	// 禁用/启用货位
	DisableRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error)
	EnableRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error)
	//站台分组
	AddRacklotGroup(ctx context.Context, in *RacklotGroupsReq, opts ...client.CallOption) (*Response, error)
	RemoveRacklotGroup(ctx context.Context, in *RacklotGroupsReq, opts ...client.CallOption) (*Response, error)
	//占用/释放货位
	OccupyRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error)
	ReleaseRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error)
	// 判断货位是否存在
	IsExist(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error)
	// 判断货位是否有效
	IsAvailable(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error)
	// 获取货位关联的货架
	GetBindedRack(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*IDReq, error)
}

type racklotService struct {
	c    client.Client
	name string
}

func NewRacklotService(name string, c client.Client) RacklotService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.racklot"
	}
	return &racklotService{
		c:    c,
		name: name,
	}
}

func (c *racklotService) AddRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*AddResp, error) {
	req := c.c.NewRequest(c.name, "RacklotService.AddRacklotType", in)
	out := new(AddResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) UpdateRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*UpdateResp, error) {
	req := c.c.NewRequest(c.name, "RacklotService.UpdateRacklotType", in)
	out := new(UpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) DeleteRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*DeleteResp, error) {
	req := c.c.NewRequest(c.name, "RacklotService.DeleteRacklotType", in)
	out := new(DeleteResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetRacklotType(ctx context.Context, in *Query, opts ...client.CallOption) (*RacklotTypes, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklotType", in)
	out := new(RacklotTypes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) AddRacklot(ctx context.Context, in *Racklot, opts ...client.CallOption) (*AddResp, error) {
	req := c.c.NewRequest(c.name, "RacklotService.AddRacklot", in)
	out := new(AddResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) BatchAddRacklots(ctx context.Context, in *Racklots, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.BatchAddRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) DeleteRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*DeleteResp, error) {
	req := c.c.NewRequest(c.name, "RacklotService.DeleteRacklot", in)
	out := new(DeleteResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, opts ...client.CallOption) (*UpdateResp, error) {
	req := c.c.NewRequest(c.name, "RacklotService.UpdateRacklot", in)
	out := new(UpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetRacklots(ctx context.Context, in *Query, opts ...client.CallOption) (*Racklots, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklots", in)
	out := new(Racklots)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetRacklotsByType(ctx context.Context, in *Query, opts ...client.CallOption) (*Racklots, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklotsByType", in)
	out := new(Racklots)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Racklot, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklot", in)
	out := new(Racklot)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) SetMultipleGroup(ctx context.Context, in *MultipleGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.SetMultipleGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) SetRacklotGroup(ctx context.Context, in *RacklotGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.SetRacklotGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) BindRack(ctx context.Context, in *RackIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.BindRack", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) UnbindRack(ctx context.Context, in *RackIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.UnbindRack", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) DisableRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.DisableRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) EnableRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.EnableRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) AddRacklotGroup(ctx context.Context, in *RacklotGroupsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.AddRacklotGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) RemoveRacklotGroup(ctx context.Context, in *RacklotGroupsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.RemoveRacklotGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) OccupyRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.OccupyRacklot", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) ReleaseRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.ReleaseRacklot", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) IsExist(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.IsExist", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) IsAvailable(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.IsAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetBindedRack(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*IDReq, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetBindedRack", in)
	out := new(IDReq)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RacklotService service

type RacklotServiceHandler interface {
	// 添加货位类型
	AddRacklotType(context.Context, *RacklotType, *AddResp) error
	// 修改货位类型
	UpdateRacklotType(context.Context, *RacklotType, *UpdateResp) error
	// 删除货位类型
	DeleteRacklotType(context.Context, *RacklotType, *DeleteResp) error
	// 获取货位类型
	GetRacklotType(context.Context, *Query, *RacklotTypes) error
	// 添加货位
	AddRacklot(context.Context, *Racklot, *AddResp) error
	// 导入货位
	BatchAddRacklots(context.Context, *Racklots, *Response) error
	// 删除货位
	DeleteRacklot(context.Context, *RacklotIDReq, *DeleteResp) error
	// 更新货位基础信息
	UpdateRacklot(context.Context, *UpdateRacklotReq, *UpdateResp) error
	// 根据查询条件获取货位
	GetRacklots(context.Context, *Query, *Racklots) error
	// 根据查询条件获取货位
	GetRacklotsByType(context.Context, *Query, *Racklots) error
	// 根据ID获取货位信息
	GetRacklot(context.Context, *RacklotIDReq, *Racklot) error
	// 设置重列组
	SetMultipleGroup(context.Context, *MultipleGroup, *Response) error
	// 设置货位组
	SetRacklotGroup(context.Context, *RacklotGroup, *Response) error
	// 为货位绑定/解绑货架
	BindRack(context.Context, *RackIDReq, *Response) error
	UnbindRack(context.Context, *RackIDReq, *Response) error
	// 禁用/启用货位
	DisableRacklots(context.Context, *RacklotIDsReq, *Response) error
	EnableRacklots(context.Context, *RacklotIDsReq, *Response) error
	//站台分组
	AddRacklotGroup(context.Context, *RacklotGroupsReq, *Response) error
	RemoveRacklotGroup(context.Context, *RacklotGroupsReq, *Response) error
	//占用/释放货位
	OccupyRacklot(context.Context, *RacklotIDReq, *Response) error
	ReleaseRacklot(context.Context, *RacklotIDReq, *Response) error
	// 判断货位是否存在
	IsExist(context.Context, *RacklotIDReq, *Response) error
	// 判断货位是否有效
	IsAvailable(context.Context, *RacklotIDReq, *Response) error
	// 获取货位关联的货架
	GetBindedRack(context.Context, *RacklotIDReq, *IDReq) error
}

func RegisterRacklotServiceHandler(s server.Server, hdlr RacklotServiceHandler, opts ...server.HandlerOption) error {
	type racklotService interface {
		AddRacklotType(ctx context.Context, in *RacklotType, out *AddResp) error
		UpdateRacklotType(ctx context.Context, in *RacklotType, out *UpdateResp) error
		DeleteRacklotType(ctx context.Context, in *RacklotType, out *DeleteResp) error
		GetRacklotType(ctx context.Context, in *Query, out *RacklotTypes) error
		AddRacklot(ctx context.Context, in *Racklot, out *AddResp) error
		BatchAddRacklots(ctx context.Context, in *Racklots, out *Response) error
		DeleteRacklot(ctx context.Context, in *RacklotIDReq, out *DeleteResp) error
		UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, out *UpdateResp) error
		GetRacklots(ctx context.Context, in *Query, out *Racklots) error
		GetRacklotsByType(ctx context.Context, in *Query, out *Racklots) error
		GetRacklot(ctx context.Context, in *RacklotIDReq, out *Racklot) error
		SetMultipleGroup(ctx context.Context, in *MultipleGroup, out *Response) error
		SetRacklotGroup(ctx context.Context, in *RacklotGroup, out *Response) error
		BindRack(ctx context.Context, in *RackIDReq, out *Response) error
		UnbindRack(ctx context.Context, in *RackIDReq, out *Response) error
		DisableRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error
		EnableRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error
		AddRacklotGroup(ctx context.Context, in *RacklotGroupsReq, out *Response) error
		RemoveRacklotGroup(ctx context.Context, in *RacklotGroupsReq, out *Response) error
		OccupyRacklot(ctx context.Context, in *RacklotIDReq, out *Response) error
		ReleaseRacklot(ctx context.Context, in *RacklotIDReq, out *Response) error
		IsExist(ctx context.Context, in *RacklotIDReq, out *Response) error
		IsAvailable(ctx context.Context, in *RacklotIDReq, out *Response) error
		GetBindedRack(ctx context.Context, in *RacklotIDReq, out *IDReq) error
	}
	type RacklotService struct {
		racklotService
	}
	h := &racklotServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RacklotService{h}, opts...))
}

type racklotServiceHandler struct {
	RacklotServiceHandler
}

func (h *racklotServiceHandler) AddRacklotType(ctx context.Context, in *RacklotType, out *AddResp) error {
	return h.RacklotServiceHandler.AddRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) UpdateRacklotType(ctx context.Context, in *RacklotType, out *UpdateResp) error {
	return h.RacklotServiceHandler.UpdateRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) DeleteRacklotType(ctx context.Context, in *RacklotType, out *DeleteResp) error {
	return h.RacklotServiceHandler.DeleteRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklotType(ctx context.Context, in *Query, out *RacklotTypes) error {
	return h.RacklotServiceHandler.GetRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) AddRacklot(ctx context.Context, in *Racklot, out *AddResp) error {
	return h.RacklotServiceHandler.AddRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) BatchAddRacklots(ctx context.Context, in *Racklots, out *Response) error {
	return h.RacklotServiceHandler.BatchAddRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) DeleteRacklot(ctx context.Context, in *RacklotIDReq, out *DeleteResp) error {
	return h.RacklotServiceHandler.DeleteRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, out *UpdateResp) error {
	return h.RacklotServiceHandler.UpdateRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklots(ctx context.Context, in *Query, out *Racklots) error {
	return h.RacklotServiceHandler.GetRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklotsByType(ctx context.Context, in *Query, out *Racklots) error {
	return h.RacklotServiceHandler.GetRacklotsByType(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklot(ctx context.Context, in *RacklotIDReq, out *Racklot) error {
	return h.RacklotServiceHandler.GetRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) SetMultipleGroup(ctx context.Context, in *MultipleGroup, out *Response) error {
	return h.RacklotServiceHandler.SetMultipleGroup(ctx, in, out)
}

func (h *racklotServiceHandler) SetRacklotGroup(ctx context.Context, in *RacklotGroup, out *Response) error {
	return h.RacklotServiceHandler.SetRacklotGroup(ctx, in, out)
}

func (h *racklotServiceHandler) BindRack(ctx context.Context, in *RackIDReq, out *Response) error {
	return h.RacklotServiceHandler.BindRack(ctx, in, out)
}

func (h *racklotServiceHandler) UnbindRack(ctx context.Context, in *RackIDReq, out *Response) error {
	return h.RacklotServiceHandler.UnbindRack(ctx, in, out)
}

func (h *racklotServiceHandler) DisableRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error {
	return h.RacklotServiceHandler.DisableRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) EnableRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error {
	return h.RacklotServiceHandler.EnableRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) AddRacklotGroup(ctx context.Context, in *RacklotGroupsReq, out *Response) error {
	return h.RacklotServiceHandler.AddRacklotGroup(ctx, in, out)
}

func (h *racklotServiceHandler) RemoveRacklotGroup(ctx context.Context, in *RacklotGroupsReq, out *Response) error {
	return h.RacklotServiceHandler.RemoveRacklotGroup(ctx, in, out)
}

func (h *racklotServiceHandler) OccupyRacklot(ctx context.Context, in *RacklotIDReq, out *Response) error {
	return h.RacklotServiceHandler.OccupyRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) ReleaseRacklot(ctx context.Context, in *RacklotIDReq, out *Response) error {
	return h.RacklotServiceHandler.ReleaseRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) IsExist(ctx context.Context, in *RacklotIDReq, out *Response) error {
	return h.RacklotServiceHandler.IsExist(ctx, in, out)
}

func (h *racklotServiceHandler) IsAvailable(ctx context.Context, in *RacklotIDReq, out *Response) error {
	return h.RacklotServiceHandler.IsAvailable(ctx, in, out)
}

func (h *racklotServiceHandler) GetBindedRack(ctx context.Context, in *RacklotIDReq, out *IDReq) error {
	return h.RacklotServiceHandler.GetBindedRack(ctx, in, out)
}