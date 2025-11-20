package tasks

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"kgplatform-backend/api/tasks/v1"
	"kgplatform-backend/internal/logic/tasks"
)

func (c *ControllerV1) ListTask(ctx context.Context, req *v1.ListTaskReq) (res *v1.ListTaskRes, err error) {
	if req.ProjectId == nil && req.PipelineId == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "请选择项目或管道")
	}
	taskLogic := tasks.New()
	output, err := taskLogic.ListTask(ctx, &tasks.ListTaskInput{
		Page:       req.Page,
		PipelineId: req.PipelineId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ListTaskRes{
		Total: output.Total,
		List:  output.Tasks,
	}, nil
}
