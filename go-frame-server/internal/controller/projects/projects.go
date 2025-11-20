package projects

import (
	"kgplatform-backend/api/projects"
	projectsLogic "kgplatform-backend/internal/logic/projects"
)

type ControllerV1 struct {
	projects *projectsLogic.Projects
}

func NewV1() projects.IProjectsV1 {
	return &ControllerV1{
		projects: projectsLogic.NewProjects(),
	}
}
