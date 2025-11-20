package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

func (c *ControllerV1) UnfavorPipeline(ctx context.Context, req *v1.UnfavorPipelineReq) (res *v1.UnfavorPipelineRes, err error) {
	res = &v1.UnfavorPipelineRes{}

	logic := pipelines.NewPipelines()
	if err = logic.UnfavorPipeline(ctx, &pipelines.UnfavorPipelineInput{
		PipelineId: req.PipelineId,
	}); err != nil {
		return nil, err
	}

	res.Message = "取消收藏成功"
	return
}
