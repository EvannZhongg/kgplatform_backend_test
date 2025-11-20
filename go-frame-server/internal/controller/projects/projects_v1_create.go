package projects

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) CreateProject(ctx context.Context, req *v1.CreateProjectReq) (res *v1.CreateProjectRes, err error) {
	output, err := c.projects.CreateProject(ctx, req.ProjectName, req.Visibility)
	if err != nil {
		return nil, gerror.New("创建项目失败")
	}

	return &v1.CreateProjectRes{
		ProjectId:  output.Id,
		PipelineId: output.PipelineId,
	}, nil
}
