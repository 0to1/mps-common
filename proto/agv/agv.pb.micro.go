// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/agv/agv.proto

package go_micro_srv_agv

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

// Client API for Agv service

type AgvService interface {
	//获取一个AGV数据//
	GetAgvByID(ctx context.Context, in *AgvReq, opts ...client.CallOption) (*AgvMessage, error)
	//获取所有的AGV数据//
	GetAgvs(ctx context.Context, in *Query, opts ...client.CallOption) (*AgvsResponse, error)
	//急停单台AGV//
	StopAgvByID(ctx context.Context, in *AgvReq, opts ...client.CallOption) (*Response, error)
	//急停所有AGV//
	StopAgvs(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type agvService struct {
	c    client.Client
	name string
}

func NewAgvService(name string, c client.Client) AgvService {
	return &agvService{
		c:    c,
		name: name,
	}
}

func (c *agvService) GetAgvByID(ctx context.Context, in *AgvReq, opts ...client.CallOption) (*AgvMessage, error) {
	req := c.c.NewRequest(c.name, "Agv.GetAgvByID", in)
	out := new(AgvMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agvService) GetAgvs(ctx context.Context, in *Query, opts ...client.CallOption) (*AgvsResponse, error) {
	req := c.c.NewRequest(c.name, "Agv.GetAgvs", in)
	out := new(AgvsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agvService) StopAgvByID(ctx context.Context, in *AgvReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Agv.StopAgvByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agvService) StopAgvs(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Agv.StopAgvs", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agv service

type AgvHandler interface {
	//获取一个AGV数据//
	GetAgvByID(context.Context, *AgvReq, *AgvMessage) error
	//获取所有的AGV数据//
	GetAgvs(context.Context, *Query, *AgvsResponse) error
	//急停单台AGV//
	StopAgvByID(context.Context, *AgvReq, *Response) error
	//急停所有AGV//
	StopAgvs(context.Context, *Request, *Response) error
}

func RegisterAgvHandler(s server.Server, hdlr AgvHandler, opts ...server.HandlerOption) error {
	type agv interface {
		GetAgvByID(ctx context.Context, in *AgvReq, out *AgvMessage) error
		GetAgvs(ctx context.Context, in *Query, out *AgvsResponse) error
		StopAgvByID(ctx context.Context, in *AgvReq, out *Response) error
		StopAgvs(ctx context.Context, in *Request, out *Response) error
	}
	type Agv struct {
		agv
	}
	h := &agvHandler{hdlr}
	return s.Handle(s.NewHandler(&Agv{h}, opts...))
}

type agvHandler struct {
	AgvHandler
}

func (h *agvHandler) GetAgvByID(ctx context.Context, in *AgvReq, out *AgvMessage) error {
	return h.AgvHandler.GetAgvByID(ctx, in, out)
}

func (h *agvHandler) GetAgvs(ctx context.Context, in *Query, out *AgvsResponse) error {
	return h.AgvHandler.GetAgvs(ctx, in, out)
}

func (h *agvHandler) StopAgvByID(ctx context.Context, in *AgvReq, out *Response) error {
	return h.AgvHandler.StopAgvByID(ctx, in, out)
}

func (h *agvHandler) StopAgvs(ctx context.Context, in *Request, out *Response) error {
	return h.AgvHandler.StopAgvs(ctx, in, out)
}
