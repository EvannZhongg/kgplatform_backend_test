package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

// GetProjectEntities 获取项目实体结构
func (c *ControllerV1) GetProjectEntities(ctx context.Context, req *v1.GetProjectEntitiesReq) (*v1.GetProjectEntitiesRes, error) {
	projectLogic := projects.NewProjects()
	output, err := projectLogic.GetProjectEntities(ctx, &projects.GetProjectEntitiesInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	res := &v1.GetProjectEntitiesRes{
		TripleEntities: output.TripleEntities,
	}
	return res, nil
}
