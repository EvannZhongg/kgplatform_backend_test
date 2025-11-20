package materials

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/materials"

	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/materials/v1"
)

func (c *ControllerV1) SwitchMaterialStatus(ctx context.Context, req *v1.SwitchMaterialStatusReq) (res *v1.SwitchMaterialStatusRes, err error) {
	materialsLogic := materials.New()
	material, err := materialsLogic.GetMaterial(ctx, &materials.GetMaterialInput{
		Id: req.Id,
	})
	if err != nil {
		g.Log().Errorf(ctx, "获取材料详情失败: %v", err)
		return nil, gerror.New("获取材料详情失败")
	}
	if material == nil {
		return nil, gerror.New("材料不存在")
	}
	enable := 0
	if material.Enable == 1 {
		enable = 0
	} else {
		enable = 1
	}
	err = materialsLogic.UpdateMaterialByGMap(ctx, req.Id, g.Map{
		"enable": enable,
	})
	if err != nil {
		g.Log().Errorf(ctx, "更新材料失败: %v", err)
		return nil, gerror.New("更新材料失败")
	}
	return &v1.SwitchMaterialStatusRes{}, nil
}
