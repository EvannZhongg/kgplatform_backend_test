package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

func (c *ControllerV1) CreatePipeline(ctx context.Context, req *v1.CreatePipelineReq) (res *v1.CreatePipelineRes, err error) {
	logic := pipelines.NewPipelines()
	pipeline, err := logic.CreatePipeline(ctx, &pipelines.CreatePipelineInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreatePipelineRes{
		Id:        pipeline.Id,
		ProjectId: pipeline.ProjectId,
	}, nil
}
