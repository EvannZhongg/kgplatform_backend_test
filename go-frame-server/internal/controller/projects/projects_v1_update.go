package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

// UpdateProject 更新项目
func (c *ControllerV1) UpdateProject(ctx context.Context, req *v1.UpdateProjectReq) (res *v1.UpdateProjectRes, err error) {
	input := &projects.UpdateProjectInput{
		ProjectName:     req.ProjectName,
		ProjectProgress: req.ProjectProgress,
		Visibility:      req.Visibility,
	}
	err = c.projects.UpdateProject(ctx, req.Id, input)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateProjectRes{
		Success: true,
	}, nil
}
