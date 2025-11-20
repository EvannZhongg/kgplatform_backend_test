package materials

import (
	"context"
	v1 "kgplatform-backend/api/materials/v1"
	materialsLogic "kgplatform-backend/internal/logic/materials"
)

func (c *ControllerV1) CreateMaterial(ctx context.Context, req *v1.CreateMaterialReq) (res *v1.CreateMaterialRes, err error) {
	result, err := c.materials.CreateMaterial(ctx, &materialsLogic.CreateMaterialInput{
		FilePath:  req.FilePath,
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateMaterialRes{
		MaterialId: result.MaterialId,
		Status:     result.Status,
	}, nil
}

func (c *ControllerV1) GetMaterial(ctx context.Context, req *v1.GetMaterialReq) (res *v1.GetMaterialRes, err error) {
	result, err := c.materials.GetMaterial(ctx, &materialsLogic.GetMaterialInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetMaterialRes{
		Id:        result.Id,
		FileUrl:   result.FileUrl,
		OcrUrl:    result.OcrUrl,
		TripleUrl: result.TripleUrl,
		ProjectId: result.ProjectId,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}
