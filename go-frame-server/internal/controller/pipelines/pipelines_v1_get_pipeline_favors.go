package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

func (c *ControllerV1) GetPipelineFavors(ctx context.Context, req *v1.GetPipelineFavorsReq) (res *v1.GetPipelineFavorsRes, err error) {
	res = &v1.GetPipelineFavorsRes{}

	logic := pipelines.NewPipelines()
	favorCount, isFavored, err := logic.GetPipelineFavors(ctx, &pipelines.GetPipelineFavorsInput{
		PipelineId: req.PipelineId,
	})
	if err != nil {
		return nil, err
	}

	res.FavorCount = favorCount
	res.IsFavored = isFavored
	return
}
