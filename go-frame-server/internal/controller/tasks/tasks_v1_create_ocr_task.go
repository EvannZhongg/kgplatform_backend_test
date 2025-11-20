package tasks

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"kgplatform-backend/api/tasks/v1"
	"kgplatform-backend/internal/consts"
	"kgplatform-backend/internal/logic/tasks"
)

func (c *ControllerV1) CreateOCRTask(ctx context.Context, req *v1.CreateOCRTaskReq) (res *v1.CreateOCRTaskRes, err error) {
	taskLogic := tasks.New()
	input := &tasks.CreateTaskInput{
		Type:           consts.TaskTypeOCR,
		PipelineId:     req.PipelineId,
		ProjectId:      req.ProjectId,
		MaterialIdList: req.MaterialIDList,
		Status:         consts.TaskStatusPending,
		UpdatedAt:      gtime.Now(),
		CreatedAt:      gtime.Now(),
	}
	taskEntity, err := taskLogic.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.CreateOCRTaskRes{
		Status: taskEntity.Status,
		TaskId: taskEntity.Id,
	}
	return res, nil
}
