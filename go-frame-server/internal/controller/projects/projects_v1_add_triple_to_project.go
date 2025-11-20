package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) AddTripleToProject(ctx context.Context, req *v1.AddTripleToProjectReq) (res *v1.AddTripleToProjectRes, err error) {
	projectLogic := projects.NewProjects()
	_, err = projectLogic.UpdateTriples(ctx, &projects.UpdateTriplesInput{
		ProjectId:  req.ProjectId,
		Triples:    req.Triples,
		TripleType: req.TripleType,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
