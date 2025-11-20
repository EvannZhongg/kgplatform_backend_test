package projects

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"kgplatform-backend/internal/logic/projects"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) GetTriplesByType(ctx context.Context, req *v1.GetTriplesByTypeReq) (res *v1.GetTriplesByTypeRes, err error) {
	projectLogic := projects.NewProjects()
	tripleType := req.TripleType.HeadType + "-" + req.TripleType.RelationshipType + "-" + req.TripleType.TailType
	output, err := projectLogic.GetTriplesByType(ctx, &projects.GetTriplesByTypeInput{
		ProjectId: req.ProjectId,
		Type:      tripleType,
	})
	if err != nil {
		return nil, gerror.New("获取项目三元组失败, " + err.Error())
	}

	return &v1.GetTriplesByTypeRes{
		Triples:     output.TripleList,
		Percentages: output.Percentages,
	}, nil
}
