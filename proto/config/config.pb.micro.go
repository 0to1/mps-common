// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/config/config.proto

package go_micro_srv_config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Client API for ConfigService service

type ConfigService interface {
	// 修改仓库信息
	UpdateWarehouse(ctx context.Context, in *Warehouse, opts ...client.CallOption) (*Response, error)
	// 获取仓库信息
	GetWarehouse(ctx context.Context, in *Request, opts ...client.CallOption) (*Warehouse, error)
	// 增加脚本参数
	AddScriptParameter(ctx context.Context, in *ScriptParameter, opts ...client.CallOption) (*Response, error)
	// 修改脚本参数
	UpdateScriptParameter(ctx context.Context, in *ScriptParameter, opts ...client.CallOption) (*Response, error)
	// 删除脚本参数
	DeleteScriptParameter(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*Response, error)
	// 查询单个脚本参数
	GetScriptParameter(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*ScriptParameter, error)
	// 查询所有的脚本参数
	GetScriptParameters(ctx context.Context, in *Query, opts ...client.CallOption) (*ScriptParameters, error)
	// 增加脚本按钮
	AddScriptButton(ctx context.Context, in *InputScriptButton, opts ...client.CallOption) (*Response, error)
	// 修改脚本按钮
	UpdateScriptButton(ctx context.Context, in *InputScriptButton, opts ...client.CallOption) (*Response, error)
	// 删除脚本按钮
	DeleteScriptButton(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*Response, error)
	// 查询单个脚本按钮
	GetScriptButton(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*ScriptButton, error)
	// 查询所有的脚本按钮
	GetScriptButtons(ctx context.Context, in *Query, opts ...client.CallOption) (*ScriptButtons, error)
}

type configService struct {
	c    client.Client
	name string
}

func NewConfigService(name string, c client.Client) ConfigService {
	return &configService{
		c:    c,
		name: name,
	}
}

func (c *configService) UpdateWarehouse(ctx context.Context, in *Warehouse, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.UpdateWarehouse", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetWarehouse(ctx context.Context, in *Request, opts ...client.CallOption) (*Warehouse, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetWarehouse", in)
	out := new(Warehouse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) AddScriptParameter(ctx context.Context, in *ScriptParameter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.AddScriptParameter", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) UpdateScriptParameter(ctx context.Context, in *ScriptParameter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.UpdateScriptParameter", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteScriptParameter(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.DeleteScriptParameter", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetScriptParameter(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*ScriptParameter, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetScriptParameter", in)
	out := new(ScriptParameter)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetScriptParameters(ctx context.Context, in *Query, opts ...client.CallOption) (*ScriptParameters, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetScriptParameters", in)
	out := new(ScriptParameters)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) AddScriptButton(ctx context.Context, in *InputScriptButton, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.AddScriptButton", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) UpdateScriptButton(ctx context.Context, in *InputScriptButton, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.UpdateScriptButton", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteScriptButton(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConfigService.DeleteScriptButton", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetScriptButton(ctx context.Context, in *IkeyRequest, opts ...client.CallOption) (*ScriptButton, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetScriptButton", in)
	out := new(ScriptButton)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) GetScriptButtons(ctx context.Context, in *Query, opts ...client.CallOption) (*ScriptButtons, error) {
	req := c.c.NewRequest(c.name, "ConfigService.GetScriptButtons", in)
	out := new(ScriptButtons)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigService service

type ConfigServiceHandler interface {
	// 修改仓库信息
	UpdateWarehouse(context.Context, *Warehouse, *Response) error
	// 获取仓库信息
	GetWarehouse(context.Context, *Request, *Warehouse) error
	// 增加脚本参数
	AddScriptParameter(context.Context, *ScriptParameter, *Response) error
	// 修改脚本参数
	UpdateScriptParameter(context.Context, *ScriptParameter, *Response) error
	// 删除脚本参数
	DeleteScriptParameter(context.Context, *IkeyRequest, *Response) error
	// 查询单个脚本参数
	GetScriptParameter(context.Context, *IkeyRequest, *ScriptParameter) error
	// 查询所有的脚本参数
	GetScriptParameters(context.Context, *Query, *ScriptParameters) error
	// 增加脚本按钮
	AddScriptButton(context.Context, *InputScriptButton, *Response) error
	// 修改脚本按钮
	UpdateScriptButton(context.Context, *InputScriptButton, *Response) error
	// 删除脚本按钮
	DeleteScriptButton(context.Context, *IkeyRequest, *Response) error
	// 查询单个脚本按钮
	GetScriptButton(context.Context, *IkeyRequest, *ScriptButton) error
	// 查询所有的脚本按钮
	GetScriptButtons(context.Context, *Query, *ScriptButtons) error
}

func RegisterConfigServiceHandler(s server.Server, hdlr ConfigServiceHandler, opts ...server.HandlerOption) error {
	type configService interface {
		UpdateWarehouse(ctx context.Context, in *Warehouse, out *Response) error
		GetWarehouse(ctx context.Context, in *Request, out *Warehouse) error
		AddScriptParameter(ctx context.Context, in *ScriptParameter, out *Response) error
		UpdateScriptParameter(ctx context.Context, in *ScriptParameter, out *Response) error
		DeleteScriptParameter(ctx context.Context, in *IkeyRequest, out *Response) error
		GetScriptParameter(ctx context.Context, in *IkeyRequest, out *ScriptParameter) error
		GetScriptParameters(ctx context.Context, in *Query, out *ScriptParameters) error
		AddScriptButton(ctx context.Context, in *InputScriptButton, out *Response) error
		UpdateScriptButton(ctx context.Context, in *InputScriptButton, out *Response) error
		DeleteScriptButton(ctx context.Context, in *IkeyRequest, out *Response) error
		GetScriptButton(ctx context.Context, in *IkeyRequest, out *ScriptButton) error
		GetScriptButtons(ctx context.Context, in *Query, out *ScriptButtons) error
	}
	type ConfigService struct {
		configService
	}
	h := &configServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConfigService{h}, opts...))
}

