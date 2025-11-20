package projects

import (
	"context"

	v1 "kgplatform-backend/api/projects/v1"
	"kgplatform-backend/internal/logic/projects"
)

func (c *ControllerV1) GetTripleSourceInfo(ctx context.Context, req *v1.GetTripleSourceInfoReq) (res *v1.GetTripleSourceInfoRes, err error) {
	result, err := c.projects.GetTripleSourceInfo(ctx, &projects.GetTripleSourceInfoInput{
		ProjectId: req.ProjectId,
		Triple:    req.Triple,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetTripleSourceInfoRes{
		MaterialId:   result.MaterialId,
		MaterialName: result.MaterialName,
		SourceText:   result.SourceText,
		ChunkIndex:   result.ChunkIndex,
	}, nil
}
