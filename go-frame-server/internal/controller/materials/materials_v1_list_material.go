package materials

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"kgplatform-backend/internal/convert"
	"kgplatform-backend/internal/logic/materials"

	"kgplatform-backend/api/materials/v1"
)

func (c *ControllerV1) ListMaterial(ctx context.Context, req *v1.ListMaterialReq) (res *v1.ListMaterialRes, err error) {
	materialLogic := materials.New()
	output, err := materialLogic.ListMaterial(ctx, &materials.ListMaterialInput{
		Page:      req.Page,
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, gerror.New("获取材料列表失败")
	}
	materialDetailList, err := convert.ConvertMaterialListToDetailList(ctx, output.Materials)
	if err != nil {
		return nil, gerror.New("转换材料列表失败")
	}
	return &v1.ListMaterialRes{
		Material: materialDetailList,
		Total:    output.Total,
	}, nil
}
