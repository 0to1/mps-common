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

// Client API for TaskService service

type TaskService interface {
	//新任务
	New(ctx context.Context, in *NewTaskReq, opts ...client.CallOption) (*Response, error)
	// 暂停任务
	Pause(ctx context.Context, in *TaskID, opts ...client.CallOption) (*Response, error)
	// 取消任务
	Cancel(ctx context.Context, in *TaskID, opts ...client.CallOption) (*Response, error)
	UpdateTask(ctx context.Context, in *UpdateTaskReq, opts ...client.CallOption) (*Response, error)
	BatchNew(ctx context.Context, in *NewTasksReq, opts ...client.CallOption) (*Response, error)
	BatchCancel(ctx context.Context, in *TaskIDs, opts ...client.CallOption) (*Response, error)
	// 增加任务参数
	AddParameters(ctx context.Context, in *Parameters, opts ...client.CallOption) (*Parameters, error)
	// 修改任务参数
	ModifyParameters(ctx context.Context, in *Parameters, opts ...client.CallOption) (*Parameters, error)
	// 删除任务参数
	DeleteParameters(ctx context.Context, in *ParameterKeys, opts ...client.CallOption) (*Parameters, error)
	// 根据查询条件获取任务
	GetOneTask(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error)
	// 根据查询条件获取历史任务
	GetOneHistory(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error)
	// 根据查询条件获取任务
	GetTasks(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error)
	// 根据查询条件获取历史任务
	GetHistories(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error)
	// 调用python脚本的函数
	Call(ctx context.Context, in *CallOptions, opts ...client.CallOption) (*CallResponse, error)
	// 执行python脚本
	Execute(ctx context.Context, in *ExecuteOptions, opts ...client.CallOption) (*ExecuteResponse, error)
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

func (c *taskService) Cancel(ctx context.Context, in *TaskID, opts ...client.CallOption) (*Response, error) {
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

func (c *taskService) AddParameters(ctx context.Context, in *Parameters, opts ...client.CallOption) (*Parameters, error) {
	req := c.c.NewRequest(c.name, "TaskService.AddParameters", in)
	out := new(Parameters)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) ModifyParameters(ctx context.Context, in *Parameters, opts ...client.CallOption) (*Parameters, error) {
	req := c.c.NewRequest(c.name, "TaskService.ModifyParameters", in)
	out := new(Parameters)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) DeleteParameters(ctx context.Context, in *ParameterKeys, opts ...client.CallOption) (*Parameters, error) {
	req := c.c.NewRequest(c.name, "TaskService.DeleteParameters", in)
	out := new(Parameters)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetOneTask(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetOneTask", in)
	out := new(TaskInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetOneHistory(ctx context.Context, in *TaskID, opts ...client.CallOption) (*TaskInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetOneHistory", in)
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

func (c *taskService) GetHistories(ctx context.Context, in *Query, opts ...client.CallOption) (*TaskInfos, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetHistories", in)
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

// Server API for TaskService service

type TaskServiceHandler interface {
	//新任务
	New(context.Context, *NewTaskReq, *Response) error
	// 暂停任务
	Pause(context.Context, *TaskID, *Response) error
	// 取消任务
	Cancel(context.Context, *TaskID, *Response) error
	UpdateTask(context.Context, *UpdateTaskReq, *Response) error
	BatchNew(context.Context, *NewTasksReq, *Response) error
	BatchCancel(context.Context, *TaskIDs, *Response) error
	// 增加任务参数
	AddParameters(context.Context, *Parameters, *Parameters) error
	// 修改任务参数
	ModifyParameters(context.Context, *Parameters, *Parameters) error
	// 删除任务参数
	DeleteParameters(context.Context, *ParameterKeys, *Parameters) error
	// 根据查询条件获取任务
	GetOneTask(context.Context, *TaskID, *TaskInfo) error
	// 根据查询条件获取历史任务
	GetOneHistory(context.Context, *TaskID, *TaskInfo) error
	// 根据查询条件获取任务
	GetTasks(context.Context, *Query, *TaskInfos) error
	// 根据查询条件获取历史任务
	GetHistories(context.Context, *Query, *TaskInfos) error
	// 调用python脚本的函数
	Call(context.Context, *CallOptions, *CallResponse) error
	// 执行python脚本
	Execute(context.Context, *ExecuteOptions, *ExecuteResponse) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		New(ctx context.Context, in *NewTaskReq, out *Response) error
		Pause(ctx context.Context, in *TaskID, out *Response) error
		Cancel(ctx context.Context, in *TaskID, out *Response) error
		UpdateTask(ctx context.Context, in *UpdateTaskReq, out *Response) error
		BatchNew(ctx context.Context, in *NewTasksReq, out *Response) error
		BatchCancel(ctx context.Context, in *TaskIDs, out *Response) error
		AddParameters(ctx context.Context, in *Parameters, out *Parameters) error
		ModifyParameters(ctx context.Context, in *Parameters, out *Parameters) error
		DeleteParameters(ctx context.Context, in *ParameterKeys, out *Parameters) error
		GetOneTask(ctx context.Context, in *TaskID, out *TaskInfo) error
		GetOneHistory(ctx context.Context, in *TaskID, out *TaskInfo) error
		GetTasks(ctx context.Context, in *Query, out *TaskInfos) error
		GetHistories(ctx context.Context, in *Query, out *TaskInfos) error
		Call(ctx context.Context, in *CallOptions, out *CallResponse) error
		Execute(ctx context.Context, in *ExecuteOptions, out *ExecuteResponse) error
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

func (h *taskServiceHandler) Cancel(ctx context.Context, in *TaskID, out *Response) error {
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

func (h *taskServiceHandler) AddParameters(ctx context.Context, in *Parameters, out *Parameters) error {
	return h.TaskServiceHandler.AddParameters(ctx, in, out)
}

func (h *taskServiceHandler) ModifyParameters(ctx context.Context, in *Parameters, out *Parameters) error {
	return h.TaskServiceHandler.ModifyParameters(ctx, in, out)
}

func (h *taskServiceHandler) DeleteParameters(ctx context.Context, in *ParameterKeys, out *Parameters) error {
	return h.TaskServiceHandler.DeleteParameters(ctx, in, out)
}

func (h *taskServiceHandler) GetOneTask(ctx context.Context, in *TaskID, out *TaskInfo) error {
	return h.TaskServiceHandler.GetOneTask(ctx, in, out)
}

func (h *taskServiceHandler) GetOneHistory(ctx context.Context, in *TaskID, out *TaskInfo) error {
	return h.TaskServiceHandler.GetOneHistory(ctx, in, out)
}

func (h *taskServiceHandler) GetTasks(ctx context.Context, in *Query, out *TaskInfos) error {
	return h.TaskServiceHandler.GetTasks(ctx, in, out)
}

func (h *taskServiceHandler) GetHistories(ctx context.Context, in *Query, out *TaskInfos) error {
	return h.TaskServiceHandler.GetHistories(ctx, in, out)
}

func (h *taskServiceHandler) Call(ctx context.Context, in *CallOptions, out *CallResponse) error {
	return h.TaskServiceHandler.Call(ctx, in, out)
}

func (h *taskServiceHandler) Execute(ctx context.Context, in *ExecuteOptions, out *ExecuteResponse) error {
	return h.TaskServiceHandler.Execute(ctx, in, out)
}
