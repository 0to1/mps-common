// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order/order.proto

package go_micro_srv_agv_order

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

// Client API for Order service

type OrderService interface {
	//取消任务
	CancelTask(ctx context.Context, in *CancelReq, opts ...client.CallOption) (*Response, error)
	//修改参数
	UpdateTask(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*Response, error)
	SendQAMessage(ctx context.Context, in *QAMessage, opts ...client.CallOption) (*Response, error)
	SendQBMessage(ctx context.Context, in *QBMessage, opts ...client.CallOption) (*Response, error)
	SendMMessage(ctx context.Context, in *MMessage, opts ...client.CallOption) (*Response, error)
	HostIsConnected(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CancelTask(ctx context.Context, in *CancelReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.CancelTask", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateTask(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateTask", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SendQAMessage(ctx context.Context, in *QAMessage, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.SendQAMessage", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SendQBMessage(ctx context.Context, in *QBMessage, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.SendQBMessage", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SendMMessage(ctx context.Context, in *MMessage, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.SendMMessage", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) HostIsConnected(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.HostIsConnected", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	//取消任务
	CancelTask(context.Context, *CancelReq, *Response) error
	//修改参数
	UpdateTask(context.Context, *UpdateReq, *Response) error
	SendQAMessage(context.Context, *QAMessage, *Response) error
	SendQBMessage(context.Context, *QBMessage, *Response) error
	SendMMessage(context.Context, *MMessage, *Response) error
	HostIsConnected(context.Context, *Request, *Response) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		CancelTask(ctx context.Context, in *CancelReq, out *Response) error
		UpdateTask(ctx context.Context, in *UpdateReq, out *Response) error
		SendQAMessage(ctx context.Context, in *QAMessage, out *Response) error
		SendQBMessage(ctx context.Context, in *QBMessage, out *Response) error
		SendMMessage(ctx context.Context, in *MMessage, out *Response) error
		HostIsConnected(ctx context.Context, in *Request, out *Response) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) CancelTask(ctx context.Context, in *CancelReq, out *Response) error {
	return h.OrderHandler.CancelTask(ctx, in, out)
}

func (h *orderHandler) UpdateTask(ctx context.Context, in *UpdateReq, out *Response) error {
	return h.OrderHandler.UpdateTask(ctx, in, out)
}

func (h *orderHandler) SendQAMessage(ctx context.Context, in *QAMessage, out *Response) error {
	return h.OrderHandler.SendQAMessage(ctx, in, out)
}

func (h *orderHandler) SendQBMessage(ctx context.Context, in *QBMessage, out *Response) error {
	return h.OrderHandler.SendQBMessage(ctx, in, out)
}

func (h *orderHandler) SendMMessage(ctx context.Context, in *MMessage, out *Response) error {
	return h.OrderHandler.SendMMessage(ctx, in, out)
}

func (h *orderHandler) HostIsConnected(ctx context.Context, in *Request, out *Response) error {
	return h.OrderHandler.HostIsConnected(ctx, in, out)
}
