// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/racklot/racklot.proto

package go_micro_srv_racklot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RacklotService service

func NewRacklotServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RacklotService service

type RacklotService interface {
	// 根据查询条件获取货位
	GetRacklots(ctx context.Context, in *RacklotQuery, opts ...client.CallOption) (*Racklots, error)
	// 根据ID获取货位信息
	GetRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*Racklot, error)
	// 获取货位类型
	GetRacklotTypes(ctx context.Context, in *RacklotTypeQuery, opts ...client.CallOption) (*RacklotTypes, error)
	// 获取货位类型
	GetRacklotType(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*RacklotType, error)
	// 添加货位
	AddRacklot(ctx context.Context, in *Racklot, opts ...client.CallOption) (*AddResp, error)
	// 批量导入货位
	BatchAddRacklots(ctx context.Context, in *Racklots, opts ...client.CallOption) (*Response, error)
	// 删除货位
	DeleteRacklot(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*DeleteResp, error)
	// 批量删除货位
	BatchDeleteRacklots(ctx context.Context, in *DeleteRacklotsReq, opts ...client.CallOption) (*Response, error)
	// 更新货位基础信息
	UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, opts ...client.CallOption) (*UpdateResp, error)
	// 批量修改货位
	BatchUpdateRacklots(ctx context.Context, in *UpdateRacklotsReq, opts ...client.CallOption) (*Response, error)
	// 添加货位类型
	AddRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*AddResp, error)
	// 修改货位类型
	UpdateRacklotType(ctx context.Context, in *UpdateRacklotTypeReq, opts ...client.CallOption) (*UpdateResp, error)
	// 删除货GetBindedRack位类型
	DeleteRacklotType(ctx context.Context, in *RacklotType, opts ...client.CallOption) (*DeleteResp, error)
	// 设置重列组
	SetMultipleGroup(ctx context.Context, in *MultipleGroup, opts ...client.CallOption) (*Response, error)
	// 为货位绑定/解绑货架
	BindRack(ctx context.Context, in *RackIDReq, opts ...client.CallOption) (*Response, error)
	UnbindRack(ctx context.Context, in *RackIDReq, opts ...client.CallOption) (*Response, error)
	// 禁用/启用货位
	DisableRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error)
	EnableRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error)
	//占用/释放货位
	OccupyRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error)
	ReleaseRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error)
	// 设置是否允许存车
	SetInbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error)
	// 设置是否允许取车
	SetOutbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error)
	AddProperties(ctx context.Context, in *PropertiesReq, opts ...client.CallOption) (*Response, error)
	GetRacklotReport(ctx context.Context, in *RacklotReportReq, opts ...client.CallOption) (*RacklotReport, error)
}

type racklotService struct {
	c    client.Client
	name string
}

func NewRacklotService(name string, c client.Client) RacklotService {
	return &racklotService{
		c:    c,
		name: name,
	}
}

