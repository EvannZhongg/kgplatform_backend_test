// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package materials

import (
	"context"

	"kgplatform-backend/api/materials/v1"
)

type IMaterialsV1 interface {
	CreateMaterial(ctx context.Context, req *v1.CreateMaterialReq) (res *v1.CreateMaterialRes, err error)
	GetMaterial(ctx context.Context, req *v1.GetMaterialReq) (res *v1.GetMaterialRes, err error)
	UpdateMaterial(ctx context.Context, req *v1.UpdateMaterialReq) (res *v1.UpdateMaterialRes, err error)
	DeleteMaterial(ctx context.Context, req *v1.DeleteMaterialReq) (res *v1.DeleteMaterialRes, err error)
	ListMaterial(ctx context.Context, req *v1.ListMaterialReq) (res *v1.ListMaterialRes, err error)
	UpdateMaterialsOCRText(ctx context.Context, req *v1.UpdateMaterialsOCRTextReq) (res *v1.UpdateMaterialsOCRTextRes, err error)
	SwitchMaterialStatus(ctx context.Context, req *v1.SwitchMaterialStatusReq) (res *v1.SwitchMaterialStatusRes, err error)
}
