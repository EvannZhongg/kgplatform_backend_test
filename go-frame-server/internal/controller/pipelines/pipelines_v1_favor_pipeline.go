package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

func (c *ControllerV1) FavorPipeline(ctx context.Context, req *v1.FavorPipelineReq) (res *v1.FavorPipelineRes, err error) {
	res = &v1.FavorPipelineRes{}

	logic := pipelines.NewPipelines()
	if err = logic.FavorPipeline(ctx, &pipelines.FavorPipelineInput{
		PipelineId: req.PipelineId,
	}); err != nil {
		return nil, err
	}

	res.Message = "收藏成功"
	return
}
