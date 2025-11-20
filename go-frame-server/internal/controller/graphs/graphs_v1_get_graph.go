package graphs

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/graphs"

	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/graphs/v1"
)

func (c *ControllerV1) GetGraph(ctx context.Context, req *v1.GetGraphReq) (res *v1.GetGraphRes, err error) {
	graphLogic := graphs.New()
	graphOutput, err := graphLogic.GetGraph(ctx, &graphs.GetGraphInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		g.Log().Errorf(ctx, "获取图谱失败: %v", err)
		return nil, gerror.New("获取图谱失败")
	}
	if graphOutput == nil {
		return nil, gerror.New("该项目的图谱暂未生成")
	}
	return &v1.GetGraphRes{
		Id:    graphOutput.Id,
		Nodes: graphOutput.Nodes,
		Edges: graphOutput.Edges,
	}, nil
}
