package projects

import (
	"context"
	v1 "kgplatform-backend/api/projects/v1"
	"kgplatform-backend/internal/logic/projects"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetProjectViewCount(ctx context.Context, req *v1.GetProjectViewCountReq) (res *v1.GetProjectViewCountRes, err error) {
	projectLogic := projects.NewProjects()

	// 检查项目可见性
	visibility, err := g.Model("projects").
		Where("id", req.ProjectId).
		Fields("visibility").
		Value()

	if err != nil {
		return nil, err
	}

	if visibility.Int() != 1 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "项目不可见或不存在")
	}

	// 获取浏览量
	viewCount, err := projectLogic.GetViewCount(ctx, req.ProjectId)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "获取浏览量失败")
	}

	return &v1.GetProjectViewCountRes{
		ViewCount: viewCount,
	}, nil
}
