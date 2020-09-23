// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/system/system.proto

package go_micro_srv_system

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for System service

func NewSystemEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for System service

type SystemService interface {
	//更新参数，调度系统的IP地址和端口//
	UpdateParam(ctx context.Context, in *Params, opts ...client.CallOption) (*Response, error)
	//获取系统参数
	GetParam(ctx context.Context, in *ParamReq, opts ...client.CallOption) (*Params, error)
	//系统操作//
	SystemOperation(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type systemService struct {
	c    client.Client
	name string
}

func NewSystemService(name string, c client.Client) SystemService {
	return &systemService{
		c:    c,
		name: name,
	}
}

func (c *systemService) UpdateParam(ctx context.Context, in *Params, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "System.UpdateParam", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemService) GetParam(ctx context.Context, in *ParamReq, opts ...client.CallOption) (*Params, error) {
	req := c.c.NewRequest(c.name, "System.GetParam", in)
	out := new(Params)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemService) SystemOperation(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "System.SystemOperation", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for System service

type SystemHandler interface {
	//更新参数，调度系统的IP地址和端口//
	UpdateParam(context.Context, *Params, *Response) error
	//获取系统参数
	GetParam(context.Context, *ParamReq, *Params) error
	//系统操作//
	SystemOperation(context.Context, *Request, *Response) error
}

func RegisterSystemHandler(s server.Server, hdlr SystemHandler, opts ...server.HandlerOption) error {
	type system interface {
		UpdateParam(ctx context.Context, in *Params, out *Response) error
		GetParam(ctx context.Context, in *ParamReq, out *Params) error
		SystemOperation(ctx context.Context, in *Request, out *Response) error
	}
	type System struct {
		system
	}
	h := &systemHandler{hdlr}
	return s.Handle(s.NewHandler(&System{h}, opts...))
}

type systemHandler struct {
	SystemHandler
}

func (h *systemHandler) UpdateParam(ctx context.Context, in *Params, out *Response) error {
	return h.SystemHandler.UpdateParam(ctx, in, out)
}

func (h *systemHandler) GetParam(ctx context.Context, in *ParamReq, out *Params) error {
	return h.SystemHandler.GetParam(ctx, in, out)
}

func (h *systemHandler) SystemOperation(ctx context.Context, in *Request, out *Response) error {
	return h.SystemHandler.SystemOperation(ctx, in, out)
}
