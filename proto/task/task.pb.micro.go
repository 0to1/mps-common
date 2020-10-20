// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/task/task.proto

package go_micro_srv_task

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

// Api Endpoints for TaskService service

func NewTaskServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TaskService service

type TaskService interface {
	//新任务
	New(ctx context.Context, in *NewTaskReq, opts ...client.CallOption) (*Response, error)
	// 暂停任务
	Pause(ctx context.Context, in *TaskID, opts ...client.CallOption) (*Response, error)
	// 取消任务
	Cancel(ctx context.Context, in *CancelReq, opts ...client.CallOption) (*Response, error)
	UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...client.CallOption) (*Response, error)
	BatchNew(ctx context.Context, in *NewTasksReq, opts ...client.CallOption) (*Response, error)
	BatchCancel(ctx context.Context, in *TaskIDs, opts ...client.CallOption) (*Response, error)
	// 设置任务参数
	SetParameters(ctx context.Context, in *LpParameters, opts ...client.CallOption) (*Response, error)
	// 删除任务参数
	DeleteParameters(ctx context.Context, in *LpParameterKeys, opts ...client.CallOption) (*Response, error)
	// 根据查询条件获取任务
	GetTask(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error)
	// 根据查询条件获取历史任务
	GetHistoryTask(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error)
	// 根据查询条件获取任务
	GetTasks(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error)
	// 根据查询条件获取历史任务
	GetHistoryTasks(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error)
	// 调用脚本的函数
	Call(ctx context.Context, in *CallOptions, opts ...client.CallOption) (*CallResponse, error)
	// 执行脚本
	Execute(ctx context.Context, in *ExecuteOptions, opts ...client.CallOption) (*ExecuteResponse, error)
	SetGP(ctx context.Context, in *GpParameter, opts ...client.CallOption) (*Response, error)
	DelGP(ctx context.Context, in *GPKey, opts ...client.CallOption) (*Response, error)
	GetGP(ctx context.Context, in *GPKey, opts ...client.CallOption) (*GpParameter, error)
	GetGPs(ctx context.Context, in *GPQuery, opts ...client.CallOption) (*GpParameters, error)
	ConfigSystemGo(ctx context.Context, in *SystemGoReq, opts ...client.CallOption) (*Response, error)
	RestartSystemGo(ctx context.Context, in *Nop, opts ...client.CallOption) (*Response, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) New(ctx context.Context, in *NewTaskReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.New", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Pause(ctx context.Context, in *TaskID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.Pause", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Cancel(ctx context.Context, in *CancelReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.Cancel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateTask", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) BatchNew(ctx context.Context, in *NewTasksReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.BatchNew", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) BatchCancel(ctx context.Context, in *TaskIDs, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.BatchCancel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) SetParameters(ctx context.Context, in *LpParameters, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.SetParameters", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) DeleteParameters(ctx context.Context, in *LpParameterKeys, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.DeleteParameters", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetTask(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetTask", in)
	out := new(TaskInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetHistoryTask(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetHistoryTask", in)
	out := new(TaskInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetTasks(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetTasks", in)
	out := new(TaskInfos)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetHistoryTasks(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetHistoryTasks", in)
	out := new(TaskInfos)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Call(ctx context.Context, in *CallOptions, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) Execute(ctx context.Context, in *ExecuteOptions, opts ...client.CallOption) (*ExecuteResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Execute", in)
	out := new(ExecuteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) SetGP(ctx context.Context, in *GpParameter, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.SetGP", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) DelGP(ctx context.Context, in *GPKey, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.DelGP", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetGP(ctx context.Context, in *GPKey, opts ...client.CallOption) (*GpParameter, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetGP", in)
	out := new(GpParameter)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetGPs(ctx context.Context, in *GPQuery, opts ...client.CallOption) (*GpParameters, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetGPs", in)
	out := new(GpParameters)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) ConfigSystemGo(ctx context.Context, in *SystemGoReq, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.ConfigSystemGo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) RestartSystemGo(ctx context.Context, in *Nop, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TaskService.RestartSystemGo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	//新任务
	New(context.Context, *NewTaskReq, *Response) error
	// 暂停任务
	Pause(context.Context, *TaskID, *Response) error
	// 取消任务
	Cancel(context.Context, *CancelReq, *Response) error
	UpdateTask(context.Context, *UpdateTaskReq, *Response) error
	BatchNew(context.Context, *NewTasksReq, *Response) error
	BatchCancel(context.Context, *TaskIDs, *Response) error
	// 设置任务参数
	SetParameters(context.Context, *LpParameters, *Response) error
	// 删除任务参数
	DeleteParameters(context.Context, *LpParameterKeys, *Response) error
	// 根据查询条件获取任务
	GetTask(context.Context, *TaskID, *TaskInfo) error
	// 根据查询条件获取历史任务
	GetHistoryTask(context.Context, *TaskID, *TaskInfo) error
	// 根据查询条件获取任务
	GetTasks(context.Context, *Query, *TaskInfos) error
	// 根据查询条件获取历史任务
	GetHistoryTasks(context.Context, *Query, *TaskInfos) error
	// 调用脚本的函数
	Call(context.Context, *CallOptions, *CallResponse) error
	// 执行脚本
	Execute(context.Context, *ExecuteOptions, *ExecuteResponse) error
	SetGP(context.Context, *GpParameter, *Response) error
	DelGP(context.Context, *GPKey, *Response) error
	GetGP(context.Context, *GPKey, *GpParameter) error
	GetGPs(context.Context, *GPQuery, *GpParameters) error
	ConfigSystemGo(context.Context, *SystemGoReq, *Response) error
	RestartSystemGo(context.Context, *Nop, *Response) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		New(ctx context.Context, in *NewTaskReq, out *Response) error
		Pause(ctx context.Context, in *TaskID, out *Response) error
		Cancel(ctx context.Context, in *CancelReq, out *Response) error
		UpdateTask(ctx context.Context, in *UpdateTaskReq, out *Response) error
		BatchNew(ctx context.Context, in *NewTasksReq, out *Response) error
		BatchCancel(ctx context.Context, in *TaskIDs, out *Response) error
		SetParameters(ctx context.Context, in *LpParameters, out *Response) error
		DeleteParameters(ctx context.Context, in *LpParameterKeys, out *Response) error
		GetTask(ctx context.Context, in *TaskID, out *TaskInfo) error
		GetHistoryTask(ctx context.Context, in *TaskID, out *TaskInfo) error
		GetTasks(ctx context.Context, in *Query, out *TaskInfos) error
		GetHistoryTasks(ctx context.Context, in *Query, out *TaskInfos) error
		Call(ctx context.Context, in *CallOptions, out *CallResponse) error
		Execute(ctx context.Context, in *ExecuteOptions, out *ExecuteResponse) error
		SetGP(ctx context.Context, in *GpParameter, out *Response) error
		DelGP(ctx context.Context, in *GPKey, out *Response) error
		GetGP(ctx context.Context, in *GPKey, out *GpParameter) error
		GetGPs(ctx context.Context, in *GPQuery, out *GpParameters) error
		ConfigSystemGo(ctx context.Context, in *SystemGoReq, out *Response) error
		RestartSystemGo(ctx context.Context, in *Nop, out *Response) error
	}
	type TaskService struct {
		taskService
	}
	h := &taskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskService{h}, opts...))
}

type taskServiceHandler struct {
	TaskServiceHandler
}

func (h *taskServiceHandler) New(ctx context.Context, in *NewTaskReq, out *Response) error {
	return h.TaskServiceHandler.New(ctx, in, out)
}

func (h *taskServiceHandler) Pause(ctx context.Context, in *TaskID, out *Response) error {
	return h.TaskServiceHandler.Pause(ctx, in, out)
}

func (h *taskServiceHandler) Cancel(ctx context.Context, in *CancelReq, out *Response) error {
	return h.TaskServiceHandler.Cancel(ctx, in, out)
}

func (h *taskServiceHandler) UpdateTask(ctx context.Context, in *UpdateTaskReq, out *Response) error {
	return h.TaskServiceHandler.UpdateTask(ctx, in, out)
}

func (h *taskServiceHandler) BatchNew(ctx context.Context, in *NewTasksReq, out *Response) error {
	return h.TaskServiceHandler.BatchNew(ctx, in, out)
}

func (h *taskServiceHandler) BatchCancel(ctx context.Context, in *TaskIDs, out *Response) error {
	return h.TaskServiceHandler.BatchCancel(ctx, in, out)
}

func (h *taskServiceHandler) SetParameters(ctx context.Context, in *LpParameters, out *Response) error {
	return h.TaskServiceHandler.SetParameters(ctx, in, out)
}

func (h *taskServiceHandler) DeleteParameters(ctx context.Context, in *LpParameterKeys, out *Response) error {
	return h.TaskServiceHandler.DeleteParameters(ctx, in, out)
}

func (h *taskServiceHandler) GetTask(ctx context.Context, in *TaskID, out *TaskInfo) error {
	return h.TaskServiceHandler.GetTask(ctx, in, out)
}

func (h *taskServiceHandler) GetHistoryTask(ctx context.Context, in *TaskID, out *TaskInfo) error {
	return h.TaskServiceHandler.GetHistoryTask(ctx, in, out)
}

func (h *taskServiceHandler) GetTasks(ctx context.Context, in *Query, out *TaskInfos) error {
	return h.TaskServiceHandler.GetTasks(ctx, in, out)
}

func (h *taskServiceHandler) GetHistoryTasks(ctx context.Context, in *Query, out *TaskInfos) error {
	return h.TaskServiceHandler.GetHistoryTasks(ctx, in, out)
}

func (h *taskServiceHandler) Call(ctx context.Context, in *CallOptions, out *CallResponse) error {
	return h.TaskServiceHandler.Call(ctx, in, out)
}

func (h *taskServiceHandler) Execute(ctx context.Context, in *ExecuteOptions, out *ExecuteResponse) error {
	return h.TaskServiceHandler.Execute(ctx, in, out)
}

func (h *taskServiceHandler) SetGP(ctx context.Context, in *GpParameter, out *Response) error {
	return h.TaskServiceHandler.SetGP(ctx, in, out)
}

func (h *taskServiceHandler) DelGP(ctx context.Context, in *GPKey, out *Response) error {
	return h.TaskServiceHandler.DelGP(ctx, in, out)
}

func (h *taskServiceHandler) GetGP(ctx context.Context, in *GPKey, out *GpParameter) error {
	return h.TaskServiceHandler.GetGP(ctx, in, out)
}

func (h *taskServiceHandler) GetGPs(ctx context.Context, in *GPQuery, out *GpParameters) error {
	return h.TaskServiceHandler.GetGPs(ctx, in, out)
}

func (h *taskServiceHandler) ConfigSystemGo(ctx context.Context, in *SystemGoReq, out *Response) error {
	return h.TaskServiceHandler.ConfigSystemGo(ctx, in, out)
}

func (h *taskServiceHandler) RestartSystemGo(ctx context.Context, in *Nop, out *Response) error {
	return h.TaskServiceHandler.RestartSystemGo(ctx, in, out)
}
