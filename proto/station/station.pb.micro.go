// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/station/station.proto

package go_micro_srv_station

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

// Client API for Station service

type StationService interface {
	//save station file
	SaveStationFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (*FileResponse, error)
	//Get station message according to station id
	GetStationByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*StationResponse, error)
	//Get all stations message
	GetStations(ctx context.Context, in *Query, opts ...client.CallOption) (*StationsResponse, error)
}

type stationService struct {
	c    client.Client
	name string
}

func NewStationService(name string, c client.Client) StationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.station"
	}
	return &stationService{
		c:    c,
		name: name,
	}
}

func (c *stationService) SaveStationFile(ctx context.Context, in *FileRequest, opts ...client.CallOption) (*FileResponse, error) {
	req := c.c.NewRequest(c.name, "Station.SaveStationFile", in)
	out := new(FileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stationService) GetStationByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*StationResponse, error) {
	req := c.c.NewRequest(c.name, "Station.GetStationByID", in)
	out := new(StationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stationService) GetStations(ctx context.Context, in *Query, opts ...client.CallOption) (*StationsResponse, error) {
	req := c.c.NewRequest(c.name, "Station.GetStations", in)
	out := new(StationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Station service

type StationHandler interface {
	//save station file
	SaveStationFile(context.Context, *FileRequest, *FileResponse) error
	//Get station message according to station id
	GetStationByID(context.Context, *IdRequest, *StationResponse) error
	//Get all stations message
	GetStations(context.Context, *Query, *StationsResponse) error
}

func RegisterStationHandler(s server.Server, hdlr StationHandler, opts ...server.HandlerOption) error {
	type station interface {
		SaveStationFile(ctx context.Context, in *FileRequest, out *FileResponse) error
		GetStationByID(ctx context.Context, in *IdRequest, out *StationResponse) error
		GetStations(ctx context.Context, in *Query, out *StationsResponse) error
	}
	type Station struct {
		station
	}
	h := &stationHandler{hdlr}
	return s.Handle(s.NewHandler(&Station{h}, opts...))
}

type stationHandler struct {
	StationHandler
}

func (h *stationHandler) SaveStationFile(ctx context.Context, in *FileRequest, out *FileResponse) error {
	return h.StationHandler.SaveStationFile(ctx, in, out)
}

func (h *stationHandler) GetStationByID(ctx context.Context, in *IdRequest, out *StationResponse) error {
	return h.StationHandler.GetStationByID(ctx, in, out)
}

func (h *stationHandler) GetStations(ctx context.Context, in *Query, out *StationsResponse) error {
	return h.StationHandler.GetStations(ctx, in, out)
}