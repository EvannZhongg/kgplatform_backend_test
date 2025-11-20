// 方法：ControllerV1.ListModels
package models

import (
	"context"

	"kgplatform-backend/api/models/v1"
	modelsLogic "kgplatform-backend/internal/logic/models"
)

func (c *ControllerV1) ListModels(ctx context.Context, req *v1.ListModelsReq) (res *v1.ListModelsRes, err error) {
	return modelsLogic.ListModels(ctx, req)
}
