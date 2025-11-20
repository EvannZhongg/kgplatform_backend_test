package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

// UnlikePipeline 取消点赞工作流
func (c *ControllerV1) UnlikePipeline(ctx context.Context, req *v1.UnlikePipelineReq) (res *v1.UnlikePipelineRes, err error) {
	logic := pipelines.NewPipelines()
	err = logic.UnlikePipeline(ctx, &pipelines.UnlikePipelineInput{
		PipelineId: req.PipelineId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UnlikePipelineRes{
		Message: "取消点赞成功",
	}, nil
}
