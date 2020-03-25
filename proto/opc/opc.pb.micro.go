// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/opc/opc.proto

package go_micro_srv_opc

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

// Client API for OpcService service

type OpcService interface {
	GetGroup(ctx context.Context, in *GrpReq, opts ...client.CallOption) (*Group, error)
	GetGroups(ctx context.Context, in *GrpQuery, opts ...client.CallOption) (*Groups, error)
	AddGroup(ctx context.Context, in *Group, opts ...client.CallOption) (*Response, error)
	UpdateGroup(ctx context.Context, in *Group, opts ...client.CallOption) (*Response, error)
	DeleteGroup(ctx context.Context, in *Group, opts ...client.CallOption) (*Response, error)
	GetItem(ctx context.Context, in *ItemReq, opts ...client.CallOption) (*Item, error)
	GetItems(ctx context.Context, in *ItemQuery, opts ...client.CallOption) (*Items, error)
	AddItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	UpdateItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	DeleteItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error)
	WriteItem(ctx context.Context, in *ItemValueReq, opts ...client.CallOption) (*Response, error)
	UpdateFile(ctx context.Context, in *FileReq, opts ...client.CallOption) (*Response, error)
}

type opcService struct {
	c    client.Client
	name string
}

func NewOpcService(name string, c client.Client) OpcService {
	return &opcService{
		c:    c,
		name: name,
	}
}

func (c *opcService) GetGroup(ctx context.Context, in *GrpReq, opts ...client.CallOption) (*Group, error) {
	req := c.c.NewRequest(c.name, "OpcService.GetGroup", in)
	out := new(Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) GetGroups(ctx context.Context, in *GrpQuery, opts ...client.CallOption) (*Groups, error) {
	req := c.c.NewRequest(c.name, "OpcService.GetGroups", in)
	out := new(Groups)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) AddGroup(ctx context.Context, in *Group, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.AddGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) UpdateGroup(ctx context.Context, in *Group, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.UpdateGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) DeleteGroup(ctx context.Context, in *Group, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.DeleteGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) GetItem(ctx context.Context, in *ItemReq, opts ...client.CallOption) (*Item, error) {
	req := c.c.NewRequest(c.name, "OpcService.GetItem", in)
	out := new(Item)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) GetItems(ctx context.Context, in *ItemQuery, opts ...client.CallOption) (*Items, error) {
	req := c.c.NewRequest(c.name, "OpcService.GetItems", in)
	out := new(Items)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) AddItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.AddItem", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) UpdateItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.UpdateItem", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) DeleteItem(ctx context.Context, in *Item, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.DeleteItem", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) WriteItem(ctx context.Context, in *ItemValueReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.WriteItem", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opcService) UpdateFile(ctx context.Context, in *FileReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "OpcService.UpdateFile", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OpcService service

type OpcServiceHandler interface {
	GetGroup(context.Context, *GrpReq, *Group) error
	GetGroups(context.Context, *GrpQuery, *Groups) error
	AddGroup(context.Context, *Group, *Response) error
	UpdateGroup(context.Context, *Group, *Response) error
	DeleteGroup(context.Context, *Group, *Response) error
	GetItem(context.Context, *ItemReq, *Item) error
	GetItems(context.Context, *ItemQuery, *Items) error
	AddItem(context.Context, *Item, *Response) error
	UpdateItem(context.Context, *Item, *Response) error
	DeleteItem(context.Context, *Item, *Response) error
	WriteItem(context.Context, *ItemValueReq, *Response) error
	UpdateFile(context.Context, *FileReq, *Response) error
}

func RegisterOpcServiceHandler(s server.Server, hdlr OpcServiceHandler, opts ...server.HandlerOption) error {
	type opcService interface {
		GetGroup(ctx context.Context, in *GrpReq, out *Group) error
		GetGroups(ctx context.Context, in *GrpQuery, out *Groups) error
		AddGroup(ctx context.Context, in *Group, out *Response) error
		UpdateGroup(ctx context.Context, in *Group, out *Response) error
		DeleteGroup(ctx context.Context, in *Group, out *Response) error
		GetItem(ctx context.Context, in *ItemReq, out *Item) error
		GetItems(ctx context.Context, in *ItemQuery, out *Items) error
		AddItem(ctx context.Context, in *Item, out *Response) error
		UpdateItem(ctx context.Context, in *Item, out *Response) error
		DeleteItem(ctx context.Context, in *Item, out *Response) error
		WriteItem(ctx context.Context, in *ItemValueReq, out *Response) error
		UpdateFile(ctx context.Context, in *FileReq, out *Response) error
	}
	type OpcService struct {
		opcService
	}
	h := &opcServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OpcService{h}, opts...))
}

type opcServiceHandler struct {
	OpcServiceHandler
}

func (h *opcServiceHandler) GetGroup(ctx context.Context, in *GrpReq, out *Group) error {
	return h.OpcServiceHandler.GetGroup(ctx, in, out)
}

func (h *opcServiceHandler) GetGroups(ctx context.Context, in *GrpQuery, out *Groups) error {
	return h.OpcServiceHandler.GetGroups(ctx, in, out)
}

func (h *opcServiceHandler) AddGroup(ctx context.Context, in *Group, out *Response) error {
	return h.OpcServiceHandler.AddGroup(ctx, in, out)
}

func (h *opcServiceHandler) UpdateGroup(ctx context.Context, in *Group, out *Response) error {
	return h.OpcServiceHandler.UpdateGroup(ctx, in, out)
}

func (h *opcServiceHandler) DeleteGroup(ctx context.Context, in *Group, out *Response) error {
	return h.OpcServiceHandler.DeleteGroup(ctx, in, out)
}

func (h *opcServiceHandler) GetItem(ctx context.Context, in *ItemReq, out *Item) error {
	return h.OpcServiceHandler.GetItem(ctx, in, out)
}

func (h *opcServiceHandler) GetItems(ctx context.Context, in *ItemQuery, out *Items) error {
	return h.OpcServiceHandler.GetItems(ctx, in, out)
}

func (h *opcServiceHandler) AddItem(ctx context.Context, in *Item, out *Response) error {
	return h.OpcServiceHandler.AddItem(ctx, in, out)
}

func (h *opcServiceHandler) UpdateItem(ctx context.Context, in *Item, out *Response) error {
	return h.OpcServiceHandler.UpdateItem(ctx, in, out)
}

func (h *opcServiceHandler) DeleteItem(ctx context.Context, in *Item, out *Response) error {
	return h.OpcServiceHandler.DeleteItem(ctx, in, out)
}

func (h *opcServiceHandler) WriteItem(ctx context.Context, in *ItemValueReq, out *Response) error {
	return h.OpcServiceHandler.WriteItem(ctx, in, out)
}

func (h *opcServiceHandler) UpdateFile(ctx context.Context, in *FileReq, out *Response) error {
	return h.OpcServiceHandler.UpdateFile(ctx, in, out)
}
