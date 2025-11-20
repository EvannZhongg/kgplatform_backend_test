package tasks

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"kgplatform-backend/internal/consts"
	"kgplatform-backend/internal/logic/tasks"
	"slices"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/tasks/v1"
)

var TaskTypeArray = []string{consts.TaskTypeOCR, consts.TaskTypeExtract, consts.TaskTypeGraph}

func (c *ControllerV1) CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error) {
	if req.Type == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "请选择任务类型")
	}
	if !slices.Contains(TaskTypeArray, req.Type) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "不支持的任务类型")
	}
	taskLogic := tasks.New()
	input := &tasks.CreateTaskInput{
		Type:           req.Type,
		PipelineId:     req.PipelineId,
		ProjectId:      req.ProjectId,
		MaterialIdList: req.MaterialIdList,
		Status:         consts.TaskStatusPending,
		UpdatedAt:      gtime.Now(),
		CreatedAt:      gtime.Now(),
		Prompt:         req.Prompt,
		ModelId:        req.ModelId,
	}
	taskEntity, err := taskLogic.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.CreateTaskRes{
		Status: taskEntity.Status,
		TaskId: taskEntity.Id,
		Type:   taskEntity.Type,
	}
	return res, nil
}
