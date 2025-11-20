package materials

import (
	"context"

	"kgplatform-backend/api/materials/v1"
)

// DeleteMaterial 删除材料
func (c *ControllerV1) DeleteMaterial(ctx context.Context, req *v1.DeleteMaterialReq) (res *v1.DeleteMaterialRes, err error) {
	err = c.materials.DeleteMaterial(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteMaterialRes{
		Success: true,
	}, nil
}
