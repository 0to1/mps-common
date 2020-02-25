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

// Client API for Station service

type StationService interface {
	//Get station message according to station id
	GetStationByID(ctx context.Context, in *IdRequest, opts ...client.CallOption) (*StationResponse, error)
	//Get all stations message
	GetStations(ctx context.Context, in *Query, opts ...client.CallOption) (*StationsResponse, error)
	//Get stations by graphql
	GetStationsByGraphql(ctx context.Context, in *GraphqlQuery, opts ...client.CallOption) (*GraphqlStations, error)
}

type stationService struct {
	c    client.Client
	name string
}

func NewStationService(name string, c client.Client) StationService {
	return &stationService{
		c:    c,
		name: name,
	}
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

func (c *stationService) GetStationsByGraphql(ctx context.Context, in *GraphqlQuery, opts ...client.CallOption) (*GraphqlStations, error) {
	req := c.c.NewRequest(c.name, "Station.GetStationsByGraphql", in)
	out := new(GraphqlStations)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Station service

type StationHandler interface {
	//Get station message according to station id
	GetStationByID(context.Context, *IdRequest, *StationResponse) error
	//Get all stations message
	GetStations(context.Context, *Query, *StationsResponse) error
	//Get stations by graphql
	GetStationsByGraphql(context.Context, *GraphqlQuery, *GraphqlStations) error
}

func RegisterStationHandler(s server.Server, hdlr StationHandler, opts ...server.HandlerOption) error {
	type station interface {
		GetStationByID(ctx context.Context, in *IdRequest, out *StationResponse) error
		GetStations(ctx context.Context, in *Query, out *StationsResponse) error
		GetStationsByGraphql(ctx context.Context, in *GraphqlQuery, out *GraphqlStations) error
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

func (h *stationHandler) GetStationByID(ctx context.Context, in *IdRequest, out *StationResponse) error {
	return h.StationHandler.GetStationByID(ctx, in, out)
}

func (h *stationHandler) GetStations(ctx context.Context, in *Query, out *StationsResponse) error {
	return h.StationHandler.GetStations(ctx, in, out)
}

func (h *stationHandler) GetStationsByGraphql(ctx context.Context, in *GraphqlQuery, out *GraphqlStations) error {
	return h.StationHandler.GetStationsByGraphql(ctx, in, out)
}
