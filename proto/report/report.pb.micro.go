// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/report/report.proto

package go_micro_srv_report

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

// Api Endpoints for ReportService service

func NewReportServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ReportService service

type ReportService interface {
	GetTaskReport(ctx context.Context, in *TaskReportReq, opts ...client.CallOption) (*TaskReportItems, error)
	GetAgvTaskReport(ctx context.Context, in *AgvTaskReportReq, opts ...client.CallOption) (*AgvTaskReportItems, error)
}

type reportService struct {
	c    client.Client
	name string
}

func NewReportService(name string, c client.Client) ReportService {
	return &reportService{
		c:    c,
		name: name,
	}
}

func (c *reportService) GetTaskReport(ctx context.Context, in *TaskReportReq, opts ...client.CallOption) (*TaskReportItems, error) {
	req := c.c.NewRequest(c.name, "ReportService.GetTaskReport", in)
	out := new(TaskReportItems)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportService) GetAgvTaskReport(ctx context.Context, in *AgvTaskReportReq, opts ...client.CallOption) (*AgvTaskReportItems, error) {
	req := c.c.NewRequest(c.name, "ReportService.GetAgvTaskReport", in)
	out := new(AgvTaskReportItems)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReportService service

type ReportServiceHandler interface {
	GetTaskReport(context.Context, *TaskReportReq, *TaskReportItems) error
	GetAgvTaskReport(context.Context, *AgvTaskReportReq, *AgvTaskReportItems) error
}

func RegisterReportServiceHandler(s server.Server, hdlr ReportServiceHandler, opts ...server.HandlerOption) error {
	type reportService interface {
		GetTaskReport(ctx context.Context, in *TaskReportReq, out *TaskReportItems) error
		GetAgvTaskReport(ctx context.Context, in *AgvTaskReportReq, out *AgvTaskReportItems) error
	}
	type ReportService struct {
		reportService
	}
	h := &reportServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ReportService{h}, opts...))
}

type reportServiceHandler struct {
	ReportServiceHandler
}

func (h *reportServiceHandler) GetTaskReport(ctx context.Context, in *TaskReportReq, out *TaskReportItems) error {
	return h.ReportServiceHandler.GetTaskReport(ctx, in, out)
}

func (h *reportServiceHandler) GetAgvTaskReport(ctx context.Context, in *AgvTaskReportReq, out *AgvTaskReportItems) error {
	return h.ReportServiceHandler.GetAgvTaskReport(ctx, in, out)
}