type configServiceHandler struct {
	ConfigServiceHandler
}

func (h *configServiceHandler) UpdateWarehouse(ctx context.Context, in *Warehouse, out *Response) error {
	return h.ConfigServiceHandler.UpdateWarehouse(ctx, in, out)
}

func (h *configServiceHandler) GetWarehouse(ctx context.Context, in *Request, out *Warehouse) error {
	return h.ConfigServiceHandler.GetWarehouse(ctx, in, out)
}

func (h *configServiceHandler) AddScriptParameter(ctx context.Context, in *ScriptParameter, out *Response) error {
	return h.ConfigServiceHandler.AddScriptParameter(ctx, in, out)
}

func (h *configServiceHandler) UpdateScriptParameter(ctx context.Context, in *ScriptParameter, out *Response) error {
	return h.ConfigServiceHandler.UpdateScriptParameter(ctx, in, out)
}

func (h *configServiceHandler) DeleteScriptParameter(ctx context.Context, in *IkeyRequest, out *Response) error {
	return h.ConfigServiceHandler.DeleteScriptParameter(ctx, in, out)
}

func (h *configServiceHandler) GetScriptParameter(ctx context.Context, in *IkeyRequest, out *ScriptParameter) error {
	return h.ConfigServiceHandler.GetScriptParameter(ctx, in, out)
}

func (h *configServiceHandler) GetScriptParameters(ctx context.Context, in *Query, out *ScriptParameters) error {
	return h.ConfigServiceHandler.GetScriptParameters(ctx, in, out)
}

func (h *configServiceHandler) AddScriptButton(ctx context.Context, in *InputScriptButton, out *Response) error {
	return h.ConfigServiceHandler.AddScriptButton(ctx, in, out)
}

func (h *configServiceHandler) UpdateScriptButton(ctx context.Context, in *InputScriptButton, out *Response) error {
	return h.ConfigServiceHandler.UpdateScriptButton(ctx, in, out)
}

func (h *configServiceHandler) DeleteScriptButton(ctx context.Context, in *IkeyRequest, out *Response) error {
	return h.ConfigServiceHandler.DeleteScriptButton(ctx, in, out)
}

func (h *configServiceHandler) GetScriptButton(ctx context.Context, in *IkeyRequest, out *ScriptButton) error {
	return h.ConfigServiceHandler.GetScriptButton(ctx, in, out)
}

func (h *configServiceHandler) GetScriptButtons(ctx context.Context, in *Query, out *ScriptButtons) error {
	return h.ConfigServiceHandler.GetScriptButtons(ctx, in, out)
}
