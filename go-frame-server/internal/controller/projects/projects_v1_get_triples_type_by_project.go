package projects

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

// GetTriplesTypeByProject 获取项目下的三元组类型
func (c *ControllerV1) GetTriplesTypeByProject(ctx context.Context, req *v1.GetTriplesTypeByProjectReq) (res *v1.GetTriplesTypeByProjectRes, err error) {
	projectLogic := projects.NewProjects()
	output, err := projectLogic.GetProjectTripleType(ctx, &projects.GetProjectTripleTypeInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, gerror.New("获取三元组失败")
	}
	res = &v1.GetTriplesTypeByProjectRes{
		TripleTypeEntity: output.TripleTypeEntity,
	}
	return res, nil
}
