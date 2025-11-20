// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package models

import (
	"context"

	"kgplatform-backend/api/models/v1"
)

type IModelsV1 interface {
	CreateModel(ctx context.Context, req *v1.CreateModelReq) (res *v1.CreateModelRes, err error)
	ListModels(ctx context.Context, req *v1.ListModelsReq) (res *v1.ListModelsRes, err error)
}
