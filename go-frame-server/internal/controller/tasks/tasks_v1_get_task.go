package tasks

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/tasks"

	"github.com/gogf/gf/v2/errors/gerror"
	"kgplatform-backend/api/tasks/v1"
)

func (c *ControllerV1) GetTask(ctx context.Context, req *v1.GetTaskReq) (res *v1.GetTaskRes, err error) {
	taskLogic := tasks.New()
	task, err := taskLogic.Get(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "获取任务失败: %v", err)
		return nil, gerror.Newf("获取任务失败")
	}
	res = &v1.GetTaskRes{
		Id:         task.Id,
		Type:       task.Type,
		Status:     task.Status,
		StartTime:  task.StartTime.String(),
		FinishTime: task.FinishTime.String(),
		CreateTime: task.CreatedAt.String(),
	}
	return res, nil
}
