// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tasks

import (
	"context"

	"kgplatform-backend/api/tasks/v1"
)

type ITasksV1 interface {
	CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error)
	GetTask(ctx context.Context, req *v1.GetTaskReq) (res *v1.GetTaskRes, err error)
	CreateOCRTask(ctx context.Context, req *v1.CreateOCRTaskReq) (res *v1.CreateOCRTaskRes, err error)
	CreateExtractTask(ctx context.Context, req *v1.CreateExtractTaskReq) (res *v1.CreateExtractTaskRes, err error)
	CreateGraphTask(ctx context.Context, req *v1.CreateGraphTaskReq) (res *v1.CreateGraphTaskRes, err error)
	ListTask(ctx context.Context, req *v1.ListTaskReq) (res *v1.ListTaskRes, err error)
}
