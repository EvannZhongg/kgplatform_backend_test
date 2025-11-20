package projects

import (
	"context"
	v1 "kgplatform-backend/api/projects/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetProjectPurchaseCount(ctx context.Context, req *v1.GetProjectPurchaseCountReq) (res *v1.GetProjectPurchaseCountRes, err error) {
	// 检查项目是否存在且可见
	var project struct {
		PurchaseCount int `json:"purchase_count"`
		Visibility    int `json:"visibility"`
	}

	err = g.Model("projects").
		Where("id", req.ProjectId).
		Fields("purchase_count, visibility").
		Scan(&project)

	if err != nil {
		return nil, err
	}

	// 如果项目不可见，返回错误
	if project.Visibility != 1 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "项目不可见或不存在")
	}

	return &v1.GetProjectPurchaseCountRes{
		PurchaseCount: project.PurchaseCount,
	}, nil
}
