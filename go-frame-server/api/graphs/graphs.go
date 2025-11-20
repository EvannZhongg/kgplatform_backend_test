// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package graphs

import (
	"context"

	"kgplatform-backend/api/graphs/v1"
)

type IGraphsV1 interface {
	GetGraph(ctx context.Context, req *v1.GetGraphReq) (res *v1.GetGraphRes, err error)
}
