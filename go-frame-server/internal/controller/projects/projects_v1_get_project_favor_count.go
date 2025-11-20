package projects

import (
	"context"
	v1 "kgplatform-backend/api/projects/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetProjectFavorCount(ctx context.Context, req *v1.GetProjectFavorCountReq) (res *v1.GetProjectFavorCountRes, err error) {
	// 检查项目是否存在且可见
	var project struct {
		Visibility int `json:"visibility"`
	}

	err = g.Model("projects").
		Where("id", req.ProjectId).
		Fields("visibility").
		Scan(&project)

	if err != nil {
		return nil, err
	}

	// 如果项目不可见，返回错误
	if project.Visibility != 1 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "项目不可见或不存在")
	}

	// 查询项目的收藏数
	var count int
	count, err = g.Model("views").
		Where("project_id", req.ProjectId).
		Count()

	if err != nil {
		return nil, err
	}

	return &v1.GetProjectFavorCountRes{
		FavorCount: count,
	}, nil
}
