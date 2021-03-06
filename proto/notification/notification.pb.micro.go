// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/notification/notification.proto

package go_micro_srv_notification

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

// Api Endpoints for NotificationService service

func NewNotificationServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NotificationService service

type NotificationService interface {
	AddNotification(ctx context.Context, in *NotificationReq, opts ...client.CallOption) (*Response, error)
	GetNotification(ctx context.Context, in *IDReq, opts ...client.CallOption) (*Notification, error)
	GetNotifications(ctx context.Context, in *NotificationQuery, opts ...client.CallOption) (*Notifications, error)
	DeleteNotifications(ctx context.Context, in *NotificationFilter, opts ...client.CallOption) (*Response, error)
	// 标记通知为已读
	MarkAsRead(ctx context.Context, in *IDsReq, opts ...client.CallOption) (*Response, error)
}

type notificationService struct {
	c    client.Client
	name string
}

func NewNotificationService(name string, c client.Client) NotificationService {
	return &notificationService{
		c:    c,
		name: name,
	}
}

func (c *notificationService) AddNotification(ctx context.Context, in *NotificationReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "NotificationService.AddNotification", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) GetNotification(ctx context.Context, in *IDReq, opts ...client.CallOption) (*Notification, error) {
	req := c.c.NewRequest(c.name, "NotificationService.GetNotification", in)
	out := new(Notification)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) GetNotifications(ctx context.Context, in *NotificationQuery, opts ...client.CallOption) (*Notifications, error) {
	req := c.c.NewRequest(c.name, "NotificationService.GetNotifications", in)
	out := new(Notifications)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) DeleteNotifications(ctx context.Context, in *NotificationFilter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "NotificationService.DeleteNotifications", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) MarkAsRead(ctx context.Context, in *IDsReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "NotificationService.MarkAsRead", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotificationService service

type NotificationServiceHandler interface {
	AddNotification(context.Context, *NotificationReq, *Response) error
	GetNotification(context.Context, *IDReq, *Notification) error
	GetNotifications(context.Context, *NotificationQuery, *Notifications) error
	DeleteNotifications(context.Context, *NotificationFilter, *Response) error
	// 标记通知为已读
	MarkAsRead(context.Context, *IDsReq, *Response) error
}

func RegisterNotificationServiceHandler(s server.Server, hdlr NotificationServiceHandler, opts ...server.HandlerOption) error {
	type notificationService interface {
		AddNotification(ctx context.Context, in *NotificationReq, out *Response) error
		GetNotification(ctx context.Context, in *IDReq, out *Notification) error
		GetNotifications(ctx context.Context, in *NotificationQuery, out *Notifications) error
		DeleteNotifications(ctx context.Context, in *NotificationFilter, out *Response) error
		MarkAsRead(ctx context.Context, in *IDsReq, out *Response) error
	}
	type NotificationService struct {
		notificationService
	}
	h := &notificationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NotificationService{h}, opts...))
}

type notificationServiceHandler struct {
	NotificationServiceHandler
}

func (h *notificationServiceHandler) AddNotification(ctx context.Context, in *NotificationReq, out *Response) error {
	return h.NotificationServiceHandler.AddNotification(ctx, in, out)
}

func (h *notificationServiceHandler) GetNotification(ctx context.Context, in *IDReq, out *Notification) error {
	return h.NotificationServiceHandler.GetNotification(ctx, in, out)
}

func (h *notificationServiceHandler) GetNotifications(ctx context.Context, in *NotificationQuery, out *Notifications) error {
	return h.NotificationServiceHandler.GetNotifications(ctx, in, out)
}

func (h *notificationServiceHandler) DeleteNotifications(ctx context.Context, in *NotificationFilter, out *Response) error {
	return h.NotificationServiceHandler.DeleteNotifications(ctx, in, out)
}

func (h *notificationServiceHandler) MarkAsRead(ctx context.Context, in *IDsReq, out *Response) error {
	return h.NotificationServiceHandler.MarkAsRead(ctx, in, out)
}
