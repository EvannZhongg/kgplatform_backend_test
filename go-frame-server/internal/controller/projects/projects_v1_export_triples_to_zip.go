package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) ExportTriplesToZip(ctx context.Context, req *v1.ExportTriplesToZipReq) (res *v1.ExportTriplesToZipRes, err error) {
	projectLogic := projects.NewProjects()
	projectLogic.ExportTriplesToExcelZip(ctx, &projects.ExportTriplesToExcelZipInput{
		ProjectId: req.ProjectId,
	})
	return nil, nil
}
