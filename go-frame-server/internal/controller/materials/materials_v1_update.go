package materials

import (
	"context"

	"kgplatform-backend/api/materials/v1"
)

// UpdateMaterial 更新材料
func (c *ControllerV1) UpdateMaterial(ctx context.Context, req *v1.UpdateMaterialReq) (res *v1.UpdateMaterialRes, err error) {
	err = c.materials.UpdateMaterial(ctx, req.Id, req.FilePath, req.OcrPath, req.TriplePath, req.ProjectId)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateMaterialRes{
		Success: true,
	}, nil
}
