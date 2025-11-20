// 方法：ControllerV1.CreateModel
package models

import (
	"context"

	"kgplatform-backend/api/models/v1"
	modelsLogic "kgplatform-backend/internal/logic/models"
)

func (c *ControllerV1) CreateModel(ctx context.Context, req *v1.CreateModelReq) (res *v1.CreateModelRes, err error) {
	return modelsLogic.CreateModel(ctx, req)
}
