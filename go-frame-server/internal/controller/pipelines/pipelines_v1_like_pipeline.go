package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

// LikePipeline 点赞工作流
func (c *ControllerV1) LikePipeline(ctx context.Context, req *v1.LikePipelineReq) (res *v1.LikePipelineRes, err error) {
	logic := pipelines.NewPipelines()
	err = logic.LikePipeline(ctx, &pipelines.LikePipelineInput{
		PipelineId: req.PipelineId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.LikePipelineRes{
		Message: "点赞成功",
	}, nil
}
