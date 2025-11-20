package pipelines

import (
	"context"
	"kgplatform-backend/api/pipelines/v1"
	"kgplatform-backend/internal/logic/pipelines"
)

// GetPipelineLikes 获取工作流点赞数
func (c *ControllerV1) GetPipelineLikes(ctx context.Context, req *v1.GetPipelineLikesReq) (res *v1.GetPipelineLikesRes, err error) {
	logic := pipelines.NewPipelines()
	likeCount, isLiked, err := logic.GetPipelineLikes(ctx, &pipelines.GetPipelineLikesInput{
		PipelineId: req.PipelineId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetPipelineLikesRes{
		LikeCount: likeCount,
		IsLiked:   isLiked,
	}, nil
}
