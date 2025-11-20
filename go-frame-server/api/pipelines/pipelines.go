// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pipelines

import (
	"context"

	"kgplatform-backend/api/pipelines/v1"
)

type IPipelinesV1 interface {
	CreatePipeline(ctx context.Context, req *v1.CreatePipelineReq) (res *v1.CreatePipelineRes, err error)
	GetPipeline(ctx context.Context, req *v1.GetPipelineReq) (res *v1.GetPipelineRes, err error)
}
