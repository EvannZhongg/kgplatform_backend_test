package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) ListProject(ctx context.Context, req *v1.ListProjectReq) (res *v1.ListProjectRes, err error) {
	logic := projects.NewProjects()
	projectOutput, err := logic.ListProject(ctx, &projects.ListProjectInput{
		Page:       req.Page,
		Visibility: req.Visibility,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ListProjectRes{
		Total:    projectOutput.Total,
		Projects: projectOutput.Projects,
	}, nil
}
