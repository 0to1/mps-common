// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/racktype/racktype.proto

package go_micro_srv_racktype

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

// Client API for RackTypeService service

type RackTypeService interface {
	AddRackType(ctx context.Context, in *RackType, opts ...client.CallOption) (*Response, error)
	DeleteRackType(ctx context.Context, in *IDReq, opts ...client.CallOption) (*Response, error)
	UpdateRackType(ctx context.Context, in *RackType, opts ...client.CallOption) (*Response, error)
	GetOneRackTypes(ctx context.Context, in *RackTypeQuery, opts ...client.CallOption) (*RackType, error)
	GetRackTypes(ctx context.Context, in *IDReq, opts ...client.CallOption) (*RackTypes, error)
	AddCellType(ctx context.Context, in *CellType, opts ...client.CallOption) (*Response, error)
	DeleteCellType(ctx context.Context, in *IDReq, opts ...client.CallOption) (*Response, error)
	UpdateCellType(ctx context.Context, in *CellType, opts ...client.CallOption) (*Response, error)
	GetCellTypes(ctx context.Context, in *IDReq, opts ...client.CallOption) (*CellType, error)
	GetOneCellTypes(ctx context.Context, in *CellTypeQuery, opts ...client.CallOption) (*CellTypes, error)
}

type rackTypeService struct {
	c    client.Client
	name string
}

func NewRackTypeService(name string, c client.Client) RackTypeService {
	return &rackTypeService{
		c:    c,
		name: name,
	}
}

func (c *rackTypeService) AddRackType(ctx context.Context, in *RackType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.AddRackType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) DeleteRackType(ctx context.Context, in *IDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.DeleteRackType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) UpdateRackType(ctx context.Context, in *RackType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.UpdateRackType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) GetOneRackTypes(ctx context.Context, in *RackTypeQuery, opts ...client.CallOption) (*RackType, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.GetOneRackTypes", in)
	out := new(RackType)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) GetRackTypes(ctx context.Context, in *IDReq, opts ...client.CallOption) (*RackTypes, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.GetRackTypes", in)
	out := new(RackTypes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) AddCellType(ctx context.Context, in *CellType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.AddCellType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) DeleteCellType(ctx context.Context, in *IDReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.DeleteCellType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) UpdateCellType(ctx context.Context, in *CellType, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.UpdateCellType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) GetCellTypes(ctx context.Context, in *IDReq, opts ...client.CallOption) (*CellType, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.GetCellTypes", in)
	out := new(CellType)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rackTypeService) GetOneCellTypes(ctx context.Context, in *CellTypeQuery, opts ...client.CallOption) (*CellTypes, error) {
	req := c.c.NewRequest(c.name, "RackTypeService.GetOneCellTypes", in)
	out := new(CellTypes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RackTypeService service

type RackTypeServiceHandler interface {
	AddRackType(context.Context, *RackType, *Response) error
	DeleteRackType(context.Context, *IDReq, *Response) error
	UpdateRackType(context.Context, *RackType, *Response) error
	GetOneRackTypes(context.Context, *RackTypeQuery, *RackType) error
	GetRackTypes(context.Context, *IDReq, *RackTypes) error
	AddCellType(context.Context, *CellType, *Response) error
	DeleteCellType(context.Context, *IDReq, *Response) error
	UpdateCellType(context.Context, *CellType, *Response) error
	GetCellTypes(context.Context, *IDReq, *CellType) error
	GetOneCellTypes(context.Context, *CellTypeQuery, *CellTypes) error
}

func RegisterRackTypeServiceHandler(s server.Server, hdlr RackTypeServiceHandler, opts ...server.HandlerOption) error {
	type rackTypeService interface {
		AddRackType(ctx context.Context, in *RackType, out *Response) error
		DeleteRackType(ctx context.Context, in *IDReq, out *Response) error
		UpdateRackType(ctx context.Context, in *RackType, out *Response) error
		GetOneRackTypes(ctx context.Context, in *RackTypeQuery, out *RackType) error
		GetRackTypes(ctx context.Context, in *IDReq, out *RackTypes) error
		AddCellType(ctx context.Context, in *CellType, out *Response) error
		DeleteCellType(ctx context.Context, in *IDReq, out *Response) error
		UpdateCellType(ctx context.Context, in *CellType, out *Response) error
		GetCellTypes(ctx context.Context, in *IDReq, out *CellType) error
		GetOneCellTypes(ctx context.Context, in *CellTypeQuery, out *CellTypes) error
	}
	type RackTypeService struct {
		rackTypeService
	}
	h := &rackTypeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RackTypeService{h}, opts...))
}

type rackTypeServiceHandler struct {
	RackTypeServiceHandler
}

func (h *rackTypeServiceHandler) AddRackType(ctx context.Context, in *RackType, out *Response) error {
	return h.RackTypeServiceHandler.AddRackType(ctx, in, out)
}

func (h *rackTypeServiceHandler) DeleteRackType(ctx context.Context, in *IDReq, out *Response) error {
	return h.RackTypeServiceHandler.DeleteRackType(ctx, in, out)
}

func (h *rackTypeServiceHandler) UpdateRackType(ctx context.Context, in *RackType, out *Response) error {
	return h.RackTypeServiceHandler.UpdateRackType(ctx, in, out)
}

func (h *rackTypeServiceHandler) GetOneRackTypes(ctx context.Context, in *RackTypeQuery, out *RackType) error {
	return h.RackTypeServiceHandler.GetOneRackTypes(ctx, in, out)
}

func (h *rackTypeServiceHandler) GetRackTypes(ctx context.Context, in *IDReq, out *RackTypes) error {
	return h.RackTypeServiceHandler.GetRackTypes(ctx, in, out)
}

func (h *rackTypeServiceHandler) AddCellType(ctx context.Context, in *CellType, out *Response) error {
	return h.RackTypeServiceHandler.AddCellType(ctx, in, out)
}

func (h *rackTypeServiceHandler) DeleteCellType(ctx context.Context, in *IDReq, out *Response) error {
	return h.RackTypeServiceHandler.DeleteCellType(ctx, in, out)
}

func (h *rackTypeServiceHandler) UpdateCellType(ctx context.Context, in *CellType, out *Response) error {
	return h.RackTypeServiceHandler.UpdateCellType(ctx, in, out)
}

func (h *rackTypeServiceHandler) GetCellTypes(ctx context.Context, in *IDReq, out *CellType) error {
	return h.RackTypeServiceHandler.GetCellTypes(ctx, in, out)
}

func (h *rackTypeServiceHandler) GetOneCellTypes(ctx context.Context, in *CellTypeQuery, out *CellTypes) error {
	return h.RackTypeServiceHandler.GetOneCellTypes(ctx, in, out)
}
