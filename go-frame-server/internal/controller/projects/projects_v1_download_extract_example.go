package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) DownloadExtractExample(ctx context.Context, req *v1.DownloadExtractExampleReq) (res *v1.DownloadExtractExampleRes, err error) {
	projectLogic := projects.NewProjects()
	projectLogic.DownloadExtractExample(ctx)
	return nil, nil
}
