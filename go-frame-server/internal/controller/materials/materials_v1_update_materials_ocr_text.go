package materials

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/materials"

	"kgplatform-backend/api/materials/v1"
)

func (c *ControllerV1) UpdateMaterialsOCRText(ctx context.Context, req *v1.UpdateMaterialsOCRTextReq) (res *v1.UpdateMaterialsOCRTextRes, err error) {
	materialsLogic := materials.New()
	data := g.Map{
		"text_url": req.FilePath,
	}
	err = materialsLogic.UpdateMaterialByGMap(ctx, req.Id, data)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateMaterialsOCRTextRes{}, nil
}