func (c *racklotService) GetRacklots(ctx context.Context, in *RacklotQuery, opts ...client.CallOption) (*Racklots, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklots", in)
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

func (c *racklotService) GetRacklotTypes(ctx context.Context, in *RacklotTypeQuery, opts ...client.CallOption) (*RacklotTypes, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklotTypes", in)
	out := new(RacklotTypes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetRacklotType(ctx context.Context, in *RacklotIDReq, opts ...client.CallOption) (*RacklotType, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklotType", in)
	out := new(RacklotType)
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

func (c *racklotService) BatchDeleteRacklots(ctx context.Context, in *DeleteRacklotsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.BatchDeleteRacklots", in)
	out := new(Response)
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

func (c *racklotService) BatchUpdateRacklots(ctx context.Context, in *UpdateRacklotsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.BatchUpdateRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *racklotService) UpdateRacklotType(ctx context.Context, in *UpdateRacklotTypeReq, opts ...client.CallOption) (*UpdateResp, error) {
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

func (c *racklotService) SetMultipleGroup(ctx context.Context, in *MultipleGroup, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.SetMultipleGroup", in)
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

func (c *racklotService) OccupyRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.OccupyRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) ReleaseRacklots(ctx context.Context, in *RacklotIDsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.ReleaseRacklots", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) SetInbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.SetInbound", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) SetOutbound(ctx context.Context, in *FlagReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.SetOutbound", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) AddProperties(ctx context.Context, in *PropertiesReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RacklotService.AddProperties", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racklotService) GetRacklotReport(ctx context.Context, in *RacklotReportReq, opts ...client.CallOption) (*RacklotReport, error) {
	req := c.c.NewRequest(c.name, "RacklotService.GetRacklotReport", in)
	out := new(RacklotReport)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RacklotService service

type RacklotServiceHandler interface {
	// 根据查询条件获取货位
	GetRacklots(context.Context, *RacklotQuery, *Racklots) error
	// 根据ID获取货位信息
	GetRacklot(context.Context, *RacklotIDReq, *Racklot) error
	// 获取货位类型
	GetRacklotTypes(context.Context, *RacklotTypeQuery, *RacklotTypes) error
	// 获取货位类型
	GetRacklotType(context.Context, *RacklotIDReq, *RacklotType) error
	// 添加货位
	AddRacklot(context.Context, *Racklot, *AddResp) error
	// 批量导入货位
	BatchAddRacklots(context.Context, *Racklots, *Response) error
	// 删除货位
	DeleteRacklot(context.Context, *RacklotIDReq, *DeleteResp) error
	// 批量删除货位
	BatchDeleteRacklots(context.Context, *DeleteRacklotsReq, *Response) error
	// 更新货位基础信息
	UpdateRacklot(context.Context, *UpdateRacklotReq, *UpdateResp) error
	// 批量修改货位
	BatchUpdateRacklots(context.Context, *UpdateRacklotsReq, *Response) error
	// 添加货位类型
	AddRacklotType(context.Context, *RacklotType, *AddResp) error
	// 修改货位类型
	UpdateRacklotType(context.Context, *UpdateRacklotTypeReq, *UpdateResp) error
	// 删除货GetBindedRack位类型
	DeleteRacklotType(context.Context, *RacklotType, *DeleteResp) error
	// 设置重列组
	SetMultipleGroup(context.Context, *MultipleGroup, *Response) error
	// 为货位绑定/解绑货架
	BindRack(context.Context, *RackIDReq, *Response) error
	UnbindRack(context.Context, *RackIDReq, *Response) error
	// 禁用/启用货位
	DisableRacklots(context.Context, *RacklotIDsReq, *Response) error
	EnableRacklots(context.Context, *RacklotIDsReq, *Response) error
	//占用/释放货位
	OccupyRacklots(context.Context, *RacklotIDsReq, *Response) error
	ReleaseRacklots(context.Context, *RacklotIDsReq, *Response) error
	// 设置是否允许存车
	SetInbound(context.Context, *FlagReq, *Response) error
	// 设置是否允许取车
	SetOutbound(context.Context, *FlagReq, *Response) error
	AddProperties(context.Context, *PropertiesReq, *Response) error
	GetRacklotReport(context.Context, *RacklotReportReq, *RacklotReport) error
}

func RegisterRacklotServiceHandler(s server.Server, hdlr RacklotServiceHandler, opts ...server.HandlerOption) error {
	type racklotService interface {
		GetRacklots(ctx context.Context, in *RacklotQuery, out *Racklots) error
		GetRacklot(ctx context.Context, in *RacklotIDReq, out *Racklot) error
		GetRacklotTypes(ctx context.Context, in *RacklotTypeQuery, out *RacklotTypes) error
		GetRacklotType(ctx context.Context, in *RacklotIDReq, out *RacklotType) error
		AddRacklot(ctx context.Context, in *Racklot, out *AddResp) error
		BatchAddRacklots(ctx context.Context, in *Racklots, out *Response) error
		DeleteRacklot(ctx context.Context, in *RacklotIDReq, out *DeleteResp) error
		BatchDeleteRacklots(ctx context.Context, in *DeleteRacklotsReq, out *Response) error
		UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, out *UpdateResp) error
		BatchUpdateRacklots(ctx context.Context, in *UpdateRacklotsReq, out *Response) error
		AddRacklotType(ctx context.Context, in *RacklotType, out *AddResp) error
		UpdateRacklotType(ctx context.Context, in *UpdateRacklotTypeReq, out *UpdateResp) error
		DeleteRacklotType(ctx context.Context, in *RacklotType, out *DeleteResp) error
		SetMultipleGroup(ctx context.Context, in *MultipleGroup, out *Response) error
		BindRack(ctx context.Context, in *RackIDReq, out *Response) error
		UnbindRack(ctx context.Context, in *RackIDReq, out *Response) error
		DisableRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error
		EnableRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error
		OccupyRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error
		ReleaseRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error
		SetInbound(ctx context.Context, in *FlagReq, out *Response) error
		SetOutbound(ctx context.Context, in *FlagReq, out *Response) error
		AddProperties(ctx context.Context, in *PropertiesReq, out *Response) error
		GetRacklotReport(ctx context.Context, in *RacklotReportReq, out *RacklotReport) error
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

func (h *racklotServiceHandler) GetRacklots(ctx context.Context, in *RacklotQuery, out *Racklots) error {
	return h.RacklotServiceHandler.GetRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklot(ctx context.Context, in *RacklotIDReq, out *Racklot) error {
	return h.RacklotServiceHandler.GetRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklotTypes(ctx context.Context, in *RacklotTypeQuery, out *RacklotTypes) error {
	return h.RacklotServiceHandler.GetRacklotTypes(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklotType(ctx context.Context, in *RacklotIDReq, out *RacklotType) error {
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

func (h *racklotServiceHandler) BatchDeleteRacklots(ctx context.Context, in *DeleteRacklotsReq, out *Response) error {
	return h.RacklotServiceHandler.BatchDeleteRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) UpdateRacklot(ctx context.Context, in *UpdateRacklotReq, out *UpdateResp) error {
	return h.RacklotServiceHandler.UpdateRacklot(ctx, in, out)
}

func (h *racklotServiceHandler) BatchUpdateRacklots(ctx context.Context, in *UpdateRacklotsReq, out *Response) error {
	return h.RacklotServiceHandler.BatchUpdateRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) AddRacklotType(ctx context.Context, in *RacklotType, out *AddResp) error {
	return h.RacklotServiceHandler.AddRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) UpdateRacklotType(ctx context.Context, in *UpdateRacklotTypeReq, out *UpdateResp) error {
	return h.RacklotServiceHandler.UpdateRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) DeleteRacklotType(ctx context.Context, in *RacklotType, out *DeleteResp) error {
	return h.RacklotServiceHandler.DeleteRacklotType(ctx, in, out)
}

func (h *racklotServiceHandler) SetMultipleGroup(ctx context.Context, in *MultipleGroup, out *Response) error {
	return h.RacklotServiceHandler.SetMultipleGroup(ctx, in, out)
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

func (h *racklotServiceHandler) OccupyRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error {
	return h.RacklotServiceHandler.OccupyRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) ReleaseRacklots(ctx context.Context, in *RacklotIDsReq, out *Response) error {
	return h.RacklotServiceHandler.ReleaseRacklots(ctx, in, out)
}

func (h *racklotServiceHandler) SetInbound(ctx context.Context, in *FlagReq, out *Response) error {
	return h.RacklotServiceHandler.SetInbound(ctx, in, out)
}

func (h *racklotServiceHandler) SetOutbound(ctx context.Context, in *FlagReq, out *Response) error {
	return h.RacklotServiceHandler.SetOutbound(ctx, in, out)
}

func (h *racklotServiceHandler) AddProperties(ctx context.Context, in *PropertiesReq, out *Response) error {
	return h.RacklotServiceHandler.AddProperties(ctx, in, out)
}

func (h *racklotServiceHandler) GetRacklotReport(ctx context.Context, in *RacklotReportReq, out *RacklotReport) error {
	return h.RacklotServiceHandler.GetRacklotReport(ctx, in, out)
}
