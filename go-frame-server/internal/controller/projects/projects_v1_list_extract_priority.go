package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) ListExtractPriority(ctx context.Context, req *v1.ListExtractPriorityReq) (res *v1.ListExtractPriorityRes, err error) {
	projectLogic := projects.NewProjects()
	priority, err := projectLogic.GetProjectExtractPriority(ctx, &projects.GetProjectExtractPriorityInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ListExtractPriorityRes{
		EntityList:       priority.EntityList,
		RelationshipList: priority.RelationshipList,
	}, nil
}
